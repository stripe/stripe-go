//
//
// File generated from our OpenAPI spec
//
//

// Package financialaccount provides the /treasury/financial_accounts APIs
package financialaccount

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /treasury/financial_accounts APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new treasury financial account.
func New(params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	return getC().New(params)
}

// New creates a new treasury financial account.
func (c Client) New(params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	financialaccount := &stripe.TreasuryFinancialAccount{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/treasury/financial_accounts",
		c.Key,
		params,
		financialaccount,
	)
	return financialaccount, err
}

// Get returns the details of a treasury financial account.
func Get(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	return getC().Get(id, params)
}

// Get returns the details of a treasury financial account.
func (c Client) Get(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	path := stripe.FormatURLPath("/v1/treasury/financial_accounts/%s", id)
	financialaccount := &stripe.TreasuryFinancialAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, financialaccount)
	return financialaccount, err
}

// Update updates a treasury financial account's properties.
func Update(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	return getC().Update(id, params)
}

// Update updates a treasury financial account's properties.
func (c Client) Update(id string, params *stripe.TreasuryFinancialAccountParams) (*stripe.TreasuryFinancialAccount, error) {
	path := stripe.FormatURLPath("/v1/treasury/financial_accounts/%s", id)
	financialaccount := &stripe.TreasuryFinancialAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, financialaccount)
	return financialaccount, err
}

// RetrieveFeatures is the method for the `GET /v1/treasury/financial_accounts/{financial_account}/features` API.
func RetrieveFeatures(id string, params *stripe.TreasuryFinancialAccountRetrieveFeaturesParams) (*stripe.TreasuryFinancialAccount, error) {
	return getC().RetrieveFeatures(id, params)
}

// RetrieveFeatures is the method for the `GET /v1/treasury/financial_accounts/{financial_account}/features` API.
func (c Client) RetrieveFeatures(id string, params *stripe.TreasuryFinancialAccountRetrieveFeaturesParams) (*stripe.TreasuryFinancialAccount, error) {
	path := stripe.FormatURLPath(
		"/v1/treasury/financial_accounts/%s/features",
		id,
	)
	financialaccount := &stripe.TreasuryFinancialAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, financialaccount)
	return financialaccount, err
}

// UpdateFeatures is the method for the `POST /v1/treasury/financial_accounts/{financial_account}/features` API.
func UpdateFeatures(id string, params *stripe.TreasuryFinancialAccountUpdateFeaturesParams) (*stripe.TreasuryFinancialAccount, error) {
	return getC().UpdateFeatures(id, params)
}

// UpdateFeatures is the method for the `POST /v1/treasury/financial_accounts/{financial_account}/features` API.
func (c Client) UpdateFeatures(id string, params *stripe.TreasuryFinancialAccountUpdateFeaturesParams) (*stripe.TreasuryFinancialAccount, error) {
	path := stripe.FormatURLPath(
		"/v1/treasury/financial_accounts/%s/features",
		id,
	)
	financialaccount := &stripe.TreasuryFinancialAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, financialaccount)
	return financialaccount, err
}

// List returns a list of treasury financial accounts.
func List(params *stripe.TreasuryFinancialAccountListParams) *Iter {
	return getC().List(params)
}

// List returns a list of treasury financial accounts.
func (c Client) List(listParams *stripe.TreasuryFinancialAccountListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.TreasuryFinancialAccountList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/treasury/financial_accounts", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for treasury financial accounts.
type Iter struct {
	*stripe.Iter
}

// TreasuryFinancialAccount returns the treasury financial account which the iterator is currently pointing to.
func (i *Iter) TreasuryFinancialAccount() *stripe.TreasuryFinancialAccount {
	return i.Current().(*stripe.TreasuryFinancialAccount)
}

// TreasuryFinancialAccountList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) TreasuryFinancialAccountList() *stripe.TreasuryFinancialAccountList {
	return i.List().(*stripe.TreasuryFinancialAccountList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
