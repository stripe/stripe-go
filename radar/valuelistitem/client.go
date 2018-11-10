// Package valuelist provides API functions related to value list items.
//
// For more details, see: https://stripe.com/docs/api/radar/list_items?lang=go
package valuelistitem

import (
	"net/http"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /radar/value_list_items APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New creates a new value list item.
func New(params *stripe.ValueListItemParams) (*stripe.ValueListItem, error) {
	return getC().New(params)
}

// New creates a new value list item.
func (c Client) New(params *stripe.ValueListItemParams) (*stripe.ValueListItem, error) {
	vli := &stripe.ValueListItem{}
	err := c.B.Call(http.MethodPost, "/radar/value_list_items", c.Key, params, vli)
	return vli, err
}

// Get returns the details of a value list item.
func Get(id string, params *stripe.ValueListItemParams) (*stripe.ValueListItem, error) {
	return getC().Get(id, params)
}

// Get returns the details of a value list item.
func (c Client) Get(id string, params *stripe.ValueListItemParams) (*stripe.ValueListItem, error) {
	path := stripe.FormatURLPath("/radar/value_list_items/%s", id)
	vli := &stripe.ValueListItem{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, vli)
	return vli, err
}

// Del removes a value list item.
func Del(id string, params *stripe.ValueListItemParams) (*stripe.ValueListItem, error) {
	return getC().Del(id, params)
}

// Del removes a value list item.
func (c Client) Del(id string, params *stripe.ValueListItemParams) (*stripe.ValueListItem, error) {
	path := stripe.FormatURLPath("/radar/value_list_items/%s", id)
	vli := &stripe.ValueListItem{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, vli)
	return vli, err
}

// List returns a list of vlis.
func List(params *stripe.ValueListItemListParams) *Iter {
	return getC().List(params)
}

// List returns a list of vlis.
func (c Client) List(listParams *stripe.ValueListItemListParams) *Iter {
	return &Iter{stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.ValueListItemList{}
		err := c.B.CallRaw(http.MethodGet, "/radar/value_list_items", c.Key, b, p, list)

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for vlis.
type Iter struct {
	*stripe.Iter
}

// ValueListItem returns the vli which the iterator is currently pointing to.
func (i *Iter) ValueListItem() *stripe.ValueListItem {
	return i.Current().(*stripe.ValueListItem)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
