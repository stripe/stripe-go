//
//
// File generated from our OpenAPI spec
//
//

// Package lineitem provides the /checkout/sessions/{session}/line_items APIs
package lineitem

import (
	stripe "github.com/stripe/stripe-go/v72"
)

// Iter is an iterator for line items.
type Iter struct {
	*stripe.Iter
}

// LineItem returns the line item which the iterator is currently pointing to.
func (i *Iter) LineItem() *stripe.LineItem {
	return i.Current().(*stripe.LineItem)
}

// LineItemList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) LineItemList() *stripe.LineItemList {
	return i.List().(*stripe.LineItemList)
}
