package stripe

import (
	"encoding/json"
	"strconv"

	"github.com/stripe/stripe-go/form"
)

// PlanInterval is the list of allowed values for a plan's interval.
type PlanInterval string

// List of values that PlanInterval can take.
const (
	PlanIntervalDay   PlanInterval = "day"
	PlanIntervalWeek  PlanInterval = "week"
	PlanIntervalMonth PlanInterval = "month"
	PlanIntervalYear  PlanInterval = "year"
)

// PlanBillingScheme is the list of allowed values for a plan's billing scheme.
type PlanBillingScheme string

// List of values that PlanBillingScheme can take.
const (
	PlanBillingSchemePerUnit PlanBillingScheme = "per_unit"
	PlanBillingSchemeTiered  PlanBillingScheme = "tiered"
)

// PlanUsageType is the list of allowed values for a plan's usage type.
type PlanUsageType string

// List of values that PlanUsageType can take.
const (
	PlanUsageTypeLicensed PlanUsageType = "licensed"
	PlanUsageTypeMetered  PlanUsageType = "metered"
)

// PlanTiersMode is the list of allowed values for a plan's tiers mode.
type PlanTiersMode string

// List of values that PlanTiersMode can take.
const (
	PlanTiersModeGraduated PlanTiersMode = "graduated"
	PlanTiersModeVolume    PlanTiersMode = "volume"
)

// PlanTransformUsageRound is the list of allowed values for a plan's transform usage round logic.
type PlanTransformUsageRound string

// List of values that PlanTransformUsageRound can take.
const (
	PlanTransformUsageRoundDown PlanTransformUsageRound = "down"
	PlanTransformUsageRoundUp   PlanTransformUsageRound = "up"
)

// PlanAggregateUsage is the list of allowed values for a plan's aggregate usage.
type PlanAggregateUsage string

// List of values that PlanAggregateUsage can take.
const (
	PlanAggregateUsageLastDuringPeriod PlanAggregateUsage = "last_during_period"
	PlanAggregateUsageLastEver         PlanAggregateUsage = "last_ever"
	PlanAggregateUsageMax              PlanAggregateUsage = "max"
	PlanAggregateUsageSum              PlanAggregateUsage = "sum"
)

// Plan is the resource representing a Stripe plan.
// For more details see https://stripe.com/docs/api#plans.
type Plan struct {
	Active          bool                `json:"active"`
	AggregateUsage  string              `json:"aggregate_usage"`
	Amount          int64               `json:"amount"`
	AmountDecimal   float64             `json:"amount_decimal,string"`
	BillingScheme   PlanBillingScheme   `json:"billing_scheme"`
	Created         int64               `json:"created"`
	Currency        Currency            `json:"currency"`
	Deleted         bool                `json:"deleted"`
	ID              string              `json:"id"`
	Interval        PlanInterval        `json:"interval"`
	IntervalCount   int64               `json:"interval_count"`
	Livemode        bool                `json:"livemode"`
	Metadata        map[string]string   `json:"metadata"`
	Nickname        string              `json:"nickname"`
	Product         *Product            `json:"product"`
	Tiers           []*PlanTier         `json:"tiers"`
	TiersMode       string              `json:"tiers_mode"`
	TransformUsage  *PlanTransformUsage `json:"transform_usage"`
	TrialPeriodDays int64               `json:"trial_period_days"`
	UsageType       PlanUsageType       `json:"usage_type"`
}

// PlanList is a list of plans as returned from a list endpoint.
type PlanList struct {
	ListMeta
	Data []*Plan `json:"data"`
}

// PlanListParams is the set of parameters that can be used when listing plans.
// For more details see https://stripe.com/docs/api#list_plans.
type PlanListParams struct {
	ListParams   `form:"*"`
	Active       *bool             `form:"active"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Product      *string           `form:"product"`
}

// PlanParams is the set of parameters that can be used when creating or updating a plan.
// For more details see https://stripe.com/docs/api#create_plan and https://stripe.com/docs/api#update_plan.
type PlanParams struct {
	Params          `form:"*"`
	Active          *bool                     `form:"active"`
	AggregateUsage  *string                   `form:"aggregate_usage"`
	Amount          *int64                    `form:"amount"`
	AmountDecimal   *float64                  `form:"amount_decimal,high_precision"`
	BillingScheme   *string                   `form:"billing_scheme"`
	Currency        *string                   `form:"currency"`
	ID              *string                   `form:"id"`
	Interval        *string                   `form:"interval"`
	IntervalCount   *int64                    `form:"interval_count"`
	Nickname        *string                   `form:"nickname"`
	Product         *PlanProductParams        `form:"product"`
	ProductID       *string                   `form:"product"`
	Tiers           []*PlanTierParams         `form:"tiers"`
	TiersMode       *string                   `form:"tiers_mode"`
	TransformUsage  *PlanTransformUsageParams `form:"transform_usage"`
	TrialPeriodDays *int64                    `form:"trial_period_days"`
	UsageType       *string                   `form:"usage_type"`
}

// PlanTier configures tiered pricing
type PlanTier struct {
	FlatAmount        int64   `json:"flat_amount"`
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	UnitAmount        int64   `json:"unit_amount"`
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	UpTo              int64   `json:"up_to"`
}

// PlanTransformUsage represents the bucket billing configuration.
type PlanTransformUsage struct {
	DivideBy int64                   `json:"divide_by"`
	Round    PlanTransformUsageRound `json:"round"`
}

// PlanTransformUsageParams represents the bucket billing configuration.
type PlanTransformUsageParams struct {
	DivideBy *int64  `form:"divide_by"`
	Round    *string `form:"round"`
}

// PlanTierParams configures tiered pricing
type PlanTierParams struct {
	Params            `form:"*"`
	FlatAmount        *int64   `form:"flat_amount"`
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	UnitAmount        *int64   `form:"unit_amount"`
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	UpTo              *int64   `form:"-"` // handled in custom AppendTo
	UpToInf           *bool    `form:"-"` // handled in custom AppendTo
}

// AppendTo implements custom up_to serialisation logic for tiers configuration
func (p *PlanTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	} else {
		body.Add(form.FormatKey(append(keyParts, "up_to")), strconv.FormatInt(Int64Value(p.UpTo), 10))
	}
}

// PlanProductParams is the set of parameters that can be used when creating a product inside a plan
// This can only be used on plan creation and won't work on plan update.
// For more details see https://stripe.com/docs/api#create_plan-product and https://stripe.com/docs/api#update_plan-product
type PlanProductParams struct {
	Active              *bool             `form:"active"`
	ID                  *string           `form:"id"`
	Name                *string           `form:"name"`
	Metadata            map[string]string `form:"metadata"`
	StatementDescriptor *string           `form:"statement_descriptor"`
	UnitLabel           *string           `form:"unit_label"`
}

// UnmarshalJSON handles deserialization of a Plan.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *Plan) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type plan Plan
	var v plan
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = Plan(v)
	return nil
}
