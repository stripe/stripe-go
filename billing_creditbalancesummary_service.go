//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v1BillingCreditBalanceSummaryService is used to invoke /v1/billing/credit_balance_summary APIs.
type v1BillingCreditBalanceSummaryService struct {
	B   Backend
	Key string
}

// Retrieves the credit balance summary for a customer.
func (c v1BillingCreditBalanceSummaryService) Retrieve(ctx context.Context, params *BillingCreditBalanceSummaryRetrieveParams) (*BillingCreditBalanceSummary, error) {
	if params == nil {
		params = &BillingCreditBalanceSummaryRetrieveParams{}
	}
	params.Context = ctx
	creditbalancesummary := &BillingCreditBalanceSummary{}
	err := c.B.Call(
		http.MethodGet, "/v1/billing/credit_balance_summary", c.Key, params, creditbalancesummary)
	return creditbalancesummary, err
}
