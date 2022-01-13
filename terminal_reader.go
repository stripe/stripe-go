//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Type of reader, one of `bbpos_chipper2x`, `bbpos_wisepos_e`, or `verifone_P400`.
type TerminalReaderDeviceType string

// List of values that TerminalReaderDeviceType can take
const (
	TerminalReaderDeviceTypeBBPOSChipper2X TerminalReaderDeviceType = "bbpos_chipper2x"
	TerminalReaderDeviceTypeBBPOSWisePOSE  TerminalReaderDeviceType = "bbpos_wisepos_e"
	TerminalReaderDeviceTypeVerifoneP400   TerminalReaderDeviceType = "verifone_P400"
)

// Updates a Reader object by setting the values of the parameters passed. Any parameters not provided will be left unchanged.
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

// Returns a list of Reader objects.
type TerminalReaderListParams struct {
	ListParams `form:"*"`
	DeviceType *string `form:"device_type"`
	Location   *string `form:"location"`
	Status     *string `form:"status"`
}

// A Reader represents a physical device for accepting payment details.
//
// Related guide: [Connecting to a Reader](https://stripe.com/docs/terminal/payments/connect-reader).
type TerminalReader struct {
	APIResource
	Deleted         bool                     `json:"deleted"`
	DeviceSwVersion string                   `json:"device_sw_version"`
	DeviceType      TerminalReaderDeviceType `json:"device_type"`
	ID              string                   `json:"id"`
	IPAddress       string                   `json:"ip_address"`
	Label           string                   `json:"label"`
	Livemode        bool                     `json:"livemode"`
	Location        string                   `json:"location"`
	Metadata        map[string]string        `json:"metadata"`
	Object          string                   `json:"object"`
	SerialNumber    string                   `json:"serial_number"`
	Status          string                   `json:"status"`
}

// TerminalReaderList is a list of Readers as retrieved from a list endpoint.
type TerminalReaderList struct {
	APIResource
	ListMeta
	Data     []*TerminalReader `json:"data"`
	Location *string           `json:"location"`
	Status   *string           `json:"status"`
}
