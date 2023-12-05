//
//
// File generated from our OpenAPI spec
//
//

// Package payout provides the /payouts APIs
package payout

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /payouts APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new payout.
func New(params *stripe.PayoutParams) (*stripe.Payout, error) {
	return getC().New(params)
}

// New creates a new payout.
func (c Client) New(params *stripe.PayoutParams) (*stripe.Payout, error) {
	payout := &stripe.Payout{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: "/v1/payouts", Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, payout)
	return payout, err
}

// Get returns the details of a payout.
func Get(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	return getC().Get(id, params)
}

// Get returns the details of a payout.
func (c Client) Get(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	path := stripe.FormatURLPath("/v1/payouts/%s", id)
	payout := &stripe.Payout{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodGet, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, payout)
	return payout, err
}

// Update updates a payout's properties.
func Update(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	return getC().Update(id, params)
}

// Update updates a payout's properties.
func (c Client) Update(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	path := stripe.FormatURLPath("/v1/payouts/%s", id)
	payout := &stripe.Payout{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, payout)
	return payout, err
}

// Cancel is the method for the `POST /v1/payouts/{payout}/cancel` API.
func Cancel(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	return getC().Cancel(id, params)
}

// Cancel is the method for the `POST /v1/payouts/{payout}/cancel` API.
func (c Client) Cancel(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	path := stripe.FormatURLPath("/v1/payouts/%s/cancel", id)
	payout := &stripe.Payout{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, payout)
	return payout, err
}

// Reverse is the method for the `POST /v1/payouts/{payout}/reverse` API.
func Reverse(id string, params *stripe.PayoutReverseParams) (*stripe.Payout, error) {
	return getC().Reverse(id, params)
}

// Reverse is the method for the `POST /v1/payouts/{payout}/reverse` API.
func (c Client) Reverse(id string, params *stripe.PayoutReverseParams) (*stripe.Payout, error) {
	path := stripe.FormatURLPath("/v1/payouts/%s/reverse", id)
	payout := &stripe.Payout{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, payout)
	return payout, err
}

// List returns a list of payouts.
func List(params *stripe.PayoutListParams) *Iter {
	return getC().List(params)
}

// List returns a list of payouts.
func (c Client) List(listParams *stripe.PayoutListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PayoutList{}
			sr := stripe.StripeRequest{
				Method: http.MethodGet,
				Path:   "/v1/payouts",
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

// Iter is an iterator for payouts.
type Iter struct {
	*stripe.Iter
}

// Payout returns the payout which the iterator is currently pointing to.
func (i *Iter) Payout() *stripe.Payout {
	return i.Current().(*stripe.Payout)
}

// PayoutList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PayoutList() *stripe.PayoutList {
	return i.List().(*stripe.PayoutList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
