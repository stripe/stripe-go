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

// v1RadarValueListService is used to invoke /v1/radar/value_lists APIs.
type v1RadarValueListService struct {
	B   Backend
	Key string
}

// Creates a new ValueList object, which can then be referenced in rules.
func (c v1RadarValueListService) Create(ctx context.Context, params *RadarValueListCreateParams) (*RadarValueList, error) {
	if params == nil {
		params = &RadarValueListCreateParams{}
	}
	params.Context = ctx
	valuelist := &RadarValueList{}
	err := c.B.Call(
		http.MethodPost, "/v1/radar/value_lists", c.Key, params, valuelist)
	return valuelist, err
}

// Retrieves a ValueList object.
func (c v1RadarValueListService) Retrieve(ctx context.Context, id string, params *RadarValueListRetrieveParams) (*RadarValueList, error) {
	if params == nil {
		params = &RadarValueListRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/radar/value_lists/%s", id)
	valuelist := &RadarValueList{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, valuelist)
	return valuelist, err
}

// Updates a ValueList object by setting the values of the parameters passed. Any parameters not provided will be left unchanged. Note that item_type is immutable.
func (c v1RadarValueListService) Update(ctx context.Context, id string, params *RadarValueListUpdateParams) (*RadarValueList, error) {
	if params == nil {
		params = &RadarValueListUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/radar/value_lists/%s", id)
	valuelist := &RadarValueList{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, valuelist)
	return valuelist, err
}

// Deletes a ValueList object, also deleting any items contained within the value list. To be deleted, a value list must not be referenced in any rules.
func (c v1RadarValueListService) Delete(ctx context.Context, id string, params *RadarValueListDeleteParams) (*RadarValueList, error) {
	if params == nil {
		params = &RadarValueListDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/radar/value_lists/%s", id)
	valuelist := &RadarValueList{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, valuelist)
	return valuelist, err
}

// Returns a list of ValueList objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1RadarValueListService) List(ctx context.Context, listParams *RadarValueListListParams) Seq2[*RadarValueList, error] {
	if listParams == nil {
		listParams = &RadarValueListListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*RadarValueList, ListContainer, error) {
		list := &RadarValueListList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/radar/value_lists", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
