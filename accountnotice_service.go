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

// v1AccountNoticeService is used to invoke /v1/account_notices APIs.
type v1AccountNoticeService struct {
	B   Backend
	Key string
}

// Retrieves an AccountNotice object.
func (c v1AccountNoticeService) Retrieve(ctx context.Context, id string, params *AccountNoticeRetrieveParams) (*AccountNotice, error) {
	if params == nil {
		params = &AccountNoticeRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/account_notices/%s", id)
	accountnotice := &AccountNotice{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountnotice)
	return accountnotice, err
}

// Updates an AccountNotice object.
func (c v1AccountNoticeService) Update(ctx context.Context, id string, params *AccountNoticeUpdateParams) (*AccountNotice, error) {
	if params == nil {
		params = &AccountNoticeUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/account_notices/%s", id)
	accountnotice := &AccountNotice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, accountnotice)
	return accountnotice, err
}

// Retrieves a list of AccountNotice objects. The objects are sorted in descending order by creation date, with the most-recently-created object appearing first.
func (c v1AccountNoticeService) List(ctx context.Context, listParams *AccountNoticeListParams) Seq2[*AccountNotice, error] {
	if listParams == nil {
		listParams = &AccountNoticeListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*AccountNotice], error) {
		list := &v1Page[*AccountNotice]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/account_notices", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
