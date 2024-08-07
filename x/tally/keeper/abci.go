package keeper

import (
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/sedaprotocol/seda-wasm-vm/tallyvm"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sedaprotocol/seda-chain/x/tally/types"
)

func (k Keeper) EndBlock(ctx sdk.Context) (err error) {
	// Use defer to prevent returning an error, which would cause
	// the chain to halt.
	defer func() {
		// Handle a panic.
		if r := recover(); r != nil {
			k.Logger(ctx).Error("recovered from panic in tally end block", "err", r)
		}
		// Handle an error.
		if err != nil {
			k.Logger(ctx).Error("error in tally end block", "err", err)
		}
		err = nil
	}()

	err = k.ProcessTallies(ctx)
	if err != nil {
		return
	}
	return
}

// ProcessTallies fetches from the core contract the list of requests
// to be tallied and then goes through it to filter and tally.
func (k Keeper) ProcessTallies(ctx sdk.Context) error {
	// Get core contract address.
	coreContract, err := k.wasmStorageKeeper.GetCoreContractAddr(ctx)
	if err != nil {
		return err
	}

	// Fetch tally-ready data requests.
	// TODO: Deal with offset and limits. (#313)
	queryRes, err := k.wasmViewKeeper.QuerySmart(ctx, coreContract, []byte(`{"get_data_requests_by_status":{"status": "tallying", "offset": 0, "limit": 100}}`))
	if err != nil {
		return err
	}
	if string(queryRes) == "[]" {
		return nil
	}

	k.Logger(ctx).Info("non-empty tally list - starting tally process")

	var tallyList []types.Request
	err = json.Unmarshal(queryRes, &tallyList)
	if err != nil {
		return err
	}

	// Loop through the list to apply filter, execute tally, and post
	// execution result.
	for _, req := range tallyList {
		// Construct barebone sudo message to be posted to the contract
		// here and populate its results fields after FilterAndTally.
		sudoMsg := types.Sudo{
			ID: req.ID,
			Result: types.DataResult{
				Version:        req.Version,
				ID:             req.ID,
				BlockHeight:    uint64(ctx.BlockHeight()),
				GasUsed:        "0", // TODO
				PaybackAddress: req.PaybackAddress,
				SedaPayload:    req.SedaPayload,
			},
		}

		vmRes, consensus, err := k.FilterAndTally(ctx, req)
		if err != nil {
			// Return with exit code 255 to signify that the tally VM
			// was not executed due to the error specified in the result
			// field.
			sudoMsg.ExitCode = 0xff
			sudoMsg.Result.ExitCode = 0xff
			sudoMsg.Result.Result = []byte(err.Error())
			sudoMsg.Result.Consensus = consensus
		} else {
			sudoMsg.ExitCode = byte(vmRes.ExitInfo.ExitCode)
			sudoMsg.Result.ExitCode = byte(vmRes.ExitInfo.ExitCode)
			sudoMsg.Result.Result = vmRes.Result
			sudoMsg.Result.Consensus = consensus
		}

		// Post results to the SEDA contract.
		msg, err := json.Marshal(struct {
			PostDataResult types.Sudo `json:"post_data_result"`
		}{
			PostDataResult: sudoMsg,
		})
		if err != nil {
			return err
		}

		k.Logger(ctx).Info(
			"posting execution results to SEDA contract",
			"request_id", req.ID,
			"execution_result", vmRes,
			"sudo_message", sudoMsg,
		)
		postRes, err := k.wasmKeeper.Sudo(ctx, coreContract, msg)
		if err != nil {
			return err
		}

		k.Logger(ctx).Info(
			"tally flow completed",
			"request_id", req.ID,
			"post_result", postRes,
		)
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeTallyCompletion,
				sdk.NewAttribute(types.AttributeDataRequestID, req.ID),
				sdk.NewAttribute(types.AttributeTypeConsensus, strconv.FormatBool(consensus)),
				sdk.NewAttribute(types.AttributeTallyVMStdOut, strings.Join(vmRes.Stdout, "\n")),
				sdk.NewAttribute(types.AttributeTallyVMStdErr, strings.Join(vmRes.Stderr, "\n")),
				sdk.NewAttribute(types.AttributeTallyExitCode, fmt.Sprintf("%02x", sudoMsg.ExitCode)),
			),
		)
	}

	return nil
}

// FilterAndTally applies filter and executes tally. It returns the
// tally VM result, consensus boolean, and error if applicable.
func (k Keeper) FilterAndTally(ctx sdk.Context, req types.Request) (tallyvm.VmResult, bool, error) {
	filter, err := base64.StdEncoding.DecodeString(req.ConsensusFilter)
	if err != nil {
		return tallyvm.VmResult{}, false, fmt.Errorf("failed to decode consensus filter: %w", err)
	}

	// Sort reveals.
	keys := make([]string, len(req.Reveals))
	i := 0
	for k := range req.Reveals {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	reveals := make([]types.RevealBody, len(req.Reveals))
	for i, k := range keys {
		reveals[i] = req.Reveals[k]
	}

	outliers, consensus, err := ApplyFilter(filter, reveals)
	if err != nil {
		return tallyvm.VmResult{}, false, fmt.Errorf("error while applying filter: %w", err)
	}

	tallyWasm, err := k.wasmStorageKeeper.GetDataRequestWasm(ctx, req.TallyBinaryID)
	if err != nil {
		return tallyvm.VmResult{}, false, err
	}
	tallyInputs, err := base64.StdEncoding.DecodeString(req.TallyInputs)
	if err != nil {
		return tallyvm.VmResult{}, false, fmt.Errorf("failed to decode tally inputs: %w", err)
	}

	args, err := tallyVMArg(tallyInputs, reveals, outliers)
	if err != nil {
		return tallyvm.VmResult{}, false, fmt.Errorf("failed to construct tally VM arguments: %w", err)
	}

	k.Logger(ctx).Info(
		"executing tally VM",
		"request_id", req.ID,
		"tally_wasm_hash", req.TallyBinaryID,
		"consensus", consensus,
		"arguments", args,
	)
	vmRes := tallyvm.ExecuteTallyVm(tallyWasm.Bytecode, args, map[string]string{
		"VM_MODE":   "tally",
		"CONSENSUS": fmt.Sprintf("%v", consensus),
	})
	return vmRes, consensus, nil
}

func tallyVMArg(inputArgs []byte, reveals []types.RevealBody, outliers []int) ([]string, error) {
	arg := []string{hex.EncodeToString(inputArgs)}

	r, err := json.Marshal(reveals)
	if err != nil {
		return nil, err
	}
	arg = append(arg, string(r))

	o, err := json.Marshal(outliers)
	if err != nil {
		return nil, err
	}
	arg = append(arg, string(o))

	return arg, err
}
