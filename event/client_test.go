package event

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestEventGet(t *testing.T) {
	event, err := Get("evt_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, event)
}

func TestEventList(t *testing.T) {
	i := List(&stripe.EventListParams{})

	// Verify that we can get at least one event
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.Event())
	assert.NotNil(t, i.EventList())
}
