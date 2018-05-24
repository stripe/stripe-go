package source

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestSourceGet(t *testing.T) {
	source, err := Get("src_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceNew(t *testing.T) {
	source, err := New(&stripe.SourceObjectParams{
		Type:     stripe.String("bitcoin"),
		Amount:   stripe.Int64(1000),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Flow:     stripe.String(string(stripe.SourceFlowReceiver)),
		Owner: &stripe.SourceOwnerParams{
			Email: stripe.String("jenny.rosen@example.com"),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceUpdate(t *testing.T) {
	source, err := Update("src_123", &stripe.SourceObjectParams{
		Owner: &stripe.SourceOwnerParams{
			Email: stripe.String("jenny.rosen@example.com"),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceDetach(t *testing.T) {
	source, err := Detach("src_123", &stripe.SourceObjectDetachParams{
		Customer: stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceSharing(t *testing.T) {
	params := &stripe.SourceObjectParams{
		Type:           stripe.String("card"),
		Customer:       stripe.String("cus_123"),
		OriginalSource: stripe.String("src_123"),
		Usage:          stripe.String(string(stripe.SourceUsageReusable)),
	}
	params.SetStripeAccount("acct_123")
	source, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, source)
}
