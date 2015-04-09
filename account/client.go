// Package account provides the /account APIs
package account

import (
	"net/url"
	"strconv"

	stripe "github.com/stripe/stripe-go"
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
	body := &url.Values{
		"managed": {strconv.FormatBool(params.Managed)},
	}

	if len(params.Country) > 0 {
		body.Add("country", params.Country)
	}

	if len(params.Email) > 0 {
		body.Add("email", params.Email)
	}

	if len(params.DefaultCurrency) > 0 {
		body.Add("default_currency", params.DefaultCurrency)
	}

	if len(params.Statement) > 0 {
		body.Add("statement_descriptor", params.Statement)
	}

	if len(params.BusinessName) > 0 {
		body.Add("business_name", params.BusinessName)
	}

	if len(params.SupportPhone) > 0 {
		body.Add("support_phone", params.SupportPhone)
	}

	if params.LegalEntity != nil {
		params.LegalEntity.AppendDetails(body)
	}

	if params.TransferSchedule != nil {
		params.TransferSchedule.AppendDetails(body)
	}

	if params.BankAccount != nil {
		params.BankAccount.AppendDetails(body)
	}

	params.AppendTo(body)

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
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &url.Values{}
		params.AppendTo(body)
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
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &url.Values{}

		if len(params.Email) > 0 {
			body.Add("email", params.Email)
		}

		if len(params.DefaultCurrency) > 0 {
			body.Add("default_currency", params.DefaultCurrency)
		}

		if len(params.Statement) > 0 {
			body.Add("statement_descriptor", params.Statement)
		}

		if len(params.BusinessName) > 0 {
			body.Add("business_name", params.BusinessName)
		}

		if len(params.SupportPhone) > 0 {
			body.Add("support_phone", params.SupportPhone)
		}

		if params.LegalEntity != nil {
			params.LegalEntity.AppendDetails(body)
		}

		if params.TransferSchedule != nil {
			params.TransferSchedule.AppendDetails(body)
		}

		if params.BankAccount != nil {
			params.BankAccount.AppendDetails(body)
		}

		params.AppendTo(body)
	}

	acct := &stripe.Account{}
	err := c.B.Call("POST", "/accounts/"+id, c.Key, body, commonParams, acct)

	return acct, err
}

// List lists your accounts.
func List(params *stripe.AccountListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.AccountListParams) *Iter {
	type accountList struct {
		stripe.ListMeta
		Values []*stripe.Account `json:"data"`
	}

	var body *url.Values
	var lp *stripe.ListParams

	if params != nil {
		body = &url.Values{}

		params.AppendTo(body)
		lp = &params.ListParams
	}

	return &Iter{stripe.GetIter(lp, body, func(b url.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &accountList{}
		err := c.B.Call("GET", "/accounts", c.Key, &b, nil, list)

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
