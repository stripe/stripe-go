//
//
// File generated from our OpenAPI spec
//
//

// Package valuelist provides the /radar/value_lists APIs
package valuelist

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /radar/value_lists APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new radar value list.
func New(params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	return getC().New(params)
}

// New creates a new radar value list.
func (c Client) New(params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	valuelist := &stripe.RadarValueList{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: "/v1/radar/value_lists", Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, valuelist)
	return valuelist, err
}

// Get returns the details of a radar value list.
func Get(id string, params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	return getC().Get(id, params)
}

// Get returns the details of a radar value list.
func (c Client) Get(id string, params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	path := stripe.FormatURLPath("/v1/radar/value_lists/%s", id)
	valuelist := &stripe.RadarValueList{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodGet, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, valuelist)
	return valuelist, err
}

// Update updates a radar value list's properties.
func Update(id string, params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	return getC().Update(id, params)
}

// Update updates a radar value list's properties.
func (c Client) Update(id string, params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	path := stripe.FormatURLPath("/v1/radar/value_lists/%s", id)
	valuelist := &stripe.RadarValueList{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, valuelist)
	return valuelist, err
}

// Del removes a radar value list.
func Del(id string, params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	return getC().Del(id, params)
}

// Del removes a radar value list.
func (c Client) Del(id string, params *stripe.RadarValueListParams) (*stripe.RadarValueList, error) {
	path := stripe.FormatURLPath("/v1/radar/value_lists/%s", id)
	valuelist := &stripe.RadarValueList{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodDelete, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, valuelist)
	return valuelist, err
}

// List returns a list of radar value lists.
func List(params *stripe.RadarValueListListParams) *Iter {
	return getC().List(params)
}

// List returns a list of radar value lists.
func (c Client) List(listParams *stripe.RadarValueListListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.RadarValueListList{}
			sr := stripe.StripeRequest{
				Method: http.MethodGet,
				Path:   "/v1/radar/value_lists",
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

// Iter is an iterator for radar value lists.
type Iter struct {
	*stripe.Iter
}

// RadarValueList returns the radar value list which the iterator is currently pointing to.
func (i *Iter) RadarValueList() *stripe.RadarValueList {
	return i.Current().(*stripe.RadarValueList)
}

// RadarValueListList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) RadarValueListList() *stripe.RadarValueListList {
	return i.List().(*stripe.RadarValueListList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
