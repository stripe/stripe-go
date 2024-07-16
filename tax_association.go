//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Status of the Tax Association.
type TaxAssociationStatus string

// List of values that TaxAssociationStatus can take
const (
	TaxAssociationStatusCommitted TaxAssociationStatus = "committed"
	TaxAssociationStatusErrored   TaxAssociationStatus = "errored"
)

// Status of the attempted Tax Transaction reversal.
type TaxAssociationStatusDetailsCommittedReversalStatus string

// List of values that TaxAssociationStatusDetailsCommittedReversalStatus can take
const (
	TaxAssociationStatusDetailsCommittedReversalStatusCommitted TaxAssociationStatusDetailsCommittedReversalStatus = "committed"
	TaxAssociationStatusDetailsCommittedReversalStatusErrored   TaxAssociationStatusDetailsCommittedReversalStatus = "errored"
)

// Details on why we could not commit the reversal Tax Transaction
type TaxAssociationStatusDetailsCommittedReversalStatusDetailsErroredReason string

// List of values that TaxAssociationStatusDetailsCommittedReversalStatusDetailsErroredReason can take
const (
	TaxAssociationStatusDetailsCommittedReversalStatusDetailsErroredReasonOriginalTransactionVoided TaxAssociationStatusDetailsCommittedReversalStatusDetailsErroredReason = "original_transaction_voided"
	TaxAssociationStatusDetailsCommittedReversalStatusDetailsErroredReasonUniqueReferenceViolation  TaxAssociationStatusDetailsCommittedReversalStatusDetailsErroredReason = "unique_reference_violation"
)

// Details on why we could not commit the Tax Transaction
type TaxAssociationStatusDetailsErroredReason string

// List of values that TaxAssociationStatusDetailsErroredReason can take
const (
	TaxAssociationStatusDetailsErroredReasonAnotherPaymentAssociatedWithCalculation TaxAssociationStatusDetailsErroredReason = "another_payment_associated_with_calculation"
	TaxAssociationStatusDetailsErroredReasonCalculationExpired                      TaxAssociationStatusDetailsErroredReason = "calculation_expired"
	TaxAssociationStatusDetailsErroredReasonCurrencyMismatch                        TaxAssociationStatusDetailsErroredReason = "currency_mismatch"
	TaxAssociationStatusDetailsErroredReasonUniqueReferenceViolation                TaxAssociationStatusDetailsErroredReason = "unique_reference_violation"
)

// Finds a tax association object by PaymentIntent id.
type TaxAssociationFindParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Valid [PaymentIntent](https://stripe.com/docs/api/payment_intents/object) id
	PaymentIntent *string `form:"payment_intent"`
}

// AddExpand appends a new field to expand.
func (p *TaxAssociationFindParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type TaxAssociationStatusDetailsCommittedReversalStatusDetailsCommitted struct {
	// The [Tax Transaction](https://stripe.com/docs/api/tax/transaction/object)
	Transaction string `json:"transaction"`
}
type TaxAssociationStatusDetailsCommittedReversalStatusDetailsErrored struct {
	// Details on why we could not commit the reversal Tax Transaction
	Reason TaxAssociationStatusDetailsCommittedReversalStatusDetailsErroredReason `json:"reason"`
	// The [Refund](https://stripe.com/docs/api/refunds/object) ID that should have created a tax reversal.
	RefundID string `json:"refund_id"`
}
type TaxAssociationStatusDetailsCommittedReversalStatusDetails struct {
	Committed *TaxAssociationStatusDetailsCommittedReversalStatusDetailsCommitted `json:"committed"`
	Errored   *TaxAssociationStatusDetailsCommittedReversalStatusDetailsErrored   `json:"errored"`
}

// Attempts to create Tax Transaction reversals
type TaxAssociationStatusDetailsCommittedReversal struct {
	// Status of the attempted Tax Transaction reversal.
	Status        TaxAssociationStatusDetailsCommittedReversalStatus         `json:"status"`
	StatusDetails *TaxAssociationStatusDetailsCommittedReversalStatusDetails `json:"status_details"`
}
type TaxAssociationStatusDetailsCommitted struct {
	// Attempts to create Tax Transaction reversals
	Reversals []*TaxAssociationStatusDetailsCommittedReversal `json:"reversals"`
	// The [Tax Transaction](https://stripe.com/docs/api/tax/transaction/object)
	Transaction string `json:"transaction"`
}
type TaxAssociationStatusDetailsErrored struct {
	// Details on why we could not commit the Tax Transaction
	Reason TaxAssociationStatusDetailsErroredReason `json:"reason"`
}
type TaxAssociationStatusDetails struct {
	Committed *TaxAssociationStatusDetailsCommitted `json:"committed"`
	Errored   *TaxAssociationStatusDetailsErrored   `json:"errored"`
}

// A Tax Association exposes the Tax Transactions that Stripe attempted to create on your behalf based on the PaymentIntent input
type TaxAssociation struct {
	APIResource
	// The [Tax Calculation](https://stripe.com/docs/api/tax/calculations/object) that was included in PaymentIntent.
	Calculation string `json:"calculation"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The [PaymentIntent](https://stripe.com/docs/api/payment_intents/object) that this Tax Association is tracking.
	PaymentIntent string `json:"payment_intent"`
	// Status of the Tax Association.
	Status        TaxAssociationStatus         `json:"status"`
	StatusDetails *TaxAssociationStatusDetails `json:"status_details"`
}
