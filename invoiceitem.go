package stripe

import "encoding/json"

// InvoiceItemParams is the set of parameters that can be used when creating or updating an invoice item.
// For more details see https://stripe.com/docs/api#create_invoiceitem and https://stripe.com/docs/api#update_invoiceitem.
type InvoiceItemParams struct {
	Params
	Customer           string
	Amount             int64
	Currency           Currency
	Invoice, Desc, Sub string
}

// InvoiceItemListparams is the set of parameters that can be used when listing invoice items.
// For more details see https://stripe.com/docs/api#list_invoiceitems.
type InvoiceItemListParams struct {
	ListParams
	Created  int64
	Customer string
}

// InvoiceItem is the resource represneting a Stripe invoice item.
// For more details see https://stripe.com/docs/api#invoiceitems.
type InvoiceItem struct {
	Id        string            `json:"id"`
	Live      bool              `json:"livemode"`
	Amount    int64             `json:"amount"`
	Currency  Currency          `json:"currency"`
	Customer  *Customer         `json:"customer"`
	Date      int64             `json:"date"`
	Proration bool              `json:"proration"`
	Desc      string            `json:"description"`
	Invoice   *Invoice          `json:"invoice"`
	Meta      map[string]string `json:"metadata"`
	Sub       string            `json:"subscription"`
}

// InvoiceItemIter is a iterator for list responses.
type InvoiceItemIter struct {
	Iter *Iter
}

// Next returns the next value in the list.
func (i *InvoiceItemIter) Next() (*InvoiceItem, error) {
	ii, err := i.Iter.Next()
	if err != nil {
		return nil, err
	}

	return ii.(*InvoiceItem), err
}

// Stop returns true if there are no more iterations to be performed.
func (i *InvoiceItemIter) Stop() bool {
	return i.Iter.Stop()
}

// Meta returns the list metadata.
func (i *InvoiceItemIter) Meta() *ListMeta {
	return i.Iter.Meta()
}

func (i *InvoiceItem) UnmarshalJSON(data []byte) error {
	type invoiceitem InvoiceItem
	var ii invoiceitem
	err := json.Unmarshal(data, &ii)
	if err == nil {
		*i = InvoiceItem(ii)
	} else {
		// the id is surrounded by escaped \, so ignore those
		i.Id = string(data[1 : len(data)-1])
	}

	return nil
}
