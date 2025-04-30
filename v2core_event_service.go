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

// v2CoreEventService is used to invoke event related APIs.
type v2CoreEventService struct {
	B   Backend
	Key string
}

// Retrieves the details of an event.
func (c v2CoreEventService) Retrieve(ctx context.Context, id string, params *V2CoreEventRetrieveParams) (V2Event, error) {
	if params == nil {
		params = &V2CoreEventRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/events/%s", id)
	raw := &V2RawEvent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, raw)
	if err != nil {
		return nil, err
	}
	return ConvertRawEvent(raw, c.B, c.Key)
}

// List events, going back up to 30 days.
func (c v2CoreEventService) List(ctx context.Context, listParams *V2CoreEventListParams) Seq2[V2Event, error] {
	if listParams == nil {
		listParams = &V2CoreEventListParams{}
	}
	listParams.Context = ctx
	return NewV2List("/v2/core/events", listParams, func(path string, p ParamsContainer) (*V2Page[V2Event], error) {
		raw := &V2Page[V2RawEvent]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, raw)
		page := &V2Page[V2Event]{}
		page.LastResponse = raw.LastResponse
		page.NextPageURL = raw.NextPageURL
		page.Data = make([]V2Event, len(raw.Data))
		for i := range raw.Data {
			page.Data[i], err = ConvertRawEvent(&raw.Data[i], c.B, c.Key)
			if err != nil {
				return nil, err
			}
		}
		return page, err
	}).All()
}
