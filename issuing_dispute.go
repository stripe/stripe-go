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

// IssuingDisputeEvidenceServiceNotAsDescribedProductType is the list of allowed product types on an issuing dispute of type service not as described.
type IssuingDisputeEvidenceServiceNotAsDescribedProductType string

// List of values that IssuingDisputeEvidenceServiceNotAsDescribedProductType can take.
const (
	IssuingDisputeEvidenceServiceNotAsDescribedProductTypeMerchandise IssuingDisputeEvidenceServiceNotAsDescribedProductType = "merchandise"
	IssuingDisputeEvidenceServiceNotAsDescribedProductTypeService     IssuingDisputeEvidenceServiceNotAsDescribedProductType = "service"
)

// The reason for filing the dispute. Its value will match the field containing the evidence.
type IssuingDisputeEvidenceReason string

// List of values that IssuingDisputeEvidenceReason can take
const (
	IssuingDisputeEvidenceReasonCanceled                  IssuingDisputeEvidenceReason = "canceled"
	IssuingDisputeEvidenceReasonDuplicate                 IssuingDisputeEvidenceReason = "duplicate"
	IssuingDisputeEvidenceReasonFraudulent                IssuingDisputeEvidenceReason = "fraudulent"
	IssuingDisputeEvidenceReasonMerchandiseNotAsDescribed IssuingDisputeEvidenceReason = "merchandise_not_as_described"
	IssuingDisputeEvidenceReasonNotReceived               IssuingDisputeEvidenceReason = "not_received"
	IssuingDisputeEvidenceReasonOther                     IssuingDisputeEvidenceReason = "other"
	IssuingDisputeEvidenceReasonServiceNotAsDescribed     IssuingDisputeEvidenceReason = "service_not_as_described"
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
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Status       *string           `form:"status"`
	Transaction  *string           `form:"transaction"`
}

// Evidence provided when `reason` is 'canceled'.
type IssuingDisputeEvidenceCanceledParams struct {
	AdditionalDocumentation    *string `form:"additional_documentation"`
	CanceledAt                 *int64  `form:"canceled_at"`
	CancellationPolicyProvided *bool   `form:"cancellation_policy_provided"`
	CancellationReason         *string `form:"cancellation_reason"`
	ExpectedAt                 *int64  `form:"expected_at"`
	Explanation                *string `form:"explanation"`
	ProductDescription         *string `form:"product_description"`
	ProductType                *string `form:"product_type"`
	ReturnedAt                 *int64  `form:"returned_at"`
	ReturnStatus               *string `form:"return_status"`
}

// Evidence provided when `reason` is 'duplicate'.
type IssuingDisputeEvidenceDuplicateParams struct {
	AdditionalDocumentation *string `form:"additional_documentation"`
	CardStatement           *string `form:"card_statement"`
	CashReceipt             *string `form:"cash_receipt"`
	CheckImage              *string `form:"check_image"`
	Explanation             *string `form:"explanation"`
	OriginalTransaction     *string `form:"original_transaction"`
}

// Evidence provided when `reason` is 'fraudulent'.
type IssuingDisputeEvidenceFraudulentParams struct {
	AdditionalDocumentation *string `form:"additional_documentation"`
	Explanation             *string `form:"explanation"`
}

// Evidence provided when `reason` is 'merchandise_not_as_described'.
type IssuingDisputeEvidenceMerchandiseNotAsDescribedParams struct {
	AdditionalDocumentation *string `form:"additional_documentation"`
	Explanation             *string `form:"explanation"`
	ReceivedAt              *int64  `form:"received_at"`
	ReturnDescription       *string `form:"return_description"`
	ReturnedAt              *int64  `form:"returned_at"`
	ReturnStatus            *string `form:"return_status"`
}

