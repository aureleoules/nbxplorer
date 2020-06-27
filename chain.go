package nbxplorer

import (
	"strconv"
)

// ChainStatus struct
type ChainStatus struct {
	BitcoinStatus struct {
		Blocks               int     `json:"blocks"`
		Headers              int     `json:"headers"`
		VerificationProgress float64 `json:"verificationProgress"`
		IsSync               bool    `json:"isSynched"`
		IncrementalRelayFee  float64 `json:"incrementalRelayFee"`
		MinRelayTxFee        float64 `json:"minRelayTxFee"`
		Capabilities         struct {
			CanScanTxOutSet            bool `json:"canSupportTxoutSet"`
			CanSupportSegwit           bool `json:"canSupportSegwit"`
			CanSupportTransactionCheck bool `json:"canSupportTransactionCheck"`
		} `json:"capabilities"`
	} `json:"bitcoinStatus"`
	RepositoryPingTime   float64  `json:"repositoryPingTime"`
	IsFullySync          bool     `json:"isfullySynched"`
	ChainHeight          int      `json:"chainHeight"`
	SyncHeight           int      `json:"syncHeight"`
	NetworkType          string   `json:"networkType"`
	CryptoCode           string   `json:"cryptoCode"`
	InstanceName         string   `json:"instanceName"`
	SupportedCryptoCodes []string `json:"supportedCryptoCodes"`
	Version              string   `json:"version"`
}

// GetStatus of node
func (c *Client) GetStatus() (ChainStatus, error) {
	var status ChainStatus
	var r ErrorResponse
	_, err := c.httpClient.R().SetError(&r).SetResult(&status).Get("/status")
	return status, err
}

// FeeRate struct
type FeeRate struct {
	FeeRate    float64 `json:"feeRate"`
	BlockCount int     `json:"blockCount"`
}

// GetFeeRate
func (c *Client) GetFeeRate(blockCount int) (FeeRate, error) {
	var feeRate FeeRate
	var r ErrorResponse
	_, err := c.httpClient.R().SetError(&r).SetResult(&feeRate).Get("/fees/" + strconv.Itoa(blockCount))
	return feeRate, err
}
