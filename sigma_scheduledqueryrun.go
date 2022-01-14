//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

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
}

// Retrieves the details of an scheduled query run.
type SigmaScheduledQueryRunParams struct {
	Params `form:"*"`
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
	Created int64 `json:"created"`
	// When the query was run, Sigma contained a snapshot of your Stripe data at this time.
	DataLoadTime int64                        `json:"data_load_time"`
	Error        *SigmaScheduledQueryRunError `json:"error"`
	// The file object representing the results of the query.
	File *File `json:"file"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	Query  string `json:"query"`
	// Time at which the result expires and is no longer available for download.
	ResultAvailableUntil int64 `json:"result_available_until"`
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
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (s *SigmaScheduledQueryRun) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		s.ID = id
		return nil
	}

	type sigmaScheduledQueryRun SigmaScheduledQueryRun
	var v sigmaScheduledQueryRun
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*s = SigmaScheduledQueryRun(v)
	return nil
}
