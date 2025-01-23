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

// Returns a full list of Report Types.
type ReportingReportTypeListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *ReportingReportTypeListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the details of a Report Type. (Certain report types require a [live-mode API key](https://stripe.com/docs/keys#test-live-modes).)
type ReportingReportTypeParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *ReportingReportTypeParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The Report Type resource corresponds to a particular type of report, such as
// the "Activity summary" or "Itemized payouts" reports. These objects are
// identified by an ID belonging to a set of enumerated values. See
// [API Access to Reports documentation](https://stripe.com/docs/reporting/statements/api)
// for those Report Type IDs, along with required and optional parameters.
//
// Note that certain report types can only be run based on your live-mode data (not test-mode
// data), and will error when queried without a [live-mode API key](https://stripe.com/docs/keys#test-live-modes).
type ReportingReportType struct {
	APIResource
	// Most recent time for which this Report Type is available. Measured in seconds since the Unix epoch.
	DataAvailableEnd time.Time `json:"data_available_end"`
	// Earliest time for which this Report Type is available. Measured in seconds since the Unix epoch.
	DataAvailableStart time.Time `json:"data_available_start"`
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
	Updated time.Time `json:"updated"`
	// Version of the Report Type. Different versions report with the same ID will have the same purpose, but may take different run parameters or have different result schemas.
	Version int64 `json:"version"`
}

// ReportingReportTypeList is a list of ReportTypes as retrieved from a list endpoint.
type ReportingReportTypeList struct {
	APIResource
	ListMeta
	Data []*ReportingReportType `json:"data"`
}

// UnmarshalJSON handles deserialization of a ReportingReportType.
// This custom unmarshaling is needed to handle the time fields correctly.
func (r *ReportingReportType) UnmarshalJSON(data []byte) error {
	type reportingReportType ReportingReportType
	v := struct {
		DataAvailableEnd   int64 `json:"data_available_end"`
		DataAvailableStart int64 `json:"data_available_start"`
		Updated            int64 `json:"updated"`
		*reportingReportType
	}{
		reportingReportType: (*reportingReportType)(r),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	r.DataAvailableEnd = time.Unix(v.DataAvailableEnd, 0)
	r.DataAvailableStart = time.Unix(v.DataAvailableStart, 0)
	r.Updated = time.Unix(v.Updated, 0)
	return nil
}

// MarshalJSON handles serialization of a ReportingReportType.
// This custom marshaling is needed to handle the time fields correctly.
func (r ReportingReportType) MarshalJSON() ([]byte, error) {
	type reportingReportType ReportingReportType
	v := struct {
		DataAvailableEnd   int64 `json:"data_available_end"`
		DataAvailableStart int64 `json:"data_available_start"`
		Updated            int64 `json:"updated"`
		reportingReportType
	}{
		reportingReportType: (reportingReportType)(r),
		DataAvailableEnd:    r.DataAvailableEnd.Unix(),
		DataAvailableStart:  r.DataAvailableStart.Unix(),
		Updated:             r.Updated.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}
