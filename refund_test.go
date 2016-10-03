package stripe

import (
	"encoding/json"
	"testing"
)

func TestRefundUnmarshal(t *testing.T) {
	refundData := map[string]interface{}{
		"id":     "re_1234",
		"object": "refund",
		"charge": "ch_1234",
	}

	bytes, err := json.Marshal(&refundData)
	if err != nil {
		t.Error(err)
	}

	var refund Refund
	err = json.Unmarshal(bytes, &refund)
	if err != nil {
		t.Error(err)
	}

	if refund.ID != "re_1234" {
		t.Errorf("Problem deserializing refund, got ID %v", refund.ID)
	}

	if refund.Charge == nil {
		t.Errorf("Problem deserializing refund, didn't get a Charge")
	}

	if refund.Charge.ID != "ch_1234" {
		t.Errorf("Problem deserializing refund.charge, wrong value for ID")
	}
}

func TestRefundUnmarshalExpandedCharge(t *testing.T) {
	refundData := map[string]interface{}{
		"id":     "re_1234",
		"object": "refund",
		"charge": map[string]interface{}{
			"id":     "ch_1234",
			"object": "charge",
		},
	}

	bytes, err := json.Marshal(&refundData)
	if err != nil {
		t.Error(err)
	}

	var refund Refund
	err = json.Unmarshal(bytes, &refund)
	if err != nil {
		t.Error(err)
	}

	if refund.ID != "re_1234" {
		t.Errorf("Problem deserializing refund, got ID %v", refund.ID)
	}

	if refund.Charge == nil {
		t.Errorf("Problem deserializing refund, didn't get a Charge")
	}

	if refund.Charge.ID != "ch_1234" {
		t.Errorf("Problem deserializing refund.charge, wrong value for ID")
	}
}
