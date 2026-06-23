package stripe

// Ref is a reference to a Stripe object. It holds the object's id and type.
type Ref struct {
	// ID is the unique identifier of the referenced object.
	ID string `json:"id"`
	// Type is the object type of the referenced object (e.g. "v2.core.account").
	Type string `json:"type"`
}
