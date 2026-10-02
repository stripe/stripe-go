//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

type RadarRule struct {
	// The action taken on the payment.
	Action string `json:"action"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object,omitempty"`
	// The predicate to evaluate the payment against.
	Predicate string `json:"predicate"`
}

// UnmarshalJSON handles deserialization of a RadarRule.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (r *RadarRule) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		r.ID = id
		return nil
	}

	type radarRule RadarRule
	var v radarRule
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*r = RadarRule(v)
	return nil
}
