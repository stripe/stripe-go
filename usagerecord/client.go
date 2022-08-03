//
//
// File generated from our OpenAPI spec
//
//

// Package usagerecord provides the /subscription_items/{subscription_item}/usage_records APIs
package usagerecord

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
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
	err := c.B.Call(http.MethodPost, path, c.Key, params, usagerecord)
	return usagerecord, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
