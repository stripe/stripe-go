//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The current status of the inquiry.
type V2RiskInquiryStatus string

// List of values that V2RiskInquiryStatus can take
const (
	V2RiskInquiryStatusClosed V2RiskInquiryStatus = "closed"
	V2RiskInquiryStatusOpen   V2RiskInquiryStatus = "open"
)

// The type of inquiry.
type V2RiskInquiryType string

// List of values that V2RiskInquiryType can take
const (
	V2RiskInquiryTypeAppeal                 V2RiskInquiryType = "appeal"
	V2RiskInquiryTypeAuthorizationDocuments V2RiskInquiryType = "authorization_documents"
	V2RiskInquiryTypeProductRemoval         V2RiskInquiryType = "product_removal"
)

// Data for appeal inquiries. Only present when type is appeal.
type V2RiskInquiryAppeal struct {
	// A text explanation for the appeal.
	Explanation string `json:"explanation"`
}

// Data for authorization_documents inquiries. Only present when type is authorization_documents.
type V2RiskInquiryAuthorizationDocuments struct {
	// IDs of uploaded files to attach as authorization documents.
	Files []string `json:"files"`
}

// Data for product_removal inquiries. Only present when type is product_removal.
type V2RiskInquiryProductRemoval struct {
	// The timestamp when the prohibited items were removed.
	ItemsRemovedAt time.Time `json:"items_removed_at"`
}

// A risk inquiry represents a request from Stripe for information about a connected account.
type V2RiskInquiry struct {
	APIResource
	// Data for appeal inquiries. Only present when type is appeal.
	Appeal *V2RiskInquiryAppeal `json:"appeal,omitempty"`
	// Data for authorization_documents inquiries. Only present when type is authorization_documents.
	AuthorizationDocuments *V2RiskInquiryAuthorizationDocuments `json:"authorization_documents,omitempty"`
	// Time at which the inquiry was closed.
	ClosedAt time.Time `json:"closed_at"`
	// Time at which the inquiry was created.
	Created time.Time `json:"created"`
	// Unique identifier for the inquiry.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Time at which the inquiry was opened.
	OpenedAt time.Time `json:"opened_at"`
	// Data for product_removal inquiries. Only present when type is product_removal.
	ProductRemoval *V2RiskInquiryProductRemoval `json:"product_removal,omitempty"`
	// The current status of the inquiry.
	Status V2RiskInquiryStatus `json:"status"`
	// The type of inquiry.
	Type V2RiskInquiryType `json:"type"`
}
