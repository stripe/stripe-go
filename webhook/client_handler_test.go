package webhook

import (
	"fmt"
	"net/http"
	"github.com/stripe/stripe-go/webhooks"
)

func Example() {
	http.HandleFunc("/webhook", func (w http.ResponseWriter, req *http.Request) {
		// Pass the request body & Stripe-Signature header to ValidateEvent, along with the webhook signing key
    event, err := webhooks.ValidateEvent(ioutil.ReadAll(req.body), req.Header.Get("Stripe-Signature"), "whsec_DaLRHCRs35vEXqOE8uTEAXGLGUOnyaFf")

    if err != nil {
			w.WriteHeader(http.StatusBadRequest) // Return a 400 error on a bad signature
			fmt.Fprintf(w, "%v", err)
			return
    }

    defer req.Body.Close()
    fmt.Fprintf(w, "Received signed event: %t", event)
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
