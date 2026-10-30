//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Retrieves a risk Form by ID.
type V2RiskFormParams struct {
	Params `form:"*"`
	// Provide this for appeal inquiries.
	Appeal *V2RiskFormAppealParams `form:"appeal" json:"appeal,omitempty"`
	// Provide this for authorization_documents inquiries.
	AuthorizationDocuments *V2RiskFormAuthorizationDocumentsParams `form:"authorization_documents" json:"authorization_documents,omitempty"`
	// Provide this for product_removal inquiries.
	ProductRemoval *V2RiskFormProductRemovalParams `form:"product_removal" json:"product_removal,omitempty"`
	// Provide this for self_funding_context inquiries.
	SelfFundingContext *V2RiskFormSelfFundingContextParams `form:"self_funding_context" json:"self_funding_context,omitempty"`
	// Provide this for transaction_context inquiries.
	TransactionContext *V2RiskFormTransactionContextParams `form:"transaction_context" json:"transaction_context,omitempty"`
}

// Provide this for appeal inquiries.
type V2RiskFormAppealParams struct {
	// A text explanation for the appeal.
	Explanation *string `form:"explanation" json:"explanation"`
}

// Provide this for authorization_documents inquiries.
type V2RiskFormAuthorizationDocumentsParams struct {
	// IDs of uploaded files to attach as authorization documents.
	Files []*string `form:"files" json:"files"`
}

// Provide this for product_removal inquiries.
type V2RiskFormProductRemovalParams struct {
	// The timestamp when the prohibited items were removed.
	ItemsRemovedAt *time.Time `form:"items_removed_at" json:"items_removed_at"`
}

// The connected account's source of wealth.
type V2RiskFormSelfFundingContextSourceOfWealthParams struct {
	// A text explanation of the connected account's source of wealth.
	Description *string `form:"description" json:"description"`
	// IDs of uploaded files evidencing the source of wealth.
	Files []*string `form:"files" json:"files"`
}

// Provide this for self_funding_context inquiries.
type V2RiskFormSelfFundingContextParams struct {
	// The connected account's source of wealth.
	SourceOfWealth *V2RiskFormSelfFundingContextSourceOfWealthParams `form:"source_of_wealth" json:"source_of_wealth,omitempty"`
}

// Where the funds for the transaction came from.
type V2RiskFormTransactionContextSourceOfFundsParams struct {
	// A text explanation of where the funds for the transaction came from.
	Description *string `form:"description" json:"description"`
	// IDs of uploaded files evidencing the source of funds.
	Files []*string `form:"files" json:"files"`
}

// Provide this for transaction_context inquiries.
type V2RiskFormTransactionContextParams struct {
	// The purpose of the transaction.
	Purpose *string `form:"purpose" json:"purpose"`
	// The connected account's relationship to the sender. Optional.
	SenderRelationship *string `form:"sender_relationship" json:"sender_relationship,omitempty"`
	// Where the funds for the transaction came from.
	SourceOfFunds *V2RiskFormTransactionContextSourceOfFundsParams `form:"source_of_funds" json:"source_of_funds,omitempty"`
	// How the funds from the transaction will be used.
	UseOfFunds *string `form:"use_of_funds" json:"use_of_funds"`
}

// Retrieves a risk Form by ID.
type V2RiskFormRetrieveParams struct {
	Params `form:"*"`
}

// Provide this for appeal inquiries.
type V2RiskFormUpdateAppealParams struct {
	// A text explanation for the appeal.
	Explanation *string `form:"explanation" json:"explanation"`
}

// Provide this for authorization_documents inquiries.
type V2RiskFormUpdateAuthorizationDocumentsParams struct {
	// IDs of uploaded files to attach as authorization documents.
	Files []*string `form:"files" json:"files"`
}

// Provide this for product_removal inquiries.
type V2RiskFormUpdateProductRemovalParams struct {
	// The timestamp when the prohibited items were removed.
	ItemsRemovedAt *time.Time `form:"items_removed_at" json:"items_removed_at"`
}

// The connected account's source of wealth.
type V2RiskFormUpdateSelfFundingContextSourceOfWealthParams struct {
	// A text explanation of the connected account's source of wealth.
	Description *string `form:"description" json:"description"`
	// IDs of uploaded files evidencing the source of wealth.
	Files []*string `form:"files" json:"files"`
}

// Provide this for self_funding_context inquiries.
type V2RiskFormUpdateSelfFundingContextParams struct {
	// The connected account's source of wealth.
	SourceOfWealth *V2RiskFormUpdateSelfFundingContextSourceOfWealthParams `form:"source_of_wealth" json:"source_of_wealth,omitempty"`
}

// Where the funds for the transaction came from.
type V2RiskFormUpdateTransactionContextSourceOfFundsParams struct {
	// A text explanation of where the funds for the transaction came from.
	Description *string `form:"description" json:"description"`
	// IDs of uploaded files evidencing the source of funds.
	Files []*string `form:"files" json:"files"`
}

// Provide this for transaction_context inquiries.
type V2RiskFormUpdateTransactionContextParams struct {
	// The purpose of the transaction.
	Purpose *string `form:"purpose" json:"purpose"`
	// The connected account's relationship to the sender. Optional.
	SenderRelationship *string `form:"sender_relationship" json:"sender_relationship,omitempty"`
	// Where the funds for the transaction came from.
	SourceOfFunds *V2RiskFormUpdateTransactionContextSourceOfFundsParams `form:"source_of_funds" json:"source_of_funds,omitempty"`
	// How the funds from the transaction will be used.
	UseOfFunds *string `form:"use_of_funds" json:"use_of_funds"`
}

// Submits a response to a risk Form.
type V2RiskFormUpdateParams struct {
	Params `form:"*"`
	// Provide this for appeal inquiries.
	Appeal *V2RiskFormUpdateAppealParams `form:"appeal" json:"appeal,omitempty"`
	// Provide this for authorization_documents inquiries.
	AuthorizationDocuments *V2RiskFormUpdateAuthorizationDocumentsParams `form:"authorization_documents" json:"authorization_documents,omitempty"`
	// Provide this for product_removal inquiries.
	ProductRemoval *V2RiskFormUpdateProductRemovalParams `form:"product_removal" json:"product_removal,omitempty"`
	// Provide this for self_funding_context inquiries.
	SelfFundingContext *V2RiskFormUpdateSelfFundingContextParams `form:"self_funding_context" json:"self_funding_context,omitempty"`
	// Provide this for transaction_context inquiries.
	TransactionContext *V2RiskFormUpdateTransactionContextParams `form:"transaction_context" json:"transaction_context,omitempty"`
}
