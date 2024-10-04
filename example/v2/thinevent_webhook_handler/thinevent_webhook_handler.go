package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v80"
	billingMeters "github.com/stripe/stripe-go/v80/billing/meter"
	"github.com/stripe/stripe-go/v80/rawrequest"
	webhook "github.com/stripe/stripe-go/v80/webhook"
)

// TODO: rewrite with map[string]interface{}, no types

/**
 * Object containing the reference to API resource relevant to the event.
 */
type ThinEventRelatedObject struct {
	/**
	 * Unique identifier for the object relevant to the event.
	 */
	ID string `json:"id"`

	/**
	 * Type of the object relevant to the event.
	 */
	Type string `json:"type"`

	/**
	 * URL to retrieve the resource.
	 */
	URL string `json:"url"`
}

type ThinEvent struct {
	/**
	 * Unique identifier for the event.
	 */
	ID string `json:"id"`

	/**
	 * String representing the object's type. Objects of the same type share the same value of the object field.
	 */
	Object string `json:"object"`

	/**
	 * Authentication context needed to fetch the event or related object.
	 */
	Context *string `json:"context,omitempty"`

	/**
	 * Time at which the object was created.
	 */
	Created string `json:"created"`

	/**
	 * Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	 */
	Livemode bool `json:"livemode"`

	/**
	 * Reason for the event.
	 */
	// reason: Event.Reason | null;

	/**
	 * The type of the event.
	 */
	Type string `json:"type"`

	/**
	 * Object containing the reference to API resource relevant to the event.
	 */
	RelatedObject *ThinEventRelatedObject `json:"related_object,omitempty"`
}

type V2Event struct {
	/**
	 * Unique identifier for the event.
	 */
	ID string `json:"id"`

	/**
	 * String representing the object's type. Objects of the same type share the same value of the object field.
	 */
	Object string `json:"object"`

	/**
	 * The type of the event.
	 */
	Type string `json:"type"`

	/**
	 * Event data for this event.  May be null/empty for some event types
	 */
	Data *map[string]interface{} `json:"data,omitempty"`

	/**
	 * Object containing the reference to API resource relevant to the event.  May be null/empty for some event types.
	 */
	RelatedObject *ThinEventRelatedObject `json:"related_object,omitempty"`
}

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
		payload, err := io.ReadAll(req.Body)
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

		thinEvent := ThinEvent{}

		if err := json.Unmarshal(payload, &thinEvent); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to parse thin event body json: %v\n", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		event := V2Event{}

		resp, err := client.RawRequest(http.MethodGet, "/v2/core/events/"+thinEvent.ID, "", nil)
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
		switch event.Type {
		case "v1.billing.meter.error_report_triggered":
			meter, err := billingMeters.Get(event.RelatedObject.ID, nil)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to get related meter object: %v\n", err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			meterID := meter.ID
			fmt.Printf("Success! %s\n", meterID)
			// Verify we can see event data
			fmt.Println(fmt.Sprint(event.Data))
		default:
			fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
		}

		w.WriteHeader(http.StatusOK)
	})
	err = http.ListenAndServe(":4242", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
