//
//
// File generated from our OpenAPI spec
//
//

// Package paymentmethodconfiguration provides the /payment_method_configurations APIs
package paymentmethodconfiguration

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /payment_method_configurations APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new payment method configuration.
func New(params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	return getC().New(params)
}

// New creates a new payment method configuration.
func (c Client) New(params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	paymentmethodconfiguration := &stripe.PaymentMethodConfiguration{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: "/v1/payment_method_configurations", Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, paymentmethodconfiguration)
	return paymentmethodconfiguration, err
}

// Get returns the details of a payment method configuration.
func Get(id string, params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	return getC().Get(id, params)
}

// Get returns the details of a payment method configuration.
func (c Client) Get(id string, params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	path := stripe.FormatURLPath("/v1/payment_method_configurations/%s", id)
	paymentmethodconfiguration := &stripe.PaymentMethodConfiguration{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodGet, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, paymentmethodconfiguration)
	return paymentmethodconfiguration, err
}

// Update updates a payment method configuration's properties.
func Update(id string, params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	return getC().Update(id, params)
}

// Update updates a payment method configuration's properties.
func (c Client) Update(id string, params *stripe.PaymentMethodConfigurationParams) (*stripe.PaymentMethodConfiguration, error) {
	path := stripe.FormatURLPath("/v1/payment_method_configurations/%s", id)
	paymentmethodconfiguration := &stripe.PaymentMethodConfiguration{}
	var err error
	sr := stripe.StripeRequest{Method: http.MethodPost, Path: path, Key: c.Key}
	err = sr.SetParams(params)
	if err != nil {
		return nil, err
	}
	err = c.B.Call(sr, paymentmethodconfiguration)
	return paymentmethodconfiguration, err
}

// List returns a list of payment method configurations.
func List(params *stripe.PaymentMethodConfigurationListParams) *Iter {
	return getC().List(params)
}

// List returns a list of payment method configurations.
func (c Client) List(listParams *stripe.PaymentMethodConfigurationListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.PaymentMethodConfigurationList{}
			sr := stripe.StripeRequest{
				Method: http.MethodGet,
				Path:   "/v1/payment_method_configurations",
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
