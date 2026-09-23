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

// v2ProvisioningPaymentProfileService is used to invoke paymentprofile related APIs.
type v2ProvisioningPaymentProfileService struct {
	B   Backend
	Key string
}

// Retrieves the payment profile for the current project.
func (c v2ProvisioningPaymentProfileService) Retrieve(ctx context.Context, params *V2ProvisioningPaymentProfileRetrieveParams) (*V2ProvisioningPaymentProfile, error) {
	if params == nil {
		params = &V2ProvisioningPaymentProfileRetrieveParams{}
	}
	params.Context = ctx
	paymentprofile := &V2ProvisioningPaymentProfile{}
	err := c.B.Call(
		http.MethodGet, "/v2/provisioning/payment_profile", c.Key, params, paymentprofile)
	return paymentprofile, err
}

// Updates the usage limit on the payment profile for a provider.
func (c v2ProvisioningPaymentProfileService) UpdateLimit(ctx context.Context, params *V2ProvisioningPaymentProfileUpdateLimitParams) (*V2ProvisioningPaymentProfile, error) {
	if params == nil {
		params = &V2ProvisioningPaymentProfileUpdateLimitParams{}
	}
	params.Context = ctx
	paymentprofile := &V2ProvisioningPaymentProfile{}
	err := c.B.Call(
		http.MethodPost, "/v2/provisioning/payment_profile/update_limit", c.Key, params, paymentprofile)
	return paymentprofile, err
}
