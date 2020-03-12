package earlyfraudwarning

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/channelmeter/stripe-go"
	_ "github.com/channelmeter/stripe-go/testing"
)

func TestRadarEarlyFraudWarningGet(t *testing.T) {
	ifr, err := Get("ifr_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, ifr)
}

func TestRadarEarlyFraudWarningList(t *testing.T) {
	i := List(&stripe.RadarEarlyFraudWarningListParams{})

	// Verify that we can get at least one issuer fraud record
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.RadarEarlyFraudWarning())
}

func TestRadarEarlyFraudWarningListByChargeID(t *testing.T) {
	i := List(&stripe.RadarEarlyFraudWarningListParams{
		Charge: stripe.String("ch_123"),
	})

	// Verify that we can get at least one issuer fraud record
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.RadarEarlyFraudWarning())
}
