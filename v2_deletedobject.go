//
//
// File generated from our OpenAPI spec
//
//

package stripe

type V2DeletedObject struct {
	APIResource
	// The ID of the object that's being deleted.
	ID string `json:"id"`
	// String representing the type of the object that has been deleted. Objects of the same type share the same value of the object field.
	Object string `json:"object,omitempty"`
}
