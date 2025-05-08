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

// v1CountrySpecService is used to invoke /v1/country_specs APIs.
type v1CountrySpecService struct {
	B   Backend
	Key string
}

// Returns a Country Spec for a given Country code.
func (c v1CountrySpecService) Retrieve(ctx context.Context, id string, params *CountrySpecRetrieveParams) (*CountrySpec, error) {
	if params == nil {
		params = &CountrySpecRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/country_specs/%s", id)
	countryspec := &CountrySpec{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, countryspec)
	return countryspec, err
}

// Lists all Country Spec objects available in the API.
func (c v1CountrySpecService) List(ctx context.Context, listParams *CountrySpecListParams) Seq2[*CountrySpec, error] {
	if listParams == nil {
		listParams = &CountrySpecListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*CountrySpec, ListContainer, error) {
		list := &CountrySpecList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/country_specs", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
