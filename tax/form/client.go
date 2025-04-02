//
//
// File generated from our OpenAPI spec
//
//

// Package form provides the /v1/tax/forms APIs
package form

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/form"
)

// Client is used to invoke /v1/tax/forms APIs.
type Client struct {
	B        stripe.Backend
	BUploads stripe.Backend
	Key      string
}

// Retrieves the details of a tax form that has previously been created. Supply the unique tax form ID that was returned from your previous request, and Stripe will return the corresponding tax form information.
func Get(id string, params *stripe.TaxFormParams) (*stripe.TaxForm, error) {
	return getC().Get(id, params)
}

// Retrieves the details of a tax form that has previously been created. Supply the unique tax form ID that was returned from your previous request, and Stripe will return the corresponding tax form information.
func (c Client) Get(id string, params *stripe.TaxFormParams) (*stripe.TaxForm, error) {
	path := stripe.FormatURLPath("/v1/tax/forms/%s", id)
	form := &stripe.TaxForm{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, form)
	return form, err
}

// Download the PDF for a tax form.
func PDF(id string, params *stripe.TaxFormPDFParams) (*stripe.APIStream, error) {
	return getC().PDF(id, params)
}

// Download the PDF for a tax form.
func (c Client) PDF(id string, params *stripe.TaxFormPDFParams) (*stripe.APIStream, error) {
	path := stripe.FormatURLPath("/v1/tax/forms/%s/pdf", id)
	stream := &stripe.APIStream{}
	err := c.BUploads.CallStreaming(http.MethodGet, path, c.Key, params, stream)
	return stream, err
}

// Returns a list of tax forms which were previously created. The tax forms are returned in sorted order, with the oldest tax forms appearing first.
func List(params *stripe.TaxFormListParams) *Iter {
	return getC().List(params)
}

// Returns a list of tax forms which were previously created. The tax forms are returned in sorted order, with the oldest tax forms appearing first.
func (c Client) List(listParams *stripe.TaxFormListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TaxFormList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/tax/forms", c.Key, []byte(b.Encode()), p, list)

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
