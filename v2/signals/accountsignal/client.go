//
//
// File generated from our OpenAPI spec
//
//

// Package accountsignal provides the accountsignal related APIs
package accountsignal

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke accountsignal related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an AccountSignal by its ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2SignalsAccountSignalParams) (*stripe.V2SignalsAccountSignal, error) {
	path := stripe.FormatURLPath("/v2/signals/account_signals/%s", id)
	accountsignal := &stripe.V2SignalsAccountSignal{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountsignal)
	return accountsignal, err
}

// Lists AccountSignals for a given account or customer, filtered by signal type.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2SignalsAccountSignalListParams) stripe.Seq2[*stripe.V2SignalsAccountSignal, error] {
	if listParams == nil {
		listParams = &stripe.V2SignalsAccountSignalListParams{}
	}
	return stripe.NewV2List("/v2/signals/account_signals", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2SignalsAccountSignal], error) {
		page := &stripe.V2Page[*stripe.V2SignalsAccountSignal]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
