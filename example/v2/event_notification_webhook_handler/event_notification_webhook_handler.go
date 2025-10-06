// event_notification_webhook_handler.go - receive and process event notifications like the
// "v1.billing.meter.error_report_triggered" event.
//
// This example uses the rawrequests module to make calls to /v2 APIs.
//
// In this example, we:
//   - parse the incoming event notification payload and get the event id
//   - get the full event from /v2/core/events/ with the event id from the event notification
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

	"github.com/stripe/stripe-go/v83"
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
			// there's basic info about the related object in the notification
			fmt.Printf("Meter w/ id %s had a problem\n", evt.RelatedObject.ID)

			// or you can fetch the full object form the API for more details
			meter, err := evt.FetchRelatedObject(context.TODO())
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error fetching related object: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			sb := fmt.Sprintf("Meter %s (%s) had a problem", meter.DisplayName, meter.ID)
			fmt.Println(sb)

			// And you can always fetch the full event:
			event, err := evt.FetchEvent(context.TODO())
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error fetching event: %v\n", err)
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			fmt.Printf("More info: %s\n", event.Data.DeveloperMessageSummary)
		case *stripe.UnknownEventNotification:
			// Events that were introduced after this SDK version release are
			// represented as `UnknownEventNotification`s.
			// They're valid, the SDK just doesn't have corresponding classes for them.
			// You must match on the "type" property instead.
			switch evt.Type {
			case "some.new.event":
				// you can still `.FetchEvent()` and `.FetchRelatedObject()`, but the latter may
				// return `nil` if that event type doesn't have a related object.
				return
			}

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
