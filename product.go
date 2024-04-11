//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v76/form"
)

// The specific type of gift_card provisioning, only `fixed_amount` currently supported.
type ProductProvisioningGiftCardType string

// List of values that ProductProvisioningGiftCardType can take
const (
	ProductProvisioningGiftCardTypeFixedAmount ProductProvisioningGiftCardType = "fixed_amount"
)

// The type of provisioning, only `gift_card` currently supported.
type ProductProvisioningType string

// List of values that ProductProvisioningType can take
const (
	ProductProvisioningTypeGiftCard ProductProvisioningType = "gift_card"
)

// The type of the product. The product is either of type `good`, which is eligible for use with Orders and SKUs, or `service`, which is eligible for use with Subscriptions and Plans.
type ProductType string

// List of values that ProductType can take
const (
	ProductTypeGood    ProductType = "good"
	ProductTypeService ProductType = "service"
)

// Delete a product. Deleting a product is only possible if it has no prices associated with it. Additionally, deleting a product with type=good is only possible if it has no SKUs associated with it.
type ProductParams struct {
	Params `form:"*"`
	// Whether the product is currently available for purchase. Defaults to `true`.
	Active *bool `form:"active"`
	// The ID of the [Price](https://stripe.com/docs/api/prices) object that is the default price for this product.
	DefaultPrice *string `form:"default_price"`
	// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object. This Price will be set as the default price for this product.
	DefaultPriceData *ProductDefaultPriceDataParams `form:"default_price_data"`
	// The product's description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.
	Description *string `form:"description"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// An identifier will be randomly generated by Stripe. You can optionally override this ID, but the ID must be unique across all products in your Stripe account.
	ID *string `form:"id"`
	// A list of up to 8 URLs of images for this product, meant to be displayable to the customer.
	Images []*string `form:"images"`
	// A list of up to 15 marketing features for this product. These are displayed in [pricing tables](https://stripe.com/docs/payments/checkout/pricing-table).
	MarketingFeatures []*ProductMarketingFeatureParams `form:"marketing_features"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The product's name, meant to be displayable to the customer.
	Name *string `form:"name"`
	// The dimensions of this product for shipping purposes.
	PackageDimensions *ProductPackageDimensionsParams `form:"package_dimensions"`
	// Provisioning configuration for this product.
	Provisioning *ProductProvisioningParams `form:"provisioning"`
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
	// A label that represents units of this product. When set, this will be included in customers' receipts, invoices, Checkout, and the customer portal. May only be set if `type=service`.
	UnitLabel *string `form:"unit_label"`
	// A URL of a publicly-accessible webpage for this product.
	URL *string `form:"url"`
}

