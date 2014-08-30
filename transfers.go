package stripe

import (
	"fmt"
	"net/url"
	"strconv"
)

// TransferStatus is the list of allowed values for the transfer's status.
// Allowed values are "paid", "pending", "failed", "canceled".
type TransferStatus string

// TransferType is the list of allowed values for the transfer's type.
// Allowed values are "card", "bank_account".
type TransferType string

// TransferFailCode is the list of allowed values for the transfer's failure code.
// Allowed values are "insufficient_funds", "account_closed", "no_account",
// "invalid_account_number", "debit_not_authorized", "bank_ownership_changed",
// "account_frozen", "could_not_process", "bank_account_restricted", "invalid_currency".
type TransferFailCode string

const (
	Paid             TransferStatus = "paid"
	Pending          TransferStatus = "pending"
	Failed           TransferStatus = "failed"
	TransferCanceled TransferStatus = "canceled"

	CardTransfer TransferType = "card"
	BankTransfer TransferType = "bank_account"

	InsufficientFunds    TransferFailCode = "insufficient_funds"
	AccountClosed        TransferFailCode = "account_closed"
	NoAccount            TransferFailCode = "no_account"
	InvalidAccountNumber TransferFailCode = "invalid_account_number"
	DebitNotAuth         TransferFailCode = "debit_not_authorized"
	BankOwnerChanged     TransferFailCode = "bank_ownership_changed"
	AccountFrozen        TransferFailCode = "account_frozen"
	CouldNotProcess      TransferFailCode = "could_not_process"
	BankAccountRestrict  TransferFailCode = "bank_account_restricted"
	InvalidCurrency      TransferFailCode = "invalid_currency"
)

// TransferParams is the set of parameters that can be used when creating or updating a transfer.
// For more details see https://stripe.com/docs/api#create_transfer and https://stripe.com/docs/api#update_transfer.
type TransferParams struct {
	Params
	Amount                      int64
	Currency                    Currency
	Recipient                   string
	Desc, Statement, Bank, Card string
}

// TransferListParams is the set of parameters that can be used when listing transfers.
// For more details see https://stripe.com/docs/api#list_transfers.
type TransferListParams struct {
	ListParams
	Created, Date int64
	Recipient     string
	Status        TransferStatus
}

// Transfer is the resource representing a Stripe transfer.
// For more details see https://stripe.com/docs/api#transfers.
type Transfer struct {
	Id        string            `json:"id"`
	Live      bool              `json:"livemode"`
	Amount    int64             `json:"amount"`
	Currency  Currency          `json:"currency"`
	Created   int64             `json:"created"`
	Date      int64             `json:"date"`
	Desc      string            `json:"description"`
	FailCode  TransferFailCode  `json:"failure_code"`
	FailMsg   string            `json:"failure_message"`
	Status    TransferStatus    `json:"status"`
	Type      TransferType      `json:"type"`
	Tx        *Transaction      `json:"balance_transaction"`
	Meta      map[string]string `json:"metadata"`
	Bank      *BankAccount      `json:"bank_account"`
	Card      *Card             `json:"card"`
	Recipient *Recipient        `json:"recipient"`
	Statement string            `json:"statement_description"`
}

// TransferList is a list object for transfers.
type TransferList struct {
	ListResponse
	Values []*Transfer `json:"data"`
}

// TransferClient is the client used to invoke /transfers APIs.
type TransferClient struct {
	api   Api
	token string
}

// Create POSTs a new transfer.
// For more details see https://stripe.com/docs/api#create_transfer.
func (c *TransferClient) Create(params *TransferParams) (*Transfer, error) {
	body := &url.Values{
		"amount":    {strconv.FormatInt(params.Amount, 10)},
		"currency":  {string(params.Currency)},
		"recipient": {params.Recipient},
	}

	if len(params.Bank) > 0 {
		body.Add("bank_account", params.Bank)
	} else if len(params.Card) > 0 {
		body.Add("card", params.Card)
	}

	if len(params.Desc) > 0 {
		body.Add("description", params.Desc)
	}

	if len(params.Statement) > 0 {
		body.Add("statement_description", params.Statement)
	}

	for k, v := range params.Meta {
		body.Add(fmt.Sprintf("metadata[%v]", k), v)
	}

	transfer := &Transfer{}
	err := c.api.Call("POST", "/transfers", c.token, body, transfer)

	return transfer, err
}

// Get returns the details of a transfer.
// For more details see https://stripe.com/docs/api#retrieve_transfer.
func (c *TransferClient) Get(id string) (*Transfer, error) {
	transfer := &Transfer{}
	err := c.api.Call("GET", "/transfers/"+id, c.token, nil, transfer)

	return transfer, err
}

// Update updates a transfer's properties.
// For more details see https://stripe.com/docs/api#update_transfer.
func (c *TransferClient) Update(id string, params *TransferParams) (*Transfer, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if len(params.Desc) > 0 {
			body.Add("description", params.Desc)
		}

		for k, v := range params.Meta {
			body.Add(fmt.Sprintf("metadata[%v]", k), v)
		}
	}

	transfer := &Transfer{}
	err := c.api.Call("POST", "/transfers/"+id, c.token, body, transfer)

	return transfer, err
}

// Cancel cancels a pending transfer.
// For more details see https://stripe.com/docs/api#cancel_transfer.
func (c *TransferClient) Cancel(id string) (*Transfer, error) {
	transfer := &Transfer{}
	err := c.api.Call("POST", fmt.Sprintf("/transfers/%v/cancel", id), c.token, nil, transfer)

	return transfer, err
}

// List returns a list of transfers.
// For more details see https://stripe.com/docs/api#list_transfers.
func (c *TransferClient) List(params *TransferListParams) (*TransferList, error) {
	var body *url.Values

	if params != nil {
		body = &url.Values{}

		if params.Created > 0 {
			body.Add("created", strconv.FormatInt(params.Created, 10))
		}

		if params.Date > 0 {
			body.Add("date", strconv.FormatInt(params.Date, 10))
		}

		if len(params.Filters.f) > 0 {
			params.Filters.appendTo(body)
		}

		if len(params.Recipient) > 0 {
			body.Add("recipient", params.Recipient)
		}

		if len(params.Status) > 0 {
			body.Add("status", string(params.Status))
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

	list := &TransferList{}
	err := c.api.Call("GET", "/transfers", c.token, body, list)

	return list, err
}
