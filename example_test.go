package stripe_test

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

func ExamplePostCharge() {
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
