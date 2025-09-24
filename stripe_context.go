package stripe

import (
	"encoding/json"
	"errors"
	"strings"
)

// StripeContext represents a context identifier that can be used to segment API requests
// and events. It consists of segments separated by forward slashes, similar to a file path.
// StripeContext is externally immutable - all methods return new instances rather than
// modifying the receiver.
type StripeContext struct {
	segments []string
}

// Parse creates a new StripeContext from a string representation.
// The string should be in the format "segment1/segment2/segment3".
// Empty strings are allowed and result in an empty context.
func ParseStripeContext(s string) *StripeContext {
	if s == "" {
		return &StripeContext{segments: []string{}}
	}

	segments := strings.Split(s, "/")
	return &StripeContext{segments: segments}
}

// Parent returns a new StripeContext with the last segment removed.
// Returns an error if the context is empty (has no segments).
func (sc *StripeContext) Parent() (*StripeContext, error) {
	if len(sc.segments) == 0 {
		return nil, errors.New("cannot get parent of empty StripeContext")
	}

	if len(sc.segments) == 1 {
		return &StripeContext{segments: []string{}}, nil
	}

	newSegments := make([]string, len(sc.segments)-1)
	copy(newSegments, sc.segments[:len(sc.segments)-1])
	return &StripeContext{segments: newSegments}, nil
}

// Child returns a new StripeContext with the given segment appended.
func (sc *StripeContext) Child(segment string) *StripeContext {
	newSegments := make([]string, len(sc.segments)+1)
	copy(newSegments, sc.segments)
	newSegments[len(newSegments)-1] = segment
	return &StripeContext{segments: newSegments}
}

// String returns the string representation of the StripeContext.
// Segments are joined with forward slashes.
func (sc *StripeContext) String() string {
	if len(sc.segments) == 0 {
		return ""
	}
	return strings.Join(sc.segments, "/")
}

// IsEmpty returns true if the context has no segments.
func (sc *StripeContext) IsEmpty() bool {
	return len(sc.segments) == 0
}

// Segments returns a copy of the segments slice to prevent external modification.
func (sc *StripeContext) Segments() []string {
	if len(sc.segments) == 0 {
		return []string{}
	}
	result := make([]string, len(sc.segments))
	copy(result, sc.segments)
	return result
}

// MarshalJSON implements the json.Marshaler interface.
// It marshals the StripeContext as a JSON string.
func (sc *StripeContext) MarshalJSON() ([]byte, error) {
	return json.Marshal(sc.String())
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// It unmarshals a JSON string into a StripeContext.
func (sc *StripeContext) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	parsed := ParseStripeContext(s)
	sc.segments = parsed.segments
	return nil
}