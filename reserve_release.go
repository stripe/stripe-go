//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Indicates which party created this ReserveRelease.
type ReserveReleaseCreatedBy string

// List of values that ReserveReleaseCreatedBy can take
const (
	ReserveReleaseCreatedByApplication ReserveReleaseCreatedBy = "application"
	ReserveReleaseCreatedByStripe      ReserveReleaseCreatedBy = "stripe"
)

// The reason for the ReserveRelease, indicating why the funds were released.
type ReserveReleaseReason string

// List of values that ReserveReleaseReason can take
const (
	ReserveReleaseReasonBulkHoldExpiry    ReserveReleaseReason = "bulk_hold_expiry"
	ReserveReleaseReasonHoldReleasedEarly ReserveReleaseReason = "hold_released_early"
	ReserveReleaseReasonHoldReversed      ReserveReleaseReason = "hold_reversed"
	ReserveReleaseReasonPlanDisabled      ReserveReleaseReason = "plan_disabled"
)

// The type of source transaction.
type ReserveReleaseSourceTransactionType string

// List of values that ReserveReleaseSourceTransactionType can take
const (
	ReserveReleaseSourceTransactionTypeDispute ReserveReleaseSourceTransactionType = "dispute"
	ReserveReleaseSourceTransactionTypeRefund  ReserveReleaseSourceTransactionType = "refund"
)

// Returns a list of ReserveReleases previously created. The ReserveReleases are returned in sorted order, with the most recent ReserveReleases appearing first.
type ReserveReleaseListParams struct {
	ListParams `form:"*"`
	// Only return ReserveReleases associated with the currency specified by this currency code. Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Only return ReserveReleases associated with the ReserveHold specified by this ReserveHold ID.
	ReserveHold *string `form:"reserve_hold"`
	// Only return ReserveReleases associated with the ReservePlan specified by this ReservePlan ID.
	ReservePlan *string `form:"reserve_plan"`
}

// AddExpand appends a new field to expand.
func (p *ReserveReleaseListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieve a ReserveRelease.
type ReserveReleaseParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *ReserveReleaseParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieve a ReserveRelease.
type ReserveReleaseRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *ReserveReleaseRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type ReserveReleaseSourceTransaction struct {
	// The ID of the dispute.
	Dispute *Dispute `json:"dispute"`
	// The ID of the refund.
	Refund *Refund `json:"refund"`
	// The type of source transaction.
	Type ReserveReleaseSourceTransactionType `json:"type"`
}

// ReserveReleases represent the release of funds from a ReserveHold.
type ReserveRelease struct {
	APIResource
	// Amount released. A positive integer representing how much is released in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Amount int64 `json:"amount"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Indicates which party created this ReserveRelease.
	CreatedBy ReserveReleaseCreatedBy `json:"created_by"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The reason for the ReserveRelease, indicating why the funds were released.
	Reason ReserveReleaseReason `json:"reason"`
	// The release timestamp of the funds.
	ReleasedAt int64 `json:"released_at"`
	// The ReserveHold this ReserveRelease is associated with.
	ReserveHold *ReserveHold `json:"reserve_hold"`
	// The ReservePlan ID this ReserveRelease is associated with. This field is only populated if a ReserveRelease is created by a ReservePlan disable operation, or from a scheduled ReservedHold expiry.
	ReservePlan       *ReservePlan                     `json:"reserve_plan"`
	SourceTransaction *ReserveReleaseSourceTransaction `json:"source_transaction"`
}

// ReserveReleaseList is a list of Releases as retrieved from a list endpoint.
type ReserveReleaseList struct {
	APIResource
	ListMeta
	Data []*ReserveRelease `json:"data"`
}
