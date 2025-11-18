//
//
// File generated from our OpenAPI spec
//
//

// Package accountnotice provides the /v1/account_notices APIs
package accountnotice

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/form"
)

// Client is used to invoke /v1/account_notices APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Retrieves an AccountNotice object.
func Get(id string, params *stripe.AccountNoticeParams) (*stripe.AccountNotice, error) {
	return getC().Get(id, params)
}

// Retrieves an AccountNotice object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.AccountNoticeParams) (*stripe.AccountNotice, error) {
	path := stripe.FormatURLPath("/v1/account_notices/%s", id)
	accountnotice := &stripe.AccountNotice{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountnotice)
	return accountnotice, err
}

// Updates an AccountNotice object.
func Update(id string, params *stripe.AccountNoticeParams) (*stripe.AccountNotice, error) {
	return getC().Update(id, params)
}

// Updates an AccountNotice object.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Update(id string, params *stripe.AccountNoticeParams) (*stripe.AccountNotice, error) {
	path := stripe.FormatURLPath("/v1/account_notices/%s", id)
	accountnotice := &stripe.AccountNotice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, accountnotice)
	return accountnotice, err
}

// Retrieves a list of AccountNotice objects. The objects are sorted in descending order by creation date, with the most-recently-created object appearing first.
func List(params *stripe.AccountNoticeListParams) *Iter {
	return getC().List(params)
}

// Retrieves a list of AccountNotice objects. The objects are sorted in descending order by creation date, with the most-recently-created object appearing first.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.AccountNoticeListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.AccountNoticeList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/account_notices", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for account notices.
type Iter struct {
	*stripe.Iter
}

// AccountNotice returns the account notice which the iterator is currently pointing to.
func (i *Iter) AccountNotice() *stripe.AccountNotice {
	return i.Current().(*stripe.AccountNotice)
}

// AccountNoticeList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) AccountNoticeList() *stripe.AccountNoticeList {
	return i.List().(*stripe.AccountNoticeList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
