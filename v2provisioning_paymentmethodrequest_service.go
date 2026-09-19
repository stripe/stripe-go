//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v2ProvisioningPaymentMethodRequestService is used to invoke paymentmethodrequest related APIs.
type v2ProvisioningPaymentMethodRequestService struct {
	B   Backend
	Key string
}

// Creates a request for a customer to authorize a new payment method.
func (c v2ProvisioningPaymentMethodRequestService) Create(ctx context.Context, params *V2ProvisioningPaymentMethodRequestCreateParams) (*V2ProvisioningPaymentMethodRequest, error) {
	if params == nil {
		params = &V2ProvisioningPaymentMethodRequestCreateParams{}
	}
	params.Context = ctx
	paymentmethodrequest := &V2ProvisioningPaymentMethodRequest{}
	err := c.B.Call(
		http.MethodPost, "/v2/provisioning/payment_method_requests", c.Key, params, paymentmethodrequest)
	return paymentmethodrequest, err
}
