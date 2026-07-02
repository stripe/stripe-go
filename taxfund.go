//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Returns a list of tax funds in reverse chronological order.
type TaxFundListParams struct {
	ListParams `form:"*"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	Created *int64 `form:"created" json:"created,omitempty"`
	// A filter on the list, based on the object `created` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with a number of different query options.
	CreatedRange *RangeQueryParams `form:"created" json:"-"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TaxFundListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a tax fund object by its ID.
type TaxFundParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TaxFundParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves a tax fund object by its ID.
type TaxFundRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TaxFundRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Associated billing or tax documents for this sweep.
type TaxFundContext struct {
	// The Checkout Session associated with this sweep, if any.
	CheckoutSession string `json:"checkout_session,omitempty"`
	// The Credit Note associated with this sweep, if any.
	CreditNote string `json:"credit_note,omitempty"`
	// The Invoice associated with this sweep, if any.
	Invoice string `json:"invoice,omitempty"`
	// The PaymentIntent associated with this sweep, if any.
	PaymentIntent string `json:"payment_intent,omitempty"`
	// The Refund associated with this sweep, if any.
	Refund string `json:"refund,omitempty"`
	// The Tax Transaction associated with this sweep, if any.
	TaxTransaction string `json:"tax_transaction,omitempty"`
}

// Details about the payments balance side of the sweep.
type TaxFundDestinationPaymentsBalance struct {
	// The balance transaction on the payments balance side of the sweep.
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
}

// Details about the tax fund financial account side of the sweep.
type TaxFundDestinationTaxFundAccount struct {
	// The tax fund financial account involved in the sweep.
	FinancialAccount string `json:"financial_account,omitempty"`
	// The financial account transaction on the tax fund account side of the sweep.
	Transaction string `json:"transaction,omitempty"`
}

// Where funds moved to.
type TaxFundDestination struct {
	// Details about the payments balance side of the sweep.
	PaymentsBalance *TaxFundDestinationPaymentsBalance `json:"payments_balance,omitempty"`
	// Details about the tax fund financial account side of the sweep.
	TaxFundAccount *TaxFundDestinationTaxFundAccount `json:"tax_fund_account,omitempty"`
	// The type of account funds moved to or from. Either `payments_balance` or `tax_fund_account`.
	Type string `json:"type"`
}

// Details about the payments balance side of the sweep.
type TaxFundSourcePaymentsBalance struct {
	// The balance transaction on the payments balance side of the sweep.
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
}

// Details about the tax fund financial account side of the sweep.
type TaxFundSourceTaxFundAccount struct {
	// The tax fund financial account involved in the sweep.
	FinancialAccount string `json:"financial_account,omitempty"`
	// The financial account transaction on the tax fund account side of the sweep.
	Transaction string `json:"transaction,omitempty"`
}

// Where funds moved from.
type TaxFundSource struct {
	// Details about the payments balance side of the sweep.
	PaymentsBalance *TaxFundSourcePaymentsBalance `json:"payments_balance,omitempty"`
	// Details about the tax fund financial account side of the sweep.
	TaxFundAccount *TaxFundSourceTaxFundAccount `json:"tax_fund_account,omitempty"`
	// The type of account funds moved to or from. Either `payments_balance` or `tax_fund_account`.
	Type string `json:"type"`
}

// What caused the sweep.
type TaxFundTrigger struct {
	// The balance transaction on the payments balance that triggered the sweep.
	BalanceTransaction *BalanceTransaction `json:"balance_transaction"`
	// The type of event that triggered the sweep. Always `balance_transaction`.
	Type string `json:"type"`
}

// A TaxFund object represents a single tax float sweep event — funds moved between
// a merchant's payments balance and their tax fund financial account for Stripe Tax obligations.
type TaxFund struct {
	APIResource
	// Amount swept, in the smallest currency unit. Always positive.
	Amount int64 `json:"amount"`
	// Associated billing or tax documents for this sweep.
	Context *TaxFundContext `json:"context,omitempty"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Where funds moved to.
	Destination *TaxFundDestination `json:"destination"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Where funds moved from.
	Source *TaxFundSource `json:"source"`
	// What caused the sweep.
	Trigger *TaxFundTrigger `json:"trigger"`
}

// TaxFundList is a list of TaxFunds as retrieved from a list endpoint.
type TaxFundList struct {
	APIResource
	ListMeta
	Data []*TaxFund `json:"data"`
}
