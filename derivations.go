package nbxplorer

import (
	"errors"
	"strconv"
)

// GetDerivationSchemeTransactions
func (c *Client) GetDerivationSchemeTransactions(derivationScheme string) ([]TransactionVerbose, error) {
	var txs []TransactionVerbose
	var r ErrorResponse
	_, err := c.httpClient.R().SetResult(&txs).SetError(&r).Get("/derivations/" + derivationScheme + "/transactions")
	return txs, err
}

// GetDerivationSchemeTransactions
func (c *Client) GetDerivationSchemeTransaction(derivationScheme string, txid string) (TransactionVerbose, error) {
	var tx TransactionVerbose
	var r ErrorResponse
	_, err := c.httpClient.R().SetResult(&tx).SetError(&r).Get("/derivations/" + derivationScheme + "/transactions/" + txid)
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
	req := c.httpClient.R().SetError(&r)

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

	_, err := c.httpClient.R().SetError(&r).SetResult(&balance).Get("/derivations/" + derivationScheme + "/balance")
	return balance, err
}

// ScriptPubKeyInfos struct
type ScriptPubKeyInfos struct {
	TrackedSource      string  `json:"trackedSource"`
	Feature            Feature `json:"feature"`
	DerivationStrategy string  `json:"derivationStrategy`
	KeyPath            string  `json:"keyPath"`
	ScriptPubKey       string  `json:"scriptPubKey"`
	Address            string  `json:"address"`
}

func (c *Client) GetScriptPubKeyInfos(derivationScheme string, script string) (ScriptPubKeyInfos, error) {
	var infos ScriptPubKeyInfos
	var r *ErrorResponse

	_, err := c.httpClient.R().SetResult(&infos).SetError(&r).Get("/derivations/" + derivationScheme + "/scripts/" + script)
	if r != nil {
		return infos, errors.New(r.Message)
	}

	return infos, err
}

// UTXO struct
type UTXO struct {
	Feature         string `json:"feature"`
	Outpoint        string `json:"outpoint"`
	Index           int    `json:"index"`
	TransactionHash string `json:"transactionHash"`
	ScriptPubKey    string `json:"scriptPubKey"`
	Value           int    `json:"value"`
	KeyPath         string `json:"keyPath"`
	Timestamp       int    `json:"timestamp"`
	Confirmations   int    `json:"confirmations"`
}

// UTXOInfos struct
type UTXOInfos struct {
	TrackedSource      string `json:"trackedSource"`
	DerivationStrategy string `json:"derivationStrategy"`
	CurrentHeight      int    `json:"currentHeight"`
	Unconfirmed        struct {
		UtxOs          []UTXO   `json:"utxOs"`
		SpentOutpoints []string `json:"spentOutpoints"`
		HasChanges     bool     `json:"hasChanges"`
	} `json:"unconfirmed"`
	Confirmed struct {
		UtxOs          []UTXO   `json:"utxOs"`
		SpentOutpoints []string `json:"spentOutpoints"`
		HasChanges     bool     `json:"hasChanges"`
	} `json:"confirmed"`
	HasChanges bool `json:"hasChanges"`
}

// GetDerivationSchemeUTXOs of derivation scheme
func (c *Client) GetDerivationSchemeUTXOs(derivationScheme string) (UTXOInfos, error) {
	var infos UTXOInfos
	var r *ErrorResponse

	_, err := c.httpClient.R().SetResult(&infos).SetError(&r).Get("/derivations/" + derivationScheme + "/utxos")
	if r != nil {
		return infos, errors.New(r.Message)
	}

	return infos, err
}

// ScanUTXOSet scans the UTXO Set for output belonging to your derivationScheme.
// In order to not consume too much RAM, NBXplorer splits the addresses to scan in several batch and scan the whole UTXO set sequentially.
// Three branches are scanned: 0/x, 1/x and x.
// If a UTXO in one branch get found at a specific x, then all addresses inferior to
// index x will be considered used and not proposed when fetching a new unused address.
func (c *Client) ScanUTXOSet(derivationScheme string, batchSize int, gapLimit int, from int) error {
	var infos UTXOInfos
	var r *ErrorResponse

	_, err := c.httpClient.R().SetResult(&infos).SetError(&r).Get("/derivations/" + derivationScheme + "/utxos/scan")
	if r != nil {
		return errors.New(r.Message)
	}

	return err
}

func (c *Client) AttachDerivationSchemeMetadata(derivationScheme string, key string, metaData interface{}) error {
	var r *ErrorResponse

	_, err := c.httpClient.R().SetError(&r).SetBody(metaData).Get("/derivations/" + derivationScheme + "/metadata/" + key)
	if r != nil {
		return errors.New(r.Message)
	}

	return err
}

func (c *Client) DetachDerivationSchemeMetadata(derivationScheme string, key string) error {
	var r *ErrorResponse

	_, err := c.httpClient.R().SetError(&r).Post("/derivations/" + derivationScheme + "/metadata/" + key)
	if r != nil {
		return errors.New(r.Message)
	}

	return err
}

func (c *Client) GetDerivationSchemeMetadata(derivationScheme string, key string) (interface{}, error) {
	var metadata interface{}
	var r *ErrorResponse

	_, err := c.httpClient.R().SetResult(&metadata).SetError(&r).Get("/derivations/" + derivationScheme + "/metadata/" + key)
	if r != nil {
		return nil, errors.New(r.Message)
	}

	return metadata, err
}

func (c *Client) PruneUTXOSet(derivationScheme string, daysToKeep int) (int, error) {
	var resp struct {
		TotalPruned int `json:"totalPruned"`
	}
	var r *ErrorResponse

	_, err := c.httpClient.R().SetResult(&resp).SetBody(map[string]string{
		"daysToKeep": strconv.Itoa(daysToKeep),
	}).SetError(&r).Post("/derivations/" + derivationScheme + "/prune")

	if r != nil {
		return 0, errors.New(r.Message)
	}

	return resp.TotalPruned, err
}

// CreateWalletOptions struct
type CreateWalletOptions struct {
	AccountNumber    int    `json:"accountNumber"`
	WordList         string `json:"wordList"`
	ExistingMnemonic string `json:"existingMnemonic"`
	WordCount        int    `json:"wordCount"`
	ScriptPubKeyType string `json:"scriptPubKeyType"`
	Passphrase       string `json:"passphrase"`
	ImportKeysToRPC  bool   `json:"importKeysToRPC"`
	SavePrivateKeys  bool   `json:"savePrivateKeys"`
}

// CreateWalletResponse struct
type CreateWalletResponse struct {
	Mnemonic         string `json:"mnemonic"`
	Passphrase       string `json:"passphrase"`
	WordList         string `json:"wordList"`
	WordCount        int    `json:"wordCount"`
	MasterHDKey      string `json:"masterHDKey"`
	AccountHDKey     string `json:"accountHDKey"`
	AccountKeyPath   string `json:"accountKeyPath"`
	DerivationScheme string `json:"derivationScheme"`
}

func (c *Client) CreateWallet(options CreateWalletOptions) (CreateWalletResponse, error) {
	var resp CreateWalletResponse
	var r *ErrorResponse

	_, err := c.httpClient.R().SetResult(&resp).SetBody(options).SetError(&r).Post("/derivations")

	if r != nil {
		return resp, errors.New(r.Message)
	}

	return resp, err
}
