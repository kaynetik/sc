package mini_oracle

import (
	"encoding/json"
	"fmt"
	abci "github.com/cometbft/cometbft/abci/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sedaprotocol/seda-chain/x/mini-oracle/types"
	"io"
	"net/http"
)

type ProposalHandler struct {
	txVerifier baseapp.ProposalTxVerifier
	txSelector baseapp.TxSelector
	fetchPrice func() (float64, error)
}

func NewDefaultProposalHandler(txVerifier baseapp.ProposalTxVerifier) *ProposalHandler {
	return &ProposalHandler{
		txVerifier: txVerifier,
		txSelector: baseapp.NewDefaultTxSelector(),
		fetchPrice: fetchBitcoinPrice,
	}
}

type CoinGeckoResponse struct {
	Bitcoin struct {
		USD float64 `json:"usd"`
	} `json:"bitcoin"`
}

func (h *ProposalHandler) PrepareProposalHandler(txConfig client.TxConfig) sdk.PrepareProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestPrepareProposal) (*abci.ResponsePrepareProposal, error) {
		price, err := h.fetchPrice()
		if err != nil {
			return nil, fmt.Errorf("failed to fetch Bitcoin price: %w", err)
		}

		msg := &types.MsgPriceReport{
			Price: int32(price * 100), // I presume you defined the interface of storing price as an integer to avoid floating point issues.
		}

		txBuilder := txConfig.NewTxBuilder()
		err = txBuilder.SetMsgs(msg)
		if err != nil {
			return nil, fmt.Errorf("failed to set message in transaction: %w", err)
		}

		txBytes, err := txConfig.TxEncoder()(txBuilder.GetTx())
		if err != nil {
			return nil, fmt.Errorf("failed to encode transaction: %w", err)
		}

		// Add the transaction to the proposal
		return &abci.ResponsePrepareProposal{Txs: [][]byte{txBytes}}, nil
	}
}

const geckoSimpleAPI = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"

func fetchBitcoinPrice() (float64, error) {
	resp, err := http.Get(geckoSimpleAPI)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var response CoinGeckoResponse
	err = json.Unmarshal(body, &response)
	if err != nil {
		return 0, err
	}

	return response.Bitcoin.USD, nil
}

func (h *ProposalHandler) ProcessProposalHandler(txConfig client.TxConfig) sdk.ProcessProposalHandler {
	return func(ctx sdk.Context, req *abci.RequestProcessProposal) (*abci.ResponseProcessProposal, error) {
		if len(req.Txs) == 0 {
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, fmt.Errorf("no transactions in proposal")
		}

		actualPrice, err := h.fetchPrice()
		if err != nil {
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, fmt.Errorf("failed to fetch Bitcoin price: %w", err)
		}

		tx, err := txConfig.TxDecoder()(req.Txs[0])
		if err != nil {
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, fmt.Errorf("failed to decode transaction: %w", err)
		}

		msg, ok := tx.GetMsgs()[0].(*types.MsgPriceReport)
		if !ok {
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, fmt.Errorf("invalid message type")
		}

		// Compare the reported price to the actual price
		reportedPrice := float64(msg.Price) / 100
		if !isWithinTolerance(actualPrice, reportedPrice) {
			return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_REJECT}, fmt.Errorf("reported price is not within tolerance")
		}

		return &abci.ResponseProcessProposal{Status: abci.ResponseProcessProposal_ACCEPT}, nil
	}
}

// Helper function to check if two prices are within 5% tolerance
func isWithinTolerance(actualPrice, reportedPrice float64) bool {
	diff := actualPrice - reportedPrice
	if diff < 0 {
		diff = -diff
	}

	tolerance := actualPrice * 0.05

	return diff <= tolerance
}
