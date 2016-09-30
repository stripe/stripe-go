package sub

import (
	"fmt"
	"math/rand"
	"testing"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/currency"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/plan"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func planAndCustomer(t *testing.T) (*stripe.Customer, *stripe.Plan) {
	customerParams := &stripe.CustomerParams{
		Source: &stripe.SourceParams{
			Card: &stripe.CardParams{
				Number: "378282246310005",
				Month:  "06",
				Year:   "20",
			},
		},
	}

	cust, err := customer.New(customerParams)
	if err != nil {
		t.Fatalf("customer creation: %s", err)
	}

	planID := fmt.Sprintf("test-%d", rand.Int63())
	planParams := &stripe.PlanParams{
		Name:     "Test Plan",
		ID:       planID,
		Amount:   99,
		Currency: currency.USD,
		Interval: plan.Month,
	}

	p, err := plan.New(planParams)
	if err != nil {
		t.Fatalf("plan creation: %s", err)
	}
	return cust, p
}

func TestSubscriptionCreateUpdateWithItems(t *testing.T) {
	cust, p := planAndCustomer(t)

	// Create
	subParams := &stripe.SubParams{
		Customer: cust.ID,
		Items: []*stripe.SubItemsParams{
			{
				Plan:     p.ID,
				Quantity: 1,
			},
		},
	}

	target, err := New(subParams)
	if err != nil {
		t.Error(err)
	}

	if len(target.Items.Values) != 1 {
		t.Fatalf("Items should be length 1, not %d\n", len(target.Items.Values))
	}

	item := target.Items.Values[0]
	if item.Plan.ID != p.ID {
		t.Errorf("Plan %v does not match expected plan %v\n", item.Plan.ID, p.ID)
	}

	if item.Quantity != 1 {
		t.Errorf("Quantity %v does not match expected quantity %v\n", item.Quantity, 1)
	}

	// Update
	subParams = &stripe.SubParams{
		Items: []*stripe.SubItemsParams{
			{
				ID:       item.ID,
				Quantity: 2,
			},
		},
	}

	target, err = Update(target.ID, subParams)
	if err != nil {
		t.Error(err)
	}

	if len(target.Items.Values) != 1 {
		t.Fatalf("Items should be length 1, not %d\n", len(target.Items.Values))
	}

	item = target.Items.Values[0]
	if item.Quantity != 2 {
		t.Errorf("Quantity %v does not match expected quantity %v\n", item.Quantity, 2)
	}

	// Update with zero
	subParams = &stripe.SubParams{
		Items: []*stripe.SubItemsParams{
			{
				ID:           item.ID,
				QuantityZero: true,
			},
		},
	}

	target, err = Update(target.ID, subParams)
	if err != nil {
		t.Error(err)
	}

	item = target.Items.Values[0]
	if item.Quantity != 0 {
		t.Errorf("Quantity %v does not match expected quantity %v\n", item.Quantity, 0)
	}
}

func TestSubscriptionCreateZeroQuantityWithItems(t *testing.T) {
	cust, p := planAndCustomer(t)

	subParams := &stripe.SubParams{
		Customer: cust.ID,
		Items: []*stripe.SubItemsParams{
			{
				Plan:         p.ID,
				QuantityZero: true,
			},
		},
	}

	target, err := New(subParams)
	if err != nil {
		t.Error(err)
	}

	if len(target.Items.Values) != 1 {
		t.Fatalf("Items should be length 1, not %d\n", len(target.Items.Values))
	}

	item := target.Items.Values[0]
	if item.Plan.ID != p.ID {
		t.Errorf("Plan %v does not match expected plan %v\n", item.Plan.ID, p.ID)
	}

	if item.Quantity != 0 {
		t.Errorf("Quantity %v does not match expected quantity %v\n", item.Quantity, 0)
	}
}
