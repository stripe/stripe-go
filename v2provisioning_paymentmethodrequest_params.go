//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Usage limit to apply to the requested payment method.
type V2ProvisioningPaymentMethodRequestUsageLimitsParams struct {
	// Three-letter ISO currency code for `max_amount`.
	Currency *string `form:"currency" json:"currency"`
	// Maximum amount that can be charged per recurring interval.
	MaxAmount *int64 `form:"max_amount" json:"max_amount,string"`
	// Interval over which `max_amount` applies.
	RecurringInterval *string `form:"recurring_interval" json:"recurring_interval"`
}

// Creates a request for a customer to authorize a new payment method.
type V2ProvisioningPaymentMethodRequestParams struct {
	Params `form:"*"`
	// Whether the billing operation should use Stripe live-mode objects. When omitted, this
	// resolves from the authenticated request context.
	Livemode *bool `form:"livemode" json:"livemode,omitempty"`
	// Owner of the requested payment method.
	PaymentMethodOwner *string `form:"payment_method_owner" json:"payment_method_owner,omitempty"`
	// Connected account to source the payment method from.
	SourceAccount *string `form:"source_account" json:"source_account,omitempty"`
	// Customer to source the payment method from.
	SourceCustomer *string `form:"source_customer" json:"source_customer,omitempty"`
	// Existing payment method to reuse instead of collecting a new one.
	SourcePaymentMethod *string `form:"source_payment_method" json:"source_payment_method,omitempty"`
	// Usage limit to apply to the requested payment method.
	UsageLimits *V2ProvisioningPaymentMethodRequestUsageLimitsParams `form:"usage_limits" json:"usage_limits,omitempty"`
}

// Usage limit to apply to the requested payment method.
type V2ProvisioningPaymentMethodRequestCreateUsageLimitsParams struct {
	// Three-letter ISO currency code for `max_amount`.
	Currency *string `form:"currency" json:"currency"`
	// Maximum amount that can be charged per recurring interval.
	MaxAmount *int64 `form:"max_amount" json:"max_amount,string"`
	// Interval over which `max_amount` applies.
	RecurringInterval *string `form:"recurring_interval" json:"recurring_interval"`
}

// Creates a request for a customer to authorize a new payment method.
type V2ProvisioningPaymentMethodRequestCreateParams struct {
	Params `form:"*"`
	// Whether the billing operation should use Stripe live-mode objects. When omitted, this
	// resolves from the authenticated request context.
	Livemode *bool `form:"livemode" json:"livemode,omitempty"`
	// Owner of the requested payment method.
	PaymentMethodOwner *string `form:"payment_method_owner" json:"payment_method_owner,omitempty"`
	// Connected account to source the payment method from.
	SourceAccount *string `form:"source_account" json:"source_account,omitempty"`
	// Customer to source the payment method from.
	SourceCustomer *string `form:"source_customer" json:"source_customer,omitempty"`
	// Existing payment method to reuse instead of collecting a new one.
	SourcePaymentMethod *string `form:"source_payment_method" json:"source_payment_method,omitempty"`
	// Usage limit to apply to the requested payment method.
	UsageLimits *V2ProvisioningPaymentMethodRequestCreateUsageLimitsParams `form:"usage_limits" json:"usage_limits,omitempty"`
}
