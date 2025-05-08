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

// v1ReviewService is used to invoke /v1/reviews APIs.
type v1ReviewService struct {
	B   Backend
	Key string
}

// Retrieves a Review object.
func (c v1ReviewService) Retrieve(ctx context.Context, id string, params *ReviewRetrieveParams) (*Review, error) {
	if params == nil {
		params = &ReviewRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/reviews/%s", id)
	review := &Review{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, review)
	return review, err
}

// Approves a Review object, closing it and removing it from the list of reviews.
func (c v1ReviewService) Approve(ctx context.Context, id string, params *ReviewApproveParams) (*Review, error) {
	if params == nil {
		params = &ReviewApproveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/reviews/%s/approve", id)
	review := &Review{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, review)
	return review, err
}

// Returns a list of Review objects that have open set to true. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1ReviewService) List(ctx context.Context, listParams *ReviewListParams) Seq2[*Review, error] {
	if listParams == nil {
		listParams = &ReviewListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Review, ListContainer, error) {
		list := &ReviewList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/reviews", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
