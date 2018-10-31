package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestPerson_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Person
		err := json.Unmarshal([]byte(`"person_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "person_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Person{ID: "person_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "person_123", v.ID)
	}
}
