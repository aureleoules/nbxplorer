package nbxplorer

import (
	"errors"
	"strconv"
)

// Output struct
type Output struct {
	KeyPath      string `json:"keyPath"`
	ScriptPubKey string `json:"scriptPubKey"`
	Index        int    `json:"index"`
	Value        int64  `json:"value"`
}

// Input struct
type Input struct{}

// TransactionVerbose struct
type TransactionVerbose struct {
	Transaction

	Outputs []Output `json:"outputs"`
	Inputs  []Input  `json:"inputs"`

	BalanceChange int64  `json:"balanceChange"`
	Replaceable   bool   `json:"replaceable"`
	Replacing     string `json:"replacing"`
	ReplaceBy     string `json:"replaceBy"`
}

// Transaction struct
type Transaction struct {
	Confirmations   int    `json:"confirmations"`
	BlockID         string `json:"blockId"`
	TransactionHash string `json:"transactionHash"`
	Transaction     string `json:"transaction"`
	Height          int    `json:"height"`
	Timestamp       int    `json:"timestamp"`
}

// GetTransaction
func (c *Client) GetTransaction(hash string, hex bool) (Transaction, error) {
	var tx Transaction
	var r ErrorResponse
	_, err := c.httpClient.R().
		SetResult(&tx).
		SetError(&r).
		SetQueryParam("includeTransaction", strconv.FormatBool(hex)).
		Get("/transactions/" + hash)

	return tx, err
}

// BroadcastTransaction
func (c *Client) BroadcastTransaction(tx []byte, testOnly bool) error {
	var r *ErrorResponse

	_, err := c.httpClient.R().
		SetError(&r).
		SetQueryParam("testMempoolAccept", strconv.FormatBool(testOnly)).
		Post("/transactions")

	if r != nil {
		return errors.New(r.Message)
	}

	return err
}

// TransactionID struct
type TransactionID struct {
	ID      string  `json:"transactionId"`
	BlockID *string `json:"blockId,omitempty"`
}

// RescanTransactions will index previous txs.
// NBXplorer does not rescan the whole blockchain when tracking a new derivation scheme.
// This means that if the derivation scheme already received UTXOs in the past,
// NBXplorer will not be aware of it and might reuse addresses already generated in the past,
// and will not show past transactions.
// By using this route, you can ask NBXplorer to rescan specific transactions found in the blockchain.
// This way, the transactions and the UTXOs present before tracking the derivation scheme will appear correctly.
// Only the transactionId is specified. Your node must run --txindex=1 for this to work
// Careful: A wrong blockId will corrupt the database.
func (c *Client) RescanTransactions(txs []TransactionID) error {
	var r *ErrorResponse

	_, err := c.httpClient.R().
		SetError(&r).
		SetBody(txs).
		Post("/rescan")

	if r != nil {
		return errors.New(r.Message)
	}

	return err
}
