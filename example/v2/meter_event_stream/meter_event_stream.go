package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	stripe "github.com/stripe/stripe-go/v80"
	rawrequest "github.com/stripe/stripe-go/v80/rawrequest"
)

var sessionAuthToken string = ""
var sessionAuthExpiresAt string = ""

func refreshMeterEventSession(client rawrequest.Client) (err error) {
	currentTime := time.Now().Format(time.RFC3339)
	// Check if session is null or expired
	if sessionAuthToken == "" || sessionAuthExpiresAt <= currentTime {
		// Create a new meter event session in case the existing session expired
		rawResp, err := client.Post("/v2/billing/meter_event_session", "", nil)
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
	_, err = sessionClient.Post("/v2/billing/meter_event_stream", content, nil)
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

	err = sendMeterEvent(client, "api_requests", customerID, "25")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("Meter event sent successfully!")
}
