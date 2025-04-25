//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"net/http"
)

// v2CoreEventService is used to invoke event related APIs.
type v2CoreEventService struct {
	B   Backend
	Key string
}

// Retrieves the details of an event.
func (c v2CoreEventService) Retrieve(id string, params *V2CoreEventParams) (V2Event, error) {
	path := FormatURLPath("/v2/core/events/%s", id)
	raw := &V2RawEvent{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, raw)
	if err != nil {
		return nil, err
	}
	return ConvertRawEvent(raw, c.B, c.Key)
}

// List events, going back up to 30 days.
func (c v2CoreEventService) All(listParams *V2CoreEventListParams) Seq2[V2Event, error] {
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
