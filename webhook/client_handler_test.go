package webhook_test

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/max-cape/stripe-go-test/webhook"
)

func Example() {
	http.HandleFunc("/webhook", func(w http.ResponseWriter, req *http.Request) {
		// Protects against a malicious client streaming us an endless request
		// body
		const MaxBodyBytes = int64(65536)
		req.Body = http.MaxBytesReader(w, req.Body, MaxBodyBytes)

		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Pass the request body & Stripe-Signature header to ConstructEvent, along with the webhook signing key
		event, err := webhook.ConstructEvent(body, req.Header.Get("Stripe-Signature"), "whsec_DaLRHCRs35vEXqOE8uTEAXGLGUOnyaFf")

		if err != nil {
			w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
			fmt.Fprintf(w, "%v", err)
			return
		}

		fmt.Fprintf(w, "Received signed event: %v", event)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
