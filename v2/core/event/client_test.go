package event_test

import (
	"encoding/json"
	"net/http"
	"testing"
	"time"

	"fmt"

	stripe "github.com/max-cape/stripe-go-test"
	"github.com/max-cape/stripe-go-test/mock"
	assert "github.com/stretchr/testify/require"
)

func TestEventGet(t *testing.T) {
	timeNow := time.Now()
	params := stripe.V2CoreEventParams{}
	testServer, sc := mock.Server(t, string(http.MethodGet), "/v2/core/events/evt_123", nil, func(p *stripe.V2CoreEventParams) []byte {
		data, err := json.Marshal(stripe.V1BillingMeterErrorReportTriggeredEvent{
			V2RawEvent: stripe.V2RawEvent{
				V2BaseEvent: stripe.V2BaseEvent{
					Created: timeNow,
					ID:      "evt_123",
					Type:    "v1.billing.meter.error_report_triggered",
				},
				RelatedObject: &stripe.RelatedObject{
					ID: "ro_123",
				},
			},
			Data: stripe.V1BillingMeterErrorReportTriggeredEventData{
				DeveloperMessageSummary: "This is a developer message",
			},
			RelatedObject: stripe.RelatedObject{
				ID: "ro_123",
			},
		})
		assert.NoError(t, err)
		return data
	})
	defer testServer.Close()

	event, err := sc.V2CoreEvents.Get("evt_123", &params)
	assert.Nil(t, err)
	v1BillingEvent, ok := event.(*stripe.V1BillingMeterErrorReportTriggeredEvent)
	assert.True(t, ok)
	assert.Equal(t, v1BillingEvent.ID, "evt_123")
	assert.Equal(t, v1BillingEvent.Type, "v1.billing.meter.error_report_triggered")
	assert.True(t, v1BillingEvent.Created.Equal(timeNow))
	assert.Equal(t, v1BillingEvent.Data.DeveloperMessageSummary, "This is a developer message")
	assert.Equal(t, v1BillingEvent.RelatedObject.ID, "ro_123")
}

func TestEventAll(t *testing.T) {
	timeNow := time.Now()
	params := &stripe.V2CoreEventListParams{}
	testServer, sc := mock.Server(t, string(http.MethodGet), "/v2/core/events", params, func(p *stripe.V2CoreEventListParams) []byte {
		data, err := json.Marshal(stripe.V2Page[stripe.V2Event]{
			Data: []stripe.V2Event{
				&stripe.V1BillingMeterErrorReportTriggeredEvent{
					V2RawEvent: stripe.V2RawEvent{
						V2BaseEvent: stripe.V2BaseEvent{
							Created: timeNow,
							ID:      "evt_1",
							Type:    "v1.billing.meter.error_report_triggered",
						},
						RelatedObject: &stripe.RelatedObject{
							ID: "ro_123",
						},
					},
					Data: stripe.V1BillingMeterErrorReportTriggeredEventData{
						DeveloperMessageSummary: "This is a developer message",
					},
					RelatedObject: stripe.RelatedObject{
						ID: "ro_123",
					},
				},
				&stripe.V1BillingMeterNoMeterFoundEvent{
					V2RawEvent: stripe.V2RawEvent{
						V2BaseEvent: stripe.V2BaseEvent{
							Created: timeNow,
							ID:      "evt_2",
							Type:    "v1.billing.meter.no_meter_found",
						},
					},
					Data: stripe.V1BillingMeterNoMeterFoundEventData{
						DeveloperMessageSummary: "This is another developer message",
					},
				},
				&stripe.V1BillingMeterNoMeterFoundEvent{
					V2RawEvent: stripe.V2RawEvent{
						V2BaseEvent: stripe.V2BaseEvent{
							Created: timeNow,
							ID:      "evt_3",
							Type:    "v1.billing.meter.no_meter_found",
						},
					},
					Data: stripe.V1BillingMeterNoMeterFoundEventData{
						DeveloperMessageSummary: "This is yet another developer message",
					},
				},
			},
		})
		assert.NoError(t, err)
		return data
	})
	defer testServer.Close()
	cnt := 1
	sc.V2CoreEvents.All(params)(func(event stripe.V2Event, err error) bool {
		assert.Nil(t, err)
		assert.NotNil(t, event)
		if cnt == 1 {
			v1BillingEvent, ok := event.(*stripe.V1BillingMeterErrorReportTriggeredEvent)
			assert.True(t, ok)
			assert.Equal(t, v1BillingEvent.ID, "evt_1")
			assert.Equal(t, v1BillingEvent.Type, "v1.billing.meter.error_report_triggered")
			assert.True(t, v1BillingEvent.Created.Equal(timeNow))
			assert.Equal(t, v1BillingEvent.Data.DeveloperMessageSummary, "This is a developer message")
			assert.Equal(t, v1BillingEvent.RelatedObject.ID, "ro_123")
		} else {
			v1BillingEvent, ok := event.(*stripe.V1BillingMeterNoMeterFoundEvent)
			assert.True(t, ok)
			assert.Equal(t, v1BillingEvent.ID, fmt.Sprintf("evt_%d", cnt))
			assert.Equal(t, v1BillingEvent.Type, "v1.billing.meter.no_meter_found")
			assert.True(t, v1BillingEvent.Created.Equal(timeNow))
			assert.Equal(t, v1BillingEvent.Data.DeveloperMessageSummary, fmt.Sprintf("This is %s developer message", map[int]string{2: "another", 3: "yet another"}[cnt]))
		}
		cnt++

		return true
	})
	assert.Equal(t, cnt, 4)
}
