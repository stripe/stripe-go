package stripe_test

import (
	"log"
	"time"
)

type PostApi struct
postApi := &PostApi{}

func (pa *PostApi) Call(method, path, token string, body *url.Values, v interface{}) error {
	v := &Charge{
		Id: "ch_example_id",
		Card: &Card {
			Id: "card_example_id",
			Month: 10,
			Year: 20,
			Funding: CreditFunding,
			LastFour: "1234",
			Brand: Visa,
			City: "San Francisco",
			Country: "US",
			State: "CA",
			CardCountry: "US",
			Name: "Go Stripe"
		},
		Created: time.Now().UnixNano(),
		Currency: "USD"
	}

	return nil
}

client := &Client{}
client.Init("key", nil, postApi)

func ExamplePostCharge() {
	params := &ChargeParams{
		Amount:   1000,
		Currency: USD,
		Card: &CardParams{
			Name:   "Go Stripe",
			Number: "4242424242424242",
			Month:  "10",
			Year:   "20",
		},
	}
	
	charge, err := client.ChargeClient.Create(params)

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%v\n", charge)
}












