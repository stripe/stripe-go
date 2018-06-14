package account

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /account APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new account.
func New(params *stripe.AccountParams) (*stripe.Account, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.AccountParams) (*stripe.Account, error) {
	acct := &stripe.Account{}
	err := c.B.Call(http.MethodPost, "/accounts", c.Key, params, acct)
	return acct, err
}

// Get returns the details of an account.
func Get() (*stripe.Account, error) {
	return getC().Get()
}

func (c Client) Get() (*stripe.Account, error) {
	account := &stripe.Account{}
	err := c.B.Call(http.MethodGet, "/account", c.Key, nil, account)
	return account, err
}

// GetByID returns the details of your account.
func GetByID(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	return getC().GetByID(id, params)
}

func (c Client) GetByID(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	path := stripe.FormatURLPath("/accounts/%s", id)
	account := &stripe.Account{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, account)
	return account, err
}

// Update updates the details of an account.
func Update(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	path := stripe.FormatURLPath("/accounts/%s", id)
	acct := &stripe.Account{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, acct)
	return acct, err
}

// Del deletes an account
func Del(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	path := stripe.FormatURLPath("/accounts/%s", id)
	acct := &stripe.Account{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, acct)
	return acct, err
}

// Reject rejects an account
func Reject(id string, params *stripe.AccountRejectParams) (*stripe.Account, error) {
	return getC().Reject(id, params)
}

func (c Client) Reject(id string, params *stripe.AccountRejectParams) (*stripe.Account, error) {
	path := stripe.FormatURLPath("/accounts/%s/reject", id)
	acct := &stripe.Account{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, acct)
	return acct, err
}

// List lists your accounts.
func List(params *stripe.AccountListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(listParams *stripe.AccountListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.AccountList{}
		err := c.B.CallRaw(http.MethodGet, "/accounts", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Accounts.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Account returns the most recent Account
// visited by a call to Next.
func (i *Iter) Account() *stripe.Account {
	return i.Current().(*stripe.Account)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
