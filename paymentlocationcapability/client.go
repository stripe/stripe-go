//
//
// File generated from our OpenAPI spec
//
//

// Package paymentlocationcapability provides the /v1/payment_location_capabilities APIs
package paymentlocationcapability

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v85"
	"github.com/stripe/stripe-go/v85/form"
)

// Client is used to invoke /v1/payment_location_capabilities APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves a payment_location capability
func Get(id string, params *stripe.PaymentLocationCapabilityParams) (*stripe.PaymentLocationCapability, error) {
	return getC().Get(id, params)
}

// Retrieves a payment_location capability
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.PaymentLocationCapabilityParams) (*stripe.PaymentLocationCapability, error) {
	path := stripe.FormatURLPath("/v1/payment_location_capabilities/%s", id)
	paymentlocationcapability := &stripe.PaymentLocationCapability{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, paymentlocationcapability)
	return paymentlocationcapability, err
}

// Updates a payment_location capability. Request or remove a payment_location capability by updating its requested parameter.
func Update(id string, params *stripe.PaymentLocationCapabilityParams) (*stripe.PaymentLocationCapability, error) {
	return getC().Update(id, params)
}

// Updates a payment_location capability. Request or remove a payment_location capability by updating its requested parameter.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.PaymentLocationCapabilityParams) (*stripe.PaymentLocationCapability, error) {
	path := stripe.FormatURLPath("/v1/payment_location_capabilities/%s", id)
	paymentlocationcapability := &stripe.PaymentLocationCapability{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, paymentlocationcapability)
	return paymentlocationcapability, err
}

// List all payment location capabilities associated with the payment location.
func List(params *stripe.PaymentLocationCapabilityListParams) *Iter {
	return getC().List(params)
}

// List all payment location capabilities associated with the payment location.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.PaymentLocationCapabilityListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PaymentLocationCapabilityList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/payment_location_capabilities", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for payment location capabilities.
type Iter struct {
	*stripe.Iter
}

// PaymentLocationCapability returns the payment location capability which the iterator is currently pointing to.
func (i *Iter) PaymentLocationCapability() *stripe.PaymentLocationCapability {
	return i.Current().(*stripe.PaymentLocationCapability)
}

// PaymentLocationCapabilityList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentLocationCapabilityList() *stripe.PaymentLocationCapabilityList {
	return i.List().(*stripe.PaymentLocationCapabilityList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
