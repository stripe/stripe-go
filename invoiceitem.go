package stripe

import "encoding/json"

// InvoiceItemParams is the set of parameters that can be used when creating or updating an invoice item.
// For more details see https://stripe.com/docs/api#create_invoiceitem and https://stripe.com/docs/api#update_invoiceitem.
type InvoiceItemParams struct {
	Params         `form:"*"`
	Customer       string   `form:"customer"`
	Amount         int64    `form:"amount"`
	Currency       Currency `form:"currency"`
	Invoice        string   `form:"invoice"`
	Desc           string   `form:"description"`
	Sub            string   `form:"subscription"`
	Discountable   bool     `form:"discountable"`
	NoDiscountable bool     `form:"discountable,invert"`
}

// InvoiceItemListParams is the set of parameters that can be used when listing invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
type InvoiceItemListParams struct {
	ListParams   `form:"*"`
	Created      int64             `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	Customer     string            `form:"customer"`
}

// InvoiceItem is the resource represneting a Stripe invoice item.
// For more details see https://stripe.com/docs/api#invoiceitems.
type InvoiceItem struct {
	ID           string            `json:"id"`
	Live         bool              `json:"livemode"`
	Amount       int64             `json:"amount"`
	Currency     Currency          `json:"currency"`
	Customer     *Customer         `json:"customer"`
	Date         int64             `json:"date"`
	Period       *Period           `json:"period"`
	Plan         *Plan             `json:"plan"`
	Proration    bool              `json:"proration"`
	Desc         string            `json:"description"`
	Invoice      *Invoice          `json:"invoice"`
	Meta         map[string]string `json:"metadata"`
	Sub          string            `json:"subscription"`
	Discountable bool              `json:"discountable"`
	Deleted      bool              `json:"deleted"`
	Quantity     int64             `json:"quantity"`
}

// InvoiceItemList is a list of invoice items as retrieved from a list endpoint.
type InvoiceItemList struct {
	ListMeta
	Values []*InvoiceItem `json:"data"`
}

// UnmarshalJSON handles deserialization of an InvoiceItem.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (i *InvoiceItem) UnmarshalJSON(data []byte) error {
	type invoiceitem InvoiceItem
	var ii invoiceitem
	err := json.Unmarshal(data, &ii)
	if err == nil {
		*i = InvoiceItem(ii)
	} else {
		// the id is surrounded by "\" characters, so strip them
		i.ID = string(data[1 : len(data)-1])
	}

	return nil
}
