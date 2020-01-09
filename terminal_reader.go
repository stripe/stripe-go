package stripe

// TerminalReaderParams is the set of parameters that can be used for creating or updating a terminal reader.
type TerminalReaderParams struct {
	Params           `form:"*"`
	Label            *string `form:"label"`
	Location         *string `form:"location"`
	RegistrationCode *string `form:"registration_code"`
}

// TerminalReaderGetParams is the set of parameters that can be used to get a terminal reader.
type TerminalReaderGetParams struct {
	Params `form:"*"`
}

// TerminalReaderListParams is the set of parameters that can be used when listing temrinal readers.
type TerminalReaderListParams struct {
	ListParams `form:"*"`
	DeviceType *string `form:"device_type"`
	Location   *string `form:"location"`
	Status     *string `form:"status"`
}

// TerminalReader is the resource representing a Stripe terminal reader.
type TerminalReader struct {
	Deleted         bool              `json:"deleted"`
	DeviceSwVersion string            `json:"device_sw_version"`
	DeviceType      string            `json:"device_type"`
	ID              string            `json:"id"`
	IPAddress       string            `json:"ip_address"`
	Label           string            `json:"label"`
	Livemode        bool              `json:"livemode"`
	Location        string            `json:"location"`
	Metadata        map[string]string `json:"metadata"`
	Object          string            `json:"object"`
	SerialNumber    string            `json:"serial_number"`
	Status          string            `json:"status"`
}

// TerminalReaderList is a list of terminal readers as retrieved from a list endpoint.
type TerminalReaderList struct {
	ListMeta
	Data     []*TerminalReader `json:"data"`
	Location *string           `json:"location"`
	Status   *string           `json:"status"`
}
