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

// v1SigmaSchemaService is used to invoke /v1/sigma/schemas APIs.
type v1SigmaSchemaService struct {
	B   Backend
	Key string
}

// Lists the schemas available to the authorized merchant in Stripe's data products
func (c v1SigmaSchemaService) List(ctx context.Context, listParams *SigmaSchemaListParams) Seq2[*SigmaSchema, error] {
	if listParams == nil {
		listParams = &SigmaSchemaListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*SigmaSchema, ListContainer, error) {
		list := &SigmaSchemaList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/sigma/schemas", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
