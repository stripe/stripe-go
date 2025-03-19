//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

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
	ClearingDate int64 `json:"clearing_date"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The total interchange received as reimbursement for the transactions.
	InterchangeFeesAmount int64 `json:"interchange_fees_amount"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The total net amount required to settle with the network.
	NetTotalAmount int64 `json:"net_total_amount"`
	// The card network for this settlement report. One of ["visa", "maestro"]
	Network IssuingSettlementNetwork `json:"network"`
	// The total amount of fees owed to the network.
	NetworkFeesAmount int64 `json:"network_fees_amount"`
	// The Settlement Identification Number assigned by the network.
	NetworkSettlementIdentifier string `json:"network_settlement_identifier"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The total amount of any additional fees assessed by the card network.
	OtherFeesAmount int64 `json:"other_fees_amount"`
	// The total number of additional fees assessed by the card network.
	OtherFeesCount int64 `json:"other_fees_count"`
	// One of `international` or `uk_national_net`.
	SettlementService string `json:"settlement_service"`
	// The current processing status of this settlement.
	Status IssuingSettlementStatus `json:"status"`
	// The total transaction amount reflected in this settlement.
	TransactionAmount int64 `json:"transaction_amount"`
	// The total number of transactions reflected in this settlement.
	TransactionCount int64 `json:"transaction_count"`
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
	var v issuingSettlement
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingSettlement(v)
	return nil
}
