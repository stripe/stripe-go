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

// v2IamActivityLogService is used to invoke activitylog related APIs.
type v2IamActivityLogService struct {
	B   Backend
	Key string
}

// List activity logs of an account.
func (c v2IamActivityLogService) List(ctx context.Context, listParams *V2IamActivityLogListParams) *V2List[*V2IamActivityLog] {
	if listParams == nil {
		listParams = &V2IamActivityLogListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/iam/activity_logs", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2IamActivityLog], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2IamActivityLog]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
