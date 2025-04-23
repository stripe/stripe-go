package eventdestination_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"fmt"

	stripe "github.com/max-cape/stripe-go-test"
	"github.com/max-cape/stripe-go-test/client"
	"github.com/max-cape/stripe-go-test/mock"
	. "github.com/max-cape/stripe-go-test/testing"
	assert "github.com/stretchr/testify/require"
)

func TestEventDestinationNew(t *testing.T) {
	timeNow := time.Now()
	params := &stripe.V2CoreEventDestinationParams{
		Name:          stripe.String("My Event Destination"),
		Description:   stripe.String("This is my event destination, I like it a lot"),
		EnabledEvents: stripe.StringSlice([]string{"v1.billing.meter.error_report_triggered"}),
		Type:          stripe.String(string(stripe.V2EventDestinationTypeWebhookEndpoint)),
		WebhookEndpoint: &stripe.V2CoreEventDestinationWebhookEndpointParams{
			URL: stripe.String("https://example.com/my/webhook/endpoint"),
		},
		EventPayload: stripe.String(string(stripe.V2EventDestinationEventPayloadThin)),
		Include:      stripe.StringSlice([]string{"webhook_endpoint.url"}),
	}
	testServer, sc := mock.Server(t, string(http.MethodPost), "/v2/core/event_destinations", params, func(p *stripe.V2CoreEventDestinationParams) []byte {
		var enabledEvents []string
		for _, event := range params.EnabledEvents {
			enabledEvents = append(enabledEvents, *event)
		}
		data, err := json.Marshal(stripe.V2EventDestination{
			Created:       timeNow,
			Description:   *params.Description,
			EnabledEvents: enabledEvents,
			EventPayload:  stripe.V2EventDestinationEventPayload(*params.EventPayload),
			EventsFrom:    []stripe.V2EventDestinationEventsFrom{stripe.V2EventDestinationEventsFromSelf},
			Name:          *params.Name,
			Status:        stripe.V2EventDestinationStatusEnabled,
			Type:          stripe.V2EventDestinationType(*params.Type),
			Updated:       timeNow,
			WebhookEndpoint: &stripe.V2EventDestinationWebhookEndpoint{
				URL: *params.WebhookEndpoint.URL,
			},
		})
		assert.NoError(t, err)
		return data
	})
	defer testServer.Close()

	dest, err := sc.V2CoreEventDestinations.New(params)
	assert.Nil(t, err)
	assert.Equal(t, dest.Name, *params.Name)
	assert.Equal(t, dest.Description, *params.Description)
	for i, event := range dest.EnabledEvents {
		assert.Equal(t, event, *params.EnabledEvents[i])
	}
	assert.Equal(t, dest.Type, stripe.V2EventDestinationType(*params.Type))
	assert.Equal(t, dest.WebhookEndpoint.URL, *params.WebhookEndpoint.URL)
	assert.Equal(t, dest.EventPayload, stripe.V2EventDestinationEventPayload(*params.EventPayload))
	assert.Equal(t, dest.Status, stripe.V2EventDestinationStatusEnabled)
	assert.True(t, dest.Created.Equal(timeNow))
	assert.True(t, dest.Updated.Equal(timeNow))
	assert.Equal(t, dest.EventsFrom, []stripe.V2EventDestinationEventsFrom{stripe.V2EventDestinationEventsFromSelf})
}

