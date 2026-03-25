//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The type of this amount. We currently only support `monetary` billing credits.
type BillingCreditGrantAmountType string

// List of values that BillingCreditGrantAmountType can take
const (
	BillingCreditGrantAmountTypeCustomPricingUnit BillingCreditGrantAmountType = "custom_pricing_unit"
	BillingCreditGrantAmountTypeMonetary          BillingCreditGrantAmountType = "monetary"
)

// The price type that credit grants can apply to. We currently only support the `metered` price type. This refers to prices that have a [Billing Meter](https://docs.stripe.com/api/billing/meter) attached to them. Cannot be used in combination with `prices`.
type BillingCreditGrantApplicabilityConfigScopePriceType string

// List of values that BillingCreditGrantApplicabilityConfigScopePriceType can take
const (
	BillingCreditGrantApplicabilityConfigScopePriceTypeMetered BillingCreditGrantApplicabilityConfigScopePriceType = "metered"
)

// The category of this credit grant. This is for tracking purposes and isn't displayed to the customer.
type BillingCreditGrantCategory string

// List of values that BillingCreditGrantCategory can take
const (
	BillingCreditGrantCategoryPaid        BillingCreditGrantCategory = "paid"
	BillingCreditGrantCategoryPromotional BillingCreditGrantCategory = "promotional"
)

