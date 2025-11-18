//
//
// File generated from our OpenAPI spec
//
//

// Package institution provides the /v1/financial_connections/institutions APIs
package institution

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/form"
)

// Client is used to invoke /v1/financial_connections/institutions APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves the details of a Financial Connections Institution.
func Get(id string, params *stripe.FinancialConnectionsInstitutionParams) (*stripe.FinancialConnectionsInstitution, error) {
	return getC().Get(id, params)
}

// Retrieves the details of a Financial Connections Institution.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.FinancialConnectionsInstitutionParams) (*stripe.FinancialConnectionsInstitution, error) {
	path := stripe.FormatURLPath("/v1/financial_connections/institutions/%s", id)
	institution := &stripe.FinancialConnectionsInstitution{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, institution)
	return institution, err
}

// Returns a list of Financial Connections Institution objects.
func List(params *stripe.FinancialConnectionsInstitutionListParams) *Iter {
	return getC().List(params)
}

// Returns a list of Financial Connections Institution objects.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.FinancialConnectionsInstitutionListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.FinancialConnectionsInstitutionList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/financial_connections/institutions", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for financial connections institutions.
type Iter struct {
	*stripe.Iter
}

// FinancialConnectionsInstitution returns the financial connections institution which the iterator is currently pointing to.
func (i *Iter) FinancialConnectionsInstitution() *stripe.FinancialConnectionsInstitution {
	return i.Current().(*stripe.FinancialConnectionsInstitution)
}

// FinancialConnectionsInstitutionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) FinancialConnectionsInstitutionList() *stripe.FinancialConnectionsInstitutionList {
	return i.List().(*stripe.FinancialConnectionsInstitutionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
