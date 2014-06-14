package stripe

import (
	"testing"
)

func TestChargeCreate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	charge := &ChargeParams{
		Amount:   1000,
		Currency: USD,
		Card: &CardParams{
			Name:   "Stripe Tester",
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	target, err := c.Charges.Create(charge)
	if err != nil {
		t.Error(err)
	}

	if target.Amount != charge.Amount {
		t.Errorf("Amount %v does not match expected amount %v\n", target.Amount, charge.Amount)
	}

	if target.Currency != charge.Currency {
		t.Errorf("Currency %q does not match expected currency %q\n", target.Currency, charge.Currency)
	}

	if target.Card.Name != charge.Card.Name {
		t.Errorf("Card name %q does not match expected name %q\n", target.Card.Name, charge.Card.Name)
	}
}

func TestChargeGet(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	charge := &ChargeParams{
		Amount:   1001,
		Currency: USD,
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	res, _ := c.Charges.Create(charge)

	target, err := c.Charges.Get(res.Id)

	if err != nil {
		t.Error(err)
	}

	if target.Id != res.Id {
		t.Errorf("Charge id %q does not match expected id %q\n", target.Id, res.Id)
	}
}

func TestChargeUpdate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	charge := &ChargeParams{
		Amount:   1002,
		Currency: USD,
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		Desc: "original description",
	}

	res, _ := c.Charges.Create(charge)

	if res.Desc != charge.Desc {
		t.Errorf("Original description %q does not match expected description %q\n", res.Desc, charge.Desc)
	}

	updated := &ChargeParams{
		Desc: "updated description",
	}

	target, err := c.Charges.Update(res.Id, updated)

	if err != nil {
		t.Error(err)
	}

	if target.Desc != updated.Desc {
		t.Errorf("Updated description %q does not match expected description %q\n", target.Desc, updated.Desc)
	}
}

func TestChargeRefund(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	charge := &ChargeParams{
		Amount:   1003,
		Currency: USD,
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	res, _ := c.Charges.Create(charge)

	// full refund
	target, err := c.Charges.Refund(res.Id, nil)

	if err != nil {
		t.Error(err)
	}

	if !target.Refunded {
		t.Errorf("Expected to have refunded this charge\n")
	}

	if len(target.Refunds) != 1 {
		t.Errorf("Expected to have a refund, but instead have %v\n", len(target.Refunds))
	}

	if target.Refunds[0].Amount != target.AmountRefunded {
		t.Errorf("Refunded amount %v does not match amount refunded %v\n", target.Refunds[0].Amount, target.AmountRefunded)
	}

	if target.Refunds[0].Currency != target.Currency {
		t.Errorf("Refunded currency %q does not match charge currency %q\n", target.Refunds[0].Currency, target.Currency)
	}

	res, err = c.Charges.Create(charge)

	// partial refund
	refund := &RefundParams{
		Amount: 253,
	}

	target, err = c.Charges.Refund(res.Id, refund)

	if err != nil {
		t.Error(err)
	}

	if target.Refunded {
		t.Errorf("Partial refund should not be marked as Refunded\n")
	}

	if target.AmountRefunded != refund.Amount {
		t.Errorf("Refunded amount %v does not match expected amount %v\n", target.AmountRefunded, refund.Amount)
	}
}

func TestChargeCapture(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	charge := &ChargeParams{
		Amount:   1004,
		Currency: USD,
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		NoCapture: true,
	}

	res, _ := c.Charges.Create(charge)

	if res.Captured {
		t.Errorf("The charge should not have been captured\n")
	}

	// full capture
	target, err := c.Charges.Capture(res.Id, nil)

	if err != nil {
		t.Error(err)
	}

	if !target.Captured {
		t.Errorf("Expected to have captured this charge after full capture\n")
	}

	res, err = c.Charges.Create(charge)

	// partial capture
	capture := &CaptureParams{
		Amount: 554,
	}

	target, err = c.Charges.Capture(res.Id, capture)

	if err != nil {
		t.Error(err)
	}

	if !target.Captured {
		t.Errorf("Expected to have captured this charge after partial capture\n")
	}

	if target.AmountRefunded != charge.Amount-capture.Amount {
		t.Errorf("Refunded amount %v does not match expected amount %v\n", target.AmountRefunded, charge.Amount-capture.Amount)
	}
}

func TestChargeList(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	params := &ChargeListParams{}
	params.Filters.AddFilter("include", "", "total_count")
	params.Filters.AddFilter("created", "gte", "1402729261")
	target, err := c.Charges.List(params)

	if err != nil {
		t.Error(err)
	}

	if target.Count == 0 {
		t.Errorf("Count is not set\n")
	}

	for _, v := range target.Values {
		if v.Amount == 0 {
			t.Errorf("Amount is not set\n")
		}
	}
}