// Retrieve a list of credit grants.
type BillingCreditGrantListParams struct {
	ListParams `form:"*"`
	// Only return credit grants for this customer.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Only return credit grants for this account representing the customer.
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *BillingCreditGrantListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The custom pricing unit amount.
type BillingCreditGrantAmountCustomPricingUnitParams struct {
	// The ID of the custom pricing unit.
	ID *string `form:"id" json:"id"`
	// A positive integer representing the amount of the credit grant.
	Value *float64 `form:"value,high_precision" json:"value,string"`
}

// The monetary amount.
type BillingCreditGrantAmountMonetaryParams struct {
	// Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) of the `value` parameter.
	Currency *string `form:"currency" json:"currency"`
	// A positive integer representing the amount of the credit grant.
	Value *int64 `form:"value" json:"value"`
}

// Amount of this credit grant.
type BillingCreditGrantAmountParams struct {
	// The custom pricing unit amount.
	CustomPricingUnit *BillingCreditGrantAmountCustomPricingUnitParams `form:"custom_pricing_unit" json:"custom_pricing_unit,omitempty"`
	// The monetary amount.
	Monetary *BillingCreditGrantAmountMonetaryParams `form:"monetary" json:"monetary,omitempty"`
	// The type of this amount. We currently only support `monetary` billing credits.
	Type *string `form:"type" json:"type"`
}

// A list of billable items that the credit grant can apply to. We currently only support metered billable items. Cannot be used in combination with `price_type` or `prices`.
type BillingCreditGrantApplicabilityConfigScopeBillableItemParams struct {
	// The billable item ID this credit grant should apply to.
	ID *string `form:"id" json:"id"`
}

// A list of prices that the credit grant can apply to. We currently only support the `metered` prices. Cannot be used in combination with `price_type`.
type BillingCreditGrantApplicabilityConfigScopePriceParams struct {
	// The price ID this credit grant should apply to.
	ID *string `form:"id" json:"id"`
}

// Specify the scope of this applicability config.
type BillingCreditGrantApplicabilityConfigScopeParams struct {
	// A list of billable items that the credit grant can apply to. We currently only support metered billable items. Cannot be used in combination with `price_type` or `prices`.
	BillableItems []*BillingCreditGrantApplicabilityConfigScopeBillableItemParams `form:"billable_items" json:"billable_items,omitempty"`
	// A list of prices that the credit grant can apply to. We currently only support the `metered` prices. Cannot be used in combination with `price_type`.
	Prices []*BillingCreditGrantApplicabilityConfigScopePriceParams `form:"prices" json:"prices,omitempty"`
	// The price type that credit grants can apply to. We currently only support the `metered` price type. Cannot be used in combination with `prices`.
	PriceType *string `form:"price_type" json:"price_type,omitempty"`
}

// Configuration specifying what this credit grant applies to. We currently only support `metered` prices that have a [Billing Meter](https://docs.stripe.com/api/billing/meter) attached to them.
type BillingCreditGrantApplicabilityConfigParams struct {
	// Specify the scope of this applicability config.
	Scope *BillingCreditGrantApplicabilityConfigScopeParams `form:"scope" json:"scope"`
}

// Creates a credit grant.
type BillingCreditGrantParams struct {
	Params `form:"*"`
	// Amount of this credit grant.
	Amount *BillingCreditGrantAmountParams `form:"amount" json:"amount,omitempty"`
	// Configuration specifying what this credit grant applies to. We currently only support `metered` prices that have a [Billing Meter](https://docs.stripe.com/api/billing/meter) attached to them.
	ApplicabilityConfig *BillingCreditGrantApplicabilityConfigParams `form:"applicability_config" json:"applicability_config,omitempty"`
	// The category of this credit grant. It defaults to `paid` if not specified.
	Category *string `form:"category" json:"category,omitempty"`
	// ID of the customer receiving the billing credits.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// ID of the account representing the customer receiving the billing credits.
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// The time when the billing credits become effective-when they're eligible for use. It defaults to the current timestamp if not specified.
	EffectiveAt *int64 `form:"effective_at" json:"effective_at,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The time when the billing credits created by this credit grant expire. If set to empty, the billing credits never expire.
	ExpiresAt *int64 `form:"expires_at" json:"expires_at,omitempty"`
	// Set of key-value pairs that you can attach to an object. You can use this to store additional information about the object (for example, cost basis) in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// A descriptive name shown in the Dashboard.
	Name *string `form:"name" json:"name,omitempty"`
	// The desired priority for applying this credit grant. If not specified, it will be set to the default value of 50. The highest priority is 0 and the lowest is 100.
	Priority    *int64                               `form:"priority" json:"priority,omitempty"`
	UnsetFields []BillingCreditGrantParamsUnsetField `form:"-" json:"-"`
}

// BillingCreditGrantParamsUnsetField is the list of fields that can be cleared/unset on BillingCreditGrantParams.
type BillingCreditGrantParamsUnsetField string

const (
	BillingCreditGrantParamsUnsetFieldExpiresAt BillingCreditGrantParamsUnsetField = "expires_at"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *BillingCreditGrantParams) AddUnsetField(field BillingCreditGrantParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *BillingCreditGrantParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *BillingCreditGrantParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Expires a credit grant.
type BillingCreditGrantExpireParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *BillingCreditGrantExpireParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Voids a credit grant.
type BillingCreditGrantVoidGrantParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *BillingCreditGrantVoidGrantParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The custom pricing unit amount.
type BillingCreditGrantCreateAmountCustomPricingUnitParams struct {
	// The ID of the custom pricing unit.
	ID *string `form:"id" json:"id"`
	// A positive integer representing the amount of the credit grant.
	Value *float64 `form:"value,high_precision" json:"value,string"`
}

// The monetary amount.
type BillingCreditGrantCreateAmountMonetaryParams struct {
	// Three-letter [ISO code for the currency](https://stripe.com/docs/currencies) of the `value` parameter.
	Currency *string `form:"currency" json:"currency"`
	// A positive integer representing the amount of the credit grant.
	Value *int64 `form:"value" json:"value"`
}

// Amount of this credit grant.
type BillingCreditGrantCreateAmountParams struct {
	// The custom pricing unit amount.
	CustomPricingUnit *BillingCreditGrantCreateAmountCustomPricingUnitParams `form:"custom_pricing_unit" json:"custom_pricing_unit,omitempty"`
	// The monetary amount.
	Monetary *BillingCreditGrantCreateAmountMonetaryParams `form:"monetary" json:"monetary,omitempty"`
	// The type of this amount. We currently only support `monetary` billing credits.
	Type *string `form:"type" json:"type"`
}

// A list of billable items that the credit grant can apply to. We currently only support metered billable items. Cannot be used in combination with `price_type` or `prices`.
type BillingCreditGrantCreateApplicabilityConfigScopeBillableItemParams struct {
	// The billable item ID this credit grant should apply to.
	ID *string `form:"id" json:"id"`
}

// A list of prices that the credit grant can apply to. We currently only support the `metered` prices. Cannot be used in combination with `price_type`.
type BillingCreditGrantCreateApplicabilityConfigScopePriceParams struct {
	// The price ID this credit grant should apply to.
	ID *string `form:"id" json:"id"`
}

// Specify the scope of this applicability config.
type BillingCreditGrantCreateApplicabilityConfigScopeParams struct {
	// A list of billable items that the credit grant can apply to. We currently only support metered billable items. Cannot be used in combination with `price_type` or `prices`.
	BillableItems []*BillingCreditGrantCreateApplicabilityConfigScopeBillableItemParams `form:"billable_items" json:"billable_items,omitempty"`
	// A list of prices that the credit grant can apply to. We currently only support the `metered` prices. Cannot be used in combination with `price_type`.
	Prices []*BillingCreditGrantCreateApplicabilityConfigScopePriceParams `form:"prices" json:"prices,omitempty"`
	// The price type that credit grants can apply to. We currently only support the `metered` price type. Cannot be used in combination with `prices`.
	PriceType *string `form:"price_type" json:"price_type,omitempty"`
}

// Configuration specifying what this credit grant applies to. We currently only support `metered` prices that have a [Billing Meter](https://docs.stripe.com/api/billing/meter) attached to them.
type BillingCreditGrantCreateApplicabilityConfigParams struct {
	// Specify the scope of this applicability config.
	Scope *BillingCreditGrantCreateApplicabilityConfigScopeParams `form:"scope" json:"scope"`
}

// Creates a credit grant.
type BillingCreditGrantCreateParams struct {
	Params `form:"*"`
	// Amount of this credit grant.
	Amount *BillingCreditGrantCreateAmountParams `form:"amount" json:"amount"`
	// Configuration specifying what this credit grant applies to. We currently only support `metered` prices that have a [Billing Meter](https://docs.stripe.com/api/billing/meter) attached to them.
	ApplicabilityConfig *BillingCreditGrantCreateApplicabilityConfigParams `form:"applicability_config" json:"applicability_config"`
	// The category of this credit grant. It defaults to `paid` if not specified.
	Category *string `form:"category" json:"category,omitempty"`
	// ID of the customer receiving the billing credits.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// ID of the account representing the customer receiving the billing credits.
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// The time when the billing credits become effective-when they're eligible for use. It defaults to the current timestamp if not specified.
	EffectiveAt *int64 `form:"effective_at" json:"effective_at,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The time when the billing credits expire. If not specified, the billing credits don't expire.
	ExpiresAt *int64 `form:"expires_at" json:"expires_at,omitempty"`
	// Set of key-value pairs that you can attach to an object. You can use this to store additional information about the object (for example, cost basis) in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// A descriptive name shown in the Dashboard.
	Name *string `form:"name" json:"name,omitempty"`
	// The desired priority for applying this credit grant. If not specified, it will be set to the default value of 50. The highest priority is 0 and the lowest is 100.
	Priority *int64 `form:"priority" json:"priority,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *BillingCreditGrantCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *BillingCreditGrantCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves a credit grant.
type BillingCreditGrantRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *BillingCreditGrantRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Updates a credit grant.
type BillingCreditGrantUpdateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The time when the billing credits created by this credit grant expire. If set to empty, the billing credits never expire.
	ExpiresAt *int64 `form:"expires_at" json:"expires_at,omitempty"`
	// Set of key-value pairs you can attach to an object. You can use this to store additional information about the object (for example, cost basis) in a structured format.
	Metadata    map[string]string                          `form:"metadata" json:"metadata,omitempty"`
	UnsetFields []BillingCreditGrantUpdateParamsUnsetField `form:"-" json:"-"`
}

// BillingCreditGrantUpdateParamsUnsetField is the list of fields that can be cleared/unset on BillingCreditGrantUpdateParams.
type BillingCreditGrantUpdateParamsUnsetField string

const (
	BillingCreditGrantUpdateParamsUnsetFieldExpiresAt BillingCreditGrantUpdateParamsUnsetField = "expires_at"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *BillingCreditGrantUpdateParams) AddUnsetField(field BillingCreditGrantUpdateParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// AddExpand appends a new field to expand.
func (p *BillingCreditGrantUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *BillingCreditGrantUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The custom pricing unit object.
type BillingCreditGrantAmountCustomPricingUnitCustomPricingUnitDetails struct {
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The name of the custom pricing unit.
	DisplayName string `json:"display_name"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// A lookup key for the custom pricing unit.
	LookupKey string `json:"lookup_key"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The status of the custom pricing unit.
	Status string `json:"status"`
}

// The custom pricing unit amount.
type BillingCreditGrantAmountCustomPricingUnit struct {
	// The custom pricing unit object.
	CustomPricingUnitDetails *BillingCreditGrantAmountCustomPricingUnitCustomPricingUnitDetails `json:"custom_pricing_unit_details"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// A positive integer representing the amount.
	Value float64 `json:"value,string"`
}

// The monetary amount.
type BillingCreditGrantAmountMonetary struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount.
	Value int64 `json:"value"`
}
type BillingCreditGrantAmount struct {
	// The custom pricing unit amount.
	CustomPricingUnit *BillingCreditGrantAmountCustomPricingUnit `json:"custom_pricing_unit,omitempty"`
	// The monetary amount.
	Monetary *BillingCreditGrantAmountMonetary `json:"monetary"`
	// The type of this amount. We currently only support `monetary` billing credits.
	Type BillingCreditGrantAmountType `json:"type"`
}

