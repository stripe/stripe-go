package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestEphemeralKey_UnmarshalJSON(t *testing.T) {
	// Test that we can extract the json even if parsing fails because the
	// frontend and backend may be using different API versions
	invalidJSON := []byte("{\"foo\":5}")
	key := &EphemeralKey{}

	err := json.Unmarshal(invalidJSON, key)
	assert.NoError(t, err)
	assert.Empty(t, key.ID)
	assert.Equal(t, invalidJSON, key.RawJSON)
}
