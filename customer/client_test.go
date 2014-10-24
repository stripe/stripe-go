package customer

import (
	"testing"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/coupon"
	"github.com/stripe/stripe-go/discount"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestCustomerNew(t *testing.T) {
	customerParams := &stripe.CustomerParams{
		Balance: -123,
		Card: &stripe.CardParams{
			Name:   "Test Card",
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		Desc:  "Test Customer",
		Email: "a@b.com",
	}

	customerParams.AddMeta("key", "value")
	target, err := New(customerParams)

	if err != nil {
		t.Error(err)
	}

	if target.Balance != customerParams.Balance {
		t.Errorf("Balance %v does not match expected balance %v\n", target.Balance, customerParams.Balance)
	}

	if target.Desc != customerParams.Desc {
		t.Errorf("Description %q does not match expected description %q\n", target.Desc, customerParams.Desc)
	}

	if target.Email != customerParams.Email {
		t.Errorf("Email %q does not match expected email %q\n", target.Email, customerParams.Email)
	}

	if target.Meta["id"] != customerParams.Meta["id"] {
		t.Errorf("Meta %v does not match expected Meta %v\n", target.Meta, customerParams.Meta)
	}

	if target.Cards == nil {
		t.Errorf("No cards recorded\n")
	}

	if target.Cards.Count != 1 {
		t.Errorf("Unexpected number of cards %v\n", target.Cards.Count)
	}

	if target.Cards.Values[0].Name != customerParams.Card.Name {
		t.Errorf("Card name %q does not match expected name %q\n", target.Cards.Values[0].Name, customerParams.Card.Name)
	}

	Del(target.ID)
}

func TestCustomerGet(t *testing.T) {
	res, _ := New(nil)

	target, err := Get(res.ID, nil)

	if err != nil {
		t.Error(err)
	}

	if target.ID != res.ID {
		t.Errorf("Customer id %q does not match expected id %q\n", target.ID, res.ID)
	}

	Del(res.ID)
}

func TestCustomerDel(t *testing.T) {
	res, _ := New(nil)

	err := Del(res.ID)

	if err != nil {
		t.Error(err)
	}
}

func TestCustomerUpdate(t *testing.T) {
	customerParams := &stripe.CustomerParams{
		Balance: 7,
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		Desc:  "Original Desc",
		Email: "first@b.com",
	}

	original, _ := New(customerParams)

	updated := &stripe.CustomerParams{
		Balance: -10,
		Card: &stripe.CardParams{
			Number: "4242424242424242",
			Month:  "10",
			Year:   "20",
		},
		Desc:  "Updated Desc",
		Email: "desc@b.com",
	}

	target, err := Update(original.ID, updated)

	if err != nil {
		t.Error(err)
	}

	if target.Balance != updated.Balance {
		t.Errorf("Balance %v does not match expected balance %v\n", target.Balance, updated.Balance)
	}

	if target.Desc != updated.Desc {
		t.Errorf("Description %q does not match expected description %q\n", target.Desc, updated.Desc)
	}

	if target.Email != updated.Email {
		t.Errorf("Email %q does not match expected email %q\n", target.Email, updated.Email)
	}

	if target.Cards == nil {
		t.Errorf("No cards recorded\n")
	}

	Del(target.ID)
}

func TestCustomerDiscount(t *testing.T) {
	couponParams := &stripe.CouponParams{
		Duration: coupon.Forever,
		ID:       "customer_coupon",
		Percent:  99,
	}

	coupon.New(couponParams)

	customerParams := &stripe.CustomerParams{
		Coupon: "customer_coupon",
	}

	target, err := New(customerParams)

	if err != nil {
		t.Error(err)
	}

	if target.Discount == nil {
		t.Errorf("Discount not found, but one was expected\n")
	}

	if target.Discount.Coupon == nil {
		t.Errorf("Coupon not found, but one was expected\n")
	}

	if target.Discount.Coupon.ID != customerParams.Coupon {
		t.Errorf("Coupon id %q does not match expected id %q\n", target.Discount.Coupon.ID, customerParams.Coupon)
	}

	err = discount.Del(target.ID)

	if err != nil {
		t.Error(err)
	}

	Del(target.ID)
	coupon.Del("customer_coupon")
}

func TestCustomerList(t *testing.T) {
	customers := make([]string, 5)

	for i := 0; i < 5; i++ {
		cust, _ := New(nil)
		customers[i] = cust.ID
	}

	i := List(nil)
	for i.Next() {
		if i.Customer() == nil {
			t.Error("No nil values expected")
		}

		if i.Meta() == nil {
			t.Error("No metadata returned")
		}
	}
	if err := i.Err(); err != nil {
		t.Error(err)
	}

	for _, v := range customers {
		Del(v)
	}
}
