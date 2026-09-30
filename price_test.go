package stripe

import (
	"encoding/json"
	"testing"

	"github.com/shopspring/decimal"
	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/v86/form"
)

func TestPrice_Unmarshal(t *testing.T) {
	priceData := map[string]interface{}{
		"id":     "price_123",
		"object": "price",
		"recurring": map[string]interface{}{
			"interval":       "month",
			"interval_count": 6,
			"usage_type":     "metered",
		},
		"tiers": []map[string]interface{}{
			{
				"flat_amount_decimal": "0.0111111111",
				"up_to":               5,
			},
			{
				"flat_amount_decimal": "0.0222222222",
				"up_to":               10,
			},
			{
				"flat_amount_decimal": "0.0333333333",
			},
		},
		"tiers_mode":          "volume",
		"unit_amount":         0,
		"unit_amount_decimal": "12345678901234567890.123456789012",
	}

	bytes, err := json.Marshal(&priceData)
	assert.NoError(t, err)

	var price Price
	err = json.Unmarshal(bytes, &price)
	assert.NoError(t, err)

	assert.Equal(t, PriceRecurringIntervalMonth, price.Recurring.Interval)
	assert.Equal(t, int64(6), price.Recurring.IntervalCount)
	assert.Equal(t, PriceRecurringUsageTypeMetered, price.Recurring.UsageType)
	assert.Equal(t, 3, len(price.Tiers))
	assert.True(t, decimal.RequireFromString("0.0111111111").Equal(price.Tiers[0].FlatAmountDecimal))
	assert.Equal(t, int64(5), price.Tiers[0].UpTo)
	assert.Equal(t, PriceTiersModeVolume, price.TiersMode)
	assert.Equal(t, "12345678901234567890.123456789012", price.UnitAmountDecimal.String())

	roundTrip, err := json.Marshal(&price)
	assert.NoError(t, err)
	var roundTripped Price
	assert.NoError(t, json.Unmarshal(roundTrip, &roundTripped))
	assert.True(t, price.UnitAmountDecimal.Equal(roundTripped.UnitAmountDecimal))
}

func TestPriceParams_ExactDecimalSerialization(t *testing.T) {
	value := decimal.RequireFromString("12345678901234567890.123456789012")
	params := &PriceParams{UnitAmountDecimal: &value}

	body := &form.Values{}
	form.AppendTo(body, params)
	assert.Equal(t, []string{"12345678901234567890.123456789012"}, body.Get("unit_amount_decimal"))

	encoded, err := json.Marshal(params)
	assert.NoError(t, err)
	assert.JSONEq(t, `{"unit_amount_decimal":"12345678901234567890.123456789012"}`, string(encoded))
}

func TestPriceTierParams_AppendTo(t *testing.T) {
	params := &PriceTierParams{
		UnitAmount: Int64(500),
		UpToInf:    Bool(true),
	}

	// TODO: figure out why this doesn't work.
	body := &form.Values{}
	form.AppendTo(body, params)
	t.Logf("body = %+v", body)
	assert.Equal(t, []string{"inf"}, body.Get("up_to"))
}
