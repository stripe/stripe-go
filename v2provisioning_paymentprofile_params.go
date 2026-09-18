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

// Retrieves the payment profile for the current project.
type V2ProvisioningPaymentProfileRetrieveParams struct {
	Params `form:"*"`
	// Whether the billing operation should use Stripe live-mode objects. When omitted, this
	// resolves from the authenticated request context.
	Livemode *bool `form:"livemode" json:"livemode,omitempty"`
}
