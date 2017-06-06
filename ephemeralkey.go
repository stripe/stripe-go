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
    AssociatedObject []struct {
        ID   string `json:"id"`
        Type string `json:"type"`
    } `json:"associated_objects"`
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
	} else {
		// the id is surrounded by "\" characters, so strip them
		e.ID = string(data[1 : len(data)-1])
	}

	return nil
}
