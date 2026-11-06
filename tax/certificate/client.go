//
//
// File generated from our OpenAPI spec
//
//

// Package certificate provides the /v1/tax/certificates APIs
package certificate

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v87"
	"github.com/stripe/stripe-go/v87/form"
)

// Client is used to invoke /v1/tax/certificates APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Create a tax certificate associated with a tax exemption.
func New(params *stripe.TaxCertificateParams) (*stripe.TaxCertificate, error) {
	return getC().New(params)
}

// Create a tax certificate associated with a tax exemption.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TaxCertificateParams) (*stripe.TaxCertificate, error) {
	certificate := &stripe.TaxCertificate{}
	err := c.B.Call(
		http.MethodPost, "/v1/tax/certificates", c.Key, params, certificate)
	return certificate, err
}

// Retrieve a tax certificate associated with a tax exemption.
func Get(id string, params *stripe.TaxCertificateParams) (*stripe.TaxCertificate, error) {
	return getC().Get(id, params)
}

// Retrieve a tax certificate associated with a tax exemption.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.TaxCertificateParams) (*stripe.TaxCertificate, error) {
	path := stripe.FormatURLPath("/v1/tax/certificates/%s", id)
	certificate := &stripe.TaxCertificate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, certificate)
	return certificate, err
}

// Delete a tax certificate associated with a tax exemption.
func Del(id string, params *stripe.TaxCertificateParams) (*stripe.TaxCertificate, error) {
	return getC().Del(id, params)
}

// Delete a tax certificate associated with a tax exemption.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Del(id string, params *stripe.TaxCertificateParams) (*stripe.TaxCertificate, error) {
	path := stripe.FormatURLPath("/v1/tax/certificates/%s", id)
	certificate := &stripe.TaxCertificate{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, certificate)
	return certificate, err
}

// List certificates for an existing tax exemption.
func List(params *stripe.TaxCertificateListParams) *Iter {
	return getC().List(params)
}

// List certificates for an existing tax exemption.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.TaxCertificateListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TaxCertificateList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/tax/certificates", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for tax certificates.
type Iter struct {
	*stripe.Iter
}

// TaxCertificate returns the tax certificate which the iterator is currently pointing to.
func (i *Iter) TaxCertificate() *stripe.TaxCertificate {
	return i.Current().(*stripe.TaxCertificate)
}

// TaxCertificateList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TaxCertificateList() *stripe.TaxCertificateList {
	return i.List().(*stripe.TaxCertificateList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
