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
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The URL to upload the JSONL file to.
	URL string `json:"url"`
}
