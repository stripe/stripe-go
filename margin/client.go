//
//
// File generated from our OpenAPI spec
//
//

// Package margin provides the /billing/margins APIs
package margin

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /billing/margins APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new margin.
func New(params *stripe.MarginParams) (*stripe.Margin, error) {
	return getC().New(params)
}

// New creates a new margin.
func (c Client) New(params *stripe.MarginParams) (*stripe.Margin, error) {
	margin := &stripe.Margin{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: "/v1/billing/margins", Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, margin)
	return margin, err
}

// Get returns the details of a margin.
func Get(id string, params *stripe.MarginParams) (*stripe.Margin, error) {
	return getC().Get(id, params)
}

// Get returns the details of a margin.
func (c Client) Get(id string, params *stripe.MarginParams) (*stripe.Margin, error) {
	path := stripe.FormatURLPath("/v1/billing/margins/%s", id)
	margin := &stripe.Margin{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodGet, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, margin)
	return margin, err
}

// Update updates a margin's properties.
func Update(id string, params *stripe.MarginParams) (*stripe.Margin, error) {
	return getC().Update(id, params)
}

// Update updates a margin's properties.
func (c Client) Update(id string, params *stripe.MarginParams) (*stripe.Margin, error) {
	path := stripe.FormatURLPath("/v1/billing/margins/%s", id)
	margin := &stripe.Margin{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, margin)
	return margin, err
}

// List returns a list of margins.
func List(params *stripe.MarginListParams) *Iter {
	return getC().List(params)
}

// List returns a list of margins.
func (c Client) List(listParams *stripe.MarginListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.MarginList{}
			err := c.B.Call(stripe.StripeRequest{
				Method: http.MethodGet,
				Path:   "/v1/billing/margins",
				Key:    c.Key,
				Params: p,
				Body:   b,
			},
				list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for margins.
type Iter struct {
	*stripe.Iter
}

// Margin returns the margin which the iterator is currently pointing to.
func (i *Iter) Margin() *stripe.Margin {
	return i.Current().(*stripe.Margin)
}

// MarginList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) MarginList() *stripe.MarginList {
	return i.List().(*stripe.MarginList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
