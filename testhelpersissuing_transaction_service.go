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

// v1TestHelpersIssuingTransactionService is used to invoke /v1/issuing/transactions APIs.
type v1TestHelpersIssuingTransactionService struct {
	B   Backend
	Key string
}

// Allows the user to capture an arbitrary amount, also known as a forced capture.
func (c v1TestHelpersIssuingTransactionService) CreateForceCapture(ctx context.Context, params *TestHelpersIssuingTransactionCreateForceCaptureParams) (*IssuingTransaction, error) {
	if params == nil {
		params = &TestHelpersIssuingTransactionCreateForceCaptureParams{}
	}
	params.Context = ctx
	transaction := &IssuingTransaction{}
	err := c.B.Call(
		http.MethodPost, "/v1/test_helpers/issuing/transactions/create_force_capture", c.Key, params, transaction)
	return transaction, err
}

// Allows the user to refund an arbitrary amount, also known as a unlinked refund.
func (c v1TestHelpersIssuingTransactionService) CreateUnlinkedRefund(ctx context.Context, params *TestHelpersIssuingTransactionCreateUnlinkedRefundParams) (*IssuingTransaction, error) {
	if params == nil {
		params = &TestHelpersIssuingTransactionCreateUnlinkedRefundParams{}
	}
	params.Context = ctx
	transaction := &IssuingTransaction{}
	err := c.B.Call(
		http.MethodPost, "/v1/test_helpers/issuing/transactions/create_unlinked_refund", c.Key, params, transaction)
	return transaction, err
}

// Refund a test-mode Transaction.
func (c v1TestHelpersIssuingTransactionService) Refund(ctx context.Context, id string, params *TestHelpersIssuingTransactionRefundParams) (*IssuingTransaction, error) {
	if params == nil {
		params = &TestHelpersIssuingTransactionRefundParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/test_helpers/issuing/transactions/%s/refund", id)
	transaction := &IssuingTransaction{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, transaction)
	return transaction, err
}
