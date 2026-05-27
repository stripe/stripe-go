//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/stripe/stripe-go/v85/form"
)

// v1RadarValueListItemService is used to invoke /v1/radar/value_list_items APIs.
type v1RadarValueListItemService struct {
	B   Backend
	Key string
}

// Creates a new ValueListItem object, which is added to the specified parent value list.
func (c v1RadarValueListItemService) Create(ctx context.Context, params *RadarValueListItemCreateParams) (*RadarValueListItem, error) {
	if params == nil {
		params = &RadarValueListItemCreateParams{}
	}
	params.Context = ctx
	valuelistitem := &RadarValueListItem{}
	err := c.B.Call(
		http.MethodPost, "/v1/radar/value_list_items", c.Key, params, valuelistitem)
	return valuelistitem, err
}

// Retrieves a ValueListItem object.
func (c v1RadarValueListItemService) Retrieve(ctx context.Context, id string, params *RadarValueListItemRetrieveParams) (*RadarValueListItem, error) {
	if params == nil {
		params = &RadarValueListItemRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/radar/value_list_items/%s", id)
	valuelistitem := &RadarValueListItem{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, valuelistitem)
	return valuelistitem, err
}

// Deletes a ValueListItem object, removing it from its parent value list.
func (c v1RadarValueListItemService) Delete(ctx context.Context, id string, params *RadarValueListItemDeleteParams) (*RadarValueListItem, error) {
	if params == nil {
		params = &RadarValueListItemDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/radar/value_list_items/%s", id)
	valuelistitem := &RadarValueListItem{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, valuelistitem)
	return valuelistitem, err
}

// Serializes a ValueListItem create request into a batch job JSONL line.
func (c v1RadarValueListItemService) MarshalBatchCreate(params *RadarValueListItemCreateParams) (string, error) {
	itemID, err := newUUID4()
	if err != nil {
		return "", err
	}

	item := struct {
		ID            string            `json:"id"`
		Context       string            `json:"context,omitempty"`
		StripeVersion string            `json:"stripe_version,omitempty"`
		PathParams    map[string]string `json:"path_params,omitempty"`
		Params        interface{}       `json:"params"`
	}{
		ID:            itemID,
		PathParams:    nil,
		StripeVersion: APIVersion,
	}
	if params != nil {
		item.Params = params
		if params.StripeContext != nil {
			item.Context = *params.StripeContext
		}
	}
	b, err := json.Marshal(item)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

// Returns a list of ValueListItem objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1RadarValueListItemService) List(ctx context.Context, listParams *RadarValueListItemListParams) *V1List[*RadarValueListItem] {
	if listParams == nil {
		listParams = &RadarValueListItemListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*RadarValueListItem], error) {
		list := &v1Page[*RadarValueListItem]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/radar/value_list_items", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
