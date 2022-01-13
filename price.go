//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v72/form"
)

// Describes how to compute the price per period. Either `per_unit` or `tiered`. `per_unit` indicates that the fixed amount (specified in `unit_amount` or `unit_amount_decimal`) will be charged per unit in `quantity` (for prices with `usage_type=licensed`), or per unit of total usage (for prices with `usage_type=metered`). `tiered` indicates that the unit pricing will be computed using a tiering strategy as defined using the `tiers` and `tiers_mode` attributes.
type PriceBillingScheme string

// List of values that PriceBillingScheme can take
const (
	PriceBillingSchemePerUnit PriceBillingScheme = "per_unit"
	PriceBillingSchemeTiered  PriceBillingScheme = "tiered"
)

// Specifies a usage aggregation strategy for prices of `usage_type=metered`. Allowed values are `sum` for summing up all usage during a period, `last_during_period` for using the last usage record reported within a period, `last_ever` for using the last usage record ever (across period bounds) or `max` which uses the usage record with the maximum reported usage during a period. Defaults to `sum`.
type PriceRecurringAggregateUsage string

// List of values that PriceRecurringAggregateUsage can take
const (
	PriceRecurringAggregateUsageLastDuringPeriod PriceRecurringAggregateUsage = "last_during_period"
	PriceRecurringAggregateUsageLastEver         PriceRecurringAggregateUsage = "last_ever"
	PriceRecurringAggregateUsageMax              PriceRecurringAggregateUsage = "max"
	PriceRecurringAggregateUsageSum              PriceRecurringAggregateUsage = "sum"
)

// The frequency at which a subscription is billed. One of `day`, `week`, `month` or `year`.
type PriceRecurringInterval string

// List of values that PriceRecurringInterval can take
const (
	PriceRecurringIntervalDay   PriceRecurringInterval = "day"
	PriceRecurringIntervalMonth PriceRecurringInterval = "month"
	PriceRecurringIntervalWeek  PriceRecurringInterval = "week"
	PriceRecurringIntervalYear  PriceRecurringInterval = "year"
)

// Configures how the quantity per period should be determined. Can be either `metered` or `licensed`. `licensed` automatically bills the `quantity` set when adding it to a subscription. `metered` aggregates the total usage based on usage records. Defaults to `licensed`.
type PriceRecurringUsageType string

// List of values that PriceRecurringUsageType can take
const (
	PriceRecurringUsageTypeLicensed PriceRecurringUsageType = "licensed"
	PriceRecurringUsageTypeMetered  PriceRecurringUsageType = "metered"
)

// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
type PriceTaxBehavior string

// List of values that PriceTaxBehavior can take
const (
	PriceTaxBehaviorExclusive   PriceTaxBehavior = "exclusive"
	PriceTaxBehaviorInclusive   PriceTaxBehavior = "inclusive"
	PriceTaxBehaviorUnspecified PriceTaxBehavior = "unspecified"
)

// Defines if the tiering price should be `graduated` or `volume` based. In `volume`-based tiering, the maximum quantity within a period determines the per unit price. In `graduated` tiering, pricing can change as the quantity grows.
type PriceTiersMode string

// List of values that PriceTiersMode can take
const (
	PriceTiersModeGraduated PriceTiersMode = "graduated"
	PriceTiersModeVolume    PriceTiersMode = "volume"
)

// After division, either round the result `up` or `down`.
type PriceTransformQuantityRound string

// List of values that PriceTransformQuantityRound can take
const (
	PriceTransformQuantityRoundDown PriceTransformQuantityRound = "down"
	PriceTransformQuantityRoundUp   PriceTransformQuantityRound = "up"
)

// One of `one_time` or `recurring` depending on whether the price is for a one-time purchase or a recurring (subscription) purchase.
type PriceType string

// List of values that PriceType can take
const (
	PriceTypeOneTime   PriceType = "one_time"
	PriceTypeRecurring PriceType = "recurring"
)

// Only return prices with these recurring fields.
type PriceRecurringListParams struct {
	Interval  *string `form:"interval"`
	UsageType *string `form:"usage_type"`
}

// Returns a list of your prices.
type PriceListParams struct {
	ListParams   `form:"*"`
	Active       *bool                     `form:"active"`
	Created      *int64                    `form:"created"`
	CreatedRange *RangeQueryParams         `form:"created"`
	Currency     *string                   `form:"currency"`
	LookupKeys   []*string                 `form:"lookup_keys"`
	Product      *string                   `form:"product"`
	Recurring    *PriceRecurringListParams `form:"recurring"`
	Type         *string                   `form:"type"`
}

// These fields can be used to create a new product that this price will belong to.
type PriceProductDataParams struct {
	Active              *bool             `form:"active"`
	ID                  *string           `form:"id"`
	Metadata            map[string]string `form:"metadata"`
	Name                *string           `form:"name"`
	StatementDescriptor *string           `form:"statement_descriptor"`
	TaxCode             *string           `form:"tax_code"`
	UnitLabel           *string           `form:"unit_label"`
}

