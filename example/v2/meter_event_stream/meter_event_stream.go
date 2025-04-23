// meter_event_stream.go - use the high-throughput meter event stream to report create billing meter events.
//
// This example uses the rawrequests module to make calls to /v2 APIs.
//
// In this example, we:
//   - create a meter event session and store the session's authentication token
//   - define an event with a payload
//   - post the event to /v2/billing/meter_event_stream to create an event stream that reports this event
//
// This example expects a billing meter with an event_name of 'alpaca_ai_tokens'.  If you have
// a different meter event name, you can change it before running this example.
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	stripe "github.com/max-cape/stripe-go-test"
	rawrequest "github.com/max-cape/stripe-go-test/rawrequest"
)

var sessionAuthToken string = ""
var sessionAuthExpiresAt string = ""

func refreshMeterEventSession(client rawrequest.Client) (err error) {
	currentTime := time.Now().Format(time.RFC3339)
	// Check if session is null or expired
	if sessionAuthToken == "" || sessionAuthExpiresAt <= currentTime {
		// Create a new meter event session in case the existing session expired
		rawResp, err := client.RawRequest(http.MethodPost, "/v2/billing/meter_event_session", "", nil)
		if err != nil {
			return err
		}
		if rawResp.StatusCode != 200 {
			return fmt.Errorf(rawResp.Status)
		}

		var resp map[string]interface{}
		err = json.Unmarshal(rawResp.RawJSON, &resp)
		if err != nil {
			return err
		}

		sessionAuthToken = resp["authentication_token"].(string)
		sessionAuthExpiresAt = resp["expires_at"].(string)

		fmt.Println("Meter event session created!")
	}
	return nil
}

func sendMeterEvent(client rawrequest.Client, eventName string, stripeCustomerID string, value string) (err error) {
	// Refresh the meter event session if necessary
	refreshMeterEventSession(client)

	if sessionAuthToken == "" {
		err = fmt.Errorf("Unable to refresh meter event session")
		return
	}

	b, err := stripe.GetRawRequestBackend(stripe.MeterEventsBackend)
	if err != nil {
		return err
	}

	sessionClient := rawrequest.Client{B: b, Key: sessionAuthToken}

	params := map[string]interface{}{
		"events": []interface{}{
			map[string]interface{}{
				"event_name": eventName,
				"payload": map[string]interface{}{
					"stripe_customer_id": stripeCustomerID,
					"value":              value,
				},
			},
		},
	}

	contentBytes, err := json.Marshal(params)
	if err != nil {
		return
	}

	content := string(contentBytes)
	_, err = sessionClient.RawRequest(http.MethodPost, "/v2/billing/meter_event_stream", content, nil)
	return
}

func main() {

	apiKey := "{{API_KEY}}"
	customerID := "{{CUSTOMER_ID}}"

	b, err := stripe.GetRawRequestBackend(stripe.APIBackend)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client := rawrequest.Client{B: b, Key: apiKey}

	err = sendMeterEvent(client, "alpaca_ai_tokens", customerID, "25")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Meter event sent successfully!")
}
