package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/form"
)

func TestAccountUnmarshal(t *testing.T) {
	accountData := map[string]interface{}{
		"id": "acct_123",
		"external_accounts": map[string]interface{}{
			"object":   "list",
			"has_more": true,
			"data": []map[string]interface{}{
				{
					"object": "bank_account",
					"id":     "ba_123",
				},
				{
					"object": "card",
					"id":     "card_123",
				},
			},
		},
	}

	bytes, err := json.Marshal(&accountData)
	assert.NoError(t, err)

	var account Account
	err = json.Unmarshal(bytes, &account)
	assert.NoError(t, err)

	assert.Equal(t, "acct_123", account.ID)
	assert.Equal(t, true, account.ExternalAccounts.More)

	assert.Equal(t, 2, len(account.ExternalAccounts.Values))
	assert.Equal(t, "ba_123", account.ExternalAccounts.Values[0].ID)
	assert.Equal(t, "card_123", account.ExternalAccounts.Values[1].ID)
}

func TestIdentityDocument_Appendto(t *testing.T) {
	document := &IdentityDocument{ID: "file_123"}
	body := &form.Values{}
	document.AppendTo(body,
		[]string{"legal_entity", "additional_owners", "0", "verification", "document"})
	assert.Equal(t,
		[]string{"file_123"},
		body.Get("legal_entity[additional_owners][0][verification][document]"),
	)
}
