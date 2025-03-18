//
//
// File generated from our OpenAPI spec
//
//

// Package invoicerenderingtemplate provides the /invoice_rendering_templates APIs
package invoicerenderingtemplate

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/form"
)

// Client is used to invoke /invoice_rendering_templates APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an invoice rendering template with the given ID. It by default returns the latest version of the template. Optionally, specify a version to see previous versions.
func Get(id string, params *stripe.InvoiceRenderingTemplateParams) (*stripe.InvoiceRenderingTemplate, error) {
	return getC().Get(id, params)
}

// Retrieves an invoice rendering template with the given ID. It by default returns the latest version of the template. Optionally, specify a version to see previous versions.
func (c Client) Get(id string, params *stripe.InvoiceRenderingTemplateParams) (*stripe.InvoiceRenderingTemplate, error) {
	path := stripe.FormatURLPath("/v1/invoice_rendering_templates/%s", id)
	invoicerenderingtemplate := &stripe.InvoiceRenderingTemplate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, invoicerenderingtemplate)
	return invoicerenderingtemplate, err
}

// Updates the status of an invoice rendering template to ‘archived' so no new Stripe objects (customers, invoices, etc.) can reference it. The template can also no longer be updated. However, if the template is already set on a Stripe object, it will continue to be applied on invoices generated by it.
func Archive(id string, params *stripe.InvoiceRenderingTemplateArchiveParams) (*stripe.InvoiceRenderingTemplate, error) {
	return getC().Archive(id, params)
}

// Updates the status of an invoice rendering template to ‘archived' so no new Stripe objects (customers, invoices, etc.) can reference it. The template can also no longer be updated. However, if the template is already set on a Stripe object, it will continue to be applied on invoices generated by it.
func (c Client) Archive(id string, params *stripe.InvoiceRenderingTemplateArchiveParams) (*stripe.InvoiceRenderingTemplate, error) {
	path := stripe.FormatURLPath("/v1/invoice_rendering_templates/%s/archive", id)
	invoicerenderingtemplate := &stripe.InvoiceRenderingTemplate{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, invoicerenderingtemplate)
	return invoicerenderingtemplate, err
}

// Unarchive an invoice rendering template so it can be used on new Stripe objects again.
func Unarchive(id string, params *stripe.InvoiceRenderingTemplateUnarchiveParams) (*stripe.InvoiceRenderingTemplate, error) {
	return getC().Unarchive(id, params)
}

// Unarchive an invoice rendering template so it can be used on new Stripe objects again.
func (c Client) Unarchive(id string, params *stripe.InvoiceRenderingTemplateUnarchiveParams) (*stripe.InvoiceRenderingTemplate, error) {
	path := stripe.FormatURLPath(
		"/v1/invoice_rendering_templates/%s/unarchive", id)
	invoicerenderingtemplate := &stripe.InvoiceRenderingTemplate{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, invoicerenderingtemplate)
	return invoicerenderingtemplate, err
}

// List all templates, ordered by creation date, with the most recently created template appearing first.
func List(params *stripe.InvoiceRenderingTemplateListParams) *Iter {
	return getC().List(params)
}

// List all templates, ordered by creation date, with the most recently created template appearing first.
func (c Client) List(listParams *stripe.InvoiceRenderingTemplateListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.InvoiceRenderingTemplateList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/invoice_rendering_templates", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for invoice rendering templates.
type Iter struct {
	*stripe.Iter
}

// InvoiceRenderingTemplate returns the invoice rendering template which the iterator is currently pointing to.
func (i *Iter) InvoiceRenderingTemplate() *stripe.InvoiceRenderingTemplate {
	return i.Current().(*stripe.InvoiceRenderingTemplate)
}

// InvoiceRenderingTemplateList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) InvoiceRenderingTemplateList() *stripe.InvoiceRenderingTemplateList {
	return i.List().(*stripe.InvoiceRenderingTemplateList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
