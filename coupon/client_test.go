package coupon

import (
	"fmt"
	"testing"
	"time"

	. "github.com/stripe/stripe-go"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	Key = GetTestKey()
}

func TestCouponCreate(t *testing.T) {
	couponParams := &CouponParams{
		Amount:         99,
		Currency:       USD,
		Duration:       Repeating,
		DurationPeriod: 3,
		Redemptions:    1,
		RedeemBy:       time.Now().AddDate(0, 0, 30).Unix(),
	}

	target, err := Create(couponParams)

	if err != nil {
		t.Error(err)
	}

	if target.Amount != couponParams.Amount {
		t.Errorf("Amount %v does not match expected amount %v\n", target.Amount, couponParams.Amount)
	}

	if target.Currency != couponParams.Currency {
		t.Errorf("Currency %q does not match expected currency %q\n", target.Currency, couponParams.Currency)
	}

	if target.Duration != couponParams.Duration {
		t.Errorf("Duration %q does not match expected duration %q\n", target.Duration, couponParams.Duration)
	}

	if target.DurationPeriod != couponParams.DurationPeriod {
		t.Errorf("Duration period %v does not match expected duration period %v\n", target.DurationPeriod, couponParams.DurationPeriod)
	}

	if target.Redemptions != couponParams.Redemptions {
		t.Errorf("Max redemptions %v does not match expected max redemptions %v\n", target.Redemptions, couponParams.Redemptions)
	}

	if target.RedeemBy != couponParams.RedeemBy {
		t.Errorf("Redeem by %v does not match expected redeem by %v\n", target.RedeemBy, couponParams.RedeemBy)
	}

	if !target.Valid {
		t.Errorf("Coupon is not valid, but was expecting it to be\n")
	}

	Delete(target.Id)
}

func TestCouponGet(t *testing.T) {
	couponParams := &CouponParams{
		Id:       "test_coupon",
		Duration: Once,
		Percent:  50,
	}

	Create(couponParams)
	target, err := Get(couponParams.Id, nil)

	if err != nil {
		t.Error(err)
	}

	if target.Id != couponParams.Id {
		t.Errorf("Id %q does not match expected id %q\n", target.Id, couponParams.Id)
	}

	if target.Percent != couponParams.Percent {
		t.Errorf("Percent %v does not match expected percent %v\n", target.Percent, couponParams.Percent)
	}

	Delete(target.Id)
}

func TestCouponList(t *testing.T) {
	for i := 0; i < 5; i++ {
		couponParams := &CouponParams{
			Id:       fmt.Sprintf("test_%v", i),
			Duration: Once,
			Percent:  50,
		}

		Create(couponParams)
	}

	target, err := List(nil)

	if err != nil {
		t.Error(err)
	}

	if len(target.Values) != 5 {
		t.Errorf("Count %v does not match expected value\n", len(target.Values))
	}

	for i := 0; i < 5; i++ {
		Delete(fmt.Sprintf("test_%v", i))
	}
}
