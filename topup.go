//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// The status of the top-up is either `canceled`, `failed`, `pending`, `reversed`, or `succeeded`.
type TopupStatus string

// List of values that TopupStatus can take
const (
	TopupStatusCanceled  TopupStatus = "canceled"
	TopupStatusFailed    TopupStatus = "failed"
	TopupStatusPending   TopupStatus = "pending"
	TopupStatusReversed  TopupStatus = "reversed"
	TopupStatusSucceeded TopupStatus = "succeeded"
)

// Top up the balance of an account
type TopupParams struct {
	Params `form:"*"`
	// A positive integer representing how much to transfer.
	Amount *int64 `form:"amount"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency *string `form:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description *string       `form:"description"`
	Source      *SourceParams `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
	// Extra information about a top-up for the source's bank statement. Limited to 15 ASCII characters.
	StatementDescriptor *string `form:"statement_descriptor"`
	// A string that identifies this top-up as part of a group.
	TransferGroup *string `form:"transfer_group"`
}

// SetSource adds valid sources to a TopupParams object,
// returning an error for unsupported sources.
func (p *TopupParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	p.Source = source
	return err
}

// Returns a list of top-ups.
type TopupListParams struct {
	ListParams `form:"*"`
	// A positive integer representing how much to transfer.
	Amount *int64 `form:"amount"`
	// A positive integer representing how much to transfer.
	AmountRange *RangeQueryParams `form:"amount"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	Created *int64 `form:"created"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	CreatedRange *RangeQueryParams `form:"created"`
	// Only return top-ups that have the given status. One of `canceled`, `failed`, `pending` or `succeeded`.
	Status *string `form:"status"`
}

// To top up your Stripe balance, you create a top-up object. You can retrieve
// individual top-ups, as well as list all top-ups. Top-ups are identified by a
// unique, random ID.
//
// Related guide: [Topping Up your Platform Account](https://stripe.com/docs/connect/top-ups).
type Topup struct {
	APIResource
	// Amount transferred.
	Amount int64 `json:"amount"`
	// ID of the balance transaction that describes the impact of this top-up on your account balance. May not be specified depending on status of top-up.
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// An arbitrary string attached to the object. Often useful for displaying to users.
	Description string `json:"description"`
	// Date the funds are expected to arrive in your Stripe account for payouts. This factors in delays like weekends or bank holidays. May not be specified depending on status of top-up.
	ExpectedAvailabilityDate int64 `json:"expected_availability_date"`
	// Error code explaining reason for top-up failure if available (see [the errors section](https://stripe.com/docs/api#errors) for a list of codes).
	FailureCode string `json:"failure_code"`
	// Message to user further explaining reason for top-up failure if available.
	FailureMessage string `json:"failure_message"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// `Source` objects allow you to accept a variety of payment methods. They
	// represent a customer's payment instrument, and can be used with the Stripe API
	// just like a `Card` object: once chargeable, they can be charged, or can be
	// attached to customers.
	//
	// Related guides: [Sources API](https://stripe.com/docs/sources) and [Sources & Customers](https://stripe.com/docs/sources/customers).
	Source *PaymentSource `json:"source"`
	// Extra information about a top-up. This will appear on your source's bank statement. It must contain at least one letter.
	StatementDescriptor string `json:"statement_descriptor"`
	// The status of the top-up is either `canceled`, `failed`, `pending`, `reversed`, or `succeeded`.
	Status TopupStatus `json:"status"`
	// A string that identifies this top-up as part of a group.
	TransferGroup string `json:"transfer_group"`

	// The following property is deprecated
	ArrivalDate int64 `json:"arrival_date"`
}

// TopupList is a list of Topups as retrieved from a list endpoint.
type TopupList struct {
	APIResource
	ListMeta
	Data []*Topup `json:"data"`
}

// UnmarshalJSON handles deserialization of a Topup.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *Topup) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		t.ID = id
		return nil
	}

	type topup Topup
	var v topup
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*t = Topup(v)
	return nil
}
