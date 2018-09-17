package stripe

// TerminalReaderRegisterParams is the set of parameters that can be used for creating a terminal reader.
type TerminalReaderRegisterParams struct {
	Params            `form:"*"`
	RegistrationCode	*string			`form:"registration_code"`
	Label							*string			`form:"label"`
	OperatorAccount   *string 		`form:"operator_account"`
	Location					*string			`form:"location"`
}

// TerminalReaderGetParams is the set of parameters that can be used to get a terminal reader.
type TerminalReaderGetParams struct {
	Params            `form:"*"`
	OperatorAccount		*string			`form:"operator_account"`
}

// TerminalReaderUpdateParams is the set of parameters that can be used to update a terminal reader.
type TerminalReaderUpdateParams struct {
	Params            `form:"*"`
	Label			        *string 		`form:"label"`
	OperatorAccount		*string			`form:"operator_account"`
}

// TerminalReaderListParams is the set of parameters that can be used when listing temrinal readers.
type TerminalReaderListParams struct {
	ListParams							`form:"*"`
	IncludeOfflineReaders		*string			`form:"include_offline_readers"`
	OperatorAccount					*string			`form:"operator_account"`
	Location								*string			`form:"location"`
}

// TerminalReader is the resource representing a Stripe terminal reader.
type TerminalReader struct {
	ID						string			`json:"id"`
	Object				string			`json:"object"`
	DeviceType		string			`json:"device_type"`
	SerialNumber	string			`json:"serial_number"`
	label					string			`json:"label"`
	IPAddress			string			`json:"ip_address"`
}

// TerminalReaderList is a list of terminal readers as retrieved from a list endpoint.
type TerminalReaderList struct {
	ListMeta
	Data []*TerminalReader `json:"data"`
}
