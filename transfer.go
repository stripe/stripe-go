package stripe

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
	ID        string            `json:"id"`
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
	Statement string            `json:"statement_descriptor"`
}
