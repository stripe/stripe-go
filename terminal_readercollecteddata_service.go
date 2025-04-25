//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v1TerminalReaderCollectedDataService is used to invoke /v1/terminal/reader_collected_data APIs.
type v1TerminalReaderCollectedDataService struct {
	B   Backend
	Key string
}

// Retrieve data collected using Reader hardware.
func (c v1TerminalReaderCollectedDataService) Retrieve(ctx context.Context, id string, params *TerminalReaderCollectedDataRetrieveParams) (*TerminalReaderCollectedData, error) {
	path := FormatURLPath("/v1/terminal/reader_collected_data/%s", id)
	readercollecteddata := &TerminalReaderCollectedData{}
	if params == nil {
		params = &TerminalReaderCollectedDataRetrieveParams{}
	}
	params.Context = ctx
	err := c.B.Call(http.MethodGet, path, c.Key, params, readercollecteddata)
	return readercollecteddata, err
}
