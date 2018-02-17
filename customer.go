package stripe

import (
	"encoding/json"
)

// CustomerParams is the set of parameters that can be used when creating or updating a customer.
// For more details see https://stripe.com/docs/api#create_customer and https://stripe.com/docs/api#update_customer.
type CustomerParams struct {
	Params         `form:"*"`
	AccountBalance *int64                   `form:"account_balance"`
	BusinessVatID  string                   `form:"business_vat_id"`
	Coupon         *string                  `form:"coupon"`
	DefaultSource  string                   `form:"default_source"`
	Description    string                   `form:"description"`
	Email          string                   `form:"email"`
	Plan           string                   `form:"plan"`
	Quantity       uint64                   `form:"quantity"`
	Shipping       *CustomerShippingDetails `form:"shipping"`
	Source         *SourceParams            `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
	TaxPercent     *float64                 `form:"tax_percent"`
	Token          string                   `form:"-"` // This doesn't seem to be used?
	TrialEnd       int64                    `form:"trial_end"`
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
	AccountBalance int64                    `json:"account_balance"`
	BusinessVatID  string                   `json:"business_vat_id"`
	Created        int64                    `json:"created"`
	Currency       Currency                 `json:"currency"`
	DefaultSource  *PaymentSource           `json:"default_source"`
	Deleted        bool                     `json:"deleted"`
	Delinquent     bool                     `json:"delinquent"`
	Description    string                   `json:"description"`
	Discount       *Discount                `json:"discount"`
	Email          string                   `json:"email"`
	ID             string                   `json:"id"`
	Livemode       bool                     `json:"livemode"`
	Metadata       map[string]string        `json:"metadata"`
	Shipping       *CustomerShippingDetails `json:"shipping"`
	Sources        *SourceList              `json:"sources"`
	Subscriptions  *SubscriptionList        `json:"subscriptions"`
}

// CustomerList is a list of customers as retrieved from a list endpoint.
type CustomerList struct {
	ListMeta
	Data []*Customer `json:"data"`
}

// CustomerShippingDetails is the structure containing shipping information.
type CustomerShippingDetails struct {
	Address Address `json:"address" form:"address"`
	Name    string  `json:"name" form:"name"`
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
