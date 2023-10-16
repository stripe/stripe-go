//
//
// File generated from our OpenAPI spec
//
//

// Package transaction provides the /issuing/transactions APIs
package transaction

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /issuing/transactions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// CreateForceCapture is the method for the `POST /v1/test_helpers/issuing/transactions/create_force_capture` API.
func CreateForceCapture(params *stripe.TestHelpersIssuingTransactionCreateForceCaptureParams) (*stripe.IssuingTransaction, error) {
	return getC().CreateForceCapture(params)
}

// CreateForceCapture is the method for the `POST /v1/test_helpers/issuing/transactions/create_force_capture` API.
func (c Client) CreateForceCapture(params *stripe.TestHelpersIssuingTransactionCreateForceCaptureParams) (*stripe.IssuingTransaction, error) {
	transaction := &stripe.IssuingTransaction{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/test_helpers/issuing/transactions/create_force_capture",
		c.Key,
		params,
		transaction,
	)
	return transaction, err
}

// CreateUnlinkedRefund is the method for the `POST /v1/test_helpers/issuing/transactions/create_unlinked_refund` API.
func CreateUnlinkedRefund(params *stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundParams) (*stripe.IssuingTransaction, error) {
	return getC().CreateUnlinkedRefund(params)
}

// CreateUnlinkedRefund is the method for the `POST /v1/test_helpers/issuing/transactions/create_unlinked_refund` API.
func (c Client) CreateUnlinkedRefund(params *stripe.TestHelpersIssuingTransactionCreateUnlinkedRefundParams) (*stripe.IssuingTransaction, error) {
	transaction := &stripe.IssuingTransaction{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/test_helpers/issuing/transactions/create_unlinked_refund",
		c.Key,
		params,
		transaction,
	)
	return transaction, err
}

// Refund is the method for the `POST /v1/test_helpers/issuing/transactions/{transaction}/refund` API.
func Refund(id string, params *stripe.TestHelpersIssuingTransactionRefundParams) (*stripe.IssuingTransaction, error) {
	return getC().Refund(id, params)
}

// Refund is the method for the `POST /v1/test_helpers/issuing/transactions/{transaction}/refund` API.
func (c Client) Refund(id string, params *stripe.TestHelpersIssuingTransactionRefundParams) (*stripe.IssuingTransaction, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/transactions/%s/refund",
		id,
	)
	transaction := &stripe.IssuingTransaction{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, transaction)
	return transaction, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
