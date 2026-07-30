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

// v2RiskInquiryService is used to invoke inquiry related APIs.
type v2RiskInquiryService struct {
	B   Backend
	Key string
}

// Retrieves a risk inquiry by ID.
func (c v2RiskInquiryService) Retrieve(ctx context.Context, id string, params *V2RiskInquiryRetrieveParams) (*V2RiskInquiry, error) {
	if params == nil {
		params = &V2RiskInquiryRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/risk/inquiries/%s", id)
	inquiry := &V2RiskInquiry{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, inquiry)
	return inquiry, err
}

// Submits a response to a risk inquiry.
func (c v2RiskInquiryService) Update(ctx context.Context, id string, params *V2RiskInquiryUpdateParams) (*V2RiskInquiry, error) {
	if params == nil {
		params = &V2RiskInquiryUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/risk/inquiries/%s", id)
	inquiry := &V2RiskInquiry{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, inquiry)
	return inquiry, err
}

// Lists risk inquiries for a connected account.
func (c v2RiskInquiryService) List(ctx context.Context, listParams *V2RiskInquiryListParams) *V2List[*V2RiskInquiry] {
	if listParams == nil {
		listParams = &V2RiskInquiryListParams{}
	}
	listParams.Context = ctx
	return newV2List(ctx, "/v2/risk/inquiries", listParams, func(ctx context.Context, path string, p ParamsContainer) (*V2Page[*V2RiskInquiry], error) {
		if p.GetParams() != nil {
			p.GetParams().Context = ctx
		}
		page := &V2Page[*V2RiskInquiry]{}
		err := c.B.Call(http.MethodGet, path, c.Key, p, page)
		return page, err
	})
}
