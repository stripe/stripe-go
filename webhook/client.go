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

	"github.com/stripe/stripe-go"
)

const (
	// Signatures older than this will be rejected by ConstructEvent
	DefaultTolerance time.Duration = 300 * time.Second
	signingVersion   string        = "v1"
)

var (
	ErrNotSigned        error = errors.New("Webhook has no Stripe-Signature header")
	ErrInvalidHeader    error = errors.New("Webhook has invalid Stripe-Signature header")
	ErrTooOld           error = errors.New("Timestamp wasn't within tolerance")
	ErrNoValidSignature error = errors.New("Webhook had no valid signature")
)

// Computes a webhook signature using Stripe's v1 signing method. See
// https://stripe.com/docs/webhooks#signatures
func computeSignature(t time.Time, payload []byte, secret string) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(fmt.Sprintf("%d", t.Unix())))
	mac.Write([]byte("."))
	mac.Write(payload)
	return mac.Sum(nil)
}

type signedHeader struct {
	timestamp  time.Time
	signatures [][]byte
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

// Initializes an Event object from a JSON webhook payload, validating the
// Stripe-Signature header using the specified signing secret. Returns an error
// if the body or Stripe-Signature header provided are unreadable, if the
// signature doesn't match, or if the timestamp for the signature is older than
// DefaultTolerance.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
//
func ConstructEvent(payload []byte, header string, secret string) (stripe.Event, error) {
	return ConstructEventWithTolerance(payload, header, secret, DefaultTolerance)
}

// Initializes an Event object from a JSON webhook payload, validating the
// signature in the Stripe-Signature header using the specified signing secret
// and tolerance window. Returns an error if the body or Stripe-Signature header
// provided are unreadable, if the signature doesn't match, or if the timestamp
// for the signature is older than the specified tolerance.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
//
func ConstructEventWithTolerance(payload []byte, header string, secret string, tolerance time.Duration) (stripe.Event, error) {
	return constructEvent(payload, header, secret, tolerance, true)
}

// Initializes an Event object from a JSON webhook payload, validating the
// Stripe-Signature header using the specified signing secret. Returns an error
// if the body or Stripe-Signature header provided are unreadable or if the
// signature doesn't match. Does not check the signature's timestamp.
//
// NOTE: Stripe will only send Webhook signing headers after you have retrieved
// your signing secret from the Stripe dashboard:
// https://dashboard.stripe.com/webhooks
//
func ConstructEventIgnoringTolerance(payload []byte, header string, secret string) (stripe.Event, error) {
	return constructEvent(payload, header, secret, 0*time.Second, false)
}

func constructEvent(payload []byte, sigHeader string, secret string, tolerance time.Duration, enforceTolerance bool) (stripe.Event, error) {
	e := stripe.Event{}

	if err := json.Unmarshal(payload, &e); err != nil {
		return e, fmt.Errorf("Failed to parse webhook body json: %s", err.Error())
	}

	header, err := parseSignatureHeader(sigHeader)
	if err != nil {
		return e, err
	}

	expectedSignature := computeSignature(header.timestamp, payload, secret)
	expiredTimestamp := time.Since(header.timestamp) > tolerance
	if enforceTolerance && expiredTimestamp {
		return e, ErrTooOld
	}

	// Check all given v1 signatures, multiple signatures will be sent temporarily in the case of a rolled signature secret
	for _, sig := range header.signatures {
		if hmac.Equal(expectedSignature, sig) {
			return e, nil
		}
	}

	return e, ErrNoValidSignature
}
