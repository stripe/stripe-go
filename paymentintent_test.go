package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestPaymentIntentSourceAction_UnmarshalJSON(t *testing.T) {
	actionData := map[string]interface{}{
		"type": "authorize_with_url",
		"value": map[string]interface{}{
			"url": "https://stripe.com",
		},
	}

	bytes, err := json.Marshal(&actionData)
	assert.NoError(t, err)

	var action PaymentIntentSourceAction
	err = json.Unmarshal(bytes, &action)
	assert.NoError(t, err)

	assert.Equal(t, PaymentIntentNextActionAuthorizeWithURL, action.Type)
	assert.Equal(t, "https://stripe.com", action.Value.AuthorizeWithURL.URL)
}

func TestPaymentIntent_UnmarshalJSON(t *testing.T) {
	intentData := map[string]interface{}{
		"id":     "pi_123",
		"object": "payment_intent",
		"allowed_source_types": []interface{}{
			"card",
		},
		"charges": map[string]interface{}{
			"object":   "list",
			"has_more": true,
			"data": []map[string]interface{}{
				{
					"id":     "ch_123",
					"object": "charge",
				},
				{
					"id":     "ch_234",
					"object": "charge",
				},
			},
		},
	}

	bytes, err := json.Marshal(&intentData)
	assert.NoError(t, err)

	var intent PaymentIntent
	err = json.Unmarshal(bytes, &intent)
	assert.NoError(t, err)

	assert.Equal(t, "pi_123", intent.ID)

	assert.Equal(t, 1, len(intent.AllowedSourceTypes))
	assert.Equal(t, 2, len(intent.Charges.Data))
}
