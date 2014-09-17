// package balance provides the /balance APIs
package balance

import (
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /balance and transaction-related APIs.
type Client struct {
	B   Backend
	Key string
}

// Get returns the details of your balance.
// For more details see https://stripe.com/docs/api#retrieve_balance.
func Get(params *BalanceParams) (*Balance, error) {
	return getC().Get(params)
}

func (c Client) Get(params *BalanceParams) (*Balance, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	balance := &Balance{}
	err := c.B.Call("GET", "/balance", c.Key, body, balance)

	return balance, err
}

// GetTx returns the details of a balance transaction.
// For more details see	https://stripe.com/docs/api#retrieve_balance_transaction.
func GetTx(id string, params *TxParams) (*Transaction, error) {
	return getC().GetTx(id, params)
}

func (c Client) GetTx(id string, params *TxParams) (*Transaction, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	balance := &Transaction{}
	err := c.B.Call("GET", "/balance/history/"+id, c.Key, body, balance)

	return balance, err
}

// List returns a list of balance transactions.
// For more details see https://stripe.com/docs/api#balance_history.
func List(params *TxListParams) *TransactionIter {
	return getC().List(params)
}

func (c Client) List(params *TxListParams) *TransactionIter {
	var body *url.Values
	var lp *ListParams

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if params.Available > 0 {
			body.Add("available_on", strconv.FormatInt(params.Available, 10))
		}

		if len(params.Currency) > 0 {
			body.Add("currency", params.Currency)
		}

		if len(params.Src) > 0 {
			body.Add("source", params.Src)
		}

		if len(params.Transfer) > 0 {
			body.Add("transfer", params.Transfer)
		}

		if len(params.Type) > 0 {
			body.Add("type", string(params.Type))
		}

		params.AppendTo(body)
		lp = &params.ListParams
	}

	return &TransactionIter{GetIter(lp, body, func(b url.Values) ([]interface{}, ListMeta, error) {
		type transactionList struct {
			ListMeta
			Values []*Transaction `json:"data"`
		}

		list := &transactionList{}
		err := c.B.Call("GET", "/balance/history", c.Key, &b, list)

		ret := make([]interface{}, len(list.Values))
		for i, v := range list.Values {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

func getC() Client {
	return Client{GetBackend(), Key}
}
