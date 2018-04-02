package stripe

import (
	"strconv"

	"github.com/stripe/stripe-go/form"
)

// PlanInterval is the list of allowed values for a plan's interval.
// Allowed values are "day", "week", "month", "year".
type PlanInterval string

const (
	// PlanBillingSchemeTiered indicates that the price per single unit is tiered
	// and can change with the total number of units.
	PlanBillingSchemeTiered string = "tiered"
	// PlanBillingSchemePerUnit indicates that each unit is billed at a fixed
	// price.
	PlanBillingSchemePerUnit string = "per_unit"
)

const (
	// PlanUsageTypeLicensed indicates that the set Quantity on a subscription item
	// will be used to bill for a subscription.
	PlanUsageTypeLicensed string = "licensed"
	// PlanUsageTypeMetered indicates that usage records must be reported instead
	// of setting a Quantity on a subscription item.
	PlanUsageTypeMetered string = "metered"
)

const (
	// PlanTransformUsageModeRoundDown represents a bucket billing configuration: a partially
	// filled bucket will count as an empty bucket.
	PlanTransformUsageModeRoundDown string = "round_down"
	// PlanTransformUsageModeRoundUp represents a bucket billing configuration: a partially
	// filled bucket will count as a full bucket.
	PlanTransformUsageModeRoundUp string = "round_up"
)

// Plan is the resource representing a Stripe plan.
// For more details see https://stripe.com/docs/api#plans.
type Plan struct {
	Amount         uint64              `json:"amount"`
	BillingScheme  string              `json:"billing_scheme"`
	Created        int64               `json:"created"`
	Currency       Currency            `json:"currency"`
	Deleted        bool                `json:"deleted"`
	ID             string              `json:"id"`
	Interval       PlanInterval        `json:"interval"`
	IntervalCount  uint64              `json:"interval_count"`
	Live           bool                `json:"livemode"`
	Meta           map[string]string   `json:"metadata"`
	Nickname       string              `json:"nickname"`
	Product        string              `json:"product"`
	Tiers          []*PlanTier         `json:"tiers"`
	TiersMode      string              `json:"tiers_mode"`
	TransformUsage *PlanTransformUsage `json:"transform_usage"`
	TrialPeriod    uint64              `json:"trial_period_days"`
	UsageType      string              `json:"usage_type"`
}

// PlanList is a list of plans as returned from a list endpoint.
type PlanList struct {
	ListMeta
	Values []*Plan `json:"data"`
}

// PlanListParams is the set of parameters that can be used when listing plans.
// For more details see https://stripe.com/docs/api#list_plans.
type PlanListParams struct {
	ListParams   `form:"*"`
	Created      int64             `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// PlanParams is the set of parameters that can be used when creating or updating a plan.
// For more details see https://stripe.com/docs/api#create_plan and https://stripe.com/docs/api#update_plan.
type PlanParams struct {
	Params         `form:"*"`
	Amount         uint64                    `form:"amount"`
	AmountZero     bool                      `form:"amount,zero"`
	BillingScheme  string                    `form:"billing_scheme"`
	Currency       Currency                  `form:"currency"`
	ID             string                    `form:"id"`
	Interval       PlanInterval              `form:"interval"`
	IntervalCount  uint64                    `form:"interval_count"`
	Nickname       string                    `form:"nickname"`
	Product        *PlanProductParams        `form:"product"`
	ProductID      *string                   `form:"product"`
	Tiers          []*PlanTierParams         `form:"tiers,indexed"`
	TiersMode      string                    `form:"tiers_mode"`
	TransformUsage *PlanTransformUsageParams `form:"transform_usage"`
	TrialPeriod    uint64                    `form:"trial_period_days"`
	UsageType      string                    `form:"usage_type"`
}

// PlanTier configures tiered pricing
type PlanTier struct {
	Amount uint64 `json:"amount"`
	UpTo   uint64 `json:"up_to"`
}

// PlanTransformUsage represents the bucket billing configuration.
type PlanTransformUsage struct {
	DivideBy int64  `json:"bucket_size"`
	Round    string `json:"round"`
}

// PlanTransformUsageParams represents the bucket billing configuration.
type PlanTransformUsageParams struct {
	DivideBy int64  `form:"bucket_size"`
	Round    string `form:"round"`
}

// PlanTierParams configures tiered pricing
type PlanTierParams struct {
	Params  `form:"*"`
	Amount  uint64 `form:"amount"`
	UpTo    uint64 `form:"-"` // handled in custom AppendTo
	UpToInf bool   `form:"-"` // handled in custom AppendTo
}

// AppendTo implements custom up_to serialisation logic for tiers configuration
func (p *PlanTierParams) AppendTo(body *form.Values, keyParts []string) {
	if p.UpToInf {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	} else {
		body.Add(form.FormatKey(append(keyParts, "up_to")), strconv.FormatUint(p.UpTo, 10))
	}
}

// PlanProductParams is the set of parameters that can be used when creating a product inside a plan
// This can only be used on plan creation and won't work on plan update.
// For more details see https://stripe.com/docs/api#create_plan-product and https://stripe.com/docs/api#update_plan-product
type PlanProductParams struct {
	ID                  string            `form:"id"`
	Name                string            `form:"name"`
	Meta                map[string]string `form:"metadata"`
	StatementDescriptor string            `form:"statement_descriptor"`
}
