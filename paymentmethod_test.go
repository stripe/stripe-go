package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestPaymentMethod_UnmarshalJSON(t *testing.T) {
	data := map[string]interface{}{
		"id": "pm_123",
		"sepa_debit": map[string]interface{}{
			"mandate": "mandate_123",
		},
	}

	bytes, err := json.Marshal(&data)
	assert.NoError(t, err)

	var paymentMethod PaymentMethod
	err = json.Unmarshal(bytes, &paymentMethod)
	assert.NoError(t, err)

	assert.Equal(t, "pm_123", paymentMethod.ID)
	assert.Equal(t, "mandate_123", paymentMethod.SepaDebit.Mandate)
}