// AddExpand appends a new field to expand.
func (p *ProductParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *ProductParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// A list of up to 15 marketing features for this product. These are displayed in [pricing tables](https://stripe.com/docs/payments/checkout/pricing-table).
type ProductMarketingFeatureParams struct {
	// The marketing feature name. Up to 80 characters long.
	Name *string `form:"name"`
}

// The dimensions of this product for shipping purposes.
type ProductPackageDimensionsParams struct {
	// Height, in inches. Maximum precision is 2 decimal places.
	Height *float64 `form:"height"`
	// Length, in inches. Maximum precision is 2 decimal places.
	Length *float64 `form:"length"`
	// Weight, in ounces. Maximum precision is 2 decimal places.
	Weight *float64 `form:"weight"`
	// Width, in inches. Maximum precision is 2 decimal places.
	Width *float64 `form:"width"`
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
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Only return products with the given IDs. Cannot be used with [starting_after](https://stripe.com/docs/api#list_products-starting_after) or [ending_before](https://stripe.com/docs/api#list_products-ending_before).
	IDs []*string `form:"ids"`
	// Only return products that can be shipped (i.e., physical, not digital products).
	Shippable *bool `form:"shippable"`
	// Only return products of this type.
	Type *string `form:"type"`
	// Only return products with the given url.
	URL *string `form:"url"`
}

// AddExpand appends a new field to expand.
func (p *ProductListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
type ProductDefaultPriceDataCurrencyOptionsCustomUnitAmountParams struct {
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
type ProductDefaultPriceDataCurrencyOptionsTierParams struct {
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

// AppendTo implements custom encoding logic for ProductDefaultPriceDataCurrencyOptionsTierParams.
func (p *ProductDefaultPriceDataCurrencyOptionsTierParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(p.UpToInf) {
		body.Add(form.FormatKey(append(keyParts, "up_to")), "inf")
	}
}

// Prices defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
type ProductDefaultPriceDataCurrencyOptionsParams struct {
	// When set, provides configuration for the amount to be adjusted by the customer during Checkout Sessions and Payment Links.
	CustomUnitAmount *ProductDefaultPriceDataCurrencyOptionsCustomUnitAmountParams `form:"custom_unit_amount"`
	// Only required if a [default tax behavior](https://stripe.com/docs/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// Each element represents a pricing tier. This parameter requires `billing_scheme` to be set to `tiered`. See also the documentation for `billing_scheme`.
	Tiers []*ProductDefaultPriceDataCurrencyOptionsTierParams `form:"tiers"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}

// The recurring components of a price such as `interval` and `interval_count`.
type ProductDefaultPriceDataRecurringParams struct {
	// Specifies billing frequency. Either `day`, `week`, `month` or `year`.
	Interval *string `form:"interval"`
	// The number of intervals between subscription billings. For example, `interval=month` and `interval_count=3` bills every 3 months. Maximum of three years interval allowed (3 years, 36 months, or 156 weeks).
	IntervalCount *int64 `form:"interval_count"`
}

// Data used to generate a new [Price](https://stripe.com/docs/api/prices) object. This Price will be set as the default price for this product.
type ProductDefaultPriceDataParams struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// Prices defined in each available currency option. Each key must be a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html) and a [supported currency](https://stripe.com/docs/currencies).
	CurrencyOptions map[string]*ProductDefaultPriceDataCurrencyOptionsParams `form:"currency_options"`
	// The recurring components of a price such as `interval` and `interval_count`.
	Recurring *ProductDefaultPriceDataRecurringParams `form:"recurring"`
	// Only required if a [default tax behavior](https://stripe.com/docs/tax/products-prices-tax-categories-tax-behavior#setting-a-default-tax-behavior-(recommended)) was not provided in the Stripe Tax settings. Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of `inclusive`, `exclusive`, or `unspecified`. Once specified as either `inclusive` or `exclusive`, it cannot be changed.
	TaxBehavior *string `form:"tax_behavior"`
	// A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge. One of `unit_amount` or `unit_amount_decimal` is required.
	UnitAmount *int64 `form:"unit_amount"`
	// Same as `unit_amount`, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of `unit_amount` and `unit_amount_decimal` can be set.
	UnitAmountDecimal *float64 `form:"unit_amount_decimal,high_precision"`
}
type ProductProvisioningGiftCardFixedAmountParams struct {
	// The initial amount with which the provisioned gift card will be created.
	Amount   *int64  `form:"amount"`
	Currency *string `form:"currency"`
}
type ProductProvisioningGiftCardParams struct {
	FixedAmount *ProductProvisioningGiftCardFixedAmountParams `form:"fixed_amount"`
	// The specific type of gift_card provisioning, only `fixed_amount` currently supported.
	Type *string `form:"type"`
}

// Provisioning configuration for this product.
type ProductProvisioningParams struct {
	GiftCard *ProductProvisioningGiftCardParams `form:"gift_card"`
	// The type of provisioning, only `gift_card` currently supported.
	Type *string `form:"type"`
}

// Search for products you've previously created using Stripe's [Search Query Language](https://stripe.com/docs/search#search-query-language).
// Don't use search in read-after-write flows where strict consistency is necessary. Under normal operating
// conditions, data is searchable in less than a minute. Occasionally, propagation of new or updated data can be up
// to an hour behind during outages. Search functionality is not available to merchants in India.
type ProductSearchParams struct {
	SearchParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// A cursor for pagination across multiple pages of results. Don't include this parameter on the first call. Use the next_page value returned in a previous response to request subsequent results.
	Page *string `form:"page"`
}

// AddExpand appends a new field to expand.
func (p *ProductSearchParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A list of up to 15 marketing features for this product. These are displayed in [pricing tables](https://stripe.com/docs/payments/checkout/pricing-table).
type ProductMarketingFeature struct {
	// The marketing feature name. Up to 80 characters long.
	Name string `json:"name"`
}

// The dimensions of this product for shipping purposes.
type ProductPackageDimensions struct {
	// Height, in inches.
	Height float64 `json:"height"`
	// Length, in inches.
	Length float64 `json:"length"`
	// Weight, in ounces.
	Weight float64 `json:"weight"`
	// Width, in inches.
	Width float64 `json:"width"`
}
type ProductProvisioningGiftCardFixedAmount struct {
	// The initial amount with which the provisioned gift card will be created.
	Amount int64 `json:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
}
type ProductProvisioningGiftCard struct {
	FixedAmount *ProductProvisioningGiftCardFixedAmount `json:"fixed_amount"`
	// The specific type of gift_card provisioning, only `fixed_amount` currently supported.
	Type ProductProvisioningGiftCardType `json:"type"`
}

// Provisioning configuration for this product.
type ProductProvisioning struct {
	GiftCard *ProductProvisioningGiftCard `json:"gift_card"`
	// The type of provisioning, only `gift_card` currently supported.
	Type ProductProvisioningType `json:"type"`
}

// Products describe the specific goods or services you offer to your customers.
// For example, you might offer a Standard and Premium version of your goods or service; each version would be a separate Product.
// They can be used in conjunction with [Prices](https://stripe.com/docs/api#prices) to configure pricing in Payment Links, Checkout, and Subscriptions.
//
// Related guides: [Set up a subscription](https://stripe.com/docs/billing/subscriptions/set-up-subscription),
// [share a Payment Link](https://stripe.com/docs/payment-links),
// [accept payments with Checkout](https://stripe.com/docs/payments/accept-a-payment#create-product-prices-upfront),
// and more about [Products and Prices](https://stripe.com/docs/products-prices/overview)
type Product struct {
	APIResource
	// Whether the product is currently available for purchase.
	Active bool `json:"active"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
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
	// A list of up to 15 marketing features for this product. These are displayed in [pricing tables](https://stripe.com/docs/payments/checkout/pricing-table).
	MarketingFeatures []*ProductMarketingFeature `json:"marketing_features"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The product's name, meant to be displayable to the customer.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The dimensions of this product for shipping purposes.
	PackageDimensions *ProductPackageDimensions `json:"package_dimensions"`
	// Provisioning configuration for this product.
	Provisioning *ProductProvisioning `json:"provisioning"`
	// Whether this product is shipped (i.e., physical goods).
	Shippable bool `json:"shippable"`
	// Extra information about a product which will appear on your customer's credit card statement. In the case that multiple products are billed at once, the first statement descriptor will be used.
	StatementDescriptor string `json:"statement_descriptor"`
	// A [tax code](https://stripe.com/docs/tax/tax-categories) ID.
	TaxCode *TaxCode `json:"tax_code"`
	// The type of the product. The product is either of type `good`, which is eligible for use with Orders and SKUs, or `service`, which is eligible for use with Subscriptions and Plans.
	Type ProductType `json:"type"`
	// A label that represents units of this product. When set, this will be included in customers' receipts, invoices, Checkout, and the customer portal.
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
