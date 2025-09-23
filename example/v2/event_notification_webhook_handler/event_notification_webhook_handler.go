// event_notification_webhook_handler.go - receive and process thin events like the
// "v1.billing.meter.error_report_triggered" event.
//
// This example uses the rawrequests module to make calls to /v2 APIs.
//
// In this example, we:
//   - parse the incoming thin event payload and get the event id
//   - get the full event from /v2/core/events/ with the event id from the thin event
//   - if the full event is a v1.billing.meter.error_report_triggered, use the
//     billing/meter package to retrieve the Billing Meter object associated with the
//     event.
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v82"
)

var apiKey = "{{API_KEY}}"
var webhookSecret = "{{WEBHOOK_SECRET}}"
var client = stripe.NewClient(apiKey)

func main() {
	http.HandleFunc("/webhook", func(w http.ResponseWriter, req *http.Request) {
		const MaxBodyBytes = int64(65536)
		req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
		payload, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		eventNotification, err := client.ParseEventNotification(payload, req.Header.Get("Stripe-Signature"), webhookSecret)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Unmarshal the event data into an appropriate struct depending on its Type
		switch evt := eventNotification.(type) {
		case *stripe.V1BillingMeterErrorReportTriggeredEventNotification:
			fmt.Printf("Got an event related to Meter ID: %s\n", evt.RelatedObject.ID)
			meter, err := evt.FetchRelatedObject(context.TODO())
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error fetching related object: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fmt.Printf("Meter ID %s uses aggregation %s.", meter.ID, meter.DefaultAggregation)

			event, err := evt.FetchEvent(context.TODO())
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error fetching event: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fmt.Printf("Event data reason: %+v\n", event.Data.Reason)

		case *stripe.UnknownEventNotification:
			fmt.Printf("Received an event the SDK doesn't have types for: %s\n", evt.Type)
			fmt.Printf("Event ID: %s; has related object? %t\n", evt.ID, evt.RelatedObject != nil)
		default:
			fmt.Fprintf(os.Stderr, "Purposefully skipping the handling of event w/ type: %s\n", evt.GetEventNotification().Type)
		}

		w.WriteHeader(http.StatusOK)
	})

	err := http.ListenAndServe(":4242", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