// The billable items that credit grants can apply to. We currently only support metered billable items. Cannot be used in combination with `price_type` or `prices`.
type BillingCreditGrantApplicabilityConfigScopeBillableItem struct {
	// Unique identifier for the object.
	ID string `json:"id"`
}

// The prices that credit grants can apply to. We currently only support `metered` prices. This refers to prices that have a [Billing Meter](https://docs.stripe.com/api/billing/meter) attached to them. Cannot be used in combination with `price_type`.
type BillingCreditGrantApplicabilityConfigScopePrice struct {
	// Unique identifier for the object.
	ID string `json:"id"`
}
type BillingCreditGrantApplicabilityConfigScope struct {
	// The billable items that credit grants can apply to. We currently only support metered billable items. Cannot be used in combination with `price_type` or `prices`.
	BillableItems []*BillingCreditGrantApplicabilityConfigScopeBillableItem `json:"billable_items,omitempty"`
	// The prices that credit grants can apply to. We currently only support `metered` prices. This refers to prices that have a [Billing Meter](https://docs.stripe.com/api/billing/meter) attached to them. Cannot be used in combination with `price_type`.
	Prices []*BillingCreditGrantApplicabilityConfigScopePrice `json:"prices,omitempty"`
	// The price type that credit grants can apply to. We currently only support the `metered` price type. This refers to prices that have a [Billing Meter](https://docs.stripe.com/api/billing/meter) attached to them. Cannot be used in combination with `prices`.
	PriceType BillingCreditGrantApplicabilityConfigScopePriceType `json:"price_type,omitempty"`
}
type BillingCreditGrantApplicabilityConfig struct {
	Scope *BillingCreditGrantApplicabilityConfigScope `json:"scope"`
}

