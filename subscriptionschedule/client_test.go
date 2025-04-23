package subscriptionschedule

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestSubscriptionScheduleCancel(t *testing.T) {
	params := &stripe.SubscriptionScheduleCancelParams{
		InvoiceNow: stripe.Bool(true),
		Prorate:    stripe.Bool(true),
	}
	schedule, err := Cancel("sub_sched_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, schedule)
}

func TestSubscriptionScheduleGet(t *testing.T) {
	schedule, err := Get("sub_sched_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, schedule)
}

func TestSubscriptionScheduleList(t *testing.T) {
	params := &stripe.SubscriptionScheduleListParams{
		ReleasedAtRange: &stripe.RangeQueryParams{
			GreaterThanOrEqual: 123456,
		},
	}
	i := List(params)

	// Verify that we can get at least one schedule
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.SubscriptionSchedule())
	assert.NotNil(t, i.SubscriptionScheduleList())
}

func TestSubscriptionScheduleNew(t *testing.T) {
	params := &stripe.SubscriptionScheduleParams{
		Customer:     stripe.String("cus_123"),
		StartDateNow: stripe.Bool(true),
		Phases: []*stripe.SubscriptionSchedulePhaseParams{
			{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					{
						Plan:     stripe.String("plan_123"),
						Quantity: stripe.Int64(10),
					},
					{
						Plan:     stripe.String("plan_456"),
						Quantity: stripe.Int64(20),
					},
				},
				Trial: stripe.Bool(true),
			},
			{
				Items: []*stripe.SubscriptionSchedulePhaseItemParams{
					{
						Plan:     stripe.String("plan_789"),
						Quantity: stripe.Int64(30),
					},
				},
			},
		},
		EndBehavior: stripe.String(string(stripe.SubscriptionScheduleEndBehaviorCancel)),
	}
	schedule, err := New(params)
	assert.Nil(t, err)
	assert.NotNil(t, schedule)
}

func TestSubscriptionScheduleRelease(t *testing.T) {
	params := &stripe.SubscriptionScheduleReleaseParams{
		PreserveCancelDate: stripe.Bool(true),
	}
	schedule, err := Release("sub_sched_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, schedule)
}

func TestSubscriptionScheduleUpdate(t *testing.T) {
	params := &stripe.SubscriptionScheduleParams{
		Params: stripe.Params{
			Metadata: map[string]string{
				"foo": "bar",
			},
		},
	}
	schedule, err := Update("sub_sched_123", params)
	assert.Nil(t, err)
	assert.NotNil(t, schedule)
}
