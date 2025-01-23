//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"time"
)

// The type of data collected by the reader.
type TerminalReaderCollectedDataType string

// List of values that TerminalReaderCollectedDataType can take
const (
	TerminalReaderCollectedDataTypeMagstripe TerminalReaderCollectedDataType = "magstripe"
)

// Retrieve data collected using Reader hardware.
type TerminalReaderCollectedDataParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *TerminalReaderCollectedDataParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The magstripe data collected by the reader.
type TerminalReaderCollectedDataMagstripe struct {
	// The raw magstripe data collected by the reader.
	Data string `json:"data"`
}

// Returns data collected by Terminal readers. This data is only stored for 24 hours.
type TerminalReaderCollectedData struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created time.Time `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// The magstripe data collected by the reader.
	Magstripe *TerminalReaderCollectedDataMagstripe `json:"magstripe"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The type of data collected by the reader.
	Type TerminalReaderCollectedDataType `json:"type"`
}

// UnmarshalJSON handles deserialization of a TerminalReaderCollectedData.
// This custom unmarshaling is needed to handle the time fields correctly.
func (t *TerminalReaderCollectedData) UnmarshalJSON(data []byte) error {
	type terminalReaderCollectedData TerminalReaderCollectedData
	v := struct {
		Created int64 `json:"created"`
		*terminalReaderCollectedData
	}{
		terminalReaderCollectedData: (*terminalReaderCollectedData)(t),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	t.Created = time.Unix(v.Created, 0)
	return nil
}

// MarshalJSON handles serialization of a TerminalReaderCollectedData.
// This custom marshaling is needed to handle the time fields correctly.
func (t TerminalReaderCollectedData) MarshalJSON() ([]byte, error) {
	type terminalReaderCollectedData TerminalReaderCollectedData
	v := struct {
		Created int64 `json:"created"`
		terminalReaderCollectedData
	}{
		terminalReaderCollectedData: (terminalReaderCollectedData)(t),
		Created:                     t.Created.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}
