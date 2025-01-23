//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"time"
)

// The card network for this settlement report. One of ["visa", "maestro"]
type IssuingSettlementNetwork string

// List of values that IssuingSettlementNetwork can take
const (
	IssuingSettlementNetworkMaestro IssuingSettlementNetwork = "maestro"
	IssuingSettlementNetworkVisa    IssuingSettlementNetwork = "visa"
)

// The current processing status of this settlement.
type IssuingSettlementStatus string

// List of values that IssuingSettlementStatus can take
const (
	IssuingSettlementStatusComplete IssuingSettlementStatus = "complete"
	IssuingSettlementStatusPending  IssuingSettlementStatus = "pending"
)

// When a non-stripe BIN is used, any use of an [issued card](https://stripe.com/docs/issuing) must be settled directly with the card network. The net amount owed is represented by an Issuing `Settlement` object.
type IssuingSettlement struct {
	// The Bank Identification Number reflecting this settlement record.
	Bin string `json:"bin"`
	// The date that the transactions are cleared and posted to user's accounts.
	ClearingDate time.Time `json:"clearing_date"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created time.Time `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The total interchange received as reimbursement for the transactions.
	InterchangeFees int64 `json:"interchange_fees"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The total net amount required to settle with the network.
	NetTotal int64 `json:"net_total"`
	// The card network for this settlement report. One of ["visa", "maestro"]
	Network IssuingSettlementNetwork `json:"network"`
	// The total amount of fees owed to the network.
	NetworkFees int64 `json:"network_fees"`
	// The Settlement Identification Number assigned by the network.
	NetworkSettlementIdentifier string `json:"network_settlement_identifier"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// One of `international` or `uk_national_net`.
	SettlementService string `json:"settlement_service"`
	// The current processing status of this settlement.
	Status IssuingSettlementStatus `json:"status"`
	// The total number of transactions reflected in this settlement.
	TransactionCount int64 `json:"transaction_count"`
	// The total transaction amount reflected in this settlement.
	TransactionVolume int64 `json:"transaction_volume"`
}

// UnmarshalJSON handles deserialization of an IssuingSettlement.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingSettlement) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type issuingSettlement IssuingSettlement
	v := struct {
		ClearingDate int64 `json:"clearing_date"`
		Created      int64 `json:"created"`
		*issuingSettlement
	}{
		issuingSettlement: (*issuingSettlement)(i),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	i.ClearingDate = time.Unix(v.ClearingDate, 0)
	i.Created = time.Unix(v.Created, 0)
	return nil
}

// MarshalJSON handles serialization of an IssuingSettlement.
// This custom marshaling is needed to handle the time fields correctly.
func (i IssuingSettlement) MarshalJSON() ([]byte, error) {
	type issuingSettlement IssuingSettlement
	v := struct {
		ClearingDate int64 `json:"clearing_date"`
		Created      int64 `json:"created"`
		issuingSettlement
	}{
		issuingSettlement: (issuingSettlement)(i),
		ClearingDate:      i.ClearingDate.Unix(),
		Created:           i.Created.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}
