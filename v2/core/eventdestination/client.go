//
//
// File generated from our OpenAPI spec
//
//

// Package eventdestination provides the eventdestination related APIs
package eventdestination

import (
	"net/http"

	stripe "github.com/max-cape/stripe-go-test"
)

// Client is used to invoke eventdestination related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a new event destination.
func (c Client) New(params *stripe.V2CoreEventDestinationParams) (*stripe.V2EventDestination, error) {
	eventdestination := &stripe.V2EventDestination{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/event_destinations", c.Key, params, eventdestination)
	return eventdestination, err
}

// Retrieves the details of an event destination.
func (c Client) Get(id string, params *stripe.V2CoreEventDestinationParams) (*stripe.V2EventDestination, error) {
	path := stripe.FormatURLPath("/v2/core/event_destinations/%s", id)
	eventdestination := &stripe.V2EventDestination{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, eventdestination)
	return eventdestination, err
}

// Update the details of an event destination.
func (c Client) Update(id string, params *stripe.V2CoreEventDestinationParams) (*stripe.V2EventDestination, error) {
	path := stripe.FormatURLPath("/v2/core/event_destinations/%s", id)
	eventdestination := &stripe.V2EventDestination{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, eventdestination)
	return eventdestination, err
}

// Delete an event destination.
func (c Client) Del(id string, params *stripe.V2CoreEventDestinationParams) (*stripe.V2EventDestination, error) {
	path := stripe.FormatURLPath("/v2/core/event_destinations/%s", id)
	eventdestination := &stripe.V2EventDestination{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, eventdestination)
	return eventdestination, err
}

// Disable an event destination.
func (c Client) Disable(id string, params *stripe.V2CoreEventDestinationDisableParams) (*stripe.V2EventDestination, error) {
	path := stripe.FormatURLPath("/v2/core/event_destinations/%s/disable", id)
	eventdestination := &stripe.V2EventDestination{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, eventdestination)
	return eventdestination, err
}

// Enable an event destination.
func (c Client) Enable(id string, params *stripe.V2CoreEventDestinationEnableParams) (*stripe.V2EventDestination, error) {
	path := stripe.FormatURLPath("/v2/core/event_destinations/%s/enable", id)
	eventdestination := &stripe.V2EventDestination{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, eventdestination)
	return eventdestination, err
}

// Send a `ping` event to an event destination.
func (c Client) Ping(id string, params *stripe.V2CoreEventDestinationPingParams) (stripe.V2Event, error) {
	path := stripe.FormatURLPath("/v2/core/event_destinations/%s/ping", id)
	raw := &stripe.V2RawEvent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, raw)
	if err != nil {
		return nil, err
	}
	return stripe.ConvertRawEvent(raw, c.B, c.Key)
}

// Lists all event destinations.
func (c Client) All(listParams *stripe.V2CoreEventDestinationListParams) stripe.Seq2[*stripe.V2EventDestination, error] {
	return stripe.NewV2List("/v2/core/event_destinations", listParams, func(path string, p stripe.ParamsContainer) (*stripe.V2Page[*stripe.V2EventDestination], error) {
		page := &stripe.V2Page[*stripe.V2EventDestination]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
