//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Retrieves a Location object.
type TerminalLocationParams struct {
	Params `form:"*"`
	// The full address of the location.
	Address *AddressParams `form:"address"`
	// The ID of a configuration that will be used to customize all readers in this location.
	ConfigurationOverrides *string `form:"configuration_overrides"`
	// A name for the location.
	DisplayName *string `form:"display_name"`
}

// Returns a list of Location objects.
type TerminalLocationListParams struct {
	ListParams `form:"*"`
}

// A Location represents a grouping of readers.
//
// Related guide: [Fleet Management](https://stripe.com/docs/terminal/fleet/locations).
type TerminalLocation struct {
	APIResource
	Address *Address `json:"address"`
	// The ID of a configuration that will be used to customize all readers in this location.
	ConfigurationOverrides string `json:"configuration_overrides"`
	Deleted                bool   `json:"deleted"`
	// The display name of the location.
	DisplayName string `json:"display_name"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}

// TerminalLocationList is a list of Locations as retrieved from a list endpoint.
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
