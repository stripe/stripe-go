//
//
// File generated from our OpenAPI spec
//
//

// Package paymentmethodconfiguration provides the /v1/payment_method_configurations APIs
package paymentmethodconfiguration

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/payment_method_configurations APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a payment method configuration
func New(params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	return getC().New(params)
}

// Creates a payment method configuration
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	paymentmethodconfiguration := &stripe.PaymentMethodConfiguration{}
	err := c.B.Call(
		http.MethodPost, "/v1/payment_method_configurations", c.Key, params, paymentmethodconfiguration)
	return paymentmethodconfiguration, err
}

// Retrieve payment method configuration
func Get(id string, params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	return getC().Get(id, params)
}

// Retrieve payment method configuration
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	path := stripe.FormatURLPath("/v1/payment_method_configurations/%s", id)
	paymentmethodconfiguration := &stripe.PaymentMethodConfiguration{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, paymentmethodconfiguration)
	return paymentmethodconfiguration, err
}

// Update payment method configuration
func Update(id string, params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	return getC().Update(id, params)
}

// Update payment method configuration
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	path := stripe.FormatURLPath("/v1/payment_method_configurations/%s", id)
	paymentmethodconfiguration := &stripe.PaymentMethodConfiguration{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, paymentmethodconfiguration)
	return paymentmethodconfiguration, err
}

// List payment method configurations
func List(params *stripe.PaymentMethodConfigurationListParams) *Iter {
	return getC().List(params)
}

// List payment method configurations
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.PaymentMethodConfigurationListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PaymentMethodConfigurationList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/payment_method_configurations", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for payment method configurations.
type Iter struct {
	*stripe.Iter
}

// PaymentMethodConfiguration returns the payment method configuration which the iterator is currently pointing to.
func (i *Iter) PaymentMethodConfiguration() *stripe.PaymentMethodConfiguration {
	return i.Current().(*stripe.PaymentMethodConfiguration)
}

// PaymentMethodConfigurationList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) PaymentMethodConfigurationList() *stripe.PaymentMethodConfigurationList {
	return i.List().(*stripe.PaymentMethodConfigurationList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
