package stripe

import (
	"encoding/json"
)

// CustomerParams is the set of parameters that can be used when creating or updating a customer.
// For more details see https://stripe.com/docs/api#create_customer and https://stripe.com/docs/api#update_customer.
type CustomerParams struct {
	Params         `form:"*"`
	Balance        int64                    `form:"account_balance"`
	BalanceZero    bool                     `form:"account_balance,zero"`
	Token          string                   `form:"-"` // This doesn't seem to be used?
	Coupon         string                   `form:"coupon"`
	CouponEmpty    bool                     `form:"coupon,empty"`
	Source         *SourceParams            `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
	Desc           string                   `form:"description"`
	Email          string                   `form:"email"`
	Plan           string                   `form:"plan"`
	Quantity       uint64                   `form:"quantity"`
	TrialEnd       int64                    `form:"trial_end"`
	DefaultSource  string                   `form:"default_source"`
	Shipping       *CustomerShippingDetails `form:"shipping"`
	BusinessVatID  string                   `form:"business_vat_id"`
	TaxPercent     float64                  `form:"tax_percent"`
	TaxPercentZero bool                     `form:"tax_percent,zero"`
}

// SetSource adds valid sources to a CustomerParams object,
// returning an error for unsupported sources.
func (cp *CustomerParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	cp.Source = source
	return err
}

// CustomerListParams is the set of parameters that can be used when listing customers.
// For more details see https://stripe.com/docs/api#list_customers.
type CustomerListParams struct {
	ListParams   `form:"*"`
	Created      int64             `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
}

// Customer is the resource representing a Stripe customer.
// For more details see https://stripe.com/docs/api#customers.
type Customer struct {
	ID            string                   `json:"id"`
	Live          bool                     `json:"livemode"`
	Sources       *SourceList              `json:"sources"`
	Created       int64                    `json:"created"`
	Balance       int64                    `json:"account_balance"`
	Currency      Currency                 `json:"currency"`
	DefaultSource *PaymentSource           `json:"default_source"`
	Delinquent    bool                     `json:"delinquent"`
	Desc          string                   `json:"description"`
	Discount      *Discount                `json:"discount"`
	Email         string                   `json:"email"`
	Meta          map[string]string        `json:"metadata"`
	Subs          *SubList                 `json:"subscriptions"`
	Deleted       bool                     `json:"deleted"`
	Shipping      *CustomerShippingDetails `json:"shipping"`
	BusinessVatID string                   `json:"business_vat_id"`
}

// CustomerList is a list of customers as retrieved from a list endpoint.
type CustomerList struct {
	ListMeta
	Values []*Customer `json:"data"`
}

// CustomerShippingDetails is the structure containing shipping information.
type CustomerShippingDetails struct {
	Name    string  `json:"name" form:"name"`
	Address Address `json:"address" form:"address"`
	Phone   string  `json:"phone" form:"phone"`
}

// UnmarshalJSON handles deserialization of a Customer.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Customer) UnmarshalJSON(data []byte) error {
	type customer Customer
	var cc customer
	err := json.Unmarshal(data, &cc)
	if err == nil {
		*c = Customer(cc)
	} else {
		// the id is surrounded by "\" characters, so strip them
		c.ID = string(data[1 : len(data)-1])
	}

	return nil
}
