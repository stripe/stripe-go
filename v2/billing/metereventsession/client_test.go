package metereventsession_test

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

func TestMeterEventSessionNew(t *testing.T) {
	timeNow := time.Now()
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPost)
		assert.Equal(t, r.URL.Path, "/v2/billing/meter_event_session")
		assert.Equal(t, r.Header.Get("Authorization"), "Bearer "+TestAPIKey)

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		params := &stripe.V2BillingMeterEventAdjustmentParams{}
		if len(body) > 0 {
			if err := json.Unmarshal(body, params); err != nil {
				assert.NoError(t, err)
			}
		}

		data, err := json.Marshal(stripe.V2BillingMeterEventSession{
			Created:             timeNow,
			ExpiresAt:           timeNow.Add(time.Hour),
			Object:              "v2.billing.meter_event_session",
			ID:                  "test_session_id",
			AuthenticationToken: "test_authentication_token",
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
	sc := client.New(TestAPIKey, backends)
	params := stripe.V2BillingMeterEventSessionParams{}
	event, err := sc.V2BillingMeterEventSessions.New(&params)
	assert.Nil(t, err)
	assert.True(t, event.Created.Equal(timeNow))
	assert.Equal(t, event.Object, "v2.billing.meter_event_session")
	assert.Equal(t, event.ID, "test_session_id")
	assert.Equal(t, event.AuthenticationToken, "test_authentication_token")
	assert.True(t, event.ExpiresAt.Equal(timeNow.Add(time.Hour)))
}
