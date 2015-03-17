// Package bankaccount provides the /bank_accounts APIs
package bankaccount

import (
	"fmt"
	"net/url"
	"strconv"

	stripe "github.com/stripe/stripe-go"
)

// Client is used to invoke /bank_accounts APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

const (
	NewAccount       stripe.BankAccountStatus = "new"
	VerifiedAccount  stripe.BankAccountStatus = "verified"
	ValidatedAccount stripe.BankAccountStatus = "validated"
	ErroredAccount   stripe.BankAccountStatus = "errored"
)

// New POSTs a new bank account.
func New(params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	body := &url.Values{
		"country":        {params.Country},
		"account_number": {params.Account},
		"routing_number": {params.Routing},
		"currency":       {params.Currency},
	}

	if params.Default {
		body.Add("default_for_currency", strconv.FormatBool(params.Default))
	}
	params.AppendTo(body)

	ba := &stripe.BankAccount{}
	err := c.B.Call("POST", fmt.Sprintf("/accounts/%v/bank_accounts", params.AccountID), c.Key, body, &params.Params, ba)

	return ba, err
}

// Get returns the details of a bank account.
func Get(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &url.Values{}
		params.AppendTo(body)
	}

	ba := &stripe.BankAccount{}
	err := c.B.Call("GET", fmt.Sprintf("/accounts/%v/bank_accounts/%v", params.AccountID, id), c.Key, body, commonParams, ba)

	return ba, err
}

// Update updates a bank account.
func Update(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	var body *url.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &url.Values{}

		if params.Default {
			body.Add("default_for_currency", strconv.FormatBool(params.Default))
		}

		params.AppendTo(body)
	}

	ba := &stripe.BankAccount{}
	err := c.B.Call("POST", fmt.Sprintf("/accounts/%v/bank_accounts/%v", params.AccountID, id), c.Key, body, commonParams, ba)

	return ba, err
}

// Del removes a bank account.
func Del(id string, params *stripe.BankAccountParams) error {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.BankAccountParams) error {
	return c.B.Call("DELETE", fmt.Sprintf("/accounts/%v/bank_accounts/%v", params.AccountID, id), c.Key, nil, nil, nil)
}

// List returns a list of bank accounts.
func List(params *stripe.BankAccountListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.BankAccountListParams) *Iter {
	body := &url.Values{}
	var lp *stripe.ListParams

	params.AppendTo(body)
	lp = &params.ListParams

	return &Iter{stripe.GetIter(lp, body, func(b url.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.BankAccountList{}
		err := c.B.Call("GET", fmt.Sprintf("/accounts/%v/bank_accounts", params.AccountID), c.Key, &b, nil, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of BankAccount.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// BankAccount returns the most recent BankAccount
// visited by a call to Next.
func (i *Iter) BankAccount() *stripe.BankAccount {
	return i.Current().(*stripe.BankAccount)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
