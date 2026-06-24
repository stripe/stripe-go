//
//
// File generated from our OpenAPI spec
//
//

// Package alert provides the alert related APIs
package alert

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke alert related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a health alert by ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreHealthAlertParams) (*stripe.V2CoreHealthAlert, error) {
	path := stripe.FormatURLPath("/v2/core/health/alerts/%s", id)
	alert := &stripe.V2CoreHealthAlert{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, alert)
	return alert, err
}

// Retrieves a list of health alerts.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2CoreHealthAlertListParams) stripe.Seq2[*stripe.V2CoreHealthAlert, error] {
	if listParams == nil {
		listParams = &stripe.V2CoreHealthAlertListParams{}
	}
	return stripe.NewV2List("/v2/core/health/alerts", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2CoreHealthAlert], error) {
		page := &stripe.V2Page[*stripe.V2CoreHealthAlert]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
