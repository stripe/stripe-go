//
//
// File generated from our OpenAPI spec
//
//

// Package requestedsession provides the /v1/delegated_checkout/requested_sessions APIs
package requestedsession

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke /v1/delegated_checkout/requested_sessions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a requested session
func New(params *stripe.DelegatedCheckoutRequestedSessionParams) (*stripe.DelegatedCheckoutRequestedSession, error) {
	return getC().New(params)
}

// Creates a requested session
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.DelegatedCheckoutRequestedSessionParams) (*stripe.DelegatedCheckoutRequestedSession, error) {
	requestedsession := &stripe.DelegatedCheckoutRequestedSession{}
	err := c.B.Call(
		http.MethodPost, "/v1/delegated_checkout/requested_sessions", c.Key, params, requestedsession)
	return requestedsession, err
}

// Retrieves a requested session
func Get(id string, params *stripe.DelegatedCheckoutRequestedSessionParams) (*stripe.DelegatedCheckoutRequestedSession, error) {
	return getC().Get(id, params)
}

// Retrieves a requested session
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.DelegatedCheckoutRequestedSessionParams) (*stripe.DelegatedCheckoutRequestedSession, error) {
	path := stripe.FormatURLPath(
		"/v1/delegated_checkout/requested_sessions/%s", id)
	requestedsession := &stripe.DelegatedCheckoutRequestedSession{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, requestedsession)
	return requestedsession, err
}

// Updates a requested session
func Update(id string, params *stripe.DelegatedCheckoutRequestedSessionParams) (*stripe.DelegatedCheckoutRequestedSession, error) {
	return getC().Update(id, params)
}

// Updates a requested session
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.DelegatedCheckoutRequestedSessionParams) (*stripe.DelegatedCheckoutRequestedSession, error) {
	path := stripe.FormatURLPath(
		"/v1/delegated_checkout/requested_sessions/%s", id)
	requestedsession := &stripe.DelegatedCheckoutRequestedSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, requestedsession)
	return requestedsession, err
}

// Confirms a requested session
func Confirm(id string, params *stripe.DelegatedCheckoutRequestedSessionConfirmParams) (*stripe.DelegatedCheckoutRequestedSession, error) {
	return getC().Confirm(id, params)
}

// Confirms a requested session
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Confirm(id string, params *stripe.DelegatedCheckoutRequestedSessionConfirmParams) (*stripe.DelegatedCheckoutRequestedSession, error) {
	path := stripe.FormatURLPath(
		"/v1/delegated_checkout/requested_sessions/%s/confirm", id)
	requestedsession := &stripe.DelegatedCheckoutRequestedSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, requestedsession)
	return requestedsession, err
}

// Expires a requested session
func Expire(id string, params *stripe.DelegatedCheckoutRequestedSessionExpireParams) (*stripe.DelegatedCheckoutRequestedSession, error) {
	return getC().Expire(id, params)
}

// Expires a requested session
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Expire(id string, params *stripe.DelegatedCheckoutRequestedSessionExpireParams) (*stripe.DelegatedCheckoutRequestedSession, error) {
	path := stripe.FormatURLPath(
		"/v1/delegated_checkout/requested_sessions/%s/expire", id)
	requestedsession := &stripe.DelegatedCheckoutRequestedSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, requestedsession)
	return requestedsession, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
