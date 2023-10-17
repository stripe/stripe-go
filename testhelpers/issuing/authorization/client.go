//
//
// File generated from our OpenAPI spec
//
//

// Package authorization provides the /issuing/authorizations APIs
package authorization

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v76"
)

// Client is used to invoke /issuing/authorizations APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new issuing authorization.
func New(params *stripe.TestHelpersIssuingAuthorizationParams) (*stripe.IssuingAuthorization, error) {
	return getC().New(params)
}

// New creates a new issuing authorization.
func (c Client) New(params *stripe.TestHelpersIssuingAuthorizationParams) (*stripe.IssuingAuthorization, error) {
	authorization := &stripe.IssuingAuthorization{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/test_helpers/issuing/authorizations",
		c.Key,
		params,
		authorization,
	)
	return authorization, err
}

// Capture is the method for the `POST /v1/test_helpers/issuing/authorizations/{authorization}/capture` API.
func Capture(id string, params *stripe.TestHelpersIssuingAuthorizationCaptureParams) (*stripe.IssuingAuthorization, error) {
	return getC().Capture(id, params)
}

// Capture is the method for the `POST /v1/test_helpers/issuing/authorizations/{authorization}/capture` API.
func (c Client) Capture(id string, params *stripe.TestHelpersIssuingAuthorizationCaptureParams) (*stripe.IssuingAuthorization, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/authorizations/%s/capture",
		id,
	)
	authorization := &stripe.IssuingAuthorization{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}

// Expire is the method for the `POST /v1/test_helpers/issuing/authorizations/{authorization}/expire` API.
func Expire(id string, params *stripe.TestHelpersIssuingAuthorizationExpireParams) (*stripe.IssuingAuthorization, error) {
	return getC().Expire(id, params)
}

// Expire is the method for the `POST /v1/test_helpers/issuing/authorizations/{authorization}/expire` API.
func (c Client) Expire(id string, params *stripe.TestHelpersIssuingAuthorizationExpireParams) (*stripe.IssuingAuthorization, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/authorizations/%s/expire",
		id,
	)
	authorization := &stripe.IssuingAuthorization{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}

// Increment is the method for the `POST /v1/test_helpers/issuing/authorizations/{authorization}/increment` API.
func Increment(id string, params *stripe.TestHelpersIssuingAuthorizationIncrementParams) (*stripe.IssuingAuthorization, error) {
	return getC().Increment(id, params)
}

// Increment is the method for the `POST /v1/test_helpers/issuing/authorizations/{authorization}/increment` API.
func (c Client) Increment(id string, params *stripe.TestHelpersIssuingAuthorizationIncrementParams) (*stripe.IssuingAuthorization, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/authorizations/%s/increment",
		id,
	)
	authorization := &stripe.IssuingAuthorization{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}

// Reverse is the method for the `POST /v1/test_helpers/issuing/authorizations/{authorization}/reverse` API.
func Reverse(id string, params *stripe.TestHelpersIssuingAuthorizationReverseParams) (*stripe.IssuingAuthorization, error) {
	return getC().Reverse(id, params)
}

// Reverse is the method for the `POST /v1/test_helpers/issuing/authorizations/{authorization}/reverse` API.
func (c Client) Reverse(id string, params *stripe.TestHelpersIssuingAuthorizationReverseParams) (*stripe.IssuingAuthorization, error) {
	path := stripe.FormatURLPath(
		"/v1/test_helpers/issuing/authorizations/%s/reverse",
		id,
	)
	authorization := &stripe.IssuingAuthorization{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, authorization)
	return authorization, err
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
