//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// TerminalLocationParams is the set of parameters that can be used when creating or updating a terminal location.
type TerminalLocationParams struct {
	Params      `form:"*"`
	Address     *AccountAddressParams `form:"address"`
	DisplayName *string               `form:"display_name"`
}

// TerminalLocationListParams is the set of parameters that can be used when listing temrinal locations.
type TerminalLocationListParams struct {
	ListParams `form:"*"`
}

// TerminalLocation is the resource representing a Stripe terminal location.
type TerminalLocation struct {
	APIResource
	Address     *AccountAddressParams `json:"address"`
	Deleted     bool                  `json:"deleted"`
	DisplayName string                `json:"display_name"`
	ID          string                `json:"id"`
	Livemode    bool                  `json:"livemode"`
	Metadata    map[string]string     `json:"metadata"`
	Object      string                `json:"object"`
}

// TerminalLocationList is a list of terminal readers as retrieved from a list endpoint.
type TerminalLocationList struct {
	APIResource
	ListMeta
	Data []*TerminalLocation `json:"data"`
}

// UnmarshalJSON handles deserialization of a TerminalLocation.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (t *TerminalLocation) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		t.ID = id
		return nil
	}

	type terminalLocation TerminalLocation
	var v terminalLocation
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*t = TerminalLocation(v)
	return nil
}
