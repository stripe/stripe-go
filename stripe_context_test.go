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
		assert.True(t, ctx.IsEmpty())
		assert.Equal(t, 0, ctx.Size())
		assert.Equal(t, "", *ctx.String())
	}

	// Test with empty segments
	{
		ctx := NewStripeContext([]string{})
		assert.NotNil(t, ctx)
		assert.True(t, ctx.IsEmpty())
		assert.Equal(t, 0, ctx.Size())
	}

	// Test with segments
	{
		segments := []string{"workspace", "account", "customer"}
		ctx := NewStripeContext(segments)
		assert.NotNil(t, ctx)
		assert.False(t, ctx.IsEmpty())
		assert.Equal(t, 3, ctx.Size())
		assert.Equal(t, "workspace/account/customer", *ctx.String())
		assert.Equal(t, segments, ctx.Segments())
	}

	// Test immutability - modifying original slice doesn't affect context
	{
		segments := []string{"a", "b"}
		ctx := NewStripeContext(segments)
		segments[0] = "modified"
		assert.Equal(t, "a", ctx.Segments()[0])
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
		assert.Equal(t, 1, ctx.Size())
		assert.Equal(t, "workspace", *ctx.String())
		assert.Equal(t, []string{"workspace"}, ctx.Segments())
	}

	// Test multiple segments
	{
		ctx := ParseStripeContext("workspace/account/customer")
		assert.Equal(t, 3, ctx.Size())
		assert.Equal(t, "workspace/account/customer", *ctx.String())
		assert.Equal(t, []string{"workspace", "account", "customer"}, ctx.Segments())
	}

	// Test empty segments in string
	{
		ctx := ParseStripeContext("a//b")
		assert.Equal(t, 3, ctx.Size())
		assert.Equal(t, []string{"a", "", "b"}, ctx.Segments())
	}
}

func TestStripeContext_Push(t *testing.T) {
	// Test push to empty context
	{
		ctx := NewStripeContext(nil)
		newCtx, err := ctx.Push("segment")
		assert.NoError(t, err)
		assert.NotNil(t, newCtx)
		assert.Equal(t, 1, newCtx.Size())
		assert.Equal(t, "segment", *newCtx.String())

		// Original context unchanged
		assert.True(t, ctx.IsEmpty())
		assert.NotEqual(t, ctx, newCtx)
	}

	// Test push to existing context
	{
		ctx := NewStripeContext([]string{"a", "b"})
		newCtx, err := ctx.Push("c")
		assert.NoError(t, err)
		assert.Equal(t, 3, newCtx.Size())
		assert.Equal(t, "a/b/c", *newCtx.String())

		// Original context unchanged
		assert.Equal(t, 2, ctx.Size())
		assert.Equal(t, "a/b", *ctx.String())
	}

	// Test push with whitespace
	{
		ctx := NewStripeContext(nil)
		newCtx, err := ctx.Push("  segment  ")
		assert.NoError(t, err)
		assert.Equal(t, "segment", *newCtx.String())
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
		assert.True(t, newCtx.IsEmpty())

		// Original unchanged
		assert.Equal(t, 1, ctx.Size())
		assert.Equal(t, "single", *ctx.String())
	}

	// Test pop multiple segments
	{
		ctx := NewStripeContext([]string{"a", "b", "c"})
		newCtx, err := ctx.Pop()
		assert.NoError(t, err)
		assert.Equal(t, 2, newCtx.Size())
		assert.Equal(t, "a/b", *newCtx.String())

		// Original unchanged
		assert.Equal(t, 3, ctx.Size())
		assert.Equal(t, "a/b/c", *ctx.String())
	}
}

func TestStripeContext_String(t *testing.T) {
	// Test empty context
	{
		ctx := NewStripeContext(nil)
		assert.Equal(t, "", *ctx.String())
	}

	// Test single segment
	{
		ctx := NewStripeContext([]string{"workspace"})
		assert.Equal(t, "workspace", *ctx.String())
	}

	// Test multiple segments
	{
		ctx := NewStripeContext([]string{"a", "b", "c"})
		assert.Equal(t, "a/b/c", *ctx.String())
	}
}

func TestStripeContext_Segments(t *testing.T) {
	segments := []string{"a", "b", "c"}
	ctx := NewStripeContext(segments)

	// Returns copy
	returned := ctx.Segments()
	assert.Equal(t, segments, returned)

	// Modifying returned slice doesn't affect original
	returned[0] = "modified"
	assert.Equal(t, "a", ctx.Segments()[0])
}

func TestStripeContext_Clone(t *testing.T) {
	ctx := NewStripeContext([]string{"a", "b", "c"})
	clone := ctx.Clone()

	assert.True(t, ctx != clone) // Different instances (check pointers)

	// Modifying clone doesn't affect original
	newClone, _ := clone.Push("d")
	assert.Equal(t, 3, ctx.Size())
	assert.Equal(t, 4, newClone.Size())
}

func TestStripeContext_Immutability(t *testing.T) {
	original := NewStripeContext([]string{"a", "b"})
	pushed, _ := original.Push("c")
	popped, err := original.Pop()
	assert.NoError(t, err)

	// Original remains unchanged
	assert.Equal(t, 2, original.Size())
	assert.Equal(t, "a/b", *original.String())

	// New instances created
	assert.Equal(t, 3, pushed.Size())
	assert.Equal(t, "a/b/c", *pushed.String())
	assert.Equal(t, 1, popped.Size())
	assert.Equal(t, "a", *popped.String())

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

	assert.Equal(t, "workspace_123", *baseContext.String())
	assert.Equal(t, "workspace_123/account_456", *child.String())
	assert.Equal(t, "workspace_123/account_456/customer_789", *grandchild.String())

	// Go back up the hierarchy
	backToChild, err := grandchild.Pop()
	assert.NoError(t, err)
	backToBase, err := backToChild.Pop()
	assert.NoError(t, err)

	assert.Equal(t, "workspace_123/account_456", *backToChild.String())
	assert.Equal(t, "workspace_123", *backToBase.String())
}

func TestParams_SetStripeContextObject(t *testing.T) {
	params := &Params{}

	// Test with nil context
	{
		params.SetStripeContextObject(nil)
		assert.Nil(t, params.StripeContext)
	}

	// Test with valid context
	{
		ctx := NewStripeContext([]string{"workspace", "account"})
		params.SetStripeContextObject(ctx)
		assert.NotNil(t, params.StripeContext)
		assert.Equal(t, "workspace/account", *params.StripeContext)
	}

	// Test with empty context
	{
		ctx := NewStripeContext(nil)
		params.SetStripeContextObject(ctx)
		assert.NotNil(t, params.StripeContext)
		assert.Equal(t, "", *params.StripeContext)
	}
}
