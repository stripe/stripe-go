package sub

import (
	"testing"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/coupon"
	"github.com/stripe/stripe-go/currency"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/discount"
	"github.com/stripe/stripe-go/plan"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestSubscriptionNew(t *testing.T) {
	customerParams := &stripe.CustomerParams{
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	planParams := &stripe.PlanParams{
		ID:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: currency.USD,
		Interval: plan.Month,
	}

	plan.New(planParams)

	subParams := &stripe.SubParams{
		Customer: cust.ID,
		Plan:     "test",
		Quantity: 10,
	}

	target, err := New(subParams)

	if err != nil {
		t.Error(err)
	}

	if target.Plan.ID != subParams.Plan {
		t.Errorf("Plan %v does not match expected plan %v\n", target.Plan, subParams.Plan)
	}

	if target.Quantity != subParams.Quantity {
		t.Errorf("Quantity %v does not match expected quantity %v\n", target.Quantity, subParams.Quantity)
	}

	customer.Del(cust.ID)
	plan.Del("test")
}

func TestSubscriptionZeroQuantity(t *testing.T) {
	customerParams := &stripe.CustomerParams{
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	planParams := &stripe.PlanParams{
		ID:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: currency.USD,
		Interval: plan.Month,
	}

	plan.New(planParams)

	subParams := &stripe.SubParams{
		Customer:     cust.ID,
		Plan:         "test",
		QuantityZero: true,
	}

	target, err := New(subParams)

	if err != nil {
		t.Error(err)
	}

	if target.Plan.ID != subParams.Plan {
		t.Errorf("Plan %v does not match expected plan %v\n", target.Plan, subParams.Plan)
	}

	if target.Quantity != 0 {
		t.Errorf("Quantity %v does not match expected quantity %v\n", target.Quantity, 0)
	}

	customer.Del(cust.ID)
	plan.Del("test")
}

func TestSubscriptionGet(t *testing.T) {
	customerParams := &stripe.CustomerParams{
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	planParams := &stripe.PlanParams{
		ID:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: currency.USD,
		Interval: plan.Month,
	}

	plan.New(planParams)

	subParams := &stripe.SubParams{
		Customer: cust.ID,
		Plan:     "test",
		Quantity: 10,
	}

	subscription, _ := New(subParams)
	target, err := Get(subscription.ID, &stripe.SubParams{Customer: cust.ID})

	if err != nil {
		t.Error(err)
	}

	if target.ID != subscription.ID {
		t.Errorf("Subscription id %q does not match expected id %q\n", target.ID, subscription.ID)
	}

	customer.Del(cust.ID)
	plan.Del("test")
}

func TestSubscriptionCancel(t *testing.T) {
	customerParams := &stripe.CustomerParams{
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	planParams := &stripe.PlanParams{
		ID:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: currency.USD,
		Interval: plan.Month,
	}

	plan.New(planParams)

	subParams := &stripe.SubParams{
		Customer: cust.ID,
		Plan:     "test",
		Quantity: 10,
	}

	subscription, _ := New(subParams)
	err := Cancel(subscription.ID, &stripe.SubParams{Customer: cust.ID})

	if err != nil {
		t.Error(err)
	}

	customer.Del(cust.ID)
	plan.Del("test")
}

func TestSubscriptionUpdate(t *testing.T) {
	customerParams := &stripe.CustomerParams{
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	planParams := &stripe.PlanParams{
		ID:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: currency.USD,
		Interval: plan.Month,
	}

	plan.New(planParams)

	subParams := &stripe.SubParams{
		Customer: cust.ID,
		Plan:     "test",
		Quantity: 10,
	}

	subscription, _ := New(subParams)
	updatedSub := &stripe.SubParams{
		Customer:  cust.ID,
		NoProrate: true,
		Quantity:  13,
	}

	target, err := Update(subscription.ID, updatedSub)

	if err != nil {
		t.Error(err)
	}

	if target.Quantity != updatedSub.Quantity {
		t.Errorf("Quantity %v does not match expected quantity %v\n", target.Quantity, updatedSub.Quantity)
	}

	customer.Del(cust.ID)
	plan.Del("test")
}

func TestSubscriptionDiscount(t *testing.T) {
	couponParams := &stripe.CouponParams{
		Duration: coupon.Forever,
		ID:       "sub_coupon",
		Percent:  99,
	}

	coupon.New(couponParams)

	customerParams := &stripe.CustomerParams{
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		Coupon: "sub_coupon",
	}

	cust, _ := customer.New(customerParams)

	planParams := &stripe.PlanParams{
		ID:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: currency.USD,
		Interval: plan.Month,
	}

	plan.New(planParams)

	subParams := &stripe.SubParams{
		Customer: cust.ID,
		Plan:     "test",
		Quantity: 10,
		Coupon:   "sub_coupon",
	}

	target, err := New(subParams)

	if err != nil {
		t.Error(err)
	}

	if target.Discount == nil {
		t.Errorf("Discount not found, but one was expected\n")
	}

	if target.Discount.Coupon == nil {
		t.Errorf("Coupon not found, but one was expected\n")
	}

	if target.Discount.Coupon.ID != subParams.Coupon {
		t.Errorf("Coupon id %q does not match expected id %q\n", target.Discount.Coupon.ID, subParams.Coupon)
	}

	err = discount.DelSub(cust.ID, target.ID)

	if err != nil {
		t.Error(err)
	}

	customer.Del(cust.ID)
	plan.Del("test")
	coupon.Del("sub_coupon")
}

func TestSubscriptionList(t *testing.T) {
	customerParams := &stripe.CustomerParams{
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	planParams := &stripe.PlanParams{
		ID:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: currency.USD,
		Interval: plan.Month,
	}

	plan.New(planParams)

	subParams := &stripe.SubParams{
		Customer: cust.ID,
		Plan:     "test",
		Quantity: 10,
	}

	for i := 0; i < 5; i++ {
		New(subParams)
	}

	i := List(&stripe.SubListParams{Customer: cust.ID})
	for i.Next() {
		if i.Sub() == nil {
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
	plan.Del("test")
}
