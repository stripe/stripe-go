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

// v2BillingLicenseFeeSubscriptionService is used to invoke licensefeesubscription related APIs.
type v2BillingLicenseFeeSubscriptionService struct {
	B   Backend
	Key string
}

// Retrieve a License Fee Subscription object.
func (c v2BillingLicenseFeeSubscriptionService) Retrieve(ctx context.Context, id string, params *V2BillingLicenseFeeSubscriptionRetrieveParams) (*V2BillingLicenseFeeSubscription, error) {
	if params == nil {
		params = &V2BillingLicenseFeeSubscriptionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/billing/license_fee_subscriptions/%s", id)
	licensefeesubscription := &V2BillingLicenseFeeSubscription{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, licensefeesubscription)
	return licensefeesubscription, err
}
