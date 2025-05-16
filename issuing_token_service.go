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

// v1IssuingTokenService is used to invoke /v1/issuing/tokens APIs.
type v1IssuingTokenService struct {
	B   Backend
	Key string
}

// Retrieves an Issuing Token object.
func (c v1IssuingTokenService) Retrieve(ctx context.Context, id string, params *IssuingTokenRetrieveParams) (*IssuingToken, error) {
	if params == nil {
		params = &IssuingTokenRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/tokens/%s", id)
	token := &IssuingToken{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, token)
	return token, err
}

// Attempts to update the specified Issuing Token object to the status specified.
func (c v1IssuingTokenService) Update(ctx context.Context, id string, params *IssuingTokenUpdateParams) (*IssuingToken, error) {
	if params == nil {
		params = &IssuingTokenUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/issuing/tokens/%s", id)
	token := &IssuingToken{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, token)
	return token, err
}

// Lists all Issuing Token objects for a given card.
func (c v1IssuingTokenService) List(ctx context.Context, listParams *IssuingTokenListParams) Seq2[*IssuingToken, error] {
	if listParams == nil {
		listParams = &IssuingTokenListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*IssuingToken, ListContainer, error) {
		list := &IssuingTokenList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/issuing/tokens", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
