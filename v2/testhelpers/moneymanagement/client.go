//
//
// File generated from our OpenAPI spec
//
//

// Package moneymanagement provides the moneymanagement related APIs
package moneymanagement

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke moneymanagement related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a RecipientVerification with a specified match result for testing purposes in a Sandbox environment.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) RecipientVerifications(params *stripe.V2TestHelpersMoneyManagementRecipientVerificationsParams) (*stripe.V2MoneyManagementRecipientVerification, error) {
	recipientverification := &stripe.V2MoneyManagementRecipientVerification{}
	err := c.B.Call(
		http.MethodPost, "/v2/test_helpers/money_management/recipient_verifications", c.Key, params, recipientverification)
	return recipientverification, err
}
