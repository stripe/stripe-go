//
//
// File generated from our OpenAPI spec
//
//

// Package recipientverification provides the recipientverification related APIs
package recipientverification

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke recipientverification related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a RecipientVerification to verify the recipient you intend to send funds to.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.V2MoneyManagementRecipientVerificationParams) (*stripe.V2MoneyManagementRecipientVerification, error) {
	recipientverification := &stripe.V2MoneyManagementRecipientVerification{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/recipient_verifications", c.Key, params, recipientverification)
	return recipientverification, err
}

// Retrieves the details of an existing RecipientVerification by passing the unique RecipientVerification ID.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.V2MoneyManagementRecipientVerificationParams) (*stripe.V2MoneyManagementRecipientVerification, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/recipient_verifications/%s", id)
	recipientverification := &stripe.V2MoneyManagementRecipientVerification{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, recipientverification)
	return recipientverification, err
}

// Acknowledges an existing RecipientVerification. Only RecipientVerification awaiting acknowledgement can be acknowledged.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Acknowledge(id string, params *stripe.V2MoneyManagementRecipientVerificationAcknowledgeParams) (*stripe.V2MoneyManagementRecipientVerification, error) {
	path := stripe.FormatURLPath(
		"/v2/money_management/recipient_verifications/%s/acknowledge", id)
	recipientverification := &stripe.V2MoneyManagementRecipientVerification{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, recipientverification)
	return recipientverification, err
}
