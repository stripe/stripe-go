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

// v1MandateService is used to invoke /v1/mandates APIs.
type v1MandateService struct {
	B   Backend
	Key string
}

// Retrieves a Mandate object.
func (c v1MandateService) Retrieve(ctx context.Context, id string, params *MandateRetrieveParams) (*Mandate, error) {
	if params == nil {
		params = &MandateRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/mandates/%s", id)
	mandate := &Mandate{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, mandate)
	return mandate, err
}

// Retrieves a list of Mandates for a given PaymentMethod.
func (c v1MandateService) List(ctx context.Context, listParams *MandateListParams) Seq2[*Mandate, error] {
	if listParams == nil {
		listParams = &MandateListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*Mandate, ListContainer, error) {
		list := &MandateList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/mandates", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
