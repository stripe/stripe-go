//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// The current status of the Form.
type V2RiskFormStatus string

// List of values that V2RiskFormStatus can take
const (
	V2RiskFormStatusClosed V2RiskFormStatus = "closed"
	V2RiskFormStatusOpen   V2RiskFormStatus = "open"
)

// The type of Form.
type V2RiskFormType string

// List of values that V2RiskFormType can take
const (
	V2RiskFormTypeAppeal                 V2RiskFormType = "appeal"
	V2RiskFormTypeAuthorizationDocuments V2RiskFormType = "authorization_documents"
	V2RiskFormTypeProductRemoval         V2RiskFormType = "product_removal"
	V2RiskFormTypeSelfFundingContext     V2RiskFormType = "self_funding_context"
	V2RiskFormTypeTransactionContext     V2RiskFormType = "transaction_context"
)

// Data for appeal inquiries. Only present when type is appeal.
type V2RiskFormAppeal struct {
	// A text explanation for the appeal.
	Explanation string `json:"explanation"`
}

// Data for authorization_documents inquiries. Only present when type is authorization_documents.
type V2RiskFormAuthorizationDocuments struct {
	// IDs of uploaded files to attach as authorization documents.
	Files []string `json:"files"`
}

// Data for product_removal inquiries. Only present when type is product_removal.
type V2RiskFormProductRemoval struct {
	// The timestamp when the prohibited items were removed.
	ItemsRemovedAt time.Time `json:"items_removed_at"`
}

// The connected account's source of wealth.
type V2RiskFormSelfFundingContextSourceOfWealth struct {
	// A text explanation of the connected account's source of wealth.
	Description string `json:"description"`
	// IDs of uploaded files evidencing the source of wealth.
	Files []string `json:"files"`
}

// Data for self_funding_context inquiries. Only present when type is self_funding_context.
type V2RiskFormSelfFundingContext struct {
	// The connected account's source of wealth.
	SourceOfWealth *V2RiskFormSelfFundingContextSourceOfWealth `json:"source_of_wealth,omitempty"`
}

// What this Form was raised about, as a typed object. Only present when the Form
// is scoped to a specific object. Absent when the Form covers the account as a whole.
type V2RiskFormSubject struct {
	// ID of the transaction this Form is about. Present only for transaction_context
	// Forms. Absent for other Form types.
	Transaction string `json:"transaction,omitempty"`
	// IDs of the transactions this inquiry is about.
	Transactions []string `json:"transactions,omitempty"`
}

// Where the funds for the transaction came from.
type V2RiskFormTransactionContextSourceOfFunds struct {
	// A text explanation of where the funds for the transaction came from.
	Description string `json:"description"`
	// IDs of uploaded files evidencing the source of funds.
	Files []string `json:"files"`
}

// Data for transaction_context inquiries. Only present when type is transaction_context.
type V2RiskFormTransactionContext struct {
	// The purpose of the transaction.
	Purpose string `json:"purpose"`
	// The connected account's relationship to the sender. Optional.
	SenderRelationship string `json:"sender_relationship,omitempty"`
	// Where the funds for the transaction came from.
	SourceOfFunds *V2RiskFormTransactionContextSourceOfFunds `json:"source_of_funds,omitempty"`
	// How the funds from the transaction will be used.
	UseOfFunds string `json:"use_of_funds"`
}

// A risk Form represents a request from Stripe for information about a connected account.
type V2RiskForm struct {
	APIResource
	// Data for appeal inquiries. Only present when type is appeal.
	Appeal *V2RiskFormAppeal `json:"appeal,omitempty"`
	// Data for authorization_documents inquiries. Only present when type is authorization_documents.
	AuthorizationDocuments *V2RiskFormAuthorizationDocuments `json:"authorization_documents,omitempty"`
	// Time at which the Form was closed.
	ClosedAt time.Time `json:"closed_at"`
	// Time at which the Form was created.
	Created time.Time `json:"created"`
	// Unique identifier for the Form.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Time at which the Form was opened.
	OpenedAt time.Time `json:"opened_at"`
	// Data for product_removal inquiries. Only present when type is product_removal.
	ProductRemoval *V2RiskFormProductRemoval `json:"product_removal,omitempty"`
	// Data for self_funding_context inquiries. Only present when type is self_funding_context.
	SelfFundingContext *V2RiskFormSelfFundingContext `json:"self_funding_context,omitempty"`
	// The current status of the Form.
	Status V2RiskFormStatus `json:"status"`
	// What this Form was raised about, as a typed object. Only present when the Form
	// is scoped to a specific object. Absent when the Form covers the account as a whole.
	Subject *V2RiskFormSubject `json:"subject,omitempty"`
	// Data for transaction_context inquiries. Only present when type is transaction_context.
	TransactionContext *V2RiskFormTransactionContext `json:"transaction_context,omitempty"`
	// The type of Form.
	Type V2RiskFormType `json:"type"`
}
