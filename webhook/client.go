// Package webhook provides utilities for parsing and validating Stripe webhook
// event payloads, including signature verification against the
// Stripe-Signature header.
package webhook

import (
	"encoding/hex"
	"fmt"
	"time"

	"github.com/stripe/stripe-go/v87"
)

//
// Public constants
//

const (
	// DefaultTolerance indicates that signatures older than this will be rejected by ConstructEvent.
	DefaultTolerance time.Duration = 300 * time.Second
)

//
// Public variables
//

// This block represents the list of errors that could be raised when using the webhook package.
// These alias the root stripe package's errors so that comparisons against
// either package's sentinel succeed.
var (
	ErrInvalidHeader    = stripe.ErrWebhookInvalidHeader
	ErrNoValidSignature = stripe.ErrWebhookNoValidSignature
	ErrNotSigned        = stripe.ErrWebhookNotSigned
	ErrTooOld           = stripe.ErrWebhookTooOld
	ErrEmptySecret      = stripe.ErrWebhookEmptySecret
)

//
// Public functions
//

// ComputeSignature computes a webhook signature using Stripe's v1 signing
// method.
//
// See https://stripe.com/docs/webhooks#signatures for more information.
func ComputeSignature(t time.Time, payload []byte, secret string) []byte {
	return stripe.ComputeSignature(t, payload, secret)
}

// ConstructEvent initializes an Event object from a JSON webhook payload, validating
// the Stripe-Signature header using the specified signing secret. Returns an error
// if the body or Stripe-Signature header provided are unreadable, if the
// signature doesn't match, or if the timestamp for the signature is older than
// DefaultTolerance.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
//
// This will return an error if the event API version does not match the
// stripe.APIVersion constant.
func ConstructEvent(payload []byte, header string, secret string) (stripe.Event, error) {
	return ConstructEventWithTolerance(payload, header, secret, DefaultTolerance)
}

// ConstructEventIgnoringTolerance initializes an Event object from a JSON webhook
// payload, validating the Stripe-Signature header using the specified signing secret.
// Returns an error if the body or Stripe-Signature header provided are unreadable or
// if the signature doesn't match. Does not check the signature's timestamp.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
//
// This will return an error if the event API version does not match the
// stripe.APIVersion constant.
func ConstructEventIgnoringTolerance(payload []byte, header string, secret string) (stripe.Event, error) {
	return constructEvent(payload, header, secret, ConstructEventOptions{IgnoreTolerance: true})
}

// ConstructEventWithTolerance initializes an Event object from a JSON webhook payload,
// validating the signature in the Stripe-Signature header using the specified signing
// secret and tolerance window. Returns an error if the body or Stripe-Signature header
// provided are unreadable, if the signature doesn't match, or if the timestamp
// for the signature is older than the specified tolerance.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
//
// This will return an error if the event API version does not match the
// stripe.APIVersion constant.
func ConstructEventWithTolerance(payload []byte, header string, secret string, tolerance time.Duration) (stripe.Event, error) {
	return constructEvent(payload, header, secret, ConstructEventOptions{Tolerance: tolerance})
}

// ConstructEventWithOptions initializes an Event object from a JSON webhook payload,
// validating the signature in the Stripe-Signature header using the specified signing
// secret and tolerance window provided by the options, if applicable.
//
// See `ConstructEventOptions` for more details on each of the options.
//
// Returns an error if the signature doesn't match, or:
//   - if `IgnoreTolerance` is false and the timestamp embedded in the event
//     header is not within the tolerance window (similar to `ConstructEventWithTolerance`)
//   - if `IgnoreAPIVersionMismatch` is false and the webhook event API version
//     does not match the API version of the stripe-go library, as defined in
//     `stripe.APIVersion`.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
func ConstructEventWithOptions(payload []byte, header string, secret string, options ConstructEventOptions) (stripe.Event, error) {
	return constructEvent(payload, header, secret, options)
}

// ValidatePayload validates the payload against the Stripe-Signature header
// using the specified signing secret. Returns an error if the body or
// Stripe-Signature header provided are unreadable, if the signature doesn't
// match, or if the timestamp for the signature is older than DefaultTolerance.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
func ValidatePayload(payload []byte, header string, secret string) error {
	return ValidatePayloadWithTolerance(payload, header, secret, DefaultTolerance)
}

