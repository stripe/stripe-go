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
	Duration *ProductCatalogTrialOfferDurationParams `form:"duration" json:"duration"`
	// Define behavior that occurs at the end of the trial.
	EndBehavior *ProductCatalogTrialOfferEndBehaviorParams `form:"end_behavior" json:"end_behavior"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// A brief, user-friendly name for the trial offer-for identification purposes.
	Name *string `form:"name" json:"name,omitempty"`
	// Price configuration during the trial period (amount, billing scheme, etc).
	Price *string `form:"price" json:"price"`
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
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The price during the trial offer.
	Price *Price `json:"price"`
}
