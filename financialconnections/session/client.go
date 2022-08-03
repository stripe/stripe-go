//
//
// File generated from our OpenAPI spec
//
//

// Package session provides the /financial_connections/sessions APIs
package session

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
)

// Client is used to invoke /financial_connections/sessions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new financial connections session.
func New(params *stripe.FinancialConnectionsSessionParams) (*stripe.FinancialConnectionsSession, error) {
	return getC().New(params)
}

// New creates a new financial connections session.
func (c Client) New(params *stripe.FinancialConnectionsSessionParams) (*stripe.FinancialConnectionsSession, error) {
	session := &stripe.FinancialConnectionsSession{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/financial_connections/sessions",
		c.Key,
		params,
		session,
	)
	return session, err
}

// Get returns the details of a financial connections session.
func Get(id string, params *stripe.FinancialConnectionsSessionParams) (*stripe.FinancialConnectionsSession, error) {
	return getC().Get(id, params)
}

// Get returns the details of a financial connections session.
func (c Client) Get(id string, params *stripe.FinancialConnectionsSessionParams) (*stripe.FinancialConnectionsSession, error) {
	path := stripe.FormatURLPath("/v1/financial_connections/sessions/%s", id)
	session := &stripe.FinancialConnectionsSession{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, session)
	return session, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
