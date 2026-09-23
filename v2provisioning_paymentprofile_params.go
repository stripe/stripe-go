//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves the payment profile for the current project.
type V2ProvisioningPaymentProfileParams struct {
	Params `form:"*"`
	// Whether the billing operation should use Stripe live-mode objects. When omitted, this
	// resolves from the authenticated request context.
	Livemode *bool `form:"livemode" json:"livemode,omitempty"`
}

// New usage limit to apply.
type V2ProvisioningPaymentProfileUpdateLimitUsageLimitsParams struct {
	// Three-letter ISO currency code for `max_amount`.
	Currency *string `form:"currency" json:"currency"`
	// Maximum amount that can be charged per recurring interval.
	MaxAmount *int64 `form:"max_amount" json:"max_amount,string"`
	// Interval over which `max_amount` applies.
	RecurringInterval *string `form:"recurring_interval" json:"recurring_interval"`
}

// Updates the usage limit on the payment profile for a provider.
type V2ProvisioningPaymentProfileUpdateLimitParams struct {
	Params `form:"*"`
	// Whether the billing operation should use Stripe live-mode objects. When omitted, this
	// resolves from the authenticated request context.
	Livemode *bool `form:"livemode" json:"livemode,omitempty"`
	// Provider to update the usage limit for.
	Provider *string `form:"provider" json:"provider,omitempty"`
	// New usage limit to apply.
	UsageLimits *V2ProvisioningPaymentProfileUpdateLimitUsageLimitsParams `form:"usage_limits" json:"usage_limits"`
}

// Retrieves the payment profile for the current project.
type V2ProvisioningPaymentProfileRetrieveParams struct {
	Params `form:"*"`
	// Whether the billing operation should use Stripe live-mode objects. When omitted, this
	// resolves from the authenticated request context.
	Livemode *bool `form:"livemode" json:"livemode,omitempty"`
}
