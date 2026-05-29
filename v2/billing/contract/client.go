//
//
// File generated from our OpenAPI spec
//
//

// Package contract provides the contract related APIs
package contract

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke contract related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Contract object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2BillingContractParams) (*stripe.V2BillingContract, error) {
	contract := &stripe.V2BillingContract{}
	err := c.B.Call(
		http.MethodPost, "/v2/billing/contracts", c.Key, params, contract)
	return contract, err
}

// Retrieve a Contract object by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2BillingContractParams) (*stripe.V2BillingContract, error) {
	path := stripe.FormatURLPath("/v2/billing/contracts/%s", id)
	contract := &stripe.V2BillingContract{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, contract)
	return contract, err
}

// Update a Contract object by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2BillingContractParams) (*stripe.V2BillingContract, error) {
	path := stripe.FormatURLPath("/v2/billing/contracts/%s", id)
	contract := &stripe.V2BillingContract{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, contract)
	return contract, err
}

// Activate a Draft Contract object by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Activate(id string, params *stripe.V2BillingContractActivateParams) (*stripe.V2BillingContract, error) {
	path := stripe.FormatURLPath("/v2/billing/contracts/%s/activate", id)
	contract := &stripe.V2BillingContract{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, contract)
	return contract, err
}

// Cancel a Contract object by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Cancel(id string, params *stripe.V2BillingContractCancelParams) (*stripe.V2BillingContract, error) {
	path := stripe.FormatURLPath("/v2/billing/contracts/%s/cancel", id)
	contract := &stripe.V2BillingContract{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, contract)
	return contract, err
}

// List Contract objects with pagination.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2BillingContractListParams) stripe.Seq2[*stripe.V2BillingContract, error] {
	if listParams == nil {
		listParams = &stripe.V2BillingContractListParams{}
	}
	return stripe.NewV2List("/v2/billing/contracts", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2BillingContract], error) {
		page := &stripe.V2Page[*stripe.V2BillingContract]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
