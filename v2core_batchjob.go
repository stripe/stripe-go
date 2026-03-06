//
//
// File generated from our OpenAPI spec
//
//

package stripe

type V2CoreBatchJob struct {
	APIResource
	// Unique identifier for the batch job.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The URL to upload the JSONL file to.
	URL string `json:"url"`
}
