package usagerecord

import (
	"testing"
	"time"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v81"
	_ "github.com/stripe/stripe-go/v81/testing"
)

func TestUsageRecordNew(t *testing.T) {
	usageRecord, err := New(&stripe.UsageRecordParams{
		Quantity:         stripe.Int64(123),
		Timestamp:        stripe.Time(time.Now()),
		Action:           stripe.String(stripe.UsageRecordActionIncrement),
		SubscriptionItem: stripe.String("si_123"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, usageRecord)
}
