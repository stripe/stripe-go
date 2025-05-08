//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1PaymentMethodConfigurationService is used to invoke /v1/payment_method_configurations APIs.
type v1PaymentMethodConfigurationService struct {
	B   Backend
	Key string
}

// Creates a payment method configuration
func (c v1PaymentMethodConfigurationService) Create(ctx context.Context, params *PaymentMethodConfigurationCreateParams) (*PaymentMethodConfiguration, error) {
	if params == nil {
		params = &PaymentMethodConfigurationCreateParams{}
	}
	params.Context = ctx
	paymentmethodconfiguration := &PaymentMethodConfiguration{}
	err := c.B.Call(
		http.MethodPost, "/v1/payment_method_configurations", c.Key, params, paymentmethodconfiguration)
	return paymentmethodconfiguration, err
}

// Retrieve payment method configuration
func (c v1PaymentMethodConfigurationService) Retrieve(ctx context.Context, id string, params *PaymentMethodConfigurationRetrieveParams) (*PaymentMethodConfiguration, error) {
	if params == nil {
		params = &PaymentMethodConfigurationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_method_configurations/%s", id)
	paymentmethodconfiguration := &PaymentMethodConfiguration{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, paymentmethodconfiguration)
	return paymentmethodconfiguration, err
}

// Update payment method configuration
func (c v1PaymentMethodConfigurationService) Update(ctx context.Context, id string, params *PaymentMethodConfigurationUpdateParams) (*PaymentMethodConfiguration, error) {
	if params == nil {
		params = &PaymentMethodConfigurationUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/payment_method_configurations/%s", id)
	paymentmethodconfiguration := &PaymentMethodConfiguration{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, paymentmethodconfiguration)
	return paymentmethodconfiguration, err
}

// List payment method configurations
func (c v1PaymentMethodConfigurationService) List(ctx context.Context, listParams *PaymentMethodConfigurationListParams) Seq2[*PaymentMethodConfiguration, error] {
	if listParams == nil {
		listParams = &PaymentMethodConfigurationListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*PaymentMethodConfiguration, ListContainer, error) {
		list := &PaymentMethodConfigurationList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/payment_method_configurations", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
