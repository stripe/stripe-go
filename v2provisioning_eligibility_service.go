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

// v2ProvisioningEligibilityService is used to invoke eligibility related APIs.
type v2ProvisioningEligibilityService struct {
	B   Backend
	Key string
}

// Checks whether a project is eligible to provision resources with a provider, including
// any outstanding KYC requirements that must be satisfied first.
func (c v2ProvisioningEligibilityService) Retrieve(ctx context.Context, params *V2ProvisioningEligibilityRetrieveParams) (*V2ProvisioningEligibility, error) {
	if params == nil {
		params = &V2ProvisioningEligibilityRetrieveParams{}
	}
	params.Context = ctx
	eligibility := &V2ProvisioningEligibility{}
	err := c.B.Call(
		http.MethodGet, "/v2/provisioning/eligibility", c.Key, params, eligibility)
	return eligibility, err
}
