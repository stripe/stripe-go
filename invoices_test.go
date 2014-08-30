package stripe

import (
	"testing"
)

// Invoices are somewhat painful to test since you need
// to first have some items, so test everything together to
// avoid unnecessary duplication
func TestAllInvoicesScenarios(t *testing.T) {
	c := &Client{}
	c.Init(key, nil, nil)

	customer := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := c.Customers.Create(customer)

	item := &InvoiceItemParams{
		Customer: cust.Id,
		Amount:   100,
		Currency: USD,
		Desc:     "Test Item",
	}

	targetItem, err := c.InvoiceItems.Create(item)

	if err != nil {
		t.Error(err)
	}

	if targetItem.Customer.Id != item.Customer {
		t.Errorf("Item customer %q does not match expected customer %q\n", targetItem.Customer.Id, item.Customer)
	}

	if targetItem.Desc != item.Desc {
		t.Errorf("Item description %q does not match expected description %q\n", targetItem.Desc, item.Desc)
	}

	if targetItem.Amount != item.Amount {
		t.Errorf("Item amount %v does not match expected amount %v\n", targetItem.Amount, item.Amount)
	}

	if targetItem.Currency != item.Currency {
		t.Errorf("Item currency %q does not match expected currency %q\n", targetItem.Currency, item.Currency)
	}

	if targetItem.Date == 0 {
		t.Errorf("Item date is not set\n")
	}

	invoice := &InvoiceParams{
		Customer:  cust.Id,
		Desc:      "Desc",
		Statement: "Statement",
	}

	targetInvoice, err := c.Invoices.Create(invoice)

	if err != nil {
		t.Error(err)
	}

	if targetInvoice.Customer.Id != invoice.Customer {
		t.Errorf("Invoice customer %q does not match expected customer %q\n", targetInvoice.Customer.Id, invoice.Customer)
	}

	if targetInvoice.Amount != targetItem.Amount {
		t.Errorf("Invoice amount %v does not match expected amount %v\n", targetInvoice.Amount, targetItem.Amount)
	}

	if targetInvoice.Currency != targetItem.Currency {
		t.Errorf("Invoice currency %q does not match expected currency %q\n", targetInvoice.Currency, targetItem.Currency)
	}

	if targetInvoice.Date == 0 {
		t.Errorf("Invoice date is not set\n")
	}

	if targetInvoice.Start == 0 {
		t.Errorf("Invoice start is not set\n")
	}

	if targetInvoice.End == 0 {
		t.Errorf("Invoice end is not set\n")
	}

	if targetInvoice.Total != targetInvoice.Amount || targetInvoice.Subtotal != targetInvoice.Amount {
		t.Errorf("Invoice total %v and subtotal %v do not match expected amount %v\n", targetInvoice.Total, targetInvoice.Subtotal, targetInvoice.Amount)
	}

	if targetInvoice.Desc != invoice.Desc {
		t.Errorf("Invoice description %q does not match expected description %q\n", targetInvoice.Desc, invoice.Desc)
	}

	if targetInvoice.Statement != invoice.Statement {
		t.Errorf("Invoice statement %q does not match expected statement %q\n", targetInvoice.Statement, invoice.Statement)
	}

	if targetInvoice.Lines == nil {
		t.Errorf("Invoice lines not found\n")
	}

	if targetInvoice.Lines.Count != 1 {
		t.Errorf("Invoice lines count %v does not match expected value\n", targetInvoice.Lines.Count)
	}

	if targetInvoice.Lines.Values == nil {
		t.Errorf("Invoice lines values not found\n")
	}

	if targetInvoice.Lines.Values[0].Amount != targetItem.Amount {
		t.Errorf("Invoice line amount %v does not match expected amount %v\n", targetInvoice.Lines.Values[0].Amount, targetItem.Amount)
	}

	if targetInvoice.Lines.Values[0].Currency != targetItem.Currency {
		t.Errorf("Invoice line currency %q does not match expected currency %q\n", targetInvoice.Lines.Values[0].Currency, targetItem.Currency)
	}

	if targetInvoice.Lines.Values[0].Desc != targetItem.Desc {
		t.Errorf("Invoice line description %q does not match expected description %q\n", targetInvoice.Lines.Values[0].Desc, targetItem.Desc)
	}

	if targetInvoice.Lines.Values[0].Type != TypeInvoiceItem {
		t.Errorf("Invoice line type %q does not match expected type\n", targetInvoice.Lines.Values[0].Type)
	}

	if targetInvoice.Lines.Values[0].Period == nil {
		t.Errorf("Invoice line period not found\n")
	}

	if targetInvoice.Lines.Values[0].Period.Start == 0 {
		t.Errorf("Invoice line period start is not set\n")
	}

	if targetInvoice.Lines.Values[0].Period.End == 0 {
		t.Errorf("Invoice line period end is not set\n")
	}

	updatedItem := &InvoiceItemParams{
		Amount: 99,
		Desc:   "Updated Desc",
	}

	targetItemUpdated, err := c.InvoiceItems.Update(targetItem.Id, updatedItem)

	if err != nil {
		t.Error(err)
	}

	if targetItemUpdated.Desc != updatedItem.Desc {
		t.Errorf("Updated item description %q does not match expected description %q\n", targetItemUpdated.Desc, updatedItem.Desc)
	}

	if targetItemUpdated.Amount != updatedItem.Amount {
		t.Errorf("Updated item amount %v does not match expected amount %v\n", targetItemUpdated.Amount, updatedItem.Amount)
	}

	updatedInvoice := &InvoiceParams{
		Desc:      "Updated Desc",
		Statement: "Updated",
	}

	targetInvoiceUpdated, err := c.Invoices.Update(targetInvoice.Id, updatedInvoice)

	if err != nil {
		t.Error(err)
	}

	if targetInvoiceUpdated.Desc != updatedInvoice.Desc {
		t.Errorf("Updated invoice description %q does not match expected description %q\n", targetInvoiceUpdated.Desc, updatedInvoice.Desc)
	}

	if targetInvoiceUpdated.Statement != updatedInvoice.Statement {
		t.Errorf("Updated invoice statement %q does not match expected statement %q\n", targetInvoiceUpdated.Statement, updatedInvoice.Statement)
	}

	_, err = c.InvoiceItems.Get(targetItem.Id)
	if err != nil {
		t.Error(err)
	}

	targetInvoiceItemList, err := c.InvoiceItems.List(&InvoiceItemListParams{Customer: cust.Id})

	if err != nil {
		t.Error(err)
	}

	if len(targetInvoiceItemList.Values) == 0 {
		t.Errorf("Invoice item list is empty\n")
	}

	targetInvoiceList, err := c.Invoices.List(&InvoiceListParams{Customer: cust.Id})

	if err != nil {
		t.Error(err)
	}

	if len(targetInvoiceList.Values) == 0 {
		t.Errorf("Invoice list is empty\n")
	}

	targetInvoiceLineList, err := c.Invoices.ListLines(&InvoiceLineListParams{Id: targetInvoice.Id, Customer: cust.Id})

	if err != nil {
		t.Error(err)
	}

	if len(targetInvoiceLineList.Values) == 0 {
		t.Errorf("Invoice line list is empty\n")
	}

	err = c.InvoiceItems.Delete(targetItem.Id)

	if err != nil {
		t.Error(err)
	}

	_, err = c.Invoices.Get(targetInvoice.Id)

	if err != nil {
		t.Error(err)
	}

	c.Customers.Delete(cust.Id)
}
