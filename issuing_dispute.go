//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Whether the product was a merchandise or service.
type IssuingDisputeEvidenceCanceledProductType string

// List of values that IssuingDisputeEvidenceCanceledProductType can take
const (
	IssuingDisputeEvidenceCanceledProductTypeMerchandise IssuingDisputeEvidenceCanceledProductType = "merchandise"
	IssuingDisputeEvidenceCanceledProductTypeService     IssuingDisputeEvidenceCanceledProductType = "service"
)

// Result of cardholder's attempt to return the product.
type IssuingDisputeEvidenceCanceledReturnStatus string

// List of values that IssuingDisputeEvidenceCanceledReturnStatus can take
const (
	IssuingDisputeEvidenceCanceledReturnStatusMerchantRejected IssuingDisputeEvidenceCanceledReturnStatus = "merchant_rejected"
	IssuingDisputeEvidenceCanceledReturnStatusSuccessful       IssuingDisputeEvidenceCanceledReturnStatus = "successful"
)

// Result of cardholder's attempt to return the product.
type IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatus string

// List of values that IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatus can take
const (
	IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatusMerchantRejected IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatus = "merchant_rejected"
	IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatusSuccessful       IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatus = "successful"
)

// Whether the product was a merchandise or service.
type IssuingDisputeEvidenceNotReceivedProductType string

// List of values that IssuingDisputeEvidenceNotReceivedProductType can take
const (
	IssuingDisputeEvidenceNotReceivedProductTypeMerchandise IssuingDisputeEvidenceNotReceivedProductType = "merchandise"
	IssuingDisputeEvidenceNotReceivedProductTypeService     IssuingDisputeEvidenceNotReceivedProductType = "service"
)

// Whether the product was a merchandise or service.
type IssuingDisputeEvidenceOtherProductType string

// List of values that IssuingDisputeEvidenceOtherProductType can take
const (
	IssuingDisputeEvidenceOtherProductTypeMerchandise IssuingDisputeEvidenceOtherProductType = "merchandise"
	IssuingDisputeEvidenceOtherProductTypeService     IssuingDisputeEvidenceOtherProductType = "service"
)

// The reason for filing the dispute. Its value will match the field containing the evidence.
type IssuingDisputeEvidenceReason string

// List of values that IssuingDisputeEvidenceReason can take
const (
	IssuingDisputeEvidenceReasonCanceled                  IssuingDisputeEvidenceReason = "canceled"
	IssuingDisputeEvidenceReasonDuplicate                 IssuingDisputeEvidenceReason = "duplicate"
	IssuingDisputeEvidenceReasonFraudulent                IssuingDisputeEvidenceReason = "fraudulent"
	IssuingDisputeEvidenceReasonMerchandiseNotAsDescribed IssuingDisputeEvidenceReason = "merchandise_not_as_described"
	IssuingDisputeEvidenceReasonNoValidAuthorization      IssuingDisputeEvidenceReason = "no_valid_authorization"
	IssuingDisputeEvidenceReasonNotReceived               IssuingDisputeEvidenceReason = "not_received"
	IssuingDisputeEvidenceReasonOther                     IssuingDisputeEvidenceReason = "other"
	IssuingDisputeEvidenceReasonServiceNotAsDescribed     IssuingDisputeEvidenceReason = "service_not_as_described"
)

// The enum that describes the dispute loss outcome. If the dispute is not lost, this field will be absent. New enum values may be added in the future, so be sure to handle unknown values.
type IssuingDisputeLossReason string

// List of values that IssuingDisputeLossReason can take
const (
	IssuingDisputeLossReasonCardholderAuthenticationIssuerLiability       IssuingDisputeLossReason = "cardholder_authentication_issuer_liability"
	IssuingDisputeLossReasonEci5TokenTransactionWithTavv                  IssuingDisputeLossReason = "eci5_token_transaction_with_tavv"
	IssuingDisputeLossReasonExcessDisputesInTimeframe                     IssuingDisputeLossReason = "excess_disputes_in_timeframe"
	IssuingDisputeLossReasonHasNotMetTheMinimumDisputeAmountRequirements  IssuingDisputeLossReason = "has_not_met_the_minimum_dispute_amount_requirements"
	IssuingDisputeLossReasonInvalidDuplicateDispute                       IssuingDisputeLossReason = "invalid_duplicate_dispute"
	IssuingDisputeLossReasonInvalidIncorrectAmountDispute                 IssuingDisputeLossReason = "invalid_incorrect_amount_dispute"
	IssuingDisputeLossReasonInvalidNoAuthorization                        IssuingDisputeLossReason = "invalid_no_authorization"
	IssuingDisputeLossReasonInvalidUseOfDisputes                          IssuingDisputeLossReason = "invalid_use_of_disputes"
	IssuingDisputeLossReasonMerchandiseDeliveredOrShipped                 IssuingDisputeLossReason = "merchandise_delivered_or_shipped"
	IssuingDisputeLossReasonMerchandiseOrServiceAsDescribed               IssuingDisputeLossReason = "merchandise_or_service_as_described"
	IssuingDisputeLossReasonNotCancelled                                  IssuingDisputeLossReason = "not_cancelled"
	IssuingDisputeLossReasonOther                                         IssuingDisputeLossReason = "other"
	IssuingDisputeLossReasonRefundIssued                                  IssuingDisputeLossReason = "refund_issued"
	IssuingDisputeLossReasonSubmittedBeyondAllowableTimeLimit             IssuingDisputeLossReason = "submitted_beyond_allowable_time_limit"
	IssuingDisputeLossReasonTransaction3dsRequired                        IssuingDisputeLossReason = "transaction_3ds_required"
	IssuingDisputeLossReasonTransactionApprovedAfterPriorFraudDispute     IssuingDisputeLossReason = "transaction_approved_after_prior_fraud_dispute"
	IssuingDisputeLossReasonTransactionAuthorized                         IssuingDisputeLossReason = "transaction_authorized"
	IssuingDisputeLossReasonTransactionElectronicallyRead                 IssuingDisputeLossReason = "transaction_electronically_read"
	IssuingDisputeLossReasonTransactionQualifiesForVisaEasyPaymentService IssuingDisputeLossReason = "transaction_qualifies_for_visa_easy_payment_service"
	IssuingDisputeLossReasonTransactionUnattended                         IssuingDisputeLossReason = "transaction_unattended"
)

// Current status of the dispute.
type IssuingDisputeStatus string

// List of values that IssuingDisputeStatus can take
const (
	IssuingDisputeStatusExpired     IssuingDisputeStatus = "expired"
	IssuingDisputeStatusLost        IssuingDisputeStatus = "lost"
	IssuingDisputeStatusSubmitted   IssuingDisputeStatus = "submitted"
	IssuingDisputeStatusUnsubmitted IssuingDisputeStatus = "unsubmitted"
	IssuingDisputeStatusWon         IssuingDisputeStatus = "won"
)

// Returns a list of Issuing Dispute objects. The objects are sorted in descending order by creation date, with the most recently created object appearing first.
type IssuingDisputeListParams struct {
	ListParams `form:"*"`
	// Only return Issuing disputes that were created during the given date interval.
	Created *int64 `form:"created"`
	// Only return Issuing disputes that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Select Issuing disputes with the given status.
	Status *string `form:"status"`
	// Select the Issuing dispute for the given transaction.
	Transaction *string `form:"transaction"`
}

