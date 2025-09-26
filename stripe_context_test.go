package stripe

import (
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestNewStripeContext(t *testing.T) {
	// Test with nil segments
	{
		ctx := NewStripeContext(nil)
		assert.NotNil(t, ctx)
		assert.Len(t, ctx.Segments, 0)
		assert.Equal(t, "", *ctx.StringPtr())
	}

	// Test with empty segments
	{
		ctx := NewStripeContext([]string{})
		assert.NotNil(t, ctx)
		assert.Len(t, ctx.Segments, 0)
		assert.Equal(t, "", *ctx.StringPtr())
	}

	// Test with segments
	{
		segments := []string{"workspace", "account", "customer"}
		ctx := NewStripeContext(segments)
		assert.NotNil(t, ctx)
		assert.Equal(t, "workspace/account/customer", *ctx.StringPtr())
		assert.Equal(t, segments, ctx.Segments)
	}

	// Test immutability - modifying original slice doesn't affect context
	{
		segments := []string{"a", "b"}
		ctx := NewStripeContext(segments)
		segments[0] = "modified"
		assert.Equal(t, "a", ctx.Segments[0])
	}
}

func TestParseStripeContext(t *testing.T) {
	// Test empty string
	{
		ctx := ParseStripeContext("")
		assert.Nil(t, ctx)
	}

	// Test single segment
	{
		ctx := ParseStripeContext("workspace")
		assert.Equal(t, 1, len(ctx.Segments))
		assert.Equal(t, "workspace", *ctx.StringPtr())
		assert.Equal(t, []string{"workspace"}, ctx.Segments)
	}

	// Test multiple segments
	{
		ctx := ParseStripeContext("workspace/account/customer")
		assert.Equal(t, 3, len(ctx.Segments))
		assert.Equal(t, "workspace/account/customer", *ctx.StringPtr())
		assert.Equal(t, []string{"workspace", "account", "customer"}, ctx.Segments)
	}

	// Test empty segments in string
	{
		ctx := ParseStripeContext("a//b")
		assert.Equal(t, 3, len(ctx.Segments))
		assert.Equal(t, []string{"a", "", "b"}, ctx.Segments)
	}
}

func TestStripeContext_Push(t *testing.T) {
	// Test push to empty context
	{
		ctx := NewStripeContext(nil)
		newCtx, err := ctx.Push("segment")
		assert.NoError(t, err)
		assert.NotNil(t, newCtx)
		assert.Equal(t, 1, len(newCtx.Segments))
		assert.Equal(t, "segment", *newCtx.StringPtr())

		// Original context unchanged
		assert.Equal(t, 0, len(ctx.Segments))
		assert.NotEqual(t, ctx, newCtx)
	}

	// Test push to existing context
	{
		ctx := NewStripeContext([]string{"a", "b"})
		newCtx, err := ctx.Push("c")
		assert.NoError(t, err)
		assert.Len(t, newCtx.Segments, 3)
		assert.Equal(t, "a/b/c", *newCtx.StringPtr())

		// Original context unchanged
		assert.Equal(t, 2, len(ctx.Segments))
		assert.Equal(t, "a/b", *ctx.StringPtr())
	}

	// Test push with whitespace
	{
		ctx := NewStripeContext(nil)
		newCtx, err := ctx.Push("  segment  ")
		assert.NoError(t, err)
		assert.Equal(t, "segment", *newCtx.StringPtr())
	}

	// Test push empty segment
	{
		ctx := NewStripeContext(nil)
		_, err := ctx.Push("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cannot be empty")
	}

	// Test push whitespace-only segment
	{
		ctx := NewStripeContext(nil)
		_, err := ctx.Push("   ")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cannot be empty")
	}
}

func TestStripeContext_Pop(t *testing.T) {
	// Test pop from empty context
	{
		ctx := NewStripeContext(nil)
		newCtx, err := ctx.Pop()
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "cannot pop from empty context")
		assert.Nil(t, newCtx)

		// Different instances (check pointers)
		assert.True(t, ctx != newCtx)
	}

	// Test pop single segment
	{
		ctx := NewStripeContext([]string{"single"})
		newCtx, err := ctx.Pop()
		assert.NoError(t, err)
		assert.Equal(t, 0, len(newCtx.Segments))

		// Original unchanged
		assert.Equal(t, 1, len(ctx.Segments))
		assert.Equal(t, "single", *ctx.StringPtr())
	}

	// Test pop multiple segments
	{
		ctx := NewStripeContext([]string{"a", "b", "c"})
		newCtx, err := ctx.Pop()
		assert.NoError(t, err)
		assert.Equal(t, 2, len(newCtx.Segments))
		assert.Equal(t, "a/b", *newCtx.StringPtr())

		// Original unchanged
		assert.Equal(t, 3, len(ctx.Segments))
		assert.Equal(t, "a/b/c", *ctx.StringPtr())
	}
}

