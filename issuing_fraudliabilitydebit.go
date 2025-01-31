//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of Issuing FraudLiabilityDebit objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
type IssuingFraudLiabilityDebitListParams struct {
	ListParams `form:"*"`
	// Only return Issuing Fraud Liability Debits that were created during the given date interval.
	Created *int64 `form:"created"`
	// Only return Issuing Fraud Liability Debits that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *IssuingFraudLiabilityDebitListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves an Issuing FraudLiabilityDebit object.
type IssuingFraudLiabilityDebitParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *IssuingFraudLiabilityDebitParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A fraud liability debit occurs when Stripe debits a platform's account for fraud losses on Issuing transactions.
type IssuingFraudLiabilityDebit struct {
	APIResource
	// Debited amount. This is equal to the disputed amount and is given in the card's currency and in the smallest currency unit.
	Amount int64 `json:"amount"`
	// ID of the [balance transaction](https://stripe.com/docs/api/balance_transactions) associated with this debit.
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The currency of the debit. Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The ID of the linked dispute.
	Dispute string `json:"dispute"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}

// IssuingFraudLiabilityDebitList is a list of FraudLiabilityDebits as retrieved from a list endpoint.
type IssuingFraudLiabilityDebitList struct {
	APIResource
	ListMeta
	Data []*IssuingFraudLiabilityDebit `json:"data"`
}
