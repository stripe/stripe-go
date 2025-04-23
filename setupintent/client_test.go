package setupintent

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestSetupIntentCancel(t *testing.T) {
	intent, err := Cancel("seti_123", &stripe.SetupIntentCancelParams{
		CancellationReason: stripe.String(string(stripe.SetupIntentCancellationReasonRequestedByCustomer)),
	})
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}

func TestSetupIntentConfirm(t *testing.T) {
	intent, err := Confirm("seti_123", &stripe.SetupIntentConfirmParams{
		ReturnURL: stripe.String("https://stripe.com/return"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}

func TestSetupIntentGet(t *testing.T) {
	intent, err := Get("seti_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}

func TestSetupIntentList(t *testing.T) {
	i := List(&stripe.SetupIntentListParams{})

	// Verify that we can get at least one setup intent
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.SetupIntent())
	assert.NotNil(t, i.SetupIntentList())
}

func TestSetupIntentNew(t *testing.T) {
	intent, err := New(&stripe.SetupIntentParams{
		Customer: stripe.String("cus_123"),
		PaymentMethodTypes: stripe.StringSlice([]string{
			"card",
		}),
	})
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}

func TestSetupIntentUpdate(t *testing.T) {
	intent, err := Update("seti_123", &stripe.SetupIntentParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	})
	assert.Nil(t, err)
	assert.NotNil(t, intent)
}
