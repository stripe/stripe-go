//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Stripe's confidence in this classification.
type FinancialConnectionsTransactionClassificationMoneyMovementConfidenceLevel string

// List of values that FinancialConnectionsTransactionClassificationMoneyMovementConfidenceLevel can take
const (
	FinancialConnectionsTransactionClassificationMoneyMovementConfidenceLevelHigh     FinancialConnectionsTransactionClassificationMoneyMovementConfidenceLevel = "high"
	FinancialConnectionsTransactionClassificationMoneyMovementConfidenceLevelLow      FinancialConnectionsTransactionClassificationMoneyMovementConfidenceLevel = "low"
	FinancialConnectionsTransactionClassificationMoneyMovementConfidenceLevelMedium   FinancialConnectionsTransactionClassificationMoneyMovementConfidenceLevel = "medium"
	FinancialConnectionsTransactionClassificationMoneyMovementConfidenceLevelVeryHigh FinancialConnectionsTransactionClassificationMoneyMovementConfidenceLevel = "very_high"
)

// Stripe's confidence in this classification.
type FinancialConnectionsTransactionClassificationPersonalFinanceConfidenceLevel string

// List of values that FinancialConnectionsTransactionClassificationPersonalFinanceConfidenceLevel can take
const (
	FinancialConnectionsTransactionClassificationPersonalFinanceConfidenceLevelHigh     FinancialConnectionsTransactionClassificationPersonalFinanceConfidenceLevel = "high"
	FinancialConnectionsTransactionClassificationPersonalFinanceConfidenceLevelLow      FinancialConnectionsTransactionClassificationPersonalFinanceConfidenceLevel = "low"
	FinancialConnectionsTransactionClassificationPersonalFinanceConfidenceLevelMedium   FinancialConnectionsTransactionClassificationPersonalFinanceConfidenceLevel = "medium"
	FinancialConnectionsTransactionClassificationPersonalFinanceConfidenceLevelVeryHigh FinancialConnectionsTransactionClassificationPersonalFinanceConfidenceLevel = "very_high"
)

// Stripe's confidence in the enriched merchant name.
type FinancialConnectionsTransactionEnrichmentsMerchantConfidenceLevel string

// List of values that FinancialConnectionsTransactionEnrichmentsMerchantConfidenceLevel can take
const (
	FinancialConnectionsTransactionEnrichmentsMerchantConfidenceLevelHigh     FinancialConnectionsTransactionEnrichmentsMerchantConfidenceLevel = "high"
	FinancialConnectionsTransactionEnrichmentsMerchantConfidenceLevelLow      FinancialConnectionsTransactionEnrichmentsMerchantConfidenceLevel = "low"
	FinancialConnectionsTransactionEnrichmentsMerchantConfidenceLevelMedium   FinancialConnectionsTransactionEnrichmentsMerchantConfidenceLevel = "medium"
	FinancialConnectionsTransactionEnrichmentsMerchantConfidenceLevelVeryHigh FinancialConnectionsTransactionEnrichmentsMerchantConfidenceLevel = "very_high"
)

// The status of the transaction.
type FinancialConnectionsTransactionStatus string

// List of values that FinancialConnectionsTransactionStatus can take
const (
	FinancialConnectionsTransactionStatusPending FinancialConnectionsTransactionStatus = "pending"
	FinancialConnectionsTransactionStatusPosted  FinancialConnectionsTransactionStatus = "posted"
	FinancialConnectionsTransactionStatusVoid    FinancialConnectionsTransactionStatus = "void"
)

// A filter on the list based on the object `transaction_refresh` field. The value can be a dictionary with the following options:
type FinancialConnectionsTransactionListTransactionRefreshParams struct {
	// Return results where the transactions were created or updated by a refresh that took place after this refresh (non-inclusive).
	After *string `form:"after" json:"after"`
}

