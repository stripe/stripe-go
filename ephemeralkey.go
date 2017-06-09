package stripe

import "encoding/json"

// EphemeralKeyParams is the set of parameters that can be used when creating
// an ephemeral key.
// For more details see https://stripe.com/docs/api#ephemeral_keys.
type EphemeralKeyParams struct {
	Params
	Customer      string
	StripeVersion string
}

// EphemeralKey is the resource representing a Stripe ephemeral key.
// For more details see https://stripe.com/docs/api#ephemeral_keys.
type EphemeralKey struct {
	ID                string `json:"id"`
	Created           int64  `json:"created"`
	Expires           int64  `json:"expires"`
	Live              bool   `json:"livemode"`
	AssociatedObjects []struct {
		ID   string `json:"id"`
		Type string `json:"type"`
	} `json:"associated_objects"`
	RawJSON []byte
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

	e.RawJSON = data

	return nil
}
