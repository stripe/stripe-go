package paymentintent

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestPaymentIntentCancel(t *testing.T) {
	intent, err := Cancel("pi_123", &stripe.PaymentIntentCancelParams{
		CancellationReason: stripe.String(string(stripe.PaymentIntentCancellationReasonRequestedByCustomer)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}

func TestPaymentIntentCapture(t *testing.T) {
	intent, err := Capture("pi_123", &stripe.PaymentIntentCaptureParams{
		AmountToCapture: stripe.Int64(123),
	})
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}

func TestPaymentIntentConfirm(t *testing.T) {
	intent, err := Confirm("pi_123", &stripe.PaymentIntentConfirmParams{
		ReturnURL:  stripe.String("https://stripe.com/return"),
		OffSession: stripe.Bool(true),
	})
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}

func TestPaymentIntentGet(t *testing.T) {
	intent, err := Get("pi_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}

func TestPaymentIntentList(t *testing.T) {
	i := List(&stripe.PaymentIntentListParams{})

	// Verify that we can get at least one payment intent
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.PaymentIntent())
	assert.NotNil(t, i.PaymentIntentList())
}

func TestPaymentIntentNew(t *testing.T) {
	intent, err := New(&stripe.PaymentIntentParams{
		Amount:   stripe.Int64(123),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	})
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}

func TestPaymentIntentUpdate(t *testing.T) {
	intent, err := Update("pi_123", &stripe.PaymentIntentParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}
