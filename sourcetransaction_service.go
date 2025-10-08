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

// v1SourceTransactionService is used to invoke sourcetransaction related APIs.
type v1SourceTransactionService struct {
	B   Backend
	Key string
}

// List source transactions for a given source.
func (c v1SourceTransactionService) List(ctx context.Context, listParams *SourceTransactionListParams) Seq2[*SourceTransaction, error] {
	if listParams == nil {
		listParams = &SourceTransactionListParams{}
	}
	listParams.Context = ctx
	path := FormatURLPath(
		"/v1/sources/%s/source_transactions", StringValue(listParams.Source))
	return newV1List(ctx, listParams, func(ctx context.Context, p *Params, b *form.Values) (*V1Page[*SourceTransaction], error) {
		list := &V1Page[*SourceTransaction]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