func TestStripeContext_String(t *testing.T) {
	// Test empty context
	{
		ctx := NewStripeContext(nil)
		assert.Equal(t, "", *ctx.StringPtr())
	}

	// Test single segment
	{
		ctx := NewStripeContext([]string{"workspace"})
		assert.Equal(t, "workspace", *ctx.StringPtr())
	}

	// Test multiple segments
	{
		ctx := NewStripeContext([]string{"a", "b", "c"})
		assert.Equal(t, "a/b/c", *ctx.StringPtr())
	}
}

func TestStripeContext_Immutability(t *testing.T) {
	original := NewStripeContext([]string{"a", "b"})
	pushed, _ := original.Push("c")
	popped, err := original.Pop()
	assert.NoError(t, err)

	// Original remains unchanged
	assert.Len(t, original.Segments, 2)
	assert.Equal(t, "a/b", *original.StringPtr())

	// New instances created
	assert.Len(t, pushed.Segments, 3)
	assert.Equal(t, "a/b/c", *pushed.StringPtr())
	assert.Len(t, popped.Segments, 1)
	assert.Equal(t, "a", *popped.StringPtr())

	// All different instances (check pointers)
	assert.True(t, original != pushed)
	assert.True(t, original != popped)
	assert.True(t, pushed != popped)
}

func TestStripeContext_UsagePattern(t *testing.T) {
	// Common usage: hierarchical context building
	baseContext := ParseStripeContext("workspace_123")
	child, err := baseContext.Push("account_456")
	assert.NoError(t, err)
	grandchild, err := child.Push("customer_789")
	assert.NoError(t, err)

	assert.Equal(t, "workspace_123", *baseContext.StringPtr())
	assert.Equal(t, "workspace_123/account_456", *child.StringPtr())
	assert.Equal(t, "workspace_123/account_456/customer_789", *grandchild.StringPtr())

	// Go back up the hierarchy
	backToChild, err := grandchild.Pop()
	assert.NoError(t, err)
	backToBase, err := backToChild.Pop()
	assert.NoError(t, err)

	assert.Equal(t, "workspace_123/account_456", *backToChild.StringPtr())
	assert.Equal(t, "workspace_123", *backToBase.StringPtr())
}

func TestParams_SetStripeContextObject(t *testing.T) {
	params := &Params{}

	// Test with nil context
	{
		params.SetStripeContextFrom(nil)
		assert.Nil(t, params.StripeContext)
	}

	// Test with valid context
	{
		ctx := NewStripeContext([]string{"workspace", "account"})
		params.SetStripeContextFrom(ctx)
		assert.NotNil(t, params.StripeContext)
		assert.Equal(t, "workspace/account", *params.StripeContext)
	}

	// Test with empty context
	{
		ctx := NewStripeContext(nil)
		params.SetStripeContextFrom(ctx)
		assert.NotNil(t, params.StripeContext)
		assert.Equal(t, "", *params.StripeContext)
	}
}
