package reportrun

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestReportingReportRunGet(t *testing.T) {
	reportrun, err := Get("frr_123", nil)
	assert.Nil(t, err)
	assert.NotNil(t, reportrun)
	assert.Equal(t, "reporting.report_run", reportrun.Object)
}

func TestReportingReportRunList(t *testing.T) {
	i := List(&stripe.ReportingReportRunListParams{})

	// Verify that we can get at least one reportrun
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.ReportingReportRun())
	assert.Equal(t, "reporting.report_run", i.ReportingReportRun().Object)
	assert.NotNil(t, i.ReportingReportRunList())
}

func TestReportingReportRunNew(t *testing.T) {
	reportrun, err := New(&stripe.ReportingReportRunParams{
		Parameters: &stripe.ReportingReportRunParametersParams{
			ConnectedAccount: stripe.String("acct_123"),
		},
		ReportType: stripe.String("activity.summary.1"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, reportrun)
	assert.Equal(t, "reporting.report_run", reportrun.Object)
}
