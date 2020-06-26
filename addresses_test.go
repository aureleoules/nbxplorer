package nbxplorer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var masterPub = "xpub68c1FHkkqwxxYjuo4WvcDTv6nmVepJXN3Db9KTiuLPzdqRH9ddnVMB3zd1Tj5YDVD4NnC6ngvc2Tic3rGpZ7pkgQQmYcZ7N5Y3GiL6GXNiH"

func TestNewUnusedAddress(t *testing.T) {
	_, err := c.NewUnusedAddress(masterPub, Deposit, 0, false)
	assert.Nil(t, err)
}
