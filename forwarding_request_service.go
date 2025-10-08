//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v83/form"
)

// v1ForwardingRequestService is used to invoke /v1/forwarding/requests APIs.
type v1ForwardingRequestService struct {
	B   Backend
	Key string
}

// Creates a ForwardingRequest object.
func (c v1ForwardingRequestService) Create(ctx context.Context, params *ForwardingRequestCreateParams) (*ForwardingRequest, error) {
	if params == nil {
		params = &ForwardingRequestCreateParams{}
	}
	params.Context = ctx
	request := &ForwardingRequest{}
	err := c.B.Call(
		http.MethodPost, "/v1/forwarding/requests", c.Key, params, request)
	return request, err
}

// Retrieves a ForwardingRequest object.
func (c v1ForwardingRequestService) Retrieve(ctx context.Context, id string, params *ForwardingRequestRetrieveParams) (*ForwardingRequest, error) {
	if params == nil {
		params = &ForwardingRequestRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/forwarding/requests/%s", id)
	request := &ForwardingRequest{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, request)
	return request, err
}

// Lists all ForwardingRequest objects.
func (c v1ForwardingRequestService) List(ctx context.Context, listParams *ForwardingRequestListParams) Seq2[*ForwardingRequest, error] {
	if listParams == nil {
		listParams = &ForwardingRequestListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*V1Page[*ForwardingRequest], error) {
		list := &V1Page[*ForwardingRequest]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/forwarding/requests", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
