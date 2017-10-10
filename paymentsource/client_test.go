package paymentsource

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestSourceGet(t *testing.T) {
	source, err := Get("card_123", &stripe.CustomerSourceParams{
		Customer: "cus_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceList(t *testing.T) {
	i := List(&stripe.SourceListParams{
		Customer: "cus_123",
	})

	// Verify that we can get at least one source
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.PaymentSource())
}

func TestSourceNew(t *testing.T) {
	params := &stripe.CustomerSourceParams{
		Customer: "cus_123",
	}
	params.SetSource("tok_123")

	source, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceUpdate(t *testing.T) {
	params := &stripe.CustomerSourceParams{
		Customer: "cus_123",
	}
	params.SetSource(&stripe.CardParams{
		Name: "Updated Name",
	})

	source, err := Update("card_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceVerify(t *testing.T) {
	source, err := Verify("ba_123", &stripe.SourceVerifyParams{
		Customer: "cus_123",
		Amounts:  [2]int64{32, 45},
	})
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceObjectVerify(t *testing.T) {
	source, err := Verify("src_123", &stripe.SourceVerifyParams{
		Values: []string{"32", "45"},
	})
	assert.Nil(t, err)
	assert.NotNil(t, source)
}
