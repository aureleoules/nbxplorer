package nbxplorer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetFeeRate(t *testing.T) {
	_, err := c.GetFeeRate(20)
	assert.Nil(t, err)
}
