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

// v1CustomerCashBalanceTransactionService is used to invoke /v1/customers/{customer}/cash_balance_transactions APIs.
type v1CustomerCashBalanceTransactionService struct {
	B   Backend
	Key string
}

// Retrieves a specific cash balance transaction, which updated the customer's [cash balance](https://docs.stripe.com/docs/payments/customer-balance).
func (c v1CustomerCashBalanceTransactionService) Retrieve(ctx context.Context, id string, params *CustomerCashBalanceTransactionRetrieveParams) (*CustomerCashBalanceTransaction, error) {
	if params == nil {
		params = &CustomerCashBalanceTransactionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/cash_balance_transactions/%s", StringValue(
			params.Customer), id)
	customercashbalancetransaction := &CustomerCashBalanceTransaction{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, customercashbalancetransaction)
	return customercashbalancetransaction, err
}

// Returns a list of transactions that modified the customer's [cash balance](https://docs.stripe.com/docs/payments/customer-balance).
func (c v1CustomerCashBalanceTransactionService) List(ctx context.Context, listParams *CustomerCashBalanceTransactionListParams) Seq2[*CustomerCashBalanceTransaction, error] {
	if listParams == nil {
		listParams = &CustomerCashBalanceTransactionListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/cash_balance_transactions", StringValue(
			listParams.Customer))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*CustomerCashBalanceTransaction, ListContainer, error) {
		list := &CustomerCashBalanceTransactionList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
