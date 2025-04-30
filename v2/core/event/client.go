//
//
// File generated from our OpenAPI spec
//
//

// Package event provides the event related APIs
package event

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
)

// Client is used to invoke event related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of an event.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreEventParams) (stripe.V2Event, error) {
	path := stripe.FormatURLPath("/v2/core/events/%s", id)
	raw := &stripe.V2RawEvent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, raw)
	if err != nil {
		return nil, err
	}
	return stripe.ConvertRawEvent(raw, c.B, c.Key)
}

// List events, going back up to 30 days.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2CoreEventListParams) stripe.Seq2[stripe.V2Event, error] {
	return stripe.NewV2List("/v2/core/events", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[stripe.V2Event], error) {
		raw := &stripe.V2Page[stripe.V2RawEvent]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, raw)
		page := &stripe.V2Page[stripe.V2Event]{}
		page.LastResponse = raw.LastResponse
		page.NextPageURL = raw.NextPageURL
		page.Data = make([]stripe.V2Event, len(raw.Data))
		for i := range raw.Data {
			page.Data[i], err = stripe.ConvertRawEvent(&raw.Data[i], c.B, c.Key)
			if err != nil {
				return nil, err
			}
		}
		return page, err
	}).All()
}
