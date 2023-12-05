//
//
// File generated from our OpenAPI spec
//
//

// Package valuelistitem provides the /radar/value_list_items APIs
// For more details, see: https://stripe.com/docs/api/radar/list_items?lang=go
package valuelistitem

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /radar/value_list_items APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new radar value list item.
func New(params *stripe.RadarValueListItemParams) (*stripe.RadarValueListItem, error) {
	return getC().New(params)
}

// New creates a new radar value list item.
func (c Client) New(params *stripe.RadarValueListItemParams) (*stripe.RadarValueListItem, error) {
	valuelistitem := &stripe.RadarValueListItem{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: "/v1/radar/value_list_items", Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, valuelistitem)
	return valuelistitem, err
}

// Get returns the details of a radar value list item.
func Get(id string, params *stripe.RadarValueListItemParams) (*stripe.RadarValueListItem, error) {
	return getC().Get(id, params)
}

// Get returns the details of a radar value list item.
func (c Client) Get(id string, params *stripe.RadarValueListItemParams) (*stripe.RadarValueListItem, error) {
	path := stripe.FormatURLPath("/v1/radar/value_list_items/%s", id)
	valuelistitem := &stripe.RadarValueListItem{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodGet, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, valuelistitem)
	return valuelistitem, err
}

// Del removes a radar value list item.
func Del(id string, params *stripe.RadarValueListItemParams) (*stripe.RadarValueListItem, error) {
	return getC().Del(id, params)
}

// Del removes a radar value list item.
func (c Client) Del(id string, params *stripe.RadarValueListItemParams) (*stripe.RadarValueListItem, error) {
	path := stripe.FormatURLPath("/v1/radar/value_list_items/%s", id)
	valuelistitem := &stripe.RadarValueListItem{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodDelete, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, valuelistitem)
	return valuelistitem, err
}

// List returns a list of radar value list items.
func List(params *stripe.RadarValueListItemListParams) *Iter {
	return getC().List(params)
}

// List returns a list of radar value list items.
func (c Client) List(listParams *stripe.RadarValueListItemListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.RadarValueListItemList{}
			sr := stripe.StripeRequest{
				Method: http.MethodGet,
				Path:   "/v1/radar/value_list_items",
				Key:    c.Key,
			}
			err := sr.SetRawForm(p, b)
			if err != nil {
				return nil, list, err
			}
			err = c.B.Call(sr, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for radar value list items.
type Iter struct {
	*stripe.Iter
}

// RadarValueListItem returns the radar value list item which the iterator is currently pointing to.
func (i *Iter) RadarValueListItem() *stripe.RadarValueListItem {
	return i.Current().(*stripe.RadarValueListItem)
}

// RadarValueListItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) RadarValueListItemList() *stripe.RadarValueListItemList {
	return i.List().(*stripe.RadarValueListItemList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