func TestEventDestinationGet(t *testing.T) {
	timeNow := time.Now()
	params := stripe.V2CoreEventDestinationParams{
		Include: stripe.StringSlice([]string{"webhook_endpoint.url"}),
	}
	testServer, sc := mock.Server(t, string(http.MethodGet), "/v2/core/event_destinations/ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6", nil, func(p *stripe.V2CoreEventDestinationParams) []byte {
		data, err := json.Marshal(stripe.V2EventDestination{
			Created:       timeNow,
			Description:   "This is my event destination, I like it a lot",
			EnabledEvents: []string{"v1.billing.meter.error_report_triggered"},
			EventPayload:  stripe.V2EventDestinationEventPayloadThin,
			EventsFrom:    []stripe.V2EventDestinationEventsFrom{stripe.V2EventDestinationEventsFromSelf},
			Name:          "My Event Destination",
			Status:        stripe.V2EventDestinationStatusEnabled,
			Type:          stripe.V2EventDestinationTypeWebhookEndpoint,
			Updated:       timeNow,
			ID:            "ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6",
		})
		assert.NoError(t, err)
		return data
	}, func(t *testing.T, r *http.Request) {
		assert.Equal(t, r.URL.Query().Get("include[0]"), "webhook_endpoint.url")
	})
	defer testServer.Close()

	dest, err := sc.V2CoreEventDestinations.Get("ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6", &params)
	assert.Nil(t, err)
	assert.Equal(t, dest.Name, "My Event Destination")
	assert.Equal(t, dest.Description, "This is my event destination, I like it a lot")
	assert.Equal(t, dest.EnabledEvents, []string{"v1.billing.meter.error_report_triggered"})
	assert.Equal(t, dest.Type, stripe.V2EventDestinationTypeWebhookEndpoint)
	assert.Equal(t, dest.EventPayload, stripe.V2EventDestinationEventPayloadThin)
	assert.Equal(t, dest.Status, stripe.V2EventDestinationStatusEnabled)
	assert.True(t, dest.Created.Equal(timeNow))
	assert.True(t, dest.Updated.Equal(timeNow))
	assert.Equal(t, dest.EventsFrom, []stripe.V2EventDestinationEventsFrom{stripe.V2EventDestinationEventsFromSelf})
	assert.Equal(t, dest.ID, "ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6")
}

func TestEventDestinationUpdate(t *testing.T) {
	timeNow := time.Now()
	params := stripe.V2CoreEventDestinationParams{
		Description: stripe.String("Better description"),
	}
	testServer, sc := mock.Server(t, string(http.MethodPost), "/v2/core/event_destinations/ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6", &params, func(p *stripe.V2CoreEventDestinationParams) []byte {
		data, err := json.Marshal(stripe.V2EventDestination{
			Created:       timeNow,
			Description:   "Better description",
			EnabledEvents: []string{"v1.billing.meter.error_report_triggered"},
			EventPayload:  stripe.V2EventDestinationEventPayloadThin,
			EventsFrom:    []stripe.V2EventDestinationEventsFrom{stripe.V2EventDestinationEventsFromSelf},
			Name:          "My Event Destination",
			Status:        stripe.V2EventDestinationStatusEnabled,
			Type:          stripe.V2EventDestinationTypeWebhookEndpoint,
			Updated:       timeNow,
			ID:            "ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6",
		})
		assert.NoError(t, err)
		return data
	})
	defer testServer.Close()

	dest, err := sc.V2CoreEventDestinations.Update("ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6", &params)
	assert.Nil(t, err)
	assert.Equal(t, dest.Name, "My Event Destination")
	assert.Equal(t, dest.Description, *params.Description)
	assert.Equal(t, dest.EnabledEvents, []string{"v1.billing.meter.error_report_triggered"})
	assert.Equal(t, dest.Type, stripe.V2EventDestinationTypeWebhookEndpoint)
	assert.Equal(t, dest.EventPayload, stripe.V2EventDestinationEventPayloadThin)
	assert.Equal(t, dest.Status, stripe.V2EventDestinationStatusEnabled)
	assert.True(t, dest.Created.Equal(timeNow))
	assert.True(t, dest.Updated.Equal(timeNow))
	assert.Equal(t, dest.EventsFrom, []stripe.V2EventDestinationEventsFrom{stripe.V2EventDestinationEventsFromSelf})
	assert.Equal(t, dest.ID, "ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6")
}

func TestEventDestinationDel(t *testing.T) {
	testServer, sc := mock.Server(t, string(http.MethodDelete), "/v2/core/event_destinations/ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6", nil, func(p *stripe.V2CoreEventDestinationParams) []byte {
		data, err := json.Marshal(stripe.V2EventDestination{
			ID: "ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6",
		})
		assert.NoError(t, err)
		return data
	})
	defer testServer.Close()
	dest, err := sc.V2CoreEventDestinations.Del("ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6", nil)
	assert.Nil(t, err)
	assert.Equal(t, dest.ID, "ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6")
}

