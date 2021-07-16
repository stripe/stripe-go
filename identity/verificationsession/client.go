//
//
// File generated from our OpenAPI spec
//
//

// Package verificationsession provides the /identity/verification_sessions APIs
package verificationsession

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/form"
)

// Client is used to invoke /identity/verification_sessions APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new identity verification session.
func New(params *stripe.IdentityVerificationSessionParams) (*stripe.IdentityVerificationSession, error) {
	return getC().New(params)
}

// New creates a new identity verification session.
func (c Client) New(params *stripe.IdentityVerificationSessionParams) (*stripe.IdentityVerificationSession, error) {
	verificationsession := &stripe.IdentityVerificationSession{}
	err := c.B.Call(
		http.MethodPost,
		"/v1/identity/verification_sessions",
		c.Key,
		params,
		verificationsession,
	)
	return verificationsession, err
}

// Get returns the details of an identity verification session.
func Get(id string, params *stripe.IdentityVerificationSessionParams) (*stripe.IdentityVerificationSession, error) {
	return getC().Get(id, params)
}

// Get returns the details of an identity verification session.
func (c Client) Get(id string, params *stripe.IdentityVerificationSessionParams) (*stripe.IdentityVerificationSession, error) {
	path := stripe.FormatURLPath("/v1/identity/verification_sessions/%s", id)
	verificationsession := &stripe.IdentityVerificationSession{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, verificationsession)
	return verificationsession, err
}

// Update updates an identity verification session's properties.
func Update(id string, params *stripe.IdentityVerificationSessionParams) (*stripe.IdentityVerificationSession, error) {
	return getC().Update(id, params)
}

// Update updates an identity verification session's properties.
func (c Client) Update(id string, params *stripe.IdentityVerificationSessionParams) (*stripe.IdentityVerificationSession, error) {
	path := stripe.FormatURLPath("/v1/identity/verification_sessions/%s", id)
	verificationsession := &stripe.IdentityVerificationSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, verificationsession)
	return verificationsession, err
}

// Cancel is the method for the `POST /v1/identity/verification_sessions/{session}/cancel` API.
func Cancel(id string, params *stripe.IdentityVerificationSessionCancelParams) (*stripe.IdentityVerificationSession, error) {
	return getC().Cancel(id, params)
}

// Cancel is the method for the `POST /v1/identity/verification_sessions/{session}/cancel` API.
func (c Client) Cancel(id string, params *stripe.IdentityVerificationSessionCancelParams) (*stripe.IdentityVerificationSession, error) {
	path := stripe.FormatURLPath(
		"/v1/identity/verification_sessions/%s/cancel",
		id,
	)
	verificationsession := &stripe.IdentityVerificationSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, verificationsession)
	return verificationsession, err
}

// Redact is the method for the `POST /v1/identity/verification_sessions/{session}/redact` API.
func Redact(id string, params *stripe.IdentityVerificationSessionRedactParams) (*stripe.IdentityVerificationSession, error) {
	return getC().Redact(id, params)
}

// Redact is the method for the `POST /v1/identity/verification_sessions/{session}/redact` API.
func (c Client) Redact(id string, params *stripe.IdentityVerificationSessionRedactParams) (*stripe.IdentityVerificationSession, error) {
	path := stripe.FormatURLPath(
		"/v1/identity/verification_sessions/%s/redact",
		id,
	)
	verificationsession := &stripe.IdentityVerificationSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, verificationsession)
	return verificationsession, err
}

// List returns a list of identity verification sessions.
func List(params *stripe.IdentityVerificationSessionListParams) *Iter {
	return getC().List(params)
}

// List returns a list of identity verification sessions.
func (c Client) List(listParams *stripe.IdentityVerificationSessionListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IdentityVerificationSessionList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/identity/verification_sessions", c.Key, b, p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for identity verification sessions.
type Iter struct {
	*stripe.Iter
}

// IdentityVerificationSession returns the identity verification session which the iterator is currently pointing to.
func (i *Iter) IdentityVerificationSession() *stripe.IdentityVerificationSession {
	return i.Current().(*stripe.IdentityVerificationSession)
}

// IdentityVerificationSessionList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IdentityVerificationSessionList() *stripe.IdentityVerificationSessionList {
	return i.List().(*stripe.IdentityVerificationSessionList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
