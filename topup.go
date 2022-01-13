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
	Params              `form:"*"`
	Amount              *int64        `form:"amount"`
	Currency            *string       `form:"currency"`
	Description         *string       `form:"description"`
	Source              *SourceParams `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
	StatementDescriptor *string       `form:"statement_descriptor"`
	TransferGroup       *string       `form:"transfer_group"`
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
	ListParams   `form:"*"`
	Amount       *int64            `form:"amount"`
	AmountRange  *RangeQueryParams `form:"amount"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Status       *string           `form:"status"`
}

// To top up your Stripe balance, you create a top-up object. You can retrieve
// individual top-ups, as well as list all top-ups. Top-ups are identified by a
// unique, random ID.
//
// Related guide: [Topping Up your Platform Account](https://stripe.com/docs/connect/top-ups).
type Topup struct {
	APIResource
	Amount                   int64               `json:"amount"`
	BalanceTransaction       *BalanceTransaction `json:"balance_transaction"`
	Created                  int64               `json:"created"`
	Currency                 Currency            `json:"currency"`
	Description              string              `json:"description"`
	ExpectedAvailabilityDate int64               `json:"expected_availability_date"`
	FailureCode              string              `json:"failure_code"`
	FailureMessage           string              `json:"failure_message"`
	ID                       string              `json:"id"`
	Livemode                 bool                `json:"livemode"`
	Metadata                 map[string]string   `json:"metadata"`
	Object                   string              `json:"object"`
	Source                   *PaymentSource      `json:"source"`
	StatementDescriptor      string              `json:"statement_descriptor"`
	Status                   TopupStatus         `json:"status"`
	TransferGroup            string              `json:"transfer_group"`

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
