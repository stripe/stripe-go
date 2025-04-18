//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Creates a meter event session to send usage on the high-throughput meter event stream. Authentication tokens are only valid for 15 minutes, so you will need to create a new meter event session when your token expires.
type V2BillingMeterEventSessionParams struct {
	Params `form:"*"`
}

// Creates a meter event session to send usage on the high-throughput meter event stream. Authentication tokens are only valid for 15 minutes, so you will need to create a new meter event session when your token expires.
type V2BillingMeterEventSessionCreateParams struct {
	Params `form:"*"`
}
