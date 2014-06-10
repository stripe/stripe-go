package stripe

import (
	"testing"
)

func TestCustomerCreate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	customer := &CustomerParams{
		Balance: -123,
		Card: &CardParams{
			Name:   "Test Card",
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		Desc:  "Test Customer",
		Email: "a@b.com",
	}

	target, err := c.Customers.Create(customer)
	if err != nil {
		t.Error(err)
	}

	if target == nil {
		t.Errorf("No customer returned\n")
	}

	if target.Balance != customer.Balance {
		t.Errorf("Balance %v does not match expected balance %v\n", target.Balance, customer.Balance)
	}

	if target.Desc != customer.Desc {
		t.Errorf("Description %q does not match expected description %q\n", target.Desc, customer.Desc)
	}

	if target.Email != customer.Email {
		t.Errorf("Email %q does not match expected email %q\n", target.Email, customer.Email)
	}

	if target.Cards == nil {
		t.Errorf("No cards recorded\n")
	}

	if target.Cards.Count != 1 {
		t.Errorf("Unexpected number of cards %v\n", target.Cards.Count)
	}

	if target.Cards.Data[0].Name != customer.Card.Name {
		t.Errorf("Card name %q does not match expected name %q\n", target.Cards.Data[0].Name, customer.Card.Name)
	}

	c.Customers.Delete(target.Id)
}

func TestCustomerGet(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	customer := &CustomerParams{}

	res, _ := c.Customers.Create(customer)

	target, err := c.Customers.Get(res.Id)
	if err != nil {
		t.Error(err)
	}

	if target == nil {
		t.Errorf("No customer returned\n")
	}

	c.Customers.Delete(res.Id)
}

func TestCustomerDelete(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	customer := &CustomerParams{}

	res, _ := c.Customers.Create(customer)

	err := c.Customers.Delete(res.Id)
	if err != nil {
		t.Error(err)
	}
}

func TestCustomerUpdate(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	customer := &CustomerParams{
		Balance: 7,
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
		Desc:  "Original Desc",
		Email: "first@b.com",
	}

	original, _ := c.Customers.Create(customer)

	updated := &CustomerParams{
		Balance: -10,
		Card: &CardParams{
			Number: "4242424242424242",
			Month:  "10",
			Year:   "20",
		},
		Desc:  "Updated Desc",
		Email: "desc@b.com",
	}

	target, err := c.Customers.Update(original.Id, updated)
	if err != nil {
		t.Error(err)
	}

	if target == nil {
		t.Errorf("No customer returned\n")
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

	c.Customers.Delete(target.Id)
}
