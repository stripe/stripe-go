//
//
// File generated from our OpenAPI spec
//
//

// Package taxtransaction provides the /tax/transactions APIs
package taxtransaction

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v74"
)

// Client is used to invoke /tax/transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new tax transaction.
func New(params *stripe.TaxTransactionParams) (*stripe.TaxTransaction, error) {
	return getC().New(params)
}

// New creates a new tax transaction.
func (c Client) New(params *stripe.TaxTransactionParams) (*stripe.TaxTransaction, error) {
	taxtransaction := &stripe.TaxTransaction{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/tax/transactions",
		c.Key,
		params,
		taxtransaction,
	)
	return taxtransaction, err
}

// Get returns the details of a tax transaction.
func Get(id string, params *stripe.TaxTransactionParams) (*stripe.TaxTransaction, error) {
	return getC().Get(id, params)
}

// Get returns the details of a tax transaction.
func (c Client) Get(id string, params *stripe.TaxTransactionParams) (*stripe.TaxTransaction, error) {
	path := stripe.FormatURLPath("/v1/tax/transactions/%s", id)
	taxtransaction := &stripe.TaxTransaction{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, taxtransaction)
	return taxtransaction, err
}

// CreateReversal is the method for the `POST /v1/tax/transactions/create_reversal` API.
func CreateReversal(params *stripe.TaxTransactionCreateReversalParams) (*stripe.TaxTransaction, error) {
	return getC().CreateReversal(params)
}

// CreateReversal is the method for the `POST /v1/tax/transactions/create_reversal` API.
func (c Client) CreateReversal(params *stripe.TaxTransactionCreateReversalParams) (*stripe.TaxTransaction, error) {
	taxtransaction := &stripe.TaxTransaction{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/tax/transactions/create_reversal",
		c.Key,
		params,
		taxtransaction,
	)
	return taxtransaction, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
