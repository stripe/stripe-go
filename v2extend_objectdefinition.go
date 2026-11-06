//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Represents a user-defined object definition that models a custom object.
// Object definitions enable the object records API to create and manage records of that object.
type V2ExtendObjectDefinition struct {
	APIResource
	// The plural API name used in this definition's object record API paths. Immutable across versions.
	APINamePlural string `json:"api_name_plural"`
	// The singular API name used in the object tag of this definition's records. Immutable across versions.
	APINameSingular string `json:"api_name_singular"`
	// Time at which the object definition was created.
	Created time.Time `json:"created"`
	// Unique identifier for the object definition.
	ID string `json:"id"`
	// The ID of the latest object definition version.
	LatestVersion string `json:"latest_version"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
}
