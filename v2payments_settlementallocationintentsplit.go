//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Type of the flow linked to the transaction which settled the SettlementAllocationIntentSplit. The field matching this value will contain the ID of the flow.
type V2PaymentsSettlementAllocationIntentSplitFlowType string

// List of values that V2PaymentsSettlementAllocationIntentSplitFlowType can take
const (
	V2PaymentsSettlementAllocationIntentSplitFlowTypeOutboundPayment  V2PaymentsSettlementAllocationIntentSplitFlowType = "outbound_payment"
	V2PaymentsSettlementAllocationIntentSplitFlowTypeOutboundTransfer V2PaymentsSettlementAllocationIntentSplitFlowType = "outbound_transfer"
	V2PaymentsSettlementAllocationIntentSplitFlowTypeReceivedCredit   V2PaymentsSettlementAllocationIntentSplitFlowType = "received_credit"
)

// The status of the SettlementAllocationIntentSplit.
type V2PaymentsSettlementAllocationIntentSplitStatus string

// List of values that V2PaymentsSettlementAllocationIntentSplitStatus can take
const (
	V2PaymentsSettlementAllocationIntentSplitStatusCanceled V2PaymentsSettlementAllocationIntentSplitStatus = "canceled"
	V2PaymentsSettlementAllocationIntentSplitStatusPending  V2PaymentsSettlementAllocationIntentSplitStatus = "pending"
	V2PaymentsSettlementAllocationIntentSplitStatusSettled  V2PaymentsSettlementAllocationIntentSplitStatus = "settled"
)

// The type of the SettlementAllocationIntentSplit.
type V2PaymentsSettlementAllocationIntentSplitType string

// List of values that V2PaymentsSettlementAllocationIntentSplitType can take
const (
	V2PaymentsSettlementAllocationIntentSplitTypeCredit V2PaymentsSettlementAllocationIntentSplitType = "credit"
	V2PaymentsSettlementAllocationIntentSplitTypeDebit  V2PaymentsSettlementAllocationIntentSplitType = "debit"
)

// The amount and currency of the SettlementAllocationIntentSplit.
type V2PaymentsSettlementAllocationIntentSplitAmount struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency,omitempty"`
	// A non-negative integer representing how much to charge in the [smallest currency unit](https://docs.stripe.com/currencies#minor-units).
	Value int64 `json:"value,omitempty"`
}

// Details about the Flow object that settled the split.
type V2PaymentsSettlementAllocationIntentSplitFlow struct {
	// If applicable, the ID of the OutboundPayment that created this transaction.
	OutboundPayment string `json:"outbound_payment,omitempty"`
	// If applicable, the ID of the OutboundTransfer that created this transaction.
	OutboundTransfer string `json:"outbound_transfer,omitempty"`
	// If applicable, the ID of the ReceivedCredit that created this transaction.
	ReceivedCredit string `json:"received_credit,omitempty"`
	// Type of the flow linked to the transaction which settled the SettlementAllocationIntentSplit. The field matching this value will contain the ID of the flow.
	Type V2PaymentsSettlementAllocationIntentSplitFlowType `json:"type"`
}

// SettlementAllocationIntentSplit resource.
type V2PaymentsSettlementAllocationIntentSplit struct {
	APIResource
	// The account id against which the SettlementAllocationIntentSplit should be settled.
	Account string `json:"account"`
	// The amount and currency of the SettlementAllocationIntentSplit.
	Amount *V2PaymentsSettlementAllocationIntentSplitAmount `json:"amount"`
	// Timestamp at which SettlementAllocationIntentSplit was created.
	Created time.Time `json:"created"`
	// Details about the Flow object that settled the split.
	Flow *V2PaymentsSettlementAllocationIntentSplitFlow `json:"flow"`
	// Unique identifier for the SettlementAllocationIntentSplit.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The ID of the SettlementAllocationIntent that this split belongs too.
	SettlementAllocationIntent string `json:"settlement_allocation_intent"`
	// The status of the SettlementAllocationIntentSplit.
	Status V2PaymentsSettlementAllocationIntentSplitStatus `json:"status"`
	// The type of the SettlementAllocationIntentSplit.
	Type V2PaymentsSettlementAllocationIntentSplitType `json:"type"`
}
