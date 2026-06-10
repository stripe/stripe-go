//
//
// File generated from our OpenAPI spec
//
//

// Package paymentlocation provides the /v1/payment_locations APIs
package paymentlocation

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/stripe-go/v86/form"
)

// Client is used to invoke /v1/payment_locations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a Payment Location.
func New(params *stripe.PaymentLocationParams) (*stripe.PaymentLocation, error) {
	return getC().New(params)
}

// Create a Payment Location.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.PaymentLocationParams) (*stripe.PaymentLocation, error) {
	paymentlocation := &stripe.PaymentLocation{}
	err := c.B.Call(
		http.MethodPost, "/v1/payment_locations", c.Key, params, paymentlocation)
	return paymentlocation, err
}

// Retrieve a Payment Location.
func Get(id string, params *stripe.PaymentLocationParams) (*stripe.PaymentLocation, error) {
	return getC().Get(id, params)
}

// Retrieve a Payment Location.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.PaymentLocationParams) (*stripe.PaymentLocation, error) {
	path := stripe.FormatURLPath("/v1/payment_locations/%s", id)
	paymentlocation := &stripe.PaymentLocation{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentlocation)
	return paymentlocation, err
}

// Update a Payment Location.
func Update(id string, params *stripe.PaymentLocationParams) (*stripe.PaymentLocation, error) {
	return getC().Update(id, params)
}

// Update a Payment Location.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.PaymentLocationParams) (*stripe.PaymentLocation, error) {
	path := stripe.FormatURLPath("/v1/payment_locations/%s", id)
	paymentlocation := &stripe.PaymentLocation{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentlocation)
	return paymentlocation, err
}

// Delete a Payment Location.
func Del(id string, params *stripe.PaymentLocationParams) (*stripe.PaymentLocation, error) {
	return getC().Del(id, params)
}

// Delete a Payment Location.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.PaymentLocationParams) (*stripe.PaymentLocation, error) {
	path := stripe.FormatURLPath("/v1/payment_locations/%s", id)
	paymentlocation := &stripe.PaymentLocation{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, paymentlocation)
	return paymentlocation, err
}

// List all Payment Locations.
func List(params *stripe.PaymentLocationListParams) *Iter {
	return getC().List(params)
}

// List all Payment Locations.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.PaymentLocationListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PaymentLocationList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/payment_locations", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for payment locations.
type Iter struct {
	*stripe.Iter
}

// PaymentLocation returns the payment location which the iterator is currently pointing to.
func (i *Iter) PaymentLocation() *stripe.PaymentLocation {
	return i.Current().(*stripe.PaymentLocation)
}

// PaymentLocationList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentLocationList() *stripe.PaymentLocationList {
	return i.List().(*stripe.PaymentLocationList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
