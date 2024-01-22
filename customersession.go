//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Configuration for buy button.
type CustomerSessionComponentsBuyButtonParams struct {
	// Whether the buy button is enabled.
	Enabled *bool `form:"enabled"`
}

// Configuration for the pricing table.
type CustomerSessionComponentsPricingTableParams struct {
	// Whether the pricing table is enabled.
	Enabled *bool `form:"enabled"`
}

// Configuration for each component. Exactly 1 component must be enabled.
type CustomerSessionComponentsParams struct {
	// Configuration for buy button.
	BuyButton *CustomerSessionComponentsBuyButtonParams `form:"buy_button"`
	// Configuration for the pricing table.
	PricingTable *CustomerSessionComponentsPricingTableParams `form:"pricing_table"`
}

// Creates a customer session object that includes a single-use client secret that you can use on your front-end to grant client-side API access for certain customer resources.
type CustomerSessionParams struct {
	Params `form:"*"`
	// Configuration for each component. Exactly 1 component must be enabled.
	Components *CustomerSessionComponentsParams `form:"components"`
	// The ID of an existing customer for which to create the customer session.
	Customer *string `form:"customer"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CustomerSessionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// This hash contains whether the buy button is enabled.
type CustomerSessionComponentsBuyButton struct {
	// Whether the buy button is enabled.
	Enabled bool `json:"enabled"`
}

// This hash contains whether the pricing table is enabled.
type CustomerSessionComponentsPricingTable struct {
	// Whether the pricing table is enabled.
	Enabled bool `json:"enabled"`
}

// Configuration for the components supported by this customer session.
type CustomerSessionComponents struct {
	// This hash contains whether the buy button is enabled.
	BuyButton *CustomerSessionComponentsBuyButton `json:"buy_button"`
	// This hash contains whether the pricing table is enabled.
	PricingTable *CustomerSessionComponentsPricingTable `json:"pricing_table"`
}

// A customer session allows you to grant client access to Stripe's frontend SDKs (like StripeJs)
// control over a customer.
type CustomerSession struct {
	APIResource
	// The client secret of this customer session. Used on the client to set up secure access to the given `customer`.
	//
	// The client secret can be used to provide access to `customer` from your frontend. It should not be stored, logged, or exposed to anyone other than the relevant customer. Make sure that you have TLS enabled on any page that includes the client secret.
	ClientSecret string `json:"client_secret"`
	// Configuration for the components supported by this customer session.
	Components *CustomerSessionComponents `json:"components"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The customer the customer session was created for.
	Customer *Customer `json:"customer"`
	// The timestamp at which this customer session will expire.
	ExpiresAt int64 `json:"expires_at"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
