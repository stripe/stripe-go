//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The current status of the Form.
type RiskFormStatus string

// List of values that RiskFormStatus can take
const (
	RiskFormStatusClosed RiskFormStatus = "closed"
	RiskFormStatusOpen   RiskFormStatus = "open"
)

// The type of information requested by Stripe.
type RiskFormType string

// List of values that RiskFormType can take
const (
	RiskFormTypeAppeal                 RiskFormType = "appeal"
	RiskFormTypeAuthorizationDocuments RiskFormType = "authorization_documents"
	RiskFormTypePaymentProcessing      RiskFormType = "payment_processing"
	RiskFormTypeProductRemoval         RiskFormType = "product_removal"
	RiskFormTypeReviewedFields         RiskFormType = "reviewed_fields"
	RiskFormTypeSelfFundingContext     RiskFormType = "self_funding_context"
	RiskFormTypeTransactionContext     RiskFormType = "transaction_context"
)

// Retrieves a risk Form by ID.
//
// # Error Codes
//
// | Code | HTTP status | Description |
// | — | — | — |
// | form_not_found | 400 | No Form matches the supplied ID. |
type RiskFormParams struct {
	Params `form:"*"`
	// Provide this for Forms of type `appeal`.
	Appeal *RiskFormAppealParams `form:"appeal" json:"appeal,omitempty"`
	// Provide this for Forms of type `authorization_documents`.
	AuthorizationDocuments *RiskFormAuthorizationDocumentsParams `form:"authorization_documents" json:"authorization_documents,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Provide this for Forms of type `payment_processing`.
	PaymentProcessing *RiskFormPaymentProcessingParams `form:"payment_processing" json:"payment_processing,omitempty"`
	// Provide this for Forms of type `product_removal`.
	ProductRemoval *RiskFormProductRemovalParams `form:"product_removal" json:"product_removal,omitempty"`
	// Provide this for Forms of type `reviewed_fields`.
	ReviewedFields *RiskFormReviewedFieldsParams `form:"reviewed_fields" json:"reviewed_fields,omitempty"`
	// Provide this for Forms of type `self_funding_context`.
	SelfFundingContext *RiskFormSelfFundingContextParams `form:"self_funding_context" json:"self_funding_context,omitempty"`
	// Provide this for Forms of type `transaction_context`.
	TransactionContext *RiskFormTransactionContextParams `form:"transaction_context" json:"transaction_context,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *RiskFormParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Provide this for Forms of type `appeal`.
type RiskFormAppealParams struct {
	// A text explanation for the appeal.
	Explanation *string `form:"explanation" json:"explanation"`
}

// Provide this for Forms of type `authorization_documents`.
type RiskFormAuthorizationDocumentsParams struct {
	// IDs of uploaded files to attach as authorization documents.
	Files []*string `form:"files" json:"files"`
}

// Files evidencing payments accepted outside Stripe. Each entry carries a tag naming the document type so a reviewer can distinguish statement types without opening the file.
type RiskFormPaymentProcessingMultiProcessorFileParams struct {
	// ID of the uploaded file. Upload using /v1/files before referencing it here.
	File *string `form:"file" json:"file"`
	// What kind of document this is. One of alt_financial_statements, recent_bank_statements, or payment_receipt.
	Tag *string `form:"tag" json:"tag,omitempty"`
}

// Payments the connected account accepts outside Stripe.
type RiskFormPaymentProcessingMultiProcessorParams struct {
	// Files evidencing payments accepted outside Stripe. Each entry carries a tag naming the document type so a reviewer can distinguish statement types without opening the file.
	Files []*RiskFormPaymentProcessingMultiProcessorFileParams `form:"files" json:"files,omitempty"`
	// Whether the connected account accepts payments outside Stripe.
	Value *string `form:"value" json:"value,omitempty"`
}

// Provide this for Forms of type `payment_processing`.
type RiskFormPaymentProcessingParams struct {
	// Payments the connected account accepts outside Stripe.
	MultiProcessor *RiskFormPaymentProcessingMultiProcessorParams `form:"multi_processor" json:"multi_processor,omitempty"`
	// How the connected account uses Stripe.
	Purpose *string `form:"purpose" json:"purpose,omitempty"`
	// Why the connected account chose Stripe.
	StripeSelectionRationale *string `form:"stripe_selection_rationale" json:"stripe_selection_rationale,omitempty"`
}

// Provide this for Forms of type `product_removal`.
type RiskFormProductRemovalParams struct {
	// The Unix timestamp when the prohibited items were removed.
	ItemsRemovedAt *int64 `form:"items_removed_at" json:"items_removed_at"`
}

// The connected account's attestation that the named account field values are accurate. Update inaccurate values through the Accounts API before attesting.
type RiskFormReviewedFieldsAttestationParams struct {
	// The Unix timestamp when the connected account attested to these fields.
	AttestedAt *int64 `form:"attested_at" json:"attested_at,omitempty"`
	// Account field paths in the currently accepted v2 Accounts API format, for example `defaults.profile.product_description`.
	Fields []*string `form:"fields" json:"fields,omitempty"`
}

// Provide this for Forms of type `reviewed_fields`.
type RiskFormReviewedFieldsParams struct {
	// The connected account's attestation that the named account field values are accurate. Update inaccurate values through the Accounts API before attesting.
	Attestation *RiskFormReviewedFieldsAttestationParams `form:"attestation" json:"attestation"`
}

// The connected account's source of wealth.
type RiskFormSelfFundingContextSourceOfWealthParams struct {
	// A text explanation of the connected account's source of wealth.
	Description *string `form:"description" json:"description,omitempty"`
	// IDs of uploaded files evidencing the source of wealth.
	Files []*string `form:"files" json:"files,omitempty"`
}

// Provide this for Forms of type `self_funding_context`.
type RiskFormSelfFundingContextParams struct {
	// The connected account's source of wealth.
	SourceOfWealth *RiskFormSelfFundingContextSourceOfWealthParams `form:"source_of_wealth" json:"source_of_wealth,omitempty"`
}

// Where the funds for the transaction came from.
type RiskFormTransactionContextSourceOfFundsParams struct {
	// A text explanation of where the funds for the transaction came from.
	Description *string `form:"description" json:"description,omitempty"`
	// IDs of uploaded files evidencing the source of funds.
	Files []*string `form:"files" json:"files,omitempty"`
}

// Provide this for Forms of type `transaction_context`.
type RiskFormTransactionContextParams struct {
	// How often the connected account expects transactions like this one. One of one_time_payment, recurring_subscription, repeat_customer, split_payment, or something_else.
	Frequency *string `form:"frequency" json:"frequency,omitempty"`
	// The purpose of the transaction.
	Purpose *string `form:"purpose" json:"purpose,omitempty"`
	// IDs of uploaded files evidencing the purpose of the transaction.
	PurposeFiles []*string `form:"purpose_files" json:"purpose_files,omitempty"`
	// The connected account's relationship to the sender.
	SenderRelationship *string `form:"sender_relationship" json:"sender_relationship,omitempty"`
	// Where the funds for the transaction came from.
	SourceOfFunds *RiskFormTransactionContextSourceOfFundsParams `form:"source_of_funds" json:"source_of_funds,omitempty"`
	// How the funds from the transaction will be used.
	UseOfFunds *string `form:"use_of_funds" json:"use_of_funds,omitempty"`
}

// Retrieves a risk Form by ID.
//
// # Error Codes
//
// | Code | HTTP status | Description |
// | — | — | — |
// | form_not_found | 400 | No Form matches the supplied ID. |
type RiskFormRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *RiskFormRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Provide this for Forms of type `appeal`.
type RiskFormUpdateAppealParams struct {
	// A text explanation for the appeal.
	Explanation *string `form:"explanation" json:"explanation"`
}

// Provide this for Forms of type `authorization_documents`.
type RiskFormUpdateAuthorizationDocumentsParams struct {
	// IDs of uploaded files to attach as authorization documents.
	Files []*string `form:"files" json:"files"`
}

// Files evidencing payments accepted outside Stripe. Each entry carries a tag naming the document type so a reviewer can distinguish statement types without opening the file.
type RiskFormUpdatePaymentProcessingMultiProcessorFileParams struct {
	// ID of the uploaded file. Upload using /v1/files before referencing it here.
	File *string `form:"file" json:"file"`
	// What kind of document this is. One of alt_financial_statements, recent_bank_statements, or payment_receipt.
	Tag *string `form:"tag" json:"tag,omitempty"`
}

// Payments the connected account accepts outside Stripe.
type RiskFormUpdatePaymentProcessingMultiProcessorParams struct {
	// Files evidencing payments accepted outside Stripe. Each entry carries a tag naming the document type so a reviewer can distinguish statement types without opening the file.
	Files []*RiskFormUpdatePaymentProcessingMultiProcessorFileParams `form:"files" json:"files,omitempty"`
	// Whether the connected account accepts payments outside Stripe.
	Value *string `form:"value" json:"value,omitempty"`
}

// Provide this for Forms of type `payment_processing`.
type RiskFormUpdatePaymentProcessingParams struct {
	// Payments the connected account accepts outside Stripe.
	MultiProcessor *RiskFormUpdatePaymentProcessingMultiProcessorParams `form:"multi_processor" json:"multi_processor,omitempty"`
	// How the connected account uses Stripe.
	Purpose *string `form:"purpose" json:"purpose,omitempty"`
	// Why the connected account chose Stripe.
	StripeSelectionRationale *string `form:"stripe_selection_rationale" json:"stripe_selection_rationale,omitempty"`
}

// Provide this for Forms of type `product_removal`.
type RiskFormUpdateProductRemovalParams struct {
	// The Unix timestamp when the prohibited items were removed.
	ItemsRemovedAt *int64 `form:"items_removed_at" json:"items_removed_at"`
}

// The connected account's attestation that the named account field values are accurate. Update inaccurate values through the Accounts API before attesting.
type RiskFormUpdateReviewedFieldsAttestationParams struct {
	// The Unix timestamp when the connected account attested to these fields.
	AttestedAt *int64 `form:"attested_at" json:"attested_at,omitempty"`
	// Account field paths in the currently accepted v2 Accounts API format, for example `defaults.profile.product_description`.
	Fields []*string `form:"fields" json:"fields,omitempty"`
}

// Provide this for Forms of type `reviewed_fields`.
type RiskFormUpdateReviewedFieldsParams struct {
	// The connected account's attestation that the named account field values are accurate. Update inaccurate values through the Accounts API before attesting.
	Attestation *RiskFormUpdateReviewedFieldsAttestationParams `form:"attestation" json:"attestation"`
}

// The connected account's source of wealth.
type RiskFormUpdateSelfFundingContextSourceOfWealthParams struct {
	// A text explanation of the connected account's source of wealth.
	Description *string `form:"description" json:"description,omitempty"`
	// IDs of uploaded files evidencing the source of wealth.
	Files []*string `form:"files" json:"files,omitempty"`
}

// Provide this for Forms of type `self_funding_context`.
type RiskFormUpdateSelfFundingContextParams struct {
	// The connected account's source of wealth.
	SourceOfWealth *RiskFormUpdateSelfFundingContextSourceOfWealthParams `form:"source_of_wealth" json:"source_of_wealth,omitempty"`
}

// Where the funds for the transaction came from.
type RiskFormUpdateTransactionContextSourceOfFundsParams struct {
	// A text explanation of where the funds for the transaction came from.
	Description *string `form:"description" json:"description,omitempty"`
	// IDs of uploaded files evidencing the source of funds.
	Files []*string `form:"files" json:"files,omitempty"`
}

// Provide this for Forms of type `transaction_context`.
type RiskFormUpdateTransactionContextParams struct {
	// How often the connected account expects transactions like this one. One of one_time_payment, recurring_subscription, repeat_customer, split_payment, or something_else.
	Frequency *string `form:"frequency" json:"frequency,omitempty"`
	// The purpose of the transaction.
	Purpose *string `form:"purpose" json:"purpose,omitempty"`
	// IDs of uploaded files evidencing the purpose of the transaction.
	PurposeFiles []*string `form:"purpose_files" json:"purpose_files,omitempty"`
	// The connected account's relationship to the sender.
	SenderRelationship *string `form:"sender_relationship" json:"sender_relationship,omitempty"`
	// Where the funds for the transaction came from.
	SourceOfFunds *RiskFormUpdateTransactionContextSourceOfFundsParams `form:"source_of_funds" json:"source_of_funds,omitempty"`
	// How the funds from the transaction will be used.
	UseOfFunds *string `form:"use_of_funds" json:"use_of_funds,omitempty"`
}

// Updates a risk Form with the connected account's response.
//
// # Error Codes
//
// | Code | HTTP status | Description |
// | — | — | — |
// | form_not_found | 400 | No Form matches the supplied ID. |
// | form_closed | 400 | The Form is closed and cannot accept updates. |
// | invalid_field_for_form_type | 400 | A submitted field is not supported by this Form type. |
// | form_field_already_submitted | 400 | A submitted field has already been submitted. |
// | form_insufficient_fields | 400 | Required fields are missing for this Form type. |
// | resource_missing | 404 | A submitted file does not exist or is not accessible to the requesting account. |
type RiskFormUpdateParams struct {
	Params `form:"*"`
	// Provide this for Forms of type `appeal`.
	Appeal *RiskFormUpdateAppealParams `form:"appeal" json:"appeal,omitempty"`
	// Provide this for Forms of type `authorization_documents`.
	AuthorizationDocuments *RiskFormUpdateAuthorizationDocumentsParams `form:"authorization_documents" json:"authorization_documents,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Provide this for Forms of type `payment_processing`.
	PaymentProcessing *RiskFormUpdatePaymentProcessingParams `form:"payment_processing" json:"payment_processing,omitempty"`
	// Provide this for Forms of type `product_removal`.
	ProductRemoval *RiskFormUpdateProductRemovalParams `form:"product_removal" json:"product_removal,omitempty"`
	// Provide this for Forms of type `reviewed_fields`.
	ReviewedFields *RiskFormUpdateReviewedFieldsParams `form:"reviewed_fields" json:"reviewed_fields,omitempty"`
	// Provide this for Forms of type `self_funding_context`.
	SelfFundingContext *RiskFormUpdateSelfFundingContextParams `form:"self_funding_context" json:"self_funding_context,omitempty"`
	// Provide this for Forms of type `transaction_context`.
	TransactionContext *RiskFormUpdateTransactionContextParams `form:"transaction_context" json:"transaction_context,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *RiskFormUpdateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type RiskFormAppeal struct {
	// A text explanation for the appeal.
	Explanation string `json:"explanation,omitempty"`
}
type RiskFormAuthorizationDocuments struct {
	// IDs of uploaded files attached as authorization documents.
	Files []string `json:"files,omitempty"`
}

// Files evidencing payments accepted outside Stripe. Each entry includes a file ID and, when supplied, a document type tag.
type RiskFormPaymentProcessingMultiProcessorFile struct {
	// ID of the uploaded file.
	File string `json:"file,omitempty"`
	// The document type: `alt_financial_statements`, `recent_bank_statements`, or `payment_receipt`.
	Tag string `json:"tag,omitempty"`
}
type RiskFormPaymentProcessingMultiProcessor struct {
	// Files evidencing payments accepted outside Stripe. Each entry includes a file ID and, when supplied, a document type tag.
	Files []*RiskFormPaymentProcessingMultiProcessorFile `json:"files,omitempty"`
	// Whether the connected account accepts payments outside Stripe.
	Value string `json:"value,omitempty"`
}
type RiskFormPaymentProcessing struct {
	MultiProcessor *RiskFormPaymentProcessingMultiProcessor `json:"multi_processor,omitempty"`
	// How the connected account uses Stripe.
	Purpose string `json:"purpose,omitempty"`
	// Why the connected account chose Stripe.
	StripeSelectionRationale string `json:"stripe_selection_rationale,omitempty"`
}
type RiskFormProductRemoval struct {
	// Time when the prohibited items were removed, measured in seconds since the Unix epoch.
	ItemsRemovedAt int64 `json:"items_removed_at,omitempty"`
}
type RiskFormReviewedFieldsAttestation struct {
	// Time when the connected account attested to these fields, measured in seconds since the Unix epoch.
	AttestedAt int64 `json:"attested_at,omitempty"`
	// Account field paths covered by this attestation. These are currently returned in v2 Accounts API format, such as `defaults.profile.product_description`.
	Fields []string `json:"fields,omitempty"`
}
type RiskFormReviewedFields struct {
	Attestation *RiskFormReviewedFieldsAttestation `json:"attestation,omitempty"`
}
type RiskFormSelfFundingContextSourceOfWealth struct {
	// A text explanation of the connected account's source of wealth.
	Description string `json:"description,omitempty"`
	// IDs of uploaded files evidencing the source of wealth.
	Files []string `json:"files,omitempty"`
}
type RiskFormSelfFundingContext struct {
	SourceOfWealth *RiskFormSelfFundingContextSourceOfWealth `json:"source_of_wealth,omitempty"`
}
type RiskFormSubject struct {
	// Account field paths the connected account must review. These are currently returned in v2 Accounts API format, such as `defaults.profile.product_description`. Present for reviewed fields Forms.
	FieldsToReview []string `json:"fields_to_review,omitempty"`
	// IDs of the transactions this Form is about. Present for transaction context Forms.
	Transactions []string `json:"transactions,omitempty"`
}
type RiskFormTransactionContextSourceOfFunds struct {
	// A text explanation of where the funds for the transaction came from.
	Description string `json:"description,omitempty"`
	// IDs of uploaded files evidencing the source of funds.
	Files []string `json:"files,omitempty"`
}
type RiskFormTransactionContext struct {
	// How often the connected account expects transactions like this one. Can be `one_time_payment`, `recurring_subscription`, `repeat_customer`, `split_payment`, or `something_else`. A custom explanation can follow `something_else: `.
	Frequency string `json:"frequency,omitempty"`
	// The purpose of the transaction.
	Purpose string `json:"purpose,omitempty"`
	// IDs of uploaded files evidencing the purpose of the transaction.
	PurposeFiles []string `json:"purpose_files,omitempty"`
	// The connected account's relationship to the sender.
	SenderRelationship string                                   `json:"sender_relationship,omitempty"`
	SourceOfFunds      *RiskFormTransactionContextSourceOfFunds `json:"source_of_funds,omitempty"`
	// How the funds from the transaction will be used.
	UseOfFunds string `json:"use_of_funds,omitempty"`
}

// A risk Form represents a request from Stripe for information about a connected account.
type RiskForm struct {
	APIResource
	Appeal                 *RiskFormAppeal                 `json:"appeal,omitempty"`
	AuthorizationDocuments *RiskFormAuthorizationDocuments `json:"authorization_documents,omitempty"`
	// Time at which the Form closed, measured in seconds since the Unix epoch.
	ClosedAt int64 `json:"closed_at"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If it exists in a sandbox, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Time at which the Form opened, measured in seconds since the Unix epoch.
	OpenedAt           int64                       `json:"opened_at"`
	PaymentProcessing  *RiskFormPaymentProcessing  `json:"payment_processing,omitempty"`
	ProductRemoval     *RiskFormProductRemoval     `json:"product_removal,omitempty"`
	ReviewedFields     *RiskFormReviewedFields     `json:"reviewed_fields,omitempty"`
	SelfFundingContext *RiskFormSelfFundingContext `json:"self_funding_context,omitempty"`
	// The current status of the Form.
	Status             RiskFormStatus              `json:"status"`
	Subject            *RiskFormSubject            `json:"subject,omitempty"`
	TransactionContext *RiskFormTransactionContext `json:"transaction_context,omitempty"`
	// The type of information requested by Stripe.
	Type RiskFormType `json:"type"`
}
