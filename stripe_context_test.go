package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestParseStripeContext(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "single segment",
			input:    "account1",
			expected: []string{"account1"},
		},
		{
			name:     "multiple segments",
			input:    "account1/subaccount2/user3",
			expected: []string{"account1", "subaccount2", "user3"},
		},
		{
			name:     "segments with special characters",
			input:    "account-1/sub_account.2/user@3",
			expected: []string{"account-1", "sub_account.2", "user@3"},
		},
		{
			name:     "segments with spaces",
			input:    "account 1/sub account 2",
			expected: []string{"account 1", "sub account 2"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := ParseStripeContext(tt.input)
			assert.Equal(t, tt.expected, ctx.segments)
			assert.Equal(t, tt.expected, ctx.Segments()) // Test the public Segments() method
		})
	}
}

func TestStripeContext_String(t *testing.T) {
	tests := []struct {
		name     string
		segments []string
		expected string
	}{
		{
			name:     "empty context",
			segments: []string{},
			expected: "",
		},
		{
			name:     "single segment",
			segments: []string{"account1"},
			expected: "account1",
		},
		{
			name:     "multiple segments",
			segments: []string{"account1", "subaccount2", "user3"},
			expected: "account1/subaccount2/user3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &StripeContext{segments: tt.segments}
			assert.Equal(t, tt.expected, ctx.String())
		})
	}
}

func TestStripeContext_Parent(t *testing.T) {
	tests := []struct {
		name          string
		segments      []string
		expectedSegs  []string
		expectError   bool
		errorContains string
	}{
		{
			name:          "empty context",
			segments:      []string{},
			expectedSegs:  nil,
			expectError:   true,
			errorContains: "cannot get parent of empty StripeContext",
		},
		{
			name:         "single segment",
			segments:     []string{"account1"},
			expectedSegs: []string{},
			expectError:  false,
		},
		{
			name:         "two segments",
			segments:     []string{"account1", "subaccount2"},
			expectedSegs: []string{"account1"},
			expectError:  false,
		},
		{
			name:         "multiple segments",
			segments:     []string{"account1", "subaccount2", "user3"},
			expectedSegs: []string{"account1", "subaccount2"},
			expectError:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &StripeContext{segments: tt.segments}
			parent, err := ctx.Parent()

			if tt.expectError {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errorContains)
				assert.Nil(t, parent)
			} else {
				assert.NoError(t, err)
				assert.NotNil(t, parent)
				assert.Equal(t, tt.expectedSegs, parent.segments)
			}
		})
	}
}

func TestStripeContext_Child(t *testing.T) {
	tests := []struct {
		name         string
		segments     []string
		childSegment string
		expectedSegs []string
	}{
		{
			name:         "empty context",
			segments:     []string{},
			childSegment: "account1",
			expectedSegs: []string{"account1"},
		},
		{
			name:         "single segment",
			segments:     []string{"account1"},
			childSegment: "subaccount2",
			expectedSegs: []string{"account1", "subaccount2"},
		},
		{
			name:         "multiple segments",
			segments:     []string{"account1", "subaccount2"},
			childSegment: "user3",
			expectedSegs: []string{"account1", "subaccount2", "user3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &StripeContext{segments: tt.segments}
			child := ctx.Child(tt.childSegment)

			assert.Equal(t, tt.expectedSegs, child.segments)
			// Ensure original context is unchanged (immutability)
			assert.Equal(t, tt.segments, ctx.segments)
		})
	}
}

func TestStripeContext_IsEmpty(t *testing.T) {
	tests := []struct {
		name     string
		segments []string
		expected bool
	}{
		{
			name:     "empty context",
			segments: []string{},
			expected: true,
		},
		{
			name:     "non-empty context",
			segments: []string{"account1"},
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &StripeContext{segments: tt.segments}
			assert.Equal(t, tt.expected, ctx.IsEmpty())
		})
	}
}

func TestStripeContext_Segments(t *testing.T) {
	original := []string{"account1", "subaccount2", "user3"}
	ctx := &StripeContext{segments: original}

	segments := ctx.Segments()
	assert.Equal(t, original, segments)

	// Modify the returned slice to ensure it doesn't affect the original
	segments[0] = "modified"
	assert.Equal(t, original, ctx.segments) // Original should be unchanged
	assert.NotEqual(t, segments, ctx.segments)
}

