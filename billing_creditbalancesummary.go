//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The type of this amount. We currently only support `monetary` billing credits.
type BillingCreditBalanceSummaryBalanceAvailableBalanceType string

// List of values that BillingCreditBalanceSummaryBalanceAvailableBalanceType can take
const (
	BillingCreditBalanceSummaryBalanceAvailableBalanceTypeCustomPricingUnit BillingCreditBalanceSummaryBalanceAvailableBalanceType = "custom_pricing_unit"
	BillingCreditBalanceSummaryBalanceAvailableBalanceTypeMonetary          BillingCreditBalanceSummaryBalanceAvailableBalanceType = "monetary"
)

// The type of this amount. We currently only support `monetary` billing credits.
type BillingCreditBalanceSummaryBalanceLedgerBalanceType string

// List of values that BillingCreditBalanceSummaryBalanceLedgerBalanceType can take
const (
	BillingCreditBalanceSummaryBalanceLedgerBalanceTypeCustomPricingUnit BillingCreditBalanceSummaryBalanceLedgerBalanceType = "custom_pricing_unit"
	BillingCreditBalanceSummaryBalanceLedgerBalanceTypeMonetary          BillingCreditBalanceSummaryBalanceLedgerBalanceType = "monetary"
)

// A list of billable items that the credit grant can apply to. We currently only support metered billable items. Cannot be used in combination with `price_type` or `prices`.
type BillingCreditBalanceSummaryFilterApplicabilityScopeBillableItemParams struct {
	// The billable item ID this credit grant should apply to.
	ID *string `form:"id"`
}

// A list of prices that the credit grant can apply to. We currently only support the `metered` prices. Cannot be used in combination with `price_type`.
type BillingCreditBalanceSummaryFilterApplicabilityScopePriceParams struct {
	// The price ID this credit grant should apply to.
	ID *string `form:"id"`
}

// The billing credit applicability scope for which to fetch credit balance summary.
type BillingCreditBalanceSummaryFilterApplicabilityScopeParams struct {
	// A list of billable items that the credit grant can apply to. We currently only support metered billable items. Cannot be used in combination with `price_type` or `prices`.
	BillableItems []*BillingCreditBalanceSummaryFilterApplicabilityScopeBillableItemParams `form:"billable_items"`
	// A list of prices that the credit grant can apply to. We currently only support the `metered` prices. Cannot be used in combination with `price_type`.
	Prices []*BillingCreditBalanceSummaryFilterApplicabilityScopePriceParams `form:"prices"`
	// The price type that credit grants can apply to. We currently only support the `metered` price type. Cannot be used in combination with `prices`.
	PriceType *string `form:"price_type"`
}

// The filter criteria for the credit balance summary.
type BillingCreditBalanceSummaryFilterParams struct {
	// The billing credit applicability scope for which to fetch credit balance summary.
	ApplicabilityScope *BillingCreditBalanceSummaryFilterApplicabilityScopeParams `form:"applicability_scope"`
	// The credit grant for which to fetch credit balance summary.
	CreditGrant *string `form:"credit_grant"`
	// Specify the type of this filter.
	Type *string `form:"type"`
}

// Retrieves the credit balance summary for a customer.
type BillingCreditBalanceSummaryParams struct {
	Params `form:"*"`
	// The customer for which to fetch credit balance summary.
	Customer *string `form:"customer"`
	// The account for which to fetch credit balance summary.
	CustomerAccount *string `form:"customer_account"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The filter criteria for the credit balance summary.
	Filter *BillingCreditBalanceSummaryFilterParams `form:"filter"`
}

// AddExpand appends a new field to expand.
func (p *BillingCreditBalanceSummaryParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A list of billable items that the credit grant can apply to. We currently only support metered billable items. Cannot be used in combination with `price_type` or `prices`.
type BillingCreditBalanceSummaryRetrieveFilterApplicabilityScopeBillableItemParams struct {
	// The billable item ID this credit grant should apply to.
	ID *string `form:"id"`
}

// A list of prices that the credit grant can apply to. We currently only support the `metered` prices. Cannot be used in combination with `price_type`.
type BillingCreditBalanceSummaryRetrieveFilterApplicabilityScopePriceParams struct {
	// The price ID this credit grant should apply to.
	ID *string `form:"id"`
}

// The billing credit applicability scope for which to fetch credit balance summary.
type BillingCreditBalanceSummaryRetrieveFilterApplicabilityScopeParams struct {
	// A list of billable items that the credit grant can apply to. We currently only support metered billable items. Cannot be used in combination with `price_type` or `prices`.
	BillableItems []*BillingCreditBalanceSummaryRetrieveFilterApplicabilityScopeBillableItemParams `form:"billable_items"`
	// A list of prices that the credit grant can apply to. We currently only support the `metered` prices. Cannot be used in combination with `price_type`.
	Prices []*BillingCreditBalanceSummaryRetrieveFilterApplicabilityScopePriceParams `form:"prices"`
	// The price type that credit grants can apply to. We currently only support the `metered` price type. Cannot be used in combination with `prices`.
	PriceType *string `form:"price_type"`
}

// The filter criteria for the credit balance summary.
type BillingCreditBalanceSummaryRetrieveFilterParams struct {
	// The billing credit applicability scope for which to fetch credit balance summary.
	ApplicabilityScope *BillingCreditBalanceSummaryRetrieveFilterApplicabilityScopeParams `form:"applicability_scope"`
	// The credit grant for which to fetch credit balance summary.
	CreditGrant *string `form:"credit_grant"`
	// Specify the type of this filter.
	Type *string `form:"type"`
}

// Retrieves the credit balance summary for a customer.
type BillingCreditBalanceSummaryRetrieveParams struct {
	Params `form:"*"`
	// The customer for which to fetch credit balance summary.
	Customer *string `form:"customer"`
	// The account for which to fetch credit balance summary.
	CustomerAccount *string `form:"customer_account"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// The filter criteria for the credit balance summary.
	Filter *BillingCreditBalanceSummaryRetrieveFilterParams `form:"filter"`
}

