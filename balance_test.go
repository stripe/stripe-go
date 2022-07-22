package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestBalance_Unmarshal(t *testing.T) {
	balanceData := map[string]interface{}{
		"available": []map[string]interface{}{
			{
				"amount":   111,
				"currency": "usd",
				"source_types": map[string]interface{}{
					"bank_account": 100,
					"card":         11,
				},
			},
			{
				"amount":   222,
				"currency": "eur",
				"source_types": map[string]interface{}{
					"card": 222,
				},
			},
		},
		"object": "balance",
		"pending": []map[string]interface{}{
			{
				"amount":   333,
				"currency": "usd",
				"source_types": map[string]interface{}{
					"bank_account": 333,
				},
			},
			{
				"amount":   444,
				"currency": "eur",
				"source_types": map[string]interface{}{
					"bank_account": 222,
					"card":         222,
				},
			},
		},
	}

	bytes, err := json.Marshal(&balanceData)
	assert.NoError(t, err)

	var balance Balance
	err = json.Unmarshal(bytes, &balance)
	assert.NoError(t, err)

	assert.Equal(t, int64(111), balance.Available[0].Amount)
	assert.Equal(t, CurrencyUSD, balance.Available[0].Currency)
	assert.Equal(t, int64(222), balance.Available[1].Amount)
	assert.Equal(t, CurrencyEUR, balance.Available[1].Currency)
	assert.Equal(t, int64(222), balance.Available[1].Amount)

	assert.Equal(t, int64(333), balance.Pending[0].Amount)
	assert.Equal(t, CurrencyUSD, balance.Pending[0].Currency)
	assert.Equal(t, int64(444), balance.Pending[1].Amount)
	assert.Equal(t, CurrencyEUR, balance.Pending[1].Currency)

	// Confirm source-type deserialization works
	assert.Equal(t, int64(100), balance.Available[0].SourceTypes[BalanceSourceTypeBankAccount])
	assert.Equal(t, int64(11), balance.Available[0].SourceTypes[BalanceSourceTypeCard])
}
