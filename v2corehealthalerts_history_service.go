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

// v2CoreHealthAlertsHistoryService is used to invoke history related APIs.
type v2CoreHealthAlertsHistoryService struct {
	B   Backend
	Key string
}

// Retrieves a list of alert history entries for a health alert.
func (c v2CoreHealthAlertsHistoryService) List(ctx context.Context, listParams *V2CoreHealthAlertsHistoryListParams) *V2List[*V2CoreHealthAlertHistoryEntry] {
	if listParams == nil {
		listParams = &V2CoreHealthAlertsHistoryListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v2/core/health/alerts/%s/history", StringValue(listParams.ID))
	return newV2List(ctx, path, listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2CoreHealthAlertHistoryEntry], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2CoreHealthAlertHistoryEntry]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
