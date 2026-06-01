//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The type of entity that collected this batch.
type V2CoreFeeBatchCollectedByType string

// List of values that V2CoreFeeBatchCollectedByType can take
const (
	V2CoreFeeBatchCollectedByTypeApplication V2CoreFeeBatchCollectedByType = "application"
	V2CoreFeeBatchCollectedByTypeNetwork     V2CoreFeeBatchCollectedByType = "network"
	V2CoreFeeBatchCollectedByTypeStripe      V2CoreFeeBatchCollectedByType = "stripe"
)

// The type of money movement object.
type V2CoreFeeBatchCollectionRecordType string

// List of values that V2CoreFeeBatchCollectionRecordType can take
const (
	V2CoreFeeBatchCollectionRecordTypeBalanceTransaction         V2CoreFeeBatchCollectionRecordType = "balance_transaction"
	V2CoreFeeBatchCollectionRecordTypeCreditTransaction          V2CoreFeeBatchCollectionRecordType = "credit_transaction"
	V2CoreFeeBatchCollectionRecordTypeMoneyManagementTransaction V2CoreFeeBatchCollectionRecordType = "money_management_transaction"
	V2CoreFeeBatchCollectionRecordTypePayableInvoice             V2CoreFeeBatchCollectionRecordType = "payable_invoice"
)

// The current state of this batch.
type V2CoreFeeBatchStatus string

// List of values that V2CoreFeeBatchStatus can take
const (
	V2CoreFeeBatchStatusBilled  V2CoreFeeBatchStatus = "billed"
	V2CoreFeeBatchStatusPending V2CoreFeeBatchStatus = "pending"
)

// The amount of tax adjusted for this batch.
type V2CoreFeeBatchAdjustmentsTaxAdjustment struct {
	// A lowercase alpha3 currency code like "usd"
	// For the taxonomy label choice, see SECURE_FRAMEWORKS-2849.
	Currency Currency `json:"currency"`
	// In major units like "1.23" for 1.23 USD
	// For the taxonomy label choice, see SECURE_FRAMEWORKS-2849.
	Value string `json:"value"`
}

// Adjustments applied to this batch.
type V2CoreFeeBatchAdjustments struct {
	// The amount of tax adjusted for this batch.
	TaxAdjustment *V2CoreFeeBatchAdjustmentsTaxAdjustment `json:"tax_adjustment,omitempty"`
}

// The total fee amount billed in this batch.
type V2CoreFeeBatchAmount struct {
	// A lowercase alpha3 currency code like "usd"
	// For the taxonomy label choice, see SECURE_FRAMEWORKS-2849.
	Currency Currency `json:"currency"`
	// In major units like "1.23" for 1.23 USD
	// For the taxonomy label choice, see SECURE_FRAMEWORKS-2849.
	Value string `json:"value"`
}

// The entity that collected this batch.
type V2CoreFeeBatchCollectedBy struct {
	// The type of entity that collected this batch.
	Type V2CoreFeeBatchCollectedByType `json:"type"`
}

// The fee amount collected via this record.
type V2CoreFeeBatchCollectionRecordAmount struct {
	// A lowercase alpha3 currency code like "usd"
	// For the taxonomy label choice, see SECURE_FRAMEWORKS-2849.
	Currency Currency `json:"currency"`
	// In major units like "1.23" for 1.23 USD
	// For the taxonomy label choice, see SECURE_FRAMEWORKS-2849.
	Value string `json:"value"`
}

// The tax amount collected via this record.
type V2CoreFeeBatchCollectionRecordTaxAmount struct {
	// A lowercase alpha3 currency code like "usd"
	// For the taxonomy label choice, see SECURE_FRAMEWORKS-2849.
	Currency Currency `json:"currency"`
	// In major units like "1.23" for 1.23 USD
	// For the taxonomy label choice, see SECURE_FRAMEWORKS-2849.
	Value string `json:"value"`
}

// The tax amount collected via this record.
type V2CoreFeeBatchCollectionRecordTax struct {
	// The tax amount collected via this record.
	Amount *V2CoreFeeBatchCollectionRecordTaxAmount `json:"amount"`
}

// The money movement records associated with collecting this batch.
type V2CoreFeeBatchCollectionRecord struct {
	// The fee amount collected via this record.
	Amount *V2CoreFeeBatchCollectionRecordAmount `json:"amount"`
	// The ID of the associated v1 balance transaction.
	BalanceTransaction string `json:"balance_transaction,omitempty"`
	// The ID of the associated credit transaction.
	CreditTransaction string `json:"credit_transaction,omitempty"`
	// The ID of the associated v2 money management transaction.
	MoneyManagementTransaction string `json:"money_management_transaction,omitempty"`
	// The ID of the associated accounts-receivable invoice.
	PayableInvoice string `json:"payable_invoice,omitempty"`
	// The tax amount collected via this record.
	Tax *V2CoreFeeBatchCollectionRecordTax `json:"tax,omitempty"`
	// The type of money movement object.
	Type V2CoreFeeBatchCollectionRecordType `json:"type"`
}

// Timestamps for each status transition.
type V2CoreFeeBatchStatusTransitions struct {
	// Timestamp of when the batch transitioned to BILLED, if applicable.
	BilledAt time.Time `json:"billed_at,omitempty"`
}

// The tax amount included in this batch.
type V2CoreFeeBatchTaxAmount struct {
	// A lowercase alpha3 currency code like "usd"
	// For the taxonomy label choice, see SECURE_FRAMEWORKS-2849.
	Currency Currency `json:"currency"`
	// In major units like "1.23" for 1.23 USD
	// For the taxonomy label choice, see SECURE_FRAMEWORKS-2849.
	Value string `json:"value"`
}

// The tax amount included in this batch.
type V2CoreFeeBatchTax struct {
	// The tax amount included in this batch.
	Amount *V2CoreFeeBatchTaxAmount `json:"amount"`
}

// A FeeBatch represents a settlement grouping of fees.
// It bridges the fee domain with actual money movement, tracking what was settled and how.
type V2CoreFeeBatch struct {
	APIResource
	// Adjustments applied to this batch.
	Adjustments *V2CoreFeeBatchAdjustments `json:"adjustments,omitempty"`
	// The total fee amount billed in this batch.
	Amount *V2CoreFeeBatchAmount `json:"amount"`
	// The entity that collected this batch.
	CollectedBy *V2CoreFeeBatchCollectedBy `json:"collected_by"`
	// The money movement records associated with collecting this batch.
	CollectionRecords []*V2CoreFeeBatchCollectionRecord `json:"collection_records"`
	// Timestamp of when this batch was created.
	Created time.Time `json:"created"`
	// Unique identifier for the FeeBatch.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode, or `false` if in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The current state of this batch.
	Status V2CoreFeeBatchStatus `json:"status"`
	// Timestamps for each status transition.
	StatusTransitions *V2CoreFeeBatchStatusTransitions `json:"status_transitions"`
	// The tax amount included in this batch.
	Tax *V2CoreFeeBatchTax `json:"tax,omitempty"`
}
