//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves the details of a Report Type. (Certain report types require a [live-mode API key](https://stripe.com/docs/keys#test-live-modes).)
type ReportTypeParams struct {
	Params `form:"*"`
}

// Returns a full list of Report Types.
type ReportTypeListParams struct {
	ListParams `form:"*"`
}

// The Report Type resource corresponds to a particular type of report, such as
// the "Activity summary" or "Itemized payouts" reports. These objects are
// identified by an ID belonging to a set of enumerated values. See
// [API Access to Reports documentation](https://stripe.com/docs/reporting/statements/api)
// for those Report Type IDs, along with required and optional parameters.
//
// Note that certain report types can only be run based on your live-mode data (not test-mode
// data), and will error when queried without a [live-mode API key](https://stripe.com/docs/keys#test-live-modes).
type ReportType struct {
	APIResource
	Created            int64    `json:"created"`
	DataAvailableEnd   int64    `json:"data_available_end"`
	DataAvailableStart int64    `json:"data_available_start"`
	DefaultColumns     []string `json:"default_columns"`
	ID                 string   `json:"id"`
	Livemode           bool     `json:"livemode"`
	Name               string   `json:"name"`
	Object             string   `json:"object"`
	Updated            int64    `json:"updated"`
	Version            int64    `json:"version"`
}

// ReportTypeList is a list of ReportTypes as retrieved from a list endpoint.
type ReportTypeList struct {
	APIResource
	ListMeta
	Data []*ReportType `json:"data"`
}
