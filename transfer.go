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

// TransferIter is a iterator for list responses.
type TransferIter struct {
	Iter *Iter
}

// Next returns the next value in the list.
func (i *TransferIter) Next() (*Transfer, error) {
	t, err := i.Iter.Next()
	if err != nil {
		return nil, err
	}

	return t.(*Transfer), err
}

// Stop returns true if there are no more iterations to be performed.
func (i *TransferIter) Stop() bool {
	return i.Iter.Stop()
}

// Meta returns the list metadata.
func (i *TransferIter) Meta() *ListMeta {
	return i.Iter.Meta()
}
