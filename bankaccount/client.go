// Package bankaccount provides the /bank_accounts APIs
package bankaccount

import (
	"errors"
	"fmt"
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

	body := &stripe.RequestValues{}
	isCustomer := len(params.Customer) > 0

	var sourceType string
	if isCustomer {
		sourceType = "source"
	} else {
		sourceType = "external_account"
	}

	// Use token (if exists) or a dictionary containing a user’s bank account details.
	if len(params.Token) > 0 {
		body.Add(sourceType, params.Token)

		if params.Default {
			body.Add("default_for_currency", strconv.FormatBool(params.Default))
		}
	} else {
		body.Add(sourceType+"[object]", "bank_account")
		body.Add(sourceType+"[country]", params.Country)
		body.Add(sourceType+"[account_number]", params.Account)
		body.Add(sourceType+"[currency]", params.Currency)

		if isCustomer {
			body.Add(sourceType+"[account_holder_name]", params.AccountHolderName)
			body.Add(sourceType+"[account_holder_type]", params.AccountHolderType)
		}

		if len(params.Routing) > 0 {
			body.Add(sourceType+"[routing_number]", params.Routing)
		}

		if params.Default {
			body.Add(sourceType+"[default_for_currency]", strconv.FormatBool(params.Default))
		}
	}
	params.AppendTo(body)

	ba := &stripe.BankAccount{}
	var err error
	if isCustomer {
		err = c.B.Call("POST", fmt.Sprintf("/customers/%v/sources", params.Customer), c.Key, body, &params.Params, ba)
	} else {
		err = c.B.Call("POST", fmt.Sprintf("/accounts/%v/bank_accounts", params.AccountID), c.Key, body, &params.Params, ba)
	}

	return ba, err
}

// Get returns the details of a bank account.
func Get(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	var body *stripe.RequestValues
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &stripe.RequestValues{}
		params.AppendTo(body)
	}

	ba := &stripe.BankAccount{}
	var err error

	if len(params.Customer) > 0 {
		err = c.B.Call("GET", fmt.Sprintf("/customers/%v/bank_accounts/%v", params.AccountID, id), c.Key, body, commonParams, ba)
	} else if len(params.AccountID) > 0 {
		err = c.B.Call("GET", fmt.Sprintf("/accounts/%v/bank_accounts/%v", params.AccountID, id), c.Key, body, commonParams, ba)
	} else {
		err = errors.New("Invalid bank account params: either Customer or AccountID need to be set")
	}

	return ba, err
}

// Update updates a bank account.
func Update(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	var body *stripe.RequestValues
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &stripe.RequestValues{}

		if params.Default {
			body.Add("default_for_currency", strconv.FormatBool(params.Default))
		}

		params.AppendTo(body)
	}

	ba := &stripe.BankAccount{}
	var err error

	if len(params.Customer) > 0 {
		err = c.B.Call("POST", fmt.Sprintf("/customers/%v/bank_accounts/%v", params.Customer, id), c.Key, body, commonParams, ba)
	} else if len(params.AccountID) > 0 {
		err = c.B.Call("POST", fmt.Sprintf("/accounts/%v/bank_accounts/%v", params.AccountID, id), c.Key, body, commonParams, ba)
	} else {
		err = errors.New("Invalid bank account params: either Customer or AccountID need to be set")
	}

	return ba, err
}

// Del removes a bank account.
func Del(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.BankAccountParams) (*stripe.BankAccount, error) {
	ba := &stripe.BankAccount{}
	var err error

	if len(params.Customer) > 0 {
		err = c.B.Call("DELETE", fmt.Sprintf("/customers/%v/bank_accounts/%v", params.Customer, id), c.Key, nil, nil, ba)
	} else if len(params.AccountID) > 0 {
		err = c.B.Call("DELETE", fmt.Sprintf("/accounts/%v/bank_accounts/%v", params.AccountID, id), c.Key, nil, nil, ba)
	} else {
		err = errors.New("Invalid bank account params: either Customer or AccountID need to be set")
	}

	return ba, err
}

// List returns a list of bank accounts.
func List(params *stripe.BankAccountListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.BankAccountListParams) *Iter {
	body := &stripe.RequestValues{}
	var lp *stripe.ListParams
	var p *stripe.Params

	params.AppendTo(body)
	lp = &params.ListParams
	p = params.ToParams()

	return &Iter{stripe.GetIter(lp, body, func(b *stripe.RequestValues) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.BankAccountList{}
		var err error

		if len(params.Customer) > 0 {
			err = c.B.Call("GET", fmt.Sprintf("/customers/%v/bank_accounts", params.Customer), c.Key, b, p, list)
		} else if len(params.AccountID) > 0 {
			err = c.B.Call("GET", fmt.Sprintf("/accounts/%v/bank_accounts", params.AccountID), c.Key, b, p, list)
		} else {
			err = errors.New("Invalid bank account params: either Customer or AccountID need to be set")
		}

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
