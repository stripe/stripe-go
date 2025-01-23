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

// Lists the recorded inferred balances for a Financial Connections Account.
type FinancialConnectionsAccountInferredBalanceListParams struct {
	ListParams `form:"*"`
	Account    *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *FinancialConnectionsAccountInferredBalanceListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A historical balance for the account on a particular day. It may be sourced from a balance snapshot provided by a financial institution, or inferred using transactions data.
type FinancialConnectionsAccountInferredBalance struct {
	// The time for which this balance was calculated, measured in seconds since the Unix epoch. If the balance was computed by Stripe and not provided directly by a financial institution, it will always be 23:59:59 UTC.
	AsOf time.Time `json:"as_of"`
	// The balances owed to (or by) the account holder, before subtracting any outbound pending transactions or adding any inbound pending transactions.
	//
	// Each key is a three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase.
	//
	// Each value is a integer amount. A positive amount indicates money owed to the account holder. A negative amount indicates money owed by the account holder.
	Current map[string]int64 `json:"current"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}

// FinancialConnectionsAccountInferredBalanceList is a list of AccountInferredBalances as retrieved from a list endpoint.
type FinancialConnectionsAccountInferredBalanceList struct {
	APIResource
	ListMeta
	Data []*FinancialConnectionsAccountInferredBalance `json:"data"`
}

// UnmarshalJSON handles deserialization of a FinancialConnectionsAccountInferredBalance.
// This custom unmarshaling is needed to handle the time fields correctly.
func (f *FinancialConnectionsAccountInferredBalance) UnmarshalJSON(data []byte) error {
	type financialConnectionsAccountInferredBalance FinancialConnectionsAccountInferredBalance
	v := struct {
		AsOf int64 `json:"as_of"`
		*financialConnectionsAccountInferredBalance
	}{
		financialConnectionsAccountInferredBalance: (*financialConnectionsAccountInferredBalance)(f),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	f.AsOf = time.Unix(v.AsOf, 0)
	return nil
}

// MarshalJSON handles serialization of a FinancialConnectionsAccountInferredBalance.
// This custom marshaling is needed to handle the time fields correctly.
func (f FinancialConnectionsAccountInferredBalance) MarshalJSON() ([]byte, error) {
	type financialConnectionsAccountInferredBalance FinancialConnectionsAccountInferredBalance
	v := struct {
		AsOf int64 `json:"as_of"`
		financialConnectionsAccountInferredBalance
	}{
		financialConnectionsAccountInferredBalance: (financialConnectionsAccountInferredBalance)(f),
		AsOf: f.AsOf.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}
