package nbxplorer

import "github.com/go-resty/resty"

// Chain enum
type Chain string

const (
	// BTC : Bitcoin
	BTC Chain = "btc"
	// LBTC : Litecoin
	LBTC Chain = "lbtc"
)

// Client struct
type Client struct {
	Chain     Chain
	userAgent string

	*resty.Client
}

// NewClient constructor
func NewClient(host string, chain Chain) *Client {
	return &Client{
		userAgent: "go-nbxplorer",
		Chain:     chain,
		Client:    resty.New().SetHostURL("http://" + host + "/v1/cryptos/" + string(chain)).EnableTrace(),
	}
}

// ChainStatus struct
type ChainStatus struct {
	BitcoinStatus struct {
		Blocks               int     `json:"blocks"`
		Headers              int     `json:"headers"`
		VerificationProgress float64 `json:"verificationProgress"`
		IsSync               bool    `json:"isSynched"`
		incrementalRelayFee  int     `json:"incrementalRelayFee"`
		MinRelayTxFee        int     `json:"minRelayTxFee"`
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
	_, err := c.R().SetError(&r).SetResult(&status).Get("/status")
	return status, err
}
