//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieve the tax integration configuration for this account.
type V2TaxIntegrationConfigurationParams struct {
	Params `form:"*"`
	// Configuration for Checkout Sessions automatic tax behavior.
	CheckoutSessions *V2TaxIntegrationConfigurationCheckoutSessionsParams `form:"checkout_sessions" json:"checkout_sessions,omitempty"`
}

// Configuration for Checkout Sessions automatic tax behavior.
type V2TaxIntegrationConfigurationCheckoutSessionsParams struct {
	// Controls the default value of automatic_tax[enabled] on new Checkout Sessions.
	AutomaticTaxDefaultValue *string `form:"automatic_tax_default_value" json:"automatic_tax_default_value"`
}

// Retrieve the tax integration configuration for this account.
type V2TaxIntegrationConfigurationRetrieveParams struct {
	Params `form:"*"`
}

// Configuration for Checkout Sessions automatic tax behavior.
type V2TaxIntegrationConfigurationUpdateCheckoutSessionsParams struct {
	// Controls the default value of automatic_tax[enabled] on new Checkout Sessions.
	AutomaticTaxDefaultValue *string `form:"automatic_tax_default_value" json:"automatic_tax_default_value"`
}

// Update the tax integration configuration for this account.
type V2TaxIntegrationConfigurationUpdateParams struct {
	Params `form:"*"`
	// Configuration for Checkout Sessions automatic tax behavior.
	CheckoutSessions *V2TaxIntegrationConfigurationUpdateCheckoutSessionsParams `form:"checkout_sessions" json:"checkout_sessions,omitempty"`
}
