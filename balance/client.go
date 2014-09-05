package balance

import (
	"net/url"
	"strconv"

	. "github.com/stripe/stripe-go"
)

// Client is used to invoke /balance and transaction-related APIs.
type Client struct {
	B     Backend
	Token string
}

var c *Client

// Get returns the details of your balance.
// For more details see https://stripe.com/docs/api#retrieve_balance.
func Get(params *BalanceParams) (*Balance, error) {
	refresh()
	return c.Get(params)
}

func (c *Client) Get(params *BalanceParams) (*Balance, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	balance := &Balance{}
	err := c.B.Call("GET", "/balance", c.Token, body, balance)

	return balance, err
}

// GetTx returns the details of a balance transaction.
// For more details see	https://stripe.com/docs/api#retrieve_balance_transaction.
func GetTx(id string, params *TxParams) (*Transaction, error) {
	refresh()
	return c.GetTx(id, params)
}

func (c *Client) GetTx(id string, params *TxParams) (*Transaction, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}
		params.AppendTo(body)
	}

	balance := &Transaction{}
	err := c.B.Call("GET", "/balance/history/"+id, c.Token, body, balance)

	return balance, err
}

// List returns a list of balance transactions.
// For more details see https://stripe.com/docs/api#balance_history.
func List(params *TxListParams) (*TransactionList, error) {
	refresh()
	return c.List(params)
}

func (c *Client) List(params *TxListParams) (*TransactionList, error) {
	var body *url.Values

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
	}

	list := &TransactionList{}
	err := c.B.Call("GET", "/balance/history", c.Token, body, list)

	return list, err
}

func refresh() {
	if c == nil {
		c = &Client{B: GetBackend()}
	}

	c.Token = Key
}
