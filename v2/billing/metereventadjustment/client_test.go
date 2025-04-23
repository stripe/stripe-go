package metereventadjustment_test

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

func TestMeterEventAdjustmentNew(t *testing.T) {
	timeNow := time.Now()
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, r.Method, http.MethodPost)
		assert.Equal(t, r.URL.Path, "/v2/billing/meter_event_adjustments")
		assert.Equal(t, r.Header.Get("Authorization"), "Bearer "+TestAPIKey)

		body, err := io.ReadAll(r.Body)
		assert.NoError(t, err)

		params := &stripe.V2BillingMeterEventAdjustmentParams{}
		if err := json.Unmarshal(body, params); err != nil {
			assert.NoError(t, err)
		}

		data, err := json.Marshal(stripe.V2BillingMeterEventAdjustment{
			Created:   timeNow,
			EventName: *params.EventName,
			Cancel: &stripe.V2BillingMeterEventAdjustmentCancel{
				Identifier: *params.Cancel.Identifier,
			},
			Type:   stripe.V2BillingMeterEventAdjustmentTypeCancel,
			Status: stripe.V2BillingMeterEventAdjustmentStatusPending,
			Object: "v2.billing.meter_event_adjustment",
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
	params := stripe.V2BillingMeterEventAdjustmentParams{
		EventName: stripe.String("test_event"),
		Cancel: &stripe.V2BillingMeterEventAdjustmentCancelParams{
			Identifier: stripe.String("test_identifier"),
		},
		Type: stripe.String(string(stripe.V2BillingMeterEventAdjustmentTypeCancel)),
	}
	sc := client.New(TestAPIKey, backends)
	event, err := sc.V2BillingMeterEventAdjustments.New(&params)
	assert.Nil(t, err)
	assert.Equal(t, event.EventName, *params.EventName)
	assert.Equal(t, event.Cancel.Identifier, *params.Cancel.Identifier)
	assert.Equal(t, event.Type, stripe.V2BillingMeterEventAdjustmentTypeCancel)
	assert.Equal(t, event.Status, stripe.V2BillingMeterEventAdjustmentStatusPending)
	assert.True(t, event.Created.Equal(timeNow))
	assert.Equal(t, event.Object, "v2.billing.meter_event_adjustment")
}
