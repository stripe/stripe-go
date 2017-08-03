// Package payout provides the /payouts APIs
package payout

import (
	"fmt"
	"strconv"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

const (
	Canceled stripe.PayoutStatus = "canceled"
	Failed   stripe.PayoutStatus = "failed"
	Paid     stripe.PayoutStatus = "paid"
	Pending  stripe.PayoutStatus = "pending"
	Transit  stripe.PayoutStatus = "in_transit"

	Bank stripe.PayoutType = "bank_account"
	Card stripe.PayoutType = "card"

	SourceAlipay  stripe.PayoutSourceType = "alipay_account"
	SourceBank    stripe.PayoutSourceType = "bank_account"
	SourceBitcoin stripe.PayoutSourceType = "bitcoin_receiver"
	SourceCard    stripe.PayoutSourceType = "card"

	AccountClosed        stripe.PayoutFailureCode = "account_closed"
	AccountFrozen        stripe.PayoutFailureCode = "account_frozen"
	BankAccountRestrict  stripe.PayoutFailureCode = "bank_account_restricted"
	BankOwnerChanged     stripe.PayoutFailureCode = "bank_ownership_changed"
	CouldNotProcess      stripe.PayoutFailureCode = "could_not_process"
	DebitNotAuth         stripe.PayoutFailureCode = "debit_not_authorized"
	InsufficientFunds    stripe.PayoutFailureCode = "insufficient_funds"
	InvalidAccountNumber stripe.PayoutFailureCode = "invalid_account_number"
	InvalidCurrency      stripe.PayoutFailureCode = "invalid_currency"
	NoAccount            stripe.PayoutFailureCode = "no_account"
)

// Client is used to invoke /payouts APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs a new payout.
// For more details see https://stripe.com/docs/api#create_payout.
func New(params *stripe.PayoutParams) (*stripe.Payout, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.PayoutParams) (*stripe.Payout, error) {
	body := &form.Values{}
	body.Add("amount", strconv.FormatInt(params.Amount, 10))
	body.Add("currency", string(params.Currency))

	if len(params.Destination) > 0 {
		body.Add("destination", params.Destination)
	}

	if len(params.Method) > 0 {
		body.Add("method", string(params.Method))
	}

	if len(params.SourceType) > 0 {
		body.Add("source_type", string(params.SourceType))
	}

	if len(params.StatementDescriptor) > 0 {
		body.Add("statement_descriptor", params.StatementDescriptor)
	}

	params.AppendTo(body)

	payout := &stripe.Payout{}
	err := c.B.Call("POST", "/payouts", c.Key, body, &params.Params, payout)

	return payout, err
}

// Get returns the details of a payout.
// For more details see https://stripe.com/docs/api#retrieve_payout.
func Get(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		params.AppendTo(body)
	}

	payout := &stripe.Payout{}
	err := c.B.Call("GET", "/payouts/"+id, c.Key, body, commonParams, payout)

	return payout, err
}

// Update updates a payout's properties.
// For more details see https://stripe.com/docs/api#update_payout.
func Update(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		params.AppendTo(body)
	}

	payout := &stripe.Payout{}
	err := c.B.Call("POST", "/payouts/"+id, c.Key, body, commonParams, payout)

	return payout, err
}

// Cancel cancels a pending payout.
// For more details see https://stripe.com/docs/api#cancel_payout.
func Cancel(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	return getC().Cancel(id, params)
}

func (c Client) Cancel(id string, params *stripe.PayoutParams) (*stripe.Payout, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params

		body = &form.Values{}
		params.AppendTo(body)
	}

	payout := &stripe.Payout{}
	err := c.B.Call("POST", fmt.Sprintf("/payouts/%v/cancel", id), c.Key, body, commonParams, payout)

	return payout, err
}

// List returns a list of payouts.
// For more details see https://stripe.com/docs/api#list_payouts.
func List(params *stripe.PayoutListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.PayoutListParams) *Iter {
	var body *form.Values
	var lp *stripe.ListParams
	var p *stripe.Params

	if params != nil {
		body = &form.Values{}

		if params.ArrivalDate > 0 {
			body.Add("created", strconv.FormatInt(params.ArrivalDate, 10))
		}

		if params.ArrivalDateRange != nil {
			params.ArrivalDateRange.AppendTo(body, "arrival_date")
		}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if params.CreatedRange != nil {
			params.CreatedRange.AppendTo(body, "created")
		}

		if len(params.Destination) > 0 {
			body.Add("destination", params.Destination)
		}

		if len(params.Status) > 0 {
			body.Add("status", string(params.Status))
		}

		params.AppendTo(body)
		lp = &params.ListParams
		p = params.ToParams()
	}

	return &Iter{stripe.GetIter(lp, body, func(b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.PayoutList{}
		err := c.B.Call("GET", "/payouts", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Payouts.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Payout returns the most recent Payout
// visited by a call to Next.
func (i *Iter) Payout() *stripe.Payout {
	return i.Current().(*stripe.Payout)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
