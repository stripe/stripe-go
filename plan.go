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
	ListParams `form:"*"`
	// Only return plans that are active or inactive (e.g., pass `false` to list all inactive plans).
	Active *bool `form:"active"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	Created *int64 `form:"created"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return plans for the given product.
	Product *string `form:"product"`
}

// The product the plan belongs to. This cannot be changed once it has been used in a subscription or subscription schedule.
type PlanProductParams struct {
	// Whether the product is currently available for purchase. Defaults to `true`.
	Active *bool `form:"active"`
	// The identifier for the product. Must be unique. If not provided, an identifier will be randomly generated.
	ID *string `form:"id"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The product's name, meant to be displayable to the customer.
	Name *string `form:"name"`
	// An arbitrary string to be displayed on your customer's credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all.
	//
	// This may be up to 22 characters. The statement description may not include `<`, `>`, `\`, `"`, `'` characters, and will appear on your customer's statement in capital letters. Non-ASCII characters are automatically stripped.
	StatementDescriptor *string `form:"statement_descriptor"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code"`
	// A label that represents units of this product in Stripe and on customers' receipts and invoices. When set, this will be included in associated invoice line item descriptions.
	UnitLabel *string `form:"unit_label"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PlanTierParams struct {
	Params `form:"*"`
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"-"` // See custom AppendTo
	UpToInf *bool  `form:"-"` // See custom AppendTo
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
	// Divide usage by this number.
	DivideBy *int64 `form:"divide_by"`
	// After division, either round the result `up` or `down`.
	Round *string `form:"round"`
}

