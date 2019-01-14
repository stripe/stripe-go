package issuerfraudrecord

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v55"
	_ "github.com/stripe/stripe-go/v55/testing"
)

func TestIssuerFraudRecordGet(t *testing.T) {
	ifr, err := Get("ifr_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, ifr)
}

func TestIssuerFraudRecordList(t *testing.T) {
	i := List(&stripe.IssuerFraudRecordListParams{})

	// Verify that we can get at least one issuer fraud record
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.IssuerFraudRecord())
}

func TestIssuerFraudRecordListByChargeID(t *testing.T) {
	i := List(&stripe.IssuerFraudRecordListParams{
		Charge: stripe.String("ch_123"),
	})

	// Verify that we can get at least one issuer fraud record
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.IssuerFraudRecord())
}
