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

// v1TestHelpersRefundService is used to invoke /v1/refunds APIs.
type v1TestHelpersRefundService struct {
	B   Backend
	Key string
}

// Expire a refund with a status of requires_action.
func (c v1TestHelpersRefundService) Expire(ctx context.Context, id string, params *TestHelpersRefundExpireParams) (*Refund, error) {
	if params == nil {
		params = &TestHelpersRefundExpireParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/test_helpers/refunds/%s/expire", id)
	refund := &Refund{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, refund)
	return refund, err
}
