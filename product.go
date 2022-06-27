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

// The type of the product. The product is either of type `good`, which is eligible for use with Orders and SKUs, or `service`, which is eligible for use with Subscriptions and Plans.
type ProductType string

// List of values that ProductType can take
const (
	ProductTypeGood    ProductType = "good"
	ProductTypeService ProductType = "service"
)

// Search for products you've previously created using Stripe's [Search Query Language](https://stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
type ProductSearchParams struct {
	SearchParams `form:"*"`
	// A cursor for pagination across multiple pages of results. Don't include this parameter on the first call. Use the next_page value returned in a previous response to request subsequent results.
	Page *string `form:"page"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsAedCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsAedTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsAedTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsAedTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in AED.
type ProductDefaultPriceDataCurrencyOptionsAedParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsAedCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsAedTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsAfnCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsAfnTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsAfnTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsAfnTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in AFN.
type ProductDefaultPriceDataCurrencyOptionsAfnParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsAfnCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsAfnTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsAllCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsAllTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsAllTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsAllTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ALL.
type ProductDefaultPriceDataCurrencyOptionsAllParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsAllCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsAllTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsAmdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsAmdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsAmdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsAmdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in AMD.
type ProductDefaultPriceDataCurrencyOptionsAmdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsAmdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsAmdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsAngCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsAngTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsAngTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsAngTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ANG.
type ProductDefaultPriceDataCurrencyOptionsAngParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsAngCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsAngTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsAoaCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsAoaTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsAoaTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsAoaTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in AOA.
type ProductDefaultPriceDataCurrencyOptionsAoaParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsAoaCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsAoaTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsArsCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsArsTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsArsTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsArsTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ARS.
type ProductDefaultPriceDataCurrencyOptionsArsParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsArsCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsArsTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsAUDCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsAUDTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsAUDTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsAUDTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in AUD.
type ProductDefaultPriceDataCurrencyOptionsAUDParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsAUDCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsAUDTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsAwgCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsAwgTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsAwgTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsAwgTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in AWG.
type ProductDefaultPriceDataCurrencyOptionsAwgParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsAwgCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsAwgTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsAznCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsAznTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsAznTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsAznTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in AZN.
type ProductDefaultPriceDataCurrencyOptionsAznParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsAznCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsAznTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsBamCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsBamTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsBamTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsBamTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BAM.
type ProductDefaultPriceDataCurrencyOptionsBamParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsBamCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsBamTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsBbdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsBbdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsBbdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsBbdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BBD.
type ProductDefaultPriceDataCurrencyOptionsBbdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsBbdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsBbdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsBdtCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsBdtTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsBdtTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsBdtTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BDT.
type ProductDefaultPriceDataCurrencyOptionsBdtParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsBdtCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsBdtTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsBgnCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsBgnTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsBgnTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsBgnTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BGN.
type ProductDefaultPriceDataCurrencyOptionsBgnParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsBgnCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsBgnTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsBhdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsBhdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsBhdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsBhdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BHD.
type ProductDefaultPriceDataCurrencyOptionsBhdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsBhdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsBhdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsBifCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsBifTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsBifTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsBifTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BIF.
type ProductDefaultPriceDataCurrencyOptionsBifParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsBifCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsBifTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsBmdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsBmdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsBmdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsBmdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BMD.
type ProductDefaultPriceDataCurrencyOptionsBmdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsBmdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsBmdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsBndCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsBndTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsBndTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsBndTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BND.
type ProductDefaultPriceDataCurrencyOptionsBndParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsBndCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsBndTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsBobCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsBobTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsBobTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsBobTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BOB.
type ProductDefaultPriceDataCurrencyOptionsBobParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsBobCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsBobTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsBrlCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsBrlTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsBrlTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsBrlTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BRL.
type ProductDefaultPriceDataCurrencyOptionsBrlParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsBrlCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsBrlTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsBsdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsBsdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsBsdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsBsdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BSD.
type ProductDefaultPriceDataCurrencyOptionsBsdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsBsdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsBsdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsBtnCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsBtnTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsBtnTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsBtnTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BTN.
type ProductDefaultPriceDataCurrencyOptionsBtnParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsBtnCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsBtnTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsBwpCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsBwpTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsBwpTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsBwpTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BWP.
type ProductDefaultPriceDataCurrencyOptionsBwpParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsBwpCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsBwpTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsBynCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsBynTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsBynTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsBynTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BYN.
type ProductDefaultPriceDataCurrencyOptionsBynParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsBynCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsBynTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsBzdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsBzdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsBzdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsBzdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in BZD.
type ProductDefaultPriceDataCurrencyOptionsBzdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsBzdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsBzdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsCADCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsCADTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsCADTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsCADTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CAD.
type ProductDefaultPriceDataCurrencyOptionsCADParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsCADCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsCADTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsCdfCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsCdfTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsCdfTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsCdfTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CDF.
type ProductDefaultPriceDataCurrencyOptionsCdfParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsCdfCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsCdfTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsCHFCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsCHFTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsCHFTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsCHFTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CHF.
type ProductDefaultPriceDataCurrencyOptionsCHFParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsCHFCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsCHFTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsClpCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsClpTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsClpTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsClpTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CLP.
type ProductDefaultPriceDataCurrencyOptionsClpParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsClpCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsClpTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsCnyCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsCnyTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsCnyTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsCnyTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CNY.
type ProductDefaultPriceDataCurrencyOptionsCnyParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsCnyCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsCnyTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsCopCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsCopTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsCopTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsCopTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in COP.
type ProductDefaultPriceDataCurrencyOptionsCopParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsCopCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsCopTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsCrcCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsCrcTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsCrcTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsCrcTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CRC.
type ProductDefaultPriceDataCurrencyOptionsCrcParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsCrcCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsCrcTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsCveCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsCveTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsCveTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsCveTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CVE.
type ProductDefaultPriceDataCurrencyOptionsCveParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsCveCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsCveTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsCZKCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsCZKTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsCZKTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsCZKTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in CZK.
type ProductDefaultPriceDataCurrencyOptionsCZKParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsCZKCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsCZKTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsDjfCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsDjfTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsDjfTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsDjfTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in DJF.
type ProductDefaultPriceDataCurrencyOptionsDjfParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsDjfCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsDjfTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsDKKCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsDKKTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsDKKTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsDKKTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in DKK.
type ProductDefaultPriceDataCurrencyOptionsDKKParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsDKKCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsDKKTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsDopCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsDopTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsDopTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsDopTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in DOP.
type ProductDefaultPriceDataCurrencyOptionsDopParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsDopCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsDopTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsDzdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsDzdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsDzdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsDzdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in DZD.
type ProductDefaultPriceDataCurrencyOptionsDzdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsDzdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsDzdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsEgpCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsEgpTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsEgpTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsEgpTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in EGP.
type ProductDefaultPriceDataCurrencyOptionsEgpParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsEgpCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsEgpTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsEtbCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsEtbTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsEtbTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsEtbTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ETB.
type ProductDefaultPriceDataCurrencyOptionsEtbParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsEtbCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsEtbTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsEURCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsEURTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsEURTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsEURTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in EUR.
type ProductDefaultPriceDataCurrencyOptionsEURParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsEURCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsEURTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsFjdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsFjdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsFjdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsFjdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in FJD.
type ProductDefaultPriceDataCurrencyOptionsFjdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsFjdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsFjdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsFkpCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsFkpTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsFkpTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsFkpTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in FKP.
type ProductDefaultPriceDataCurrencyOptionsFkpParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsFkpCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsFkpTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsGBPCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsGBPTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsGBPTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsGBPTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GBP.
type ProductDefaultPriceDataCurrencyOptionsGBPParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsGBPCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsGBPTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsGelCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsGelTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsGelTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsGelTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GEL.
type ProductDefaultPriceDataCurrencyOptionsGelParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsGelCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsGelTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsGhsCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsGhsTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsGhsTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsGhsTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GHS.
type ProductDefaultPriceDataCurrencyOptionsGhsParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsGhsCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsGhsTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsGipCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsGipTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsGipTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsGipTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GIP.
type ProductDefaultPriceDataCurrencyOptionsGipParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsGipCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsGipTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsGmdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsGmdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsGmdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsGmdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GMD.
type ProductDefaultPriceDataCurrencyOptionsGmdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsGmdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsGmdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsGnfCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsGnfTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsGnfTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsGnfTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GNF.
type ProductDefaultPriceDataCurrencyOptionsGnfParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsGnfCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsGnfTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsGtqCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsGtqTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsGtqTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsGtqTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GTQ.
type ProductDefaultPriceDataCurrencyOptionsGtqParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsGtqCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsGtqTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsGydCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsGydTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsGydTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsGydTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in GYD.
type ProductDefaultPriceDataCurrencyOptionsGydParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsGydCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsGydTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsHKDCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsHKDTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsHKDTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsHKDTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in HKD.
type ProductDefaultPriceDataCurrencyOptionsHKDParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsHKDCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsHKDTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsHnlCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsHnlTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsHnlTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsHnlTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in HNL.
type ProductDefaultPriceDataCurrencyOptionsHnlParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsHnlCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsHnlTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsHrkCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsHrkTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsHrkTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsHrkTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in HRK.
type ProductDefaultPriceDataCurrencyOptionsHrkParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsHrkCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsHrkTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsHtgCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsHtgTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsHtgTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsHtgTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in HTG.
type ProductDefaultPriceDataCurrencyOptionsHtgParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsHtgCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsHtgTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsHufCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsHufTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsHufTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsHufTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in HUF.
type ProductDefaultPriceDataCurrencyOptionsHufParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsHufCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsHufTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsIdrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsIdrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsIdrTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsIdrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in IDR.
type ProductDefaultPriceDataCurrencyOptionsIdrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsIdrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsIdrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsIlsCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsIlsTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsIlsTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsIlsTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ILS.
type ProductDefaultPriceDataCurrencyOptionsIlsParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsIlsCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsIlsTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsInrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsInrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsInrTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsInrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in INR.
type ProductDefaultPriceDataCurrencyOptionsInrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsInrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsInrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsIskCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsIskTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsIskTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsIskTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ISK.
type ProductDefaultPriceDataCurrencyOptionsIskParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsIskCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsIskTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsJmdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsJmdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsJmdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsJmdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in JMD.
type ProductDefaultPriceDataCurrencyOptionsJmdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsJmdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsJmdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsJodCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsJodTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsJodTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsJodTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in JOD.
type ProductDefaultPriceDataCurrencyOptionsJodParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsJodCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsJodTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsJpyCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsJpyTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsJpyTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsJpyTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in JPY.
type ProductDefaultPriceDataCurrencyOptionsJpyParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsJpyCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsJpyTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsKesCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsKesTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsKesTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsKesTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KES.
type ProductDefaultPriceDataCurrencyOptionsKesParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsKesCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsKesTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsKgsCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsKgsTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsKgsTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsKgsTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KGS.
type ProductDefaultPriceDataCurrencyOptionsKgsParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsKgsCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsKgsTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsKhrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsKhrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsKhrTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsKhrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KHR.
type ProductDefaultPriceDataCurrencyOptionsKhrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsKhrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsKhrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsKmfCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsKmfTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsKmfTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsKmfTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KMF.
type ProductDefaultPriceDataCurrencyOptionsKmfParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsKmfCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsKmfTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsKrwCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsKrwTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsKrwTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsKrwTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KRW.
type ProductDefaultPriceDataCurrencyOptionsKrwParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsKrwCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsKrwTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsKwdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsKwdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsKwdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsKwdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KWD.
type ProductDefaultPriceDataCurrencyOptionsKwdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsKwdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsKwdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsKydCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsKydTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsKydTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsKydTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KYD.
type ProductDefaultPriceDataCurrencyOptionsKydParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsKydCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsKydTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsKztCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsKztTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsKztTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsKztTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in KZT.
type ProductDefaultPriceDataCurrencyOptionsKztParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsKztCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsKztTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsLakCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsLakTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsLakTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsLakTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in LAK.
type ProductDefaultPriceDataCurrencyOptionsLakParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsLakCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsLakTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsLbpCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsLbpTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsLbpTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsLbpTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in LBP.
type ProductDefaultPriceDataCurrencyOptionsLbpParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsLbpCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsLbpTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsLkrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsLkrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsLkrTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsLkrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in LKR.
type ProductDefaultPriceDataCurrencyOptionsLkrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsLkrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsLkrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsLrdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsLrdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsLrdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsLrdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in LRD.
type ProductDefaultPriceDataCurrencyOptionsLrdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsLrdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsLrdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsLslCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsLslTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsLslTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsLslTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in LSL.
type ProductDefaultPriceDataCurrencyOptionsLslParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsLslCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsLslTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsMadCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsMadTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsMadTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsMadTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MAD.
type ProductDefaultPriceDataCurrencyOptionsMadParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsMadCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsMadTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsMdlCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsMdlTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsMdlTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsMdlTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MDL.
type ProductDefaultPriceDataCurrencyOptionsMdlParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsMdlCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsMdlTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsMgaCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsMgaTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsMgaTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsMgaTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MGA.
type ProductDefaultPriceDataCurrencyOptionsMgaParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsMgaCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsMgaTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsMkdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsMkdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsMkdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsMkdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MKD.
type ProductDefaultPriceDataCurrencyOptionsMkdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsMkdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsMkdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsMmkCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsMmkTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsMmkTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsMmkTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MMK.
type ProductDefaultPriceDataCurrencyOptionsMmkParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsMmkCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsMmkTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsMntCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsMntTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsMntTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsMntTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MNT.
type ProductDefaultPriceDataCurrencyOptionsMntParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsMntCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsMntTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsMopCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsMopTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsMopTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsMopTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MOP.
type ProductDefaultPriceDataCurrencyOptionsMopParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsMopCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsMopTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsMroCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsMroTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsMroTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsMroTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MRO.
type ProductDefaultPriceDataCurrencyOptionsMroParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsMroCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsMroTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsMurCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsMurTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsMurTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsMurTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MUR.
type ProductDefaultPriceDataCurrencyOptionsMurParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsMurCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsMurTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsMvrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsMvrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsMvrTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsMvrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MVR.
type ProductDefaultPriceDataCurrencyOptionsMvrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsMvrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsMvrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsMwkCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsMwkTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsMwkTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsMwkTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MWK.
type ProductDefaultPriceDataCurrencyOptionsMwkParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsMwkCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsMwkTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsMxnCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsMxnTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsMxnTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsMxnTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MXN.
type ProductDefaultPriceDataCurrencyOptionsMxnParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsMxnCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsMxnTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsMYRCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsMYRTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsMYRTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsMYRTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MYR.
type ProductDefaultPriceDataCurrencyOptionsMYRParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsMYRCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsMYRTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsMznCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsMznTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsMznTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsMznTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in MZN.
type ProductDefaultPriceDataCurrencyOptionsMznParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsMznCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsMznTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsNadCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsNadTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsNadTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsNadTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in NAD.
type ProductDefaultPriceDataCurrencyOptionsNadParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsNadCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsNadTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsNgnCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsNgnTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsNgnTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsNgnTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in NGN.
type ProductDefaultPriceDataCurrencyOptionsNgnParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsNgnCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsNgnTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsNioCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsNioTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsNioTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsNioTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in NIO.
type ProductDefaultPriceDataCurrencyOptionsNioParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsNioCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsNioTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsNOKCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsNOKTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsNOKTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsNOKTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in NOK.
type ProductDefaultPriceDataCurrencyOptionsNOKParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsNOKCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsNOKTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsNprCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsNprTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsNprTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsNprTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in NPR.
type ProductDefaultPriceDataCurrencyOptionsNprParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsNprCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsNprTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsNZDCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsNZDTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsNZDTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsNZDTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in NZD.
type ProductDefaultPriceDataCurrencyOptionsNZDParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsNZDCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsNZDTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsOmrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsOmrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsOmrTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsOmrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in OMR.
type ProductDefaultPriceDataCurrencyOptionsOmrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsOmrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsOmrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsPabCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsPabTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsPabTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsPabTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in PAB.
type ProductDefaultPriceDataCurrencyOptionsPabParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsPabCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsPabTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsPenCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsPenTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsPenTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsPenTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in PEN.
type ProductDefaultPriceDataCurrencyOptionsPenParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsPenCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsPenTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsPgkCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsPgkTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsPgkTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsPgkTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in PGK.
type ProductDefaultPriceDataCurrencyOptionsPgkParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsPgkCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsPgkTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsPhpCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsPhpTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsPhpTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsPhpTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in PHP.
type ProductDefaultPriceDataCurrencyOptionsPhpParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsPhpCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsPhpTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsPkrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsPkrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsPkrTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsPkrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in PKR.
type ProductDefaultPriceDataCurrencyOptionsPkrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsPkrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsPkrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsPlnCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsPlnTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsPlnTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsPlnTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in PLN.
type ProductDefaultPriceDataCurrencyOptionsPlnParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsPlnCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsPlnTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsPygCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsPygTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsPygTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsPygTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in PYG.
type ProductDefaultPriceDataCurrencyOptionsPygParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsPygCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsPygTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsQarCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsQarTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsQarTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsQarTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in QAR.
type ProductDefaultPriceDataCurrencyOptionsQarParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsQarCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsQarTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsRonCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsRonTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsRonTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsRonTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in RON.
type ProductDefaultPriceDataCurrencyOptionsRonParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsRonCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsRonTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsRsdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsRsdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsRsdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsRsdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in RSD.
type ProductDefaultPriceDataCurrencyOptionsRsdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsRsdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsRsdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsRubCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsRubTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsRubTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsRubTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in RUB.
type ProductDefaultPriceDataCurrencyOptionsRubParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsRubCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsRubTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsRwfCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsRwfTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsRwfTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsRwfTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in RWF.
type ProductDefaultPriceDataCurrencyOptionsRwfParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsRwfCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsRwfTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsSarCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsSarTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsSarTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsSarTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SAR.
type ProductDefaultPriceDataCurrencyOptionsSarParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsSarCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsSarTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsSbdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsSbdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsSbdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsSbdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SBD.
type ProductDefaultPriceDataCurrencyOptionsSbdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsSbdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsSbdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsScrCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsScrTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsScrTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsScrTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SCR.
type ProductDefaultPriceDataCurrencyOptionsScrParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsScrCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsScrTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsSEKCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsSEKTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsSEKTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsSEKTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SEK.
type ProductDefaultPriceDataCurrencyOptionsSEKParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsSEKCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsSEKTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsSGDCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsSGDTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsSGDTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsSGDTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SGD.
type ProductDefaultPriceDataCurrencyOptionsSGDParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsSGDCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsSGDTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsShpCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsShpTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsShpTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsShpTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SHP.
type ProductDefaultPriceDataCurrencyOptionsShpParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsShpCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsShpTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsSllCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsSllTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsSllTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsSllTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SLL.
type ProductDefaultPriceDataCurrencyOptionsSllParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsSllCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsSllTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsSosCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsSosTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsSosTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsSosTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SOS.
type ProductDefaultPriceDataCurrencyOptionsSosParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsSosCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsSosTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsSrdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsSrdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsSrdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsSrdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SRD.
type ProductDefaultPriceDataCurrencyOptionsSrdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsSrdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsSrdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsStdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsStdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsStdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsStdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in STD.
type ProductDefaultPriceDataCurrencyOptionsStdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsStdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsStdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsSzlCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsSzlTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsSzlTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsSzlTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in SZL.
type ProductDefaultPriceDataCurrencyOptionsSzlParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsSzlCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsSzlTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsThbCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsThbTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsThbTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsThbTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in THB.
type ProductDefaultPriceDataCurrencyOptionsThbParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsThbCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsThbTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsTjsCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsTjsTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsTjsTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsTjsTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in TJS.
type ProductDefaultPriceDataCurrencyOptionsTjsParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsTjsCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsTjsTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsTndCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsTndTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsTndTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsTndTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in TND.
type ProductDefaultPriceDataCurrencyOptionsTndParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsTndCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsTndTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsTopCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsTopTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsTopTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsTopTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in TOP.
type ProductDefaultPriceDataCurrencyOptionsTopParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsTopCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsTopTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsTryCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsTryTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsTryTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsTryTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in TRY.
type ProductDefaultPriceDataCurrencyOptionsTryParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsTryCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsTryTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsTtdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsTtdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsTtdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsTtdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in TTD.
type ProductDefaultPriceDataCurrencyOptionsTtdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsTtdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsTtdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsTwdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsTwdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsTwdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsTwdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in TWD.
type ProductDefaultPriceDataCurrencyOptionsTwdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsTwdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsTwdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsTzsCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsTzsTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsTzsTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsTzsTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in TZS.
type ProductDefaultPriceDataCurrencyOptionsTzsParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsTzsCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsTzsTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsUahCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsUahTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsUahTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsUahTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in UAH.
type ProductDefaultPriceDataCurrencyOptionsUahParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsUahCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsUahTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsUgxCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsUgxTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsUgxTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsUgxTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in UGX.
type ProductDefaultPriceDataCurrencyOptionsUgxParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsUgxCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsUgxTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsUSDCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsUSDTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsUSDTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsUSDTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in USD.
type ProductDefaultPriceDataCurrencyOptionsUSDParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsUSDCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsUSDTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsUsdcCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsUsdcTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsUsdcTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsUsdcTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in USDC.
type ProductDefaultPriceDataCurrencyOptionsUsdcParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsUsdcCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsUsdcTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsUyuCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsUyuTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsUyuTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsUyuTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in UYU.
type ProductDefaultPriceDataCurrencyOptionsUyuParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsUyuCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsUyuTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsUzsCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsUzsTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsUzsTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsUzsTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in UZS.
type ProductDefaultPriceDataCurrencyOptionsUzsParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsUzsCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsUzsTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsVndCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsVndTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsVndTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsVndTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in VND.
type ProductDefaultPriceDataCurrencyOptionsVndParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsVndCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsVndTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsVuvCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsVuvTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsVuvTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsVuvTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in VUV.
type ProductDefaultPriceDataCurrencyOptionsVuvParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsVuvCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsVuvTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsWstCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsWstTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsWstTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsWstTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in WST.
type ProductDefaultPriceDataCurrencyOptionsWstParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsWstCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsWstTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsXafCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsXafTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsXafTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsXafTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in XAF.
type ProductDefaultPriceDataCurrencyOptionsXafParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsXafCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsXafTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsXcdCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsXcdTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsXcdTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsXcdTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in XCD.
type ProductDefaultPriceDataCurrencyOptionsXcdParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsXcdCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsXcdTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsXofCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsXofTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsXofTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsXofTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in XOF.
type ProductDefaultPriceDataCurrencyOptionsXofParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsXofCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsXofTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsXpfCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsXpfTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsXpfTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsXpfTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in XPF.
type ProductDefaultPriceDataCurrencyOptionsXpfParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsXpfCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsXpfTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsYerCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsYerTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsYerTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsYerTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in YER.
type ProductDefaultPriceDataCurrencyOptionsYerParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsYerCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsYerTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsZarCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsZarTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsZarTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsZarTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ZAR.
type ProductDefaultPriceDataCurrencyOptionsZarParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsZarCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsZarTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsZmwCustomUnitAmountParams struct {
	// Pass in `true` to enable `custom_unit_amount`, otherwise omit `custom_unit_amount`.
	Enabled *bool `form:"enabled"`
	// The maximum unit amount the customer can specify for this item.
	Maximum *int64 `form:"maximum"`
	// The minimum unit amount the customer can specify for this item. Must be at least the minimum charge amount.
	Minimum *int64 `form:"minimum"`
	// The starting unit amount which can be updated by the customer.
	Preset *int64 `form:"preset"`
}

// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
type ProductDefaultPriceDataCurrencyOptionsZmwTierParams struct {
	// The flat billing amount for an entire tier, regardless of the number of units in the tier.
	FlatAmount *int64 `form:"flat_amount"`
	// Same as `flat_amount`, but accepts a decimal value representing an integer in the minor units of the currency. Only one of `flat_amount` and `flat_amount_decimal` can be set.
	FlatAmountDecimal *float64 `form:"flat_amount_decimal,high_precision"`
	// The per unit billing amount for each individual unit for which this tier applies.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
	// Specifies the upper bound of this tier. The lower bound of a tier is the upper bound of the previous tier adding one. Use `inf` to define a fallback tier.
	UpTo    *int64 `form:"up_to"`
	UpToInf *bool  `form:"-"` // See custom AppendTo
}

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsZmwTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsZmwTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// The price defined in ZMW.
type ProductDefaultPriceDataCurrencyOptionsZmwParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsZmwCustomUnitAmountParams `form:"custom_unit_amount"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsZmwTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// Prices defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
type ProductDefaultPriceDataCurrencyOptionsParams struct {
	// The price defined in AED.
	Aed *ProductDefaultPriceDataCurrencyOptionsAedParams `form:"aed"`
	// The price defined in AFN.
	Afn *ProductDefaultPriceDataCurrencyOptionsAfnParams `form:"afn"`
	// The price defined in ALL.
	All *ProductDefaultPriceDataCurrencyOptionsAllParams `form:"all"`
	// The price defined in AMD.
	Amd *ProductDefaultPriceDataCurrencyOptionsAmdParams `form:"amd"`
	// The price defined in ANG.
	Ang *ProductDefaultPriceDataCurrencyOptionsAngParams `form:"ang"`
	// The price defined in AOA.
	Aoa *ProductDefaultPriceDataCurrencyOptionsAoaParams `form:"aoa"`
	// The price defined in ARS.
	Ars *ProductDefaultPriceDataCurrencyOptionsArsParams `form:"ars"`
	// The price defined in AUD.
	AUD *ProductDefaultPriceDataCurrencyOptionsAUDParams `form:"aud"`
	// The price defined in AWG.
	Awg *ProductDefaultPriceDataCurrencyOptionsAwgParams `form:"awg"`
	// The price defined in AZN.
	Azn *ProductDefaultPriceDataCurrencyOptionsAznParams `form:"azn"`
	// The price defined in BAM.
	Bam *ProductDefaultPriceDataCurrencyOptionsBamParams `form:"bam"`
	// The price defined in BBD.
	Bbd *ProductDefaultPriceDataCurrencyOptionsBbdParams `form:"bbd"`
	// The price defined in BDT.
	Bdt *ProductDefaultPriceDataCurrencyOptionsBdtParams `form:"bdt"`
	// The price defined in BGN.
	Bgn *ProductDefaultPriceDataCurrencyOptionsBgnParams `form:"bgn"`
	// The price defined in BHD.
	Bhd *ProductDefaultPriceDataCurrencyOptionsBhdParams `form:"bhd"`
	// The price defined in BIF.
	Bif *ProductDefaultPriceDataCurrencyOptionsBifParams `form:"bif"`
	// The price defined in BMD.
	Bmd *ProductDefaultPriceDataCurrencyOptionsBmdParams `form:"bmd"`
	// The price defined in BND.
	Bnd *ProductDefaultPriceDataCurrencyOptionsBndParams `form:"bnd"`
	// The price defined in BOB.
	Bob *ProductDefaultPriceDataCurrencyOptionsBobParams `form:"bob"`
	// The price defined in BRL.
	Brl *ProductDefaultPriceDataCurrencyOptionsBrlParams `form:"brl"`
	// The price defined in BSD.
	Bsd *ProductDefaultPriceDataCurrencyOptionsBsdParams `form:"bsd"`
	// The price defined in BTN.
	Btn *ProductDefaultPriceDataCurrencyOptionsBtnParams `form:"btn"`
	// The price defined in BWP.
	Bwp *ProductDefaultPriceDataCurrencyOptionsBwpParams `form:"bwp"`
	// The price defined in BYN.
	Byn *ProductDefaultPriceDataCurrencyOptionsBynParams `form:"byn"`
	// The price defined in BZD.
	Bzd *ProductDefaultPriceDataCurrencyOptionsBzdParams `form:"bzd"`
	// The price defined in CAD.
	CAD *ProductDefaultPriceDataCurrencyOptionsCADParams `form:"cad"`
	// The price defined in CDF.
	Cdf *ProductDefaultPriceDataCurrencyOptionsCdfParams `form:"cdf"`
	// The price defined in CHF.
	CHF *ProductDefaultPriceDataCurrencyOptionsCHFParams `form:"chf"`
	// The price defined in CLP.
	Clp *ProductDefaultPriceDataCurrencyOptionsClpParams `form:"clp"`
	// The price defined in CNY.
	Cny *ProductDefaultPriceDataCurrencyOptionsCnyParams `form:"cny"`
	// The price defined in COP.
	Cop *ProductDefaultPriceDataCurrencyOptionsCopParams `form:"cop"`
	// The price defined in CRC.
	Crc *ProductDefaultPriceDataCurrencyOptionsCrcParams `form:"crc"`
	// The price defined in CVE.
	Cve *ProductDefaultPriceDataCurrencyOptionsCveParams `form:"cve"`
	// The price defined in CZK.
	CZK *ProductDefaultPriceDataCurrencyOptionsCZKParams `form:"czk"`
	// The price defined in DJF.
	Djf *ProductDefaultPriceDataCurrencyOptionsDjfParams `form:"djf"`
	// The price defined in DKK.
	DKK *ProductDefaultPriceDataCurrencyOptionsDKKParams `form:"dkk"`
	// The price defined in DOP.
	Dop *ProductDefaultPriceDataCurrencyOptionsDopParams `form:"dop"`
	// The price defined in DZD.
	Dzd *ProductDefaultPriceDataCurrencyOptionsDzdParams `form:"dzd"`
	// The price defined in EGP.
	Egp *ProductDefaultPriceDataCurrencyOptionsEgpParams `form:"egp"`
	// The price defined in ETB.
	Etb *ProductDefaultPriceDataCurrencyOptionsEtbParams `form:"etb"`
	// The price defined in EUR.
	EUR *ProductDefaultPriceDataCurrencyOptionsEURParams `form:"eur"`
	// The price defined in FJD.
	Fjd *ProductDefaultPriceDataCurrencyOptionsFjdParams `form:"fjd"`
	// The price defined in FKP.
	Fkp *ProductDefaultPriceDataCurrencyOptionsFkpParams `form:"fkp"`
	// The price defined in GBP.
	GBP *ProductDefaultPriceDataCurrencyOptionsGBPParams `form:"gbp"`
	// The price defined in GEL.
	Gel *ProductDefaultPriceDataCurrencyOptionsGelParams `form:"gel"`
	// The price defined in GHS.
	Ghs *ProductDefaultPriceDataCurrencyOptionsGhsParams `form:"ghs"`
	// The price defined in GIP.
	Gip *ProductDefaultPriceDataCurrencyOptionsGipParams `form:"gip"`
	// The price defined in GMD.
	Gmd *ProductDefaultPriceDataCurrencyOptionsGmdParams `form:"gmd"`
	// The price defined in GNF.
	Gnf *ProductDefaultPriceDataCurrencyOptionsGnfParams `form:"gnf"`
	// The price defined in GTQ.
	Gtq *ProductDefaultPriceDataCurrencyOptionsGtqParams `form:"gtq"`
	// The price defined in GYD.
	Gyd *ProductDefaultPriceDataCurrencyOptionsGydParams `form:"gyd"`
	// The price defined in HKD.
	HKD *ProductDefaultPriceDataCurrencyOptionsHKDParams `form:"hkd"`
	// The price defined in HNL.
	Hnl *ProductDefaultPriceDataCurrencyOptionsHnlParams `form:"hnl"`
	// The price defined in HRK.
	Hrk *ProductDefaultPriceDataCurrencyOptionsHrkParams `form:"hrk"`
	// The price defined in HTG.
	Htg *ProductDefaultPriceDataCurrencyOptionsHtgParams `form:"htg"`
	// The price defined in HUF.
	Huf *ProductDefaultPriceDataCurrencyOptionsHufParams `form:"huf"`
	// The price defined in IDR.
	Idr *ProductDefaultPriceDataCurrencyOptionsIdrParams `form:"idr"`
	// The price defined in ILS.
	Ils *ProductDefaultPriceDataCurrencyOptionsIlsParams `form:"ils"`
	// The price defined in INR.
	Inr *ProductDefaultPriceDataCurrencyOptionsInrParams `form:"inr"`
	// The price defined in ISK.
	Isk *ProductDefaultPriceDataCurrencyOptionsIskParams `form:"isk"`
	// The price defined in JMD.
	Jmd *ProductDefaultPriceDataCurrencyOptionsJmdParams `form:"jmd"`
	// The price defined in JOD.
	Jod *ProductDefaultPriceDataCurrencyOptionsJodParams `form:"jod"`
	// The price defined in JPY.
	Jpy *ProductDefaultPriceDataCurrencyOptionsJpyParams `form:"jpy"`
	// The price defined in KES.
	Kes *ProductDefaultPriceDataCurrencyOptionsKesParams `form:"kes"`
	// The price defined in KGS.
	Kgs *ProductDefaultPriceDataCurrencyOptionsKgsParams `form:"kgs"`
	// The price defined in KHR.
	Khr *ProductDefaultPriceDataCurrencyOptionsKhrParams `form:"khr"`
	// The price defined in KMF.
	Kmf *ProductDefaultPriceDataCurrencyOptionsKmfParams `form:"kmf"`
	// The price defined in KRW.
	Krw *ProductDefaultPriceDataCurrencyOptionsKrwParams `form:"krw"`
	// The price defined in KWD.
	Kwd *ProductDefaultPriceDataCurrencyOptionsKwdParams `form:"kwd"`
	// The price defined in KYD.
	Kyd *ProductDefaultPriceDataCurrencyOptionsKydParams `form:"kyd"`
	// The price defined in KZT.
	Kzt *ProductDefaultPriceDataCurrencyOptionsKztParams `form:"kzt"`
	// The price defined in LAK.
	Lak *ProductDefaultPriceDataCurrencyOptionsLakParams `form:"lak"`
	// The price defined in LBP.
	Lbp *ProductDefaultPriceDataCurrencyOptionsLbpParams `form:"lbp"`
	// The price defined in LKR.
	Lkr *ProductDefaultPriceDataCurrencyOptionsLkrParams `form:"lkr"`
	// The price defined in LRD.
	Lrd *ProductDefaultPriceDataCurrencyOptionsLrdParams `form:"lrd"`
	// The price defined in LSL.
	Lsl *ProductDefaultPriceDataCurrencyOptionsLslParams `form:"lsl"`
	// The price defined in MAD.
	Mad *ProductDefaultPriceDataCurrencyOptionsMadParams `form:"mad"`
	// The price defined in MDL.
	Mdl *ProductDefaultPriceDataCurrencyOptionsMdlParams `form:"mdl"`
	// The price defined in MGA.
	Mga *ProductDefaultPriceDataCurrencyOptionsMgaParams `form:"mga"`
	// The price defined in MKD.
	Mkd *ProductDefaultPriceDataCurrencyOptionsMkdParams `form:"mkd"`
	// The price defined in MMK.
	Mmk *ProductDefaultPriceDataCurrencyOptionsMmkParams `form:"mmk"`
	// The price defined in MNT.
	Mnt *ProductDefaultPriceDataCurrencyOptionsMntParams `form:"mnt"`
	// The price defined in MOP.
	Mop *ProductDefaultPriceDataCurrencyOptionsMopParams `form:"mop"`
	// The price defined in MRO.
	Mro *ProductDefaultPriceDataCurrencyOptionsMroParams `form:"mro"`
	// The price defined in MUR.
	Mur *ProductDefaultPriceDataCurrencyOptionsMurParams `form:"mur"`
	// The price defined in MVR.
	Mvr *ProductDefaultPriceDataCurrencyOptionsMvrParams `form:"mvr"`
	// The price defined in MWK.
	Mwk *ProductDefaultPriceDataCurrencyOptionsMwkParams `form:"mwk"`
	// The price defined in MXN.
	Mxn *ProductDefaultPriceDataCurrencyOptionsMxnParams `form:"mxn"`
	// The price defined in MYR.
	MYR *ProductDefaultPriceDataCurrencyOptionsMYRParams `form:"myr"`
	// The price defined in MZN.
	Mzn *ProductDefaultPriceDataCurrencyOptionsMznParams `form:"mzn"`
	// The price defined in NAD.
	Nad *ProductDefaultPriceDataCurrencyOptionsNadParams `form:"nad"`
	// The price defined in NGN.
	Ngn *ProductDefaultPriceDataCurrencyOptionsNgnParams `form:"ngn"`
	// The price defined in NIO.
	Nio *ProductDefaultPriceDataCurrencyOptionsNioParams `form:"nio"`
	// The price defined in NOK.
	NOK *ProductDefaultPriceDataCurrencyOptionsNOKParams `form:"nok"`
	// The price defined in NPR.
	Npr *ProductDefaultPriceDataCurrencyOptionsNprParams `form:"npr"`
	// The price defined in NZD.
	NZD *ProductDefaultPriceDataCurrencyOptionsNZDParams `form:"nzd"`
	// The price defined in OMR.
	Omr *ProductDefaultPriceDataCurrencyOptionsOmrParams `form:"omr"`
	// The price defined in PAB.
	Pab *ProductDefaultPriceDataCurrencyOptionsPabParams `form:"pab"`
	// The price defined in PEN.
	Pen *ProductDefaultPriceDataCurrencyOptionsPenParams `form:"pen"`
	// The price defined in PGK.
	Pgk *ProductDefaultPriceDataCurrencyOptionsPgkParams `form:"pgk"`
	// The price defined in PHP.
	Php *ProductDefaultPriceDataCurrencyOptionsPhpParams `form:"php"`
	// The price defined in PKR.
	Pkr *ProductDefaultPriceDataCurrencyOptionsPkrParams `form:"pkr"`
	// The price defined in PLN.
	Pln *ProductDefaultPriceDataCurrencyOptionsPlnParams `form:"pln"`
	// The price defined in PYG.
	Pyg *ProductDefaultPriceDataCurrencyOptionsPygParams `form:"pyg"`
	// The price defined in QAR.
	Qar *ProductDefaultPriceDataCurrencyOptionsQarParams `form:"qar"`
	// The price defined in RON.
	Ron *ProductDefaultPriceDataCurrencyOptionsRonParams `form:"ron"`
	// The price defined in RSD.
	Rsd *ProductDefaultPriceDataCurrencyOptionsRsdParams `form:"rsd"`
	// The price defined in RUB.
	Rub *ProductDefaultPriceDataCurrencyOptionsRubParams `form:"rub"`
	// The price defined in RWF.
	Rwf *ProductDefaultPriceDataCurrencyOptionsRwfParams `form:"rwf"`
	// The price defined in SAR.
	Sar *ProductDefaultPriceDataCurrencyOptionsSarParams `form:"sar"`
	// The price defined in SBD.
	Sbd *ProductDefaultPriceDataCurrencyOptionsSbdParams `form:"sbd"`
	// The price defined in SCR.
	Scr *ProductDefaultPriceDataCurrencyOptionsScrParams `form:"scr"`
	// The price defined in SEK.
	SEK *ProductDefaultPriceDataCurrencyOptionsSEKParams `form:"sek"`
	// The price defined in SGD.
	SGD *ProductDefaultPriceDataCurrencyOptionsSGDParams `form:"sgd"`
	// The price defined in SHP.
	Shp *ProductDefaultPriceDataCurrencyOptionsShpParams `form:"shp"`
	// The price defined in SLL.
	Sll *ProductDefaultPriceDataCurrencyOptionsSllParams `form:"sll"`
	// The price defined in SOS.
	Sos *ProductDefaultPriceDataCurrencyOptionsSosParams `form:"sos"`
	// The price defined in SRD.
	Srd *ProductDefaultPriceDataCurrencyOptionsSrdParams `form:"srd"`
	// The price defined in STD.
	Std *ProductDefaultPriceDataCurrencyOptionsStdParams `form:"std"`
	// The price defined in SZL.
	Szl *ProductDefaultPriceDataCurrencyOptionsSzlParams `form:"szl"`
	// The price defined in THB.
	Thb *ProductDefaultPriceDataCurrencyOptionsThbParams `form:"thb"`
	// The price defined in TJS.
	Tjs *ProductDefaultPriceDataCurrencyOptionsTjsParams `form:"tjs"`
	// The price defined in TND.
	Tnd *ProductDefaultPriceDataCurrencyOptionsTndParams `form:"tnd"`
	// The price defined in TOP.
	Top *ProductDefaultPriceDataCurrencyOptionsTopParams `form:"top"`
	// The price defined in TRY.
	Try *ProductDefaultPriceDataCurrencyOptionsTryParams `form:"try"`
	// The price defined in TTD.
	Ttd *ProductDefaultPriceDataCurrencyOptionsTtdParams `form:"ttd"`
	// The price defined in TWD.
	Twd *ProductDefaultPriceDataCurrencyOptionsTwdParams `form:"twd"`
	// The price defined in TZS.
	Tzs *ProductDefaultPriceDataCurrencyOptionsTzsParams `form:"tzs"`
	// The price defined in UAH.
	Uah *ProductDefaultPriceDataCurrencyOptionsUahParams `form:"uah"`
	// The price defined in UGX.
	Ugx *ProductDefaultPriceDataCurrencyOptionsUgxParams `form:"ugx"`
	// The price defined in USD.
	USD *ProductDefaultPriceDataCurrencyOptionsUSDParams `form:"usd"`
	// The price defined in USDC.
	Usdc *ProductDefaultPriceDataCurrencyOptionsUsdcParams `form:"usdc"`
	// The price defined in UYU.
	Uyu *ProductDefaultPriceDataCurrencyOptionsUyuParams `form:"uyu"`
	// The price defined in UZS.
	Uzs *ProductDefaultPriceDataCurrencyOptionsUzsParams `form:"uzs"`
	// The price defined in VND.
	Vnd *ProductDefaultPriceDataCurrencyOptionsVndParams `form:"vnd"`
	// The price defined in VUV.
	Vuv *ProductDefaultPriceDataCurrencyOptionsVuvParams `form:"vuv"`
	// The price defined in WST.
	Wst *ProductDefaultPriceDataCurrencyOptionsWstParams `form:"wst"`
	// The price defined in XAF.
	Xaf *ProductDefaultPriceDataCurrencyOptionsXafParams `form:"xaf"`
	// The price defined in XCD.
	Xcd *ProductDefaultPriceDataCurrencyOptionsXcdParams `form:"xcd"`
	// The price defined in XOF.
	Xof *ProductDefaultPriceDataCurrencyOptionsXofParams `form:"xof"`
	// The price defined in XPF.
	Xpf *ProductDefaultPriceDataCurrencyOptionsXpfParams `form:"xpf"`
	// The price defined in YER.
	Yer *ProductDefaultPriceDataCurrencyOptionsYerParams `form:"yer"`
	// The price defined in ZAR.
	Zar *ProductDefaultPriceDataCurrencyOptionsZarParams `form:"zar"`
	// The price defined in ZMW.
	Zmw *ProductDefaultPriceDataCurrencyOptionsZmwParams `form:"zmw"`
}

// The recurring components of a price such as `interval` and `interval_count`.
type ProductDefaultPriceDataRecurringParams struct {
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks).
	IntervalCount *int64 `form:"interval_count"`
}

// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object. This Price will be set as the default price for this product.
type ProductDefaultPriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Prices defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions *ProductDefaultPriceDataCurrencyOptionsParams `form:"currency_options"`
	// The recurring components of a price such as `interval` and `interval_count`.
	Recurring *ProductDefaultPriceDataRecurringParams `form:"recurring"`
	// Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge. One of `unit_amount` or `unit_amount_decimal` is required.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// The dimensions of this product for shipping purposes.
type PackageDimensionsParams struct {
	// Height, in inches. Maximum precision is 2 decimal places.
	Height *float64 `form:"height"`
	// Length, in inches. Maximum precision is 2 decimal places.
	Length *float64 `form:"length"`
	// Weight, in ounces. Maximum precision is 2 decimal places.
	Weight *float64 `form:"weight"`
	// Width, in inches. Maximum precision is 2 decimal places.
	Width *float64 `form:"width"`
}

// Creates a new product object.
type ProductParams struct {
	Params `form:"*"`
	// Whether the product is available for purchase.
	Active *bool `form:"active"`
	// A list of up to 5 alphanumeric attributes that each SKU can provide values for (e.g., `["color", "size"]`). If a value for `attributes` is specified, the list specified will replace the existing attributes list on this product. Any attributes not present after the update will be deleted from the SKUs for this product.
	Attributes []*string `form:"attributes"`
	// A short one-line description of the product, meant to be displayable to the customer. May only be set if `type=good`.
	Caption *string `form:"caption"`
	// An array of Connect application names or identifiers that should not be able to order the SKUs for this product. May only be set if `type=good`.
	DeactivateOn []*string `form:"deactivate_on"`
	// The ID of the [Price](https://stripe.com/docs/api/prices) object that is the default price for this product.
	DefaultPrice *string `form:"default_price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object. This Price will be set as the default price for this product.
	DefaultPriceData *ProductDefaultPriceDataParams `form:"default_price_data"`
	// The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description *string `form:"description"`
	// An identifier will be randomly generated by Stripe. You can optionally override this ID, but the ID must be unique across all products in your Stripe account.
	ID *string `form:"id"`
	// A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
	Images []*string `form:"images"`
	// The product's name, meant to be displayable to the customer.
	Name *string `form:"name"`
	// The dimensions of this product for shipping purposes.
	PackageDimensions *PackageDimensionsParams `form:"package_dimensions"`
	// Whether this product is shipped (i.e., physical goods).
	Shippable *bool `form:"shippable"`
	// An arbitrary string to be displayed on your customer's credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all.
	//
	// This may be up to 22 characters. The statement description may not include `<`, `>`, `\`, `"`, `'` characters, and will appear on your customer's statement in capital letters. Non-ASCII characters are automatically stripped.
	//  It must contain at least one letter. May only be set if `type=service`.
	StatementDescriptor *string `form:"statement_descriptor"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID.
	TaxCode *string `form:"tax_code"`
	// The type of the product. Defaults to `service` if not explicitly specified, enabling use of this product with Subscriptions and Plans. Set this parameter to `good` to use this product with Orders and SKUs. On API versions before `2018-02-05`, this field defaults to `good` for compatibility reasons.
	Type *string `form:"type"`
	// A label that represents units of this product in Stripe and on customers' receipts and invoices. When set, this will be included in associated invoice line item descriptions. May only be set if `type=service`.
	UnitLabel *string `form:"unit_label"`
	// A URL of a publicly-accessible webpage for this product.
	URL *string `form:"url"`
}

// Returns a list of your products. The products are returned sorted by creation date, with the most recently created products appearing first.
type ProductListParams struct {
	ListParams `form:"*"`
	// Only return products that are active or inactive (e.g., pass `false` to list all inactive products).
	Active *bool `form:"active"`
	// Only return products that were created during the given date interval.
	Created *int64 `form:"created"`
	// Only return products that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return products with the given IDs.
	IDs []*string `form:"ids"`
	// Only return products that can be shipped (i.e., physical, not digital products).
	Shippable *bool `form:"shippable"`
	// Only return products of this type.
	Type *string `form:"type"`
	// Only return products with the given url.
	URL *string `form:"url"`
}

// The dimensions of this product for shipping purposes.
type PackageDimensions struct {
	// Height, in inches.
	Height float64 `json:"height"`
	// Length, in inches.
	Length float64 `json:"length"`
	// Weight, in ounces.
	Weight float64 `json:"weight"`
	// Width, in inches.
	Width float64 `json:"width"`
}

// Products describe the specific goods or services you offer to your customers.
// For example, you might offer a Standard and Premium version of your goods or service; each version would be a separate Product.
// They can be used in conjunction with [Prices](https://stripe.com/docs/api#prices) to configure pricing in Payment Links, Checkout, and Subscriptions.
//
// Related guides: [Set up a subscription](https://stripe.com/docs/billing/subscriptions/set-up-subscription),
// [share a Payment Link](https://stripe.com/docs/payments/payment-links/overview),
// [accept payments with Checkout](https://stripe.com/docs/payments/accept-a-payment#create-product-prices-upfront),
// and more about [Products and Prices](https://stripe.com/docs/products-prices/overview)
type Product struct {
	APIResource
	// Whether the product is currently available for purchase.
	Active bool `json:"active"`
	// A list of up to 5 attributes that each SKU can provide values for (e.g., `["color", "size"]`).
	Attributes []string `json:"attributes"`
	// A short one-line description of the product, meant to be displayable to the customer. Only applicable to products of `type=good`.
	Caption string `json:"caption"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// An array of connect application identifiers that cannot purchase this product. Only applicable to products of `type=good`.
	DeactivateOn []string `json:"deactivate_on"`
	// The ID of the [Price](https://stripe.com/docs/api/prices) object that is the default price for this product.
	DefaultPrice *Price `json:"default_price"`
	Deleted      bool   `json:"deleted"`
	// The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description string `json:"description"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
	Images []string `json:"images"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The product's name, meant to be displayable to the customer.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The dimensions of this product for shipping purposes.
	PackageDimensions *PackageDimensions `json:"package_dimensions"`
	// Whether this product is shipped (i.e., physical goods).
	Shippable bool `json:"shippable"`
	// Extra information about a product which will appear on your customer's credit card statement. In the case that multiple products are billed at once, the first statement descriptor will be used.
	StatementDescriptor string `json:"statement_descriptor"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID.
	TaxCode *TaxCode `json:"tax_code"`
	// The type of the product. The product is either of type `good`, which is eligible for use with Orders and SKUs, or `service`, which is eligible for use with Subscriptions and Plans.
	Type ProductType `json:"type"`
	// A label that represents units of this product in Stripe and on customers' receipts and invoices. When set, this will be included in associated invoice line item descriptions.
	UnitLabel string `json:"unit_label"`
	// Time at which the object was last updated. Measured in seconds since the Unix epoch.
	Updated int64 `json:"updated"`
	// A URL of a publicly-accessible webpage for this product.
	URL string `json:"url"`
}

// ProductList is a list of Products as retrieved from a list endpoint.
type ProductList struct {
	APIResource
	ListMeta
	Data []*Product `json:"data"`
}

// ProductSearchResult is a list of Product search results as retrieved from a search endpoint.
type ProductSearchResult struct {
	APIResource
	SearchMeta
	Data []*Product `json:"data"`
}

// UnmarshalJSON handles deserialization of a Product.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *Product) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type product Product
	var v product
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = Product(v)
	return nil
}
