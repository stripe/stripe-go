//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v72/form"
	"strconv"
)

// Specifies a usage aggregation strategy for plans of `usage_type=metered`. Allowed values are `sum` for summing up all usage during a period, `last_during_period` for using the last usage record reported within a period, `last_ever` for using the last usage record ever (across period bounds) or `max` which uses the usage record with the maximum reported usage during a period. Defaults to `sum`.
type PlanAggregateUsage string

// List of values that PlanAggregateUsage can take
const (
	PlanAggregateUsageLastDuringPeriod PlanAggregateUsage = "last_during_period"
	PlanAggregateUsageLastEver         PlanAggregateUsage = "last_ever"
	PlanAggregateUsageMax              PlanAggregateUsage = "max"
	PlanAggregateUsageSum              PlanAggregateUsage = "sum"
)

// Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `amount`) will be charged per unit in `quantity` (for plans with `usage_type=licensed`), or per unit of total usage (for plans with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.
type PlanBillingScheme string

// List of values that PlanBillingScheme can take
const (
	PlanBillingSchemePerUnit PlanBillingScheme = "per_unit"
	PlanBillingSchemeTiered  PlanBillingScheme = "tiered"
)

// The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.
type PlanInterval string

// List of values that PlanInterval can take
const (
	PlanIntervalDay   PlanInterval = "day"
	PlanIntervalMonth PlanInterval = "month"
	PlanIntervalWeek  PlanInterval = "week"
	PlanIntervalYear  PlanInterval = "year"
)

// Defines if the tiering price should be `graduated` or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per unit price. In `graduated` tiering, pricing can change as the quantity grows.
type PlanTiersMode string

// List of values that PlanTiersMode can take
const (
	PlanTiersModeGraduated PlanTiersMode = "graduated"
	PlanTiersModeVolume    PlanTiersMode = "volume"
)

// After division, either round the result `up` or `down`.
type PlanTransformUsageRound string

// List of values that PlanTransformUsageRound can take
const (
	PlanTransformUsageRoundDown PlanTransformUsageRound = "down"
	PlanTransformUsageRoundUp   PlanTransformUsageRound = "up"
)

// Configures how the quantity per period should be determined. Can be either `metered` or `licensed`. `licensed` automatically bills the `quantity` set when adding it to a subscription. `metered` aggregates the total usage based on usage records. Defaults to `licensed`.
type PlanUsageType string

// List of values that PlanUsageType can take
const (
	PlanUsageTypeLicensed PlanUsageType = "licensed"
	PlanUsageTypeMetered  PlanUsageType = "metered"
)

// Returns a list of your plans.
type PlanListParams struct {
	ListParams   `form:"*"`
	Active       *bool             `form:"active"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Product      *string           `form:"product"`
}

// The product the plan belongs to. This cannot be changed once it has been used in a subscription or subscription schedule.
type PlanProductParams struct {
	Active              *bool             `form:"active"`
	ID                  *string           `form:"id"`
	Metadata            map[string]string `form:"metadata"`
	Name                *string           `form:"name"`
	StatementDescriptor *string           `form:"statement_descriptor"`
	TaxCode             *string           `form:"tax_code"`
	UnitLabel           *string           `form:"unit_label"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PlanTierParams struct {
	Params            `form:"*"`
	FlatAmount        *int64   `form:"flat_amount"`
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	UnitAmount        *int64   `form:"unit_amount"`
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	UpTo              *int64   `form:"-"` // See custom AppendTo
	UpToInf           *bool    `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PlanTierParams.
func (p *PlanTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	} else {
		body.Add(
			form.FormatKey(append(keyParts, "up_to")),
			strconv.FormatInt(Int64Value(p.UpTo), 10),
		)
	}
}

// Apply a transformation to the reported usage or set quantity before computing the billed price. Cannot be combined with `tiers`.
type PlanTransformUsageParams struct {
	DivideBy *int64  `form:"divide_by"`
	Round    *string `form:"round"`
}

// You can now model subscriptions more flexibly using the [Prices API](https://stripe.com/docs/api#prices). It replaces the Plans API and is backwards compatible to simplify your migration.
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

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PlanTier struct {
	FlatAmount        int64   `json:"flat_amount"`
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	UnitAmount        int64   `json:"unit_amount"`
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	UpTo              int64   `json:"up_to"`
}

// Apply a transformation to the reported usage or set quantity before computing the amount billed. Cannot be combined with `tiers`.
type PlanTransformUsage struct {
	DivideBy int64                   `json:"divide_by"`
	Round    PlanTransformUsageRound `json:"round"`
}

// You can now model subscriptions more flexibly using the [Prices API](https://stripe.com/docs/api#prices). It replaces the Plans API and is backwards compatible to simplify your migration.
//
// Plans define the base price, currency, and billing cycle for recurring purchases of products.
// [Products](https://stripe.com/docs/api#products) help you track inventory or provisioning, and plans help you track pricing. Different physical goods or levels of service should be represented by products, and pricing options should be represented by plans. This approach lets you change prices without having to change your provisioning scheme.
//
// For example, you might have a single "gold" product that has plans for $10/month, $100/year, €9/month, and €90/year.
//
// Related guides: [Set up a subscription](https://stripe.com/docs/billing/subscriptions/set-up-subscription) and more about [products and prices](https://stripe.com/docs/billing/prices-guide).
type Plan struct {
	APIResource
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
	Object          string              `json:"object"`
	Product         *Product            `json:"product"`
	Tiers           []*PlanTier         `json:"tiers"`
	TiersMode       string              `json:"tiers_mode"`
	TransformUsage  *PlanTransformUsage `json:"transform_usage"`
	TrialPeriodDays int64               `json:"trial_period_days"`
	UsageType       PlanUsageType       `json:"usage_type"`
}

// PlanList is a list of Plans as retrieved from a list endpoint.
type PlanList struct {
	APIResource
	ListMeta
	Data []*Plan `json:"data"`
}

// UnmarshalJSON handles deserialization of a Plan.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *Plan) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type plan Plan
	var v plan
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = Plan(v)
	return nil
}
