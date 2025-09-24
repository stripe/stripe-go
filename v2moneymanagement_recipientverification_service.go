//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v2MoneyManagementRecipientVerificationService is used to invoke recipientverification related APIs.
type v2MoneyManagementRecipientVerificationService struct {
	B   Backend
	Key string
}

// Creates a RecipientVerification to verify the recipient you intend to send funds to.
func (c v2MoneyManagementRecipientVerificationService) Create(ctx context.Context, params *V2MoneyManagementRecipientVerificationCreateParams) (*V2MoneyManagementRecipientVerification, error) {
	if params == nil {
		params = &V2MoneyManagementRecipientVerificationCreateParams{}
	}
	params.Context = ctx
	recipientverification := &V2MoneyManagementRecipientVerification{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/recipient_verifications", c.Key, params, recipientverification)
	return recipientverification, err
}

// Retrieves the details of an existing RecipientVerification by passing the unique RecipientVerification ID.
func (c v2MoneyManagementRecipientVerificationService) Retrieve(ctx context.Context, id string, params *V2MoneyManagementRecipientVerificationRetrieveParams) (*V2MoneyManagementRecipientVerification, error) {
	if params == nil {
		params = &V2MoneyManagementRecipientVerificationRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v2/money_management/recipient_verifications/%s", id)
	recipientverification := &V2MoneyManagementRecipientVerification{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, recipientverification)
	return recipientverification, err
}

// Acknowledges an existing RecipientVerification. Only RecipientVerification awaiting acknowledgement can be acknowledged.
func (c v2MoneyManagementRecipientVerificationService) Acknowledge(ctx context.Context, id string, params *V2MoneyManagementRecipientVerificationAcknowledgeParams) (*V2MoneyManagementRecipientVerification, error) {
	if params == nil {
		params = &V2MoneyManagementRecipientVerificationAcknowledgeParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v2/money_management/recipient_verifications/%s/acknowledge", id)
	recipientverification := &V2MoneyManagementRecipientVerification{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, recipientverification)
	return recipientverification, err
}
