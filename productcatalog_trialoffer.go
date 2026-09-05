//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of trial offer duration.
type ProductCatalogTrialOfferDurationType string

// List of values that ProductCatalogTrialOfferDurationType can take
const (
	ProductCatalogTrialOfferDurationTypeRelative  ProductCatalogTrialOfferDurationType = "relative"
	ProductCatalogTrialOfferDurationTypeTimestamp ProductCatalogTrialOfferDurationType = "timestamp"
)

// The type of behavior when the trial offer ends.
type ProductCatalogTrialOfferEndBehaviorType string

// List of values that ProductCatalogTrialOfferEndBehaviorType can take
const (
	ProductCatalogTrialOfferEndBehaviorTypeTransition ProductCatalogTrialOfferEndBehaviorType = "transition"
)

// Returns a list of trial offers.
type ProductCatalogTrialOfferListParams struct {
	ListParams `form:"*"`
	// Only return trial offers that are active (`true`) or archived (`false`). If omitted, both active and archived trial offers are returned.
	Active *bool `form:"active" json:"active,omitempty"`
	// Only return trial offers that were created during the given date interval.
	Created *int64 `form:"created" json:"created,omitempty"`
	// Only return trial offers that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created" json:"-"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Only return trial offers that reference these prices (during the trial period).
	Prices []*string `form:"prices" json:"prices,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *ProductCatalogTrialOfferListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The relative duration of the trial period computed as the number of recurring price intervals.
type ProductCatalogTrialOfferDurationRelativeParams struct {
	// The number of recurring price's interval to apply for the trial period.
	Iterations *int64 `form:"iterations" json:"iterations"`
}

// Duration of one service period of the trial.
type ProductCatalogTrialOfferDurationParams struct {
	// The relative duration of the trial period computed as the number of recurring price intervals.
	Relative *ProductCatalogTrialOfferDurationRelativeParams `form:"relative" json:"relative,omitempty"`
	// Specifies how the trial offer duration is determined.
	Type *string `form:"type" json:"type"`
}

// The transition to apply when the trial offer ends.
type ProductCatalogTrialOfferEndBehaviorTransitionParams struct {
	// The price to transition the recurring item to when the trial offer ends.
	Price *string `form:"price" json:"price"`
}

// Define behavior that occurs at the end of the trial.
type ProductCatalogTrialOfferEndBehaviorParams struct {
	// The transition to apply when the trial offer ends.
	Transition *ProductCatalogTrialOfferEndBehaviorTransitionParams `form:"transition" json:"transition"`
}

// Creates a trial offer.
type ProductCatalogTrialOfferParams struct {
	Params `form:"*"`
	// Duration of one service period of the trial.
	Duration *ProductCatalogTrialOfferDurationParams `form:"duration" json:"duration,omitempty"`
	// Define behavior that occurs at the end of the trial.
	EndBehavior *ProductCatalogTrialOfferEndBehaviorParams `form:"end_behavior" json:"end_behavior,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// A brief, user-friendly name for the trial offer-for identification purposes.
	Name *string `form:"name" json:"name,omitempty"`
	// Price configuration during the trial period (amount, billing scheme, etc).
	Price *string `form:"price" json:"price,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *ProductCatalogTrialOfferParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The relative duration of the trial period computed as the number of recurring price intervals.
type ProductCatalogTrialOfferCreateDurationRelativeParams struct {
	// The number of recurring price's interval to apply for the trial period.
	Iterations *int64 `form:"iterations" json:"iterations"`
}

// Duration of one service period of the trial.
type ProductCatalogTrialOfferCreateDurationParams struct {
	// The relative duration of the trial period computed as the number of recurring price intervals.
	Relative *ProductCatalogTrialOfferCreateDurationRelativeParams `form:"relative" json:"relative,omitempty"`
	// Specifies how the trial offer duration is determined.
	Type *string `form:"type" json:"type"`
}

// The transition to apply when the trial offer ends.
type ProductCatalogTrialOfferCreateEndBehaviorTransitionParams struct {
	// The price to transition the recurring item to when the trial offer ends.
	Price *string `form:"price" json:"price"`
}

// Define behavior that occurs at the end of the trial.
type ProductCatalogTrialOfferCreateEndBehaviorParams struct {
	// The transition to apply when the trial offer ends.
	Transition *ProductCatalogTrialOfferCreateEndBehaviorTransitionParams `form:"transition" json:"transition"`
}

// Creates a trial offer.
type ProductCatalogTrialOfferCreateParams struct {
	Params `form:"*"`
	// Duration of one service period of the trial.
	Duration *ProductCatalogTrialOfferCreateDurationParams `form:"duration" json:"duration"`
	// Define behavior that occurs at the end of the trial.
	EndBehavior *ProductCatalogTrialOfferCreateEndBehaviorParams `form:"end_behavior" json:"end_behavior"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// A brief, user-friendly name for the trial offer-for identification purposes.
	Name *string `form:"name" json:"name,omitempty"`
	// Price configuration during the trial period (amount, billing scheme, etc).
	Price *string `form:"price" json:"price"`
}

// AddExpand appends a new field to expand.
func (p *ProductCatalogTrialOfferCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the trial offer with the given ID.
type ProductCatalogTrialOfferRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *ProductCatalogTrialOfferRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type ProductCatalogTrialOfferDurationRelative struct {
	// The number of iterations of the price's interval for this trial offer.
	Iterations int64 `json:"iterations"`
}
type ProductCatalogTrialOfferDuration struct {
	Relative *ProductCatalogTrialOfferDurationRelative `json:"relative,omitempty"`
	// The type of trial offer duration.
	Type ProductCatalogTrialOfferDurationType `json:"type"`
}
type ProductCatalogTrialOfferEndBehaviorTransition struct {
	// The new price to use at the end of the trial offer period.
	Price *Price `json:"price"`
}
type ProductCatalogTrialOfferEndBehavior struct {
	Transition *ProductCatalogTrialOfferEndBehaviorTransition `json:"transition"`
	// The type of behavior when the trial offer ends.
	Type ProductCatalogTrialOfferEndBehaviorType `json:"type"`
}

// Trial offers let you define free or paid introductory pricing for a subscription item.
// A TrialOffer specifies the price to charge during the trial, how long the trial lasts
// (a fixed end timestamp or a number of billing intervals), and what price the subscription
// item transitions to when the trial ends. You attach a TrialOffer to a subscription item
// using `items[current_trial][trial_offer]` when creating or updating a subscription.
type ProductCatalogTrialOffer struct {
	APIResource
	Duration    *ProductCatalogTrialOfferDuration    `json:"duration"`
	EndBehavior *ProductCatalogTrialOfferEndBehavior `json:"end_behavior"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// A brief, user-friendly name for the trial offer-for identification purposes.
	Name string `json:"name,omitempty"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The price during the trial offer.
	Price *Price `json:"price"`
}

// ProductCatalogTrialOfferList is a list of TrialOffers as retrieved from a list endpoint.
type ProductCatalogTrialOfferList struct {
	APIResource
	ListMeta
	Data []*ProductCatalogTrialOffer `json:"data"`
}
