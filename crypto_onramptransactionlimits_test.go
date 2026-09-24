package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCryptoOnrampTransactionLimits_UnmarshalJSON(t *testing.T) {
	var limits CryptoOnrampTransactionLimits
	err := json.Unmarshal([]byte(`{"limits":{"usd":{"card":{"standard":{"remaining":1250}}}}}`), &limits)
	assert.NoError(t, err)

	usd := limits.Limits["usd"].(map[string]interface{})
	card := usd["card"].(map[string]interface{})
	standard := card["standard"].(map[string]interface{})
	assert.Equal(t, float64(1250), standard["remaining"])
}
