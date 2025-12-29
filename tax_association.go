//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Details on why we couldn't commit the tax transaction.
type TaxAssociationTaxTransactionAttemptErroredReason string

// List of values that TaxAssociationTaxTransactionAttemptErroredReason can take
const (
	TaxAssociationTaxTransactionAttemptErroredReasonAnotherPaymentAssociatedWithCalculation TaxAssociationTaxTransactionAttemptErroredReason = "another_payment_associated_with_calculation"
	TaxAssociationTaxTransactionAttemptErroredReasonCalculationExpired                      TaxAssociationTaxTransactionAttemptErroredReason = "calculation_expired"
	TaxAssociationTaxTransactionAttemptErroredReasonCurrencyMismatch                        TaxAssociationTaxTransactionAttemptErroredReason = "currency_mismatch"
	TaxAssociationTaxTransactionAttemptErroredReasonOriginalTransactionVoided               TaxAssociationTaxTransactionAttemptErroredReason = "original_transaction_voided"
	TaxAssociationTaxTransactionAttemptErroredReasonUniqueReferenceViolation                TaxAssociationTaxTransactionAttemptErroredReason = "unique_reference_violation"
)

// Finds a tax association object by PaymentIntent id.
type TaxAssociationFindParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Valid [PaymentIntent](https://docs.stripe.com/api/payment_intents/object) id
	PaymentIntent *string `form:"payment_intent"`
}

// AddExpand appends a new field to expand.
func (p *TaxAssociationFindParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type TaxAssociationTaxTransactionAttemptCommitted struct {
	// The [Tax Transaction](https://docs.stripe.com/api/tax/transaction/object)
	Transaction string `json:"transaction"`
}
type TaxAssociationTaxTransactionAttemptErrored struct {
	// Details on why we couldn't commit the tax transaction.
	Reason TaxAssociationTaxTransactionAttemptErroredReason `json:"reason"`
}

// Information about the tax transactions linked to this payment intent
type TaxAssociationTaxTransactionAttempt struct {
	Committed *TaxAssociationTaxTransactionAttemptCommitted `json:"committed"`
	Errored   *TaxAssociationTaxTransactionAttemptErrored   `json:"errored"`
	// The source of the tax transaction attempt. This is either a refund or a payment intent.
	Source string `json:"source"`
	// The status of the transaction attempt. This can be `errored` or `committed`.
	Status string `json:"status"`
}

// A Tax Association exposes the Tax Transactions that Stripe attempted to create on your behalf based on the PaymentIntent input
type TaxAssociation struct {
	APIResource
	// The [Tax Calculation](https://docs.stripe.com/api/tax/calculations/object) that was included in PaymentIntent.
	Calculation string `json:"calculation"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The [PaymentIntent](https://docs.stripe.com/api/payment_intents/object) that this Tax Association is tracking.
	PaymentIntent string `json:"payment_intent"`
	// Information about the tax transactions linked to this payment intent
	TaxTransactionAttempts []*TaxAssociationTaxTransactionAttempt `json:"tax_transaction_attempts"`
}
