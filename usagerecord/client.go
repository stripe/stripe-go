// Package usage_record provides the /subscription_items/{SUBSCRIPTION_ITEM_ID}/usage_records APIs
package usagerecord

import (
	"fmt"
	"net/url"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /plans APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new usage record.
// For more details see https://stripe.com/docs/api#usage_records
func New(params *stripe.UsageRecordParams) (*stripe.UsageRecord, error) {
	return getC().New(params)
}

// New internal implementation to create a new usage record.
func (c Client) New(params *stripe.UsageRecordParams) (*stripe.UsageRecord, error) {
	body := &form.Values{}
	form.AppendTo(body, params)

	url := fmt.Sprintf("/subscription_items/%s/usage_records", url.QueryEscape(params.SubscriptionItem))
	record := &stripe.UsageRecord{}
	err := c.B.Call("POST", url, c.Key, body, &params.Params, record)

	return record, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