func TestStripeContext_ImmutabilityChain(t *testing.T) {
	// Test that chaining operations doesn't modify original contexts
	original := ParseStripeContext("account1/subaccount2/user3")
	originalStr := original.String()

	child := original.Child("extra")
	parent, err := original.Parent()
	assert.NoError(t, err)

	// Original should be unchanged
	assert.Equal(t, originalStr, original.String())
	assert.Equal(t, "account1/subaccount2/user3/extra", child.String())
	assert.Equal(t, "account1/subaccount2", parent.String())

	// Test multiple operations on same original
	child2 := original.Child("another")
	parent2, err := original.Parent()
	assert.NoError(t, err)

	assert.Equal(t, originalStr, original.String())
	assert.Equal(t, "account1/subaccount2/user3/another", child2.String())
	assert.Equal(t, "account1/subaccount2", parent2.String())
}

func TestStripeContext_ParseAndStringRoundTrip(t *testing.T) {
	testCases := []string{
		"",
		"account1",
		"account1/subaccount2",
		"account1/subaccount2/user3",
		"complex-account_1/sub.account@2/user with spaces",
	}

	for _, original := range testCases {
		t.Run("roundtrip_"+original, func(t *testing.T) {
			ctx := ParseStripeContext(original)
			result := ctx.String()
			assert.Equal(t, original, result)
		})
	}
}

func TestStripeContext_JSON(t *testing.T) {
	tests := []struct {
		name     string
		context  *StripeContext
		jsonStr  string
	}{
		{
			name:    "empty context",
			context: ParseStripeContext(""),
			jsonStr: `""`,
		},
		{
			name:    "single segment",
			context: ParseStripeContext("account1"),
			jsonStr: `"account1"`,
		},
		{
			name:    "multiple segments",
			context: ParseStripeContext("account1/subaccount2/user3"),
			jsonStr: `"account1/subaccount2/user3"`,
		},
		{
			name:    "nil context",
			context: nil,
			jsonStr: `null`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_marshal", func(t *testing.T) {
			data, err := json.Marshal(tt.context)
			assert.NoError(t, err)
			assert.Equal(t, tt.jsonStr, string(data))
		})

		if tt.context != nil {
			t.Run(tt.name+"_unmarshal", func(t *testing.T) {
				var ctx StripeContext
				err := json.Unmarshal([]byte(tt.jsonStr), &ctx)
				assert.NoError(t, err)
				assert.Equal(t, tt.context.String(), ctx.String())
				assert.Equal(t, tt.context.segments, ctx.segments)
			})
		}
	}
}

func TestStripeContext_JSONInStruct(t *testing.T) {
	type TestStruct struct {
		Context *StripeContext `json:"context,omitempty"`
		Name    string         `json:"name"`
	}

	tests := []struct {
		name     string
		data     TestStruct
		expected string
	}{
		{
			name: "with context",
			data: TestStruct{
				Context: ParseStripeContext("account1/user2"),
				Name:    "test",
			},
			expected: `{"context":"account1/user2","name":"test"}`,
		},
		{
			name: "without context",
			data: TestStruct{
				Name: "test",
			},
			expected: `{"name":"test"}`,
		},
		{
			name: "with empty context",
			data: TestStruct{
				Context: ParseStripeContext(""),
				Name:    "test",
			},
			expected: `{"context":"","name":"test"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_marshal", func(t *testing.T) {
			data, err := json.Marshal(tt.data)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, string(data))
		})

		t.Run(tt.name+"_unmarshal", func(t *testing.T) {
			var result TestStruct
			err := json.Unmarshal([]byte(tt.expected), &result)
			assert.NoError(t, err)

			if tt.data.Context == nil {
				assert.Nil(t, result.Context)
			} else {
				assert.NotNil(t, result.Context)
				assert.Equal(t, tt.data.Context.String(), result.Context.String())
			}
			assert.Equal(t, tt.data.Name, result.Name)
		})
	}
}

func TestStripeContext_JSONUnmarshalError(t *testing.T) {
	var ctx StripeContext

	// Test invalid JSON
	err := json.Unmarshal([]byte(`123`), &ctx)
	assert.Error(t, err)

	// Test malformed JSON
	err = json.Unmarshal([]byte(`{`), &ctx)
	assert.Error(t, err)
}