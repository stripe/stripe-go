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

// v2CoreHealthAlertService is used to invoke alert related APIs.
type v2CoreHealthAlertService struct {
	B   Backend
	Key string
}

// Retrieves a health alert by ID.
func (c v2CoreHealthAlertService) Retrieve(ctx context.Context, id string, params *V2CoreHealthAlertRetrieveParams) (*V2CoreHealthAlert, error) {
	if params == nil {
		params = &V2CoreHealthAlertRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/core/health/alerts/%s", id)
	alert := &V2CoreHealthAlert{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, alert)
	return alert, err
}

// Retrieves a list of health alerts.
func (c v2CoreHealthAlertService) List(ctx context.Context, listParams *V2CoreHealthAlertListParams) *V2List[*V2CoreHealthAlert] {
	if listParams == nil {
		listParams = &V2CoreHealthAlertListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/core/health/alerts", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2CoreHealthAlert], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2CoreHealthAlert]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
