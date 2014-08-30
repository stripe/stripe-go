package stripe

import (
	"encoding/json"
	"net/url"
	"strconv"
)

// TransactionStatus is the list of allowed values for the transaction's status.
// Allowed values are "available", "pending".
type TransactionStatus string

// TransactionType is the list of allowed values for the transaction's type.
// Allowed values are "charge", "refund", "adjustment", "application_fee",
// "application_fee_refund", "transfer", "transfer_cancel", "transfer_failure".
type TransactionType string

const (
	TxAvailable TransactionStatus = "available"
	TxPending   TransactionStatus = "pending"

	TxCharge         TransactionType = "charge"
	TxRefund         TransactionType = "refund"
	TxAdjust         TransactionType = "adjustment"
	TxFee            TransactionType = "application_fee"
	TxFeeRefund      TransactionType = "application_fee_refund"
	TxTransfer       TransactionType = "transfer"
	TxTransferCancel TransactionType = "transfer_cancel"
	TxTransferFail   TransactionType = "transfer_failure"
)

// TxListParams is the set of parameters that can be used when listing balance transactions.
// For more details see https://stripe.com/docs/api/#balance_history.
type TxListParams struct {
	ListParams
	Created, Available      int64
	Currency, Src, Transfer string
	Type                    TransactionType
}

// Balance is the resource representing your Stripe balance.
// For more details see https://stripe.com/docs/api/#balance.
type Balance struct {
	// Live indicates the live mode.
	Live      bool     `json:"livemode"`
	Available []Amount `json:"available"`
	Pending   []Amount `json:"pending"`
}

// Transaction is the resource representing the balance transaction.
// For more details see https://stripe.com/docs/api/#balance.
type Transaction struct {
	Id         string            `json:"id"`
	Amount     int64             `json:"amount"`
	Currency   Currency          `json:"currency"`
	Available  int64             `json:"available_on"`
	Created    int64             `json:"created"`
	Fee        int64             `json:"fee"`
	FeeDetails []Fee             `json:"fee_details"`
	Net        int64             `json:"net"`
	Status     TransactionStatus `json:"status"`
	Type       TransactionType   `json:"type"`
	Desc       string            `json:"description"`
	Src        string            `json:"source"`
	Recipient  string            `json:"recipient"`
}

// Amount is a structure wrapping an amount value and its currency.
type Amount struct {
	Value    int64    `json:"amount"`
	Currency Currency `json:"currency"`
}

// Fee is a structure that breaks down the fees in a transaction.
type Fee struct {
	Amount      int64    `json:"amount"`
	Currency    Currency `json:"currency"`
	Type        string   `json:"type"`
	Desc        string   `json:"description"`
	Application string   `json:"application"`
}

// TransactionList is a list object for transactions.
type TransactionList struct {
	ListResponse
	Values []*Transaction `json:"data"`
}

// BalaneClient is the client used to invoke /balance and transaction-related APIs.
type BalanceClient struct {
	api   Api
	token string
}

// Get returns the details of your balance.
// For more details see https://stripe.com/docs/api#retrieve_balance.
func (c *BalanceClient) Get() (*Balance, error) {
	balance := &Balance{}
	err := c.api.Call("GET", "/balance", c.token, nil, balance)

	return balance, err
}

// GetTx returns the details of a balance transaction.
// For more details see	https://stripe.com/docs/api#retrieve_balance_transaction.
func (c *BalanceClient) GetTx(id string) (*Transaction, error) {
	balance := &Transaction{}
	err := c.api.Call("GET", "/balance/history/"+id, c.token, nil, balance)

	return balance, err
}

// List returns a list of balance transactions.
// For more details see https://stripe.com/docs/api#balance_history.
func (c *BalanceClient) List(params *TxListParams) (*TransactionList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if params.Available > 0 {
			body.Add("available_on", strconv.FormatInt(params.Available, 10))
		}

		if len(params.Filters.f) > 0 {
			params.Filters.appendTo(body)
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

		if len(params.Start) > 0 {
			body.Add("starting_after", params.Start)
		}

		if len(params.End) > 0 {
			body.Add("ending_before", params.End)
		}

		if params.Limit > 0 {
			if params.Limit > 100 {
				params.Limit = 100
			}

			body.Add("limit", strconv.FormatUint(params.Limit, 10))
		}
	}

	list := &TransactionList{}
	err := c.api.Call("GET", "/balance/history", c.token, body, list)

	return list, err
}

func (t *Transaction) UnmarshalJSON(data []byte) error {
	type transaction Transaction
	var tt transaction
	err := json.Unmarshal(data, &tt)
	if err == nil {
		*t = Transaction(tt)
	} else {
		// the id is surrounded by escaped \, so ignore those
		t.Id = string(data[1 : len(data)-1])
	}

	return nil
}
