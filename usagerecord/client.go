//
//
// File generated from our OpenAPI spec
//
//

// Package usagerecord provides the /subscription_items/{subscription_item}/usage_records APIs
package usagerecord

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /subscription_items/{subscription_item}/usage_records APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new usage record.
func New(params *stripe.UsageRecordParams) (*stripe.UsageRecord, error) {
	return getC().New(params)
}

// New creates a new usage record.
func (c Client) New(params *stripe.UsageRecordParams) (*stripe.UsageRecord, error) {
	path := stripe.FormatURLPath(
		"/v1/subscription_items/%s/usage_records",
		stripe.StringValue(params.SubscriptionItem),
	)
	usagerecord := &stripe.UsageRecord{}
	var err error
	sr := stripe.NewStripeRequest(http.MethodPost, path, c.Key)
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, usagerecord)
	return usagerecord, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
