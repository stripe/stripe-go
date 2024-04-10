package webhook

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/stripe/stripe-go/v78"
)

//
// Public constants
//

const (
	// DefaultTolerance indicates that signatures older than this will be rejected by ConstructEvent.
	DefaultTolerance time.Duration = 300 * time.Second
	// signingVersion represents the version of the signature we currently use.
	signingVersion string = "v1"
)

//
// Public variables
//

// This block represents the list of errors that could be raised when using the webhook package.
var (
	ErrInvalidHeader    = errors.New("webhook has invalid Stripe-Signature header")
	ErrNoValidSignature = errors.New("webhook had no valid signature")
	ErrNotSigned        = errors.New("webhook has no Stripe-Signature header")
	ErrTooOld           = errors.New("timestamp wasn't within tolerance")
)

//
// Public functions
//

// ComputeSignature computes a webhook signature using Stripe's v1 signing
// method.
//
// See https://stripe.com/docs/webhooks#signatures for more information.
func ComputeSignature(t time.Time, payload []byte, secret string) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(fmt.Sprintf("%d", t.Unix())))
	mac.Write([]byte("."))
	mac.Write(payload)
	return mac.Sum(nil)
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
// Private types
//

type signedHeader struct {
	timestamp  time.Time
	signatures [][]byte
}

//
// Private functions
//

func constructEvent(payload []byte, sigHeader string, secret string, options ConstructEventOptions) (stripe.Event, error) {
	e := stripe.Event{}

	tolerance := options.Tolerance
	if options.Tolerance == 0 && !options.IgnoreTolerance {
		tolerance = DefaultTolerance
	}

	if err := validatePayload(payload, sigHeader, secret, tolerance, !options.IgnoreTolerance); err != nil {
		return e, err
	}

	if err := json.Unmarshal(payload, &e); err != nil {
		return e, fmt.Errorf("Failed to parse webhook body json: %s", err.Error())
	}

	if !options.IgnoreAPIVersionMismatch && e.APIVersion != stripe.APIVersion {
		return e, fmt.Errorf("Received event with API version %s, but stripe-go %s expects API version %s. We recommend that you create a WebhookEndpoint with this API version. Otherwise, you can disable this error by using `ConstructEventWithOptions(..., ConstructEventOptions{..., ignoreAPIVersionMismatch: true})`  but be wary that objects may be incorrectly deserialized.", e.APIVersion, stripe.ClientVersion, stripe.APIVersion)
	}

	return e, nil

}

func parseSignatureHeader(header string) (*signedHeader, error) {
	sh := &signedHeader{}

	if header == "" {
		return sh, ErrNotSigned
	}

	// Signed header looks like "t=1495999758,v1=ABC,v1=DEF,v0=GHI"
	pairs := strings.Split(header, ",")
	for _, pair := range pairs {
		parts := strings.Split(pair, "=")
		if len(parts) != 2 {
			return sh, ErrInvalidHeader
		}

		switch parts[0] {
		case "t":
			timestamp, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return sh, ErrInvalidHeader
			}
			sh.timestamp = time.Unix(timestamp, 0)

		case signingVersion:
			sig, err := hex.DecodeString(parts[1])
			if err != nil {
				continue // Ignore invalid signatures
			}

			sh.signatures = append(sh.signatures, sig)

		default:
			continue // Ignore unknown parts of the header
		}
	}

	if len(sh.signatures) == 0 {
		return sh, ErrNoValidSignature
	}

	return sh, nil
}

func validatePayload(payload []byte, sigHeader string, secret string, tolerance time.Duration, enforceTolerance bool) error {

	header, err := parseSignatureHeader(sigHeader)
	if err != nil {
		return err
	}

	expectedSignature := ComputeSignature(header.timestamp, payload, secret)
	expiredTimestamp := time.Since(header.timestamp) > tolerance
	if enforceTolerance && expiredTimestamp {
		return ErrTooOld
	}

	// Check all given v1 signatures, multiple signatures will be sent temporarily in the case of a rolled signature secret
	for _, sig := range header.signatures {
		if hmac.Equal(expectedSignature, sig) {
			return nil
		}
	}

	return ErrNoValidSignature
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

	if signedPayload.Timestamp == (time.Time{}) {
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
