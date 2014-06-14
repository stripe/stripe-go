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

	sub := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	target, err := c.Subs.Create(sub)

	if err != nil {
		t.Error(err)
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

	sub := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	subscription, _ := c.Subs.Create(sub)
	target, err := c.Subs.Get(subscription.Id, &SubParams{Customer: cust.Id})

	if err != nil {
		t.Error(err)
	}

	if target.Id != subscription.Id {
		t.Errorf("Subscription id %q does not match expected id %q\n", target.Id, subscription.Id)
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

	sub := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	subscription, _ := c.Subs.Create(sub)
	err := c.Subs.Cancel(subscription.Id, &SubParams{Customer: cust.Id})

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

	sub := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	subscription, _ := c.Subs.Create(sub)
	updatedSub := &SubParams{
		Customer:  cust.Id,
		NoProrate: true,
		Quantity:  13,
	}

	target, err := c.Subs.Update(subscription.Id, updatedSub)

	if err != nil {
		t.Error(err)
	}

	if target.Quantity != updatedSub.Quantity {
		t.Errorf("Quantity %v does not match expected quantity $v\n", target.Quantity, updatedSub.Quantity)
	}

	c.Customers.Delete(cust.Id)
	c.Plans.Delete("test")
}

func TestSubscriptionDiscount(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	coupon := &CouponParams{
		Duration: Forever,
		Id:       "sub_coupon",
		Percent:  99,
	}
	c.Coupons.Create(coupon)

	customer := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		Coupon: "sub_coupon",
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

	sub := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
		Coupon:   "sub_coupon",
	}

	target, err := c.Subs.Create(sub)

	if err != nil {
		t.Error(err)
	}

	if target.Discount == nil {
		t.Errorf("Discount not found, but one was expected\n")
	}

	if target.Discount.Coupon == nil {
		t.Errorf("Coupon not found, but one was expected\n")
	}

	if target.Discount.Coupon.Id != sub.Coupon {
		t.Errorf("Coupon id %q does not match expected id %q\n", target.Discount.Coupon.Id, sub.Coupon)
	}

	err = c.Discounts.DeleteSubscription(cust.Id, target.Id)

	if err != nil {
		t.Error(err)
	}

	c.Customers.Delete(cust.Id)
	c.Plans.Delete("test")
	c.Coupons.Delete("sub_coupon")
}

func TestSubscriptionList(t *testing.T) {
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

	sub := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	for i := 0; i < 5; i++ {
		c.Subs.Create(sub)
	}

	target, err := c.Subs.List(&SubListParams{Customer: cust.Id})

	if err != nil {
		t.Error(err)
	}

	if len(target.Values) != 5 {
		t.Errorf("Count %v does not match expected value\n", len(target.Values))
	}

	c.Customers.Delete(cust.Id)
	c.Plans.Delete("test")
}
