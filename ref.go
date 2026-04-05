package stripe

import (
	"context"
	"errors"
	"net/http"
)

// Ref is a typed reference to a Stripe API object. It carries the object's
// type, ID, and URL, and can fetch the full object via Fetch.
//
// Ref values are typically embedded in v2 event payloads. They are inert
// (Fetch returns an error) when deserialized without an attached backend.
type Ref[T any] struct {
	// ID is the unique identifier of the referenced object.
	ID string `json:"id"`
	// Type is the object type of the referenced object (e.g. "v2.billing.meter").
	Type string `json:"type"`
	// URL is the API URL of the referenced object.
	URL string `json:"url"`

	// backend and key are set by the SDK after deserialization via SetBackend.
	backend Backend
	key     string
}

// SetBackend attaches a Backend and API key to this Ref so that Fetch can
// make API calls. Generated service code calls this after unmarshaling event
// payloads.
func (r *Ref[T]) SetBackend(b Backend, key string) {
	r.backend = b
	r.key = key
}

// Fetch retrieves the full object that this Ref points to.
//
// It returns an error if the Ref was not constructed with a backend (e.g. it
// was deserialized directly from JSON without going through the SDK client).
func (r *Ref[T]) Fetch(ctx context.Context) (*T, error) {
	if r.backend == nil {
		return nil, errors.New("stripe: Ref.Fetch called on a Ref without an attached backend; use stripe.Client to obtain Ref values")
	}

	obj := new(T)
	params := &Params{Context: ctx}
	err := r.backend.Call(http.MethodGet, r.URL, r.key, params, any(obj).(LastResponseSetter))
	if err != nil {
		return nil, err
	}
	return obj, nil
}
