//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v86/form"
)

// v1BillingFeedbackOptionService is used to invoke /v1/billing/feedback_options APIs.
type v1BillingFeedbackOptionService struct {
	B   Backend
	Key string
}

// Creates a new feedback option.
func (c v1BillingFeedbackOptionService) Create(ctx context.Context, params *BillingFeedbackOptionCreateParams) (*BillingFeedbackOption, error) {
	if params == nil {
		params = &BillingFeedbackOptionCreateParams{}
	}
	params.Context = ctx
	feedbackoption := &BillingFeedbackOption{}
	err := c.B.Call(
		http.MethodPost, "/v1/billing/feedback_options", c.Key, params, feedbackoption)
	return feedbackoption, err
}

// Retrieves a feedback options object given an ID.
func (c v1BillingFeedbackOptionService) Retrieve(ctx context.Context, id string, params *BillingFeedbackOptionRetrieveParams) (*BillingFeedbackOption, error) {
	if params == nil {
		params = &BillingFeedbackOptionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/feedback_options/%s", id)
	feedbackoption := &BillingFeedbackOption{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, feedbackoption)
	return feedbackoption, err
}

// Updates the description of an existing feedback option.
func (c v1BillingFeedbackOptionService) Update(ctx context.Context, id string, params *BillingFeedbackOptionUpdateParams) (*BillingFeedbackOption, error) {
	if params == nil {
		params = &BillingFeedbackOptionUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/feedback_options/%s", id)
	feedbackoption := &BillingFeedbackOption{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, feedbackoption)
	return feedbackoption, err
}

// Deactivates a feedback option. Deactivated feedback options cannot be used in portal configurations.
func (c v1BillingFeedbackOptionService) Deactivate(ctx context.Context, id string, params *BillingFeedbackOptionDeactivateParams) (*BillingFeedbackOption, error) {
	if params == nil {
		params = &BillingFeedbackOptionDeactivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/billing/feedback_options/%s/deactivate", id)
	feedbackoption := &BillingFeedbackOption{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, feedbackoption)
	return feedbackoption, err
}

// An API method for listing the feedback options model
func (c v1BillingFeedbackOptionService) List(ctx context.Context, listParams *BillingFeedbackOptionListParams) *V1List[*BillingFeedbackOption] {
	if listParams == nil {
		listParams = &BillingFeedbackOptionListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*v1Page[*BillingFeedbackOption], error) {
		list := &v1Page[*BillingFeedbackOption]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/billing/feedback_options", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
