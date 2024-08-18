package mini_oracle

//
//import (
//	"bytes"
//	"cosmossdk.io/log"
//	storetypes "cosmossdk.io/store/types"
//	"errors"
//	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
//	abci "github.com/cometbft/cometbft/abci/types"
//	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
//	sdkintegration "github.com/cosmos/cosmos-sdk/testutil/integration"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
//	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
//	sdkstakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
//	"github.com/golang/mock/gomock"
//	"github.com/rs/zerolog"
//	"github.com/sedaprotocol/seda-chain/x/mini-oracle/types"
//	wasmstoragetypes "github.com/sedaprotocol/seda-chain/x/wasm-storage/types"
//	"github.com/stretchr/testify/require"
//	"testing"
//	"time"
//)
//
//func TestPrepareProposalHandler(t *testing.T) {
//	ctrl := gomock.NewController(t)
//	defer ctrl.Finish()
//
//	mockTxConfig := NewMockTxConfig(ctrl)
//	mockTxBuilder := NewMockTxBuilder(ctrl)
//	mockTx := NewMockTx(ctrl) //🤦‍♂️ At this point I realized that I'll also have to understand mocking of the signing process...
//
//	mockEncoder := func(tx sdk.Tx) ([]byte, error) {
//		return []byte("mock-tx-bytes"), nil
//	}
//
//	// Note: Used from one of your integration tests.
//	keys := storetypes.NewKVStoreKeys(
//		authtypes.StoreKey, banktypes.StoreKey, sdkstakingtypes.StoreKey, wasmstoragetypes.StoreKey, wasmtypes.StoreKey,
//	)
//	buf := &bytes.Buffer{}
//	logger := log.NewLogger(buf, log.LevelOption(zerolog.DebugLevel))
//	cms := sdkintegration.CreateMultiStore(keys, logger)
//	ctx := sdk.NewContext(cms, cmtproto.Header{Time: time.Now().UTC()}, true, logger)
//
//	cases := []struct {
//		name          string
//		mockPrice     float64
//		mockPriceErr  error
//		expectTxBytes bool
//		expectError   bool
//	}{
//		{
//			name:          "Successful case",
//			mockPrice:     50000.0,
//			mockPriceErr:  nil,
//			expectTxBytes: true,
//			expectError:   false,
//		},
//		{
//			name:          "API error case",
//			mockPrice:     0.0,
//			mockPriceErr:  errors.New("failed to fetch Bitcoin price"),
//			expectTxBytes: false,
//			expectError:   true,
//		},
//		{
//			name:          "Price within tolerance",
//			mockPrice:     100.0,
//			mockPriceErr:  nil,
//			expectTxBytes: true,
//			expectError:   false,
//		},
//		{
//			name:          "Price out of tolerance",
//			mockPrice:     100.0,
//			mockPriceErr:  nil,
//			expectTxBytes: true,
//			expectError:   false,
//		},
//		{
//			name:          "No transactions in proposal",
//			mockPrice:     100.0,
//			mockPriceErr:  nil,
//			expectTxBytes: false,
//			expectError:   true,
//		},
//	}
//
//	for _, tc := range cases {
//		t.Run(tc.name, func(t *testing.T) {
//			mockTxConfig.EXPECT().NewTxBuilder().Return(mockTxBuilder)
//			mockTxConfig.EXPECT().TxEncoder().Return(mockEncoder)
//
//			if tc.expectTxBytes {
//				mockTxBuilder.EXPECT().SetMsgs(gomock.Any()).Return(nil)
//				mockTx.EXPECT().GetMsgs().Return([]sdk.Msg{&types.MsgPriceReport{
//					Price: int32(tc.mockPrice * 100),
//				}})
//				mockTxBuilder.EXPECT().GetTx().Return(mockTx)
//			}
//
//			handler := ProposalHandler{
//				fetchPrice: func() (float64, error) {
//					return tc.mockPrice, tc.mockPriceErr
//				},
//			}
//
//			req := &abci.RequestPrepareProposal{
//				MaxTxBytes: 1000,
//			}
//
//			resp, err := handler.PrepareProposalHandler(mockTxConfig)(ctx, req)
//
//			if tc.expectError {
//				require.Error(t, err)
//			} else {
//				require.NoError(t, err)
//				require.NotNil(t, resp)
//
//				if tc.expectTxBytes {
//					require.Len(t, resp.Txs, 1)
//					require.Equal(t, []byte("mock-tx-bytes"), resp.Txs[0])
//				} else {
//					require.Len(t, resp.Txs, 0)
//				}
//			}
//		})
//	}
//}
