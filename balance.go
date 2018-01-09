package stripe

import "encoding/json"

// TransactionStatus is the list of allowed values for the transaction's status.
// Allowed values are "available", "pending".
type TransactionStatus string

// TransactionType is the list of allowed values for the transaction's type.
// Allowed values are "charge", "refund", "adjustment", "application_fee",
// "application_fee_refund", "transfer", "transfer_cancel", "transfer_failure".
type TransactionType string

// TransactionSourceType consts represent valid balance transaction sources.
type TransactionSourceType string

const (
	// TransactionSourceCharge is a constant representing a transaction source of charge
	TransactionSourceCharge TransactionSourceType = "charge"

	// TransactionSourceDispute is a constant representing a transaction source of dispute
	TransactionSourceDispute TransactionSourceType = "dispute"

	// TransactionSourceFee is a constant representing a transaction source of application_fee
	TransactionSourceFee TransactionSourceType = "application_fee"

	// TransactionSourcePayout is a constant representing a transaction source of payout
	TransactionSourcePayout TransactionSourceType = "payout"

	// TransactionSourceRecipientTransfer is a constant representing a transaction source of recipient_transfer
	TransactionSourceRecipientTransfer TransactionSourceType = "recipient_transfer"

	// TransactionSourceRefund is a constant representing a transaction source of refund
	TransactionSourceRefund TransactionSourceType = "refund"

	// TransactionSourceReversal is a constant representing a transaction source of reversal
	TransactionSourceReversal TransactionSourceType = "reversal"

	// TransactionSourceTransfer is a constant representing a transaction source of transfer
	TransactionSourceTransfer TransactionSourceType = "transfer"
)

// TransactionSource describes the source of a balance Transaction.
// The Type should indicate which object is fleshed out.
// For more details see https://stripe.com/docs/api#retrieve_balance_transaction
type TransactionSource struct {
	Charge            *Charge               `json:"-"`
	Dispute           *Dispute              `json:"-"`
	Fee               *Fee                  `json:"-"`
	ID                string                `json:"id"`
	Payout            *Payout               `json:"-"`
	RecipientTransfer *RecipientTransfer    `json:"-"`
	Refund            *Refund               `json:"-"`
	Reversal          *Reversal             `json:"-"`
	Transfer          *Transfer             `json:"-"`
	Type              TransactionSourceType `json:"object"`
}

// BalanceParams is the set of parameters that can be used when retrieving a balance.
// For more details see https://stripe.com/docs/api#balance.
type BalanceParams struct {
	Params `form:"*"`
}

// TxParams is the set of parameters that can be used when retrieving a transaction.
// For more details see https://stripe.com/docs/api#retrieve_balance_transaction.
type TxParams struct {
	Params `form:"*"`
}

// TxListParams is the set of parameters that can be used when listing balance transactions.
// For more details see https://stripe.com/docs/api/#balance_history.
type TxListParams struct {
	ListParams     `form:"*"`
	Available      int64             `form:"available_on"`
	AvailableRange *RangeQueryParams `form:"available_on"`
	Created        int64             `form:"created"`
	CreatedRange   *RangeQueryParams `form:"created"`
	Currency       string            `form:"currency"`
	Payout         string            `form:"payout"`
	Src            string            `form:"source"`
	Type           TransactionType   `form:"type"`
}

// Balance is the resource representing your Stripe balance.
// For more details see https://stripe.com/docs/api/#balance.
type Balance struct {
	Available []Amount `json:"available"`
	Live      bool     `json:"livemode"`
	Pending   []Amount `json:"pending"`
}

// Transaction is the resource representing the balance transaction.
// For more details see https://stripe.com/docs/api/#balance.
type Transaction struct {
	Amount     int64             `json:"amount"`
	Available  int64             `json:"available_on"`
	Created    int64             `json:"created"`
	Currency   Currency          `json:"currency"`
	Desc       string            `json:"description"`
	ID         string            `json:"id"`
	Fee        int64             `json:"fee"`
	FeeDetails []TxFee           `json:"fee_details"`
	Net        int64             `json:"net"`
	Recipient  string            `json:"recipient"`
	Src        TransactionSource `json:"source"`
	Status     TransactionStatus `json:"status"`
	Type       TransactionType   `json:"type"`
}

// TransactionList is a list of transactions as returned from a list endpoint.
type TransactionList struct {
	ListMeta
	Values []*Transaction `json:"data"`
}

// Amount is a structure wrapping an amount value and its currency.
type Amount struct {
	Value    int64    `json:"amount"`
	Currency Currency `json:"currency"`
}

// TxFee is a structure that breaks down the fees in a transaction.
type TxFee struct {
	Application string   `json:"application"`
	Amount      int64    `json:"amount"`
	Currency    Currency `json:"currency"`
	Desc        string   `json:"description"`
	Type        string   `json:"type"`
}

// UnmarshalJSON handles deserialization of a Transaction.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *Transaction) UnmarshalJSON(data []byte) error {
	type transaction Transaction
	var tt transaction
	err := json.Unmarshal(data, &tt)
	if err == nil {
		*t = Transaction(tt)
	} else {
		// the id is surrounded by "\" characters, so strip them
		t.ID = string(data[1 : len(data)-1])
	}

	return nil
}

// UnmarshalJSON handles deserialization of a TransactionSource.
// This custom unmarshaling is needed because the specific
// type of transaction source it refers to is specified in the JSON
func (s *TransactionSource) UnmarshalJSON(data []byte) error {
	type source TransactionSource
	var ss source
	err := json.Unmarshal(data, &ss)
	if err == nil {
		*s = TransactionSource(ss)

		switch s.Type {
		case TransactionSourceCharge:
			err = json.Unmarshal(data, &s.Charge)
		case TransactionSourceDispute:
			err = json.Unmarshal(data, &s.Dispute)
		case TransactionSourceFee:
			err = json.Unmarshal(data, &s.Fee)
		case TransactionSourcePayout:
			err = json.Unmarshal(data, &s.Payout)
		case TransactionSourceRecipientTransfer:
			err = json.Unmarshal(data, &s.RecipientTransfer)
		case TransactionSourceRefund:
			err = json.Unmarshal(data, &s.Refund)
		case TransactionSourceReversal:
			err = json.Unmarshal(data, &s.Reversal)
		case TransactionSourceTransfer:
			err = json.Unmarshal(data, &s.Transfer)
		}

		if err != nil {
			return err
		}
	} else {
		// the id is surrounded by "\" characters, so strip them
		s.ID = string(data[1 : len(data)-1])
	}

	return nil
}

// MarshalJSON handles serialization of a TransactionSource.
func (s *TransactionSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.ID)
}
