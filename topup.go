//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// TopupStatus is a status of a Topup.
type TopupStatus string

// List of values that TopupStatus can take.
const (
	TopupStatusCanceled  TopupStatus = "canceled"
	TopupStatusFailed    TopupStatus = "failed"
	TopupStatusPending   TopupStatus = "pending"
	TopupStatusReversed  TopupStatus = "reversed"
	TopupStatusSucceeded TopupStatus = "succeeded"
)

// TopupParams is the set of parameters that can be used when creating or updating a top-up.
// For more details see https://stripe.com/docs/api#create_topup and https://stripe.com/docs/api#update_topup.
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

// TopupListParams is the set of parameters that can be used when listing top-ups.
// For more details see https://stripe.com/docs/api#list_topups.
type TopupListParams struct {
	ListParams   `form:"*"`
	Amount       *int64            `form:"amount"`
	AmountRange  *RangeQueryParams `form:"amount"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Status       *string           `form:"status"`
}

// Topup is the resource representing a Stripe top-up.
// For more details see https://stripe.com/docs/api#topups.
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

// TopupList is a list of top-ups as retrieved from a list endpoint.
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
