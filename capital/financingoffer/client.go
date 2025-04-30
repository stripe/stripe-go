//
//
// File generated from our OpenAPI spec
//
//

// Package financingoffer provides the /v1/capital/financing_offers APIs
package financingoffer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/capital/financing_offers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Get the details of the financing offer
func Get(id string, params *stripe.CapitalFinancingOfferParams) (*stripe.CapitalFinancingOffer, error) {
	return getC().Get(id, params)
}

// Get the details of the financing offer
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.CapitalFinancingOfferParams) (*stripe.CapitalFinancingOffer, error) {
	path := stripe.FormatURLPath("/v1/capital/financing_offers/%s", id)
	financingoffer := &stripe.CapitalFinancingOffer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, financingoffer)
	return financingoffer, err
}

// Acknowledges that platform has received and delivered the financing_offer to
// the intended merchant recipient.
func MarkDelivered(id string, params *stripe.CapitalFinancingOfferMarkDeliveredParams) (*stripe.CapitalFinancingOffer, error) {
	return getC().MarkDelivered(id, params)
}

// Acknowledges that platform has received and delivered the financing_offer to
// the intended merchant recipient.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) MarkDelivered(id string, params *stripe.CapitalFinancingOfferMarkDeliveredParams) (*stripe.CapitalFinancingOffer, error) {
	path := stripe.FormatURLPath(
		"/v1/capital/financing_offers/%s/mark_delivered", id)
	financingoffer := &stripe.CapitalFinancingOffer{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, financingoffer)
	return financingoffer, err
}

// Retrieves the financing offers available for Connected accounts that belong to your platform.
func List(params *stripe.CapitalFinancingOfferListParams) *Iter {
	return getC().List(params)
}

// Retrieves the financing offers available for Connected accounts that belong to your platform.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.CapitalFinancingOfferListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CapitalFinancingOfferList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/capital/financing_offers", c.Key, []byte(b.Encode()), p, list)

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
