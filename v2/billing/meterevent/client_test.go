package meterevent_test

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	stripe "github.com/max-cape/stripe-go-test"
	"github.com/max-cape/stripe-go-test/client"
	. "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestMeterEventNew(t *testing.T) {
	timeNow := time.Now()
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPost)
		assert.Equal(t, r.URL.Path, "/v2/billing/meter_events")
		assert.Equal(t, r.Header.Get("Authorization"), "Bearer "+TestAPIKey)

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		params := &stripe.V2BillingMeterEventParams{}
		if err := json.Unmarshal(body, params); err != nil {
			assert.NoError(t, err)
		}

		data, err := json.Marshal(stripe.V2BillingMeterEvent{
			Created:    timeNow,
			Timestamp:  timeNow,
			EventName:  *params.EventName,
			Identifier: *params.Identifier,
			Payload:    params.Payload,
		})
		assert.NoError(t, err)

		_, err = w.Write(data)
		assert.NoError(t, err)
	}))
	defer testServer.Close()

	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{
			URL: stripe.String(testServer.URL),
		},
	)
	params := stripe.V2BillingMeterEventParams{
		EventName: stripe.String("test_event"),
		Payload: map[string]string{
			"stripe_customer_id": "cus_123",
			"value":              "100",
		},
		Identifier: stripe.String("test_identifier"),
	}
	sc := client.New(TestAPIKey, backends)
	event, err := sc.V2BillingMeterEvents.New(&params)
	assert.Nil(t, err)
	assert.Equal(t, event.EventName, *params.EventName)
	assert.Equal(t, event.Identifier, *params.Identifier)
	assert.Equal(t, event.Payload, params.Payload)
	assert.True(t, event.Created.Equal(timeNow))
	assert.True(t, event.Timestamp.Equal(timeNow))
}
