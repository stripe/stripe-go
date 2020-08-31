package stripe

import (
	"encoding/json"
	"strconv"

	"github.com/stripe/stripe-go/v72/form"
)

// PriceBillingScheme is the list of allowed values for a price's billing scheme.
type PriceBillingScheme string

// List of values that PriceBillingScheme can take.
const (
	PriceBillingSchemePerUnit PriceBillingScheme = "per_unit"
	PriceBillingSchemeTiered  PriceBillingScheme = "tiered"
)

// PriceRecurringAggregateUsage is the list of allowed values for a price's aggregate usage.
type PriceRecurringAggregateUsage string

// List of values that PriceRecurringAggregateUsage can take.
const (
	PriceRecurringAggregateUsageLastDuringPeriod PriceRecurringAggregateUsage = "last_during_period"
	PriceRecurringAggregateUsageLastEver         PriceRecurringAggregateUsage = "last_ever"
	PriceRecurringAggregateUsageMax              PriceRecurringAggregateUsage = "max"
	PriceRecurringAggregateUsageSum              PriceRecurringAggregateUsage = "sum"
)

// PriceRecurringInterval is the list of allowed values for a price's recurring interval.
type PriceRecurringInterval string

// List of values that PriceRecurringInterval can take.
const (
	PriceRecurringIntervalDay   PriceRecurringInterval = "day"
	PriceRecurringIntervalWeek  PriceRecurringInterval = "week"
	PriceRecurringIntervalMonth PriceRecurringInterval = "month"
	PriceRecurringIntervalYear  PriceRecurringInterval = "year"
)

// PriceRecurringUsageType is the list of allowed values for a price's usage type.
type PriceRecurringUsageType string

// List of values that PriceUsageType can take.
const (
	PriceRecurringUsageTypeLicensed PriceRecurringUsageType = "licensed"
	PriceRecurringUsageTypeMetered  PriceRecurringUsageType = "metered"
)

// PriceTiersMode is the list of allowed values for a price's tiers mode.
type PriceTiersMode string

// List of values that PriceTiersMode can take.
const (
	PriceTiersModeGraduated PriceTiersMode = "graduated"
	PriceTiersModeVolume    PriceTiersMode = "volume"
)

// PriceTransformQuantityRound is the list of allowed values for a price's transform quantity rounding logic.
type PriceTransformQuantityRound string

// List of values that PriceTransformQuantityRound can take.
const (
	PriceTransformQuantityRoundDown PriceTransformQuantityRound = "down"
	PriceTransformQuantityRoundUp   PriceTransformQuantityRound = "up"
)

// PriceType is the list of allowed values for a price's type.
type PriceType string

// List of values that PriceType can take.
const (
	PriceTypeOneTime   PriceType = "one_time"
	PriceTypeRecurring PriceType = "recurring"
)

// PriceProductDataParams is the set of parameters that can be used when creating a product inside a price.
type PriceProductDataParams struct {
	Active              *bool             `form:"active"`
	ID                  *string           `form:"id"`
	Name                *string           `form:"name"`
	Metadata            map[string]string `form:"metadata"`
	StatementDescriptor *string           `form:"statement_descriptor"`
	UnitLabel           *string           `form:"unit_label"`
}

// PriceRecurringParams is the set of parameters for the recurring components of a price.
type PriceRecurringParams struct {
	AggregateUsage  *string `form:"aggregate_usage"`
	Interval        *string `form:"interval"`
	IntervalCount   *int64  `form:"interval_count"`
	TrialPeriodDays *int64  `form:"trial_period_days"`
	UsageType       *string `form:"usage_type"`
}

// PriceTierParams configures tiered pricing
type PriceTierParams struct {
	FlatAmount        *int64   `form:"flat_amount"`
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	UnitAmount        *int64   `form:"unit_amount"`
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	UpTo              *int64   `form:"-"` // handled in custom AppendTo
	UpToInf           *bool    `form:"-"` // handled in custom AppendTo
}

// AppendTo implements custom up_to serialisation logic for tiers configuration
func (p *PriceTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	} else {
		body.Add(form.FormatKey(append(keyParts, "up_to")), strconv.FormatInt(Int64Value(p.UpTo), 10))
	}
}

// PriceTransformQuantityParams represents the parameter controlling the calculation of the final amount based on the quantity.
type PriceTransformQuantityParams struct {
	DivideBy *int64  `form:"divide_by"`
	Round    *string `form:"round"`
}

// PriceParams is the set of parameters that can be used when creating or updating a price.
// For more details see https://stripe.com/docs/api#create_price and https://stripe.com/docs/api#update_price.
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
	Tiers             []*PriceTierParams            `form:"tiers"`
	TiersMode         *string                       `form:"tiers_mode"`
	TransferLookupKey *bool                         `form:"transfer_lookup_key"`
	TransformQuantity *PriceTransformQuantityParams `form:"transform_quantity"`
	UnitAmount        *int64                        `form:"unit_amount"`
	UnitAmountDecimal *float64                      `form:"unit_amount_decimal,high_precision"`
}

// PriceRecurringListParams is the set of parameters that can be used when listing recurring prices.
type PriceRecurringListParams struct {
	Interval  *string `form:"interval"`
	UsageType *string `form:"usage_type"`
}

// PriceListParams is the set of parameters that can be used when listing prices.
// For more details see https://stripe.com/docs/api#list_prices.
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

// PriceRecurring represents the recurring components of a price.
type PriceRecurring struct {
	AggregateUsage  PriceRecurringAggregateUsage `json:"aggregate_usage"`
	Interval        PriceRecurringInterval       `json:"interval"`
	IntervalCount   int64                        `json:"interval_count"`
	TrialPeriodDays int64                        `json:"trial_period_days"`
	UsageType       PriceRecurringUsageType      `json:"usage_type"`
}

// PriceTier configures tiered pricing
type PriceTier struct {
	FlatAmount        int64   `json:"flat_amount"`
	FlatAmountDecimal float64 `json:"flat_amount_decimal,string"`
	UnitAmount        int64   `json:"unit_amount"`
	UnitAmountDecimal float64 `json:"unit_amount_decimal,string"`
	UpTo              int64   `json:"up_to"`
}

// PriceTransformQuantity represents how to calculate the final amount based on the quantity.
type PriceTransformQuantity struct {
	DivideBy int64                       `json:"divide_by"`
	Round    PriceTransformQuantityRound `json:"round"`
}

// Price is the resource representing a Stripe price.
// For more details see https://stripe.com/docs/api#prices.
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
	Tiers             []*PriceTier            `json:"tiers"`
	TiersMode         PriceTiersMode          `json:"tiers_mode"`
	TransformQuantity *PriceTransformQuantity `json:"transform_quantity"`
	Type              PriceType               `json:"type"`
	UnitAmount        int64                   `json:"unit_amount"`
	UnitAmountDecimal float64                 `json:"unit_amount_decimal,string"`
}

// PriceList is a list of prices as returned from a list endpoint.
type PriceList struct {
	APIResource
	ListMeta
	Data []*Price `json:"data"`
}

// UnmarshalJSON handles deserialization of a Price.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *Price) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type price Price
	var v price
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = Price(v)
	return nil
}