// The recurring components of a price such as `interval` and `usage_type`.
type PriceRecurringParams struct {
	AggregateUsage  *string `form:"aggregate_usage"`
	Interval        *string `form:"interval"`
	IntervalCount   *int64  `form:"interval_count"`
	TrialPeriodDays *int64  `form:"trial_period_days"`
	UsageType       *string `form:"usage_type"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceTierParams struct {
	FlatAmount        *int64   `form:"flat_amount"`
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	UnitAmount        *int64   `form:"unit_amount"`
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	UpTo              *int64   `form:"up_to"`
	UpToInf           *bool    `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for PriceTierParams.
func (p *PriceTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// Apply a transformation to the reported usage or set quantity before computing the billed price. Cannot be combined with `tiers`.
type PriceTransformQuantityParams struct {
	DivideBy *int64  `form:"divide_by"`
	Round    *string `form:"round"`
}

// Creates a new price for an existing product. The price can be recurring or one-time.
type PriceParams struct {
	Params            `form:"*"`
	Active            *bool                         `form:"active"`
	BillingScheme     *string                       `form:"billing_scheme"`
	Currency          *string                       `form:"currency"`
	LookupKey         *string                       `form:"lookup_key"`
	Nickname          *string                       `form:"nickname"`
	Product           *string                       `form:"product"`
	ProductData       *PriceProductDataParams       `form:"product_data"`
	Recurring         *PriceRecurringParams         `form:"recurring"`
	TaxBehavior       *string                       `form:"tax_behavior"`
	Tiers             []*PriceTierParams            `form:"tiers"`
	TiersMode         *string                       `form:"tiers_mode"`
	TransferLookupKey *bool                         `form:"transfer_lookup_key"`
	TransformQuantity *PriceTransformQuantityParams `form:"transform_quantity"`
	UnitAmount        *int64                        `form:"unit_amount"`
	UnitAmountDecimal *float64                      `form:"unit_amount_decimal,high_precision"`
}

// The recurring components of a price such as `interval` and `usage_type`.
type PriceRecurring struct {
	AggregateUsage  PriceRecurringAggregateUsage `json:"aggregate_usage"`
	Interval        PriceRecurringInterval       `json:"interval"`
	IntervalCount   int64                        `json:"interval_count"`
	TrialPeriodDays int64                        `json:"trial_period_days"`
	UsageType       PriceRecurringUsageType      `json:"usage_type"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type PriceTier struct {
	FlatAmount        int64   `json:"flat_amount"`
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	UnitAmount        int64   `json:"unit_amount"`
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	UpTo              int64   `json:"up_to"`
}

// Apply a transformation to the reported usage or set quantity before computing the amount billed. Cannot be combined with `tiers`.
type PriceTransformQuantity struct {
	DivideBy int64                       `json:"divide_by"`
	Round    PriceTransformQuantityRound `json:"round"`
}

// Prices define the unit cost, currency, and (optional) billing cycle for both recurring and one-time purchases of products.
// [Products](https://stripe.com/docs/api#products) help you track inventory or provisioning, and prices help you track payment terms. Different physical goods or levels of service should be represented by products, and pricing options should be represented by prices. This approach lets you change prices without having to change your provisioning scheme.
//
// For example, you might have a single "gold" product that has prices for $10/month, $100/year, and €9 once.
//
// Related guides: [Set up a subscription](https://stripe.com/docs/billing/subscriptions/set-up-subscription), [create an invoice](https://stripe.com/docs/billing/invoices/create), and more about [products and prices](https://stripe.com/docs/billing/prices-guide).
type Price struct {
	APIResource
	Active            bool                    `json:"active"`
	BillingScheme     PriceBillingScheme      `json:"billing_scheme"`
	Created           int64                   `json:"created"`
	Currency          Currency                `json:"currency"`
	Deleted           bool                    `json:"deleted"`
	ID                string                  `json:"id"`
	Livemode          bool                    `json:"livemode"`
	LookupKey         string                  `json:"lookup_key"`
	Metadata          map[string]string       `json:"metadata"`
	Nickname          string                  `json:"nickname"`
	Object            string                  `json:"object"`
	Product           *Product                `json:"product"`
	Recurring         *PriceRecurring         `json:"recurring"`
	TaxBehavior       PriceTaxBehavior        `json:"tax_behavior"`
	Tiers             []*PriceTier            `json:"tiers"`
	TiersMode         PriceTiersMode          `json:"tiers_mode"`
	TransformQuantity *PriceTransformQuantity `json:"transform_quantity"`
	Type              PriceType               `json:"type"`
	UnitAmount        int64                   `json:"unit_amount"`
	UnitAmountDecimal float64                 `json:"unit_amount_decimal,string"`
}

// PriceList is a list of Prices as retrieved from a list endpoint.
type PriceList struct {
	APIResource
	ListMeta
	Data []*Price `json:"data"`
}

// UnmarshalJSON handles deserialization of a Price.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *Price) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type price Price
	var v price
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = Price(v)
	return nil
}
