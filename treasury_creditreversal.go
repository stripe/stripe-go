//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"time"
)

// The rails used to reverse the funds.
type TreasuryCreditReversalNetwork string

// List of values that TreasuryCreditReversalNetwork can take
const (
	TreasuryCreditReversalNetworkACH    TreasuryCreditReversalNetwork = "ach"
	TreasuryCreditReversalNetworkStripe TreasuryCreditReversalNetwork = "stripe"
)

// Status of the CreditReversal
type TreasuryCreditReversalStatus string

// List of values that TreasuryCreditReversalStatus can take
const (
	TreasuryCreditReversalStatusCanceled   TreasuryCreditReversalStatus = "canceled"
	TreasuryCreditReversalStatusPosted     TreasuryCreditReversalStatus = "posted"
	TreasuryCreditReversalStatusProcessing TreasuryCreditReversalStatus = "processing"
)

// Returns a list of CreditReversals.
type TreasuryCreditReversalListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Returns objects associated with this FinancialAccount.
	FinancialAccount *string `form:"financial_account"`
	// Only return CreditReversals for the ReceivedCredit ID.
	ReceivedCredit *string `form:"received_credit"`
	// Only return CreditReversals for a given status.
	Status *string `form:"status"`
}

// AddExpand appends a new field to expand.
func (p *TreasuryCreditReversalListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Reverses a ReceivedCredit and creates a CreditReversal object.
type TreasuryCreditReversalParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The ReceivedCredit to reverse.
	ReceivedCredit *string `form:"received_credit"`
}

// AddExpand appends a new field to expand.
func (p *TreasuryCreditReversalParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TreasuryCreditReversalParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

type TreasuryCreditReversalStatusTransitions struct {
	// Timestamp describing when the CreditReversal changed status to `posted`
	PostedAt time.Time `json:"posted_at"`
}

// You can reverse some [ReceivedCredits](https://stripe.com/docs/api#received_credits) depending on their network and source flow. Reversing a ReceivedCredit leads to the creation of a new object known as a CreditReversal.
type TreasuryCreditReversal struct {
	APIResource
	// Amount (in cents) transferred.
	Amount int64 `json:"amount"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created time.Time `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The FinancialAccount to reverse funds from.
	FinancialAccount string `json:"financial_account"`
	// A [hosted transaction receipt](https://stripe.com/docs/treasury/moving-money/regulatory-receipts) URL that is provided when money movement is considered regulated under Stripe's money transmission licenses.
	HostedRegulatoryReceiptURL string `json:"hosted_regulatory_receipt_url"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The rails used to reverse the funds.
	Network TreasuryCreditReversalNetwork `json:"network"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The ReceivedCredit being reversed.
	ReceivedCredit string `json:"received_credit"`
	// Status of the CreditReversal
	Status            TreasuryCreditReversalStatus             `json:"status"`
	StatusTransitions *TreasuryCreditReversalStatusTransitions `json:"status_transitions"`
	// The Transaction associated with this object.
	Transaction *TreasuryTransaction `json:"transaction"`
}

// TreasuryCreditReversalList is a list of CreditReversals as retrieved from a list endpoint.
type TreasuryCreditReversalList struct {
	APIResource
	ListMeta
	Data []*TreasuryCreditReversal `json:"data"`
}

// UnmarshalJSON handles deserialization of a TreasuryCreditReversalStatusTransitions.
// This custom unmarshaling is needed to handle the time fields correctly.
func (t *TreasuryCreditReversalStatusTransitions) UnmarshalJSON(data []byte) error {
	type treasuryCreditReversalStatusTransitions TreasuryCreditReversalStatusTransitions
	v := struct {
		PostedAt int64 `json:"posted_at"`
		*treasuryCreditReversalStatusTransitions
	}{
		treasuryCreditReversalStatusTransitions: (*treasuryCreditReversalStatusTransitions)(t),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.PostedAt = time.Unix(v.PostedAt, 0)
	return nil
}

// UnmarshalJSON handles deserialization of a TreasuryCreditReversal.
// This custom unmarshaling is needed to handle the time fields correctly.
func (t *TreasuryCreditReversal) UnmarshalJSON(data []byte) error {
	type treasuryCreditReversal TreasuryCreditReversal
	v := struct {
		Created int64 `json:"created"`
		*treasuryCreditReversal
	}{
		treasuryCreditReversal: (*treasuryCreditReversal)(t),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.Created = time.Unix(v.Created, 0)
	return nil
}

// MarshalJSON handles serialization of a TreasuryCreditReversalStatusTransitions.
// This custom marshaling is needed to handle the time fields correctly.
func (t TreasuryCreditReversalStatusTransitions) MarshalJSON() ([]byte, error) {
	type treasuryCreditReversalStatusTransitions TreasuryCreditReversalStatusTransitions
	v := struct {
		PostedAt int64 `json:"posted_at"`
		treasuryCreditReversalStatusTransitions
	}{
		treasuryCreditReversalStatusTransitions: (treasuryCreditReversalStatusTransitions)(t),
		PostedAt:                                t.PostedAt.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}

// MarshalJSON handles serialization of a TreasuryCreditReversal.
// This custom marshaling is needed to handle the time fields correctly.
func (t TreasuryCreditReversal) MarshalJSON() ([]byte, error) {
	type treasuryCreditReversal TreasuryCreditReversal
	v := struct {
		Created int64 `json:"created"`
		treasuryCreditReversal
	}{
		treasuryCreditReversal: (treasuryCreditReversal)(t),
		Created:                t.Created.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}
