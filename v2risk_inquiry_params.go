//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Lists risk inquiries for a connected account.
type V2RiskInquiryListParams struct {
	Params `form:"*"`
	// The account to list inquiries for.
	Account *string `form:"account" json:"account"`
	// Maximum number of results to return. Default: 10. Valid range: 1-100.
	Limit *int64 `form:"limit" json:"limit,omitempty"`
}

// Retrieves a risk inquiry by ID.
type V2RiskInquiryParams struct {
	Params `form:"*"`
	// Provide this for appeal inquiries.
	Appeal *V2RiskInquiryAppealParams `form:"appeal" json:"appeal,omitempty"`
	// Provide this for authorization_documents inquiries.
	AuthorizationDocuments *V2RiskInquiryAuthorizationDocumentsParams `form:"authorization_documents" json:"authorization_documents,omitempty"`
	// Provide this for product_removal inquiries.
	ProductRemoval *V2RiskInquiryProductRemovalParams `form:"product_removal" json:"product_removal,omitempty"`
}

// Provide this for appeal inquiries.
type V2RiskInquiryAppealParams struct {
	// A text explanation for the appeal.
	Explanation *string `form:"explanation" json:"explanation"`
}

// Provide this for authorization_documents inquiries.
type V2RiskInquiryAuthorizationDocumentsParams struct {
	// IDs of uploaded files to attach as authorization documents.
	Files []*string `form:"files" json:"files"`
}

// Provide this for product_removal inquiries.
type V2RiskInquiryProductRemovalParams struct {
	// The timestamp when the prohibited items were removed.
	ItemsRemovedAt *time.Time `form:"items_removed_at" json:"items_removed_at"`
}

// Retrieves a risk inquiry by ID.
type V2RiskInquiryRetrieveParams struct {
	Params `form:"*"`
}

// Provide this for appeal inquiries.
type V2RiskInquiryUpdateAppealParams struct {
	// A text explanation for the appeal.
	Explanation *string `form:"explanation" json:"explanation"`
}

// Provide this for authorization_documents inquiries.
type V2RiskInquiryUpdateAuthorizationDocumentsParams struct {
	// IDs of uploaded files to attach as authorization documents.
	Files []*string `form:"files" json:"files"`
}

// Provide this for product_removal inquiries.
type V2RiskInquiryUpdateProductRemovalParams struct {
	// The timestamp when the prohibited items were removed.
	ItemsRemovedAt *time.Time `form:"items_removed_at" json:"items_removed_at"`
}

// Submits a response to a risk inquiry.
type V2RiskInquiryUpdateParams struct {
	Params `form:"*"`
	// Provide this for appeal inquiries.
	Appeal *V2RiskInquiryUpdateAppealParams `form:"appeal" json:"appeal,omitempty"`
	// Provide this for authorization_documents inquiries.
	AuthorizationDocuments *V2RiskInquiryUpdateAuthorizationDocumentsParams `form:"authorization_documents" json:"authorization_documents,omitempty"`
	// Provide this for product_removal inquiries.
	ProductRemoval *V2RiskInquiryUpdateProductRemovalParams `form:"product_removal" json:"product_removal,omitempty"`
}
