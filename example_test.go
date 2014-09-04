package stripe

import (
	"fmt"
	"log"
	"net/url"

	"github.com/stripe/stripe"
)

type Post struct{}

func (pa *Post) Call(method, path, token string, body *url.Values, v interface{}) error {
	fmt.Println("ch_example_id")
	return nil
}

func ExamplePost() {
	client := &stripe.Client{}
	client.Init("key", nil, &Post{})

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

	charge, err := client.Charges.Create(params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", charge.Id)
	// Output:
	// ch_example_id
}

type Get struct{}

func (pa *Get) Call(method, path, token string, body *url.Values, v interface{}) error {
	fmt.Println("ch_example_id")
	return nil
}

func ExampleGet() {
	client := &stripe.Client{}
	client.Init("key", nil, &Get{})

	charge, err := client.Charges.Get("ch_example_id")

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", charge.Id)
	// Output:
	// ch_example_id
}

type Update struct{}

func (pa *Update) Call(method, path, token string, body *url.Values, v interface{}) error {
	fmt.Println("updated description")
	return nil
}

func ExampleUpdate() {
	client := &stripe.Client{}
	client.Init("key", nil, &Update{})

	params := &stripe.ChargeParams{
		Desc: "updated description",
	}

	charge, err := client.Charges.Update("ch_example_id", params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", charge.Desc)
	// Output:
	// updated description
}

type Delete struct{}

func (pa *Delete) Call(method, path, token string, body *url.Values, v interface{}) error {
	return nil
}

func ExampleDelete() {
	client := &stripe.Client{}
	client.Init("key", nil, &Delete{})

	err := client.Customers.Delete("acct_example_id")

	if err != nil {
		log.Fatal(err)
	}

	// Output:
}

type List struct{}

func (pa *List) Call(method, path, token string, body *url.Values, v interface{}) error {
	fmt.Println("7")
	return nil
}

func ExampleList() {
	client := &stripe.Client{}
	client.Init("key", nil, &List{})

	params := &stripe.ChargeListParams{}
	params.Filters.AddFilter("limit", "", "7")

	charges, err := client.Charges.List(params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", len(charges.Values))
	// Output:
	// 7
}
