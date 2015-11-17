package order

import (
	"math/rand"
	"testing"
	"time"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/coupon"
	"github.com/stripe/stripe-go/currency"
	"github.com/stripe/stripe-go/orderitem"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
	rand.Seed(time.Now().UTC().UnixNano())
}

func TestOrder(t *testing.T) {
	o, err := New(&stripe.OrderParams{
		Currency: "usd",
		Items: []*stripe.OrderItemParams{
			&stripe.OrderItemParams{
				Type:   "sku",
				Parent: "sku_7EjB0wuV6CBT8x",
			},
		},
		Shipping: &stripe.ShippingParams{
			Name: "Jenny Rosen",
			Address: &stripe.AddressParams{
				Line1:      "1234 Main Street",
				City:       "Anytown",
				Country:    "US",
				PostalCode: "123456",
			},
			Phone: "6504244242",
		},
		Email: "jenny@ros.en",
		Params: stripe.Params{
			Meta: map[string]string{
				"foo": "bar",
			},
		},
	})

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if o.ID == "" {
		t.Errorf("ID is not set %v", o.ID)
	}

	if o.Created == 0 {
		t.Errorf("Created date is not set")
	}

	if o.Updated == 0 {
		t.Errorf("Updated is not set")
	}

	if o.Currency != "usd" {
		t.Errorf("Currency is invalid: %v", o.Currency)
	}

	if o.Email != "jenny@ros.en" {
		t.Errorf("Email is invalid: %v", o.Email)
	}

	if o.Shipping.Name != "Jenny Rosen" {
		t.Errorf("Name is invalid: %v", o.Shipping.Name)
	}

	if o.Shipping.Address.Line1 != "1234 Main Street" {
		t.Errorf(
			"Shipping address line 1 is invalid: %v",
			o.Shipping.Address.Line1,
		)
	}

	if o.Shipping.Phone != "6504244242" {
		t.Errorf("Shipping phone is invalid: %v", o.Shipping.Phone)
	}

	if o.Status != "created" {
		t.Errorf("Order status is invalid: %v", o.Status)
	}

	if len(o.Items) == 0 {
		t.Error("Order has no items")
	}

	if o.Meta["foo"] != "bar" {
		t.Error("Order metadata not set")
	}
}

func TestOrderUpdate(t *testing.T) {
	o, err := New(&stripe.OrderParams{
		Currency: currency.USD,
		Items: []*stripe.OrderItemParams{
			&stripe.OrderItemParams{
				Type:   "sku",
				Parent: "sku_7EjB0wuV6CBT8x",
			},
		},
		Shipping: &stripe.ShippingParams{
			Name: "Jenny Rosen",
			Address: &stripe.AddressParams{
				Line1:      "1234 Main Street",
				City:       "Anytown",
				Country:    "US",
				PostalCode: "123456",
			},
			Phone: "6504244242",
		},
		Email: "jenny@ros.en",
		Params: stripe.Params{
			Meta: map[string]string{
				"foo": "bar",
			},
		},
	})

	if err != nil {
		t.Fatalf("%+v", err)
	}

	couponParams := &stripe.CouponParams{
		Amount:         99,
		Currency:       currency.USD,
		Duration:       coupon.Once,
		DurationPeriod: 3,
		Redemptions:    1,
		RedeemBy:       time.Now().AddDate(0, 0, 30).Unix(),
	}

	c, err := coupon.New(couponParams)

	updated, err := Update(
		o.ID,
		&stripe.OrderUpdateParams{
			Coupon: c.ID,
			Params: stripe.Params{
				Meta: map[string]string{"foo": "baz"},
			},
			Status: stripe.StatusCanceled,
		},
	)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if updated.Status != stripe.StatusCanceled {
		t.Errorf("Order status not updated: %v", updated.Status)
	}

	hasCoupon := false
	for _, item := range updated.Items {
		if item.Type == orderitem.Discount {
			hasCoupon = true
		}
	}
	if !hasCoupon {
		t.Errorf("Discount code not applied: %v", updated.Items)
	}

	if updated.Meta["foo"] != "baz" {
		t.Errorf("Order metadata not updated: %v", updated.Meta["foo"])
	}

	coupon.Del(c.ID)
}

func TestOrderPay(t *testing.T) {
	o, err := New(&stripe.OrderParams{
		Currency: currency.USD,
		Items: []*stripe.OrderItemParams{
			&stripe.OrderItemParams{
				Type:   "sku",
				Parent: "sku_7EjB0wuV6CBT8x",
			},
		},
		Shipping: &stripe.ShippingParams{
			Name: "Jenny Rosen",
			Address: &stripe.AddressParams{
				Line1:      "1234 Main Street",
				City:       "Anytown",
				Country:    "US",
				PostalCode: "123456",
			},
			Phone: "6504244242",
		},
		Email: "jenny@ros.en",
	})

	if err != nil {
		t.Fatalf("%+v", err)
	}

	params := &stripe.OrderPayParams{
		Params: stripe.Params{
			Meta: map[string]string{"order_id": "137"},
		},
	}
	params.SetSource(&stripe.CardParams{
		Name:     "Stripe Tester",
		Number:   "4242424242424242",
		Month:    "06",
		Year:     "20",
		Address1: "1234 Main Street",
		Address2: "Apt 1",
		City:     "Anytown",
		State:    "CA",
	})

	order, err := Pay(o.ID, params)

	if err != nil {
		t.Fatalf("%+v", err)
	}

	if order.Meta["order_id"] != "137" {
		t.Errorf("Order metadata not set: %v", order.Meta["order_id"])
	}

	if order.Status != stripe.StatusPaid {
		t.Errorf("Order status not set to paid: %v", order.Status)
	}
}

func TestOrderList(t *testing.T) {
	params := &stripe.OrderParams{
		Currency: currency.USD,
		Items: []*stripe.OrderItemParams{
			&stripe.OrderItemParams{
				Type:   "sku",
				Parent: "sku_7EjB0wuV6CBT8x",
			},
		},
		Shipping: &stripe.ShippingParams{
			Name: "Jenny Rosen",
			Address: &stripe.AddressParams{
				Line1:      "1234 Main Street",
				City:       "Anytown",
				Country:    "US",
				PostalCode: "123456",
			},
			Phone: "6504244242",
		},
		Email: "jenny@ros.en",
	}

	first, err := New(params)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	_, err = Update(
		first.ID,
		&stripe.OrderUpdateParams{
			Status: stripe.StatusCanceled,
		},
	)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	second, err := New(params)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	listParams := &stripe.OrderListParams{
		IDs:    []string{first.ID, second.ID},
		Status: stripe.StatusCanceled,
	}

	i := List(listParams)
	count := 0
	for i.Next() {
		target := i.Order()

		if target.Status != stripe.StatusCanceled {
			t.Errorf(
				"Order list should only include status=%v, got %v\n",
				stripe.StatusCanceled,
				target.Status,
			)
		}
		count++

	}
	if count != 1 {
		t.Errorf("Expected to get 1 object, got %v", count)
	}
	if err := i.Err(); err != nil {
		t.Error(err)
	}
}
