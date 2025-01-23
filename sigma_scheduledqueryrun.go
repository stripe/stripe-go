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

// The query's execution status, which will be `completed` for successful runs, and `canceled`, `failed`, or `timed_out` otherwise.
type SigmaScheduledQueryRunStatus string

// List of values that SigmaScheduledQueryRunStatus can take
const (
	SigmaScheduledQueryRunStatusCanceled  SigmaScheduledQueryRunStatus = "canceled"
	SigmaScheduledQueryRunStatusCompleted SigmaScheduledQueryRunStatus = "completed"
	SigmaScheduledQueryRunStatusFailed    SigmaScheduledQueryRunStatus = "failed"
	SigmaScheduledQueryRunStatusTimedOut  SigmaScheduledQueryRunStatus = "timed_out"
)

// Returns a list of scheduled query runs.
type SigmaScheduledQueryRunListParams struct {
	ListParams `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *SigmaScheduledQueryRunListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the details of an scheduled query run.
type SigmaScheduledQueryRunParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *SigmaScheduledQueryRunParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

type SigmaScheduledQueryRunError struct {
	// Information about the run failure.
	Message string `json:"message"`
}

// If you have [scheduled a Sigma query](https://stripe.com/docs/sigma/scheduled-queries), you'll
// receive a `sigma.scheduled_query_run.created` webhook each time the query
// runs. The webhook contains a `ScheduledQueryRun` object, which you can use to
// retrieve the query results.
type SigmaScheduledQueryRun struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created time.Time `json:"created"`
	// When the query was run, Sigma contained a snapshot of your Stripe data at this time.
	DataLoadTime time.Time                    `json:"data_load_time"`
	Error        *SigmaScheduledQueryRunError `json:"error"`
	// The file object representing the results of the query.
	File *File `json:"file"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Time at which the result expires and is no longer available for download.
	ResultAvailableUntil time.Time `json:"result_available_until"`
	// SQL for the query.
	SQL string `json:"sql"`
	// The query's execution status, which will be `completed` for successful runs, and `canceled`, `failed`, or `timed_out` otherwise.
	Status SigmaScheduledQueryRunStatus `json:"status"`
	// Title of the query.
	Title string `json:"title"`
}

// SigmaScheduledQueryRunList is a list of ScheduledQueryRuns as retrieved from a list endpoint.
type SigmaScheduledQueryRunList struct {
	APIResource
	ListMeta
	Data []*SigmaScheduledQueryRun `json:"data"`
}

// UnmarshalJSON handles deserialization of a SigmaScheduledQueryRun.
// This custom unmarshaling is needed to handle the time fields correctly.
func (s *SigmaScheduledQueryRun) UnmarshalJSON(data []byte) error {
	type sigmaScheduledQueryRun SigmaScheduledQueryRun
	v := struct {
		Created              int64 `json:"created"`
		DataLoadTime         int64 `json:"data_load_time"`
		ResultAvailableUntil int64 `json:"result_available_until"`
		*sigmaScheduledQueryRun
	}{
		sigmaScheduledQueryRun: (*sigmaScheduledQueryRun)(s),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	s.Created = time.Unix(v.Created, 0)
	s.DataLoadTime = time.Unix(v.DataLoadTime, 0)
	s.ResultAvailableUntil = time.Unix(v.ResultAvailableUntil, 0)
	return nil
}

// MarshalJSON handles serialization of a SigmaScheduledQueryRun.
// This custom marshaling is needed to handle the time fields correctly.
func (s SigmaScheduledQueryRun) MarshalJSON() ([]byte, error) {
	type sigmaScheduledQueryRun SigmaScheduledQueryRun
	v := struct {
		Created              int64 `json:"created"`
		DataLoadTime         int64 `json:"data_load_time"`
		ResultAvailableUntil int64 `json:"result_available_until"`
		sigmaScheduledQueryRun
	}{
		sigmaScheduledQueryRun: (sigmaScheduledQueryRun)(s),
		Created:                s.Created.Unix(),
		DataLoadTime:           s.DataLoadTime.Unix(),
		ResultAvailableUntil:   s.ResultAvailableUntil.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}
