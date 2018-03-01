package stripe

import (
	"strconv"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/form"
)

func TestUsageRecordParams_AppendTo(t *testing.T) {
	testCases := []struct {
		field  string
		params *UsageRecordParams
		want   interface{}
	}{
		{"quantity", &UsageRecordParams{Quantity: 2000}, strconv.FormatUint(2000, 10)},
		{"quantity", &UsageRecordParams{QuantityZero: true}, strconv.FormatUint(0, 10)},
		{"timestamp", &UsageRecordParams{Timestamp: 123123123}, strconv.FormatUint(123123123, 10)},
		{"action", &UsageRecordParams{Action: "increment"}, "increment"},
	}
	for _, tc := range testCases {
		t.Run(tc.field, func(t *testing.T) {
			body := &form.Values{}
			form.AppendTo(body, tc.params)
			values := body.ToValues()
			assert.Equal(t, tc.want, values.Get(tc.field))
		})
	}
}
