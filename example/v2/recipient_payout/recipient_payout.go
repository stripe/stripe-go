package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v84"
	"github.com/stripe/stripe-go/v84/rawrequest"
)

func main() {

	apiKey := "{{API_KEY}}"

	b, err := stripe.GetRawRequestBackend(stripe.APIBackend)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client := rawrequest.Client{B: b, Key: apiKey}

	params := map[string]interface{}{
		"from": map[string]interface{}{
			"financial_account": "{{FINANCIAL_ACCOUNT_ID}}",
			"balance_type":      "storage",
		},
		"to": map[string]interface{}{
			"recipient":   "{{ACCOUNT_ID}}",
			"destination": "{{US_BANK_ACCOUNT_ID}}",
		},
		"money_movement_amounts": map[string]interface{}{
			"source": map[string]interface{}{
				"currency": "usd",
				"value":    50,
			},
		},
		"recipient_notification": map[string]interface{}{
			"setting": "configured",
		},
	}

	contentBytes, err := json.Marshal(params)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	content := string(contentBytes)

	// Send a POST request to create an outbound payment
	resp, err := client.RawRequest(http.MethodPost, "/v2/outbound_payments", content, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var newPayment map[string]interface{}
	if err := json.Unmarshal(resp.RawJSON, &newPayment); err != nil {
		fmt.Println(err)
		os.Exit(1)
		return
	}
	newPaymentID := newPayment["id"].(string)

	// Send a GET request to retrieve a payment
	response, err := client.RawRequest(http.MethodGet, "/v2/outbound_payments/"+newPaymentID, "", nil)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	var payment map[string]interface{}
	json.Unmarshal(response.RawJSON, &payment)
	fmt.Println(payment)
}
