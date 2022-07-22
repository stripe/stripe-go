package stripe_test

import (
	"log"

	stripe "github.com/stripe/stripe-go/v72"
	"github.com/stripe/stripe-go/v72/charge"
	"github.com/stripe/stripe-go/v72/customer"
	"github.com/stripe/stripe-go/v72/invoice"
	"github.com/stripe/stripe-go/v72/plan"
)

func ExampleCharge_new() {
	stripe.Key = "sk_key"

	params := &stripe.ChargeParams{
		Amount:   stripe.Int64(1000),
		Currency: stripe.String(string(stripe.CurrencyUSD)),
		Source:   stripe.String("tok_visa"),
	}
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
