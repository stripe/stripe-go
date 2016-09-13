package stripe

import (
	"encoding/json"
	"testing"
)

func TestAccountUnmarshal(t *testing.T) {
	accountData := map[string]interface{}{
		"id": "acct_1234",
		"external_accounts": map[string]interface{}{
			"object":   "list",
			"has_more": true,
			"data": []map[string]interface{}{
				{
					"object": "bank_account",
					"id":     "ba_1234",
				},
				{
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

	if len(account.ExternalAccounts.Values) != 2 {
		t.Errorf("Problem deserializing account, got wrong number of external accounts")
	}

	bankAccount := account.ExternalAccounts.Values[0]
	if bankAccount == nil {
		t.Errorf("Problem deserializing account, didn't get a bank account")
	}
	if "ba_1234" != bankAccount.ID {
		t.Errorf("Problem deserializing account, got bank account ID %v", bankAccount.ID)
	}

	card := account.ExternalAccounts.Values[1]
	if card == nil {
		t.Errorf("Problem deserializing account, didn't get a card")
	}

	if "card_1234" != card.ID {
		t.Errorf("Problem deserializing account, got card ID %v", card.ID)
	}
}

func TestAddressAppendDetails(t *testing.T) {
	actual := &RequestValues{}
	expected := &RequestValues{}

	address := &Address{}
	address.AppendDetails(actual, "test_address")

	// First test the empty case. All empty fields should be respected while
	// encoding.
	if expected.Encode() != actual.Encode() {
		t.Errorf("Problem encoding address, got: %v", actual.Encode())
	}

	address = &Address{
		Line1:   "test_line1",
		Line2:   "test_line2",
		City:    "test_city",
		State:   "test_state",
		Zip:     "test_zip",
		Country: "test_country",
		Town:    "test_town",
	}
	address.AppendDetails(actual, "test_address")

	expected.Add("test_address[line1]", "test_line1")
	expected.Add("test_address[line2]", "test_line2")
	expected.Add("test_address[city]", "test_city")
	expected.Add("test_address[state]", "test_state")
	expected.Add("test_address[postal_code]", "test_zip")
	expected.Add("test_address[country]", "test_country")
	expected.Add("test_address[town]", "test_town")

	if expected.Encode() != actual.Encode() {
		t.Errorf("Problem encoding address, got: %v", actual.Encode())
	}
}
