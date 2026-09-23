//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Status of the payment method request.
type V2ProvisioningPaymentMethodRequestStatus string

// List of values that V2ProvisioningPaymentMethodRequestStatus can take
const (
	V2ProvisioningPaymentMethodRequestStatusCheckoutInitiated V2ProvisioningPaymentMethodRequestStatus = "checkout_initiated"
	V2ProvisioningPaymentMethodRequestStatusComplete          V2ProvisioningPaymentMethodRequestStatus = "complete"
)

// The result of an in-progress request for a customer to authorize a new payment method.
type V2ProvisioningPaymentMethodRequest struct {
	APIResource
	// URL for the customer to complete payment method authorization.
	CheckoutSessionURL string `json:"checkout_session_url"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Status of the payment method request.
	Status V2ProvisioningPaymentMethodRequestStatus `json:"status"`
}
