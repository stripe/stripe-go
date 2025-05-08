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

// v1TreasuryFinancialAccountService is used to invoke /v1/treasury/financial_accounts APIs.
type v1TreasuryFinancialAccountService struct {
	B   Backend
	Key string
}

// Creates a new FinancialAccount. Each connected account can have up to three FinancialAccounts by default.
func (c v1TreasuryFinancialAccountService) Create(ctx context.Context, params *TreasuryFinancialAccountCreateParams) (*TreasuryFinancialAccount, error) {
	if params == nil {
		params = &TreasuryFinancialAccountCreateParams{}
	}
	params.Context = ctx
	financialaccount := &TreasuryFinancialAccount{}
	err := c.B.Call(
		http.MethodPost, "/v1/treasury/financial_accounts", c.Key, params, financialaccount)
	return financialaccount, err
}

// Retrieves the details of a FinancialAccount.
func (c v1TreasuryFinancialAccountService) Retrieve(ctx context.Context, id string, params *TreasuryFinancialAccountRetrieveParams) (*TreasuryFinancialAccount, error) {
	if params == nil {
		params = &TreasuryFinancialAccountRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/financial_accounts/%s", id)
	financialaccount := &TreasuryFinancialAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, financialaccount)
	return financialaccount, err
}

// Updates the details of a FinancialAccount.
func (c v1TreasuryFinancialAccountService) Update(ctx context.Context, id string, params *TreasuryFinancialAccountUpdateParams) (*TreasuryFinancialAccount, error) {
	if params == nil {
		params = &TreasuryFinancialAccountUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/financial_accounts/%s", id)
	financialaccount := &TreasuryFinancialAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, financialaccount)
	return financialaccount, err
}

// Closes a FinancialAccount. A FinancialAccount can only be closed if it has a zero balance, has no pending InboundTransfers, and has canceled all attached Issuing cards.
func (c v1TreasuryFinancialAccountService) Close(ctx context.Context, id string, params *TreasuryFinancialAccountCloseParams) (*TreasuryFinancialAccount, error) {
	if params == nil {
		params = &TreasuryFinancialAccountCloseParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/financial_accounts/%s/close", id)
	financialaccount := &TreasuryFinancialAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, financialaccount)
	return financialaccount, err
}

// Retrieves Features information associated with the FinancialAccount.
func (c v1TreasuryFinancialAccountService) RetrieveFeatures(ctx context.Context, id string, params *TreasuryFinancialAccountRetrieveFeaturesParams) (*TreasuryFinancialAccountFeatures, error) {
	if params == nil {
		params = &TreasuryFinancialAccountRetrieveFeaturesParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/financial_accounts/%s/features", id)
	financialaccountfeatures := &TreasuryFinancialAccountFeatures{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, financialaccountfeatures)
	return financialaccountfeatures, err
}

// Updates the Features associated with a FinancialAccount.
func (c v1TreasuryFinancialAccountService) UpdateFeatures(ctx context.Context, id string, params *TreasuryFinancialAccountUpdateFeaturesParams) (*TreasuryFinancialAccountFeatures, error) {
	if params == nil {
		params = &TreasuryFinancialAccountUpdateFeaturesParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/treasury/financial_accounts/%s/features", id)
	financialaccountfeatures := &TreasuryFinancialAccountFeatures{}
	err := c.B.Call(
		http.MethodPost, path, c.Key, params, financialaccountfeatures)
	return financialaccountfeatures, err
}

// Returns a list of FinancialAccounts.
func (c v1TreasuryFinancialAccountService) List(ctx context.Context, listParams *TreasuryFinancialAccountListParams) Seq2[*TreasuryFinancialAccount, error] {
	if listParams == nil {
		listParams = &TreasuryFinancialAccountListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*TreasuryFinancialAccount, ListContainer, error) {
		list := &TreasuryFinancialAccountList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/treasury/financial_accounts", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
