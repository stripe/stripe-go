//
//
// File generated from our OpenAPI spec
//
//

// Package eventdestination provides the eventdestination related APIs
package eventdestination

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke eventdestination related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a new event destination.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreEventDestinationParams) (*stripe.V2CoreEventDestination, error) {
	eventdestination := &stripe.V2CoreEventDestination{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/event_destinations", c.Key, params, eventdestination)
	return eventdestination, err
}

// Retrieves the details of an event destination.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2CoreEventDestinationParams) (*stripe.V2CoreEventDestination, error) {
	path := stripe.FormatURLPath("/v2/core/event_destinations/%s", id)
	eventdestination := &stripe.V2CoreEventDestination{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, eventdestination)
	return eventdestination, err
}

// Update the details of an event destination.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.V2CoreEventDestinationParams) (*stripe.V2CoreEventDestination, error) {
	path := stripe.FormatURLPath("/v2/core/event_destinations/%s", id)
	eventdestination := &stripe.V2CoreEventDestination{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, eventdestination)
	return eventdestination, err
}

// Delete an event destination.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.V2CoreEventDestinationParams) (*stripe.V2DeletedObject, error) {
	path := stripe.FormatURLPath("/v2/core/event_destinations/%s", id)
	deletedObj := &stripe.V2DeletedObject{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, deletedObj)
	return deletedObj, err
}

// Disable an event destination.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Disable(id string, params *stripe.V2CoreEventDestinationDisableParams) (*stripe.V2CoreEventDestination, error) {
	path := stripe.FormatURLPath("/v2/core/event_destinations/%s/disable", id)
	eventdestination := &stripe.V2CoreEventDestination{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, eventdestination)
	return eventdestination, err
}

// Enable an event destination.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Enable(id string, params *stripe.V2CoreEventDestinationEnableParams) (*stripe.V2CoreEventDestination, error) {
	path := stripe.FormatURLPath("/v2/core/event_destinations/%s/enable", id)
	eventdestination := &stripe.V2CoreEventDestination{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, eventdestination)
	return eventdestination, err
}

// Send a `ping` event to an event destination.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Ping(id string, params *stripe.V2CoreEventDestinationPingParams) (stripe.V2CoreEvent, error) {
	path := stripe.FormatURLPath("/v2/core/event_destinations/%s/ping", id)
	raw := &stripe.V2CoreRawEvent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, raw)
	if err != nil {
		return nil, err
	}
	return stripe.ConvertRawEvent(raw, c.B, c.Key)
}

// Lists all event destinations.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) All(listParams *stripe.V2CoreEventDestinationListParams) stripe.Seq2[*stripe.V2CoreEventDestination, error] {
	return stripe.NewV2List("/v2/core/event_destinations", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2CoreEventDestination], error) {
		page := &stripe.V2Page[*stripe.V2CoreEventDestination]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All(listParams.Context)
}
