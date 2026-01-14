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

// The relative duration of the trial period computed as the number of recurring price intervals.
type ProductCatalogTrialOfferDurationRelativeParams struct {
	// The number of recurring price's interval to apply for the trial period.
	Iterations *int64 `form:"iterations"`
}

// Duration of one service period of the trial.
type ProductCatalogTrialOfferDurationParams struct {
	// The relative duration of the trial period computed as the number of recurring price intervals.
	Relative *ProductCatalogTrialOfferDurationRelativeParams `form:"relative"`
	// Specifies how the trial offer duration is determined.
	Type *string `form:"type"`
}

// The transition to apply when the trial offer ends.
type ProductCatalogTrialOfferEndBehaviorTransitionParams struct {
	// The price to transition the recurring item to when the trial offer ends.
	Price *string `form:"price"`
}

// Define behavior that occurs at the end of the trial.
type ProductCatalogTrialOfferEndBehaviorParams struct {
	// The transition to apply when the trial offer ends.
	Transition *ProductCatalogTrialOfferEndBehaviorTransitionParams `form:"transition"`
}

// Creates a trial offer.
type ProductCatalogTrialOfferParams struct {
	Params `form:"*"`
	// Duration of one service period of the trial.
	Duration *ProductCatalogTrialOfferDurationParams `form:"duration"`
	// Define behavior that occurs at the end of the trial.
	EndBehavior *ProductCatalogTrialOfferEndBehaviorParams `form:"end_behavior"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// A brief, user-friendly name for the trial offer-for identification purposes.
	Name *string `form:"name"`
	// Price configuration during the trial period (amount, billing scheme, etc).
	Price *string `form:"price"`
}

// AddExpand appends a new field to expand.
func (p *ProductCatalogTrialOfferParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The relative duration of the trial period computed as the number of recurring price intervals.
type ProductCatalogTrialOfferCreateDurationRelativeParams struct {
	// The number of recurring price's interval to apply for the trial period.
	Iterations *int64 `form:"iterations"`
}

// Duration of one service period of the trial.
type ProductCatalogTrialOfferCreateDurationParams struct {
	// The relative duration of the trial period computed as the number of recurring price intervals.
	Relative *ProductCatalogTrialOfferCreateDurationRelativeParams `form:"relative"`
	// Specifies how the trial offer duration is determined.
	Type *string `form:"type"`
}

// The transition to apply when the trial offer ends.
type ProductCatalogTrialOfferCreateEndBehaviorTransitionParams struct {
	// The price to transition the recurring item to when the trial offer ends.
	Price *string `form:"price"`
}

// Define behavior that occurs at the end of the trial.
type ProductCatalogTrialOfferCreateEndBehaviorParams struct {
	// The transition to apply when the trial offer ends.
	Transition *ProductCatalogTrialOfferCreateEndBehaviorTransitionParams `form:"transition"`
}

// Creates a trial offer.
type ProductCatalogTrialOfferCreateParams struct {
	Params `form:"*"`
	// Duration of one service period of the trial.
	Duration *ProductCatalogTrialOfferCreateDurationParams `form:"duration"`
	// Define behavior that occurs at the end of the trial.
	EndBehavior *ProductCatalogTrialOfferCreateEndBehaviorParams `form:"end_behavior"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// A brief, user-friendly name for the trial offer-for identification purposes.
	Name *string `form:"name"`
	// Price configuration during the trial period (amount, billing scheme, etc).
	Price *string `form:"price"`
}

// AddExpand appends a new field to expand.
func (p *ProductCatalogTrialOfferCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type ProductCatalogTrialOfferDurationRelative struct {
	// The number of iterations of the price's interval for this trial offer.
	Iterations int64 `json:"iterations"`
}
type ProductCatalogTrialOfferDuration struct {
	Relative *ProductCatalogTrialOfferDurationRelative `json:"relative"`
	// The type of trial offer duration.
	Type ProductCatalogTrialOfferDurationType `json:"type"`
}
type ProductCatalogTrialOfferEndBehaviorTransition struct {
	// The new price to use at the end of the trial offer period.
	Price string `json:"price"`
}
type ProductCatalogTrialOfferEndBehavior struct {
	Transition *ProductCatalogTrialOfferEndBehaviorTransition `json:"transition"`
	// The type of behavior when the trial offer ends.
	Type ProductCatalogTrialOfferEndBehaviorType `json:"type"`
}

// Resource for the TrialOffer API, used to describe a subscription item's trial period settings.
// Renders a TrialOffer object that describes the price, duration, end_behavior of a trial offer.
type ProductCatalogTrialOffer struct {
	APIResource
	Duration    *ProductCatalogTrialOfferDuration    `json:"duration"`
	EndBehavior *ProductCatalogTrialOfferEndBehavior `json:"end_behavior"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// A brief, user-friendly name for the trial offer-for identification purposes.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The price during the trial offer.
	Price *Price `json:"price"`
}
