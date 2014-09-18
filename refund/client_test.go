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

func TestRefundNew(t *testing.T) {
	chargeParams := &ChargeParams{
		Amount:   1000,
		Currency: USD,
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	res, _ := charge.New(chargeParams)

	// full refund
	ref, err := New(&RefundParams{Charge: res.Id})

	if err != nil {
		t.Error(err)
	}

	if ref.Charge != res.Id {
		t.Errorf("Refund charge %q does not match expected value %v\n", ref.Charge, res.Id)
	}

	target, _ := charge.Get(res.Id, nil)

	if !target.Refunded || target.Refunds == nil {
		t.Errorf("Expected to have refunded this charge\n")
	}

	if len(target.Refunds.Values) != 1 {
		t.Errorf("Expected to have a refund, but instead have %v\n", len(target.Refunds.Values))
	}

	if target.Refunds.Values[0].Amount != target.AmountRefunded {
		t.Errorf("Refunded amount %v does not match amount refunded %v\n", target.Refunds.Values[0].Amount, target.AmountRefunded)
	}

	if target.Refunds.Values[0].Currency != target.Currency {
		t.Errorf("Refunded currency %q does not match charge currency %q\n", target.Refunds.Values[0].Currency, target.Currency)
	}

	if len(target.Refunds.Values[0].Tx.Id) == 0 {
		t.Errorf("Refund transaction not set\n")
	}

	if target.Refunds.Values[0].Charge != target.Id {
		t.Errorf("Refund charge %q does not match expected value %v\n", target.Refunds.Values[0].Charge, target.Id)
	}

	res, err = charge.New(chargeParams)

	// partial refund
	refundParams := &RefundParams{
		Charge: res.Id,
		Amount: 253,
	}

	New(refundParams)

	target, _ = charge.Get(res.Id, nil)

	if target.Refunded {
		t.Errorf("Partial refund should not be marked as Refunded\n")
	}

	if target.AmountRefunded != refundParams.Amount {
		t.Errorf("Refunded amount %v does not match expected amount %v\n", target.AmountRefunded, refundParams.Amount)
	}
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

	ch, _ := charge.New(chargeParams)
	ref, _ := New(&RefundParams{Charge: ch.Id})

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

	ch, _ := charge.New(chargeParams)
	for i := 0; i < 4; i++ {
		New(&RefundParams{Charge: ch.Id, Amount: 200})
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
