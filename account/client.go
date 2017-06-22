package account

import (
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
	body := &form.Values{}

	// Type is now required on creation and not allowed on update
	// It can't be passed if you pass `from_recipient` though
	if len(params.FromRecipient) == 0 {
		body.Add("type", string(params.Type))
	}

	form.AppendTo(body, params)

	acct := &stripe.Account{}
	err := c.B.Call("POST", "/accounts", c.Key, body, &params.Params, acct)

	return acct, err
}

// Get returns the details of an account.
func Get() (*stripe.Account, error) {
	return getC().Get()
}

func (c Client) Get() (*stripe.Account, error) {
	account := &stripe.Account{}
	err := c.B.Call("GET", "/account", c.Key, nil, nil, account)

	return account, err
}

// GetByID returns the details of your account.
func GetByID(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	return getC().GetByID(id, params)
}

func (c Client) GetByID(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	account := &stripe.Account{}
	err := c.B.Call("GET", "/accounts/"+id, c.Key, body, commonParams, account)

	return account, err
}

// Update updates the details of an account.
func Update(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	acct := &stripe.Account{}
	err := c.B.Call("POST", "/accounts/"+id, c.Key, body, commonParams, acct)

	return acct, err
}

// Del deletes an account
func Del(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.AccountParams) (*stripe.Account, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		body = &form.Values{}
		form.AppendTo(body, params)
		commonParams = &params.Params
	}

	acct := &stripe.Account{}
	err := c.B.Call("DELETE", "/accounts/"+id, c.Key, body, commonParams, acct)

	return acct, err
}

// Reject rejects an account
func Reject(id string, params *stripe.AccountRejectParams) (*stripe.Account, error) {
	return getC().Reject(id, params)
}

func (c Client) Reject(id string, params *stripe.AccountRejectParams) (*stripe.Account, error) {
	body := &form.Values{}
	if len(params.Reason) > 0 {
		body.Add("reason", params.Reason)
	}
	acct := &stripe.Account{}
	err := c.B.Call("POST", "/accounts/"+id+"/reject", c.Key, body, nil, acct)

	return acct, err
}

// List lists your accounts.
func List(params *stripe.AccountListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.AccountListParams) *Iter {
	var body *form.Values
	var lp *stripe.ListParams
	var p *stripe.Params

	if params != nil {
		body = &form.Values{}

		form.AppendTo(body, params)
		lp = &params.ListParams
		p = params.ToParams()
	}

	return &Iter{stripe.GetIter(lp, body, func(b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.AccountList{}
		err := c.B.Call("GET", "/accounts", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
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
