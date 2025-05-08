//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v2CoreEventDestinationService is used to invoke eventdestination related APIs.
type v2CoreEventDestinationService struct {
	B   Backend
	Key string
}

// Create a new event destination.
func (c v2CoreEventDestinationService) Create(ctx context.Context, params *V2CoreEventDestinationCreateParams) (*V2EventDestination, error) {
	if params == nil {
		params = &V2CoreEventDestinationCreateParams{}
	}
	params.Context = ctx
	eventdestination := &V2EventDestination{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/event_destinations", c.Key, params, eventdestination)
	return eventdestination, err
}

// Retrieves the details of an event destination.
func (c v2CoreEventDestinationService) Retrieve(ctx context.Context, id string, params *V2CoreEventDestinationRetrieveParams) (*V2EventDestination, error) {
	if params == nil {
		params = &V2CoreEventDestinationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/event_destinations/%s", id)
	eventdestination := &V2EventDestination{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, eventdestination)
	return eventdestination, err
}

// Update the details of an event destination.
func (c v2CoreEventDestinationService) Update(ctx context.Context, id string, params *V2CoreEventDestinationUpdateParams) (*V2EventDestination, error) {
	if params == nil {
		params = &V2CoreEventDestinationUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/event_destinations/%s", id)
	eventdestination := &V2EventDestination{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, eventdestination)
	return eventdestination, err
}

// Delete an event destination.
func (c v2CoreEventDestinationService) Delete(ctx context.Context, id string, params *V2CoreEventDestinationDeleteParams) (*V2EventDestination, error) {
	if params == nil {
		params = &V2CoreEventDestinationDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/event_destinations/%s", id)
	eventdestination := &V2EventDestination{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, eventdestination)
	return eventdestination, err
}

// Disable an event destination.
func (c v2CoreEventDestinationService) Disable(ctx context.Context, id string, params *V2CoreEventDestinationDisableParams) (*V2EventDestination, error) {
	if params == nil {
		params = &V2CoreEventDestinationDisableParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/event_destinations/%s/disable", id)
	eventdestination := &V2EventDestination{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, eventdestination)
	return eventdestination, err
}

// Enable an event destination.
func (c v2CoreEventDestinationService) Enable(ctx context.Context, id string, params *V2CoreEventDestinationEnableParams) (*V2EventDestination, error) {
	if params == nil {
		params = &V2CoreEventDestinationEnableParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/event_destinations/%s/enable", id)
	eventdestination := &V2EventDestination{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, eventdestination)
	return eventdestination, err
}

// Send a `ping` event to an event destination.
func (c v2CoreEventDestinationService) Ping(id string, params *V2CoreEventDestinationPingParams) (V2Event, error) {
	path := FormatURLPath("/v2/core/event_destinations/%s/ping", id)
	raw := &V2RawEvent{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, raw)
	if err != nil {
		return nil, err
	}
	return ConvertRawEvent(raw, c.B, c.Key)
}

// Lists all event destinations.
func (c v2CoreEventDestinationService) List(ctx context.Context, listParams *V2CoreEventDestinationListParams) Seq2[*V2EventDestination, error] {
	if listParams == nil {
		listParams = &V2CoreEventDestinationListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/core/event_destinations", listParams, func(path string, p ParamsContainer) (*V2Page[*V2EventDestination], error) {
		page := &V2Page[*V2EventDestination]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	}).All()
}
