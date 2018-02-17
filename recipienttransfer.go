package stripe

import "encoding/json"

// RecipientTransferDestinationType consts represent valid recipient_transfer destinations.
type RecipientTransferDestinationType string

// RecipientTransferFailCode is the list of allowed values for the recipient_transfer's failure code.
// Allowed values are "insufficient_funds", "account_closed", "no_account",
// "invalid_account_number", "debit_not_authorized", "bank_ownership_changed",
// "account_frozen", "could_not_process", "bank_account_restricted", "invalid_currency".
type RecipientTransferFailCode string

// RecipientTransferSourceType is the list of allowed values for the recipient_transfer's source_type field.
// Allowed values are "alipay_account", bank_account", "bitcoin_receiver", "card".
type RecipientTransferSourceType string

// RecipientTransferStatus is the list of allowed values for the recipient_transfer's status.
// Allowed values are "paid", "pending", "in_transit",  "failed".
type RecipientTransferStatus string

// RecipientTransferType is the list of allowed values for the recipient_transfer's type.
// Allowed values are "bank_account" or "card".
type RecipientTransferType string

const (
	// RecipientTransferDestinationBankAccount is a constant representing a recipient_transfer destination
	// which is a bank account.
	RecipientTransferDestinationBankAccount RecipientTransferDestinationType = "bank_account"

	// RecipientTransferDestinationCard is a constant representing a recipient_transfer destination
	// which is a card.
	RecipientTransferDestinationCard RecipientTransferDestinationType = "card"
)

// RecipientTransferMethodType represents the type of recipient_transfer
type RecipientTransferMethodType string

const (
	// RecipientTransferMethodInstant is a constant representing an instant recipient_transfer
	RecipientTransferMethodInstant RecipientTransferMethodType = "instant"

	// RecipientTransferMethodStandard is a constant representing a standard recipient_transfer
	RecipientTransferMethodStandard RecipientTransferMethodType = "standard"
)

// RecipientTransferDestination describes the destination of a RecipientTransfer.
// The Type should indicate which object is fleshed out
// For more details see https://stripe.com/docs/api/go#recipient_transfer_object
type RecipientTransferDestination struct {
	BankAccount *BankAccount                     `json:"-"`
	Card        *Card                            `json:"-"`
	ID          string                           `json:"id"`
	Type        RecipientTransferDestinationType `json:"object"`
}

// RecipientTransfer is the resource representing a Stripe recipient_transfer.
// For more details see https://stripe.com/docs/api#recipient_transfers.
type RecipientTransfer struct {
	Amount              int64                        `json:"amount"`
	AmountReversed      int64                        `json:"amount_reversed"`
	BalanceTransaction  *BalanceTransaction          `json:"balance_transaction"`
	BankAccount         *BankAccount                 `json:"bank_account"`
	Card                *Card                        `json:"card"`
	Created             int64                        `json:"created"`
	Currency            Currency                     `json:"currency"`
	Date                int64                        `json:"date"`
	Description         string                       `json:"description"`
	Destination         RecipientTransferDestination `json:"destination"`
	FailureCode         RecipientTransferFailCode    `json:"failure_code"`
	FailureMessage      string                       `json:"failure_message"`
	ID                  string                       `json:"id"`
	Livemode            bool                         `json:"livemode"`
	Metadata            map[string]string            `json:"metadata"`
	Method              RecipientTransferMethodType  `json:"method"`
	Recipient           *Recipient                   `json:"recipient"`
	Reversals           *ReversalList                `json:"reversals"`
	Reversed            bool                         `json:"reversed"`
	SourceTransaction   *BalanceTransactionSource    `json:"source_transaction"`
	SourceType          RecipientTransferSourceType  `json:"source_type"`
	StatementDescriptor string                       `json:"statement_descriptor"`
	Status              RecipientTransferStatus      `json:"status"`
	Type                RecipientTransferType        `json:"type"`
}

// UnmarshalJSON handles deserialization of a RecipientTransfer.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *RecipientTransfer) UnmarshalJSON(data []byte) error {
	type transfer RecipientTransfer
	var tb transfer
	err := json.Unmarshal(data, &tb)
	if err == nil {
		*t = RecipientTransfer(tb)
	} else {
		// the id is surrounded by "\" characters, so strip them
		t.ID = string(data[1 : len(data)-1])
	}

	return nil
}

// UnmarshalJSON handles deserialization of a RecipientTransferDestination.
// This custom unmarshaling is needed because the specific
// type of destination it refers to is specified in the JSON
func (d *RecipientTransferDestination) UnmarshalJSON(data []byte) error {
	type dest RecipientTransferDestination
	var dd dest
	err := json.Unmarshal(data, &dd)
	if err == nil {
		*d = RecipientTransferDestination(dd)

		switch d.Type {
		case RecipientTransferDestinationBankAccount:
			err = json.Unmarshal(data, &d.BankAccount)
		case RecipientTransferDestinationCard:
			err = json.Unmarshal(data, &d.Card)
		}

		if err != nil {
			return err
		}
	} else {
		// the id is surrounded by "\" characters, so strip them
		d.ID = string(data[1 : len(data)-1])
	}

	return nil
}

// MarshalJSON handles serialization of a RecipientTransferDestination.
// This custom marshaling is needed because we can only send a string
// ID as a destination, even though it can be expanded to a full
// object when retrieving
func (d *RecipientTransferDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.ID)
}
