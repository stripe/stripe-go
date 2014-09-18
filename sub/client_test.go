package sub

import (
	"testing"

	. "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/coupon"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/discount"
	"github.com/stripe/stripe-go/plan"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	Key = GetTestKey()
}

func TestSubscriptionNew(t *testing.T) {
	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	planParams := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}

	plan.New(planParams)

	subParams := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	target, err := New(subParams)

	if err != nil {
		t.Error(err)
	}

	if target.Plan.Id != subParams.Plan {
		t.Errorf("Plan %q does not match expected plan %q\n", target.Plan, subParams.Plan)
	}

	if target.Quantity != subParams.Quantity {
		t.Errorf("Quantity %v does not match expected quantity %v\n", target.Quantity, subParams.Quantity)
	}

	customer.Del(cust.Id)
	plan.Del("test")
}

func TestSubscriptionGet(t *testing.T) {
	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	planParams := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}

	plan.New(planParams)

	subParams := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	subscription, _ := New(subParams)
	target, err := Get(subscription.Id, &SubParams{Customer: cust.Id})

	if err != nil {
		t.Error(err)
	}

	if target.Id != subscription.Id {
		t.Errorf("Subscription id %q does not match expected id %q\n", target.Id, subscription.Id)
	}

	customer.Del(cust.Id)
	plan.Del("test")
}

func TestSubscriptionCancel(t *testing.T) {
	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	planParams := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}

	plan.New(planParams)

	subParams := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	subscription, _ := New(subParams)
	err := Cancel(subscription.Id, &SubParams{Customer: cust.Id})

	if err != nil {
		t.Error(err)
	}

	customer.Del(cust.Id)
	plan.Del("test")
}

func TestSubscriptionUpdate(t *testing.T) {
	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	planParams := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}

	plan.New(planParams)

	subParams := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	subscription, _ := New(subParams)
	updatedSub := &SubParams{
		Customer:  cust.Id,
		NoProrate: true,
		Quantity:  13,
	}

	target, err := Update(subscription.Id, updatedSub)

	if err != nil {
		t.Error(err)
	}

	if target.Quantity != updatedSub.Quantity {
		t.Errorf("Quantity %v does not match expected quantity $v\n", target.Quantity, updatedSub.Quantity)
	}

	customer.Del(cust.Id)
	plan.Del("test")
}

func TestSubscriptionDiscount(t *testing.T) {
	couponParams := &CouponParams{
		Duration: Forever,
		Id:       "sub_coupon",
		Percent:  99,
	}

	coupon.New(couponParams)

	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		Coupon: "sub_coupon",
	}

	cust, _ := customer.New(customerParams)

	planParams := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}

	plan.New(planParams)

	subParams := &SubParams{
		Customer: cust.Id,
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

	if target.Discount.Coupon.Id != subParams.Coupon {
		t.Errorf("Coupon id %q does not match expected id %q\n", target.Discount.Coupon.Id, subParams.Coupon)
	}

	err = discount.DelSub(cust.Id, target.Id)

	if err != nil {
		t.Error(err)
	}

	customer.Del(cust.Id)
	plan.Del("test")
	coupon.Del("sub_coupon")
}

func TestSubscriptionList(t *testing.T) {
	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	planParams := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}

	plan.New(planParams)

	subParams := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	for i := 0; i < 5; i++ {
		New(subParams)
	}

	i := List(&SubListParams{Customer: cust.Id})
	for !i.Stop() {
		target, err := i.Next()

		if err != nil {
			t.Error(err)
		}

		if target == nil {
			t.Error("No nil values expected")
		}

		if i.Meta() == nil {
			t.Error("No metadata returned")
		}
	}

	customer.Del(cust.Id)
	plan.Del("test")
}
