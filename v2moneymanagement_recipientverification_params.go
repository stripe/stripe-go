//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates a RecipientVerification to verify the recipient you intend to send funds to.
type V2MoneyManagementRecipientVerificationParams struct {
	Params `form:"*"`
	// ID of the payout method.
	PayoutMethod *string `form:"payout_method" json:"payout_method,omitempty"`
	// ID of the recipient account. Required if the recipient is different from the sender. Omit for payouts to your own account.
	Recipient *string `form:"recipient" json:"recipient,omitempty"`
}

// Acknowledges an existing RecipientVerification in the `awaiting_acknowledgement` state.
type V2MoneyManagementRecipientVerificationAcknowledgeParams struct {
	Params `form:"*"`
}

// Creates a RecipientVerification to verify the recipient you intend to send funds to.
type V2MoneyManagementRecipientVerificationCreateParams struct {
	Params `form:"*"`
	// ID of the payout method.
	PayoutMethod *string `form:"payout_method" json:"payout_method"`
	// ID of the recipient account. Required if the recipient is different from the sender. Omit for payouts to your own account.
	Recipient *string `form:"recipient" json:"recipient,omitempty"`
}

// Retrieves the details of an existing RecipientVerification by passing the unique RecipientVerification ID.
type V2MoneyManagementRecipientVerificationRetrieveParams struct {
	Params `form:"*"`
}
