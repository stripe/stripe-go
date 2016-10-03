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

	source_tx := transfer.SourceTx
	if source_tx == nil {
		t.Errorf("Problem deserializing transfer, didn't get a SourceTx")
	}

	if source_tx.ID != "ch_1234" {
		t.Errorf("Problem deserializing transfer.source_transaction, wrong value for ID")
	}

	if source_tx.Type != TransactionSourceCharge {
		t.Errorf("Problem deserializing transfer.source_transaction, wrong value for Type")
	}

	if source_tx.Charge == nil {
		t.Errorf("Problem deserializing transfer.source_transaction, didn't get a Charge")
	}

	if source_tx.Charge.ID != "ch_1234" {
		t.Errorf("Problem deserializing transfer.source_transaction, wrong value for Charge.ID")
	}
}
