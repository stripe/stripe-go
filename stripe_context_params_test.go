package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
)

func TestParams_StripeContextIntegration(t *testing.T) {
	tests := []struct {
		name            string
		setupParams     func(*Params)
		expectedContext string
	}{
		{
			name: "string context only",
			setupParams: func(p *Params) {
				p.SetStripeContext("account1/user2")
			},
			expectedContext: "account1/user2",
		},
		{
			name: "StripeContext only",
			setupParams: func(p *Params) {
				ctx := ParseStripeContext("account1/subaccount2/user3")
				p.SetStripeContextValue(ctx)
			},
			expectedContext: "account1/subaccount2/user3",
		},
		{
			name: "both set - StripeContextValue takes precedence",
			setupParams: func(p *Params) {
				p.SetStripeContext("old_context")
				ctx := ParseStripeContext("new_context/priority")
				p.SetStripeContextValue(ctx)
			},
			expectedContext: "new_context/priority",
		},
		{
			name: "empty StripeContext takes precedence over string",
			setupParams: func(p *Params) {
				p.SetStripeContext("fallback_context")
				ctx := ParseStripeContext("")
				p.SetStripeContextValue(ctx)
			},
			expectedContext: "",
		},
		{
			name: "neither set",
			setupParams: func(p *Params) {
				// Don't set anything
			},
			expectedContext: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			params := &Params{}
			tt.setupParams(params)

			result := params.GetEffectiveStripeContext()
			assert.Equal(t, tt.expectedContext, result)
		})
	}
}

func TestParams_StripeContextValueImmutability(t *testing.T) {
	// Test that modifying a StripeContext after setting it doesn't affect the Params
	originalCtx := ParseStripeContext("account1/user2")
	params := &Params{}
	params.SetStripeContextValue(originalCtx)

	// Modify the original context (this should create a new instance)
	childCtx := originalCtx.Child("extra")

	// The params should still have the original context
	assert.Equal(t, "account1/user2", params.GetEffectiveStripeContext())
	assert.Equal(t, "account1/user2/extra", childCtx.String())
}

func TestThinEvent_JSONWithStripeContext(t *testing.T) {
	tests := []struct {
		name     string
		event    ThinEvent
		expected string
	}{
		{
			name: "event with context",
			event: ThinEvent{
				ID:      "evt_123",
				Type:    "test.event",
				Context: ParseStripeContext("account1/user2"),
			},
			expected: `{"id":"evt_123","object":"","type":"test.event","livemode":false,"created":"0001-01-01T00:00:00Z","related_object":null,"context":"account1/user2","reason":null}`,
		},
		{
			name: "event without context",
			event: ThinEvent{
				ID:   "evt_456",
				Type: "test.event",
			},
			expected: `{"id":"evt_456","object":"","type":"test.event","livemode":false,"created":"0001-01-01T00:00:00Z","related_object":null,"context":null,"reason":null}`,
		},
		{
			name: "event with empty context",
			event: ThinEvent{
				ID:      "evt_789",
				Type:    "test.event",
				Context: ParseStripeContext(""),
			},
			expected: `{"id":"evt_789","object":"","type":"test.event","livemode":false,"created":"0001-01-01T00:00:00Z","related_object":null,"context":"","reason":null}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_marshal", func(t *testing.T) {
			data, err := json.Marshal(tt.event)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, string(data))
		})

		t.Run(tt.name+"_unmarshal", func(t *testing.T) {
			var event ThinEvent
			err := json.Unmarshal([]byte(tt.expected), &event)
			assert.NoError(t, err)

			assert.Equal(t, tt.event.ID, event.ID)
			assert.Equal(t, tt.event.Type, event.Type)

			if tt.event.Context == nil {
				assert.Nil(t, event.Context)
			} else {
				assert.NotNil(t, event.Context)
				assert.Equal(t, tt.event.Context.String(), event.Context.String())
			}
		})
	}
}

func TestV2BaseEvent_JSONWithStripeContext(t *testing.T) {
	tests := []struct {
		name     string
		event    V2BaseEvent
		expected string
	}{
		{
			name: "event with context",
			event: V2BaseEvent{
				ID:      "evt_123",
				Type:    "test.event",
				Context: ParseStripeContext("account1/user2"),
			},
			expected: `{"context":"account1/user2","created":"0001-01-01T00:00:00Z","id":"evt_123","livemode":false,"object":"","reason":null,"type":"test.event"}`,
		},
		{
			name: "event without context",
			event: V2BaseEvent{
				ID:   "evt_456",
				Type: "test.event",
			},
			expected: `{"created":"0001-01-01T00:00:00Z","id":"evt_456","livemode":false,"object":"","reason":null,"type":"test.event"}`,
		},
		{
			name: "event with empty context",
			event: V2BaseEvent{
				ID:      "evt_789",
				Type:    "test.event",
				Context: ParseStripeContext(""),
			},
			expected: `{"context":"","created":"0001-01-01T00:00:00Z","id":"evt_789","livemode":false,"object":"","reason":null,"type":"test.event"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_marshal", func(t *testing.T) {
			data, err := json.Marshal(tt.event)
			assert.NoError(t, err)
			assert.Equal(t, tt.expected, string(data))
		})

		t.Run(tt.name+"_unmarshal", func(t *testing.T) {
			var event V2BaseEvent
			err := json.Unmarshal([]byte(tt.expected), &event)
			assert.NoError(t, err)

			assert.Equal(t, tt.event.ID, event.ID)
			assert.Equal(t, tt.event.Type, event.Type)

			if tt.event.Context == nil {
				assert.Nil(t, event.Context)
			} else {
				assert.NotNil(t, event.Context)
				assert.Equal(t, tt.event.Context.String(), event.Context.String())
			}
		})
	}
}

