package stripe

import "encoding/json"

// InvoiceItemParams is the set of parameters that can be used when creating or updating an invoice item.
// For more details see https://stripe.com/docs/api#create_invoiceitem and https://stripe.com/docs/api#update_invoiceitem.
type InvoiceItemParams struct {
	Params         `form:"*"`
	Amount         int64    `form:"amount"`
	Currency       Currency `form:"currency"`
	Customer       string   `form:"customer"`
	Desc           string   `form:"description"`
	Discountable   bool     `form:"discountable"`
	Invoice        string   `form:"invoice"`
	NoDiscountable bool     `form:"discountable,invert"`
	Sub            string   `form:"subscription"`
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
	Amount       int64             `json:"amount"`
	Currency     Currency          `json:"currency"`
	Customer     *Customer         `json:"customer"`
	Date         int64             `json:"date"`
	Deleted      bool              `json:"deleted"`
	Desc         string            `json:"description"`
	Discountable bool              `json:"discountable"`
	ID           string            `json:"id"`
	Invoice      *Invoice          `json:"invoice"`
	Live         bool              `json:"livemode"`
	Meta         map[string]string `json:"metadata"`
	Period       *Period           `json:"period"`
	Plan         *Plan             `json:"plan"`
	Proration    bool              `json:"proration"`
	Quantity     int64             `json:"quantity"`
	Sub          string            `json:"subscription"`
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
