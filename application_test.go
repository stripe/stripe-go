package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestApplicationUnmarshal(t *testing.T) {
	applicationData := map[string]interface{}{
		"id":   "ca_123",
		"name": "My Application Name",
	}

	bytes, err := json.Marshal(&applicationData)
	assert.NoError(t, err)

	var application Application
	err = json.Unmarshal(bytes, &application)
	assert.NoError(t, err)

	assert.Equal(t, "ca_123", application.ID)
	assert.Equal(t, "My Application Name", application.Name)
}
