package stripe_test

import (
	"log"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/invoice"
	"github.com/stripe/stripe-go/plan"
)

func ExampleCharge_new() {
	stripe.Key = "sk_key"

	params := &stripe.ChargeParams{
		Amount:   1000,
		Currency: stripe.USD,
		Card: &stripe.CardParams{
			Name:   "Go Stripe",
			Number: "4242424242424242",
			Month:  "10",
			Year:   "20",
		},
	}

	ch, err := charge.New(params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", ch.Id)
}

func ExampleCharge_get() {
	stripe.Key = "sk_key"

	params := &stripe.ChargeParams{}
	params.Expand("customer")
	params.Expand("balance_transaction")

	ch, err := charge.Get("ch_example_id", params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", ch.Id)
}

func ExampleInvoice_update() {
	stripe.Key = "sk_key"

	params := &stripe.InvoiceParams{
		Desc: "updated description",
	}

	inv, err := invoice.Update("sub_example_id", params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", inv.Desc)
}

func ExampleCustomer_delete() {
	stripe.Key = "sk_key"

	err := customer.Del("acct_example_id")

	if err != nil {
		log.Fatal(err)
	}
}

func ExamplePlan_list() {
	stripe.Key = "sk_key"

	params := &stripe.PlanListParams{}
	params.Filters.AddFilter("limit", "", "3")
	params.Single = true

	i := plan.List(params)
	for !i.Stop() {
		target, err := i.Next()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("%v ", target.Name)
	}
}
