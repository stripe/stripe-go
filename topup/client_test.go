package topup

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestTopupGet(t *testing.T) {
	topup, err := Get("tu_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, topup)
}

func TestTopupNew(t *testing.T) {
	topup, err := New(&stripe.TopupParams{
		Amount:    123,
		Currency:  "usd",
		Source:    &stripe.SourceParams{Token: "src_123"},
		Desc:      "creating a topup",
		Statement: "topup",
	})
	assert.Nil(t, err)
	assert.NotNil(t, topup)
}

func TestTopupNew_WithSetSource(t *testing.T) {
	params := stripe.TopupParams{
		Amount:    123,
		Currency:  "usd",
		Desc:      "creating a topup",
		Statement: "topup",
	}
	params.SetSource("src_123")

	topup, err := New(&params)
	assert.Nil(t, err)
	assert.NotNil(t, topup)
}

func TestTopupUpdate(t *testing.T) {
	params := &stripe.TopupParams{}
	params.AddMeta("key", "value")
	topup, err := Update("tu_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, topup)
}
