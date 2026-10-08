//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Controls the default value of automatic_tax[enabled] on new Checkout Sessions.
type V2TaxIntegrationConfigurationCheckoutSessionsAutomaticTaxDefaultValue string

// List of values that V2TaxIntegrationConfigurationCheckoutSessionsAutomaticTaxDefaultValue can take
const (
	V2TaxIntegrationConfigurationCheckoutSessionsAutomaticTaxDefaultValueDisabled            V2TaxIntegrationConfigurationCheckoutSessionsAutomaticTaxDefaultValue = "disabled"
	V2TaxIntegrationConfigurationCheckoutSessionsAutomaticTaxDefaultValueEnabledWhenPossible V2TaxIntegrationConfigurationCheckoutSessionsAutomaticTaxDefaultValue = "enabled_when_possible"
)

// Controls the default value of automatic_tax[enabled] on new standalone Invoices.
type V2TaxIntegrationConfigurationInvoicesAutomaticTaxDefaultValue string

// List of values that V2TaxIntegrationConfigurationInvoicesAutomaticTaxDefaultValue can take
const (
	V2TaxIntegrationConfigurationInvoicesAutomaticTaxDefaultValueDisabled            V2TaxIntegrationConfigurationInvoicesAutomaticTaxDefaultValue = "disabled"
	V2TaxIntegrationConfigurationInvoicesAutomaticTaxDefaultValueEnabledWhenPossible V2TaxIntegrationConfigurationInvoicesAutomaticTaxDefaultValue = "enabled_when_possible"
)

// Configuration for Checkout Sessions automatic tax behavior.
type V2TaxIntegrationConfigurationCheckoutSessions struct {
	// Controls the default value of automatic_tax[enabled] on new Checkout Sessions.
	AutomaticTaxDefaultValue V2TaxIntegrationConfigurationCheckoutSessionsAutomaticTaxDefaultValue `json:"automatic_tax_default_value"`
}

// Configuration for standalone Invoices automatic tax behavior.
type V2TaxIntegrationConfigurationInvoices struct {
	// Controls the default value of automatic_tax[enabled] on new standalone Invoices.
	AutomaticTaxDefaultValue V2TaxIntegrationConfigurationInvoicesAutomaticTaxDefaultValue `json:"automatic_tax_default_value"`
}

// Per-account configuration controlling implicit behavior of Stripe Tax
// across supported integration surfaces.
type V2TaxIntegrationConfiguration struct {
	APIResource
	// Configuration for Checkout Sessions automatic tax behavior.
	CheckoutSessions *V2TaxIntegrationConfigurationCheckoutSessions `json:"checkout_sessions"`
	// Configuration for standalone Invoices automatic tax behavior.
	Invoices *V2TaxIntegrationConfigurationInvoices `json:"invoices"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
}
