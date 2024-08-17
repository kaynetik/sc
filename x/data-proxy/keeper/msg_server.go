package keeper

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/secp256k1"

	"cosmossdk.io/collections"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sedaprotocol/seda-chain/x/data-proxy/types"
)

type msgServer struct {
	Keeper
}

var _ types.MsgServer = msgServer{}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

func (m msgServer) RegisterDataProxy(goCtx context.Context, msg *types.MsgRegisterDataProxy) (*types.MsgRegisterDataProxyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.Validate(); err != nil {
		return nil, err
	}

	if _, err := sdk.AccAddressFromBech32(msg.AdminAddress); err != nil {
		return nil, types.ErrInvalidAddress.Wrapf("invalid admin address %s", err)
	}

	pubKeyBytes, err := hex.DecodeString(msg.PubKey)
	if err != nil {
		return nil, types.ErrInvalidHex.Wrap(err.Error())
	}

	signatureBytes, err := hex.DecodeString(msg.Signature)
	if err != nil {
		return nil, types.ErrInvalidHex.Wrap(err.Error())
	}

	found, err := m.DataProxyConfigs.Has(ctx, pubKeyBytes)
	if err != nil {
		return nil, err
	}
	if found {
		return nil, types.ErrAlreadyExists
	}

	feeBytes := []byte(msg.Fee.String())
	payoutAddressBytes := []byte(msg.PayoutAddress)
	memoBytes := []byte(msg.Memo)

	payload := make([]byte, 0, len(feeBytes)+len(payoutAddressBytes)+len(memoBytes))
	payload = append(payload, feeBytes...)
	payload = append(payload, payoutAddressBytes...)
	payload = append(payload, memoBytes...)

	if valid := secp256k1.VerifySignature(pubKeyBytes, crypto.Keccak256(payload), signatureBytes); !valid {
		return nil, types.ErrInvalidSignature
	}

	proxyConfig := types.ProxyConfig{
		PayoutAddress: msg.PayoutAddress,
		Fee:           msg.Fee,
		Memo:          msg.Memo,
		FeeUpdate:     nil,
		AdminAddress:  msg.AdminAddress,
	}

	err = proxyConfig.Validate()
	if err != nil {
		return nil, err
	}

	err = m.DataProxyConfigs.Set(ctx, pubKeyBytes, proxyConfig)
	if err != nil {
		return nil, err
	}

	return &types.MsgRegisterDataProxyResponse{}, nil
}

func (m msgServer) EditDataProxy(goCtx context.Context, msg *types.MsgEditDataProxy) (*types.MsgEditDataProxyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := msg.Validate(); err != nil {
		return nil, err
	}

	pubKeyBytes, err := hex.DecodeString(msg.PubKey)
	if err != nil {
		return nil, types.ErrInvalidHex.Wrap(err.Error())
	}

	proxyConfig, err := m.DataProxyConfigs.Get(ctx, pubKeyBytes)
	if err != nil {
		if errors.Is(err, collections.ErrNotFound) {
			return nil, types.ErrUnknownDataProxy.Wrapf("no data proxy registered for %s", msg.PubKey)
		}
		return nil, err
	}

	if msg.Sender != proxyConfig.AdminAddress {
		return nil, types.ErrUnauthorized
	}

	err = proxyConfig.UpdateBasic(msg.NewPayoutAddress, msg.NewMemo)
	if err != nil {
		return nil, err
	}

	// If there is no new fee we can terminate early
	if msg.NewFee == nil {
		err = m.DataProxyConfigs.Set(ctx, pubKeyBytes, proxyConfig)
		if err != nil {
			return nil, err
		}

		return &types.MsgEditDataProxyResponse{}, nil
	}

	return m.ProcessProxyFeeUpdate(ctx, pubKeyBytes, &proxyConfig, msg)
}

func (m msgServer) ProcessProxyFeeUpdate(ctx sdk.Context, pubKeyBytes []byte, proxyConfig *types.ProxyConfig, msg *types.MsgEditDataProxy) (*types.MsgEditDataProxyResponse, error) {
	params, err := m.Keeper.Params.Get(ctx)
	if err != nil {
		return nil, err
	}

	updateDelay := params.MinFeeUpdateDelay
	// Validate custom delay if passed
	if msg.FeeUpdateDelay != types.UseMinimumDelay {
		if msg.FeeUpdateDelay < params.MinFeeUpdateDelay {
			return nil, types.ErrInvalidDelay.Wrapf("minimum delay %d, got %d", params.MinFeeUpdateDelay, msg.FeeUpdateDelay)
		}

		updateDelay = msg.FeeUpdateDelay
	}

	// Determine update height
	updateHeight := ctx.BlockHeight() + int64(updateDelay)
	feeUpdate := &types.FeeUpdate{
		NewFee:       *msg.NewFee,
		UpdateHeight: updateHeight,
	}

	// Retain previous pending update so it can be removed
	previousPendingUpdate := proxyConfig.FeeUpdate
	proxyConfig.FeeUpdate = feeUpdate

	// Schedule new update
	err = m.Keeper.FeeUpdates.Set(ctx, collections.Join(updateHeight, pubKeyBytes))
	if err != nil {
		return nil, err
	}

	// Delete previous pending update, if applicable
	if previousPendingUpdate != nil {
		err = m.Keeper.FeeUpdates.Remove(ctx, collections.Join(previousPendingUpdate.UpdateHeight, pubKeyBytes))
		if err != nil {
			return nil, err
		}
	}

	err = m.DataProxyConfigs.Set(ctx, pubKeyBytes, *proxyConfig)
	if err != nil {
		return nil, err
	}

	return &types.MsgEditDataProxyResponse{
		UpdateHeight: updateHeight,
	}, nil
}

func (m msgServer) UpdateParams(goCtx context.Context, req *types.MsgUpdateParams) (*types.MsgUpdateParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if _, err := sdk.AccAddressFromBech32(req.Authority); err != nil {
		return nil, fmt.Errorf("invalid authority address: %s", err)
	}
	if m.GetAuthority() != req.Authority {
		return nil, fmt.Errorf("unauthorized authority; expected %s, got %s", m.GetAuthority(), req.Authority)
	}

	if err := req.Params.Validate(); err != nil {
		return nil, err
	}
	if err := m.Params.Set(ctx, req.Params); err != nil {
		return nil, err
	}

	return &types.MsgUpdateParamsResponse{}, nil
}
