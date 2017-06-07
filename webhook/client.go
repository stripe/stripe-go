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

const SigningVersion string = "v1"
const TestmodeSigningVersion string = "v0"

// computeSignature computes a signature using stripe's v1 signing method.
func computeSignature(t time.Time, payload []byte, secret string) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(fmt.Sprintf("%d", t.Unix())))
	mac.Write([]byte("."))
	mac.Write(payload)
	return mac.Sum(nil)
}

const DefaultTolerance time.Duration = 300 * time.Second

var ErrNotSigned error = errors.New("Webhook has no Stripe-Signature header")
var ErrInvalidHeader error = errors.New("Webhook has invalid Stripe-Signature header")
var ErrTooOld error = errors.New("Timestamp wasn't within tolerance")
var ErrNoValidSignature error = errors.New("Webhook had no valid signature")

type signedHeader struct {
	timestamp time.Time
	signatures [][]byte
	testmode bool
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

		case SigningVersion:
			sig, err := hex.DecodeString(parts[1])
			if err != nil {
				return sh, ErrInvalidHeader
			}

			sh.signatures = append(sh.signatures, sig)

		case TestmodeSigningVersion:
			sh.testmode = true

		default:
			return sh, ErrInvalidHeader
		}
	}

	return sh, nil
}

func ValidateEvent(payload []byte, sigHeader string, secret string) (stripe.Event, error) {
	return ValidateEventWithTolerance(payload, sigHeader, secret, DefaultTolerance)
}

func ValidateEventWithTolerance(payload []byte, sigHeader string, secret string, tolerance time.Duration) (stripe.Event, error) {
	return validateEvent(payload, sigHeader, secret, tolerance, true)
}

func ValidateEventIgnoringTolerance(payload []byte, sigHeader string, secret string) (stripe.Event, error) {
	return validateEvent(payload, sigHeader, secret, 0 * time.Second, false)
}


// ConstructEvent initializes an Event object from a JSON payload.
// It returns an non-nil error when the payload is not valid JSON or when signature verification fails.
// payload is the webhook request body, i.e. `ioutil.ReadAll(r.Body)`.
// sigHeader is the webhook Stripe-Signature header, i.e. `r.Header.Get("Stripe-Signature")`.
// secret is your Signing Secret, i.e. `"whsec_XYZ"`.  See https://dashboard.stripe.com/webhooks
// tolerance (suggested 300) is the max difference in seconds between now and Stripe-Signature's timestamp. If the difference is greater than this tolerance, the signature is rejected and a non-nil error is returned.  If tolerance is 0 or less, then the timestamp is not checked.
// NOTE: your requests will only have Stripe-Signature if you have clicked to reveal your secret
func validateEvent(payload []byte, sigHeader string, secret string, tolerance time.Duration, enforceTolerance bool) (stripe.Event, error) {	e := stripe.Event{}

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
