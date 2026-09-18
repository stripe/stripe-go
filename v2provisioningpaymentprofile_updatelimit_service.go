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

// v2ProvisioningPaymentProfileUpdateLimitService is used to invoke updatelimit related APIs.
type v2ProvisioningPaymentProfileUpdateLimitService struct {
	B   Backend
	Key string
}

// Updates the usage limit on the payment profile for a provider.
func (c v2ProvisioningPaymentProfileUpdateLimitService) Update(ctx context.Context, params *V2ProvisioningPaymentProfileUpdateLimitUpdateParams) (*V2ProvisioningPaymentProfile, error) {
	if params == nil {
		params = &V2ProvisioningPaymentProfileUpdateLimitUpdateParams{}
	}
	params.Context = ctx
	paymentprofile := &V2ProvisioningPaymentProfile{}
	err := c.B.Call(
		http.MethodPost, "/v2/provisioning/payment_profile/update_limit", c.Key, params, paymentprofile)
	return paymentprofile, err
}
