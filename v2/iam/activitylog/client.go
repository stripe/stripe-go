//
//
// File generated from our OpenAPI spec
//
//

// Package activitylog provides the activitylog related APIs
package activitylog

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
)

// Client is used to invoke activitylog related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// List activity logs of an account.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2IamActivityLogListParams) stripe.Seq2[*stripe.V2IamActivityLog, error] {
	if listParams == nil {
		listParams = &stripe.V2IamActivityLogListParams{}
	}
	return stripe.NewV2List("/v2/iam/activity_logs", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2IamActivityLog], error) {
		page := &stripe.V2Page[*stripe.V2IamActivityLog]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
