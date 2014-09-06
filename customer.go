package stripe

import "encoding/json"

// CustomerParams is the set of parameters that can be used when creating or updating a customer.
// For more details see https://stripe.com/docs/api#create_customer and https://stripe.com/docs/api#update_customer.
type CustomerParams struct {
	Params
	Balance       int64
	Token, Coupon string
	Card          *CardParams
	Desc, Email   string
	Plan          string
	Quantity      uint64
	TrialEnd      int64
	DefaultCard   string
}

// CustomerListParams is the set of parameters that can be used when listing customers.
// For more details see https://stripe.com/docs/api#list_customers.
type CustomerListParams struct {
	ListParams
	Created int64
}

// Customer is the resource representing a Stripe customer.
// For more details see https://stripe.com/docs/api#customers.
type Customer struct {
	Id          string            `json:"id"`
	Live        bool              `json:"livemode"`
	Cards       *CardList         `json:"cards"`
	Created     int64             `json:"created"`
	Balance     int64             `json:"account_balance"`
	Currency    Currency          `json:"currency"`
	DefaultCard *Card             `json:"default_card"`
	Delinquent  bool              `json:"delinquent"`
	Desc        string            `json:"description"`
	Discount    *Discount         `json:"discount"`
	Email       string            `json:"email"`
	Meta        map[string]string `json:"metadata"`
	Subs        *SubList          `json:"subscriptions"`
}

// CustomerIter is a iterator for list responses.
type CustomerIter struct {
	Iter *Iter
}

// Next returns the next value in the list.
func (i *CustomerIter) Next() (*Customer, error) {
	c, err := i.Iter.Next()
	if err != nil {
		return nil, err
	}

	return c.(*Customer), err
}

// Stop returns true if there are no more iterations to be performed.
func (i *CustomerIter) Stop() bool {
	return i.Iter.Stop()
}

// Meta returns the list metadata.
func (i *CustomerIter) Meta() *ListMeta {
	return i.Iter.Meta()
}
func (c *Customer) UnmarshalJSON(data []byte) error {
	type customer Customer
	var cc customer
	err := json.Unmarshal(data, &cc)
	if err == nil {
		*c = Customer(cc)
	} else {
		c.Id = string(data[1 : len(data)-1])
	}

	return nil
}
