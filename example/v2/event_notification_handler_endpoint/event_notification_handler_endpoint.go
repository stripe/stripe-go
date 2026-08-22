// event_notification_handler_endpoint.go - receive and process event notifications like the
// "v1.billing.meter.error_report_triggered" event.
//
// This example uses the rawrequests module to make calls to /v2 APIs.
//
// In this example, we:
//   - write a fallback callback to handle unrecognized event notifications
//   - create a stripe.Client called client
//   - Initialize an EventNotificationHandler with the client, webhook secret, and fallback callback
//   - register a PreHandle hook that deduplicates events by ID before any callback runs
//   - register a specific handler for the "v1.billing.meter.error_report_triggered" event notification type
//   - use handler.Handle() to process the received notification webhook body

package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"

	"github.com/stripe/stripe-go/v86"
)

var apiKey = "{{API_KEY}}"
var webhookSecret = "{{WEBHOOK_SECRET}}"
var client = stripe.NewClient(apiKey)
var handler = stripe.NewEventNotificationHandler(client, webhookSecret, func(context context.Context, notif stripe.EventNotificationContainer, client *stripe.Client, details stripe.UnhandledNotificationDetails) error {
	fmt.Printf("Received unhandled event with type %s", notif.GetEventNotification().Type)
	return nil
})

// seenEventIDs tracks event IDs we've already processed so that redelivered webhooks
// (e.g. after a slow response causes Stripe to retry) don't get handled twice.
var (
	seenEventIDsMu sync.Mutex
	seenEventIDs   = map[string]bool{}
)

func init() {
	// PreHandle runs after Handle() parses the payload but before any callback (registered
	// or fallback) fires. Returning false here skips every callback for this event, which is
	// exactly what we want for a duplicate delivery.
	handler.PreHandle(func(ctx context.Context, notif stripe.EventNotificationContainer, client *stripe.Client) (bool, error) {
		eventID := notif.GetEventNotification().ID

		seenEventIDsMu.Lock()
		defer seenEventIDsMu.Unlock()

		if seenEventIDs[eventID] {
			fmt.Printf("Skipping duplicate delivery of event %s\n", eventID)
			return false, nil
		}
		seenEventIDs[eventID] = true
		return true, nil
	})
}

// Handles events delivered through a channel that has already authenticated them, such as
// AWS EventBridge or Azure Event Grid. Those payloads carry no Stripe-Signature header, so
// this handler skips verification. Callbacks are registered separately from the one above.
var unverifiedHandler = stripe.NewEventNotificationHandlerWithoutVerification(client, func(context context.Context, notif stripe.EventNotificationContainer, client *stripe.Client, details stripe.UnhandledNotificationDetails) error {
	fmt.Printf("Received unhandled event with type %s", notif.GetEventNotification().Type)
	return nil
})

func handleMeterErrors(ctx context.Context, notif *stripe.V1BillingMeterErrorReportTriggeredEventNotification, client *stripe.Client) error {
	meter, err := notif.FetchRelatedObject(ctx)
	if err != nil {
		return fmt.Errorf("error fetching related object: %v", err)
	}
	fmt.Printf("Meter %s (%s) had a problem\n", meter.DisplayName, meter.ID)

	return nil
}

func main() {
	handler.OnV1BillingMeterErrorReportTriggered(handleMeterErrors)

	http.HandleFunc("/webhook", func(w http.ResponseWriter, req *http.Request) {
		const MaxBodyBytes = int64(65536)
		req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
		payload, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = handler.Handle(req.Context(), payload, req.Header.Get("Stripe-Signature"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error handling event notification: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	unverifiedHandler.OnV1BillingMeterErrorReportTriggered(handleMeterErrors)

	http.HandleFunc("/webhook-from-cloud-provider", func(w http.ResponseWriter, req *http.Request) {
		const MaxBodyBytes = int64(65536)
		req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
		payload, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Handle takes only the body here; there's no signature to check
		err = unverifiedHandler.Handle(req.Context(), payload)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error handling event notification: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	err := http.ListenAndServe(":4242", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
