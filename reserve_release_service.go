//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v84/form"
)

// v1ReserveReleaseService is used to invoke /v1/reserve/releases APIs.
type v1ReserveReleaseService struct {
	B   Backend
	Key string
}

// Retrieve a ReserveRelease.
func (c v1ReserveReleaseService) Retrieve(ctx context.Context, id string, params *ReserveReleaseRetrieveParams) (*ReserveRelease, error) {
	if params == nil {
		params = &ReserveReleaseRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/reserve/releases/%s", id)
	release := &ReserveRelease{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, release)
	return release, err
}

// Returns a list of ReserveReleases previously created. The ReserveReleases are returned in sorted order, with the most recent ReserveReleases appearing first.
func (c v1ReserveReleaseService) List(ctx context.Context, listParams *ReserveReleaseListParams) Seq2[*ReserveRelease, error] {
	if listParams == nil {
		listParams = &ReserveReleaseListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*ReserveRelease], error) {
		list := &v1Page[*ReserveRelease]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/reserve/releases", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
