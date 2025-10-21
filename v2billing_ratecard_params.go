//
//
// File generated from our OpenAPI spec
//
//

package stripe

// List all Rate Card objects.
type V2BillingRateCardListParams struct {
	Params `form:"*"`
	// Optionally filter to active/inactive RateCards.
	Active *bool `form:"active" json:"active,omitempty"`
	// Optionally set the maximum number of results per page. Defaults to 20.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
	// Filter by lookup keys.
	// You can specify up to 10 lookup keys.
	LookupKeys []*string `form:"lookup_keys,flat_array" json:"lookup_keys,omitempty"`
}

// Create a Rate Card object.
type V2BillingRateCardParams struct {
	Params `form:"*"`
	// Sets whether the RateCard is active. Inactive RateCards cannot be used in new activations or have new rates added.
	Active *bool `form:"active" json:"active,omitempty"`
	// The currency of this RateCard.
	Currency *string `form:"currency" json:"currency,omitempty"`
	// A customer-facing name for the RateCard.
	// This name is used in Stripe-hosted products like the Customer Portal and Checkout. It does not show up on Invoices.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Changes the version that new RateCard activations will use. Providing `live_version = "latest"` will set the
	// RateCard's `live_version` to its latest version.
	LiveVersion *string `form:"live_version" json:"live_version,omitempty"`
	// An internal key you can use to search for a particular RateCard. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The interval for assessing service. For example, a monthly RateCard with a rate of $1 for the first 10 "workloads"
	// and $2 thereafter means "$1 per workload up to 10 workloads during a month of service." This is similar to but
	// distinct from billing interval; the service interval deals with the rate at which the customer accumulates fees,
	// while the billing interval in Cadence deals with the rate the customer is billed.
	ServiceInterval *string `form:"service_interval" json:"service_interval,omitempty"`
	// The length of the interval for assessing service. For example, set this to 3 and `service_interval` to `"month"` in
	// order to specify quarterly service.
	ServiceIntervalCount *int64 `form:"service_interval_count" json:"service_interval_count,omitempty"`
	// The Stripe Tax tax behavior - whether the rates are inclusive or exclusive of tax.
	TaxBehavior *string `form:"tax_behavior" json:"tax_behavior,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingRateCardParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Create a Rate Card object.
type V2BillingRateCardCreateParams struct {
	Params `form:"*"`
	// The currency of this RateCard.
	Currency *string `form:"currency" json:"currency"`
	// A customer-facing name for the RateCard.
	// This name is used in Stripe-hosted products like the Customer Portal and Checkout. It does not show up on Invoices.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name"`
	// An internal key you can use to search for a particular RateCard. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The interval for assessing service. For example, a monthly RateCard with a rate of $1 for the first 10 "workloads"
	// and $2 thereafter means "$1 per workload up to 10 workloads during a month of service." This is similar to but
	// distinct from billing interval; the service interval deals with the rate at which the customer accumulates fees,
	// while the billing interval in Cadence deals with the rate the customer is billed.
	ServiceInterval *string `form:"service_interval" json:"service_interval"`
	// The length of the interval for assessing service. For example, set this to 3 and `service_interval` to `"month"` in
	// order to specify quarterly service.
	ServiceIntervalCount *int64 `form:"service_interval_count" json:"service_interval_count"`
	// The Stripe Tax tax behavior - whether the rates are inclusive or exclusive of tax.
	TaxBehavior *string `form:"tax_behavior" json:"tax_behavior"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingRateCardCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieve the latest version of a Rate Card object.
type V2BillingRateCardRetrieveParams struct {
	Params `form:"*"`
}

// Update a Rate Card object.
type V2BillingRateCardUpdateParams struct {
	Params `form:"*"`
	// Sets whether the RateCard is active. Inactive RateCards cannot be used in new activations or have new rates added.
	Active *bool `form:"active" json:"active,omitempty"`
	// A customer-facing name for the RateCard.
	// This name is used in Stripe-hosted products like the Customer Portal and Checkout. It does not show up on Invoices.
	// Maximum length of 250 characters.
	DisplayName *string `form:"display_name" json:"display_name,omitempty"`
	// Changes the version that new RateCard activations will use. Providing `live_version = "latest"` will set the
	// RateCard's `live_version` to its latest version.
	LiveVersion *string `form:"live_version" json:"live_version,omitempty"`
	// An internal key you can use to search for a particular RateCard. Maximum length of 200 characters.
	LookupKey *string `form:"lookup_key" json:"lookup_key,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *V2BillingRateCardUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}
