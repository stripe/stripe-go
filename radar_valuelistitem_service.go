//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
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

// Returns a list of ValueListItem objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1RadarValueListItemService) List(ctx context.Context, listParams *RadarValueListItemListParams) Seq2[*RadarValueListItem, error] {
	if listParams == nil {
		listParams = &RadarValueListItemListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*RadarValueListItem, ListContainer, error) {
		list := &RadarValueListItemList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/radar/value_list_items", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
