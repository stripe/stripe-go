//
//
// File generated from our OpenAPI spec
//
//

// Package history provides the history related APIs
package history

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke history related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a list of alert history entries for a health alert.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2CoreHealthAlertsHistoryListParams) stripe.Seq2[*stripe.V2CoreHealthAlertHistoryEntry, error] {
	if listParams == nil {
		listParams = &stripe.V2CoreHealthAlertsHistoryListParams{}
	}
	path := stripe.FormatURLPath(
		"/v2/core/health/alerts/%s/history", stripe.StringValue(listParams.ID))
	return stripe.NewV2List(path, listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2CoreHealthAlertHistoryEntry], error) {
		page := &stripe.V2Page[*stripe.V2CoreHealthAlertHistoryEntry]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
