//
//
// File generated from our OpenAPI spec
//
//

package stripe

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
	AsOf int64 `json:"as_of"`
	// The balances owed to (or by) the account holder.
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
