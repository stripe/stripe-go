//
//
// File generated from our OpenAPI spec
//
//

// Package onboardinglink provides the /v1/terminal/onboarding_links APIs
package onboardinglink

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
)

// Client is used to invoke /v1/terminal/onboarding_links APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a new OnboardingLink object that contains a redirect_url used for onboarding onto Tap to Pay on iPhone.
func New(params *stripe.TerminalOnboardingLinkParams) (*stripe.TerminalOnboardingLink, error) {
	return getC().New(params)
}

// Creates a new OnboardingLink object that contains a redirect_url used for onboarding onto Tap to Pay on iPhone.
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.TerminalOnboardingLinkParams) (*stripe.TerminalOnboardingLink, error) {
	onboardinglink := &stripe.TerminalOnboardingLink{}
	err := c.B.Call(
		http.MethodPost, "/v1/terminal/onboarding_links", c.Key, params, onboardinglink)
	return onboardinglink, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
