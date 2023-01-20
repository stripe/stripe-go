package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestPaymentIntentNextAction_UnmarshalJSON(t *testing.T) {
	actionData := map[string]interface{}{
		"redirect_to_url": map[string]interface{}{
			"return_url": "https://stripe.com/return",
			"url":        "https://stripe.com",
		},
		"type": "redirect_to_url",
	}

	bytes, err := json.Marshal(&actionData)
	assert.NoError(t, err)

	var action PaymentIntentNextAction
	err = json.Unmarshal(bytes, &action)
	assert.NoError(t, err)

	assert.Equal(t, PaymentIntentNextActionTypeRedirectToURL, action.Type)
	assert.Equal(t, "https://stripe.com", action.RedirectToURL.URL)
	assert.Equal(t, "https://stripe.com/return", action.RedirectToURL.ReturnURL)
}

func TestPaymentIntent_UnmarshalJSON(t *testing.T) {
	intentData := map[string]interface{}{
		"id":     "pi_123",
		"object": "payment_intent",
		"latest_charge": map[string]interface{}{
			"id":     "ch_123",
			"object": "charge",
		},
		"payment_method_types": []interface{}{
			"card",
		},
	}

	bytes, err := json.Marshal(&intentData)
	assert.NoError(t, err)

	var intent PaymentIntent
	err = json.Unmarshal(bytes, &intent)
	assert.NoError(t, err)

	assert.Equal(t, "pi_123", intent.ID)

	assert.Equal(t, "ch_123", intent.LatestCharge.ID)
	assert.Equal(t, 1, len(intent.PaymentMethodTypes))
}
