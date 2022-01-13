//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

type PayoutDestinationType string

// List of values that PayoutDestinationType can take
const (
	PayoutDestinationTypeBankAccount PayoutDestinationType = "bank_account"
	PayoutDestinationTypeCard        PayoutDestinationType = "card"
)

// Error code explaining reason for payout failure if available. See [Types of payout failures](https://stripe.com/docs/api#payout_failures) for a list of failure codes.
type PayoutFailureCode string

// List of values that PayoutFailureCode can take
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

// The method used to send this payout, which can be `standard` or `instant`. `instant` is only supported for payouts to debit cards. (See [Instant payouts for marketplaces](https://stripe.com/blog/instant-payouts-for-marketplaces) for more information.)
type PayoutMethodType string

// List of values that PayoutMethodType can take
const (
	PayoutMethodInstant  PayoutMethodType = "instant"
	PayoutMethodStandard PayoutMethodType = "standard"
)

// The source balance this payout came from. One of `card`, `fpx`, or `bank_account`.
type PayoutSourceType string

// List of values that PayoutSourceType can take
const (
	PayoutSourceTypeBankAccount PayoutSourceType = "bank_account"
	PayoutSourceTypeCard        PayoutSourceType = "card"
	PayoutSourceTypeFPX         PayoutSourceType = "fpx"
)

// Current status of the payout: `paid`, `pending`, `in_transit`, `canceled` or `failed`. A payout is `pending` until it is submitted to the bank, when it becomes `in_transit`. The status then changes to `paid` if the transaction goes through, or to `failed` or `canceled` (within 5 business days). Some failed payouts may initially show as `paid` but then change to `failed`.
type PayoutStatus string

// List of values that PayoutStatus can take
const (
	PayoutStatusCanceled  PayoutStatus = "canceled"
	PayoutStatusFailed    PayoutStatus = "failed"
	PayoutStatusInTransit PayoutStatus = "in_transit"
	PayoutStatusPaid      PayoutStatus = "paid"
	PayoutStatusPending   PayoutStatus = "pending"
)

// Can be `bank_account` or `card`.
type PayoutType string

// List of values that PayoutType can take
const (
	PayoutTypeBank PayoutType = "bank_account"
	PayoutTypeCard PayoutType = "card"
)

// Retrieves the details of an existing payout. Supply the unique payout ID from either a payout creation request or the payout list, and Stripe will return the corresponding payout information.
type PayoutParams struct {
	Params              `form:"*"`
	Amount              *int64  `form:"amount"`
	Currency            *string `form:"currency"`
	Description         *string `form:"description"`
	Destination         *string `form:"destination"`
	Method              *string `form:"method"`
	SourceType          *string `form:"source_type"`
	StatementDescriptor *string `form:"statement_descriptor"`
}

// Returns a list of existing payouts sent to third-party bank accounts or that Stripe has sent you. The payouts are returned in sorted order, with the most recently created payouts appearing first.
type PayoutListParams struct {
	ListParams       `form:"*"`
	ArrivalDate      *int64            `form:"arrival_date"`
	ArrivalDateRange *RangeQueryParams `form:"arrival_date"`
	Created          *int64            `form:"created"`
	CreatedRange     *RangeQueryParams `form:"created"`
	Destination      *string           `form:"destination"`
	Status           *string           `form:"status"`
}

// Reverses a payout by debiting the destination bank account. Only payouts for connected accounts to US bank accounts may be reversed at this time. If the payout is in the pending status, /v1/payouts/:id/cancel should be used instead.
//
// By requesting a reversal via /v1/payouts/:id/reverse, you confirm that the authorized signatory of the selected bank account has authorized the debit on the bank account and that no other authorization is required.
type PayoutReverseParams struct {
	Params `form:"*"`
}

// A `Payout` object is created when you receive funds from Stripe, or when you
// initiate a payout to either a bank account or debit card of a [connected
// Stripe account](https://stripe.com/docs/connect/bank-debit-card-payouts). You can retrieve individual payouts,
// as well as list all payouts. Payouts are made on [varying
// schedules](https://stripe.com/docs/connect/manage-payout-schedule), depending on your country and
// industry.
//
// Related guide: [Receiving Payouts](https://stripe.com/docs/payouts).
type Payout struct {
	APIResource
	Amount                    int64               `json:"amount"`
	ArrivalDate               int64               `json:"arrival_date"`
	Automatic                 bool                `json:"automatic"`
	BalanceTransaction        *BalanceTransaction `json:"balance_transaction"`
	BankAccount               *BankAccount        `json:"bank_account"`
	Card                      *Card               `json:"card"`
	Created                   int64               `json:"created"`
	Currency                  Currency            `json:"currency"`
	Description               *string             `json:"description"`
	Destination               *PayoutDestination  `json:"destination"`
	FailureBalanceTransaction *BalanceTransaction `json:"failure_balance_transaction"`
	FailureCode               PayoutFailureCode   `json:"failure_code"`
	FailureMessage            string              `json:"failure_message"`
	ID                        string              `json:"id"`
	Livemode                  bool                `json:"livemode"`
	Metadata                  map[string]string   `json:"metadata"`
	Method                    PayoutMethodType    `json:"method"`
	Object                    string              `json:"object"`
	OriginalPayout            *Payout             `json:"original_payout"`
	ReversedBy                *Payout             `json:"reversed_by"`
	SourceType                PayoutSourceType    `json:"source_type"`
	StatementDescriptor       string              `json:"statement_descriptor"`
	Status                    PayoutStatus        `json:"status"`
	Type                      PayoutType          `json:"type"`
}
type PayoutDestination struct {
	ID   string                `json:"id"`
	Type PayoutDestinationType `json:"object"`

	BankAccount *BankAccount `json:"-"`
	Card        *Card        `json:"-"`
}

// PayoutList is a list of Payouts as retrieved from a list endpoint.
type PayoutList struct {
	APIResource
	ListMeta
	Data []*Payout `json:"data"`
}

// UnmarshalJSON handles deserialization of a Payout.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *Payout) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type payout Payout
	var v payout
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = Payout(v)
	return nil
}

// UnmarshalJSON handles deserialization of a PayoutDestination.
// This custom unmarshaling is needed because the specific type of
// PayoutDestination it refers to is specified in the JSON
func (p *PayoutDestination) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type payoutDestination PayoutDestination
	var v payoutDestination
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = PayoutDestination(v)
	var err error

	switch p.Type {
	case PayoutDestinationTypeBankAccount:
		err = json.Unmarshal(data, &p.BankAccount)
	case PayoutDestinationTypeCard:
		err = json.Unmarshal(data, &p.Card)
	}
	return err
}
