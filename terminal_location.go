package stripe

// TerminalLocationParams is the set of parameters that can be used when creating or updating a terminal location.
type TerminalLocationParams struct {
	Params            `form:"*"`
	DisplayName				*string											`form:"display_name"`
	Address           *AccountAddressParams       `form:"address"`
	OperatorAccount   *string 										`form:"operator_account"`
}

// TerminalLocationListParams is the set of parameters that can be used when listing temrinal locations.
type TerminalLocationListParams struct {
	ListParams				`form:"*"`
	OperatorAccount   *string 		`form:"operator_account"`
}

// TerminalLocation is the resource representing a Stripe terminal location.
type TerminalLocation struct {
	ID				  	string											`json:"id"`
	Object				string											`json:"object"`
	Address       *AccountAddressParams       `json:"address"`
}

// TerminalLocationList is a list of terminal readers as retrieved from a list endpoint.
type TerminalLocationList struct {
	ListMeta
	Data []*TerminalLocation	 `json:"data"`
}
