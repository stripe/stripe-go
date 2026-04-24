//
//
// File generated from our OpenAPI spec
//
//

// Package agreement provides the agreement related APIs
package agreement

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke agreement related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a new Agreement.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2OrchestratedCommerceAgreementParams) (*stripe.V2OrchestratedCommerceAgreement, error) {
	agreement := &stripe.V2OrchestratedCommerceAgreement{}
	err := c.B.Call(
		http.MethodPost, "/v2/orchestrated_commerce/agreements", c.Key, params, agreement)
	return agreement, err
}

// Retrieve an Agreement by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2OrchestratedCommerceAgreementParams) (*stripe.V2OrchestratedCommerceAgreement, error) {
	path := stripe.FormatURLPath("/v2/orchestrated_commerce/agreements/%s", id)
	agreement := &stripe.V2OrchestratedCommerceAgreement{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, agreement)
	return agreement, err
}

// Confirm an Agreement.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Confirm(id string, params *stripe.V2OrchestratedCommerceAgreementConfirmParams) (*stripe.V2OrchestratedCommerceAgreement, error) {
	path := stripe.FormatURLPath(
		"/v2/orchestrated_commerce/agreements/%s/confirm", id)
	agreement := &stripe.V2OrchestratedCommerceAgreement{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, agreement)
	return agreement, err
}

// Terminate an Agreement.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Terminate(id string, params *stripe.V2OrchestratedCommerceAgreementTerminateParams) (*stripe.V2OrchestratedCommerceAgreement, error) {
	path := stripe.FormatURLPath(
		"/v2/orchestrated_commerce/agreements/%s/terminate", id)
	agreement := &stripe.V2OrchestratedCommerceAgreement{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, agreement)
	return agreement, err
}

// List Agreements for the profile associated with the authenticated merchant.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2OrchestratedCommerceAgreementListParams) stripe.Seq2[*stripe.V2OrchestratedCommerceAgreement, error] {
	if listParams == nil {
		listParams = &stripe.V2OrchestratedCommerceAgreementListParams{}
	}
	return stripe.NewV2List("/v2/orchestrated_commerce/agreements", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2OrchestratedCommerceAgreement], error) {
		page := &stripe.V2Page[*stripe.V2OrchestratedCommerceAgreement]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
