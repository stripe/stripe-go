package invoice

import (
	"testing"

	. "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/customer"
	"github.com/stripe/stripe-go/invoiceitem"
	. "github.com/stripe/stripe-go/utils"
)

func init() {
	Key = GetTestKey()
}

// Invoices are somewhat painful to test since you need
// to first have some items, so test everything together to
// avoid unnecessary duplication
func TestAllInvoicesScenarios(t *testing.T) {
	customerParams := &CustomerParams{
		Card: &CardParams{
			Number: "378282246310005",
			Month:  "06",
			Year:   "20",
		},
	}

	cust, _ := customer.New(customerParams)

	item := &InvoiceItemParams{
		Customer: cust.Id,
		Amount:   100,
		Currency: USD,
		Desc:     "Test Item",
	}

	targetItem, err := invoiceitem.New(item)

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

	invoiceParams := &InvoiceParams{
		Customer:  cust.Id,
		Desc:      "Desc",
		Statement: "Statement",
	}

	targetInvoice, err := New(invoiceParams)

	if err != nil {
		t.Error(err)
	}

	if targetInvoice.Customer.Id != invoiceParams.Customer {
		t.Errorf("Invoice customer %q does not match expected customer %q\n", targetInvoice.Customer.Id, invoiceParams.Customer)
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

	if targetInvoice.Desc != invoiceParams.Desc {
		t.Errorf("Invoice description %q does not match expected description %q\n", targetInvoice.Desc, invoiceParams.Desc)
	}

	if targetInvoice.Statement != invoiceParams.Statement {
		t.Errorf("Invoice statement %q does not match expected statement %q\n", targetInvoice.Statement, invoiceParams.Statement)
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

	targetItemUpdated, err := invoiceitem.Update(targetItem.Id, updatedItem)

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

	targetInvoiceUpdated, err := Update(targetInvoice.Id, updatedInvoice)

	if err != nil {
		t.Error(err)
	}

	if targetInvoiceUpdated.Desc != updatedInvoice.Desc {
		t.Errorf("Updated invoice description %q does not match expected description %q\n", targetInvoiceUpdated.Desc, updatedInvoice.Desc)
	}

	if targetInvoiceUpdated.Statement != updatedInvoice.Statement {
		t.Errorf("Updated invoice statement %q does not match expected statement %q\n", targetInvoiceUpdated.Statement, updatedInvoice.Statement)
	}

	_, err = invoiceitem.Get(targetItem.Id, nil)
	if err != nil {
		t.Error(err)
	}

	ii := invoiceitem.List(&InvoiceItemListParams{Customer: cust.Id})
	for !ii.Stop() {
		targetInvoiceItemList, err := ii.Next()

		if err != nil {
			t.Error(err)
		}

		if targetInvoiceItemList == nil {
			t.Error("No nil values expected")
		}

		if ii.Meta() == nil {
			t.Error("No metadata returned")
		}
	}

	i := List(&InvoiceListParams{Customer: cust.Id})
	for !i.Stop() {
		targetInvoiceList, err := i.Next()

		if err != nil {
			t.Error(err)
		}

		if targetInvoiceList == nil {
			t.Error("No nil values expected")
		}

		if i.Meta() == nil {
			t.Error("No metadata returned")
		}
	}

	il := ListLines(&InvoiceLineListParams{Id: targetInvoice.Id, Customer: cust.Id})
	for !il.Stop() {
		targetLineList, err := il.Next()

		if err != nil {
			t.Error(err)
		}

		if targetLineList == nil {
			t.Error("No nil values expected")
		}

		if il.Meta() == nil {
			t.Error("No metadata returned")
		}
	}

	err = invoiceitem.Del(targetItem.Id)

	if err != nil {
		t.Error(err)
	}

	_, err = Get(targetInvoice.Id, nil)

	if err != nil {
		t.Error(err)
	}

	customer.Del(cust.Id)
}