// AddExpand appends a new field to expand.
func (p *BillingCreditBalanceSummaryRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The custom pricing unit object.
type BillingCreditBalanceSummaryBalanceAvailableBalanceCustomPricingUnitCustomPricingUnitDetails struct {
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
type BillingCreditBalanceSummaryBalanceAvailableBalanceCustomPricingUnit struct {
	// The custom pricing unit object.
	CustomPricingUnitDetails *BillingCreditBalanceSummaryBalanceAvailableBalanceCustomPricingUnitCustomPricingUnitDetails `json:"custom_pricing_unit_details"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// A positive integer representing the amount.
	Value float64 `json:"value,string"`
}

// The monetary amount.
type BillingCreditBalanceSummaryBalanceAvailableBalanceMonetary struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount.
	Value int64 `json:"value"`
}
type BillingCreditBalanceSummaryBalanceAvailableBalance struct {
	// The custom pricing unit amount.
	CustomPricingUnit *BillingCreditBalanceSummaryBalanceAvailableBalanceCustomPricingUnit `json:"custom_pricing_unit"`
	// The monetary amount.
	Monetary *BillingCreditBalanceSummaryBalanceAvailableBalanceMonetary `json:"monetary"`
	// The type of this amount. We currently only support `monetary` billing credits.
	Type BillingCreditBalanceSummaryBalanceAvailableBalanceType `json:"type"`
}

// The custom pricing unit object.
type BillingCreditBalanceSummaryBalanceLedgerBalanceCustomPricingUnitCustomPricingUnitDetails struct {
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
type BillingCreditBalanceSummaryBalanceLedgerBalanceCustomPricingUnit struct {
	// The custom pricing unit object.
	CustomPricingUnitDetails *BillingCreditBalanceSummaryBalanceLedgerBalanceCustomPricingUnitCustomPricingUnitDetails `json:"custom_pricing_unit_details"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// A positive integer representing the amount.
	Value float64 `json:"value,string"`
}

// The monetary amount.
type BillingCreditBalanceSummaryBalanceLedgerBalanceMonetary struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// A positive integer representing the amount.
	Value int64 `json:"value"`
}
type BillingCreditBalanceSummaryBalanceLedgerBalance struct {
	// The custom pricing unit amount.
	CustomPricingUnit *BillingCreditBalanceSummaryBalanceLedgerBalanceCustomPricingUnit `json:"custom_pricing_unit"`
	// The monetary amount.
	Monetary *BillingCreditBalanceSummaryBalanceLedgerBalanceMonetary `json:"monetary"`
	// The type of this amount. We currently only support `monetary` billing credits.
	Type BillingCreditBalanceSummaryBalanceLedgerBalanceType `json:"type"`
}

// The billing credit balances. One entry per credit grant currency. If a customer only has credit grants in a single currency, then this will have a single balance entry.
type BillingCreditBalanceSummaryBalance struct {
	AvailableBalance *BillingCreditBalanceSummaryBalanceAvailableBalance `json:"available_balance"`
	LedgerBalance    *BillingCreditBalanceSummaryBalanceLedgerBalance    `json:"ledger_balance"`
}

// Indicates the billing credit balance for billing credits granted to a customer.
type BillingCreditBalanceSummary struct {
	APIResource
	// The billing credit balances. One entry per credit grant currency. If a customer only has credit grants in a single currency, then this will have a single balance entry.
	Balances []*BillingCreditBalanceSummaryBalance `json:"balances"`
	// The customer the balance is for.
	Customer *Customer `json:"customer"`
	// The account the balance is for.
	CustomerAccount string `json:"customer_account"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
