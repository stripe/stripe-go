package stripe

import (
	"fmt"
	"testing"
	"time"
)

func TestCouponCreate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	coupon := &CouponParams{
		Amount:         99,
		Currency:       USD,
		Duration:       Repeating,
		DurationPeriod: 3,
		Redemptions:    1,
		RedeemBy:       time.Now().AddDate(0, 0, 30).Unix(),
	}

	target, err := c.Coupons.Create(coupon)

	if err != nil {
		t.Error(err)
	}

	if target.Amount != coupon.Amount {
		t.Errorf("Amount %v does not match expected amount %v\n", target.Amount, coupon.Amount)
	}

	if target.Currency != coupon.Currency {
		t.Errorf("Currency %q does not match expected currency %q\n", target.Currency, coupon.Currency)
	}

	if target.Duration != coupon.Duration {
		t.Errorf("Duration %q does not match expected duration %q\n", target.Duration, coupon.Duration)
	}

	if target.DurationPeriod != coupon.DurationPeriod {
		t.Errorf("Duration period %v does not match expected duration period %v\n", target.DurationPeriod, coupon.DurationPeriod)
	}

	if target.Redemptions != coupon.Redemptions {
		t.Errorf("Max redemptions %v does not match expected max redemptions %v\n", target.Redemptions, coupon.Redemptions)
	}

	if target.RedeemBy != coupon.RedeemBy {
		t.Errorf("Redeem by %v does not match expected redeem by %v\n", target.RedeemBy, coupon.RedeemBy)
	}

	if !target.Valid {
		t.Errorf("Coupon is not valid, but was expecting it to be\n")
	}

	c.Coupons.Delete(target.Id)
}

func TestCouponGet(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	coupon := &CouponParams{
		Id:       "test_coupon",
		Duration: Once,
		Percent:  50,
	}

	c.Coupons.Create(coupon)
	target, err := c.Coupons.Get(coupon.Id)

	if err != nil {
		t.Error(err)
	}

	if target.Id != coupon.Id {
		t.Errorf("Id %q does not match expected id %q\n", target.Id, coupon.Id)
	}

	if target.Percent != coupon.Percent {
		t.Errorf("Percent %v does not match expected percent %v\n", target.Percent, coupon.Percent)
	}

	c.Coupons.Delete(target.Id)
}

func TestCouponList(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	for i := 0; i < 5; i++ {
		coupon := &CouponParams{
			Id:       fmt.Sprintf("test_%v", i),
			Duration: Once,
			Percent:  50,
		}

		c.Coupons.Create(coupon)
	}

	target, err := c.Coupons.List(nil)

	if err != nil {
		t.Error(err)
	}

	if len(target.Values) != 5 {
		t.Errorf("Count %v does not match expected value\n", len(target.Values))
	}

	for i := 0; i < 5; i++ {
		c.Coupons.Delete(fmt.Sprintf("test_%v", i))
	}
}
