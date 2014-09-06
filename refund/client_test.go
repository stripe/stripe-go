package refund

import (
	"testing"

	. "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	Key = GetTestKey()
}

func TestRefundGet(t *testing.T) {
	chargeParams := &ChargeParams{
		Amount:   1000,
		Currency: USD,
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	ch, _ := charge.Create(chargeParams)
	ref, _ := charge.RefundCharge(&RefundParams{Charge: ch.Id})

	target, err := Get(ref.Id, &RefundParams{Charge: ch.Id})

	if err != nil {
		t.Error(err)
	}

	if target.Charge != ch.Id {
		t.Errorf("Refund charge %q does not match expected value %v\n", target.Charge, ch.Id)
	}
}

func TestRefundList(t *testing.T) {
	chargeParams := &ChargeParams{
		Amount:   1000,
		Currency: USD,
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	ch, _ := charge.Create(chargeParams)
	for i := 0; i < 4; i++ {
		charge.RefundCharge(&RefundParams{Charge: ch.Id, Amount: 200})
	}

	i := List(&RefundListParams{Charge: ch.Id})
	for !i.Stop() {
		target, err := i.Next()

		if err != nil {
			t.Error(err)
		}

		if target.Amount != 200 {
			t.Error("Amount %v does not match expected value\n", target.Amount)
		}

		if target.Charge != ch.Id {
			t.Errorf("Refund charge %q does not match expected value %q\n", target.Charge, ch.Id)
		}

		if i.Meta() == nil {
			t.Error("No metadata returned")
		}
	}
}
