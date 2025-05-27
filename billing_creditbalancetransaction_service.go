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

// v1BillingCreditBalanceTransactionService is used to invoke /v1/billing/credit_balance_transactions APIs.
type v1BillingCreditBalanceTransactionService struct {
	B   Backend
	Key string
}

// Retrieves a credit balance transaction.
func (c v1BillingCreditBalanceTransactionService) Retrieve(ctx context.Context, id string, params *BillingCreditBalanceTransactionRetrieveParams) (*BillingCreditBalanceTransaction, error) {
	if params == nil {
		params = &BillingCreditBalanceTransactionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/credit_balance_transactions/%s", id)
	creditbalancetransaction := &BillingCreditBalanceTransaction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, creditbalancetransaction)
	return creditbalancetransaction, err
}

// Retrieve a list of credit balance transactions.
func (c v1BillingCreditBalanceTransactionService) List(ctx context.Context, listParams *BillingCreditBalanceTransactionListParams) Seq2[*BillingCreditBalanceTransaction, error] {
	if listParams == nil {
		listParams = &BillingCreditBalanceTransactionListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*BillingCreditBalanceTransaction, ListContainer, error) {
		list := &BillingCreditBalanceTransactionList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/billing/credit_balance_transactions", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
