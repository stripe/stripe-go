package verificationreport

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v75"
	_ "github.com/stripe/stripe-go/v75/testing"
)

func TestIdentityVerificationReportGet(t *testing.T) {
	report, err := Get("vr_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, report)
	assert.Equal(t, "identity.verification_report", report.Object)
}

func TestIdentityVerificationReportList(t *testing.T) {
	i := List(&stripe.IdentityVerificationReportListParams{})

	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.IdentityVerificationReport())
	assert.Equal(t, "identity.verification_report", i.IdentityVerificationReport().Object)
	assert.NotNil(t, i.IdentityVerificationReportList())
}
