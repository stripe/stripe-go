package stripe

import (
	"encoding/json"
	"testing"
)

func TestRecipientTransferUnmarshal(t *testing.T) {
	recipientTransferData := map[string]interface{}{
		"id":                   "tr_1234",
		"object":               "recipient_transfer",
		"amount":               1234,
		"amount_reversed":      0,
		"balance_transaction":  "txn_xxx",
		"card":                 "card_1234",
		"created":              1490723772,
		"currency":             "usd",
		"date":                 1490723772,
		"description":          "Payout Transaction xyz",
		"destination":          "card_1234",
		"failure_code":         "null",
		"failure_message":      "null",
		"livemode":             false,
		"method":               "instant",
		"recipient":            "rp_AAABBB",
		"reversed":             false,
		"source_type":          "card",
		"statement_descriptor": "Some co xyz",
		"status":               "pending",
		"type":                 "card",
	}

	bytes, err := json.Marshal(&recipientTransferData)
	if err != nil {
		t.Error(err)
	}

	var recipientTransfer RecipientTransfer
	err = json.Unmarshal(bytes, &recipientTransfer)
	if err != nil {
		t.Error(err)
	}

	if recipientTransfer.ID != "tr_1234" {
		t.Errorf("Problem deserializing the ID, got %v", recipientTransfer.ID)
	}

	if recipientTransfer.BalanceTransaction == nil {
		t.Errorf("Problem deserializing balance_transaction, got nothing.")
	} else if recipientTransfer.BalanceTransaction.ID != "txn_xxx" {
		t.Errorf("Problem deserializing balance_transaction, got %v", recipientTransfer.BalanceTransaction)
	}

	if recipientTransfer.Card == nil {
		t.Errorf("Problem deserializing card as it's null")
	} else if recipientTransfer.Card.ID != "card_1234" {
		t.Errorf("Problem deserializing card id, got %v", recipientTransfer.Card.ID)
	}
}
