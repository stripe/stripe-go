// Package balance provides the /balance APIs
package balance

import (
	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

const (
	TxAvailable stripe.TransactionStatus = "available"
	TxPending   stripe.TransactionStatus = "pending"

	TxCharge                  stripe.TransactionType = "charge"
	TxRefund                  stripe.TransactionType = "refund"
	TxAdjust                  stripe.TransactionType = "adjustment"
	TxAppFee                  stripe.TransactionType = "application_fee"
	TxFeeRefund               stripe.TransactionType = "application_fee_refund"
	TxRecipientTransfer       stripe.TransactionType = "recipient_transfer"
	TxRecipientTransferCancel stripe.TransactionType = "recipient_transfer_cancel"
	TxRecipientTransferFail   stripe.TransactionType = "recipient_transfer_failure"
	TxPayout                  stripe.TransactionType = "payout"
	TxPayoutCancel            stripe.TransactionType = "payout_cancel"
	TxPayoutFail              stripe.TransactionType = "payout_failure"
	TxTransfer                stripe.TransactionType = "transfer"
	TxTransferCancel          stripe.TransactionType = "transfer_refund"
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

// GetTx returns the details of a balance transaction.
// For more details see	https://stripe.com/docs/api#retrieve_balance_transaction.
func GetTx(id string, params *stripe.TxParams) (*stripe.Transaction, error) {
	return getC().GetTx(id, params)
}

func (c Client) GetTx(id string, params *stripe.TxParams) (*stripe.Transaction, error) {
	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	balance := &stripe.Transaction{}
	err := c.B.Call("GET", "/balance/history/"+id, c.Key, body, commonParams, balance)

	return balance, err
}

// List returns a list of balance transactions.
// For more details see https://stripe.com/docs/api#balance_history.
func List(params *stripe.TxListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.TxListParams) *Iter {
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
		list := &stripe.TransactionList{}
		err := c.B.Call("GET", "/balance/history", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Transactions.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Charge returns the most recent Transaction
// visited by a call to Next.
func (i *Iter) Transaction() *stripe.Transaction {
	return i.Current().(*stripe.Transaction)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
