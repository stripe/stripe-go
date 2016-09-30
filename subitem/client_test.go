package subitem

import (
	"fmt"
	"math/rand"
	"testing"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/currency"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/plan"
	"github.com/stripe/stripe-go/sub"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func createSubItem(t *testing.T) (*stripe.Sub, *stripe.SubItem, func()) {
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
		ID:       planID,
		Name:     "Test Plan",
		Amount:   99,
		Currency: currency.USD,
		Interval: plan.Month,
	}

	p, err := plan.New(planParams)
	if err != nil {
		t.Fatalf("plan creation: %s", err)
	}

	subParams := &stripe.SubParams{
		Customer: cust.ID,
		Items: []*stripe.SubItemsParams{
			{
				Plan:     p.ID,
				Quantity: 1,
			},
		},
	}

	target, err := sub.New(subParams)
	if err != nil {
		t.Fatalf("subscription creation: %s", err)
	}
	if target.Items == nil {
		t.Fatalf("no items for sub %s", target.ID)
	}
	if len(target.Items.Values) != 1 {
		t.Fatalf("missing items: %#v", target)
	}
	return target, target.Items.Values[0], func() {
		sub.Cancel(target.ID, nil)
		plan.Del(p.ID)
		customer.Del(cust.ID)
	}
}

func TestUpdateItem(t *testing.T) {
	_, item, cleanup := createSubItem(t)
	defer cleanup()

	item, err := Update(item.ID, &stripe.SubItemParams{
		Quantity: 10,
	})
	if err != nil {
		t.Errorf("update err: %s", err)
	}
	if item.Quantity != 10 {
		t.Errorf("quantity should be 10, not %d", item.Quantity)
	}
}

func TestGetItem(t *testing.T) {
	_, item, cleanup := createSubItem(t)
	defer cleanup()

	got, err := Get(item.ID, nil)
	if err != nil {
		t.Errorf("get err: %s", err)
	}

	if got.ID != item.ID {
		t.Errorf("incorrect ID %s != %s", got.ID, item.ID)
	}
	if got.Quantity != item.Quantity {
		t.Errorf("incorrect quantity %d != %d", got.Quantity, item.Quantity)
	}
	if got.Created != item.Created {
		t.Errorf("incorrect created %d != %d", got.Created, item.Created)
	}
}

func TestListItems(t *testing.T) {
	sub, _, cleanup := createSubItem(t)
	defer cleanup()

	// Create a new plan
	planID := fmt.Sprintf("test-%d", rand.Int63())
	planParams := &stripe.PlanParams{
		ID:       planID,
		Name:     "Expensive plan",
		Amount:   1000,
		Currency: currency.USD,
		Interval: plan.Month,
	}
	p, err := plan.New(planParams)
	if err != nil {
		t.Fatal(err)
	}
	defer plan.Del(p.ID)

	item, err := New(&stripe.SubItemParams{
		Sub:      sub.ID,
		Plan:     p.ID,
		Quantity: 2,
	})
	if err != nil {
		t.Errorf("new err: %s", err)
	}
	if item.Quantity != 2 {
		t.Errorf("quantity should be 2 after update, not %d", item.Quantity)
	}

	if item.ID != "" {
		item, err := Del(item.ID, nil)
		if err != nil {
			t.Errorf("del err: %s", err)
		}
		if !item.Deleted {
			t.Errorf("item should be deleted %#v", item)
		}
	}
}

func TestCreateAndDelItem(t *testing.T) {
	sub, item, cleanup := createSubItem(t)
	defer cleanup()

	item, err := New(&stripe.SubItemParams{
		Sub:      sub.ID,
		Plan:     item.Plan.ID,
		Quantity: 1,
	})
	if err != nil {
		t.Errorf("create err: %s", err)
	}
	if item.Quantity != 1 {
		t.Errorf("quantity should be 1, not %d", item.Quantity)
	}

	item, err = Del(item.ID, nil)
	if err != nil {
		t.Errorf("create err: %s", err)
	}
	if !item.Deleted {
		t.Error("item should be deleted")
	}
}