// You can now model subscriptions more flexibly using the [Prices API](https://stripe.com/docs/api#prices). It replaces the Plans API and is backwards compatible to simplify your migration.
type PlanParams struct {
	Params `form:"*"`
	// Whether the plan is currently available for new subscriptions.
	Active *bool `form:"active"`
	// Specifies a usage aggregation strategy for plans of `usage_type=metered`. Allowed values are `sum` for summing up all usage during a period, `last_during_period` for using the last usage record reported within a period, `last_ever` for using the last usage record ever (across period bounds) or `max` which uses the usage record with the maximum reported usage during a period. Defaults to `sum`.
	AggregateUsage *string `form:"aggregate_usage"`
	// A positive integer in cents (or local equivalent) (or 0 for a free plan) representing how much to charge on a recurring basis.
	Amount *int64 `form:"amount"`
	// Same as `amount`, but accepts a decimal value with at most 12 decimal places. Only one of `amount` and `amount_decimal` can be set.
	AmountDecimal *float64 `form:"amount_decimal,high_precision"`
	// Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `amount`) will be charged per unit in `quantity` (for plans with `usage_type=licensed`), or per unit of total usage (for plans with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.
	BillingScheme *string `form:"billing_scheme"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// An identifier randomly generated by Stripe. Used to identify this plan when subscribing a customer. You can optionally override this ID, but the ID must be unique across all plans in your Stripe account. You can, however, use the same plan ID in both live and test modes.
	ID *string `form:"id"`
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks).
	IntervalCount *int64 `form:"interval_count"`
	// A brief description of the plan, hidden from customers.
	Nickname  *string            `form:"nickname"`
	Product   *PlanProductParams `form:"product"`
	ProductID *string            `form:"product"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PlanTierParams `form:"tiers"`
	// Defines if the tiering price should be `graduated` or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per unit price, in `graduated` tiering pricing can successively change as the quantity grows.
	TiersMode *string `form:"tiers_mode"`
	// Apply a transformation to the reported usage or set quantity before computing the billed price. Cannot be combined with `tiers`.
	TransformUsage *PlanTransformUsageParams `form:"transform_usage"`
	// Default number of trial days when subscribing a customer to this plan using [`trial_from_plan=true`](https://stripe.com/docs/api#create_subscription-trial_from_plan).
	TrialPeriodDays *int64 `form:"trial_period_days"`
	// Configures how the quantity per period should be determined. Can be either `metered` or `licensed`. `licensed` automatically bills the `quantity` set when adding it to a subscription. `metered` aggregates the total usage based on usage records. Defaults to `licensed`.
	UsageType *string `form:"usage_type"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PlanTier struct {
	// Price for the entire tier.
	FlatAmount int64 `json:"flat_amount"`
	// Same as `flat_amount`, but contains a decimal value with at most 12 decimal places.
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	// Per unit price for units relevant to the tier.
	UnitAmount int64 `json:"unit_amount"`
	// Same as `unit_amount`, but contains a decimal value with at most 12 decimal places.
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	// Up to and including to this quantity will be contained in the tier.
	UpTo int64 `json:"up_to"`
}

// Apply a transformation to the reported usage or set quantity before computing the amount billed. Cannot be combined with `tiers`.
type PlanTransformUsage struct {
	// Divide usage by this number.
	DivideBy int64 `json:"divide_by"`
	// After division, either round the result `up` or `down`.
	Round PlanTransformUsageRound `json:"round"`
}

// You can now model subscriptions more flexibly using the [Prices API](https://stripe.com/docs/api#prices). It replaces the Plans API and is backwards compatible to simplify your migration.
//
// Plans define the base price, currency, and billing cycle for recurring purchases of products.
// [Products](https://stripe.com/docs/api#products) help you track inventory or provisioning, and plans help you track pricing. Different physical goods or levels of service should be represented by products, and pricing options should be represented by plans. This approach lets you change prices without having to change your provisioning scheme.
//
// For example, you might have a single "gold" product that has plans for $10/month, $100/year, €9/month, and €90/year.
//
// Related guides: [Set up a subscription](https://stripe.com/docs/billing/subscriptions/set-up-subscription) and more about [products and prices](https://stripe.com/docs/products-prices/overview).
type Plan struct {
	APIResource
	// Whether the plan can be used for new purchases.
	Active bool `json:"active"`
	// Specifies a usage aggregation strategy for plans of `usage_type=metered`. Allowed values are `sum` for summing up all usage during a period, `last_during_period` for using the last usage record reported within a period, `last_ever` for using the last usage record ever (across period bounds) or `max` which uses the usage record with the maximum reported usage during a period. Defaults to `sum`.
	AggregateUsage string `json:"aggregate_usage"`
	// The unit amount in %s to be charged, represented as a whole integer if possible. Only set if `billing_scheme=per_unit`.
	Amount int64 `json:"amount"`
	// The unit amount in %s to be charged, represented as a decimal string with at most 12 decimal places. Only set if `billing_scheme=per_unit`.
	AmountDecimal float64 `json:"amount_decimal,string"`
	// Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `amount`) will be charged per unit in `quantity` (for plans with `usage_type=licensed`), or per unit of total usage (for plans with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.
	BillingScheme PlanBillingScheme `json:"billing_scheme"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	Deleted  bool     `json:"deleted"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.
	Interval PlanInterval `json:"interval"`
	// The number of intervals (specified in the `interval` attribute) between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months.
	IntervalCount int64 `json:"interval_count"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// A brief description of the plan, hidden from customers.
	Nickname string `json:"nickname"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The product whose pricing this plan determines.
	Product *Product `json:"product"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*PlanTier `json:"tiers"`
	// Defines if the tiering price should be `graduated` or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per unit price. In `graduated` tiering, pricing can change as the quantity grows.
	TiersMode string `json:"tiers_mode"`
	// Apply a transformation to the reported usage or set quantity before computing the amount billed. Cannot be combined with `tiers`.
	TransformUsage *PlanTransformUsage `json:"transform_usage"`
	// Default number of trial days when subscribing a customer to this plan using [`trial_from_plan=true`](https://stripe.com/docs/api#create_subscription-trial_from_plan).
	TrialPeriodDays int64 `json:"trial_period_days"`
	// Configures how the quantity per period should be determined. Can be either `metered` or `licensed`. `licensed` automatically bills the `quantity` set when adding it to a subscription. `metered` aggregates the total usage based on usage records. Defaults to `licensed`.
	UsageType PlanUsageType `json:"usage_type"`
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
