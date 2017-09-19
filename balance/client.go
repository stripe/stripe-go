// Package balance provides the /balance APIs
package balance

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

const (
	BalanceTransactionAvailable stripe.BalanceTransactionStatus = "available"
	BalanceTransactionPending   stripe.BalanceTransactionStatus = "pending"

	BalanceTransactionCharge                  stripe.BalanceTransactionType = "charge"
	BalanceTransactionRefund                  stripe.BalanceTransactionType = "refund"
	BalanceTransactionAdjust                  stripe.BalanceTransactionType = "adjustment"
	BalanceTransactionAppFee                  stripe.BalanceTransactionType = "application_fee"
	BalanceTransactionFeeRefund               stripe.BalanceTransactionType = "application_fee_refund"
	BalanceTransactionRecipientTransfer       stripe.BalanceTransactionType = "recipient_transfer"
	BalanceTransactionRecipientTransferCancel stripe.BalanceTransactionType = "recipient_transfer_cancel"
	BalanceTransactionRecipientTransferFail   stripe.BalanceTransactionType = "recipient_transfer_failure"
	BalanceTransactionPayout                  stripe.BalanceTransactionType = "payout"
	BalanceTransactionPayoutCancel            stripe.BalanceTransactionType = "payout_cancel"
	BalanceTransactionPayoutFail              stripe.BalanceTransactionType = "payout_failure"
	BalanceTransactionTransfer                stripe.BalanceTransactionType = "transfer"
	BalanceTransactionTransferCancel          stripe.BalanceTransactionType = "transfer_refund"
)

// Client is used to invoke /balance and transaction-related APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// Get returns the details of your balance.
// For more details see https://stripe.com/docs/api#retrieve_balance.
func Get(params *stripe.BalanceParams) (*stripe.Balance, error) {
	return getC().Get(params)
}

func (c Client) Get(params *stripe.BalanceParams) (*stripe.Balance, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	balance := &stripe.Balance{}
	err := c.B.Call("GET", "/balance", c.Key, body, commonParams, balance)

	return balance, err
}

// GetBalanceTransaction returns the details of a balance transaction.
// For more details see	https://stripe.com/docs/api#retrieve_balance_transaction.
func GetBalanceTransaction(id string, params *stripe.BalanceTransactionParams) (*stripe.BalanceTransaction, error) {
	return getC().GetBalanceTransaction(id, params)
}

func (c Client) GetBalanceTransaction(id string, params *stripe.BalanceTransactionParams) (*stripe.BalanceTransaction, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	balance := &stripe.BalanceTransaction{}
	err := c.B.Call("GET", "/balance/history/"+id, c.Key, body, commonParams, balance)

	return balance, err
}

// List returns a list of balance transactions.
// For more details see https://stripe.com/docs/api#balance_history.
func List(params *stripe.BalanceTransactionListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.BalanceTransactionListParams) *Iter {
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
		list := &stripe.BalanceTransactionList{}
		err := c.B.Call("GET", "/balance/history", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of BalanceTransactions.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Charge returns the most recent BalanceTransaction
// visited by a call to Next.
func (i *Iter) BalanceTransaction() *stripe.BalanceTransaction {
	return i.Current().(*stripe.BalanceTransaction)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
