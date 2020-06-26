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
	_, err := c.R().
		SetResult(&tx).
		SetError(&r).
		SetQueryParam("includeTransaction", strconv.FormatBool(hex)).
		Get("/transactions/" + hash)

	return tx, err
}

// BroadcastTransaction
func (c *Client) BroadcastTransaction(tx []byte, testOnly bool) error {
	var r *ErrorResponse

	_, err := c.R().
		SetError(&r).
		SetQueryParam("testMempoolAccept", strconv.FormatBool(testOnly)).
		Post("/transactions")

	if r != nil {
		return errors.New(r.Message)
	}

	return err
}
