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

// v1IssuingPhysicalBundleService is used to invoke /v1/issuing/physical_bundles APIs.
type v1IssuingPhysicalBundleService struct {
	B   Backend
	Key string
}

// Retrieves a physical bundle object.
func (c v1IssuingPhysicalBundleService) Retrieve(ctx context.Context, id string, params *IssuingPhysicalBundleRetrieveParams) (*IssuingPhysicalBundle, error) {
	if params == nil {
		params = &IssuingPhysicalBundleRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/physical_bundles/%s", id)
	physicalbundle := &IssuingPhysicalBundle{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, physicalbundle)
	return physicalbundle, err
}

// Returns a list of physical bundle objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
func (c v1IssuingPhysicalBundleService) List(ctx context.Context, listParams *IssuingPhysicalBundleListParams) Seq2[*IssuingPhysicalBundle, error] {
	if listParams == nil {
		listParams = &IssuingPhysicalBundleListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*IssuingPhysicalBundle, ListContainer, error) {
		list := &IssuingPhysicalBundleList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/issuing/physical_bundles", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
