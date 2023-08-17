//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The reason for the financing transaction (if applicable).
type CapitalFinancingTransactionDetailsReason string

// List of values that CapitalFinancingTransactionDetailsReason can take
const (
	CapitalFinancingTransactionDetailsReasonAutomaticWithholding       CapitalFinancingTransactionDetailsReason = "automatic_withholding"
	CapitalFinancingTransactionDetailsReasonAutomaticWithholdingRefund CapitalFinancingTransactionDetailsReason = "automatic_withholding_refund"
	CapitalFinancingTransactionDetailsReasonCollection                 CapitalFinancingTransactionDetailsReason = "collection"
	CapitalFinancingTransactionDetailsReasonCollectionFailure          CapitalFinancingTransactionDetailsReason = "collection_failure"
	CapitalFinancingTransactionDetailsReasonFinancingCancellation      CapitalFinancingTransactionDetailsReason = "financing_cancellation"
	CapitalFinancingTransactionDetailsReasonRefill                     CapitalFinancingTransactionDetailsReason = "refill"
	CapitalFinancingTransactionDetailsReasonRequestedByUser            CapitalFinancingTransactionDetailsReason = "requested_by_user"
	CapitalFinancingTransactionDetailsReasonUserInitiated              CapitalFinancingTransactionDetailsReason = "user_initiated"
)

// The type of the financing transaction.
type CapitalFinancingTransactionType string

// List of values that CapitalFinancingTransactionType can take
const (
	CapitalFinancingTransactionTypePayment  CapitalFinancingTransactionType = "payment"
	CapitalFinancingTransactionTypePayout   CapitalFinancingTransactionType = "payout"
	CapitalFinancingTransactionTypeReversal CapitalFinancingTransactionType = "reversal"
)

// Retrieves a financing transaction for a financing offer.
type CapitalFinancingTransactionParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *CapitalFinancingTransactionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Returns a list of financing transactions. The transactions are returned in sorted order,
// with the most recent transactions appearing first.
type CapitalFinancingTransactionListParams struct {
	ListParams `form:"*"`
	// For transactions of type `paydown` and reason `automatic_withholding` only, only returns transactions that were created as a result of this charge.
	Charge *string `form:"charge"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Returns transactions that were created that apply to this financing offer ID.
	FinancingOffer *string `form:"financing_offer"`
	// Only returns transactions that are responsible for reversing this financing transaction ID.
	ReversedTransaction *string `form:"reversed_transaction"`
	// For transactions of type `paydown` and reason `automatic_withholding` only, only returns transactions that were created as a result of this Treasury Transaction.
	TreasuryTransaction *string `form:"treasury_transaction"`
}

// AddExpand appends a new field to expand.
func (p *CapitalFinancingTransactionListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// This is an object representing a linked transaction on a Capital Financing Transaction.
type CapitalFinancingTransactionDetailsTransaction struct {
	// The linked payment ID.
	Charge string `json:"charge"`
	// The linked Treasury Financing Transaction ID.
	TreasuryTransaction string `json:"treasury_transaction"`
}

// This is an object representing a transaction on a Capital financing offer.
type CapitalFinancingTransactionDetails struct {
	// The advance amount being repaid, paid out, or reversed in minor units.
	AdvanceAmount int64 `json:"advance_amount"`
	// The currency of the financing transaction.
	Currency Currency `json:"currency"`
	// The fee amount being repaid, paid out, or reversed in minor units.
	FeeAmount int64 `json:"fee_amount"`
	// The linked payment for the transaction. This field only applies to financing transactions of type `paydown` and reason `automatic_withholding`.
	LinkedPayment string `json:"linked_payment"`
	// The reason for the financing transaction (if applicable).
	Reason CapitalFinancingTransactionDetailsReason `json:"reason"`
	// The reversed transaction. This field only applies to financing
	// transactions of type `reversal`.
	ReversedTransaction string `json:"reversed_transaction"`
	// The advance and fee amount being repaid, paid out, or reversed in minor units.
	TotalAmount int64 `json:"total_amount"`
	// This is an object representing a linked transaction on a Capital Financing Transaction.
	Transaction *CapitalFinancingTransactionDetailsTransaction `json:"transaction"`
}

// This is an object representing the details of a transaction on a Capital financing object.
type CapitalFinancingTransaction struct {
	APIResource
	// The ID of the merchant associated with this financing transaction.
	Account string `json:"account"`
	// Time at which the financing transaction was created. Given in seconds since unix epoch.
	CreatedAt int64 `json:"created_at"`
	// This is an object representing a transaction on a Capital financing offer.
	Details *CapitalFinancingTransactionDetails `json:"details"`
	// The Capital financing offer for this financing transaction.
	FinancingOffer string `json:"financing_offer"`
	// A unique identifier for the financing transaction object.
	ID string `json:"id"`
	// The Capital transaction object that predates the Financing Transactions API and
	// corresponds with the balance transaction that was created as a result of this
	// financing transaction.
	LegacyBalanceTransactionSource string `json:"legacy_balance_transaction_source"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The object type: financing_transaction
	Object string `json:"object"`
	// The type of the financing transaction.
	Type CapitalFinancingTransactionType `json:"type"`
	// A human-friendly description of the financing transaction.
	UserFacingDescription string `json:"user_facing_description"`
}

// CapitalFinancingTransactionList is a list of FinancingTransactions as retrieved from a list endpoint.
type CapitalFinancingTransactionList struct {
	APIResource
	ListMeta
	Data []*CapitalFinancingTransaction `json:"data"`
}
