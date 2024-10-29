// thinevent_webhook_handler.go - receive and process thin events like the
// v1.billing.meter.error_report_triggered event.
//
// TODO: change this
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
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	webhook "github.com/stripe/stripe-go/v80/webhook"
)

var apiKey = "{{API_KEY}}"
var webhookSecret = "{{WEBHOOK_SECRET}}"

func main() {

	http.HandleFunc("/webhook", func(w http.ResponseWriter, req *http.Request) {
		const MaxBodyBytes = int64(65536)
		req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
		payload, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		event, err := webhook.ConstructEvent(payload, req.Header.Get("Stripe-Signature"), webhookSecret)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Println(event.Type)
		w.WriteHeader(http.StatusOK)
	})
	err := http.ListenAndServe(":4242", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