func TestStripeContext_UsageExamples(t *testing.T) {
	// Example usage patterns that would be common in real code

	t.Run("hierarchical account structure", func(t *testing.T) {
		// Start with a main account
		mainAccount := ParseStripeContext("acct_main")
		assert.Equal(t, "acct_main", mainAccount.String())

		// Add a sub-account
		subAccount := mainAccount.Child("acct_sub1")
		assert.Equal(t, "acct_main/acct_sub1", subAccount.String())

		// Add a user to the sub-account
		user := subAccount.Child("user_123")
		assert.Equal(t, "acct_main/acct_sub1/user_123", user.String())

		// Get back to the sub-account
		backToSub, err := user.Parent()
		assert.NoError(t, err)
		assert.Equal(t, "acct_main/acct_sub1", backToSub.String())

		// Get back to the main account
		backToMain, err := backToSub.Parent()
		assert.NoError(t, err)
		assert.Equal(t, "acct_main", backToMain.String())
	})

	t.Run("context-aware API calls", func(t *testing.T) {
		// Simulate creating parameters for an API call
		params := &Params{}

		// Set context using the new StripeContext
		accountContext := ParseStripeContext("organization1/team2/project3")
		params.SetStripeContextValue(accountContext)

		// Verify the effective context
		assert.Equal(t, "organization1/team2/project3", params.GetEffectiveStripeContext())

		// For a child operation, create a child context
		childContext := accountContext.Child("operation4")
		childParams := &Params{}
		childParams.SetStripeContextValue(childContext)

		assert.Equal(t, "organization1/team2/project3/operation4", childParams.GetEffectiveStripeContext())
	})

	t.Run("parsing webhook context", func(t *testing.T) {
		// Simulate receiving a webhook with context
		webhookJSON := `{
			"id": "evt_webhook_123",
			"type": "payment_intent.succeeded",
			"context": "merchant1/location2/terminal3"
		}`

		var event ThinEvent
		err := json.Unmarshal([]byte(webhookJSON), &event)
		assert.NoError(t, err)

		assert.NotNil(t, event.Context)
		assert.Equal(t, "merchant1/location2/terminal3", event.Context.String())

		// Extract individual segments
		segments := event.Context.Segments()
		assert.Equal(t, []string{"merchant1", "location2", "terminal3"}, segments)

		// Get parent contexts
		locationContext, err := event.Context.Parent()
		assert.NoError(t, err)
		assert.Equal(t, "merchant1/location2", locationContext.String())

		merchantContext, err := locationContext.Parent()
		assert.NoError(t, err)
		assert.Equal(t, "merchant1", merchantContext.String())
	})
}

func TestStripeContext_EdgeCases(t *testing.T) {
	t.Run("parent of empty context", func(t *testing.T) {
		empty := ParseStripeContext("")
		parent, err := empty.Parent()
		assert.Error(t, err)
		assert.Nil(t, parent)
		assert.Contains(t, err.Error(), "empty StripeContext")
	})

	t.Run("parent of single segment returns empty", func(t *testing.T) {
		single := ParseStripeContext("single")
		parent, err := single.Parent()
		assert.NoError(t, err)
		assert.NotNil(t, parent)
		assert.True(t, parent.IsEmpty())
		assert.Equal(t, "", parent.String())
	})

	t.Run("child of empty context", func(t *testing.T) {
		empty := ParseStripeContext("")
		child := empty.Child("first")
		assert.Equal(t, "first", child.String())
		assert.Equal(t, []string{"first"}, child.Segments())
	})

	t.Run("context with forward slash in segment", func(t *testing.T) {
		// Note: This might be an edge case that the API doesn't support,
		// but our implementation should handle it consistently
		ctx := ParseStripeContext("account1/segment with / slash/account3")
		segments := ctx.Segments()
		expected := []string{"account1", "segment with ", " slash", "account3"}
		assert.Equal(t, expected, segments)
		assert.Equal(t, "account1/segment with / slash/account3", ctx.String())
	})
}