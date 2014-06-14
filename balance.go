package stripe

import (
	"net/url"
	"strconv"
)

type Balance struct {
	Live      bool     `json:"livemode"`
	Available []Amount `json:"available"`
	Pending   []Amount `json:"pending"`
}

type TransactionStatus string
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

type TxListParams struct {
	Created, Available                  int64
	Filters                             Filters
	Currency, Src, Transfer, Start, End string
	Type                                TransactionType
	Limit                               uint64
}

type Transaction struct {
	Id         string            `json:"id"`
	Amount     int64             `json:"amount"`
	Currency   Currency          `json:"currency"`
	Available  int64             `json:"available_on"`
	Created    int64             `json:"created"`
	Fee        uint64            `json:"fee"`
	FeeDetails []Fee             `json:"fee_details"`
	Net        int64             `json:"net"`
	Status     TransactionStatus `json:"status"`
	Type       TransactionType   `json:"type"`
	Desc       string            `json:"description"`
	Src        string            `json:"source"`
	Recipient  string            `json:"recipient"`
}

type Amount struct {
	Value    int64    `json:"amount"`
	Currency Currency `json:"currency"`
}

type Fee struct {
	Amount      int64    `json:"amount"`
	Currency    Currency `json:"currency"`
	Desc        string   `json:"description"`
	Application string   `json:"application"`
}

type TransactionList struct {
	Count  uint16         `json:"total_count"`
	More   bool           `json:"has_more"`
	Url    string         `json:"url"`
	Values []*Transaction `json:"data"`
}

type BalanceClient struct {
	api   Api
	token string
}

func (c *BalanceClient) Get() (*Balance, error) {
	balance := &Balance{}
	err := c.api.Call("GET", "/balance", c.token, nil, balance)

	return balance, err
}

func (c *BalanceClient) GetTx(id string) (*Transaction, error) {
	balance := &Transaction{}
	err := c.api.Call("GET", "/balance/history/"+id, c.token, nil, balance)

	return balance, err
}

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
