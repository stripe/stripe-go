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

// v1OrchestrationPaymentAttemptService is used to invoke /v1/orchestration/payment_attempts APIs.
type v1OrchestrationPaymentAttemptService struct {
	B   Backend
	Key string
}

// Retrieves orchestration information for the given payment attempt record (e.g. return url).
func (c v1OrchestrationPaymentAttemptService) Retrieve(ctx context.Context, id string, params *OrchestrationPaymentAttemptRetrieveParams) (*OrchestrationPaymentAttempt, error) {
	if params == nil {
		params = &OrchestrationPaymentAttemptRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/orchestration/payment_attempts/%s", id)
	paymentattempt := &OrchestrationPaymentAttempt{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, paymentattempt)
	return paymentattempt, err
}
