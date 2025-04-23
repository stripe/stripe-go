package paymentmethod

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestPaymentMethodAttach(t *testing.T) {
	params := &stripe.PaymentMethodAttachParams{
		Customer: stripe.String("cus_123"),
	}
	pm, err := Attach("pm_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, pm)
}

func TestPaymentMethodDetach(t *testing.T) {
	params := &stripe.PaymentMethodDetachParams{}
	pm, err := Detach("pm_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, pm)
}

func TestPaymentMethodGet(t *testing.T) {
	pm, err := Get("pm_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, pm)
}

func TestPaymentMethodList(t *testing.T) {
	params := &stripe.PaymentMethodListParams{
		Customer: stripe.String("cus_123"),
		Type:     stripe.String(string(stripe.PaymentMethodTypeCard)),
	}
	i := List(params)

	// Verify that we can get at least one pm
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.PaymentMethod())
	assert.NotNil(t, i.PaymentMethodList())
}

func TestPaymentMethodNew(t *testing.T) {
	pm, err := New(&stripe.PaymentMethodParams{
		Type: stripe.String(string(stripe.PaymentMethodTypeCard)),
		Card: &stripe.PaymentMethodCardParams{
			Token: stripe.String("tok_123"),
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, pm)
}

func TestPaymentMethodUpdate(t *testing.T) {
	params := &stripe.PaymentMethodParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	}
	pm, err := Update("pm_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, pm)
}
