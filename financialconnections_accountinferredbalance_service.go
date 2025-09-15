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

// v1FinancialConnectionsAccountInferredBalanceService is used to invoke /v1/financial_connections/accounts/{account}/inferred_balances APIs.
type v1FinancialConnectionsAccountInferredBalanceService struct {
	B   Backend
	Key string
}

// Lists the recorded inferred balances for a Financial Connections Account.
func (c v1FinancialConnectionsAccountInferredBalanceService) List(ctx context.Context, listParams *FinancialConnectionsAccountInferredBalanceListParams) Seq2[*FinancialConnectionsAccountInferredBalance, error] {
	if listParams == nil {
		listParams = &FinancialConnectionsAccountInferredBalanceListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/financial_connections/accounts/%s/inferred_balances", StringValue(
			listParams.Account))
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*FinancialConnectionsAccountInferredBalance], error) {
		list := &v1Page[*FinancialConnectionsAccountInferredBalance]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
