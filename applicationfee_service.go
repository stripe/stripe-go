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

// v1ApplicationFeeService is used to invoke /v1/application_fees APIs.
type v1ApplicationFeeService struct {
	B   Backend
	Key string
}

// Retrieves the details of an application fee that your account has collected. The same information is returned when refunding the application fee.
func (c v1ApplicationFeeService) Retrieve(ctx context.Context, id string, params *ApplicationFeeRetrieveParams) (*ApplicationFee, error) {
	if params == nil {
		params = &ApplicationFeeRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/application_fees/%s", id)
	applicationfee := &ApplicationFee{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, applicationfee)
	return applicationfee, err
}

// Returns a list of application fees you've previously collected. The application fees are returned in sorted order, with the most recent fees appearing first.
func (c v1ApplicationFeeService) List(ctx context.Context, listParams *ApplicationFeeListParams) Seq2[*ApplicationFee, error] {
	return c.ListWithPage(ctx, listParams).All()
}

// Returns a list of application fees you've previously collected. The application fees are returned in sorted order, with the most recent fees appearing first.
func (c v1ApplicationFeeService) ListWithPage(ctx context.Context, listParams *ApplicationFeeListParams) *V1List[*ApplicationFee] {
	if listParams == nil {
		listParams = &ApplicationFeeListParams{}
	}
	listParams.Context = ctx
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*V1Page[*ApplicationFee], error) {
		list := &V1Page[*ApplicationFee]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/application_fees", c.Key, []byte(b.Encode()), p, list)
		return list, err
	})
}
