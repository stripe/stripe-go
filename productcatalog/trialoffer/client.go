//
//
// File generated from our OpenAPI spec
//
//

// Package trialoffer provides the /v1/product_catalog/trial_offers APIs
package trialoffer

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/stripe-go/v86/form"
)

// Client is used to invoke /v1/product_catalog/trial_offers APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a trial offer.
func New(params *stripe.ProductCatalogTrialOfferParams) (*stripe.ProductCatalogTrialOffer, error) {
	return getC().New(params)
}

// Creates a trial offer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.ProductCatalogTrialOfferParams) (*stripe.ProductCatalogTrialOffer, error) {
	trialoffer := &stripe.ProductCatalogTrialOffer{}
	err := c.B.Call(
		http.MethodPost, "/v1/product_catalog/trial_offers", c.Key, params, trialoffer)
	return trialoffer, err
}

// Retrieves the trial offer with the given ID.
func Get(id string, params *stripe.ProductCatalogTrialOfferParams) (*stripe.ProductCatalogTrialOffer, error) {
	return getC().Get(id, params)
}

// Retrieves the trial offer with the given ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.ProductCatalogTrialOfferParams) (*stripe.ProductCatalogTrialOffer, error) {
	path := stripe.FormatURLPath("/v1/product_catalog/trial_offers/%s", id)
	trialoffer := &stripe.ProductCatalogTrialOffer{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, trialoffer)
	return trialoffer, err
}

// Returns a list of trial offers.
func List(params *stripe.ProductCatalogTrialOfferListParams) *Iter {
	return getC().List(params)
}

// Returns a list of trial offers.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.ProductCatalogTrialOfferListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.ProductCatalogTrialOfferList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/product_catalog/trial_offers", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for product catalog trial offers.
type Iter struct {
	*stripe.Iter
}

// ProductCatalogTrialOffer returns the product catalog trial offer which the iterator is currently pointing to.
func (i *Iter) ProductCatalogTrialOffer() *stripe.ProductCatalogTrialOffer {
	return i.Current().(*stripe.ProductCatalogTrialOffer)
}

// ProductCatalogTrialOfferList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) ProductCatalogTrialOfferList() *stripe.ProductCatalogTrialOfferList {
	return i.List().(*stripe.ProductCatalogTrialOfferList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