// A credit grant is an API resource that documents the allocation of some billing credits to a customer.
//
// Related guide: [Billing credits](https://docs.stripe.com/billing/subscriptions/usage-based/billing-credits)
type BillingCreditGrant struct {
	APIResource
	Amount              *BillingCreditGrantAmount              `json:"amount"`
	ApplicabilityConfig *BillingCreditGrantApplicabilityConfig `json:"applicability_config"`
	// The category of this credit grant. This is for tracking purposes and isn't displayed to the customer.
	Category BillingCreditGrantCategory `json:"category"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// ID of the customer receiving the billing credits.
	Customer *Customer `json:"customer"`
	// ID of the account representing the customer receiving the billing credits
	CustomerAccount string `json:"customer_account"`
	// The time when the billing credits become effective-when they're eligible for use.
	EffectiveAt int64 `json:"effective_at"`
	// The time when the billing credits expire. If not present, the billing credits don't expire.
	ExpiresAt int64 `json:"expires_at"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// A descriptive name shown in dashboard.
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The priority for applying this credit grant. The highest priority is 0 and the lowest is 100.
	Priority int64 `json:"priority,omitempty"`
	// ID of the test clock this credit grant belongs to.
	TestClock *TestHelpersTestClock `json:"test_clock"`
	// Time at which the object was last updated. Measured in seconds since the Unix epoch.
	Updated int64 `json:"updated"`
	// The time when this credit grant was voided. If not present, the credit grant hasn't been voided.
	VoidedAt int64 `json:"voided_at"`
}

// BillingCreditGrantList is a list of CreditGrants as retrieved from a list endpoint.
type BillingCreditGrantList struct {
	APIResource
	ListMeta
	Data []*BillingCreditGrant `json:"data"`
}

// UnmarshalJSON handles deserialization of a BillingCreditGrant.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (b *BillingCreditGrant) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		b.ID = id
		return nil
	}

	type billingCreditGrant BillingCreditGrant
	var v billingCreditGrant
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*b = BillingCreditGrant(v)
	return nil
}
