package reporttype

import (
	"testing"

	stripe "github.com/max-cape/stripe-go-test"
	_ "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestReportTestGet(t *testing.T) {
	reporttype, err := Get("activity.summary.1", nil)
	assert.Nil(t, err)
	assert.NotNil(t, reporttype)
	assert.Equal(t, "reporting.report_type", reporttype.Object)
}

func TestReportTestList(t *testing.T) {
	i := List(&stripe.ReportingReportTypeListParams{})

	// Verify that we can get at least one reporttype
	assert.True(t, i.Next())
	assert.Nil(t, i.Err())
	assert.NotNil(t, i.ReportingReportType())
	assert.Equal(t, "reporting.report_type", i.ReportingReportType().Object)
	assert.NotNil(t, i.ReportingReportTypeList())
}
