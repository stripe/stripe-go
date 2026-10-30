//
//
// File generated from our OpenAPI spec
//
//

// Package quantitychange provides the quantitychange related APIs
package quantitychange

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke quantitychange related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// List quantity changes for a licensed recurring-price term.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListContractTermQuantityChanges(listParams *stripe.V2BillingContractsTermsQuantityChangeListContractTermQuantityChangesParams) stripe.Seq2[*stripe.V2BillingContractTermQuantityChange, error] {
	if listParams == nil {
		listParams = &stripe.V2BillingContractsTermsQuantityChangeListContractTermQuantityChangesParams{}
	}
	path := stripe.FormatURLPath(
		"/v2/billing/contracts/%s/terms/%s/quantity_changes", stripe.StringValue(
			listParams.ContractID), stripe.StringValue(listParams.Term))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingContractTermQuantityChange], error) {
		page := &stripe.V2Page[*stripe.V2BillingContractTermQuantityChange]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
