package usagerecord

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v75"
	_ "github.com/stripe/stripe-go/v75/testing"
)

func TestUsageRecordNew(t *testing.T) {
	now := int64(time.Now().Unix())
	usageRecord, err := New(&stripe.UsageRecordParams{
		Quantity:         stripe.Int64(123),
		Timestamp:        stripe.Int64(now),
		Action:           stripe.String(stripe.UsageRecordActionIncrement),
		SubscriptionItem: stripe.String("si_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, usageRecord)
}
