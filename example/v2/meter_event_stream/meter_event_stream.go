package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	stripe "github.com/stripe/stripe-go/v80"
	rawrequest "github.com/stripe/stripe-go/v80/rawrequest"
)

type MeterEventSessionResponse struct {
	Object              string // "v2.billing.meter_event_session"
	AuthenticationToken string `json:"authentication_token"`
	ExpiresAt           string `json:"expires_at"`
}

type MeterEventStreamEventPayload struct {
	StripeCustomerID string `json:"stripe_customer_id"`
	Value            string `json:"value"`
}

type MeterEventStreamEventParams struct {
	EventName string                       `json:"event_name"`
	Payload   MeterEventStreamEventPayload `json:"payload"`
}
type MeterEventStreamParams struct {
	Events []MeterEventStreamEventParams `json:"events"`
}

var sessionAuthToken string = ""
var sessionAuthExpiresAt string = ""

func refreshMeterEventSession() (err error) {
	currentTime := time.Now().Format(time.RFC3339)
	// Check if session is null or expired
	if sessionAuthToken == "" || sessionAuthExpiresAt <= currentTime {
		// Create a new meter event session in case the existing session expired
		rawResp, err := rawrequest.Post("/v2/billing/meter_event_session", "", nil)
		if err != nil {
			return err
		}
		if rawResp.StatusCode != 200 {
			return fmt.Errorf(rawResp.Status)
		}

		/* Example using map[string]interface{}
		var resp map[string]interface{}
		err = json.Unmarshal(rawResp.RawJSON, &resp)
		if err != nil {
			return err
		}
		// sessionAuthToken = resp["authentication_token"].(string)
		// sessionAuthExpiresAt = resp["expires_at"].(string)
		*/

		// Example using structs
		var resp MeterEventSessionResponse
		// var resp map[string]interface{}
		err = json.Unmarshal(rawResp.RawJSON, &resp)
		if err != nil {
			return err
		}
		sessionAuthToken = resp.AuthenticationToken
		sessionAuthExpiresAt = resp.ExpiresAt
		fmt.Println("Meter event session created!")
	}
	return nil
}

var oldKey = ""
var meterEventsURL = "https://meter-events.stripe.com"

func prepareMeterEventsCall(authToken string) {
	fmt.Println("Preparing for meter events call")
	backend := stripe.GetBackendWithConfig(
		stripe.APIBackend,
		&stripe.BackendConfig{
			HTTPClient:        nil,
			LeveledLogger:     nil, // Set by GetBackendWithConfiguation when nil
			MaxNetworkRetries: nil, // Set by GetBackendWithConfiguation when nil
			URL:               &meterEventsURL,
		},
	)

	stripe.SetBackend(stripe.APIBackend, backend)
	oldKey = stripe.Key
	stripe.Key = authToken
}

func cleanupMeterEventsCall() {
	fmt.Println("Cleaning up after meter events call")
	stripe.SetBackend(stripe.APIBackend, nil)
	stripe.Key = oldKey
	oldKey = ""
}

func sendMeterEvent(eventName string, stripeCustomerID string, value string) (err error) {
	// Refresh the meter event session if necessary
	refreshMeterEventSession()

	if sessionAuthToken == "" {
		err = fmt.Errorf("Unable to refresh meter event session")
		return
	}

	// Change the APIBackend to point to meter-events.stripe.com
	// to use rawrequest.Post to call a meter event stream API
	prepareMeterEventsCall(sessionAuthToken)
	defer cleanupMeterEventsCall()

	/* Example using map[string]interface{}
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
	*/
	// Example using structs
	params := MeterEventStreamParams{
		Events: []MeterEventStreamEventParams{
			{
				EventName: eventName,
				Payload: MeterEventStreamEventPayload{
					StripeCustomerID: stripeCustomerID,
					Value:            value,
				},
			},
		},
	}

	contentBytes, err := json.Marshal(params)
	if err != nil {
		return
	}
	content := string(contentBytes)
	_, err = rawrequest.Post("/v2/billing/meter_event_stream", content, nil)
	return
}

func main() {

	stripe.Key = "{{API_KEY}}"
	customerID := "{{CUSTOMER_ID}}"

	err := sendMeterEvent("api_requests", customerID, "25")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Meter event sent successfully!")
}
