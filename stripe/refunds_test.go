package stripe

import "testing"

func TestRefundGet(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	charge := &ChargeParams{
		Amount:   1000,
		Currency: USD,
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	ch, _ := c.Charges.Create(charge)
	refund, _ := c.Charges.Refund(&RefundParams{Charge: ch.Id})

	target, err := c.Refunds.Get(refund.Id, &RefundParams{Charge: ch.Id})

	if err != nil {
		t.Error(err)
	}

	if target.Charge != ch.Id {
		t.Errorf("Refund charge %q does not match expected value %v\n", target.Charge, ch.Id)
	}
}

func TestRefundList(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	charge := &ChargeParams{
		Amount:   1000,
		Currency: USD,
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	ch, _ := c.Charges.Create(charge)
	for i := 0; i < 4; i++ {
		c.Charges.Refund(&RefundParams{Charge: ch.Id, Amount: 200})
	}

	target, err := c.Refunds.List(&RefundListParams{Charge: ch.Id})

	if err != nil {
		t.Error(err)
	}

	if len(target.Values) != 4 {
		t.Errorf("The number of refunds %v does not match expected value\n", len(target.Values))
	}

	for _, r := range target.Values {
		if r.Amount != 200 {
			t.Errorf("Refund amount %v does not match expected value\n", r.Amount)
		}

		if r.Charge != ch.Id {
			t.Errorf("Refund charge %q does not match expected value %q\n", r.Charge, ch.Id)
		}
	}
}
