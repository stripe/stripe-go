package usagerecord

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestUsageRecordNew(t *testing.T) {
	now := uint64(time.Now().Unix())
	usageRecord, err := New(&stripe.UsageRecordParams{
		Quantity:         123,
		Timestamp:        now,
		Action:           stripe.UsageRecordParamsActionIncrement,
		SubscriptionItem: "si_123",
	})
	assert.Nil(t, err)
	assert.NotNil(t, usageRecord)
}
