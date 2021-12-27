//
//
// File generated from our OpenAPI spec
//
//

// Package paymentsource provides the /customers/{customer}/sources APIs
package paymentsource

import (
	"fmt"
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /customers/{customer}/sources APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new payment source.
func New(params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	return getC().New(params)
}

// New creates a new payment source.
func (c Client) New(params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	if params == nil {
		return nil, fmt.Errorf("params should not be nil")
	}
	if params.Customer == nil {
		return nil, fmt.Errorf("Invalid source params: customer needs to be set")
	}
	path := stripe.FormatURLPath(
		"/v1/customers/%s/sources",
		stripe.StringValue(params.Customer),
	)
	paymentsource := &stripe.PaymentSource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentsource)
	return paymentsource, err
}

// Get returns the details of a payment source.
func Get(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	return getC().Get(id, params)
}

// Get returns the details of a payment source.
func (c Client) Get(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	if params == nil {
		return nil, fmt.Errorf("params should not be nil")
	}
	if params.Customer == nil {
		return nil, fmt.Errorf("Invalid source params: customer needs to be set")
	}
	path := stripe.FormatURLPath(
		"/v1/customers/%s/sources/%s",
		stripe.StringValue(params.Customer),
		id,
	)
	paymentsource := &stripe.PaymentSource{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentsource)
	return paymentsource, err
}

// Update updates a payment source's properties.
func Update(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	return getC().Update(id, params)
}

// Update updates a payment source's properties.
func (c Client) Update(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	if params == nil {
		return nil, fmt.Errorf("params should not be nil")
	}
	if params.Customer == nil {
		return nil, fmt.Errorf("Invalid source params: customer needs to be set")
	}
	path := stripe.FormatURLPath(
		"/v1/customers/%s/sources/%s",
		stripe.StringValue(params.Customer),
		id,
	)
	paymentsource := &stripe.PaymentSource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, paymentsource)
	return paymentsource, err
}

// Del removes a payment source.
func Del(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	return getC().Del(id, params)
}

// Del removes a payment source.
func (c Client) Del(id string, params *stripe.CustomerSourceParams) (*stripe.PaymentSource, error) {
	if params == nil {
		return nil, fmt.Errorf("params should not be nil")
	}
	if params.Customer == nil {
		return nil, fmt.Errorf("Invalid source params: customer needs to be set")
	}
	path := stripe.FormatURLPath(
		"/v1/customers/%s/sources/%s",
		stripe.StringValue(params.Customer),
		id,
	)
	paymentsource := &stripe.PaymentSource{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, paymentsource)
	return paymentsource, err
}

// Verify verifies a source which is used for bank accounts.
func Verify(id string, params *stripe.SourceVerifyParams) (*stripe.PaymentSource, error) {
	return getC().Verify(id, params)
}

// Verify verifies a source which is used for bank accounts.
func (c Client) Verify(id string, params *stripe.SourceVerifyParams) (*stripe.PaymentSource, error) {
	if params == nil {
		return nil, fmt.Errorf("params should not be nil")
	}

	var path string
	if params.Customer != nil {
		path = stripe.FormatURLPath("/v1/customers/%s/sources/%s/verify",
			stripe.StringValue(params.Customer), id)
	} else if len(params.Values) > 0 {
		path = stripe.FormatURLPath("/v1/sources/%s/verify", id)
	} else {
		return nil, fmt.Errorf("Only customer bank accounts or sources can be verified in this manner")
	}

	source := &stripe.PaymentSource{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, source)
	return source, err
}

// List returns a list of payment sources.
func List(params *stripe.SourceListParams) *Iter {
	return getC().List(params)
}

// List returns a list of payment sources.
func (c Client) List(listParams *stripe.SourceListParams) *Iter {
	var outerErr error
	var path string

	if listParams == nil {
		outerErr = fmt.Errorf("params should not be nil")
	} else if listParams.Customer == nil {
		outerErr = fmt.Errorf("Invalid source params: customer needs to be set")
	} else {
		path = stripe.FormatURLPath("/v1/customers/%s/sources",
			stripe.StringValue(listParams.Customer))
	}
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.SourceList{}

			if outerErr != nil {
				return nil, list, outerErr
			}

			err := c.B.CallRaw(http.MethodGet, path, c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for payment sources.
type Iter struct {
	*stripe.Iter
}

// PaymentSource returns the payment source which the iterator is currently pointing to.
func (i *Iter) PaymentSource() *stripe.PaymentSource {
	return i.Current().(*stripe.PaymentSource)
}

// SourceList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) SourceList() *stripe.SourceList {
	return i.List().(*stripe.SourceList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
