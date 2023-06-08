//
//
// File generated from our OpenAPI spec
//
//

package stripe

type AcceptedFinancingFinancingType string

// List of values that AcceptedFinancingFinancingType can take
const (
	AcceptedFinancingFinancingTypeCashAdvance AcceptedFinancingFinancingType = "cash_advance"
	AcceptedFinancingFinancingTypeFlexLoan    AcceptedFinancingFinancingType = "flex_loan"
)

// Returns the Accepted Financing offers for a merchant
type AcceptedFinancingParams struct {
	ListParams `form:"*"`
}

// This is an object representing financing that has been accepted by a merchant.
type AcceptedFinancing struct {
	AcceptedAdvanceAmount int64                          `json:"accepted_advance_amount"`
	AcceptedAt            int64                          `json:"accepted_at"`
	AcceptedPremiumAmount int64                          `json:"accepted_premium_amount"`
	AcceptedWithholdRate  float64                        `json:"accepted_withhold_rate"`
	Currency              Currency                       `json:"currency"`
	FinancingType         AcceptedFinancingFinancingType `json:"financing_type"`
	ID                    string                         `json:"id"`
	IsRefill              bool                           `json:"is_refill"`
	Merchant              string                         `json:"merchant"`
	Object                string                         `json:"object"`
	State                 string                         `json:"state"`
}

// AcceptedFinancingList is a list of AcceptedFinancings as retrieved from a list endpoint.
type AcceptedFinancingList struct {
	APIResource
	ListMeta
	Data []*AcceptedFinancing `json:"data"`
}
