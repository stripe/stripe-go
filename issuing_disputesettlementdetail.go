//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of event corresponding to this dispute settlement detail, representing the stage in the dispute network lifecycle.
type IssuingDisputeSettlementDetailEventType string

// List of values that IssuingDisputeSettlementDetailEventType can take
const (
	IssuingDisputeSettlementDetailEventTypeFiling        IssuingDisputeSettlementDetailEventType = "filing"
	IssuingDisputeSettlementDetailEventTypeLoss          IssuingDisputeSettlementDetailEventType = "loss"
	IssuingDisputeSettlementDetailEventTypeRepresentment IssuingDisputeSettlementDetailEventType = "representment"
	IssuingDisputeSettlementDetailEventTypeWin           IssuingDisputeSettlementDetailEventType = "win"
)

// The card network for this dispute settlement detail. One of ["visa", "mastercard", "maestro"]
type IssuingDisputeSettlementDetailNetwork string

// List of values that IssuingDisputeSettlementDetailNetwork can take
const (
	IssuingDisputeSettlementDetailNetworkMaestro    IssuingDisputeSettlementDetailNetwork = "maestro"
	IssuingDisputeSettlementDetailNetworkMastercard IssuingDisputeSettlementDetailNetwork = "mastercard"
	IssuingDisputeSettlementDetailNetworkVisa       IssuingDisputeSettlementDetailNetwork = "visa"
)

// Returns a list of Issuing DisputeSettlementDetail objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
type IssuingDisputeSettlementDetailListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Select the Issuing dispute settlement details for the given settlement.
	Settlement *string `form:"settlement"`
}

// AddExpand appends a new field to expand.
func (p *IssuingDisputeSettlementDetailListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves an Issuing DisputeSettlementDetail object.
type IssuingDisputeSettlementDetailParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *IssuingDisputeSettlementDetailParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Represents a record from the card network of a money movement or change in state for an Issuing dispute. These records are included in the settlement reports that we receive from networks and expose to users as Settlement objects.
type IssuingDisputeSettlementDetail struct {
	APIResource
	// Disputed amount in the card's currency and in the smallest currency unit. Usually the amount of the transaction, but can differ (usually because of currency fluctuation).
	Amount int64 `json:"amount"`
	// The card used to make the original transaction.
	Card string `json:"card"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The currency the original transaction was made in. Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The ID of the linked dispute.
	Dispute string `json:"dispute"`
	// The type of event corresponding to this dispute settlement detail, representing the stage in the dispute network lifecycle.
	EventType IssuingDisputeSettlementDetailEventType `json:"event_type"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The card network for this dispute settlement detail. One of ["visa", "mastercard", "maestro"]
	Network IssuingDisputeSettlementDetailNetwork `json:"network"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The ID of the linked card network settlement.
	Settlement string `json:"settlement"`
}

// IssuingDisputeSettlementDetailList is a list of DisputeSettlementDetails as retrieved from a list endpoint.
type IssuingDisputeSettlementDetailList struct {
	APIResource
	ListMeta
	Data []*IssuingDisputeSettlementDetail `json:"data"`
}
