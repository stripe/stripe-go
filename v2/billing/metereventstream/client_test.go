package metereventstream_test

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

func TestMeterEventStreamNew(t *testing.T) {
	timeNow := time.Now()
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPost)
		assert.Equal(t, r.URL.Path, "/v2/billing/meter_event_stream")
		assert.Equal(t, r.Header.Get("Authorization"), "Bearer "+TestAPIKey)

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		params := &stripe.V2BillingMeterEventStreamParams{}
		if err := json.Unmarshal(body, params); err != nil {
			assert.NoError(t, err)
		}

		assert.Equal(t, params.Events[0].EventName, stripe.String("test_event"))
		assert.Equal(t, params.Events[0].Payload["stripe_customer_id"], "cus_123")
		assert.Equal(t, params.Events[0].Payload["value"], "100")
		assert.Equal(t, params.Events[0].Identifier, stripe.String("test_identifier"))
		assert.True(t, params.Events[0].Timestamp.Equal(timeNow))

		_, err = w.Write([]byte("{}"))
		assert.NoError(t, err)
	}))
	defer testServer.Close()

	backends := stripe.NewBackendsWithConfig(
		&stripe.BackendConfig{
			URL: stripe.String(testServer.URL),
		},
	)
	sc := client.New(TestAPIKey, backends)
	params := stripe.V2BillingMeterEventStreamParams{
		Events: []*stripe.V2BillingMeterEventStreamEventParams{
			{
				EventName: stripe.String("test_event"),
				Payload: map[string]string{
					"stripe_customer_id": "cus_123",
					"value":              "100",
				},
				Identifier: stripe.String("test_identifier"),
				Timestamp:  &timeNow,
			},
		},
	}
	err := sc.V2BillingMeterEventStreams.New(&params)
	assert.Nil(t, err)
}
