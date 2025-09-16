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

// v1TaxFormService is used to invoke /v1/tax/forms APIs.
type v1TaxFormService struct {
	B        Backend
	BUploads Backend
	Key      string
}

// Retrieves the details of a tax form that has previously been created. Supply the unique tax form ID that was returned from your previous request, and Stripe will return the corresponding tax form information.
func (c v1TaxFormService) Retrieve(ctx context.Context, id string, params *TaxFormRetrieveParams) (*TaxForm, error) {
	if params == nil {
		params = &TaxFormRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/tax/forms/%s", id)
	form := &TaxForm{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, form)
	return form, err
}

// Download the PDF for a tax form.
func (c v1TaxFormService) PDF(ctx context.Context, id string, params *TaxFormPDFParams) (*APIStream, error) {
	if params == nil {
		params = &TaxFormPDFParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/tax/forms/%s/pdf", id)
	stream := &APIStream{}
	err := c.BUploads.CallStreaming(http.MethodGet, path, c.Key, params, stream)
	return stream, err
}

// Returns a list of tax forms which were previously created. The tax forms are returned in sorted order, with the oldest tax forms appearing first.
func (c v1TaxFormService) List(ctx context.Context, listParams *TaxFormListParams) Seq2[*TaxForm, error] {
	if listParams == nil {
		listParams = &TaxFormListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TaxForm, ListContainer, error) {
		list := &TaxFormList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/tax/forms", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
