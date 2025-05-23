//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1IdentityVerificationSessionService is used to invoke /v1/identity/verification_sessions APIs.
type v1IdentityVerificationSessionService struct {
	B   Backend
	Key string
}

// Creates a VerificationSession object.
//
// After the VerificationSession is created, display a verification modal using the session client_secret or send your users to the session's url.
//
// If your API key is in test mode, verification checks won't actually process, though everything else will occur as if in live mode.
//
// Related guide: [Verify your users' identity documents](https://docs.stripe.com/docs/identity/verify-identity-documents)
func (c v1IdentityVerificationSessionService) Create(ctx context.Context, params *IdentityVerificationSessionCreateParams) (*IdentityVerificationSession, error) {
	if params == nil {
		params = &IdentityVerificationSessionCreateParams{}
	}
	params.Context = ctx
	verificationsession := &IdentityVerificationSession{}
	err := c.B.Call(
		http.MethodPost, "/v1/identity/verification_sessions", c.Key, params, verificationsession)
	return verificationsession, err
}

// Retrieves the details of a VerificationSession that was previously created.
//
// When the session status is requires_input, you can use this method to retrieve a valid
// client_secret or url to allow re-submission.
func (c v1IdentityVerificationSessionService) Retrieve(ctx context.Context, id string, params *IdentityVerificationSessionRetrieveParams) (*IdentityVerificationSession, error) {
	if params == nil {
		params = &IdentityVerificationSessionRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/identity/verification_sessions/%s", id)
	verificationsession := &IdentityVerificationSession{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, verificationsession)
	return verificationsession, err
}

// Updates a VerificationSession object.
//
// When the session status is requires_input, you can use this method to update the
// verification check and options.
func (c v1IdentityVerificationSessionService) Update(ctx context.Context, id string, params *IdentityVerificationSessionUpdateParams) (*IdentityVerificationSession, error) {
	if params == nil {
		params = &IdentityVerificationSessionUpdateParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/identity/verification_sessions/%s", id)
	verificationsession := &IdentityVerificationSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, verificationsession)
	return verificationsession, err
}

// A VerificationSession object can be canceled when it is in requires_input [status](https://docs.stripe.com/docs/identity/how-sessions-work).
//
// Once canceled, future submission attempts are disabled. This cannot be undone. [Learn more](https://docs.stripe.com/docs/identity/verification-sessions#cancel).
func (c v1IdentityVerificationSessionService) Cancel(ctx context.Context, id string, params *IdentityVerificationSessionCancelParams) (*IdentityVerificationSession, error) {
	if params == nil {
		params = &IdentityVerificationSessionCancelParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/identity/verification_sessions/%s/cancel", id)
	verificationsession := &IdentityVerificationSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, verificationsession)
	return verificationsession, err
}

// Redact a VerificationSession to remove all collected information from Stripe. This will redact
// the VerificationSession and all objects related to it, including VerificationReports, Events,
// request logs, etc.
//
// A VerificationSession object can be redacted when it is in requires_input or verified
// [status](https://docs.stripe.com/docs/identity/how-sessions-work). Redacting a VerificationSession in requires_action
// state will automatically cancel it.
//
// The redaction process may take up to four days. When the redaction process is in progress, the
// VerificationSession's redaction.status field will be set to processing; when the process is
// finished, it will change to redacted and an identity.verification_session.redacted event
// will be emitted.
//
// Redaction is irreversible. Redacted objects are still accessible in the Stripe API, but all the
// fields that contain personal data will be replaced by the string [redacted] or a similar
// placeholder. The metadata field will also be erased. Redacted objects cannot be updated or
// used for any purpose.
//
// [Learn more](https://docs.stripe.com/docs/identity/verification-sessions#redact).
func (c v1IdentityVerificationSessionService) Redact(ctx context.Context, id string, params *IdentityVerificationSessionRedactParams) (*IdentityVerificationSession, error) {
	if params == nil {
		params = &IdentityVerificationSessionRedactParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/identity/verification_sessions/%s/redact", id)
	verificationsession := &IdentityVerificationSession{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, verificationsession)
	return verificationsession, err
}

// Returns a list of VerificationSessions
func (c v1IdentityVerificationSessionService) List(ctx context.Context, listParams *IdentityVerificationSessionListParams) Seq2[*IdentityVerificationSession, error] {
	if listParams == nil {
		listParams = &IdentityVerificationSessionListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*IdentityVerificationSession, ListContainer, error) {
		list := &IdentityVerificationSessionList{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/identity/verification_sessions", c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
