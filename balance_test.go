package stripe

import (
	"testing"
)

func TestBalanceGet(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	target, err := c.Balance.Get()

	if err != nil {
		t.Error(err)
	}

	if target.Available == nil || len(target.Available) != 1 {
		t.Errorf("Available array is not set\n")
	}

	if target.Pending == nil || len(target.Pending) != 1 {
		t.Errorf("Pending array is not set\n")
	}

	if len(target.Available[0].Currency) == 0 {
		t.Errorf("Available currency is not set\n")
	}

	if len(target.Pending[0].Currency) == 0 {
		t.Errorf("Pending currency is not set\n")
	}
}

func TestBalanceGetTx(t *testing.T) {
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
		Desc: "charge transaction",
	}

	res, _ := c.Charges.Create(charge)

	target, err := c.Balance.GetTx(res.Tx.Id)

	if err != nil {
		t.Error(err)
	}

	if uint64(target.Amount) != charge.Amount {
		t.Errorf("Amount %v does not match expected amount %v\n", target.Amount, charge.Amount)
	}

	if target.Currency != charge.Currency {
		t.Errorf("Currency %q does not match expected currency %q\n", target.Currency, charge.Currency)
	}

	if target.Desc != charge.Desc {
		t.Errorf("Description %q does not match expected description %q\n", target.Desc, charge.Desc)
	}

	if target.Available == 0 {
		t.Errorf("Available date is not set\n")
	}

	if target.Created == 0 {
		t.Errorf("Created date is not set\n")
	}

	if target.Fee == 0 {
		t.Errorf("Fee is not set\n")
	}

	if target.FeeDetails == nil || len(target.FeeDetails) != 1 {
		t.Errorf("Fee details are not set")
	}

	if target.FeeDetails[0].Amount == 0 {
		t.Errorf("Fee detail amount is not set\n")
	}

	if len(target.FeeDetails[0].Currency) == 0 {
		t.Errorf("Fee detail currency is not set\n")
	}

	if len(target.FeeDetails[0].Desc) == 0 {
		t.Errorf("Fee detail description is not set\n")
	}

	if target.Net == 0 {
		t.Errorf("Net is not set\n")
	}

	if target.Status != TxPending {
		t.Errorf("Status %v does not match expected value\n", target.Status)
	}

	if target.Type != TxCharge {
		t.Errorf("Type %v does not match expected value\n", target.Type)
	}

	if target.Src != res.Id {
		t.Errorf("Source %q does not match expeted value %q\n", target.Src, res.Id)
	}
}

func TestBalanceList(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	target, err := c.Balance.List(nil)

	if err != nil {
		t.Error(err)
	}

	if len(target.Values) == 0 {
		t.Errorf("Balance list is empty\n")
	}
}
