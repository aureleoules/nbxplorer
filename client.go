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
