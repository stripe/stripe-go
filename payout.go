package stripe

import "encoding/json"

// PayoutStatus is the list of allowed values for the payout's status.
// Allowed values are "paid", "pending", "in_transit",  "failed", "canceled".
type PayoutStatus string

// PayoutType is the list of allowed values for the payout's type.
// Allowed values are "bank_account" or "card".
type PayoutType string

// PayoutSourceType is the list of allowed values for the payout's source_type field.
// Allowed values are "alipay_account", bank_account", "bitcoin_receiver", "card".
type PayoutSourceType string

// PayoutFailureCode is the list of allowed values for the payout's failure code.
// Allowed values are "insufficient_funds", "account_closed", "no_account",
// "invalid_account_number", "debit_not_authorized", "bank_ownership_changed",
// "account_frozen", "could_not_process", "bank_account_restricted", "invalid_currency".
type PayoutFailureCode string

// PayoutDestinationType consts represent valid payout destinations.
type PayoutDestinationType string

const (
	// PayoutDestinationBankAccount is a constant representing a payout destination
	// which is a bank account.
	PayoutDestinationBankAccount PayoutDestinationType = "bank_account"

	// PayoutDestinationCard is a constant representing a payout destination
	// which is a card.
	PayoutDestinationCard PayoutDestinationType = "card"
)

// PayoutMethodType represents the type of payout
type PayoutMethodType string

const (
	// PayoutMethodInstant is a constant representing an instant payout
	PayoutMethodInstant PayoutMethodType = "instant"

	// PayoutMethodStandard is a constant representing a standard payout
	PayoutMethodStandard PayoutMethodType = "standard"
)

// PayoutDestination describes the destination of a Payout.
// The Type should indicate which object is fleshed out
// For more details see https://stripe.com/docs/api/go#payout_object
type PayoutDestination struct {
	Type        PayoutDestinationType `json:"object"`
	ID          string                `json:"id"`
	BankAccount *BankAccount          `json:"-"`
	Card        *Card                 `json:"-"`
}

// PayoutParams is the set of parameters that can be used when creating or updating a payout.
// For more details see https://stripe.com/docs/api#create_payout and https://stripe.com/docs/api#update_payout.
type PayoutParams struct {
	Params
	Amount              int64
	Currency            Currency
	Destination         string
	Method              PayoutMethodType
	SourceType          PayoutSourceType
	StatementDescriptor string
}

// PayoutListParams is the set of parameters that can be used when listing payouts.
// For more details see https://stripe.com/docs/api#list_payouts.
type PayoutListParams struct {
	ListParams
	ArrivalDate      int64
	ArrivalDateRange *RangeQueryParams
	Created          int64
	CreatedRange     *RangeQueryParams
	Destination      string
	Status           PayoutStatus
}

// Payout is the resource representing a Stripe payout.
// For more details see https://stripe.com/docs/api#payouts.
type Payout struct {
	ID                        string            `json:"id"`
	Live                      bool              `json:"livemode"`
	Amount                    int64             `json:"amount"`
	ArrivalDate               int64             `json:"arrival_date"`
	BalanceTransaction        *Transaction      `json:"balance_transaction"`
	Bank                      *BankAccount      `json:"bank_account"`
	Card                      *Card             `json:"card"`
	Created                   int64             `json:"created"`
	Currency                  Currency          `json:"currency"`
	Destination               PayoutDestination `json:"destination"`
	FailureBalanceTransaction *Transaction      `json:"failure_balance_transaction"`
	FailCode                  PayoutFailureCode `json:"failure_code"`
	FailMessage               string            `json:"failure_message"`
	Meta                      map[string]string `json:"metadata"`
	Method                    PayoutMethodType  `json:"method"`
	SourceType                PayoutSourceType  `json:"source_type"`
	StatementDescriptor       string            `json:"statement_descriptor"`
	Status                    PayoutStatus      `json:"status"`
	Type                      PayoutType        `json:"type"`
}

// PayoutList is a list of payouts as retrieved from a list endpoint.
type PayoutList struct {
	ListMeta
	Values []*Payout `json:"data"`
}

// UnmarshalJSON handles deserialization of a Payout.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *Payout) UnmarshalJSON(data []byte) error {
	type payout Payout
	var tb payout
	err := json.Unmarshal(data, &tb)
	if err == nil {
		*t = Payout(tb)
	} else {
		// the id is surrounded by "\" characters, so strip them
		t.ID = string(data[1 : len(data)-1])
	}

	return nil
}

// UnmarshalJSON handles deserialization of a PayoutDestination.
// This custom unmarshaling is needed because the specific
// type of destination it refers to is specified in the JSON
func (d *PayoutDestination) UnmarshalJSON(data []byte) error {
	type dest PayoutDestination
	var dd dest
	err := json.Unmarshal(data, &dd)
	if err == nil {
		*d = PayoutDestination(dd)

		switch d.Type {
		case PayoutDestinationBankAccount:
			json.Unmarshal(data, &d.BankAccount)
		case PayoutDestinationCard:
			json.Unmarshal(data, &d.Card)
		}
	} else {
		// the id is surrounded by "\" characters, so strip them
		d.ID = string(data[1 : len(data)-1])
	}

	return nil
}

// MarshalJSON handles serialization of a PayoutDestination.
// This custom marshaling is needed because we can only send a string
// ID as a destination, even though it can be expanded to a full
// object when retrieving
func (d *PayoutDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.ID)
}
