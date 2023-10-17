//
//
// File generated from our OpenAPI spec
//
//

// Package accountnotice provides the /account_notices APIs
package accountnotice

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
	"github.com/stripe/stripe-go/v76/form"
)

// Client is used to invoke /account_notices APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of an account notice.
func Get(id string, params *stripe.AccountNoticeParams) (*stripe.AccountNotice, error) {
	return getC().Get(id, params)
}

// Get returns the details of an account notice.
func (c Client) Get(id string, params *stripe.AccountNoticeParams) (*stripe.AccountNotice, error) {
	path := stripe.FormatURLPath("/v1/account_notices/%s", id)
	accountnotice := &stripe.AccountNotice{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, accountnotice)
	return accountnotice, err
}

// Update updates an account notice's properties.
func Update(id string, params *stripe.AccountNoticeParams) (*stripe.AccountNotice, error) {
	return getC().Update(id, params)
}

// Update updates an account notice's properties.
func (c Client) Update(id string, params *stripe.AccountNoticeParams) (*stripe.AccountNotice, error) {
	path := stripe.FormatURLPath("/v1/account_notices/%s", id)
	accountnotice := &stripe.AccountNotice{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, accountnotice)
	return accountnotice, err
}

// List returns a list of account notices.
func List(params *stripe.AccountNoticeListParams) *Iter {
	return getC().List(params)
}

// List returns a list of account notices.
func (c Client) List(listParams *stripe.AccountNoticeListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.AccountNoticeList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/account_notices", c.Key, b, p, list)

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
