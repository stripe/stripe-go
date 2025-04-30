//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1IssuingAuthorizationService is used to invoke /v1/issuing/authorizations APIs.
type v1IssuingAuthorizationService struct {
	B   Backend
	Key string
}

// Retrieves an Issuing Authorization object.
func (c v1IssuingAuthorizationService) Retrieve(ctx context.Context, id string, params *IssuingAuthorizationRetrieveParams) (*IssuingAuthorization, error) {
	path := FormatURLPath("/v1/issuing/authorizations/%s", id)
	authorization := &IssuingAuthorization{}
	if params == nil {
		params = &IssuingAuthorizationRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, authorization)
	return authorization, err
}

// Updates the specified Issuing Authorization object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
func (c v1IssuingAuthorizationService) Update(ctx context.Context, id string, params *IssuingAuthorizationUpdateParams) (*IssuingAuthorization, error) {
	path := FormatURLPath("/v1/issuing/authorizations/%s", id)
	authorization := &IssuingAuthorization{}
	if params == nil {
		params = &IssuingAuthorizationUpdateParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}

// [Deprecated] Approves a pending Issuing Authorization object. This request should be made within the timeout window of the [real-time authorization](https://stripe.com/docs/issuing/controls/real-time-authorizations) flow.
// This method is deprecated. Instead, [respond directly to the webhook request to approve an authorization](https://stripe.com/docs/issuing/controls/real-time-authorizations#authorization-handling).
// Deprecated:
func (c v1IssuingAuthorizationService) Approve(ctx context.Context, id string, params *IssuingAuthorizationApproveParams) (*IssuingAuthorization, error) {
	path := FormatURLPath("/v1/issuing/authorizations/%s/approve", id)
	authorization := &IssuingAuthorization{}
	if params == nil {
		params = &IssuingAuthorizationApproveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}

// [Deprecated] Declines a pending Issuing Authorization object. This request should be made within the timeout window of the [real time authorization](https://stripe.com/docs/issuing/controls/real-time-authorizations) flow.
// This method is deprecated. Instead, [respond directly to the webhook request to decline an authorization](https://stripe.com/docs/issuing/controls/real-time-authorizations#authorization-handling).
// Deprecated:
func (c v1IssuingAuthorizationService) Decline(ctx context.Context, id string, params *IssuingAuthorizationDeclineParams) (*IssuingAuthorization, error) {
	path := FormatURLPath("/v1/issuing/authorizations/%s/decline", id)
	authorization := &IssuingAuthorization{}
	if params == nil {
		params = &IssuingAuthorizationDeclineParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}

// Returns a list of Issuing Authorization objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingAuthorizationService) List(ctx context.Context, listParams *IssuingAuthorizationListParams) Seq2[*IssuingAuthorization, error] {
	if listParams == nil {
		listParams = &IssuingAuthorizationListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*IssuingAuthorization, ListContainer, error) {
		list := &IssuingAuthorizationList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/issuing/authorizations", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
