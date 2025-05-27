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

// v1CustomerBalanceTransactionService is used to invoke /v1/customers/{customer}/balance_transactions APIs.
type v1CustomerBalanceTransactionService struct {
	B   Backend
	Key string
}

// Creates an immutable transaction that updates the customer's credit [balance](https://docs.stripe.com/docs/billing/customer/balance).
func (c v1CustomerBalanceTransactionService) Create(ctx context.Context, params *CustomerBalanceTransactionCreateParams) (*CustomerBalanceTransaction, error) {
	if params == nil {
		params = &CustomerBalanceTransactionCreateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/balance_transactions", StringValue(params.Customer))
	customerbalancetransaction := &CustomerBalanceTransaction{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, customerbalancetransaction)
	return customerbalancetransaction, err
}

// Retrieves a specific customer balance transaction that updated the customer's [balances](https://docs.stripe.com/docs/billing/customer/balance).
func (c v1CustomerBalanceTransactionService) Retrieve(ctx context.Context, id string, params *CustomerBalanceTransactionRetrieveParams) (*CustomerBalanceTransaction, error) {
	if params == nil {
		params = &CustomerBalanceTransactionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/balance_transactions/%s", StringValue(
			params.Customer), id)
	customerbalancetransaction := &CustomerBalanceTransaction{}
	err := c.B.Call(
		http.MethodGet, path, c.Key, params, customerbalancetransaction)
	return customerbalancetransaction, err
}

// Most credit balance transaction fields are immutable, but you may update its description and metadata.
func (c v1CustomerBalanceTransactionService) Update(ctx context.Context, id string, params *CustomerBalanceTransactionUpdateParams) (*CustomerBalanceTransaction, error) {
	if params == nil {
		params = &CustomerBalanceTransactionUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/balance_transactions/%s", StringValue(
			params.Customer), id)
	customerbalancetransaction := &CustomerBalanceTransaction{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, customerbalancetransaction)
	return customerbalancetransaction, err
}

// Returns a list of transactions that updated the customer's [balances](https://docs.stripe.com/docs/billing/customer/balance).
func (c v1CustomerBalanceTransactionService) List(ctx context.Context, listParams *CustomerBalanceTransactionListParams) Seq2[*CustomerBalanceTransaction, error] {
	if listParams == nil {
		listParams = &CustomerBalanceTransactionListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/customers/%s/balance_transactions", StringValue(listParams.Customer))
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*CustomerBalanceTransaction, ListContainer, error) {
		list := &CustomerBalanceTransactionList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
