package issuerfraudrecord

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go"
	_ "github.com/stripe/stripe-go/testing"
)

func TestIssuerFraudRecordGet(t *testing.T) {
	topup, err := Get("ifr_123")
	assert.Nil(t, err)
	assert.NotNil(t, topup)
}

func TestIssuerFraudRecordList(t *testing.T) {
	i := List(&stripe.IssuerFraudRecordListParams{})

	// Verify that we can get at least one issuer fraud record
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.IssuerFraudRecord())
}

func TestIssuerFraudRecordListByChargeID(t *testing.T) {
	i := List(&stripe.IssuerFraudRecordListParams{Charge: "ch_123"})

	// Verify that we can get at least one issuer fraud record
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.IssuerFraudRecord())
}
