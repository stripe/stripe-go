package stripe

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
	ListMeta
	Data []*TerminalLocation `json:"data"`
}
