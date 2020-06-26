package nbxplorer

import (
	"testing"

	"github.com/davecgh/go-spew/spew"
	"github.com/stretchr/testify/assert"
)

func TestGetEventStream(t *testing.T) {
	events, err := c.GetEventStream(0, false, nil)
	assert.Nil(t, err)

	spew.Dump(events)
}
func TestGetRecentEventStream(t *testing.T) {
	events, err := c.GetRecentEventStream(1)
	assert.Nil(t, err)
	spew.Dump(events)
}
