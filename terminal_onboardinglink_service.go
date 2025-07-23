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

// v1TerminalOnboardingLinkService is used to invoke /v1/terminal/onboarding_links APIs.
type v1TerminalOnboardingLinkService struct {
	B   Backend
	Key string
}

// Creates a new OnboardingLink object that contains a redirect_url used for onboarding onto Tap to Pay on iPhone.
func (c v1TerminalOnboardingLinkService) Create(ctx context.Context, params *TerminalOnboardingLinkCreateParams) (*TerminalOnboardingLink, error) {
	if params == nil {
		params = &TerminalOnboardingLinkCreateParams{}
	}
	params.Context = ctx
	onboardinglink := &TerminalOnboardingLink{}
	err := c.B.Call(
		http.MethodPost, "/v1/terminal/onboarding_links", c.Key, params, onboardinglink)
	return onboardinglink, err
}
