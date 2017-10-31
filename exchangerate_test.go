package stripe

import (
	"encoding/json"
	"testing"
)

func TestExchangeRateUnmarshal(t *testing.T) {
	exchangeRateData := map[string]interface{}{
		"id":     "usd",
		"object": "exchange_rate",
		"rates": map[string]interface{}{
			"eur": 0.845876,
		},
	}

	bytes, err := json.Marshal(&exchangeRateData)
	if err != nil {
		t.Error(err)
	}

	var exchangeRate ExchangeRate
	err = json.Unmarshal(bytes, &exchangeRate)
	if err != nil {
		t.Error(err)
	}

	if exchangeRate.Rates == nil {
		t.Errorf("Problem deserializing rates, got nothing.")
	}

	if exchangeRate.Rates["eur"] != 0.845876 {
		t.Errorf("Problem deserializing rates[eur], got %v", exchangeRate.Rates["eur"])
	}
}
