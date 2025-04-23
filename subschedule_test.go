package stripe

import (
	"encoding/json"
	"testing"

	"github.com/max-cape/stripe-go-test/form"
	assert "github.com/stretchr/testify/require"
)

func TestSubscriptionScheduleParams_AppendTo(t *testing.T) {
	{
		params := &SubscriptionScheduleParams{StartDateNow: Bool(true)}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"now"}, body.Get("start_date"))
	}
}

func TestSubscriptionSchedule_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v SubscriptionSchedule
		err := json.Unmarshal([]byte(`"sub_sched_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "sub_sched_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := SubscriptionSchedule{ID: "sub_sched_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "sub_sched_123", v.ID)
	}
}
