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

// v1AccountNoticeService is used to invoke /v1/account_notices APIs.
type v1AccountNoticeService struct {
	B   Backend
	Key string
}

// Retrieves an AccountNotice object.
func (c v1AccountNoticeService) Retrieve(ctx context.Context, id string, params *AccountNoticeRetrieveParams) (*AccountNotice, error) {
	path := FormatURLPath("/v1/account_notices/%s", id)
	accountnotice := &AccountNotice{}
	if params == nil {
		params = &AccountNoticeRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountnotice)
	return accountnotice, err
}

// Updates an AccountNotice object.
func (c v1AccountNoticeService) Update(ctx context.Context, id string, params *AccountNoticeUpdateParams) (*AccountNotice, error) {
	path := FormatURLPath("/v1/account_notices/%s", id)
	accountnotice := &AccountNotice{}
	if params == nil {
		params = &AccountNoticeUpdateParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodPost, path, c.Key, params, accountnotice)
	return accountnotice, err
}

// Retrieves a list of AccountNotice objects. The objects are sorted in descending order by creation date, with the most-recently-created object appearing first.
func (c v1AccountNoticeService) List(ctx context.Context, listParams *AccountNoticeListParams) Seq2[*AccountNotice, error] {
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*AccountNotice, ListContainer, error) {
		list := &AccountNoticeList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/account_notices", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
