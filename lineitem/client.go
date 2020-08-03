// Package lineitem provides the tools needs to interact with the LineItem resource.
package lineitem

import (
	stripe "github.com/stripe/stripe-go/v71"
)

// Iter is an iterator for line items across various resources.
type Iter struct {
	*stripe.Iter
}

// LineItem returns the line item which the iterator is currently pointing to.
func (i *Iter) LineItem() *stripe.LineItem {
	return i.Current().(*stripe.LineItem)
}

// LineItemList returns the line item which the iterator is currently pointing to.
func (i *Iter) LineItemList() *stripe.LineItemList {
	return i.Current().(*stripe.LineItemList)
}
