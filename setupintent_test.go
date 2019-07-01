package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestSetupIntentNextAction_UnmarshalJSON(t *testing.T) {
	actionData := map[string]interface{}{
		"redirect_to_url": map[string]interface{}{
			"return_url": "https://stripe.com/return",
			"url":        "https://stripe.com",
		},
		"type": "redirect_to_url",
	}

	bytes, err := json.Marshal(&actionData)
	assert.NoError(t, err)

	var action SetupIntentNextAction
	err = json.Unmarshal(bytes, &action)
	assert.NoError(t, err)

	assert.Equal(t, SetupIntentNextActionTypeRedirectToURL, action.Type)
	assert.Equal(t, "https://stripe.com", action.RedirectToURL.URL)
	assert.Equal(t, "https://stripe.com/return", action.RedirectToURL.ReturnURL)
}
