package paymentsource

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestSourceGet(t *testing.T) {
	source, err := Get("card_123", &stripe.PaymentSourceParams{
		Customer: stripe.String("cus_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceList(t *testing.T) {
	i := List(&stripe.PaymentSourceListParams{
		Customer: stripe.String("cus_123"),
	})

	// Verify that we can get at least one source
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.PaymentSource())
	assert.NotNil(t, i.PaymentSourceList())
}

func TestSourceNew(t *testing.T) {
	params := &stripe.PaymentSourceParams{
		Customer: stripe.String("cus_123"),
		Source:   &stripe.PaymentSourceSourceParams{Token: stripe.String("tok_123")},
	}

	source, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceUpdate(t *testing.T) {
	params := &stripe.PaymentSourceParams{
		Customer: stripe.String("cus_123"),
	}
	params.AddMetadata("key", "value")

	source, err := Update("card_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceVerify(t *testing.T) {
	source, err := Verify("ba_123", &stripe.PaymentSourceVerifyParams{
		Customer: stripe.String("cus_123"),
		Amounts:  [2]int64{32, 45},
	})
	assert.Nil(t, err)
	assert.NotNil(t, source)
}

func TestSourceObjectVerify(t *testing.T) {
	source, err := Verify("src_123", &stripe.PaymentSourceVerifyParams{
		Values: stripe.StringSlice([]string{
			"32",
			"45",
		}),
	})
	assert.Nil(t, err)
	assert.NotNil(t, source)
}
