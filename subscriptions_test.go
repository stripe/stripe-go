package stripe

import (
	"testing"
)

func TestSubscriptionCreate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	customer := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := c.Customers.Create(customer)

	plan := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}
	c.Plans.Create(plan)

	sub := &SubscriptionParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	target, err := c.Subs.Create(sub)
	if err != nil {
		t.Error(err)
	}

	if target == nil {
		t.Errorf("No subscription returned\n")
	}

	if target.Plan.Id != sub.Plan {
		t.Errorf("Plan %q does not match expected plan %q\n", target.Plan, sub.Plan)
	}

	if target.Quantity != sub.Quantity {
		t.Errorf("Quantity %v does not match expected quantity %v\n", target.Quantity, sub.Quantity)
	}

	c.Customers.Delete(cust.Id)
	c.Plans.Delete("test")
}

func TestSubscriptionGet(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	customer := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := c.Customers.Create(customer)

	plan := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}
	c.Plans.Create(plan)

	sub := &SubscriptionParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	subscription, _ := c.Subs.Create(sub)
	target, err := c.Subs.Get(subscription.Id, &SubscriptionParams{Customer: cust.Id})

	if err != nil {
		t.Error(err)
	}

	if target == nil {
		t.Errorf("No subscription returned\n")
	}

	c.Customers.Delete(cust.Id)
	c.Plans.Delete("test")
}

func TestSubscriptionCancel(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	customer := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := c.Customers.Create(customer)

	plan := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}
	c.Plans.Create(plan)

	sub := &SubscriptionParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	subscription, _ := c.Subs.Create(sub)
	err := c.Subs.Cancel(subscription.Id, &SubscriptionParams{Customer: cust.Id})

	if err != nil {
		t.Error(err)
	}

	c.Customers.Delete(cust.Id)
	c.Plans.Delete("test")
}

func TestSubscriptionUpdate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	customer := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := c.Customers.Create(customer)

	plan := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}
	c.Plans.Create(plan)

	sub := &SubscriptionParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	subscription, _ := c.Subs.Create(sub)
	updatedSub := &SubscriptionParams{
		Customer:  cust.Id,
		NoProrate: true,
		Quantity:  13,
	}

	target, err := c.Subs.Update(subscription.Id, updatedSub)

	if err != nil {
		t.Error(err)
	}

	if target == nil {
		t.Errorf("No subscription returned\n")
	}

	if target.Quantity != updatedSub.Quantity {
		t.Errorf("Quantity %v does not match expected quantity $v\n", target.Quantity, updatedSub.Quantity)
	}

	c.Customers.Delete(cust.Id)
	c.Plans.Delete("test")
}
