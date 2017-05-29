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

// computeSignature computes a signature using stripe's v1 signing method.
func computeSignature(t string, payload []byte, secret string) []byte {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(t))
	mac.Write([]byte("."))
	mac.Write(payload)
	return mac.Sum(nil)
}

const ToleranceDefault int64 = 300
const ToleranceIgnoreTimestamp int64 = 0

var ErrNotSigned error = errors.New("Webhook has no Stripe-Signature header")
var ErrNoTimestamp error = errors.New("Webhook has no timestamp")
var ErrTooOld error = errors.New("Webhook had valid signature but timestamp wasn't within tolerance")
var ErrNoValidSignature error = errors.New("Webhook had no valid signature")

// ConstructEvent initializes an Event object from a JSON payload.
// It returns an non-nil error when the payload is not valid JSON or when signature verification fails.
// payload is the webhook request body, i.e. `ioutil.ReadAll(r.Body)`.
// sigHeader is the webhook Stripe-Signature header, i.e. `r.Header.Get("Stripe-Signature")`.
// secret is your Signing Secret, i.e. `"whsec_XYZ"`.  See https://dashboard.stripe.com/webhooks
// tolerance (suggested 300) is the max difference in seconds between now and Stripe-Signature's timestamp. If the difference is greater than this tolerance, the signature is rejected and a non-nil error is returned.  If tolerance is 0 or less, then the timestamp is not checked.
// NOTE: your requests will only have Stripe-Signature if you have clicked to reveal your secret
func ConstructEvent(payload []byte, sigHeader string, secret string, tolerance int64) (stripe.Event, error) {
	e := stripe.Event{}

	if err := json.Unmarshal(payload, &e); err != nil {
		return e, fmt.Errorf("Failed to parse webhook body json: %s", err.Error())
	}

	if sigHeader == "" {
		return e, ErrNotSigned
	}

	// sigHeader looks like "t=1495999758,v1=ABC,v1=DEF,v0=GHI"

	// First extract the timestamp
	var t string
	pairs := strings.Split(sigHeader, ",")
	for _, pair := range pairs {
		parts := strings.Split(pair, "=")
		if len(parts) != 2 || parts[0] != "t" {
			continue
		}

		t = parts[1]
		break
	}

	if t == "" {
		return e, ErrNoTimestamp
	}

	invalidTimestamp := false
	if tolerance > 0 {
		timestamp, err := strconv.ParseInt(t, 10, 64)
		if err == nil {
			currentTimestamp := time.Now().Unix()

			diff := timestamp - currentTimestamp
			if diff < 0 {
				diff = -diff
			}

			if diff > tolerance {
				invalidTimestamp = true
			}
		} else {
			invalidTimestamp = true
		}
	}

	expectedSignature := computeSignature(t, payload, secret)

	// Check all given v1 signatures since multiple v1 can happen in case of rolled secret.
	for _, pair := range pairs {
		parts := strings.Split(pair, "=")
		if len(parts) != 2 || parts[0] != SigningVersion {
			continue
		}

		sig, err := hex.DecodeString(parts[1])
		if err != nil {
			// Ignore signature which isn't valid hex.
			continue
		}

		if hmac.Equal(expectedSignature, sig) {
			if invalidTimestamp {
				return e, ErrTooOld
			}

			// OK
			return e, nil
		}
	}

	return e, ErrNoValidSignature
}
