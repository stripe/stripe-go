//
//
// File generated from our OpenAPI spec
//
//

// Package form provides the /tax/forms APIs
package form

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /tax/forms APIs.
type Client struct {
	B        stripe.Backend
	BUploads stripe.Backend
	Key      string
}

// Get returns the details of a tax form.
func Get(id string, params *stripe.TaxFormParams) (*stripe.TaxForm, error) {
	return getC().Get(id, params)
}

// Get returns the details of a tax form.
func (c Client) Get(id string, params *stripe.TaxFormParams) (*stripe.TaxForm, error) {
	path := stripe.FormatURLPath("/v1/tax/forms/%s", id)
	form := &stripe.TaxForm{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, form)
	return form, err
}

// PDF is the method for the `GET /v1/tax/forms/{id}/pdf` API.
func PDF(id string, params *stripe.TaxFormPDFParams) (*stripe.APIStream, error) {
	return getC().PDF(id, params)
}

// PDF is the method for the `GET /v1/tax/forms/{id}/pdf` API.
func (c Client) PDF(id string, params *stripe.TaxFormPDFParams) (*stripe.APIStream, error) {
	path := stripe.FormatURLPath("/v1/tax/forms/%s/pdf", id)
	stream := &stripe.APIStream{}
	err := c.BUploads.CallStreaming(http.MethodGet, path, c.Key, params, stream)
	return stream, err
}

// List returns a list of tax forms.
func List(params *stripe.TaxFormListParams) *Iter {
	return getC().List(params)
}

// List returns a list of tax forms.
func (c Client) List(listParams *stripe.TaxFormListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TaxFormList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/tax/forms", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for tax forms.
type Iter struct {
	*stripe.Iter
}

// TaxForm returns the tax form which the iterator is currently pointing to.
func (i *Iter) TaxForm() *stripe.TaxForm {
	return i.Current().(*stripe.TaxForm)
}

// TaxFormList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TaxFormList() *stripe.TaxFormList {
	return i.List().(*stripe.TaxFormList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.GetBackend(stripe.UploadsBackend), stripe.Key}
}