// Evidence provided when `reason` is 'not_received'.
type IssuingDisputeEvidenceNotReceivedParams struct {
	AdditionalDocumentation *string `form:"additional_documentation"`
	ExpectedAt              *int64  `form:"expected_at"`
	Explanation             *string `form:"explanation"`
	ProductDescription      *string `form:"product_description"`
	ProductType             *string `form:"product_type"`
}

// Evidence provided when `reason` is 'other'.
type IssuingDisputeEvidenceOtherParams struct {
	AdditionalDocumentation *string `form:"additional_documentation"`
	Explanation             *string `form:"explanation"`
	ProductDescription      *string `form:"product_description"`
	ProductType             *string `form:"product_type"`
}

// Evidence provided when `reason` is 'service_not_as_described'.
type IssuingDisputeEvidenceServiceNotAsDescribedParams struct {
	AdditionalDocumentation *string `form:"additional_documentation"`
	CanceledAt              *int64  `form:"canceled_at"`
	CancellationReason      *string `form:"cancellation_reason"`
	Explanation             *string `form:"explanation"`
	ProductDescription      *string `form:"product_description"`
	ProductType             *string `form:"product_type"`
	ReceivedAt              *int64  `form:"received_at"`
}

// Evidence provided for the dispute.
type IssuingDisputeEvidenceParams struct {
	Canceled                  *IssuingDisputeEvidenceCanceledParams                  `form:"canceled"`
	Duplicate                 *IssuingDisputeEvidenceDuplicateParams                 `form:"duplicate"`
	Fraudulent                *IssuingDisputeEvidenceFraudulentParams                `form:"fraudulent"`
	MerchandiseNotAsDescribed *IssuingDisputeEvidenceMerchandiseNotAsDescribedParams `form:"merchandise_not_as_described"`
	NotReceived               *IssuingDisputeEvidenceNotReceivedParams               `form:"not_received"`
	Other                     *IssuingDisputeEvidenceOtherParams                     `form:"other"`
	Reason                    *string                                                `form:"reason"`
	ServiceNotAsDescribed     *IssuingDisputeEvidenceServiceNotAsDescribedParams     `form:"service_not_as_described"`
}

// Creates an Issuing Dispute object. Individual pieces of evidence within the evidence object are optional at this point. Stripe only validates that required evidence is present during submission. Refer to [Dispute reasons and evidence](https://stripe.com/docs/issuing/purchases/disputes#dispute-reasons-and-evidence) for more details about evidence requirements.
type IssuingDisputeParams struct {
	Params      `form:"*"`
	Evidence    *IssuingDisputeEvidenceParams `form:"evidence"`
	Transaction *string                       `form:"transaction"`
}

