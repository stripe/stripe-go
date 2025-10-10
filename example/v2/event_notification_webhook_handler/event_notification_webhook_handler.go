package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/stripe/stripe-go/v81"
	"github.com/stripe/stripe-go/v81/webhookendpoint"
)

func main() {
	// This is your Stripe CLI webhook secret for testing your endpoint locally.
	endpointSecret := os.Getenv("STRIPE_WEBHOOK_SECRET")

	http.HandleFunc("/webhook", func(w http.ResponseWriter, req *http.Request) {
		const MaxBodyBytes = int64(65536)
		req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)
		body, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
			w.WriteHeader(http.StatusServiceUnavailable)
			return
		}

		event, err := webhookendpoint.ConstructEventWithOptions(body, req.Header.Get("Stripe-Signature"), endpointSecret, webhookendpoint.ConstructEventOptions{
			IgnoreAPIVersionMismatch: true,
		})

		if err != nil {
			fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
			w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
			return
		}

		// Unmarshal the event data into an appropriate struct depending on its Type
		switch event.Type {
		case "v1.billing.meter.error_report_triggered":
			var meterErrorReport stripe.V1BillingMeterErrorReportTriggeredEvent
			err := json.Unmarshal(body, &meterErrorReport)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			log.Printf("Received event: %v", meterErrorReport)
		default:
			fmt.Fprintf(os.Stderr, "Unhandled event type: %s\n", event.Type)
		}

		w.WriteHeader(http.StatusOK)
	})

	// For production use, you should use proper TLS certificates signed by a trusted CA.
	// For local testing, you can generate self-signed certificates with:
	//   openssl req -x509 -newkey rsa:4096 -keyout server.key -out server.crt -days 365 -nodes -subj '/CN=localhost'
	// 
	// Note: Self-signed certificates should only be used for development/testing.
	log.Fatal(http.ListenAndServeTLS(":4242", "server.crt", "server.key", nil))
}
