package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestCharge_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Charge
		err := json.Unmarshal([]byte(`"ch_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "ch_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Charge{ID: "ch_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "ch_123", v.ID)
	}
}

func TestChargeOutcomeRule_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v ChargeOutcomeRule
		err := json.Unmarshal([]byte(`"ssr_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "ssr_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := ChargeOutcomeRule{ID: "ssr_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "ssr_123", v.ID)
	}
}
