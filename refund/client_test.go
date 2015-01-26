package refund

import (
	"testing"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/charge"
	"github.com/stripe/stripe-go/currency"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	stripe.Key = GetTestKey()
}

func TestRefundNew(t *testing.T) {
	chargeParams := &stripe.ChargeParams{
		Amount:   1000,
		Currency: currency.USD,
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	res, _ := charge.New(chargeParams)

	// full refund
	ref, err := New(&stripe.RefundParams{Charge: res.ID})

	if err != nil {
		t.Error(err)
	}

	if ref.Charge != res.ID {
		t.Errorf("Refund charge %q does not match expected value %v\n", ref.Charge, res.ID)
	}

	target, _ := charge.Get(res.ID, nil)

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

	if len(target.Refunds.Values[0].Tx.ID) == 0 {
		t.Errorf("Refund transaction not set\n")
	}

	if target.Refunds.Values[0].Charge != target.ID {
		t.Errorf("Refund charge %q does not match expected value %v\n", target.Refunds.Values[0].Charge, target.ID)
	}

	res, err = charge.New(chargeParams)

	// partial refund
	refundParams := &stripe.RefundParams{
		Charge: res.ID,
		Amount: 253,
	}

	New(refundParams)

	target, _ = charge.Get(res.ID, nil)

	if target.Refunded {
		t.Errorf("Partial refund should not be marked as Refunded\n")
	}

	if target.AmountRefunded != refundParams.Amount {
		t.Errorf("Refunded amount %v does not match expected amount %v\n", target.AmountRefunded, refundParams.Amount)
	}

	// refund with reason
	res, err = charge.New(chargeParams)
	New(&stripe.RefundParams{Charge: res.ID, Reason: RefundFraudulent})
	target, _ = charge.Get(res.ID, nil)

	if target.FraudDetails.UserReport != "fraudulent" {
		t.Errorf("Expected a fraudulent UserReport for charge refunded with reason=fraudulent but got: %s",
			target.FraudDetails.UserReport)
	}
}

func TestRefundGet(t *testing.T) {
	chargeParams := &stripe.ChargeParams{
		Amount:   1000,
		Currency: currency.USD,
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	ch, _ := charge.New(chargeParams)
	ref, _ := New(&stripe.RefundParams{Charge: ch.ID})

	target, err := Get(ref.ID, &stripe.RefundParams{Charge: ch.ID})

	if err != nil {
		t.Error(err)
	}

	if target.Charge != ch.ID {
		t.Errorf("Refund charge %q does not match expected value %v\n", target.Charge, ch.ID)
	}
}

func TestRefundList(t *testing.T) {
	chargeParams := &stripe.ChargeParams{
		Amount:   1000,
		Currency: currency.USD,
		Card: &stripe.CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	ch, _ := charge.New(chargeParams)
	for i := 0; i < 4; i++ {
		New(&stripe.RefundParams{Charge: ch.ID, Amount: 200})
	}

	i := List(&stripe.RefundListParams{Charge: ch.ID})
	for i.Next() {
		target := i.Refund()

		if target.Amount != 200 {
			t.Errorf("Amount %v does not match expected value\n", target.Amount)
		}

		if target.Charge != ch.ID {
			t.Errorf("Refund charge %q does not match expected value %q\n", target.Charge, ch.ID)
		}

		if i.Meta() == nil {
			t.Error("No metadata returned")
		}
	}
	if err := i.Err(); err != nil {
		t.Error(err)
	}
}
