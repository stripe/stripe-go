package stripe_test

import (
	"log"

	stripe "github.com/stripe/stripe-go/v83"
	"github.com/stripe/stripe-go/v83/charge"
	"github.com/stripe/stripe-go/v83/customer"
	"github.com/stripe/stripe-go/v83/invoice"
	"github.com/stripe/stripe-go/v83/plan"
)

func ExampleCharge_new() {
	stripe.Key = "sk_key"

	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(1000),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
	}
	params.SetSource("tok_visa")
	params.AddMetadata("key", "value")

	ch, err := charge.New(params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", ch.ID)
}

func ExampleCharge_get() {
	stripe.Key = "sk_key"

	params := &stripe.ChargeParams{}
	params.AddExpand("customer")
	params.AddExpand("application")
	params.AddExpand("balance_transaction")

	ch, err := charge.Get("ch_example_id", params)

	if err != nil {
		log.Fatal(err)
	}

	if ch.Application != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", ch.ID)
}

func ExampleInvoice_update() {
	stripe.Key = "sk_key"

	params := &stripe.InvoiceParams{
		Description: stripe.String("updated description"),
	}

	inv, err := invoice.Update("sub_example_id", params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", inv.Description)
}

func ExampleCustomer_delete() {
	stripe.Key = "sk_key"

	customerDel, err := customer.Del("cus_example_id", nil)

	if err != nil {
		log.Fatal(err)
	}

	if !customerDel.Deleted {
		log.Fatal("Customer doesn't appear deleted while it should be")
	}
}

func ExamplePlan_list() {
	stripe.Key = "sk_key"

	params := &stripe.PlanListParams{}
	params.Filters.AddFilter("limit", "", "3")
	params.Single = true

	it := plan.List(params)
	for it.Next() {
		log.Printf("%v ", it.Plan().Nickname)
	}
	if err := it.Err(); err != nil {
		log.Fatal(err)
	}
}

// This example demonstrates how to enable OpenTelemetry tracing for Stripe API calls.
func ExampleSetTracerProvider() {
	// Get the global tracer provider from your OpenTelemetry setup
	// (You would typically initialize this in your main() or init() function)
	// For example:
	//   import "go.opentelemetry.io/otel"
	//   tp := otel.GetTracerProvider()

	// Enable OpenTelemetry tracing for all Stripe API calls
	// stripe.SetTracerProvider(tp)

	// Now all Stripe API calls will create spans with proper attributes
	// including http.method, url.template, stripe.request_id, etc.

	// To disable tracing later:
	// stripe.SetTracerProvider(nil)
}

// This example shows how to use FormatURLPathWithTemplate to ensure
// proper URL template tracking in OpenTelemetry spans.
func ExampleFormatURLPathWithTemplate() {
	stripe.Key = "sk_test_example"

	chargeID := "ch_1234567890"

	// Use FormatURLPathWithTemplate instead of FormatURLPath
	// This ensures the OpenTelemetry span gets the template "/v1/charges/{id}"
	// instead of the expanded path "/v1/charges/ch_1234567890"
	params := &stripe.ChargeParams{}

	// The helper function formats the path and sets URLTemplate in params
	path := stripe.FormatURLPathWithTemplate(
		&params.Params,
		"/v1/charges/{id}",
		chargeID,
	)

	log.Printf("Formatted path: %s", path)
	// Output: path is "/v1/charges/ch_1234567890"
	// The resulting OpenTelemetry span will have url.template="/v1/charges/{id}"
}
