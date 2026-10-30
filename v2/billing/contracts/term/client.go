//
//
// File generated from our OpenAPI spec
//
//

// Package term provides the term related APIs
package term

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
)

// Client is used to invoke term related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieve a term on a contract.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) RetrieveContractTerm(id string, params *stripe.V2BillingContractsTermRetrieveContractTermParams) (*stripe.V2BillingContractTerm, error) {
	path := stripe.FormatURLPath(
		"/v2/billing/contracts/%s/terms/%s", stripe.StringValue(
			params.ContractID), id)
	contractterm := &stripe.V2BillingContractTerm{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, contractterm)
	return contractterm, err
}

// List the terms on a contract, optionally filtered by type.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) ListContractTerms(listParams *stripe.V2BillingContractsTermListContractTermsParams) stripe.Seq2[*stripe.V2BillingContractTerm, error] {
	if listParams == nil {
		listParams = &stripe.V2BillingContractsTermListContractTermsParams{}
	}
	path := stripe.FormatURLPath(
		"/v2/billing/contracts/%s/terms", stripe.StringValue(listParams.ContractID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingContractTerm], error) {
		page := &stripe.V2Page[*stripe.V2BillingContractTerm]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
