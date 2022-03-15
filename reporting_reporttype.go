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
	Created int64 `json:"created"`
	// Most recent time for which this Report Type is available. Measured in seconds since the Unix epoch.
	DataAvailableEnd int64 `json:"data_available_end"`
	// Earliest time for which this Report Type is available. Measured in seconds since the Unix epoch.
	DataAvailableStart int64 `json:"data_available_start"`
	// List of column names that are included by default when this Report Type gets run. (If the Report Type doesn't support the `columns` parameter, this will be null.)
	DefaultColumns []string `json:"default_columns"`
	// The [ID of the Report Type](https://stripe.com/docs/reporting/statements/api#available-report-types), such as `balance.summary.1`.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Human-readable name of the Report Type
	Name string `json:"name"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// When this Report Type was latest updated. Measured in seconds since the Unix epoch.
	Updated int64 `json:"updated"`
	// Version of the Report Type. Different versions report with the same ID will have the same purpose, but may take different run parameters or have different result schemas.
	Version int64 `json:"version"`
}

// ReportTypeList is a list of ReportTypes as retrieved from a list endpoint.
type ReportTypeList struct {
	APIResource
	ListMeta
	Data []*ReportType `json:"data"`
}
