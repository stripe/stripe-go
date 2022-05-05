//
//
// File generated from our OpenAPI spec
//
//

// Package account provides the /financial_connections/accounts APIs
package account

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke /financial_connections/accounts APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// GetByID returns the details of a financial connections account.
func GetByID(id string, params *stripe.FinancialConnectionsAccountParams) (*stripe.FinancialConnectionsAccount, error) {
	return getC().GetByID(id, params)
}

// GetByID returns the details of a financial connections account.
func (c Client) GetByID(id string, params *stripe.FinancialConnectionsAccountParams) (*stripe.FinancialConnectionsAccount, error) {
	path := stripe.FormatURLPath("/v1/financial_connections/accounts/%s", id)
	account := &stripe.FinancialConnectionsAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, account)
	return account, err
}

// Disconnect is the method for the `POST /v1/financial_connections/accounts/{account}/disconnect` API.
func Disconnect(id string, params *stripe.FinancialConnectionsAccountDisconnectParams) (*stripe.FinancialConnectionsAccount, error) {
	return getC().Disconnect(id, params)
}

// Disconnect is the method for the `POST /v1/financial_connections/accounts/{account}/disconnect` API.
func (c Client) Disconnect(id string, params *stripe.FinancialConnectionsAccountDisconnectParams) (*stripe.FinancialConnectionsAccount, error) {
	path := stripe.FormatURLPath(
		"/v1/financial_connections/accounts/%s/disconnect",
		id,
	)
	account := &stripe.FinancialConnectionsAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

// Refresh is the method for the `POST /v1/financial_connections/accounts/{account}/refresh` API.
func Refresh(id string, params *stripe.FinancialConnectionsAccountRefreshParams) (*stripe.FinancialConnectionsAccount, error) {
	return getC().Refresh(id, params)
}

// Refresh is the method for the `POST /v1/financial_connections/accounts/{account}/refresh` API.
func (c Client) Refresh(id string, params *stripe.FinancialConnectionsAccountRefreshParams) (*stripe.FinancialConnectionsAccount, error) {
	path := stripe.FormatURLPath(
		"/v1/financial_connections/accounts/%s/refresh",
		id,
	)
	account := &stripe.FinancialConnectionsAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, account)
	return account, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