// ValidatePayloadIgnoringTolerance validates the payload against the Stripe-Signature header
// using the specified signing secret. Returns an error if the body or
// Stripe-Signature header provided are unreadable or if the signature doesn't match.
// Does not check the signature's timestamp.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
func ValidatePayloadIgnoringTolerance(payload []byte, header string, secret string) error {
	return validatePayload(payload, header, secret, 0*time.Second, false)
}

// ValidatePayloadWithTolerance validates the payload against the Stripe-Signature header
// using the specified signing secret and tolerance window. Returns an error if the body
// or Stripe-Signature header provided are unreadable, if the signature doesn't match, or
// if the timestamp for the signature is older than the specified tolerance.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
func ValidatePayloadWithTolerance(payload []byte, header string, secret string, tolerance time.Duration) error {
	return validatePayload(payload, header, secret, tolerance, true)
}

type ConstructEventOptions struct {
	// Validates event timestamps using a custom Tolerance window. If this is
	// not set and `IgnoreTolerance` is false, will default to
	// `DefaultTolerance`.
	Tolerance time.Duration

	// If set to true, will ignore the `tolerance` option entirely and will not
	// check the event signature's timestamp. Defaults to false. When false,
	// constructing an event will fail with an error if the timestamp is not
	// within the `Tolerance` window.
	IgnoreTolerance bool

	// If set to true, will ignore validating whether an event's API version
	// matches the stripe-go API version. Defaults to false, returning an error
	// when there is a mismatch.
	IgnoreAPIVersionMismatch bool
}

//
// Private functions
//

func constructEvent(payload []byte, sigHeader string, secret string, options ConstructEventOptions) (stripe.Event, error) {
	// Event construction lives in the root stripe package, like signature
	// verification; this package only maps its options. An unset Tolerance
	// falls through to stripe's default, which equals DefaultTolerance.
	var opts []stripe.WebhookOption
	if options.IgnoreTolerance {
		opts = append(opts, stripe.WithIgnoreTolerance())
	} else if options.Tolerance != 0 {
		opts = append(opts, stripe.WithTolerance(options.Tolerance))
	}
	if options.IgnoreAPIVersionMismatch {
		opts = append(opts, stripe.WithIgnoreAPIVersionMismatch())
	}

	return stripe.ConstructEvent(payload, sigHeader, secret, opts...)
}

func validatePayload(payload []byte, sigHeader string, secret string, tolerance time.Duration, enforceTolerance bool) error {
	// Signature verification lives in the root stripe package; this package
	// delegates so there is only one implementation to maintain. Always pass an
	// explicit tolerance option: stripe.ValidatePayload defaults to a zero
	// tolerance that is enforced, which would reject nearly everything.
	opt := stripe.WithIgnoreTolerance()
	if enforceTolerance {
		opt = stripe.WithTolerance(tolerance)
	}

	return stripe.ValidatePayload(payload, sigHeader, secret, opt)
}

// For mocking webhook events
type UnsignedPayload struct {
	Payload   []byte
	Secret    string
	Timestamp time.Time
	Scheme    string
}

type SignedPayload struct {
	UnsignedPayload

	Signature []byte
	Header    string
}

func GenerateTestSignedPayload(options *UnsignedPayload) *SignedPayload {
	signedPayload := &SignedPayload{UnsignedPayload: *options}

	if signedPayload.Timestamp.IsZero() {
		signedPayload.Timestamp = time.Now()
	}

	if signedPayload.Scheme == "" {
		signedPayload.Scheme = "v1"
	}

	signedPayload.Signature = ComputeSignature(signedPayload.Timestamp, signedPayload.Payload, signedPayload.Secret)
	signedPayload.Header = generateHeader(*signedPayload)

	return signedPayload
}

func generateHeader(p SignedPayload) string {
	return fmt.Sprintf("t=%d,%s=%s", p.Timestamp.Unix(), p.Scheme, hex.EncodeToString(p.Signature))
}
