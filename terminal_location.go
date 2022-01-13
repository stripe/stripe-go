//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves a Location object.
type TerminalLocationParams struct {
	Params      `form:"*"`
	Address     *AccountAddressParams `form:"address"`
	DisplayName *string               `form:"display_name"`
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
	Address     *AccountAddressParams `json:"address"`
	Deleted     bool                  `json:"deleted"`
	DisplayName string                `json:"display_name"`
	ID          string                `json:"id"`
	Livemode    bool                  `json:"livemode"`
	Metadata    map[string]string     `json:"metadata"`
	Object      string                `json:"object"`
}

// TerminalLocationList is a list of Locations as retrieved from a list endpoint.
type TerminalLocationList struct {
	APIResource
	ListMeta
	Data []*TerminalLocation `json:"data"`
}
