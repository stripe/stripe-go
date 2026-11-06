//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v87/form"
)

// v1TaxCertificateService is used to invoke /v1/tax/certificates APIs.
type v1TaxCertificateService struct {
	B   Backend
	Key string
}

// Create a tax certificate associated with a tax exemption.
func (c v1TaxCertificateService) Create(ctx context.Context, params *TaxCertificateCreateParams) (*TaxCertificate, error) {
	if params == nil {
		params = &TaxCertificateCreateParams{}
	}
	params.Context = ctx
	certificate := &TaxCertificate{}
	err := c.B.Call(
		http.MethodPost, "/v1/tax/certificates", c.Key, params, certificate)
	return certificate, err
}

// Retrieve a tax certificate associated with a tax exemption.
func (c v1TaxCertificateService) Retrieve(ctx context.Context, id string, params *TaxCertificateRetrieveParams) (*TaxCertificate, error) {
	if params == nil {
		params = &TaxCertificateRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/tax/certificates/%s", id)
	certificate := &TaxCertificate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, certificate)
	return certificate, err
}

// Delete a tax certificate associated with a tax exemption.
func (c v1TaxCertificateService) Delete(ctx context.Context, id string, params *TaxCertificateDeleteParams) (*TaxCertificate, error) {
	if params == nil {
		params = &TaxCertificateDeleteParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/tax/certificates/%s", id)
	certificate := &TaxCertificate{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, certificate)
	return certificate, err
}

// List certificates for an existing tax exemption.
func (c v1TaxCertificateService) List(ctx context.Context, listParams *TaxCertificateListParams) *V1List[*TaxCertificate] {
	if listParams == nil {
		listParams = &TaxCertificateListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*TaxCertificate], error) {
		list := &v1Page[*TaxCertificate]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/tax/certificates", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
