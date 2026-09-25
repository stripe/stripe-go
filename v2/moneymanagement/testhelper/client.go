//
//
// File generated from our OpenAPI spec
//
//

// Package testhelper provides the testhelper related APIs
package testhelper

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v86"
)

// Client is used to invoke testhelper related APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates an EarnedCredit in a Sandbox environment for testing purposes.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) EarnedCredits(params *stripe.V2MoneyManagementTestHelperEarnedCreditsParams) (*stripe.V2MoneyManagementEarnedCreditSimulation, error) {
	earnedcreditsimulation := &stripe.V2MoneyManagementEarnedCreditSimulation{}
	err := c.B.Call(
		http.MethodPost, "/v2/money_management/test_helpers/earned_credits", c.Key, params, earnedcreditsimulation)
	return earnedcreditsimulation, err
}
