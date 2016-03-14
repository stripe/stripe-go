package stripe

import (
	"encoding/json"
	"testing"
)

func TestAccountUnmarshal(t *testing.T) {
	accountData := map[string]interface{}{
		"id": "acct_1234",
		"external_accounts": map[string]interface{}{
			"object": "list",
			"has_more": true,
			"data": []map[string]interface{}{
				map[string]interface{}{
					"object": "bank_account",
					"id":     "ba_1234",
				},
				map[string]interface{}{
					"object": "card",
					"id":     "card_1234",
				},
			},
		},
	}

	bytes, err := json.Marshal(&accountData)
	if err != nil {
		t.Error(err)
	}

	var account Account
	err = json.Unmarshal(bytes, &account)
	if err != nil {
		t.Error(err)
	}

	if account.ID != "acct_1234" {
		t.Errorf("Problem deserializing account, got ID %v", account.ID)
	}

	if !account.ExternalAccounts.More {
		t.Errorf("Problem deserializing account, wrong value for list More")
	}

	if len(account.ExternalAccounts.BankAccountValues) != 1 {
		t.Errorf("Problem deserializing account, got wrong number of bank accounts")
	}

	bankAccountID := account.ExternalAccounts.BankAccountValues[0].ID
	if "ba_1234" != bankAccountID {
		t.Errorf("Problem deserializing account, got bank account ID %v", bankAccountID)
	}

	if len(account.ExternalAccounts.CardValues) != 1 {
		t.Errorf("Problem deserializing account, got wrong number of cards")
	}

	cardID := account.ExternalAccounts.CardValues[0].ID
	if "card_1234" != cardID {
		t.Errorf("Problem deserializing account, got card ID %v", cardID)
	}
}
