//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v84/form"
)

// v1FinancialConnectionsInstitutionService is used to invoke /v1/financial_connections/institutions APIs.
type v1FinancialConnectionsInstitutionService struct {
	B   Backend
	Key string
}

// Retrieves the details of a Financial Connections Institution.
func (c v1FinancialConnectionsInstitutionService) Retrieve(ctx context.Context, id string, params *FinancialConnectionsInstitutionRetrieveParams) (*FinancialConnectionsInstitution, error) {
	if params == nil {
		params = &FinancialConnectionsInstitutionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/financial_connections/institutions/%s", id)
	institution := &FinancialConnectionsInstitution{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, institution)
	return institution, err
}

// Returns a list of Financial Connections Institution objects.
func (c v1FinancialConnectionsInstitutionService) List(ctx context.Context, listParams *FinancialConnectionsInstitutionListParams) Seq2[*FinancialConnectionsInstitution, error] {
	if listParams == nil {
		listParams = &FinancialConnectionsInstitutionListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*FinancialConnectionsInstitution], error) {
		list := &v1Page[*FinancialConnectionsInstitution]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/financial_connections/institutions", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
