package nbxplorer

import "errors"

// GetDerivationSchemeTransactions
func (c *Client) GetDerivationSchemeTransactions(derivationScheme string) ([]TransactionVerbose, error) {
	var txs []TransactionVerbose
	var r ErrorResponse
	_, err := c.R().SetResult(&txs).SetError(&r).Get("/derivations/" + derivationScheme + "/transactions")
	return txs, err
}

// GetDerivationSchemeTransactions
func (c *Client) GetDerivationSchemeTransaction(derivationScheme string, txid string) (TransactionVerbose, error) {
	var tx TransactionVerbose
	var r ErrorResponse
	_, err := c.R().SetResult(&tx).SetError(&r).Get("/derivations/" + derivationScheme + "/transactions/" + txid)
	return tx, err
}

// DerivationOptions struct
type DerivationOptions struct {
	Feature      Feature `json:"feature"`
	MinAddresses *int    `json:"minAddresses"`
	MaxAddresses *int    `json:"maxAddresses"`
}

// TrackDerivationSchemeOptions struct
type TrackDerivationSchemeOptions struct {
	DerivationOptions DerivationOptions `json:"derivationOptions"`
	Wait              bool              `json:"wait"`
}

// TrackDerivationScheme
func (c *Client) TrackDerivationScheme(derivationScheme string, options *TrackDerivationSchemeOptions) error {
	var r ErrorResponse
	req := c.R().SetError(&r)

	if options != nil {
		req.SetBody(options)
	}
	resp, err := req.Post("/derivations/" + derivationScheme)
	if err != nil {
		return err
	}
	if len(resp.Body()) == 0 {
		return nil
	}

	return errors.New(r.Message)
}

// Balance struct
type Balance struct {
	Unconfirmed int64
	Confirmed   int64
	Total       int64
}

// GetCurrentBalance
func (c *Client) GetCurrentBalance(derivationScheme string) (Balance, error) {
	var balance Balance
	var r ErrorResponse

	_, err := c.R().SetError(&r).SetResult(&balance).Get("/derivations/" + derivationScheme + "/balance")
	return balance, err
}
