package stripe

import "encoding/json"

// EphemeralKeyParams is the set of parameters that can be used when creating
// an ephemeral key.
// For more details see https://stripe.com/docs/api#ephemeral_keys.
type EphemeralKeyParams struct {
	Params
	Customer string
}

// EphemeralKey is the resource representing a Stripe ephemeral key.
// For more details see https://stripe.com/docs/api#ephemeral_keys.
type EphemeralKey struct {
	ID      string `json:"id"`
	Created int64  `json:"created"`
	Expires int64  `json:"expires"`
	Live    bool   `json:"livemode"`
    AssociatedObjects []struct {
        ID   string `json:"id"`
        Type string `json:"type"`
    } `json:"associated_objects"`
    RawJSON []byte
}

// UnmarshalJSON handles deserialization of an EphemeralKey.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (e *EphemeralKey) UnmarshalJSON(data []byte) error {
	type ephemeralKey EphemeralKey
	var ee ephemeralKey
	err := json.Unmarshal(data, &ee)
	if err == nil {
		*e = EphemeralKey(ee)
	}

    // Store the raw JSON so it can be passed back to the frontend
    e.RawJSON = data

	return nil
}