// AddExpand appends a new field to expand.
func (p *IssuingDisputeListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Evidence provided when `reason` is 'canceled'.
type IssuingDisputeEvidenceCanceledParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Date when order was canceled.
	CanceledAt *int64 `form:"canceled_at"`
	// Whether the cardholder was provided with a cancellation policy.
	CancellationPolicyProvided *bool `form:"cancellation_policy_provided"`
	// Reason for canceling the order.
	CancellationReason *string `form:"cancellation_reason"`
	// Date when the cardholder expected to receive the product.
	ExpectedAt *int64 `form:"expected_at"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Description of the merchandise or service that was purchased.
	ProductDescription *string `form:"product_description"`
	// Whether the product was a merchandise or service.
	ProductType *string `form:"product_type"`
	// Date when the product was returned or attempted to be returned.
	ReturnedAt *int64 `form:"returned_at"`
	// Result of cardholder's attempt to return the product.
	ReturnStatus *string                                          `form:"return_status"`
	UnsetFields  []IssuingDisputeEvidenceCanceledParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeEvidenceCanceledParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeEvidenceCanceledParams.
type IssuingDisputeEvidenceCanceledParamsUnsetField string

const (
	IssuingDisputeEvidenceCanceledParamsUnsetFieldAdditionalDocumentation    IssuingDisputeEvidenceCanceledParamsUnsetField = "additional_documentation"
	IssuingDisputeEvidenceCanceledParamsUnsetFieldCanceledAt                 IssuingDisputeEvidenceCanceledParamsUnsetField = "canceled_at"
	IssuingDisputeEvidenceCanceledParamsUnsetFieldCancellationPolicyProvided IssuingDisputeEvidenceCanceledParamsUnsetField = "cancellation_policy_provided"
	IssuingDisputeEvidenceCanceledParamsUnsetFieldCancellationReason         IssuingDisputeEvidenceCanceledParamsUnsetField = "cancellation_reason"
	IssuingDisputeEvidenceCanceledParamsUnsetFieldExpectedAt                 IssuingDisputeEvidenceCanceledParamsUnsetField = "expected_at"
	IssuingDisputeEvidenceCanceledParamsUnsetFieldExplanation                IssuingDisputeEvidenceCanceledParamsUnsetField = "explanation"
	IssuingDisputeEvidenceCanceledParamsUnsetFieldProductDescription         IssuingDisputeEvidenceCanceledParamsUnsetField = "product_description"
	IssuingDisputeEvidenceCanceledParamsUnsetFieldProductType                IssuingDisputeEvidenceCanceledParamsUnsetField = "product_type"
	IssuingDisputeEvidenceCanceledParamsUnsetFieldReturnStatus               IssuingDisputeEvidenceCanceledParamsUnsetField = "return_status"
	IssuingDisputeEvidenceCanceledParamsUnsetFieldReturnedAt                 IssuingDisputeEvidenceCanceledParamsUnsetField = "returned_at"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeEvidenceCanceledParams) AddUnsetField(field IssuingDisputeEvidenceCanceledParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'duplicate'.
type IssuingDisputeEvidenceDuplicateParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Copy of the card statement showing that the product had already been paid for.
	CardStatement *string `form:"card_statement"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Copy of the receipt showing that the product had been paid for in cash.
	CashReceipt *string `form:"cash_receipt"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Image of the front and back of the check that was used to pay for the product.
	CheckImage *string `form:"check_image"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Transaction (e.g., ipi_...) that the disputed transaction is a duplicate of. Of the two or more transactions that are copies of each other, this is original undisputed one.
	OriginalTransaction *string                                           `form:"original_transaction"`
	UnsetFields         []IssuingDisputeEvidenceDuplicateParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeEvidenceDuplicateParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeEvidenceDuplicateParams.
type IssuingDisputeEvidenceDuplicateParamsUnsetField string

const (
	IssuingDisputeEvidenceDuplicateParamsUnsetFieldAdditionalDocumentation IssuingDisputeEvidenceDuplicateParamsUnsetField = "additional_documentation"
	IssuingDisputeEvidenceDuplicateParamsUnsetFieldCardStatement           IssuingDisputeEvidenceDuplicateParamsUnsetField = "card_statement"
	IssuingDisputeEvidenceDuplicateParamsUnsetFieldCashReceipt             IssuingDisputeEvidenceDuplicateParamsUnsetField = "cash_receipt"
	IssuingDisputeEvidenceDuplicateParamsUnsetFieldCheckImage              IssuingDisputeEvidenceDuplicateParamsUnsetField = "check_image"
	IssuingDisputeEvidenceDuplicateParamsUnsetFieldExplanation             IssuingDisputeEvidenceDuplicateParamsUnsetField = "explanation"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeEvidenceDuplicateParams) AddUnsetField(field IssuingDisputeEvidenceDuplicateParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'fraudulent'.
type IssuingDisputeEvidenceFraudulentParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string                                            `form:"explanation"`
	UnsetFields []IssuingDisputeEvidenceFraudulentParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeEvidenceFraudulentParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeEvidenceFraudulentParams.
type IssuingDisputeEvidenceFraudulentParamsUnsetField string

const (
	IssuingDisputeEvidenceFraudulentParamsUnsetFieldAdditionalDocumentation IssuingDisputeEvidenceFraudulentParamsUnsetField = "additional_documentation"
	IssuingDisputeEvidenceFraudulentParamsUnsetFieldExplanation             IssuingDisputeEvidenceFraudulentParamsUnsetField = "explanation"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeEvidenceFraudulentParams) AddUnsetField(field IssuingDisputeEvidenceFraudulentParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'merchandise_not_as_described'.
type IssuingDisputeEvidenceMerchandiseNotAsDescribedParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Date when the product was received.
	ReceivedAt *int64 `form:"received_at"`
	// Description of the cardholder's attempt to return the product.
	ReturnDescription *string `form:"return_description"`
	// Date when the product was returned or attempted to be returned.
	ReturnedAt *int64 `form:"returned_at"`
	// Result of cardholder's attempt to return the product.
	ReturnStatus *string                                                           `form:"return_status"`
	UnsetFields  []IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeEvidenceMerchandiseNotAsDescribedParams.
type IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetField string

const (
	IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetFieldAdditionalDocumentation IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetField = "additional_documentation"
	IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetFieldExplanation             IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetField = "explanation"
	IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetFieldReceivedAt              IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetField = "received_at"
	IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetFieldReturnDescription       IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetField = "return_description"
	IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetFieldReturnStatus            IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetField = "return_status"
	IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetFieldReturnedAt              IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetField = "returned_at"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeEvidenceMerchandiseNotAsDescribedParams) AddUnsetField(field IssuingDisputeEvidenceMerchandiseNotAsDescribedParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'no_valid_authorization'.
type IssuingDisputeEvidenceNoValidAuthorizationParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string                                                      `form:"explanation"`
	UnsetFields []IssuingDisputeEvidenceNoValidAuthorizationParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeEvidenceNoValidAuthorizationParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeEvidenceNoValidAuthorizationParams.
type IssuingDisputeEvidenceNoValidAuthorizationParamsUnsetField string

const (
	IssuingDisputeEvidenceNoValidAuthorizationParamsUnsetFieldAdditionalDocumentation IssuingDisputeEvidenceNoValidAuthorizationParamsUnsetField = "additional_documentation"
	IssuingDisputeEvidenceNoValidAuthorizationParamsUnsetFieldExplanation             IssuingDisputeEvidenceNoValidAuthorizationParamsUnsetField = "explanation"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeEvidenceNoValidAuthorizationParams) AddUnsetField(field IssuingDisputeEvidenceNoValidAuthorizationParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'not_received'.
type IssuingDisputeEvidenceNotReceivedParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Date when the cardholder expected to receive the product.
	ExpectedAt *int64 `form:"expected_at"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Description of the merchandise or service that was purchased.
	ProductDescription *string `form:"product_description"`
	// Whether the product was a merchandise or service.
	ProductType *string                                             `form:"product_type"`
	UnsetFields []IssuingDisputeEvidenceNotReceivedParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeEvidenceNotReceivedParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeEvidenceNotReceivedParams.
type IssuingDisputeEvidenceNotReceivedParamsUnsetField string

const (
	IssuingDisputeEvidenceNotReceivedParamsUnsetFieldAdditionalDocumentation IssuingDisputeEvidenceNotReceivedParamsUnsetField = "additional_documentation"
	IssuingDisputeEvidenceNotReceivedParamsUnsetFieldExpectedAt              IssuingDisputeEvidenceNotReceivedParamsUnsetField = "expected_at"
	IssuingDisputeEvidenceNotReceivedParamsUnsetFieldExplanation             IssuingDisputeEvidenceNotReceivedParamsUnsetField = "explanation"
	IssuingDisputeEvidenceNotReceivedParamsUnsetFieldProductDescription      IssuingDisputeEvidenceNotReceivedParamsUnsetField = "product_description"
	IssuingDisputeEvidenceNotReceivedParamsUnsetFieldProductType             IssuingDisputeEvidenceNotReceivedParamsUnsetField = "product_type"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeEvidenceNotReceivedParams) AddUnsetField(field IssuingDisputeEvidenceNotReceivedParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'other'.
type IssuingDisputeEvidenceOtherParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Description of the merchandise or service that was purchased.
	ProductDescription *string `form:"product_description"`
	// Whether the product was a merchandise or service.
	ProductType *string                                       `form:"product_type"`
	UnsetFields []IssuingDisputeEvidenceOtherParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeEvidenceOtherParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeEvidenceOtherParams.
type IssuingDisputeEvidenceOtherParamsUnsetField string

const (
	IssuingDisputeEvidenceOtherParamsUnsetFieldAdditionalDocumentation IssuingDisputeEvidenceOtherParamsUnsetField = "additional_documentation"
	IssuingDisputeEvidenceOtherParamsUnsetFieldExplanation             IssuingDisputeEvidenceOtherParamsUnsetField = "explanation"
	IssuingDisputeEvidenceOtherParamsUnsetFieldProductDescription      IssuingDisputeEvidenceOtherParamsUnsetField = "product_description"
	IssuingDisputeEvidenceOtherParamsUnsetFieldProductType             IssuingDisputeEvidenceOtherParamsUnsetField = "product_type"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeEvidenceOtherParams) AddUnsetField(field IssuingDisputeEvidenceOtherParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'service_not_as_described'.
type IssuingDisputeEvidenceServiceNotAsDescribedParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Date when order was canceled.
	CanceledAt *int64 `form:"canceled_at"`
	// Reason for canceling the order.
	CancellationReason *string `form:"cancellation_reason"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Date when the product was received.
	ReceivedAt  *int64                                                        `form:"received_at"`
	UnsetFields []IssuingDisputeEvidenceServiceNotAsDescribedParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeEvidenceServiceNotAsDescribedParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeEvidenceServiceNotAsDescribedParams.
type IssuingDisputeEvidenceServiceNotAsDescribedParamsUnsetField string

const (
	IssuingDisputeEvidenceServiceNotAsDescribedParamsUnsetFieldAdditionalDocumentation IssuingDisputeEvidenceServiceNotAsDescribedParamsUnsetField = "additional_documentation"
	IssuingDisputeEvidenceServiceNotAsDescribedParamsUnsetFieldCanceledAt              IssuingDisputeEvidenceServiceNotAsDescribedParamsUnsetField = "canceled_at"
	IssuingDisputeEvidenceServiceNotAsDescribedParamsUnsetFieldCancellationReason      IssuingDisputeEvidenceServiceNotAsDescribedParamsUnsetField = "cancellation_reason"
	IssuingDisputeEvidenceServiceNotAsDescribedParamsUnsetFieldExplanation             IssuingDisputeEvidenceServiceNotAsDescribedParamsUnsetField = "explanation"
	IssuingDisputeEvidenceServiceNotAsDescribedParamsUnsetFieldReceivedAt              IssuingDisputeEvidenceServiceNotAsDescribedParamsUnsetField = "received_at"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeEvidenceServiceNotAsDescribedParams) AddUnsetField(field IssuingDisputeEvidenceServiceNotAsDescribedParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided for the dispute.
type IssuingDisputeEvidenceParams struct {
	// Evidence provided when `reason` is 'canceled'.
	Canceled *IssuingDisputeEvidenceCanceledParams `form:"canceled"`
	// Evidence provided when `reason` is 'duplicate'.
	Duplicate *IssuingDisputeEvidenceDuplicateParams `form:"duplicate"`
	// Evidence provided when `reason` is 'fraudulent'.
	Fraudulent *IssuingDisputeEvidenceFraudulentParams `form:"fraudulent"`
	// Evidence provided when `reason` is 'merchandise_not_as_described'.
	MerchandiseNotAsDescribed *IssuingDisputeEvidenceMerchandiseNotAsDescribedParams `form:"merchandise_not_as_described"`
	// Evidence provided when `reason` is 'not_received'.
	NotReceived *IssuingDisputeEvidenceNotReceivedParams `form:"not_received"`
	// Evidence provided when `reason` is 'no_valid_authorization'.
	NoValidAuthorization *IssuingDisputeEvidenceNoValidAuthorizationParams `form:"no_valid_authorization"`
	// Evidence provided when `reason` is 'other'.
	Other *IssuingDisputeEvidenceOtherParams `form:"other"`
	// The reason for filing the dispute. The evidence should be submitted in the field of the same name.
	Reason *string `form:"reason"`
	// Evidence provided when `reason` is 'service_not_as_described'.
	ServiceNotAsDescribed *IssuingDisputeEvidenceServiceNotAsDescribedParams `form:"service_not_as_described"`
	UnsetFields           []IssuingDisputeEvidenceParamsUnsetField           `form:"-" json:"-"`
}

// IssuingDisputeEvidenceParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeEvidenceParams.
type IssuingDisputeEvidenceParamsUnsetField string

const (
	IssuingDisputeEvidenceParamsUnsetFieldCanceled                  IssuingDisputeEvidenceParamsUnsetField = "canceled"
	IssuingDisputeEvidenceParamsUnsetFieldDuplicate                 IssuingDisputeEvidenceParamsUnsetField = "duplicate"
	IssuingDisputeEvidenceParamsUnsetFieldFraudulent                IssuingDisputeEvidenceParamsUnsetField = "fraudulent"
	IssuingDisputeEvidenceParamsUnsetFieldMerchandiseNotAsDescribed IssuingDisputeEvidenceParamsUnsetField = "merchandise_not_as_described"
	IssuingDisputeEvidenceParamsUnsetFieldNoValidAuthorization      IssuingDisputeEvidenceParamsUnsetField = "no_valid_authorization"
	IssuingDisputeEvidenceParamsUnsetFieldNotReceived               IssuingDisputeEvidenceParamsUnsetField = "not_received"
	IssuingDisputeEvidenceParamsUnsetFieldOther                     IssuingDisputeEvidenceParamsUnsetField = "other"
	IssuingDisputeEvidenceParamsUnsetFieldServiceNotAsDescribed     IssuingDisputeEvidenceParamsUnsetField = "service_not_as_described"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeEvidenceParams) AddUnsetField(field IssuingDisputeEvidenceParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Params for disputes related to Treasury FinancialAccounts
type IssuingDisputeTreasuryParams struct {
	// The ID of the ReceivedDebit to initiate an Issuings dispute for.
	ReceivedDebit *string `form:"received_debit"`
}

// Creates an Issuing Dispute object. Individual pieces of evidence within the evidence object are optional at this point. Stripe only validates that required evidence is present during submission. Refer to [Dispute reasons and evidence](https://docs.stripe.com/docs/issuing/purchases/disputes#dispute-reasons-and-evidence) for more details about evidence requirements.
type IssuingDisputeParams struct {
	Params `form:"*"`
	// The dispute amount in the card's currency and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal). If not set, defaults to the full transaction amount.
	Amount *int64 `form:"amount"`
	// Evidence provided for the dispute.
	Evidence *IssuingDisputeEvidenceParams `form:"evidence"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The ID of the issuing transaction to create a dispute for. For transaction on Treasury FinancialAccounts, use `treasury.received_debit`.
	Transaction *string `form:"transaction"`
	// Params for disputes related to Treasury FinancialAccounts
	Treasury *IssuingDisputeTreasuryParams `form:"treasury"`
}

// IssuingDisputeParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeParams.
type IssuingDisputeParamsUnsetField string

const (
	IssuingDisputeParamsUnsetFieldMetadata IssuingDisputeParamsUnsetField = "metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeParams) AddUnsetField(field IssuingDisputeParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, string(field))
}

// AddExpand appends a new field to expand.
func (p *IssuingDisputeParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *IssuingDisputeParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Submits an Issuing Dispute to the card network. Stripe validates that all evidence fields required for the dispute's reason are present. For more details, see [Dispute reasons and evidence](https://docs.stripe.com/docs/issuing/purchases/disputes#dispute-reasons-and-evidence).
type IssuingDisputeSubmitParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
}

// IssuingDisputeSubmitParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeSubmitParams.
type IssuingDisputeSubmitParamsUnsetField string

const (
	IssuingDisputeSubmitParamsUnsetFieldMetadata IssuingDisputeSubmitParamsUnsetField = "metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeSubmitParams) AddUnsetField(field IssuingDisputeSubmitParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, string(field))
}

// AddExpand appends a new field to expand.
func (p *IssuingDisputeSubmitParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *IssuingDisputeSubmitParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Evidence provided when `reason` is 'canceled'.
type IssuingDisputeCreateEvidenceCanceledParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Date when order was canceled.
	CanceledAt *int64 `form:"canceled_at"`
	// Whether the cardholder was provided with a cancellation policy.
	CancellationPolicyProvided *bool `form:"cancellation_policy_provided"`
	// Reason for canceling the order.
	CancellationReason *string `form:"cancellation_reason"`
	// Date when the cardholder expected to receive the product.
	ExpectedAt *int64 `form:"expected_at"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Description of the merchandise or service that was purchased.
	ProductDescription *string `form:"product_description"`
	// Whether the product was a merchandise or service.
	ProductType *string `form:"product_type"`
	// Date when the product was returned or attempted to be returned.
	ReturnedAt *int64 `form:"returned_at"`
	// Result of cardholder's attempt to return the product.
	ReturnStatus *string                                                `form:"return_status"`
	UnsetFields  []IssuingDisputeCreateEvidenceCanceledParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeCreateEvidenceCanceledParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeCreateEvidenceCanceledParams.
type IssuingDisputeCreateEvidenceCanceledParamsUnsetField string

const (
	IssuingDisputeCreateEvidenceCanceledParamsUnsetFieldAdditionalDocumentation    IssuingDisputeCreateEvidenceCanceledParamsUnsetField = "additional_documentation"
	IssuingDisputeCreateEvidenceCanceledParamsUnsetFieldCanceledAt                 IssuingDisputeCreateEvidenceCanceledParamsUnsetField = "canceled_at"
	IssuingDisputeCreateEvidenceCanceledParamsUnsetFieldCancellationPolicyProvided IssuingDisputeCreateEvidenceCanceledParamsUnsetField = "cancellation_policy_provided"
	IssuingDisputeCreateEvidenceCanceledParamsUnsetFieldCancellationReason         IssuingDisputeCreateEvidenceCanceledParamsUnsetField = "cancellation_reason"
	IssuingDisputeCreateEvidenceCanceledParamsUnsetFieldExpectedAt                 IssuingDisputeCreateEvidenceCanceledParamsUnsetField = "expected_at"
	IssuingDisputeCreateEvidenceCanceledParamsUnsetFieldExplanation                IssuingDisputeCreateEvidenceCanceledParamsUnsetField = "explanation"
	IssuingDisputeCreateEvidenceCanceledParamsUnsetFieldProductDescription         IssuingDisputeCreateEvidenceCanceledParamsUnsetField = "product_description"
	IssuingDisputeCreateEvidenceCanceledParamsUnsetFieldProductType                IssuingDisputeCreateEvidenceCanceledParamsUnsetField = "product_type"
	IssuingDisputeCreateEvidenceCanceledParamsUnsetFieldReturnStatus               IssuingDisputeCreateEvidenceCanceledParamsUnsetField = "return_status"
	IssuingDisputeCreateEvidenceCanceledParamsUnsetFieldReturnedAt                 IssuingDisputeCreateEvidenceCanceledParamsUnsetField = "returned_at"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeCreateEvidenceCanceledParams) AddUnsetField(field IssuingDisputeCreateEvidenceCanceledParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'duplicate'.
type IssuingDisputeCreateEvidenceDuplicateParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Copy of the card statement showing that the product had already been paid for.
	CardStatement *string `form:"card_statement"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Copy of the receipt showing that the product had been paid for in cash.
	CashReceipt *string `form:"cash_receipt"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Image of the front and back of the check that was used to pay for the product.
	CheckImage *string `form:"check_image"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Transaction (e.g., ipi_...) that the disputed transaction is a duplicate of. Of the two or more transactions that are copies of each other, this is original undisputed one.
	OriginalTransaction *string                                                 `form:"original_transaction"`
	UnsetFields         []IssuingDisputeCreateEvidenceDuplicateParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeCreateEvidenceDuplicateParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeCreateEvidenceDuplicateParams.
type IssuingDisputeCreateEvidenceDuplicateParamsUnsetField string

const (
	IssuingDisputeCreateEvidenceDuplicateParamsUnsetFieldAdditionalDocumentation IssuingDisputeCreateEvidenceDuplicateParamsUnsetField = "additional_documentation"
	IssuingDisputeCreateEvidenceDuplicateParamsUnsetFieldCardStatement           IssuingDisputeCreateEvidenceDuplicateParamsUnsetField = "card_statement"
	IssuingDisputeCreateEvidenceDuplicateParamsUnsetFieldCashReceipt             IssuingDisputeCreateEvidenceDuplicateParamsUnsetField = "cash_receipt"
	IssuingDisputeCreateEvidenceDuplicateParamsUnsetFieldCheckImage              IssuingDisputeCreateEvidenceDuplicateParamsUnsetField = "check_image"
	IssuingDisputeCreateEvidenceDuplicateParamsUnsetFieldExplanation             IssuingDisputeCreateEvidenceDuplicateParamsUnsetField = "explanation"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeCreateEvidenceDuplicateParams) AddUnsetField(field IssuingDisputeCreateEvidenceDuplicateParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'fraudulent'.
type IssuingDisputeCreateEvidenceFraudulentParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string                                                  `form:"explanation"`
	UnsetFields []IssuingDisputeCreateEvidenceFraudulentParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeCreateEvidenceFraudulentParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeCreateEvidenceFraudulentParams.
type IssuingDisputeCreateEvidenceFraudulentParamsUnsetField string

const (
	IssuingDisputeCreateEvidenceFraudulentParamsUnsetFieldAdditionalDocumentation IssuingDisputeCreateEvidenceFraudulentParamsUnsetField = "additional_documentation"
	IssuingDisputeCreateEvidenceFraudulentParamsUnsetFieldExplanation             IssuingDisputeCreateEvidenceFraudulentParamsUnsetField = "explanation"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeCreateEvidenceFraudulentParams) AddUnsetField(field IssuingDisputeCreateEvidenceFraudulentParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'merchandise_not_as_described'.
type IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Date when the product was received.
	ReceivedAt *int64 `form:"received_at"`
	// Description of the cardholder's attempt to return the product.
	ReturnDescription *string `form:"return_description"`
	// Date when the product was returned or attempted to be returned.
	ReturnedAt *int64 `form:"returned_at"`
	// Result of cardholder's attempt to return the product.
	ReturnStatus *string                                                                 `form:"return_status"`
	UnsetFields  []IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParams.
type IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetField string

const (
	IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetFieldAdditionalDocumentation IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetField = "additional_documentation"
	IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetFieldExplanation             IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetField = "explanation"
	IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetFieldReceivedAt              IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetField = "received_at"
	IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetFieldReturnDescription       IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetField = "return_description"
	IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetFieldReturnStatus            IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetField = "return_status"
	IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetFieldReturnedAt              IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetField = "returned_at"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParams) AddUnsetField(field IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'no_valid_authorization'.
type IssuingDisputeCreateEvidenceNoValidAuthorizationParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string                                                            `form:"explanation"`
	UnsetFields []IssuingDisputeCreateEvidenceNoValidAuthorizationParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeCreateEvidenceNoValidAuthorizationParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeCreateEvidenceNoValidAuthorizationParams.
type IssuingDisputeCreateEvidenceNoValidAuthorizationParamsUnsetField string

const (
	IssuingDisputeCreateEvidenceNoValidAuthorizationParamsUnsetFieldAdditionalDocumentation IssuingDisputeCreateEvidenceNoValidAuthorizationParamsUnsetField = "additional_documentation"
	IssuingDisputeCreateEvidenceNoValidAuthorizationParamsUnsetFieldExplanation             IssuingDisputeCreateEvidenceNoValidAuthorizationParamsUnsetField = "explanation"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeCreateEvidenceNoValidAuthorizationParams) AddUnsetField(field IssuingDisputeCreateEvidenceNoValidAuthorizationParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'not_received'.
type IssuingDisputeCreateEvidenceNotReceivedParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Date when the cardholder expected to receive the product.
	ExpectedAt *int64 `form:"expected_at"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Description of the merchandise or service that was purchased.
	ProductDescription *string `form:"product_description"`
	// Whether the product was a merchandise or service.
	ProductType *string                                                   `form:"product_type"`
	UnsetFields []IssuingDisputeCreateEvidenceNotReceivedParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeCreateEvidenceNotReceivedParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeCreateEvidenceNotReceivedParams.
type IssuingDisputeCreateEvidenceNotReceivedParamsUnsetField string

const (
	IssuingDisputeCreateEvidenceNotReceivedParamsUnsetFieldAdditionalDocumentation IssuingDisputeCreateEvidenceNotReceivedParamsUnsetField = "additional_documentation"
	IssuingDisputeCreateEvidenceNotReceivedParamsUnsetFieldExpectedAt              IssuingDisputeCreateEvidenceNotReceivedParamsUnsetField = "expected_at"
	IssuingDisputeCreateEvidenceNotReceivedParamsUnsetFieldExplanation             IssuingDisputeCreateEvidenceNotReceivedParamsUnsetField = "explanation"
	IssuingDisputeCreateEvidenceNotReceivedParamsUnsetFieldProductDescription      IssuingDisputeCreateEvidenceNotReceivedParamsUnsetField = "product_description"
	IssuingDisputeCreateEvidenceNotReceivedParamsUnsetFieldProductType             IssuingDisputeCreateEvidenceNotReceivedParamsUnsetField = "product_type"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeCreateEvidenceNotReceivedParams) AddUnsetField(field IssuingDisputeCreateEvidenceNotReceivedParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'other'.
type IssuingDisputeCreateEvidenceOtherParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Description of the merchandise or service that was purchased.
	ProductDescription *string `form:"product_description"`
	// Whether the product was a merchandise or service.
	ProductType *string                                             `form:"product_type"`
	UnsetFields []IssuingDisputeCreateEvidenceOtherParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeCreateEvidenceOtherParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeCreateEvidenceOtherParams.
type IssuingDisputeCreateEvidenceOtherParamsUnsetField string

const (
	IssuingDisputeCreateEvidenceOtherParamsUnsetFieldAdditionalDocumentation IssuingDisputeCreateEvidenceOtherParamsUnsetField = "additional_documentation"
	IssuingDisputeCreateEvidenceOtherParamsUnsetFieldExplanation             IssuingDisputeCreateEvidenceOtherParamsUnsetField = "explanation"
	IssuingDisputeCreateEvidenceOtherParamsUnsetFieldProductDescription      IssuingDisputeCreateEvidenceOtherParamsUnsetField = "product_description"
	IssuingDisputeCreateEvidenceOtherParamsUnsetFieldProductType             IssuingDisputeCreateEvidenceOtherParamsUnsetField = "product_type"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeCreateEvidenceOtherParams) AddUnsetField(field IssuingDisputeCreateEvidenceOtherParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'service_not_as_described'.
type IssuingDisputeCreateEvidenceServiceNotAsDescribedParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Date when order was canceled.
	CanceledAt *int64 `form:"canceled_at"`
	// Reason for canceling the order.
	CancellationReason *string `form:"cancellation_reason"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Date when the product was received.
	ReceivedAt  *int64                                                              `form:"received_at"`
	UnsetFields []IssuingDisputeCreateEvidenceServiceNotAsDescribedParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeCreateEvidenceServiceNotAsDescribedParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeCreateEvidenceServiceNotAsDescribedParams.
type IssuingDisputeCreateEvidenceServiceNotAsDescribedParamsUnsetField string

const (
	IssuingDisputeCreateEvidenceServiceNotAsDescribedParamsUnsetFieldAdditionalDocumentation IssuingDisputeCreateEvidenceServiceNotAsDescribedParamsUnsetField = "additional_documentation"
	IssuingDisputeCreateEvidenceServiceNotAsDescribedParamsUnsetFieldCanceledAt              IssuingDisputeCreateEvidenceServiceNotAsDescribedParamsUnsetField = "canceled_at"
	IssuingDisputeCreateEvidenceServiceNotAsDescribedParamsUnsetFieldCancellationReason      IssuingDisputeCreateEvidenceServiceNotAsDescribedParamsUnsetField = "cancellation_reason"
	IssuingDisputeCreateEvidenceServiceNotAsDescribedParamsUnsetFieldExplanation             IssuingDisputeCreateEvidenceServiceNotAsDescribedParamsUnsetField = "explanation"
	IssuingDisputeCreateEvidenceServiceNotAsDescribedParamsUnsetFieldReceivedAt              IssuingDisputeCreateEvidenceServiceNotAsDescribedParamsUnsetField = "received_at"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeCreateEvidenceServiceNotAsDescribedParams) AddUnsetField(field IssuingDisputeCreateEvidenceServiceNotAsDescribedParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided for the dispute.
type IssuingDisputeCreateEvidenceParams struct {
	// Evidence provided when `reason` is 'canceled'.
	Canceled *IssuingDisputeCreateEvidenceCanceledParams `form:"canceled"`
	// Evidence provided when `reason` is 'duplicate'.
	Duplicate *IssuingDisputeCreateEvidenceDuplicateParams `form:"duplicate"`
	// Evidence provided when `reason` is 'fraudulent'.
	Fraudulent *IssuingDisputeCreateEvidenceFraudulentParams `form:"fraudulent"`
	// Evidence provided when `reason` is 'merchandise_not_as_described'.
	MerchandiseNotAsDescribed *IssuingDisputeCreateEvidenceMerchandiseNotAsDescribedParams `form:"merchandise_not_as_described"`
	// Evidence provided when `reason` is 'not_received'.
	NotReceived *IssuingDisputeCreateEvidenceNotReceivedParams `form:"not_received"`
	// Evidence provided when `reason` is 'no_valid_authorization'.
	NoValidAuthorization *IssuingDisputeCreateEvidenceNoValidAuthorizationParams `form:"no_valid_authorization"`
	// Evidence provided when `reason` is 'other'.
	Other *IssuingDisputeCreateEvidenceOtherParams `form:"other"`
	// The reason for filing the dispute. The evidence should be submitted in the field of the same name.
	Reason *string `form:"reason"`
	// Evidence provided when `reason` is 'service_not_as_described'.
	ServiceNotAsDescribed *IssuingDisputeCreateEvidenceServiceNotAsDescribedParams `form:"service_not_as_described"`
	UnsetFields           []IssuingDisputeCreateEvidenceParamsUnsetField           `form:"-" json:"-"`
}

// IssuingDisputeCreateEvidenceParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeCreateEvidenceParams.
type IssuingDisputeCreateEvidenceParamsUnsetField string

const (
	IssuingDisputeCreateEvidenceParamsUnsetFieldCanceled                  IssuingDisputeCreateEvidenceParamsUnsetField = "canceled"
	IssuingDisputeCreateEvidenceParamsUnsetFieldDuplicate                 IssuingDisputeCreateEvidenceParamsUnsetField = "duplicate"
	IssuingDisputeCreateEvidenceParamsUnsetFieldFraudulent                IssuingDisputeCreateEvidenceParamsUnsetField = "fraudulent"
	IssuingDisputeCreateEvidenceParamsUnsetFieldMerchandiseNotAsDescribed IssuingDisputeCreateEvidenceParamsUnsetField = "merchandise_not_as_described"
	IssuingDisputeCreateEvidenceParamsUnsetFieldNoValidAuthorization      IssuingDisputeCreateEvidenceParamsUnsetField = "no_valid_authorization"
	IssuingDisputeCreateEvidenceParamsUnsetFieldNotReceived               IssuingDisputeCreateEvidenceParamsUnsetField = "not_received"
	IssuingDisputeCreateEvidenceParamsUnsetFieldOther                     IssuingDisputeCreateEvidenceParamsUnsetField = "other"
	IssuingDisputeCreateEvidenceParamsUnsetFieldServiceNotAsDescribed     IssuingDisputeCreateEvidenceParamsUnsetField = "service_not_as_described"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeCreateEvidenceParams) AddUnsetField(field IssuingDisputeCreateEvidenceParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Params for disputes related to Treasury FinancialAccounts
type IssuingDisputeCreateTreasuryParams struct {
	// The ID of the ReceivedDebit to initiate an Issuings dispute for.
	ReceivedDebit *string `form:"received_debit"`
}

// Creates an Issuing Dispute object. Individual pieces of evidence within the evidence object are optional at this point. Stripe only validates that required evidence is present during submission. Refer to [Dispute reasons and evidence](https://docs.stripe.com/docs/issuing/purchases/disputes#dispute-reasons-and-evidence) for more details about evidence requirements.
type IssuingDisputeCreateParams struct {
	Params `form:"*"`
	// The dispute amount in the card's currency and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal). If not set, defaults to the full transaction amount.
	Amount *int64 `form:"amount"`
	// Evidence provided for the dispute.
	Evidence *IssuingDisputeCreateEvidenceParams `form:"evidence"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// The ID of the issuing transaction to create a dispute for. For transaction on Treasury FinancialAccounts, use `treasury.received_debit`.
	Transaction *string `form:"transaction"`
	// Params for disputes related to Treasury FinancialAccounts
	Treasury *IssuingDisputeCreateTreasuryParams `form:"treasury"`
}

// AddExpand appends a new field to expand.
func (p *IssuingDisputeCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *IssuingDisputeCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves an Issuing Dispute object.
type IssuingDisputeRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *IssuingDisputeRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Evidence provided when `reason` is 'canceled'.
type IssuingDisputeUpdateEvidenceCanceledParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Date when order was canceled.
	CanceledAt *int64 `form:"canceled_at"`
	// Whether the cardholder was provided with a cancellation policy.
	CancellationPolicyProvided *bool `form:"cancellation_policy_provided"`
	// Reason for canceling the order.
	CancellationReason *string `form:"cancellation_reason"`
	// Date when the cardholder expected to receive the product.
	ExpectedAt *int64 `form:"expected_at"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Description of the merchandise or service that was purchased.
	ProductDescription *string `form:"product_description"`
	// Whether the product was a merchandise or service.
	ProductType *string `form:"product_type"`
	// Date when the product was returned or attempted to be returned.
	ReturnedAt *int64 `form:"returned_at"`
	// Result of cardholder's attempt to return the product.
	ReturnStatus *string                                                `form:"return_status"`
	UnsetFields  []IssuingDisputeUpdateEvidenceCanceledParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeUpdateEvidenceCanceledParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeUpdateEvidenceCanceledParams.
type IssuingDisputeUpdateEvidenceCanceledParamsUnsetField string

const (
	IssuingDisputeUpdateEvidenceCanceledParamsUnsetFieldAdditionalDocumentation    IssuingDisputeUpdateEvidenceCanceledParamsUnsetField = "additional_documentation"
	IssuingDisputeUpdateEvidenceCanceledParamsUnsetFieldCanceledAt                 IssuingDisputeUpdateEvidenceCanceledParamsUnsetField = "canceled_at"
	IssuingDisputeUpdateEvidenceCanceledParamsUnsetFieldCancellationPolicyProvided IssuingDisputeUpdateEvidenceCanceledParamsUnsetField = "cancellation_policy_provided"
	IssuingDisputeUpdateEvidenceCanceledParamsUnsetFieldCancellationReason         IssuingDisputeUpdateEvidenceCanceledParamsUnsetField = "cancellation_reason"
	IssuingDisputeUpdateEvidenceCanceledParamsUnsetFieldExpectedAt                 IssuingDisputeUpdateEvidenceCanceledParamsUnsetField = "expected_at"
	IssuingDisputeUpdateEvidenceCanceledParamsUnsetFieldExplanation                IssuingDisputeUpdateEvidenceCanceledParamsUnsetField = "explanation"
	IssuingDisputeUpdateEvidenceCanceledParamsUnsetFieldProductDescription         IssuingDisputeUpdateEvidenceCanceledParamsUnsetField = "product_description"
	IssuingDisputeUpdateEvidenceCanceledParamsUnsetFieldProductType                IssuingDisputeUpdateEvidenceCanceledParamsUnsetField = "product_type"
	IssuingDisputeUpdateEvidenceCanceledParamsUnsetFieldReturnStatus               IssuingDisputeUpdateEvidenceCanceledParamsUnsetField = "return_status"
	IssuingDisputeUpdateEvidenceCanceledParamsUnsetFieldReturnedAt                 IssuingDisputeUpdateEvidenceCanceledParamsUnsetField = "returned_at"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeUpdateEvidenceCanceledParams) AddUnsetField(field IssuingDisputeUpdateEvidenceCanceledParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'duplicate'.
type IssuingDisputeUpdateEvidenceDuplicateParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Copy of the card statement showing that the product had already been paid for.
	CardStatement *string `form:"card_statement"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Copy of the receipt showing that the product had been paid for in cash.
	CashReceipt *string `form:"cash_receipt"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Image of the front and back of the check that was used to pay for the product.
	CheckImage *string `form:"check_image"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Transaction (e.g., ipi_...) that the disputed transaction is a duplicate of. Of the two or more transactions that are copies of each other, this is original undisputed one.
	OriginalTransaction *string                                                 `form:"original_transaction"`
	UnsetFields         []IssuingDisputeUpdateEvidenceDuplicateParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeUpdateEvidenceDuplicateParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeUpdateEvidenceDuplicateParams.
type IssuingDisputeUpdateEvidenceDuplicateParamsUnsetField string

const (
	IssuingDisputeUpdateEvidenceDuplicateParamsUnsetFieldAdditionalDocumentation IssuingDisputeUpdateEvidenceDuplicateParamsUnsetField = "additional_documentation"
	IssuingDisputeUpdateEvidenceDuplicateParamsUnsetFieldCardStatement           IssuingDisputeUpdateEvidenceDuplicateParamsUnsetField = "card_statement"
	IssuingDisputeUpdateEvidenceDuplicateParamsUnsetFieldCashReceipt             IssuingDisputeUpdateEvidenceDuplicateParamsUnsetField = "cash_receipt"
	IssuingDisputeUpdateEvidenceDuplicateParamsUnsetFieldCheckImage              IssuingDisputeUpdateEvidenceDuplicateParamsUnsetField = "check_image"
	IssuingDisputeUpdateEvidenceDuplicateParamsUnsetFieldExplanation             IssuingDisputeUpdateEvidenceDuplicateParamsUnsetField = "explanation"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeUpdateEvidenceDuplicateParams) AddUnsetField(field IssuingDisputeUpdateEvidenceDuplicateParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'fraudulent'.
type IssuingDisputeUpdateEvidenceFraudulentParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string                                                  `form:"explanation"`
	UnsetFields []IssuingDisputeUpdateEvidenceFraudulentParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeUpdateEvidenceFraudulentParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeUpdateEvidenceFraudulentParams.
type IssuingDisputeUpdateEvidenceFraudulentParamsUnsetField string

const (
	IssuingDisputeUpdateEvidenceFraudulentParamsUnsetFieldAdditionalDocumentation IssuingDisputeUpdateEvidenceFraudulentParamsUnsetField = "additional_documentation"
	IssuingDisputeUpdateEvidenceFraudulentParamsUnsetFieldExplanation             IssuingDisputeUpdateEvidenceFraudulentParamsUnsetField = "explanation"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeUpdateEvidenceFraudulentParams) AddUnsetField(field IssuingDisputeUpdateEvidenceFraudulentParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'merchandise_not_as_described'.
type IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Date when the product was received.
	ReceivedAt *int64 `form:"received_at"`
	// Description of the cardholder's attempt to return the product.
	ReturnDescription *string `form:"return_description"`
	// Date when the product was returned or attempted to be returned.
	ReturnedAt *int64 `form:"returned_at"`
	// Result of cardholder's attempt to return the product.
	ReturnStatus *string                                                                 `form:"return_status"`
	UnsetFields  []IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParams.
type IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetField string

const (
	IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetFieldAdditionalDocumentation IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetField = "additional_documentation"
	IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetFieldExplanation             IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetField = "explanation"
	IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetFieldReceivedAt              IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetField = "received_at"
	IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetFieldReturnDescription       IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetField = "return_description"
	IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetFieldReturnStatus            IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetField = "return_status"
	IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetFieldReturnedAt              IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetField = "returned_at"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParams) AddUnsetField(field IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'no_valid_authorization'.
type IssuingDisputeUpdateEvidenceNoValidAuthorizationParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string                                                            `form:"explanation"`
	UnsetFields []IssuingDisputeUpdateEvidenceNoValidAuthorizationParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeUpdateEvidenceNoValidAuthorizationParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeUpdateEvidenceNoValidAuthorizationParams.
type IssuingDisputeUpdateEvidenceNoValidAuthorizationParamsUnsetField string

const (
	IssuingDisputeUpdateEvidenceNoValidAuthorizationParamsUnsetFieldAdditionalDocumentation IssuingDisputeUpdateEvidenceNoValidAuthorizationParamsUnsetField = "additional_documentation"
	IssuingDisputeUpdateEvidenceNoValidAuthorizationParamsUnsetFieldExplanation             IssuingDisputeUpdateEvidenceNoValidAuthorizationParamsUnsetField = "explanation"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeUpdateEvidenceNoValidAuthorizationParams) AddUnsetField(field IssuingDisputeUpdateEvidenceNoValidAuthorizationParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'not_received'.
type IssuingDisputeUpdateEvidenceNotReceivedParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Date when the cardholder expected to receive the product.
	ExpectedAt *int64 `form:"expected_at"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Description of the merchandise or service that was purchased.
	ProductDescription *string `form:"product_description"`
	// Whether the product was a merchandise or service.
	ProductType *string                                                   `form:"product_type"`
	UnsetFields []IssuingDisputeUpdateEvidenceNotReceivedParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeUpdateEvidenceNotReceivedParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeUpdateEvidenceNotReceivedParams.
type IssuingDisputeUpdateEvidenceNotReceivedParamsUnsetField string

const (
	IssuingDisputeUpdateEvidenceNotReceivedParamsUnsetFieldAdditionalDocumentation IssuingDisputeUpdateEvidenceNotReceivedParamsUnsetField = "additional_documentation"
	IssuingDisputeUpdateEvidenceNotReceivedParamsUnsetFieldExpectedAt              IssuingDisputeUpdateEvidenceNotReceivedParamsUnsetField = "expected_at"
	IssuingDisputeUpdateEvidenceNotReceivedParamsUnsetFieldExplanation             IssuingDisputeUpdateEvidenceNotReceivedParamsUnsetField = "explanation"
	IssuingDisputeUpdateEvidenceNotReceivedParamsUnsetFieldProductDescription      IssuingDisputeUpdateEvidenceNotReceivedParamsUnsetField = "product_description"
	IssuingDisputeUpdateEvidenceNotReceivedParamsUnsetFieldProductType             IssuingDisputeUpdateEvidenceNotReceivedParamsUnsetField = "product_type"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeUpdateEvidenceNotReceivedParams) AddUnsetField(field IssuingDisputeUpdateEvidenceNotReceivedParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'other'.
type IssuingDisputeUpdateEvidenceOtherParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Description of the merchandise or service that was purchased.
	ProductDescription *string `form:"product_description"`
	// Whether the product was a merchandise or service.
	ProductType *string                                             `form:"product_type"`
	UnsetFields []IssuingDisputeUpdateEvidenceOtherParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeUpdateEvidenceOtherParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeUpdateEvidenceOtherParams.
type IssuingDisputeUpdateEvidenceOtherParamsUnsetField string

const (
	IssuingDisputeUpdateEvidenceOtherParamsUnsetFieldAdditionalDocumentation IssuingDisputeUpdateEvidenceOtherParamsUnsetField = "additional_documentation"
	IssuingDisputeUpdateEvidenceOtherParamsUnsetFieldExplanation             IssuingDisputeUpdateEvidenceOtherParamsUnsetField = "explanation"
	IssuingDisputeUpdateEvidenceOtherParamsUnsetFieldProductDescription      IssuingDisputeUpdateEvidenceOtherParamsUnsetField = "product_description"
	IssuingDisputeUpdateEvidenceOtherParamsUnsetFieldProductType             IssuingDisputeUpdateEvidenceOtherParamsUnsetField = "product_type"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeUpdateEvidenceOtherParams) AddUnsetField(field IssuingDisputeUpdateEvidenceOtherParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided when `reason` is 'service_not_as_described'.
type IssuingDisputeUpdateEvidenceServiceNotAsDescribedParams struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *string `form:"additional_documentation"`
	// Date when order was canceled.
	CanceledAt *int64 `form:"canceled_at"`
	// Reason for canceling the order.
	CancellationReason *string `form:"cancellation_reason"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation *string `form:"explanation"`
	// Date when the product was received.
	ReceivedAt  *int64                                                              `form:"received_at"`
	UnsetFields []IssuingDisputeUpdateEvidenceServiceNotAsDescribedParamsUnsetField `form:"-" json:"-"`
}

// IssuingDisputeUpdateEvidenceServiceNotAsDescribedParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeUpdateEvidenceServiceNotAsDescribedParams.
type IssuingDisputeUpdateEvidenceServiceNotAsDescribedParamsUnsetField string

const (
	IssuingDisputeUpdateEvidenceServiceNotAsDescribedParamsUnsetFieldAdditionalDocumentation IssuingDisputeUpdateEvidenceServiceNotAsDescribedParamsUnsetField = "additional_documentation"
	IssuingDisputeUpdateEvidenceServiceNotAsDescribedParamsUnsetFieldCanceledAt              IssuingDisputeUpdateEvidenceServiceNotAsDescribedParamsUnsetField = "canceled_at"
	IssuingDisputeUpdateEvidenceServiceNotAsDescribedParamsUnsetFieldCancellationReason      IssuingDisputeUpdateEvidenceServiceNotAsDescribedParamsUnsetField = "cancellation_reason"
	IssuingDisputeUpdateEvidenceServiceNotAsDescribedParamsUnsetFieldExplanation             IssuingDisputeUpdateEvidenceServiceNotAsDescribedParamsUnsetField = "explanation"
	IssuingDisputeUpdateEvidenceServiceNotAsDescribedParamsUnsetFieldReceivedAt              IssuingDisputeUpdateEvidenceServiceNotAsDescribedParamsUnsetField = "received_at"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeUpdateEvidenceServiceNotAsDescribedParams) AddUnsetField(field IssuingDisputeUpdateEvidenceServiceNotAsDescribedParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Evidence provided for the dispute.
type IssuingDisputeUpdateEvidenceParams struct {
	// Evidence provided when `reason` is 'canceled'.
	Canceled *IssuingDisputeUpdateEvidenceCanceledParams `form:"canceled"`
	// Evidence provided when `reason` is 'duplicate'.
	Duplicate *IssuingDisputeUpdateEvidenceDuplicateParams `form:"duplicate"`
	// Evidence provided when `reason` is 'fraudulent'.
	Fraudulent *IssuingDisputeUpdateEvidenceFraudulentParams `form:"fraudulent"`
	// Evidence provided when `reason` is 'merchandise_not_as_described'.
	MerchandiseNotAsDescribed *IssuingDisputeUpdateEvidenceMerchandiseNotAsDescribedParams `form:"merchandise_not_as_described"`
	// Evidence provided when `reason` is 'not_received'.
	NotReceived *IssuingDisputeUpdateEvidenceNotReceivedParams `form:"not_received"`
	// Evidence provided when `reason` is 'no_valid_authorization'.
	NoValidAuthorization *IssuingDisputeUpdateEvidenceNoValidAuthorizationParams `form:"no_valid_authorization"`
	// Evidence provided when `reason` is 'other'.
	Other *IssuingDisputeUpdateEvidenceOtherParams `form:"other"`
	// The reason for filing the dispute. The evidence should be submitted in the field of the same name.
	Reason *string `form:"reason"`
	// Evidence provided when `reason` is 'service_not_as_described'.
	ServiceNotAsDescribed *IssuingDisputeUpdateEvidenceServiceNotAsDescribedParams `form:"service_not_as_described"`
	UnsetFields           []IssuingDisputeUpdateEvidenceParamsUnsetField           `form:"-" json:"-"`
}

// IssuingDisputeUpdateEvidenceParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeUpdateEvidenceParams.
type IssuingDisputeUpdateEvidenceParamsUnsetField string

const (
	IssuingDisputeUpdateEvidenceParamsUnsetFieldCanceled                  IssuingDisputeUpdateEvidenceParamsUnsetField = "canceled"
	IssuingDisputeUpdateEvidenceParamsUnsetFieldDuplicate                 IssuingDisputeUpdateEvidenceParamsUnsetField = "duplicate"
	IssuingDisputeUpdateEvidenceParamsUnsetFieldFraudulent                IssuingDisputeUpdateEvidenceParamsUnsetField = "fraudulent"
	IssuingDisputeUpdateEvidenceParamsUnsetFieldMerchandiseNotAsDescribed IssuingDisputeUpdateEvidenceParamsUnsetField = "merchandise_not_as_described"
	IssuingDisputeUpdateEvidenceParamsUnsetFieldNoValidAuthorization      IssuingDisputeUpdateEvidenceParamsUnsetField = "no_valid_authorization"
	IssuingDisputeUpdateEvidenceParamsUnsetFieldNotReceived               IssuingDisputeUpdateEvidenceParamsUnsetField = "not_received"
	IssuingDisputeUpdateEvidenceParamsUnsetFieldOther                     IssuingDisputeUpdateEvidenceParamsUnsetField = "other"
	IssuingDisputeUpdateEvidenceParamsUnsetFieldServiceNotAsDescribed     IssuingDisputeUpdateEvidenceParamsUnsetField = "service_not_as_described"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeUpdateEvidenceParams) AddUnsetField(field IssuingDisputeUpdateEvidenceParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Updates the specified Issuing Dispute object by setting the values of the parameters passed. Any parameters not provided will be left unchanged. Properties on the evidence object can be unset by passing in an empty string.
type IssuingDisputeUpdateParams struct {
	Params `form:"*"`
	// The dispute amount in the card's currency and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal).
	Amount *int64 `form:"amount"`
	// Evidence provided for the dispute.
	Evidence *IssuingDisputeUpdateEvidenceParams `form:"evidence"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
}

// IssuingDisputeUpdateParamsUnsetField is the list of fields that can be cleared/unset on IssuingDisputeUpdateParams.
type IssuingDisputeUpdateParamsUnsetField string

const (
	IssuingDisputeUpdateParamsUnsetFieldMetadata IssuingDisputeUpdateParamsUnsetField = "metadata"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *IssuingDisputeUpdateParams) AddUnsetField(field IssuingDisputeUpdateParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, string(field))
}

// AddExpand appends a new field to expand.
func (p *IssuingDisputeUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *IssuingDisputeUpdateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

type IssuingDisputeEvidenceCanceled struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *File `json:"additional_documentation"`
	// Date when order was canceled.
	CanceledAt int64 `json:"canceled_at"`
	// Whether the cardholder was provided with a cancellation policy.
	CancellationPolicyProvided bool `json:"cancellation_policy_provided"`
	// Reason for canceling the order.
	CancellationReason string `json:"cancellation_reason"`
	// Date when the cardholder expected to receive the product.
	ExpectedAt int64 `json:"expected_at"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation string `json:"explanation"`
	// Description of the merchandise or service that was purchased.
	ProductDescription string `json:"product_description"`
	// Whether the product was a merchandise or service.
	ProductType IssuingDisputeEvidenceCanceledProductType `json:"product_type"`
	// Date when the product was returned or attempted to be returned.
	ReturnedAt int64 `json:"returned_at"`
	// Result of cardholder's attempt to return the product.
	ReturnStatus IssuingDisputeEvidenceCanceledReturnStatus `json:"return_status"`
}
type IssuingDisputeEvidenceDuplicate struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *File `json:"additional_documentation"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Copy of the card statement showing that the product had already been paid for.
	CardStatement *File `json:"card_statement"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Copy of the receipt showing that the product had been paid for in cash.
	CashReceipt *File `json:"cash_receipt"`
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Image of the front and back of the check that was used to pay for the product.
	CheckImage *File `json:"check_image"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation string `json:"explanation"`
	// Transaction (e.g., ipi_...) that the disputed transaction is a duplicate of. Of the two or more transactions that are copies of each other, this is original undisputed one.
	OriginalTransaction string `json:"original_transaction"`
}
type IssuingDisputeEvidenceFraudulent struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *File `json:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation string `json:"explanation"`
}
type IssuingDisputeEvidenceMerchandiseNotAsDescribed struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *File `json:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation string `json:"explanation"`
	// Date when the product was received.
	ReceivedAt int64 `json:"received_at"`
	// Description of the cardholder's attempt to return the product.
	ReturnDescription string `json:"return_description"`
	// Date when the product was returned or attempted to be returned.
	ReturnedAt int64 `json:"returned_at"`
	// Result of cardholder's attempt to return the product.
	ReturnStatus IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatus `json:"return_status"`
}
type IssuingDisputeEvidenceNoValidAuthorization struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *File `json:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation string `json:"explanation"`
}
type IssuingDisputeEvidenceNotReceived struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *File `json:"additional_documentation"`
	// Date when the cardholder expected to receive the product.
	ExpectedAt int64 `json:"expected_at"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation string `json:"explanation"`
	// Description of the merchandise or service that was purchased.
	ProductDescription string `json:"product_description"`
	// Whether the product was a merchandise or service.
	ProductType IssuingDisputeEvidenceNotReceivedProductType `json:"product_type"`
}
type IssuingDisputeEvidenceOther struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *File `json:"additional_documentation"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation string `json:"explanation"`
	// Description of the merchandise or service that was purchased.
	ProductDescription string `json:"product_description"`
	// Whether the product was a merchandise or service.
	ProductType IssuingDisputeEvidenceOtherProductType `json:"product_type"`
}
type IssuingDisputeEvidenceServiceNotAsDescribed struct {
	// (ID of a [file upload](https://stripe.com/docs/guides/file-upload)) Additional documentation supporting the dispute.
	AdditionalDocumentation *File `json:"additional_documentation"`
	// Date when order was canceled.
	CanceledAt int64 `json:"canceled_at"`
	// Reason for canceling the order.
	CancellationReason string `json:"cancellation_reason"`
	// Explanation of why the cardholder is disputing this transaction.
	Explanation string `json:"explanation"`
	// Date when the product was received.
	ReceivedAt int64 `json:"received_at"`
}
type IssuingDisputeEvidence struct {
	Canceled                  *IssuingDisputeEvidenceCanceled                  `json:"canceled"`
	Duplicate                 *IssuingDisputeEvidenceDuplicate                 `json:"duplicate"`
	Fraudulent                *IssuingDisputeEvidenceFraudulent                `json:"fraudulent"`
	MerchandiseNotAsDescribed *IssuingDisputeEvidenceMerchandiseNotAsDescribed `json:"merchandise_not_as_described"`
	NotReceived               *IssuingDisputeEvidenceNotReceived               `json:"not_received"`
	NoValidAuthorization      *IssuingDisputeEvidenceNoValidAuthorization      `json:"no_valid_authorization"`
	Other                     *IssuingDisputeEvidenceOther                     `json:"other"`
	// The reason for filing the dispute. Its value will match the field containing the evidence.
	Reason                IssuingDisputeEvidenceReason                 `json:"reason"`
	ServiceNotAsDescribed *IssuingDisputeEvidenceServiceNotAsDescribed `json:"service_not_as_described"`
}

// [Treasury](https://docs.stripe.com/api/treasury) details related to this dispute if it was created on a [FinancialAccount](/docs/api/treasury/financial_accounts
type IssuingDisputeTreasury struct {
	// The Treasury [DebitReversal](https://docs.stripe.com/api/treasury/debit_reversals) representing this Issuing dispute
	DebitReversal string `json:"debit_reversal"`
	// The Treasury [ReceivedDebit](https://docs.stripe.com/api/treasury/received_debits) that is being disputed.
	ReceivedDebit string `json:"received_debit"`
}

// As a [card issuer](https://docs.stripe.com/issuing), you can dispute transactions that the cardholder does not recognize, suspects to be fraudulent, or has other issues with.
//
// Related guide: [Issuing disputes](https://docs.stripe.com/issuing/purchases/disputes)
type IssuingDispute struct {
	APIResource
	// Disputed amount in the card's currency and in the [smallest currency unit](https://docs.stripe.com/currencies#zero-decimal). Usually the amount of the `transaction`, but can differ (usually because of currency fluctuation).
	Amount int64 `json:"amount"`
	// List of balance transactions associated with the dispute.
	BalanceTransactions []*BalanceTransaction `json:"balance_transactions"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The currency the `transaction` was made in.
	Currency Currency                `json:"currency"`
	Evidence *IssuingDisputeEvidence `json:"evidence"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// The enum that describes the dispute loss outcome. If the dispute is not lost, this field will be absent. New enum values may be added in the future, so be sure to handle unknown values.
	LossReason IssuingDisputeLossReason `json:"loss_reason"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Current status of the dispute.
	Status IssuingDisputeStatus `json:"status"`
	// The transaction being disputed.
	Transaction *IssuingTransaction `json:"transaction"`
	// [Treasury](https://docs.stripe.com/api/treasury) details related to this dispute if it was created on a [FinancialAccount](/docs/api/treasury/financial_accounts
	Treasury *IssuingDisputeTreasury `json:"treasury"`
}

// IssuingDisputeList is a list of Disputes as retrieved from a list endpoint.
type IssuingDisputeList struct {
	APIResource
	ListMeta
	Data []*IssuingDispute `json:"data"`
}

// UnmarshalJSON handles deserialization of an IssuingDispute.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *IssuingDispute) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		i.ID = id
		return nil
	}

	type issuingDispute IssuingDispute
	var v issuingDispute
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*i = IssuingDispute(v)
	return nil
}