// Returns a list of Financial Connections Transaction objects.
type FinancialConnectionsTransactionListParams struct {
	ListParams `form:"*"`
	// The ID of the Financial Connections Account whose transactions will be retrieved.
	Account *string `form:"account" json:"account"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// A filter on the list based on the object `transacted_at` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with the following options:
	TransactedAt *int64 `form:"transacted_at" json:"transacted_at,omitempty"`
	// A filter on the list based on the object `transacted_at` field. The value can be a string with an integer Unix timestamp, or it can be a dictionary with the following options:
	TransactedAtRange *RangeQueryParams `form:"transacted_at" json:"-"`
	// A filter on the list based on the object `transaction_refresh` field. The value can be a dictionary with the following options:
	TransactionRefresh *FinancialConnectionsTransactionListTransactionRefreshParams `form:"transaction_refresh" json:"transaction_refresh,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *FinancialConnectionsTransactionListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the details of a Financial Connections Transaction
type FinancialConnectionsTransactionParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *FinancialConnectionsTransactionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the details of a Financial Connections Transaction
type FinancialConnectionsTransactionRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *FinancialConnectionsTransactionRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Money movement classification labels for this transaction.
type FinancialConnectionsTransactionClassificationMoneyMovement struct {
	// Stripe's confidence in this classification.
	ConfidenceLevel FinancialConnectionsTransactionClassificationMoneyMovementConfidenceLevel `json:"confidence_level"`
	// The detailed category label for this transaction.
	DetailedLabel string `json:"detailed_label"`
	// The primary category label for this transaction.
	PrimaryLabel string `json:"primary_label"`
}

// Personal finance classification labels for this transaction.
type FinancialConnectionsTransactionClassificationPersonalFinance struct {
	// Stripe's confidence in this classification.
	ConfidenceLevel FinancialConnectionsTransactionClassificationPersonalFinanceConfidenceLevel `json:"confidence_level"`
	// The detailed category label for this transaction.
	DetailedLabel string `json:"detailed_label"`
	// The primary category label for this transaction.
	PrimaryLabel string `json:"primary_label"`
}

// Classification labels for this transaction, one entry per subscribed use case.
type FinancialConnectionsTransactionClassification struct {
	// Money movement classification labels for this transaction.
	MoneyMovement *FinancialConnectionsTransactionClassificationMoneyMovement `json:"money_movement"`
	// Personal finance classification labels for this transaction.
	PersonalFinance *FinancialConnectionsTransactionClassificationPersonalFinance `json:"personal_finance"`
	// The taxonomy type for this classification entry.
	Type string `json:"type"`
}
type FinancialConnectionsTransactionEnrichmentsMerchant struct {
	// Stripe's confidence in the enriched merchant name.
	ConfidenceLevel FinancialConnectionsTransactionEnrichmentsMerchantConfidenceLevel `json:"confidence_level"`
	// The normalized merchant name for this transaction.
	Name string `json:"name"`
}

// Enriched merchant information for this transaction.
type FinancialConnectionsTransactionEnrichments struct {
	Merchant *FinancialConnectionsTransactionEnrichmentsMerchant `json:"merchant"`
}
type FinancialConnectionsTransactionStatusTransitions struct {
	// Time at which this transaction posted. Measured in seconds since the Unix epoch.
	PostedAt int64 `json:"posted_at"`
	// Time at which this transaction was voided. Measured in seconds since the Unix epoch.
	VoidAt int64 `json:"void_at"`
}

// A Transaction represents a real transaction that affects a Financial Connections Account balance.
type FinancialConnectionsTransaction struct {
	APIResource
	// The ID of the Financial Connections Account this transaction belongs to.
	Account string `json:"account"`
	// The amount of this transaction, in cents (or local equivalent).
	Amount int64 `json:"amount"`
	// Classification labels for this transaction, one entry per subscribed use case.
	Classifications []*FinancialConnectionsTransactionClassification `json:"classifications,omitempty"`
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// The description of this transaction.
	Description string `json:"description"`
	// Enriched merchant information for this transaction.
	Enrichments *FinancialConnectionsTransactionEnrichments `json:"enrichments,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The status of the transaction.
	Status            FinancialConnectionsTransactionStatus             `json:"status"`
	StatusTransitions *FinancialConnectionsTransactionStatusTransitions `json:"status_transitions"`
	// Time at which the transaction was transacted. Measured in seconds since the Unix epoch.
	TransactedAt int64 `json:"transacted_at"`
	// The token of the transaction refresh that last updated or created this transaction.
	TransactionRefresh string `json:"transaction_refresh"`
	// Time at which the object was last updated. Measured in seconds since the Unix epoch.
	Updated int64 `json:"updated"`
}

// FinancialConnectionsTransactionList is a list of Transactions as retrieved from a list endpoint.
type FinancialConnectionsTransactionList struct {
	APIResource
	ListMeta
	Data []*FinancialConnectionsTransaction `json:"data"`
}
