package stripe

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Context provides a container and convenience methods for interacting with the `Stripe-Context` header. All methods return a new instance of Context.
// You can use its `StringPtr` method whenever you're initializing a `StripeClient` or sending `StripeContext` with a request. It's also found in the `EventNotification.Context` property.
type Context struct {
	Segments []string
}

// NewStripeContext creates a new StripeContext with the given segments.
// If segments is nil or empty, creates an empty context.
func NewStripeContext(segments []string) *Context {
	if len(segments) == 0 {
		return &Context{
			Segments: []string{},
		}
	}

	// Create a copy to ensure immutability
	segmentsCopy := make([]string, len(segments))
	copy(segmentsCopy, segments)

	return &Context{
		Segments: segmentsCopy,
	}
}

// ParseStripeContext parses a context string into a StripeContext instance.
// If contextStr is empty, returns an empty StripeContext.
func ParseStripeContext(contextStr string) *Context {
	if contextStr == "" {
		return nil
	}

	return NewStripeContext(strings.Split(contextStr, "/"))
}

// Push creates a new StripeContext with an additional segment appended.
// Returns an error if the segment is empty.
func (c *Context) Push(segment string) (*Context, error) {
	newSegment := strings.TrimSpace(segment)
	if newSegment == "" {
		return nil, fmt.Errorf("segment cannot be empty or whitespace")
	}

	newSegments := make([]string, len(c.Segments)+1)
	copy(newSegments, c.Segments)
	newSegments[len(c.Segments)] = newSegment

	return NewStripeContext(newSegments), nil
}

// Pop creates a new StripeContext with the last segment removed.
// If there are no segments, returns a new empty StripeContext.
func (c *Context) Pop() (*Context, error) {
	if len(c.Segments) == 0 {
		return nil, fmt.Errorf("cannot pop from empty context")
	}

	return NewStripeContext(c.Segments[:len(c.Segments)-1]), nil
}

// StringPtr returns the string representation of the StripeContext.
// Segments are joined with "/" as the separator.
func (c *Context) StringPtr() *string {
	result := strings.Join(c.Segments, "/")
	return &result
}

// UnmarshalJSON implements the [encoding/json.Unmarshaler] interface for StripeContext.
func (c *Context) UnmarshalJSON(data []byte) error {
	var contextStr string
	if err := json.Unmarshal(data, &contextStr); err != nil {
		return err
	}

	if contextStr != "" {
		*c = *ParseStripeContext(contextStr)
	}
	return nil
}
