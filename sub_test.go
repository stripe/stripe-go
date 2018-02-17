package stripe

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/form"
)

func TestSubscriptionParams_AppendTo(t *testing.T) {
	{
		params := &SubscriptionParams{BillingCycleAnchorNow: true}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"now"}, body.Get("billing_cycle_anchor"))
	}

	{
		params := &SubscriptionParams{BillingCycleAnchorUnchanged: true}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"unchanged"}, body.Get("billing_cycle_anchor"))
	}

	{
		params := &SubscriptionParams{TrialEndNow: true}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"now"}, body.Get("trial_end"))
	}
}
