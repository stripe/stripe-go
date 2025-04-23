// thinevent_webhook_handler.go - receive and process thin events like the
// v1.billing.meter.error_report_triggered event.
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
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/max-cape/stripe-go-test"
	billingMeters "github.com/max-cape/stripe-go-test/billing/meter"
	"github.com/max-cape/stripe-go-test/rawrequest"
	webhook "github.com/max-cape/stripe-go-test/webhook"
)

var apiKey = "{{API_KEY}}"
var webhookSecret = "{{WEBHOOK_SECRET}}"

func main() {
	b, err := stripe.GetRawRequestBackend(stripe.APIBackend)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client := rawrequest.Client{B: b, Key: apiKey}

	http.HandleFunc("/webhook", func(w http.ResponseWriter, req *http.Request) {
		const MaxBodyBytes = int64(65536)
		req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
		payload, err := ioutil.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = webhook.ValidatePayload(payload, req.Header.Get("Stripe-Signature"), webhookSecret)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var thinEvent map[string]interface{}

		if err := json.Unmarshal(payload, &thinEvent); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse thin event body json: %v\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		eventID := thinEvent["id"].(string)

		var event map[string]interface{}

		resp, err := client.RawRequest(http.MethodGet, "/v2/core/events/"+eventID, "", nil)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to get pull event: %v\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if err := json.Unmarshal(resp.RawJSON, &event); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse pull event body json: %v\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Unmarshal the event data into an appropriate struct depending on its Type
		switch t := event["type"].(string); t {
		case "v1.billing.meter.error_report_triggered":
			relatedObject := event["related_object"].(map[string]interface{})
			meter, err := billingMeters.Get(relatedObject["id"].(string), nil)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to get related meter object: %v\n", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			meterID := meter.ID
			fmt.Printf("Success! %s\n", meterID)
			// Verify we can see event data
			fmt.Println(fmt.Sprint(event["data"]))
		default:
			fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", t)
		}

		w.WriteHeader(http.StatusOK)
	})
	err = http.ListenAndServe(":4242", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
