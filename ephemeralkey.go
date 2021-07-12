//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// EphemeralKeyParams is the set of parameters that can be used when creating
// an ephemeral key.
type EphemeralKeyParams struct {
	Params        `form:"*"`
	Customer      *string `form:"customer"`
	IssuingCard   *string `form:"issuing_card"`
	StripeVersion *string `form:"-"` // This goes in the `Stripe-Version` header
}

// EphemeralKey is the resource representing a Stripe ephemeral key. This is used by Mobile SDKs
// to for example manage a Customer's payment methods.
type EphemeralKey struct {
	APIResource
	AssociatedObjects []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"associated_objects"`
	Created  int64  `json:"created"`
	Expires  int64  `json:"expires"`
	ID       string `json:"id"`
	Livemode bool   `json:"livemode"`
	Object   string `json:"object"`
	Secret   string `json:"secret"`
	// RawJSON is provided so that it may be passed back to the frontend
	// unchanged.  Ephemeral keys are issued on behalf of another client which
	// may be running a different version of the bindings and thus expect a
	// different JSON structure.  This ensures that if the structure differs
	// from the version of these bindings, we can still pass back a compatible
	// key.
	RawJSON []byte `json:"-"`
}

// UnmarshalJSON handles deserialization of an EphemeralKey.
// This custom unmarshaling is needed because we need to store the
// raw JSON on the object so it may be passed back to the frontend.

func (e *EphemeralKey) UnmarshalJSON(data []byte) error {
	type ephemeralKey EphemeralKey
	var ee ephemeralKey
	err := json.Unmarshal(data, &ee)
	if err == nil {
		*e = EphemeralKey(ee)
	}

	// Go does guarantee the longevity of `data`, so copy when assigning `RawJSON`
	// See https://golang.org/pkg/encoding/json/#Unmarshaler
	// and https://github.com/stripe/stripe-go/pull/1142
	e.RawJSON = append(e.RawJSON[:0], data...)

	return nil
}
