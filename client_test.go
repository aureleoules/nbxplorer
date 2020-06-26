package nbxplorer

import "testing"

var c *Client

func TestMain(t *testing.M) {
	c = NewClient("localhost:7000", BTC)

	t.Run()
}
