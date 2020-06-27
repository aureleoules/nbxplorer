package nbxplorer

import "github.com/go-resty/resty"

// Chain enum
type Chain string

const (
	BTC   Chain = "btc"
	LBTC  Chain = "lbtc"
	AGM   Chain = "agm"
	BTX   Chain = "btx"
	LTC   Chain = "ltc"
	Doge  Chain = "doge"
	BCH   Chain = "bch"
	GRS   Chain = "grs"
	BTG   Chain = "btg"
	Dash  Chain = "dash"
	TRC   Chain = "trc"
	Polis Chain = "polis"
	Mona  Chain = "mona"
	FTC   Chain = "ftc"
	UFO   Chain = "ufo"
	VIA   Chain = "via"
	XMCC  Chain = "xmcc"
	BGX   Chain = "bgx"
	COLX  Chain = "colx"
	QTUM  Chain = "qtum"
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
