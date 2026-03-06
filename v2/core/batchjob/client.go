//
//
// File generated from our OpenAPI spec
//
//

// Package batchjob provides the batchjob related APIs
package batchjob

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
)

// Client is used to invoke batchjob related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new batch job.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2CoreBatchJobParams) (*stripe.V2CoreBatchJob, error) {
	batchjob := &stripe.V2CoreBatchJob{}
	err := c.B.Call(
		http.MethodPost, "/v2/core/batch_jobs", c.Key, params, batchjob)
	return batchjob, err
}
