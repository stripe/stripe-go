package stripe

import "encoding/json"

// PayoutDestinationType consts represent valid payout destinations.
type PayoutDestinationType string

const (
	PayoutDestinationTypeBankAccount PayoutDestinationType = "bank_account"
	PayoutDestinationTypeCard        PayoutDestinationType = "card"
)

// PayoutFailureCode is the list of allowed values for the payout's failure code.
type PayoutFailureCode string

const (
	PayoutFailureCodeAccountClosed         PayoutFailureCode = "account_closed"
	PayoutFailureCodeAccountFrozen         PayoutFailureCode = "account_frozen"
	PayoutFailureCodeBankAccountRestricted PayoutFailureCode = "bank_account_restricted"
	PayoutFailureCodeBankOwnershipChanged  PayoutFailureCode = "bank_ownership_changed"
	PayoutFailureCodeCouldNotProcess       PayoutFailureCode = "could_not_process"
	PayoutFailureCodeDebitNotAuthorized    PayoutFailureCode = "debit_not_authorized"
	PayoutFailureCodeInsufficientFunds     PayoutFailureCode = "insufficient_funds"
	PayoutFailureCodeInvalidAccountNumber  PayoutFailureCode = "invalid_account_number"
	PayoutFailureCodeInvalidCurrency       PayoutFailureCode = "invalid_currency"
	PayoutFailureCodeNoAccount             PayoutFailureCode = "no_account"
)

// PayoutSourceType is the list of allowed values for the payout's source_type field.
type PayoutSourceType string

const (
	PayoutSourceTypeAlipayAccount   PayoutSourceType = "alipay_account"
	PayoutSourceTypeBankAccount     PayoutSourceType = "bank_account"
	PayoutSourceTypeBitcoinReceiver PayoutSourceType = "bitcoin_receiver"
	PayoutSourceTypeCard            PayoutSourceType = "card"
)

// PayoutStatus is the list of allowed values for the payout's status.
type PayoutStatus string

const (
	PayoutStatusCanceled  PayoutStatus = "canceled"
	PayoutStatusFailed    PayoutStatus = "failed"
	PayoutStatusInTransit PayoutStatus = "in_transit"
	PayoutStatusPaid      PayoutStatus = "paid"
	PayoutStatusPending   PayoutStatus = "pending"
)

// PayoutType is the list of allowed values for the payout's type.
type PayoutType string

const (
	PayoutTypeBank PayoutType = "bank_account"
	PayoutTypeCard PayoutType = "card"
)

// PayoutMethodType represents the type of payout
type PayoutMethodType string

const (
	PayoutMethodInstant  PayoutMethodType = "instant"
	PayoutMethodStandard PayoutMethodType = "standard"
)

// PayoutDestination describes the destination of a Payout.
// The Type should indicate which object is fleshed out
// For more details see https://stripe.com/docs/api/go#payout_object
type PayoutDestination struct {
	BankAccount *BankAccount          `json:"-"`
	Card        *Card                 `json:"-"`
	ID          string                `json:"id"`
	Type        PayoutDestinationType `json:"object"`
}

// PayoutParams is the set of parameters that can be used when creating or updating a payout.
// For more details see https://stripe.com/docs/api#create_payout and https://stripe.com/docs/api#update_payout.
type PayoutParams struct {
	Params              `form:"*"`
	Amount              *int64  `form:"amount"`
	Currency            *string `form:"currency"`
	Destination         *string `form:"destination"`
	Method              *string `form:"method"`
	SourceType          *string `form:"source_type"`
	StatementDescriptor *string `form:"statement_descriptor"`
}

// PayoutListParams is the set of parameters that can be used when listing payouts.
// For more details see https://stripe.com/docs/api#list_payouts.
type PayoutListParams struct {
	ListParams       `form:"*"`
	ArrivalDate      *int64            `form:"arrival_date"`
	ArrivalDateRange *RangeQueryParams `form:"arrival_date"`
	Created          *int64            `form:"created"`
	CreatedRange     *RangeQueryParams `form:"created"`
	Destination      *string           `form:"destination"`
	Status           *string           `form:"status"`
}

// Payout is the resource representing a Stripe payout.
// For more details see https://stripe.com/docs/api#payouts.
type Payout struct {
	Amount                    int64               `json:"amount"`
	ArrivalDate               int64               `json:"arrival_date"`
	Automatic                 bool                `json:"automatic"`
	BalanceTransaction        *BalanceTransaction `json:"balance_transaction"`
	BankAccount               *BankAccount        `json:"bank_account"`
	Card                      *Card               `json:"card"`
	Created                   int64               `json:"created"`
	Currency                  Currency            `json:"currency"`
	Destination               string              `json:"destination"`
	FailureBalanceTransaction *BalanceTransaction `json:"failure_balance_transaction"`
	FailureCode               PayoutFailureCode   `json:"failure_code"`
	FailureMessage            string              `json:"failure_message"`
	ID                        string              `json:"id"`
	Livemode                  bool                `json:"livemode"`
	Metadata                  map[string]string   `json:"metadata"`
	Method                    PayoutMethodType    `json:"method"`
	SourceType                PayoutSourceType    `json:"source_type"`
	StatementDescriptor       string              `json:"statement_descriptor"`
	Status                    PayoutStatus        `json:"status"`
	Type                      PayoutType          `json:"type"`
}

// PayoutList is a list of payouts as retrieved from a list endpoint.
type PayoutList struct {
	ListMeta
	Data []*Payout `json:"data"`
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
		case PayoutDestinationTypeBankAccount:
			err = json.Unmarshal(data, &d.BankAccount)
		case PayoutDestinationTypeCard:
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

// MarshalJSON handles serialization of a PayoutDestination.
// This custom marshaling is needed because we can only send a string
// ID as a destination, even though it can be expanded to a full
// object when retrieving
func (d *PayoutDestination) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.ID)
}
