package stripe

import (
	"encoding/json"
	"testing"

	"github.com/max-cape/stripe-go-test/form"
	assert "github.com/stretchr/testify/require"
)

func TestSubscriptionParams_AppendTo(t *testing.T) {
	{
		params := &SubscriptionParams{BillingCycleAnchorNow: Bool(true)}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"now"}, body.Get("billing_cycle_anchor"))
	}

	{
		params := &SubscriptionParams{BillingCycleAnchorUnchanged: Bool(true)}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"unchanged"}, body.Get("billing_cycle_anchor"))
	}

	{
		params := &SubscriptionParams{TrialEndNow: Bool(true)}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"now"}, body.Get("trial_end"))
	}
}

func TestSubscription_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Subscription
		err := json.Unmarshal([]byte(`"sub_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "sub_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Subscription{ID: "sub_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "sub_123", v.ID)
	}
}
