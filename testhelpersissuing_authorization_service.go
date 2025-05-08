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

// v1TestHelpersIssuingAuthorizationService is used to invoke /v1/issuing/authorizations APIs.
type v1TestHelpersIssuingAuthorizationService struct {
	B   Backend
	Key string
}

// Create a test-mode authorization.
func (c v1TestHelpersIssuingAuthorizationService) Create(ctx context.Context, params *TestHelpersIssuingAuthorizationCreateParams) (*IssuingAuthorization, error) {
	if params == nil {
		params = &TestHelpersIssuingAuthorizationCreateParams{}
	}
	params.Context = ctx
	authorization := &IssuingAuthorization{}
	err := c.B.Call(
		http.MethodPost, "/v1/test_helpers/issuing/authorizations", c.Key, params, authorization)
	return authorization, err
}

// Capture a test-mode authorization.
func (c v1TestHelpersIssuingAuthorizationService) Capture(ctx context.Context, id string, params *TestHelpersIssuingAuthorizationCaptureParams) (*IssuingAuthorization, error) {
	if params == nil {
		params = &TestHelpersIssuingAuthorizationCaptureParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/issuing/authorizations/%s/capture", id)
	authorization := &IssuingAuthorization{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}

// Expire a test-mode Authorization.
func (c v1TestHelpersIssuingAuthorizationService) Expire(ctx context.Context, id string, params *TestHelpersIssuingAuthorizationExpireParams) (*IssuingAuthorization, error) {
	if params == nil {
		params = &TestHelpersIssuingAuthorizationExpireParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/test_helpers/issuing/authorizations/%s/expire", id)
	authorization := &IssuingAuthorization{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}

// Finalize the amount on an Authorization prior to capture, when the initial authorization was for an estimated amount.
func (c v1TestHelpersIssuingAuthorizationService) FinalizeAmount(ctx context.Context, id string, params *TestHelpersIssuingAuthorizationFinalizeAmountParams) (*IssuingAuthorization, error) {
	if params == nil {
		params = &TestHelpersIssuingAuthorizationFinalizeAmountParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/issuing/authorizations/%s/finalize_amount", id)
	authorization := &IssuingAuthorization{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}

// Increment a test-mode Authorization.
func (c v1TestHelpersIssuingAuthorizationService) Increment(ctx context.Context, id string, params *TestHelpersIssuingAuthorizationIncrementParams) (*IssuingAuthorization, error) {
	if params == nil {
		params = &TestHelpersIssuingAuthorizationIncrementParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/issuing/authorizations/%s/increment", id)
	authorization := &IssuingAuthorization{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}

// Respond to a fraud challenge on a testmode Issuing authorization, simulating either a confirmation of fraud or a correction of legitimacy.
func (c v1TestHelpersIssuingAuthorizationService) Respond(ctx context.Context, id string, params *TestHelpersIssuingAuthorizationRespondParams) (*IssuingAuthorization, error) {
	if params == nil {
		params = &TestHelpersIssuingAuthorizationRespondParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/issuing/authorizations/%s/fraud_challenges/respond", id)
	authorization := &IssuingAuthorization{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}

// Reverse a test-mode Authorization.
func (c v1TestHelpersIssuingAuthorizationService) Reverse(ctx context.Context, id string, params *TestHelpersIssuingAuthorizationReverseParams) (*IssuingAuthorization, error) {
	if params == nil {
		params = &TestHelpersIssuingAuthorizationReverseParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/issuing/authorizations/%s/reverse", id)
	authorization := &IssuingAuthorization{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}
