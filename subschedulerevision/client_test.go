package subschedulerevision

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestSubscriptionScheduleRevisionGet(t *testing.T) {
	params := &stripe.SubscriptionScheduleRevisionParams{
		Schedule: stripe.String("sub_sched_123"),
	}
	revision, err := Get("sub_sched_rev_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, revision)
}

func TestSubscriptionScheduleRevisionList(t *testing.T) {
	params := &stripe.SubscriptionScheduleRevisionListParams{
		Schedule: stripe.String("sub_sched_123"),
	}
	i := List(params)

	// Verify that we can get at least one revision
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.SubscriptionScheduleRevision())
}
