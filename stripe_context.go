package stripe

import (
	"encoding/json"
	"fmt"
	"strings"
)

// StripeContext provides an immutable container and convenience methods for interacting with the `Stripe-Context` header. All methods return a new instance of StripeContext.
// You can use it whenever you're initializing a `StripeClient` or sending `stripe_context` with a request. It's also found in the `EventNotification.context` property.
type StripeContext struct {
	segments []string
}

// NewStripeContext creates a new StripeContext with the given segments.
// If segments is nil or empty, creates an empty context.
func NewStripeContext(segments []string) *StripeContext {
	if len(segments) == 0 {
		return &StripeContext{
			segments: []string{},
		}
	}

	// Create a copy to ensure immutability
	segmentsCopy := make([]string, len(segments))
	copy(segmentsCopy, segments)

	return &StripeContext{
		segments: segmentsCopy,
	}
}

// ParseStripeContext parses a context string into a StripeContext instance.
// If contextStr is empty, returns an empty StripeContext.
func ParseStripeContext(contextStr string) *StripeContext {
	if contextStr == "" {
		return nil
	}

	return NewStripeContext(strings.Split(contextStr, "/"))
}

// Push creates a new StripeContext with an additional segment appended.
// Returns an error if the segment is empty.
func (c *StripeContext) Push(segment string) (*StripeContext, error) {
	newSegment := strings.TrimSpace(segment)
	if newSegment == "" {
		return nil, fmt.Errorf("segment cannot be empty or whitespace")
	}

	newSegments := make([]string, len(c.segments)+1)
	copy(newSegments, c.segments)
	newSegments[len(c.segments)] = newSegment

	return NewStripeContext(newSegments), nil
}

// Pop creates a new StripeContext with the last segment removed.
// If there are no segments, returns a new empty StripeContext.
func (c *StripeContext) Pop() (*StripeContext, error) {
	if len(c.segments) == 0 {
		return nil, fmt.Errorf("cannot pop from empty context")
	}

	return NewStripeContext(c.segments[:len(c.segments)-1]), nil
}

// String returns the string representation of the StripeContext.
// Segments are joined with "/" as the separator.
func (c *StripeContext) String() *string {
	result := strings.Join(c.segments, "/")
	return &result
}

// Segments returns a copy of the segments slice.
// The returned slice is safe to modify without affecting the original StripeContext.
func (c *StripeContext) Segments() []string {
	segments := make([]string, len(c.segments))
	copy(segments, c.segments)
	return segments
}

// IsEmpty returns true if the StripeContext has no segments.
func (c *StripeContext) IsEmpty() bool {
	return len(c.segments) == 0
}

// Size returns the number of segments in the StripeContext.
func (c *StripeContext) Size() int {
	return len(c.segments)
}

// Clone creates a deep copy of the StripeContext.
func (c *StripeContext) Clone() *StripeContext {
	return NewStripeContext(c.segments)
}

// UnmarshalJSON implements the [encoding/json.Unmarshaler] interface for StripeContext.
func (c *StripeContext) UnmarshalJSON(data []byte) error {
	var contextStr string
	if err := json.Unmarshal(data, &contextStr); err != nil {
		return err
	}

	if contextStr != "" {
		*c = *ParseStripeContext(contextStr)
	}
	return nil
}
