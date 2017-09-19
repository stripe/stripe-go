package stripe

import (
	"encoding/json"
	"testing"
)

func TestTransferUnmarshal(t *testing.T) {
	transferData := map[string]interface{}{
		"id":     "tr_1234",
		"object": "transfer",
		"source_transaction": map[string]interface{}{
			"id":     "ch_1234",
			"object": "charge",
		},
	}

	bytes, err := json.Marshal(&transferData)
	if err != nil {
		t.Error(err)
	}

	var transfer Transfer
	err = json.Unmarshal(bytes, &transfer)
	if err != nil {
		t.Error(err)
	}

	if transfer.ID != "tr_1234" {
		t.Errorf("Problem deserializing transfer, got ID %v", transfer.ID)
	}

	source_transaction := transfer.SourceTransaction
	if source_transaction == nil {
		t.Errorf("Problem deserializing transfer, didn't get a SourceTransaction")
	}

	if source_transaction.ID != "ch_1234" {
		t.Errorf("Problem deserializing transfer.source_transaction, wrong value for ID")
	}

	if source_transaction.Type != BalanceTransactionSourceTypeCharge {
		t.Errorf("Problem deserializing transfer.source_transaction, wrong value for Type")
	}

	if source_transaction.Charge == nil {
		t.Errorf("Problem deserializing transfer.source_transaction, didn't get a Charge")
	}

	if source_transaction.Charge.ID != "ch_1234" {
		t.Errorf("Problem deserializing transfer.source_transaction, wrong value for Charge.ID")
	}
}
