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
	PayoutFailureCodeAccountClosed                 PayoutFailureCode = "account_closed"
	PayoutFailureCodeAccountFrozen                 PayoutFailureCode = "account_frozen"
	PayoutFailureCodeBankAccountRestricted         PayoutFailureCode = "bank_account_restricted"
	PayoutFailureCodeBankOwnershipChanged          PayoutFailureCode = "bank_ownership_changed"
	PayoutFailureCodeCouldNotProcess               PayoutFailureCode = "could_not_process"
	PayoutFailureCodeDebitNotAuthorized            PayoutFailureCode = "debit_not_authorized"
	PayoutFailureCodeDeclined                      PayoutFailureCode = "declined"
	PayoutFailureCodeInsufficientFunds             PayoutFailureCode = "insufficient_funds"
	PayoutFailureCodeInvalidAccountNumber          PayoutFailureCode = "invalid_account_number"
	PayoutFailureCodeIncorrectAccountHolderName    PayoutFailureCode = "incorrect_account_holder_name"
	PayoutFailureCodeIncorrectAccountHolderAddress PayoutFailureCode = "incorrect_account_holder_address"
	PayoutFailureCodeIncorrectAccountHolderTaxID   PayoutFailureCode = "incorrect_account_holder_tax_id"
	PayoutFailureCodeInvalidCurrency               PayoutFailureCode = "invalid_currency"
	PayoutFailureCodeNoAccount                     PayoutFailureCode = "no_account"
	PayoutFailureCodeUnsupportedCard               PayoutFailureCode = "unsupported_card"
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
	Params `form:"*"`
	// A positive integer in cents representing how much to payout.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `form:"description"`
	// The ID of a bank account or a card to send the payout to. If no destination is supplied, the default external account for the specified currency will be used.
	Destination *string `form:"destination"`
	// The method used to send this payout, which can be `standard` or `instant`. `instant` is only supported for payouts to debit cards. (See [Instant payouts for marketplaces for more information](https://stripe.com/blog/instant-payouts-for-marketplaces).)
	Method *string `form:"method"`
	// The balance type of your Stripe balance to draw this payout from. Balances for different payment sources are kept separately. You can find the amounts with the balances API. One of `bank_account`, `card`, or `fpx`.
	SourceType *string `form:"source_type"`
	// A string to be displayed on the recipient's bank or card statement. This may be at most 22 characters. Attempting to use a `statement_descriptor` longer than 22 characters will return an error. Note: Most banks will truncate this information and/or display it inconsistently. Some may not display it at all.
	StatementDescriptor *string `form:"statement_descriptor"`
}

// Returns a list of existing payouts sent to third-party bank accounts or that Stripe has sent you. The payouts are returned in sorted order, with the most recently created payouts appearing first.
type PayoutListParams struct {
	ListParams       `form:"*"`
	ArrivalDate      *int64            `form:"arrival_date"`
	ArrivalDateRange *RangeQueryParams `form:"arrival_date"`
	Created          *int64            `form:"created"`
	CreatedRange     *RangeQueryParams `form:"created"`
	// The ID of an external account - only return payouts sent to this external account.
	Destination *string `form:"destination"`
	// Only return payouts that have the given status: `pending`, `paid`, `failed`, or `canceled`.
	Status *string `form:"status"`
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
	// Amount (in %s) to be transferred to your bank account or debit card.
	Amount int64 `json:"amount"`
	// Date the payout is expected to arrive in the bank. This factors in delays like weekends or bank holidays.
	ArrivalDate int64 `json:"arrival_date"`
	// Returns `true` if the payout was created by an [automated payout schedule](https://stripe.com/docs/payouts#payout-schedule), and `false` if it was [requested manually](https://stripe.com/docs/payouts#manual-payouts).
	Automatic bool `json:"automatic"`
	// ID of the balance transaction that describes the impact of this payout on your account balance.
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
	BankAccount        *BankAccount        `json:"bank_account"`
	Card               *Card               `json:"card"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string `json:"description"`
	// ID of the bank account or card the payout was sent to.
	Destination *PayoutDestination `json:"destination"`
	// If the payout failed or was canceled, this will be the ID of the balance transaction that reversed the initial balance transaction, and puts the funds from the failed payout back in your balance.
	FailureBalanceTransaction *BalanceTransaction `json:"failure_balance_transaction"`
	// Error code explaining reason for payout failure if available. See [Types of payout failures](https://stripe.com/docs/api#payout_failures) for a list of failure codes.
	FailureCode PayoutFailureCode `json:"failure_code"`
	// Message to user further explaining reason for payout failure if available.
	FailureMessage string `json:"failure_message"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The method used to send this payout, which can be `standard` or `instant`. `instant` is only supported for payouts to debit cards. (See [Instant payouts for marketplaces](https://stripe.com/blog/instant-payouts-for-marketplaces) for more information.)
	Method PayoutMethodType `json:"method"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// If the payout reverses another, this is the ID of the original payout.
	OriginalPayout *Payout `json:"original_payout"`
	// If the payout was reversed, this is the ID of the payout that reverses this payout.
	ReversedBy *Payout `json:"reversed_by"`
	// The source balance this payout came from. One of `card`, `fpx`, or `bank_account`.
	SourceType PayoutSourceType `json:"source_type"`
	// Extra information about a payout to be displayed on the user's bank statement.
	StatementDescriptor string `json:"statement_descriptor"`
	// Current status of the payout: `paid`, `pending`, `in_transit`, `canceled` or `failed`. A payout is `pending` until it is submitted to the bank, when it becomes `in_transit`. The status then changes to `paid` if the transaction goes through, or to `failed` or `canceled` (within 5 business days). Some failed payouts may initially show as `paid` but then change to `failed`.
	Status PayoutStatus `json:"status"`
	// Can be `bank_account` or `card`.
	Type PayoutType `json:"type"`
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
