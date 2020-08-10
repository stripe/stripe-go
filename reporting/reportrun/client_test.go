package reportrun

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/stripe/stripe-go/v71"
	_ "github.com/stripe/stripe-go/v71/testing"
)

func TestReportRunGet(t *testing.T) {
	reportrun, err := Get("frr_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, reportrun)
	assert.Equal(t, "reporting.report_run", reportrun.Object)
}

func TestReportRunList(t *testing.T) {
	i := List(&stripe.ReportRunListParams{})

	// Verify that we can get at least one reportrun
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.ReportRun())
	assert.Equal(t, "reporting.report_run", i.ReportRun().Object)
	assert.NotNil(t, i.ReportRunList())
}

func TestReportRunNew(t *testing.T) {
	reportrun, err := New(&stripe.ReportRunParams{
		Parameters: &stripe.ReportRunParametersParams{
			ConnectedAccount: stripe.String("acct_123"),
		},
		ReportType: stripe.String("activity.summary.1"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reportrun)
	assert.Equal(t, "reporting.report_run", reportrun.Object)
}
