package stripe

// TerminalConnectionTokenParams is the set of parameters that can be used when creating a terminal connection token.
type TerminalConnectionTokenParams struct {
	Params   `form:"*"`
	Location string `form:"location"`
}

// TerminalConnectionToken is the resource representing a Stripe terminal connection token.
type TerminalConnectionToken struct {
	Location string `json:"location"`
	Object   string `json:"object"`
	Secret   string `json:"secret"`
}
