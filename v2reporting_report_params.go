//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves metadata about a specific `Report` template, including its name, description,
// and the parameters it accepts. It's useful for understanding the capabilities and
// requirements of a particular `Report` before requesting a `ReportRun`.
type V2ReportingReportParams struct {
	Params `form:"*"`
}

// Retrieves metadata about a specific `Report` template, including its name, description,
// and the parameters it accepts. It's useful for understanding the capabilities and
// requirements of a particular `Report` before requesting a `ReportRun`.
type V2ReportingReportRetrieveParams struct {
	Params `form:"*"`
}
