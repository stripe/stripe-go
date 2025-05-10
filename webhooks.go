package stripe

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
)

//
// Public constants
//

const (
	// DefaultTolerance indicates that signatures older than this will be rejected by ConstructEvent.
	WebhookDefaultTolerance time.Duration = 300 * time.Second
	// signingVersion represents the version of the signature we currently use.
	signingVersion string = "v1"
)

//
// Public variables
//

// This block represents the list of errors that could be raised when using the webhook package.
var (
	ErrWebhookInvalidHeader    = errors.New("webhook has invalid Stripe-Signature header")
	ErrWebhookNoValidSignature = errors.New("webhook had no valid signature")
	ErrWebhookNotSigned        = errors.New("webhook has no Stripe-Signature header")
	ErrWebhookTooOld           = errors.New("timestamp wasn't within tolerance")
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
// APIVersion constant.
func ConstructEvent(payload []byte, header string, secret string, opts ...WebhookOption) (Event, error) {
	cfg := webhookConfig{
		Tolerance: WebhookDefaultTolerance,
	}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(&cfg)
	}
	return constructEvent(payload, header, secret, cfg)
}

// ValidatePayload validates the payload against the Stripe-Signature header
// using the specified signing secret. Returns an error if the body or
// Stripe-Signature header provided are unreadable, if the signature doesn't
// match, or if the timestamp for the signature is older than DefaultTolerance.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
func ValidatePayload(payload []byte, header string, secret string, opts ...WebhookOption) error {
	cfg := webhookConfig{
		Tolerance: WebhookDefaultTolerance,
	}
	for _, opt := range opts {
		if opt == nil {
			continue
		}
		opt(&cfg)
	}
	return validatePayload(payload, header, secret, cfg)
}

type WebhookOption func(*webhookConfig)

// WithTolerance validates event timestamps using a custom Tolerance window. If this is
// not set and `IgnoreTolerance` is false, will default to
// `WebhookDefaultTolerance`.
func WithTolerance(tolerance time.Duration) WebhookOption {
	return func(w *webhookConfig) {
		w.Tolerance = tolerance
	}
}

// WithIgnoreTolerance will ignore the the event signature's timestamp.
func WithIgnoreTolerance() WebhookOption {
	return func(w *webhookConfig) {
		w.IgnoreTolerance = true
	}
}

// WithIgnoreAPIVersionMismatch will ignore validating whether an event's API version
// matches the stripe-go API version. This is currently only used for ConstructEvent.
func WithIgnoreAPIVersionMismatch() WebhookOption {
	return func(w *webhookConfig) {
		w.IgnoreAPIVersionMismatch = true
	}
}

type webhookConfig struct {
	Tolerance                time.Duration
	IgnoreTolerance          bool
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

func isCompatibleAPIVersion(sdkAPIVersion, eventAPIVersion string) bool {
	// If the event api version is from before we started adding
	// a release train, there's no way its compatible with this
	// version
	if !strings.Contains(eventAPIVersion, ".") {
		return false
	}

	// if the SDK is pinned to a preview version, the event's API version must match exactly
	var currentReleaseTrain = strings.Split(sdkAPIVersion, ".")[1]
	if currentReleaseTrain == "preview" {
		return sdkAPIVersion == eventAPIVersion
	}

	// versions are yyyy-MM-dd.train
	var eventReleaseTrain = strings.Split(eventAPIVersion, ".")[1]
	return eventReleaseTrain == currentReleaseTrain
}

func constructEvent(payload []byte, sigHeader string, secret string, cfg webhookConfig) (Event, error) {
	e := Event{}

	if err := validatePayload(payload, sigHeader, secret, cfg); err != nil {
		return e, err
	}

	if err := json.Unmarshal(payload, &e); err != nil {
		return e, fmt.Errorf("Failed to parse webhook body json: %s", err.Error())
	}

	if !cfg.IgnoreAPIVersionMismatch && !isCompatibleAPIVersion(APIVersion, e.APIVersion) {
		return e, fmt.Errorf("Received event with API version %s, but stripe-go %s expects API version %s. We recommend that you create a WebhookEndpoint with this API version. Otherwise, you can disable this error by using `ConstructEventWithOptions(..., ConstructEventOptions{..., ignoreAPIVersionMismatch: true})`  but be wary that objects may be incorrectly deserialized.", e.APIVersion, ClientVersion, APIVersion)
	}

	return e, nil

}

func parseSignatureHeader(header string) (*signedHeader, error) {
	sh := &signedHeader{}

	if header == "" {
		return sh, ErrWebhookNotSigned
	}

	// Signed header looks like "t=1495999758,v1=ABC,v1=DEF,v0=GHI"
	pairs := strings.Split(header, ",")
	for _, pair := range pairs {
		parts := strings.Split(pair, "=")
		if len(parts) != 2 {
			return sh, ErrWebhookInvalidHeader
		}

		switch parts[0] {
		case "t":
			timestamp, err := strconv.ParseInt(parts[1], 10, 64)
			if err != nil {
				return sh, ErrWebhookInvalidHeader
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
		return sh, ErrWebhookNoValidSignature
	}

	return sh, nil
}

func validatePayload(payload []byte, sigHeader string, secret string, cfg webhookConfig) error {

	header, err := parseSignatureHeader(sigHeader)
	if err != nil {
		return err
	}

	expiredTimestamp := time.Since(header.timestamp) > cfg.Tolerance
	if !cfg.IgnoreTolerance && expiredTimestamp {
		return ErrWebhookTooOld
	}

	expectedSignature := ComputeSignature(header.timestamp, payload, secret)
	// Check all given v1 signatures, multiple signatures will be sent temporarily in the case of a rolled signature secret
	for _, sig := range header.signatures {
		if hmac.Equal(expectedSignature, sig) {
			return nil
		}
	}

	return ErrWebhookNoValidSignature
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