// Submits an Issuing Dispute to the card network. Stripe validates that all evidence fields required for the dispute's reason are present. For more details, see [Dispute reasons and evidence](https://stripe.com/docs/issuing/purchases/disputes#dispute-reasons-and-evidence).
type IssuingDisputeSubmitParams struct {
	Params `form:"*"`
}
type IssuingDisputeEvidenceCanceled struct {
	AdditionalDocumentation    *File                                      `json:"additional_documentation"`
	CanceledAt                 int64                                      `json:"canceled_at"`
	CancellationPolicyProvided bool                                       `json:"cancellation_policy_provided"`
	CancellationReason         string                                     `json:"cancellation_reason"`
	ExpectedAt                 int64                                      `json:"expected_at"`
	Explanation                string                                     `json:"explanation"`
	ProductDescription         string                                     `json:"product_description"`
	ProductType                IssuingDisputeEvidenceCanceledProductType  `json:"product_type"`
	ReturnedAt                 int64                                      `json:"returned_at"`
	ReturnStatus               IssuingDisputeEvidenceCanceledReturnStatus `json:"return_status"`
}
type IssuingDisputeEvidenceDuplicate struct {
	AdditionalDocumentation *File  `json:"additional_documentation"`
	CardStatement           *File  `json:"card_statement"`
	CashReceipt             *File  `json:"cash_receipt"`
	CheckImage              *File  `json:"check_image"`
	Explanation             string `json:"explanation"`
	OriginalTransaction     string `json:"original_transaction"`
}
type IssuingDisputeEvidenceFraudulent struct {
	AdditionalDocumentation *File  `json:"additional_documentation"`
	Explanation             string `json:"explanation"`
}
type IssuingDisputeEvidenceMerchandiseNotAsDescribed struct {
	AdditionalDocumentation *File                                                       `json:"additional_documentation"`
	Explanation             string                                                      `json:"explanation"`
	ReceivedAt              int64                                                       `json:"received_at"`
	ReturnDescription       string                                                      `json:"return_description"`
	ReturnedAt              int64                                                       `json:"returned_at"`
	ReturnStatus            IssuingDisputeEvidenceMerchandiseNotAsDescribedReturnStatus `json:"return_status"`
}
type IssuingDisputeEvidenceNotReceived struct {
	AdditionalDocumentation *File                                        `json:"additional_documentation"`
	ExpectedAt              int64                                        `json:"expected_at"`
	Explanation             string                                       `json:"explanation"`
	ProductDescription      string                                       `json:"product_description"`
	ProductType             IssuingDisputeEvidenceNotReceivedProductType `json:"product_type"`
}
type IssuingDisputeEvidenceOther struct {
	AdditionalDocumentation *File                                  `json:"additional_documentation"`
	Explanation             string                                 `json:"explanation"`
	ProductDescription      string                                 `json:"product_description"`
	ProductType             IssuingDisputeEvidenceOtherProductType `json:"product_type"`
}
type IssuingDisputeEvidenceServiceNotAsDescribed struct {
	AdditionalDocumentation *File                                                  `json:"additional_documentation"`
	CanceledAt              int64                                                  `json:"canceled_at"`
	CancellationReason      string                                                 `json:"cancellation_reason"`
	Explanation             string                                                 `json:"explanation"`
	ProductDescription      string                                                 `json:"product_description"`
	ProductType             IssuingDisputeEvidenceServiceNotAsDescribedProductType `json:"product_type"`
	ReceivedAt              int64                                                  `json:"received_at"`
}
type IssuingDisputeEvidence struct {
	Canceled                  *IssuingDisputeEvidenceCanceled                  `json:"canceled"`
	Duplicate                 *IssuingDisputeEvidenceDuplicate                 `json:"duplicate"`
	Fraudulent                *IssuingDisputeEvidenceFraudulent                `json:"fraudulent"`
	MerchandiseNotAsDescribed *IssuingDisputeEvidenceMerchandiseNotAsDescribed `json:"merchandise_not_as_described"`
	NotReceived               *IssuingDisputeEvidenceNotReceived               `json:"not_received"`
	Other                     *IssuingDisputeEvidenceOther                     `json:"other"`
	Reason                    IssuingDisputeEvidenceReason                     `json:"reason"`
	ServiceNotAsDescribed     *IssuingDisputeEvidenceServiceNotAsDescribed     `json:"service_not_as_described"`
}

// As a [card issuer](https://stripe.com/docs/issuing), you can dispute transactions that the cardholder does not recognize, suspects to be fraudulent, or has other issues with.
//
// Related guide: [Disputing Transactions](https://stripe.com/docs/issuing/purchases/disputes)
type IssuingDispute struct {
	APIResource
	Amount              int64                   `json:"amount"`
	BalanceTransactions []*BalanceTransaction   `json:"balance_transactions"`
	Created             int64                   `json:"created"`
	Currency            Currency                `json:"currency"`
	Evidence            *IssuingDisputeEvidence `json:"evidence"`
	ID                  string                  `json:"id"`
	Livemode            bool                    `json:"livemode"`
	Metadata            map[string]string       `json:"metadata"`
	Object              string                  `json:"object"`
	Status              *IssuingDisputeStatus   `json:"status"`
	Transaction         *IssuingTransaction     `json:"transaction"`
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
