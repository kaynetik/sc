package e2e

import (
	"bytes"
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/ethereum/go-ethereum/crypto"

	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"

	"github.com/sedaprotocol/seda-chain/x/wasm-storage/types"
)

func (s *IntegrationTestSuite) testWasmStorageStoreExecutorWasm() {
	proposalCounter++
	proposalID := proposalCounter

	senderAddress, err := s.chain.validators[0].keyInfo.GetAddress()
	s.Require().NoError(err)
	sender := senderAddress.String()

	bytecode, err := os.ReadFile(filepath.Join(localWasmDirPath, executorWasm))
	if err != nil {
		panic("failed to read data request Wasm file")
	}
	executorHashBytes := crypto.Keccak256(bytecode)
	if executorHashBytes == nil {
		panic("failed to compute hash")
	}
	executorHashStr := hex.EncodeToString(executorHashBytes)

	s.execWasmStorageStoreExecutor(s.chain, 0, executorWasm, "clean_title", "sustainable_summary", "data-request-executor", sender, standardFees.String(), false, proposalID)
	s.execGovVoteYes(s.chain, 0, sender, standardFees.String(), false, proposalID)

	s.Require().Eventually(
		func() bool {
			executorWasmRes, err := queryExecutorWasm(s.endpoint, executorHashStr)
			s.Require().NoError(err)
			s.Require().True(bytes.Equal(executorHashBytes, executorWasmRes.Wasm.Hash))

			wasms, err := queryExecutorWasms(s.endpoint)
			s.Require().NoError(err)

			return executorHashStr == wasms.List[0]
		},
		30*time.Second,
		5*time.Second,
	)
}

func (s *IntegrationTestSuite) execWasmStorageStoreExecutor(
	c *chain,
	valIdx int,
	executorWasm,
	title,
	summary,
	wasmType,
	from,
	fees string,
	expectErr bool,
	proposalID int,
	opt ...flagOption,
) {
	opt = append(opt, withKeyValue(flagFees, fees))
	opt = append(opt, withKeyValue(flagFrom, from))
	opt = append(opt, withKeyValue(flagWasmType, wasmType))
	opt = append(opt, withKeyValue(flagTitle, title))
	opt = append(opt, withKeyValue(flagSummary, summary))
	opt = append(opt, withKeyValue(flagDeposit, "10000000aseda"))
	opt = append(opt, withKeyValue(flagAuthority, authtypes.NewModuleAddress("gov").String()))
	opts := applyTxOptions(c.id, opt)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	wasmFilePath := filepath.Join(containerWasmDirPath, executorWasm)

	command := []string{
		binary,
		txCommand,
		types.ModuleName,
		"submit-proposal",
		"store-executor-wasm",
		wasmFilePath,
		"-y",
	}
	for flag, value := range opts {
		command = append(command, fmt.Sprintf("--%s=%v", flag, value))
	}

	s.T().Logf("proposing to store executor wasm %s on chain %s", wasmFilePath, c.id)

	s.executeTx(ctx, c, command, valIdx, s.expectErrExecValidation(c, valIdx, expectErr))

	s.Require().Eventually(
		func() bool {
			proposal, err := queryGovProposal(s.endpoint, proposalID)
			s.Require().NoError(err)

			return proposal.GetProposal().Status == govtypesv1.StatusVotingPeriod
		},
		15*time.Second,
		5*time.Second,
	)
}

func (s *IntegrationTestSuite) execGovVoteYes(
	c *chain,
	valIdx int,
	from,
	fees string,
	expectErr bool,
	proposalID int,
	opt ...flagOption,
) {
	opt = append(opt, withKeyValue(flagFees, fees))
	opt = append(opt, withKeyValue(flagFrom, from))
	opts := applyTxOptions(c.id, opt)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	command := []string{
		binary,
		txCommand,
		govtypes.ModuleName,
		"vote",
		strconv.Itoa(proposalID),
		"yes",
		"-y",
	}
	for flag, value := range opts {
		command = append(command, fmt.Sprintf("--%s=%v", flag, value))
	}

	s.T().Logf("voting yes to proposal %s on chain %s", strconv.Itoa(proposalID), c.id)

	s.executeTx(ctx, c, command, valIdx, s.expectErrExecValidation(c, valIdx, expectErr))

	s.Require().Eventually(
		func() bool {
			proposal, err := queryGovProposal(s.endpoint, proposalID)
			s.Require().NoError(err)

			return proposal.GetProposal().Status == govtypesv1.StatusPassed
		},
		30*time.Second,
		5*time.Second,
	)
}

func (s *IntegrationTestSuite) execGetSeedQuery(
	c *chain,
	valIdx int,
	coreContractAddr string,
	_ bool,
	opt ...flagOption,
) {
	opts := applyQueryOptions(c.id, opt)

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	// sedad query wasm contract-state smart $CORE_CONTRACT_ADDRESS '{"query_seed_request":{}}' --node $RPC_URL --output json
	command := []string{
		binary,
		queryCommand,
		wasmtypes.ModuleName,
		"contract-state",
		"smart",
		coreContractAddr,
		`{"query_seed_request":{}}`,
	}
	for flag, value := range opts {
		command = append(command, fmt.Sprintf("--%s=%v", flag, value))
	}

	s.T().Logf("getting seed from contract %s on chain %s", coreContractAddr, c.id)

	s.executeQuery(ctx, c, command, valIdx, s.validateGetSeedResponse(false))
}

func (s *IntegrationTestSuite) validateGetSeedResponse(expectEmpty bool) func([]byte, []byte) bool {
	return func(stdOut, stdErr []byte) bool {
		var getSeedResponse struct {
			Data struct {
				BlockHeight int    `json:"block_height"`
				Seed        string `json:"seed"`
			} `json:"data"`
		}

		err := json.Unmarshal(stdOut, &getSeedResponse)
		s.Require().NoError(err)
		if expectEmpty {
			s.Require().Empty(getSeedResponse.Data.Seed)
		} else {
			s.Require().NotEmpty(getSeedResponse.Data.Seed)
		}
		return true
	}
}
