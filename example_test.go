package stripe_test

import (
	"fmt"
	"log"
	"net/url"

	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/invoice"
	"github.com/stripe/stripe-go/invoiceitem"
)

type Post struct{}

func (self *Post) Call(method, path, token string, body *url.Values, v interface{}) error {
	fmt.Println("ch_example_id")
	return nil
}

func ExampleCharge_post() {
	stripe.Key = "sk_key"
	stripe.SetBackend(&Post{}) // mocking backend for example

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

	ch, err := charge.Create(params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", ch.Id)
	// Output:
	// ch_example_id
}

type Get struct{}

func (self *Get) Call(method, path, token string, body *url.Values, v interface{}) error {
	fmt.Println("ch_example_id")
	return nil
}

func ExampleCharge_get() {
	stripe.Key = "sk_key"
	stripe.SetBackend(&Get{}) // mocking backend for example

	params := &stripe.ChargeParams{}
	params.Expand("customer")
	params.Expand("balance_transaction")

	ch, err := charge.Get("ch_example_id", params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", ch.Id)
	// Output:
	// ch_example_id
}

type Update struct{}

func (self *Update) Call(method, path, token string, body *url.Values, v interface{}) error {
	fmt.Println("updated description")
	return nil
}

func ExampleInvoice_update() {
	stripe.Key = "sk_key"
	stripe.SetBackend(&Update{}) // mocking backend for example

	params := &stripe.InvoiceParams{
		Desc: "updated description",
	}

	inv, err := invoice.Update("sub_example_id", params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", inv.Desc)
	// Output:
	// updated description
}

type Delete struct{}

func (self *Delete) Call(method, path, token string, body *url.Values, v interface{}) error {
	return nil
}

func ExampleCustomer_delete() {
	stripe.Key = "sk_key"
	stripe.SetBackend(&Delete{}) // mocking backend for example

	err := customer.Delete("acct_example_id")

	if err != nil {
		log.Fatal(err)
	}

	// Output:
}

type List struct{}

func (self *List) Call(method, path, token string, body *url.Values, v interface{}) error {
	fmt.Println("7")
	return nil
}

func ExampleInvoiceItem_list() {
	stripe.Key = "sk_key"
	stripe.SetBackend(&List{}) // mocking backend for example

	params := &stripe.InvoiceItemListParams{}
	params.Filters.AddFilter("limit", "", "7")

	items, err := invoiceitem.List(params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", len(items.Values))
	// Output:
	// 7
}
