package paymentsource

import (
	"testing"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/card"
	"github.com/stripe/stripe-go/customer"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestSourceNew(t *testing.T) {
	customerParams := &stripe.CustomerParams{}
	cust, err := customer.New(customerParams)

	if err != nil {
		t.Error(err)
	}

	sourceParams := &stripe.CustomerSourceParams{
		Customer: cust.ID,
	}
	sourceParams.SetSource(&stripe.CardParams{
		Number: "4242424242424242",
		Month:  "10",
		Year:   "20",
		CVC:    "1234",
	})

	source, err := New(sourceParams)

	if err != nil {
		t.Error(err)
	}

	target := source.Card

	if target.LastFour != "4242" {
		t.Errorf("Unexpected last four %q for card number %v\n", target.LastFour, sourceParams.Source.Card.Number)
	}

	if target.CVCCheck != card.Pass {
		t.Errorf("CVC check %q does not match expected status\n", target.ZipCheck)
	}

	targetCust, err := customer.Get(cust.ID, nil)

	if err != nil {
		t.Error(err)
	}

	if targetCust.Sources.Count != 1 {
		t.Errorf("Unexpected number of sources %v\n", targetCust.Sources.Count)
	}

	customer.Del(cust.ID)
}

func TestSourceGet(t *testing.T) {
	customerParams := &stripe.CustomerParams{
		Email: "SomethingIdentifiable@gmail.om",
	}
	customerParams.SetSource(&stripe.CardParams{
		Number: "4242424242424242",
		Month:  "06",
		Year:   "20",
	})
	cust, err := customer.New(customerParams)

	if err != nil {
		t.Error(err)
	}

	source, err := Get(cust.DefaultSource.ID, &stripe.CustomerSourceParams{Customer: cust.ID})

	if err != nil {
		t.Error(err)
	}

	target := source.Card

	if target.LastFour != "4242" {
		t.Errorf("Unexpected last four %q for card number %v\n", target.LastFour, customerParams.Source.Card.Number)
	}

	if target.Brand != card.Visa {
		t.Errorf("Card brand %q does not match expected value\n", target.Brand)
	}

	customer.Del(cust.ID)
}

func TestSourceDel(t *testing.T) {
	customerParams := &stripe.CustomerParams{}
	customerParams.SetSource(&stripe.CardParams{
		Number: "378282246310005",
		Month:  "06",
		Year:   "20",
	})

	cust, _ := customer.New(customerParams)

	err := Del(cust.DefaultSource.ID, &stripe.CustomerSourceParams{Customer: cust.ID})
	if err != nil {
		t.Error(err)
	}

	customer.Del(cust.ID)
}

func TestSourceUpdate(t *testing.T) {
	customerParams := &stripe.CustomerParams{}
	customerParams.SetSource(&stripe.CardParams{
		Number: "378282246310005",
		Month:  "06",
		Year:   "20",
		Name:   "Original Name",
	})

	cust, err := customer.New(customerParams)

	if err != nil {
		t.Error(err)
	}

	sourceParams := &stripe.CustomerSourceParams{
		Customer: cust.ID,
	}
	sourceParams.SetSource(&stripe.CardParams{
		Name: "Updated Name",
	})

	source, err := Update(cust.DefaultSource.ID, sourceParams)

	if err != nil {
		t.Error(err)
		return
	}

	target := source.Card

	if target.Name != sourceParams.Source.Card.Name {
		t.Errorf("Card name %q does not match expected name %q\n", target.Name, sourceParams.Source.Card.Name)
	}

	customer.Del(cust.ID)
}

func TestSourceList(t *testing.T) {
	customerParams := &stripe.CustomerParams{}
	customerParams.SetSource(&stripe.CardParams{
		Number: "378282246310005",
		Month:  "06",
		Year:   "20",
	})

	cust, _ := customer.New(customerParams)

	sourceParams := &stripe.CustomerSourceParams{
		Customer: cust.ID,
	}
	sourceParams.SetSource(&stripe.CardParams{
		Number: "4242424242424242",
		Month:  "10",
		Year:   "20",
	})

	New(sourceParams)

	i := List(&stripe.SourceListParams{Customer: cust.ID})
	for i.Next() {
		paymentSource := i.PaymentSource()

		if paymentSource == nil {
			t.Error("No nil values expected")
		}

		if paymentSource.Card == nil {
			t.Error("No nil values expected")
		}

		if i.Meta() == nil {
			t.Error("No metadata returned")
		}
	}
	if err := i.Err(); err != nil {
		t.Error(err)
	}

	customer.Del(cust.ID)
}
