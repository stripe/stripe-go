//
//
// File generated from our OpenAPI spec
//
//

// Package customertaxexemption provides the /v1/customers/{customer}/tax_exemptions APIs
package customertaxexemption

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
	"github.com/stripe/stripe-go/v86/form"
)

// Client is used to invoke /v1/customers/{customer}/tax_exemptions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a location specific tax exemption for a customer.
func New(params *stripe.CustomerTaxExemptionParams) (*stripe.CustomerTaxExemption, error) {
	return getC().New(params)
}

// Create a location specific tax exemption for a customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.CustomerTaxExemptionParams) (*stripe.CustomerTaxExemption, error) {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/tax_exemptions", stripe.StringValue(params.Customer))
	customertaxexemption := &stripe.CustomerTaxExemption{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, customertaxexemption)
	return customertaxexemption, err
}

// Retrieve a location specific tax exemption for a customer.
func Get(id string, params *stripe.CustomerTaxExemptionParams) (*stripe.CustomerTaxExemption, error) {
	return getC().Get(id, params)
}

// Retrieve a location specific tax exemption for a customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.CustomerTaxExemptionParams) (*stripe.CustomerTaxExemption, error) {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/tax_exemptions/%s", stripe.StringValue(
			params.Customer), id)
	customertaxexemption := &stripe.CustomerTaxExemption{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, customertaxexemption)
	return customertaxexemption, err
}

// Delete a location specific tax exemption for a customer.
func Del(id string, params *stripe.CustomerTaxExemptionParams) (*stripe.CustomerTaxExemption, error) {
	return getC().Del(id, params)
}

// Delete a location specific tax exemption for a customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.CustomerTaxExemptionParams) (*stripe.CustomerTaxExemption, error) {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/tax_exemptions/%s", stripe.StringValue(
			params.Customer), id)
	customertaxexemption := &stripe.CustomerTaxExemption{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, customertaxexemption)
	return customertaxexemption, err
}

// List all location specific tax exemptions for a customer.
func List(params *stripe.CustomerTaxExemptionListParams) *Iter {
	return getC().List(params)
}

// List all location specific tax exemptions for a customer.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.CustomerTaxExemptionListParams) *Iter {
	path := stripe.FormatURLPath(
		"/v1/customers/%s/tax_exemptions", stripe.StringValue(listParams.Customer))
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.CustomerTaxExemptionList{}
			err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for customer tax exemptions.
type Iter struct {
	*stripe.Iter
}

// CustomerTaxExemption returns the customer tax exemption which the iterator is currently pointing to.
func (i *Iter) CustomerTaxExemption() *stripe.CustomerTaxExemption {
	return i.Current().(*stripe.CustomerTaxExemption)
}

// CustomerTaxExemptionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) CustomerTaxExemptionList() *stripe.CustomerTaxExemptionList {
	return i.List().(*stripe.CustomerTaxExemptionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