func TestEventDestinationDisable(t *testing.T) {
	timeNow := time.Now()
	testServer, sc := mock.Server(t, string(http.MethodPost), "/v2/core/event_destinations/ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6/disable", nil, func(p *stripe.V2CoreEventDestinationParams) []byte {
		data, err := json.Marshal(stripe.V2EventDestination{
			Created:       timeNow,
			Description:   "This is my event destination, I like it a lot",
			EnabledEvents: []string{"v1.billing.meter.error_report_triggered"},
			EventPayload:  stripe.V2EventDestinationEventPayloadThin,
			EventsFrom:    []stripe.V2EventDestinationEventsFrom{stripe.V2EventDestinationEventsFromSelf},
			Name:          "My Event Destination",
			Status:        stripe.V2EventDestinationStatusDisabled,
			Type:          stripe.V2EventDestinationTypeWebhookEndpoint,
			Updated:       timeNow,
			ID:            "ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6",
		})
		assert.NoError(t, err)
		return data
	})
	defer testServer.Close()
	dest, err := sc.V2CoreEventDestinations.Disable("ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6", nil)
	assert.Nil(t, err)
	assert.Equal(t, dest.ID, "ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6")
	assert.Equal(t, dest.Status, stripe.V2EventDestinationStatusDisabled)
}

func TestEventDestinationEnable(t *testing.T) {
	timeNow := time.Now()
	testServer, sc := mock.Server(t, string(http.MethodPost), "/v2/core/event_destinations/ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6/enable", nil, func(p *stripe.V2CoreEventDestinationParams) []byte {
		data, err := json.Marshal(stripe.V2EventDestination{
			Created:       timeNow,
			Description:   "This is my event destination, I like it a lot",
			EnabledEvents: []string{"v1.billing.meter.error_report_triggered"},
			EventPayload:  stripe.V2EventDestinationEventPayloadThin,
			EventsFrom:    []stripe.V2EventDestinationEventsFrom{stripe.V2EventDestinationEventsFromSelf},
			Name:          "My Event Destination",
			Status:        stripe.V2EventDestinationStatusEnabled,
			Type:          stripe.V2EventDestinationTypeWebhookEndpoint,
			Updated:       timeNow,
			ID:            "ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6",
		})
		assert.NoError(t, err)
		return data
	})
	defer testServer.Close()
	dest, err := sc.V2CoreEventDestinations.Enable("ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6", nil)
	assert.Nil(t, err)
	assert.Equal(t, dest.ID, "ed_test_61RM8ltWcTW4mbsxf16RJyfa2xSQLHJJh1sxm7H0KVT6")
	assert.Equal(t, dest.Status, stripe.V2EventDestinationStatusEnabled)
}

func TestEventDestinationPing(t *testing.T) {
	timeNow := time.Now()
	testServer, sc := mock.Server(t, string(http.MethodPost), "/v2/core/event_destinations/evt_test_65RM8sQH2oXnebF5Rpc16RJyfa2xSQLHJJh1sxm7H0KI92/ping", nil, func(p *stripe.V2CoreEventDestinationParams) []byte {
		data, err := json.Marshal(stripe.V2RawEvent{V2BaseEvent: stripe.V2BaseEvent{
			ID:      "evt_test_65RM8sQH2oXnebF5Rpc16RJyfa2xSQLHJJh1sxm7H0KI92",
			Created: timeNow,
			Object:  "v2.core.event",
			Type:    "v2.core.event_destination.ping",
		}})
		assert.NoError(t, err)
		return data
	})
	defer testServer.Close()
	event, err := sc.V2CoreEventDestinations.Ping("evt_test_65RM8sQH2oXnebF5Rpc16RJyfa2xSQLHJJh1sxm7H0KI92", nil)
	assert.Nil(t, err)
	assert.Equal(t, event.GetBaseEvent().ID, "evt_test_65RM8sQH2oXnebF5Rpc16RJyfa2xSQLHJJh1sxm7H0KI92")
	assert.Equal(t, event.GetBaseEvent().Type, "v2.core.event_destination.ping")
	assert.True(t, event.GetBaseEvent().Created.Equal(timeNow))
	assert.Equal(t, event.GetBaseEvent().Object, "v2.core.event")
}

