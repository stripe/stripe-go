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

func TestSubscriptionCreate(t *testing.T) {
	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.Create(customerParams)

	planParams := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}

	plan.Create(planParams)

	subParams := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	target, err := Create(subParams)

	if err != nil {
		t.Error(err)
	}

	if target.Plan.Id != subParams.Plan {
		t.Errorf("Plan %q does not match expected plan %q\n", target.Plan, subParams.Plan)
	}

	if target.Quantity != subParams.Quantity {
		t.Errorf("Quantity %v does not match expected quantity %v\n", target.Quantity, subParams.Quantity)
	}

	customer.Delete(cust.Id)
	plan.Delete("test")
}

func TestSubscriptionGet(t *testing.T) {
	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.Create(customerParams)

	planParams := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}

	plan.Create(planParams)

	subParams := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	subscription, _ := Create(subParams)
	target, err := Get(subscription.Id, &SubParams{Customer: cust.Id})

	if err != nil {
		t.Error(err)
	}

	if target.Id != subscription.Id {
		t.Errorf("Subscription id %q does not match expected id %q\n", target.Id, subscription.Id)
	}

	customer.Delete(cust.Id)
	plan.Delete("test")
}

func TestSubscriptionCancel(t *testing.T) {
	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.Create(customerParams)

	planParams := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}

	plan.Create(planParams)

	subParams := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	subscription, _ := Create(subParams)
	err := Cancel(subscription.Id, &SubParams{Customer: cust.Id})

	if err != nil {
		t.Error(err)
	}

	customer.Delete(cust.Id)
	plan.Delete("test")
}

func TestSubscriptionUpdate(t *testing.T) {
	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.Create(customerParams)

	planParams := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}

	plan.Create(planParams)

	subParams := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	subscription, _ := Create(subParams)
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

	customer.Delete(cust.Id)
	plan.Delete("test")
}

func TestSubscriptionDiscount(t *testing.T) {
	couponParams := &CouponParams{
		Duration: Forever,
		Id:       "sub_coupon",
		Percent:  99,
	}

	coupon.Create(couponParams)

	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		Coupon: "sub_coupon",
	}

	cust, _ := customer.Create(customerParams)

	planParams := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}

	plan.Create(planParams)

	subParams := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
		Coupon:   "sub_coupon",
	}

	target, err := Create(subParams)

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

	err = discount.DeleteSubscription(cust.Id, target.Id)

	if err != nil {
		t.Error(err)
	}

	customer.Delete(cust.Id)
	plan.Delete("test")
	coupon.Delete("sub_coupon")
}

func TestSubscriptionList(t *testing.T) {
	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.Create(customerParams)

	planParams := &PlanParams{
		Id:       "test",
		Name:     "Test Plan",
		Amount:   99,
		Currency: USD,
		Interval: Month,
	}

	plan.Create(planParams)

	subParams := &SubParams{
		Customer: cust.Id,
		Plan:     "test",
		Quantity: 10,
	}

	for i := 0; i < 5; i++ {
		Create(subParams)
	}

	target, err := List(&SubListParams{Customer: cust.Id})

	if err != nil {
		t.Error(err)
	}

	if len(target.Values) != 5 {
		t.Errorf("Count %v does not match expected value\n", len(target.Values))
	}

	customer.Delete(cust.Id)
	plan.Delete("test")
}
