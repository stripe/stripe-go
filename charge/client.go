//
//
// File generated from our OpenAPI spec
//
//

// Package charge provides the /charges APIs
package charge

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /charges APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new charge.
func New(params *stripe.ChargeParams) (*stripe.Charge, error) {
	return getC().New(params)
}

// New creates a new charge.
func (c Client) New(params *stripe.ChargeParams) (*stripe.Charge, error) {
	charge := &stripe.Charge{}
	err := c.B.Call(http.MethodPost, "/v1/charges", c.Key, params, charge)
	return charge, err
}

// Get returns the details of a charge.
func Get(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	return getC().Get(id, params)
}

// Get returns the details of a charge.
func (c Client) Get(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	path := stripe.FormatURLPath("/v1/charges/%s", id)
	charge := &stripe.Charge{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, charge)
	return charge, err
}

// Update updates a charge's properties.
func Update(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	return getC().Update(id, params)
}

// Update updates a charge's properties.
func (c Client) Update(id string, params *stripe.ChargeParams) (*stripe.Charge, error) {
	path := stripe.FormatURLPath("/v1/charges/%s", id)
	charge := &stripe.Charge{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, charge)
	return charge, err
}

// Capture is the method for the `POST /v1/charges/{charge}/capture` API.
func Capture(id string, params *stripe.CaptureParams) (*stripe.Charge, error) {
	return getC().Capture(id, params)
}

// Capture is the method for the `POST /v1/charges/{charge}/capture` API.
func (c Client) Capture(id string, params *stripe.CaptureParams) (*stripe.Charge, error) {
	path := stripe.FormatURLPath("/v1/charges/%s/capture", id)
	charge := &stripe.Charge{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, charge)
	return charge, err
}

// List returns a list of charges.
func List(params *stripe.ChargeListParams) *Iter {
	return getC().List(params)
}

// List returns a list of charges.
func (c Client) List(listParams *stripe.ChargeListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ChargeList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/charges", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for charges.
type Iter struct {
	*stripe.Iter
}

// Charge returns the charge which the iterator is currently pointing to.
func (i *Iter) Charge() *stripe.Charge {
	return i.Current().(*stripe.Charge)
}

// ChargeList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ChargeList() *stripe.ChargeList {
	return i.List().(*stripe.ChargeList)
}

// Search returns a search result containing charges.
func Search(params *stripe.ChargeSearchParams) *SearchIter {
	return getC().Search(params)
}

// Search returns a search result containing charges.
func (c Client) Search(params *stripe.ChargeSearchParams) *SearchIter {
	return &SearchIter{
		SearchIter: stripe.GetSearchIter(params, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.SearchContainer, error) {
			list := &stripe.ChargeSearchResult{}
			err := c.B.CallRaw(http.MethodGet, "/v1/charges/search", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// SearchIter is an iterator for charges.
type SearchIter struct {
	*stripe.SearchIter
}

// Charge returns the charge which the iterator is currently pointing to.
func (i *SearchIter) Charge() *stripe.Charge {
	return i.Current().(*stripe.Charge)
}

// ChargeSearchResult returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *SearchIter) ChargeSearchResult() *stripe.ChargeSearchResult {
	return i.SearchResult().(*stripe.ChargeSearchResult)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