func TestEventDestinationList_SinglePage(t *testing.T) {
	params := &stripe.V2CoreEventDestinationListParams{}
	testServer, sc := mock.Server(t, string(http.MethodGet), "/v2/core/event_destinations", params, func(p *stripe.V2CoreEventDestinationListParams) []byte {
		data, err := json.Marshal(stripe.V2Page[stripe.V2EventDestination]{
			Data: []stripe.V2EventDestination{
				{
					Description:   "Event destination 1",
					EnabledEvents: []string{"v1.billing.meter.error_report_triggered"},
					EventPayload:  stripe.V2EventDestinationEventPayloadThin,
					ID:            "ed_test_1",
				},
				{
					Description:   "Event destination 2",
					EnabledEvents: []string{"v1.billing.meter.error_report_triggered"},
					EventPayload:  stripe.V2EventDestinationEventPayloadThin,
					ID:            "ed_test_2",
				},
				{
					Description:   "Event destination 3",
					EnabledEvents: []string{"v1.billing.meter.error_report_triggered"},
					EventPayload:  stripe.V2EventDestinationEventPayloadThin,
					ID:            "ed_test_3",
				},
			},
		})
		assert.NoError(t, err)
		return data
	})
	defer testServer.Close()
	cnt := 1
	sc.V2CoreEventDestinations.All(params)(func(dest *stripe.V2EventDestination, err error) bool {
		assert.Nil(t, err)
		assert.NotNil(t, dest)
		assert.Equal(t, dest.Description, fmt.Sprintf("Event destination %d", cnt))
		assert.Equal(t, dest.EnabledEvents, []string{"v1.billing.meter.error_report_triggered"})
		assert.Equal(t, dest.EventPayload, stripe.V2EventDestinationEventPayloadThin)
		assert.Equal(t, dest.ID, fmt.Sprintf("ed_test_%d", cnt))
		cnt++
		return true
	})
	assert.Equal(t, cnt, 4)
}

func TestEventDestinationList_MultiplePages(t *testing.T) {
	params := &stripe.V2CoreEventDestinationListParams{
		Limit: stripe.Int64(1),
	}
	cnt := 1
	testServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "1", r.URL.Query().Get("limit"))
		assert.Equal(t, "/v2/core/event_destinations", r.URL.Path)
		page := r.URL.Query().Get("page")
		var data []byte
		var err error
		switch page {
		case "", "page_1":
			data, err = json.Marshal(stripe.V2Page[stripe.V2EventDestination]{
				NextPageURL: fmt.Sprintf("/v2/core/event_destinations?limit=1&page=page_%d", cnt),
				Data: []stripe.V2EventDestination{
					{
						Description:   fmt.Sprintf("Event destination %d", cnt),
						EnabledEvents: []string{"v1.billing.meter.error_report_triggered"},
						EventPayload:  stripe.V2EventDestinationEventPayloadThin,
						ID:            fmt.Sprintf("ed_test_%d", cnt),
					},
				},
			})
		case "page_2":
			data, err = json.Marshal(stripe.V2Page[stripe.V2EventDestination]{
				Data: []stripe.V2EventDestination{
					{
						Description:   fmt.Sprintf("Event destination %d", cnt),
						EnabledEvents: []string{"v1.billing.meter.error_report_triggered"},
						EventPayload:  stripe.V2EventDestinationEventPayloadThin,
						ID:            fmt.Sprintf("ed_test_%d", cnt),
					},
				},
			})
		default:
			assert.Fail(t, fmt.Sprintf("unexpected page: %s", page))
		}
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

	sc.V2CoreEventDestinations.All(params)(func(dest *stripe.V2EventDestination, err error) bool {
		assert.Nil(t, err)
		assert.NotNil(t, dest)
		assert.Equal(t, dest.Description, fmt.Sprintf("Event destination %d", cnt))
		assert.Equal(t, dest.EnabledEvents, []string{"v1.billing.meter.error_report_triggered"})
		assert.Equal(t, dest.EventPayload, stripe.V2EventDestinationEventPayloadThin)
		assert.Equal(t, dest.ID, fmt.Sprintf("ed_test_%d", cnt))
		cnt++
		return true
	})
	assert.Equal(t, cnt, 4)
}
