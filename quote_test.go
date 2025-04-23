package stripe

import (
	"encoding/json"
	"testing"

	"github.com/max-cape/stripe-go-test/form"
	assert "github.com/stretchr/testify/require"
)

func TestQuoteSubscriptionDataParams_AppendTo(t *testing.T) {
	{
		params := &QuoteSubscriptionDataParams{EffectiveDateCurrentPeriodEnd: Bool(true)}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"current_period_end"}, body.Get("effective_date"))
	}
}

func TestQuote_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Quote
		err := json.Unmarshal([]byte(`"qt_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "qt_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Quote{ID: "qt_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "qt_123", v.ID)
	}
}
