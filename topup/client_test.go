package topup

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestTopupCancel(t *testing.T) {
	topup, err := Cancel("tu_123", &stripe.TopupParams{})
	assert.Nil(t, err)
	assert.NotNil(t, topup)
}

func TestTopupGet(t *testing.T) {
	topup, err := Get("tu_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, topup)
}

func TestTopupNew(t *testing.T) {
	topup, err := New(&stripe.TopupParams{
		Amount:              stripe.Int64(123),
		Currency:            stripe.String(string(stripe.CurrencyUSD)),
		Source:              stripe.String("src_123"),
		Description:         stripe.String("creating a topup"),
		StatementDescriptor: stripe.String("topup"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, topup)
}

func TestTopupUpdate(t *testing.T) {
	params := &stripe.TopupParams{}
	params.AddMetadata("key", "value")
	topup, err := Update("tu_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, topup)
}
