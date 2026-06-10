//
//
// File generated from our OpenAPI spec
//
//

// Package quantitychange provides the quantitychange related APIs
package quantitychange

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke quantitychange related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// List license quantity changes for a contract given a license pricing ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListQuantityChanges(listParams *stripe.V2BillingContractsLicensePricingQuantityChangeListQuantityChangesParams) stripe.Seq2[*stripe.V2BillingContractLicensePricingQuantityChange, error] {
	if listParams == nil {
		listParams = &stripe.V2BillingContractsLicensePricingQuantityChangeListQuantityChangesParams{}
	}
	path := stripe.FormatURLPath(
		"/v2/billing/contracts/%s/license_pricing/%s/quantity_changes", stripe.StringValue(
			listParams.ContractID), stripe.StringValue(listParams.LicensePricingID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingContractLicensePricingQuantityChange], error) {
		page := &stripe.V2Page[*stripe.V2BillingContractLicensePricingQuantityChange]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
