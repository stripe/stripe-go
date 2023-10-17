//
//
// File generated from our OpenAPI spec
//
//

// Package financingoffer provides the /capital/financing_offers APIs
package financingoffer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /capital/financing_offers APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of a capital financing offer.
func Get(id string, params *stripe.CapitalFinancingOfferParams) (*stripe.CapitalFinancingOffer, error) {
	return getC().Get(id, params)
}

// Get returns the details of a capital financing offer.
func (c Client) Get(id string, params *stripe.CapitalFinancingOfferParams) (*stripe.CapitalFinancingOffer, error) {
	path := stripe.FormatURLPath("/v1/capital/financing_offers/%s", id)
	financingoffer := &stripe.CapitalFinancingOffer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, financingoffer)
	return financingoffer, err
}

// MarkDelivered is the method for the `POST /v1/capital/financing_offers/{financing_offer}/mark_delivered` API.
func MarkDelivered(id string, params *stripe.CapitalFinancingOfferMarkDeliveredParams) (*stripe.CapitalFinancingOffer, error) {
	return getC().MarkDelivered(id, params)
}

// MarkDelivered is the method for the `POST /v1/capital/financing_offers/{financing_offer}/mark_delivered` API.
func (c Client) MarkDelivered(id string, params *stripe.CapitalFinancingOfferMarkDeliveredParams) (*stripe.CapitalFinancingOffer, error) {
	path := stripe.FormatURLPath(
		"/v1/capital/financing_offers/%s/mark_delivered",
		id,
	)
	financingoffer := &stripe.CapitalFinancingOffer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, financingoffer)
	return financingoffer, err
}

// List returns a list of capital financing offers.
func List(params *stripe.CapitalFinancingOfferListParams) *Iter {
	return getC().List(params)
}

// List returns a list of capital financing offers.
func (c Client) List(listParams *stripe.CapitalFinancingOfferListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CapitalFinancingOfferList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/capital/financing_offers", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for capital financing offers.
type Iter struct {
	*stripe.Iter
}

// CapitalFinancingOffer returns the capital financing offer which the iterator is currently pointing to.
func (i *Iter) CapitalFinancingOffer() *stripe.CapitalFinancingOffer {
	return i.Current().(*stripe.CapitalFinancingOffer)
}

// CapitalFinancingOfferList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CapitalFinancingOfferList() *stripe.CapitalFinancingOfferList {
	return i.List().(*stripe.CapitalFinancingOfferList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
