// Package card provides the /cards APIs
package card

import (
	"errors"

	stripe "github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/form"
)

// Client is used to invoke /cards APIs.
type Client struct {
	B   stripe.Backend
	Key string
}

// New POSTs new cards either for a customer or recipient.
// For more details see https://stripe.com/docs/api#create_card.
func New(params *stripe.CardParams) (*stripe.Card, error) {
	return getC().New(params)
}

func (c Client) New(params *stripe.CardParams) (*stripe.Card, error) {
	if params == nil {
		return nil, errors.New("params should not be nil")
	}

	body := &form.Values{}

	// Note that we call this special append method instead of the standard one
	// from the form package. We should not use form's because doing so will
	// include some parameters that are undesirable here.
	params.AppendToAsCardSourceOrExternalAccount(body, nil)

	card := &stripe.Card{}
	var err error

	if params.Account != nil {
		err = c.B.Call("POST", stripe.FormatURLPath("/accounts/%s/external_accounts", stripe.StringValue(params.Account)), c.Key, body, &params.Params, card)
	} else if params.Customer != nil {
		err = c.B.Call("POST", stripe.FormatURLPath("/customers/%s/sources", stripe.StringValue(params.Customer)), c.Key, body, &params.Params, card)
	} else if params.Recipient != nil {
		err = c.B.Call("POST", stripe.FormatURLPath("/recipients/%s/cards", stripe.StringValue(params.Recipient)), c.Key, body, &params.Params, card)
	} else {
		err = errors.New("Invalid card params: either account, customer or recipient need to be set")
	}

	return card, err
}

// Get returns the details of a card.
// For more details see https://stripe.com/docs/api#retrieve_card.
func Get(id string, params *stripe.CardParams) (*stripe.Card, error) {
	return getC().Get(id, params)
}

func (c Client) Get(id string, params *stripe.CardParams) (*stripe.Card, error) {
	if params == nil {
		return nil, errors.New("params should not be nil")
	}

	var body *form.Values
	var commonParams *stripe.Params

	if params != nil {
		commonParams = &params.Params
		body = &form.Values{}
		form.AppendTo(body, params)
	}

	card := &stripe.Card{}
	var err error

	if params.Account != nil {
		err = c.B.Call("GET", stripe.FormatURLPath("/accounts/%s/external_accounts/%s", stripe.StringValue(params.Account), id), c.Key, body, commonParams, card)
	} else if params.Customer != nil {
		err = c.B.Call("GET", stripe.FormatURLPath("/customers/%s/sources/%s", stripe.StringValue(params.Customer), id), c.Key, body, commonParams, card)
	} else if params.Recipient != nil {
		err = c.B.Call("GET", stripe.FormatURLPath("/recipients/%s/cards/%s", stripe.StringValue(params.Recipient), id), c.Key, body, commonParams, card)
	} else {
		err = errors.New("Invalid card params: either account, customer or recipient need to be set")
	}

	return card, err
}

// Update updates a card's properties.
// For more details see	https://stripe.com/docs/api#update_card.
func Update(id string, params *stripe.CardParams) (*stripe.Card, error) {
	return getC().Update(id, params)
}

func (c Client) Update(id string, params *stripe.CardParams) (*stripe.Card, error) {
	if params == nil {
		return nil, errors.New("params should not be nil")
	}

	body := &form.Values{}
	form.AppendTo(body, params)

	card := &stripe.Card{}
	var err error

	if params.Account != nil {
		err = c.B.Call("POST", stripe.FormatURLPath("/accounts/%s/external_accounts/%s", stripe.StringValue(params.Account), id), c.Key, body, &params.Params, card)
	} else if params.Customer != nil {
		err = c.B.Call("POST", stripe.FormatURLPath("/customers/%s/sources/%s", stripe.StringValue(params.Customer), id), c.Key, body, &params.Params, card)
	} else if params.Recipient != nil {
		err = c.B.Call("POST", stripe.FormatURLPath("/recipients/%s/cards/%s", stripe.StringValue(params.Recipient), id), c.Key, body, &params.Params, card)
	} else {
		err = errors.New("Invalid card params: either account, customer or recipient need to be set")
	}

	return card, err
}

// Del removes a card.
// For more details see https://stripe.com/docs/api#delete_card.
func Del(id string, params *stripe.CardParams) (*stripe.Card, error) {
	return getC().Del(id, params)
}

func (c Client) Del(id string, params *stripe.CardParams) (*stripe.Card, error) {
	if params == nil {
		return nil, errors.New("params should not be nil")
	}

	var body *form.Values
	var commonParams *stripe.Params

	body = &form.Values{}
	form.AppendTo(body, params)
	commonParams = &params.Params

	card := &stripe.Card{}
	var err error

	if params.Account != nil {
		err = c.B.Call("DELETE", stripe.FormatURLPath("/accounts/%s/external_accounts/%s", stripe.StringValue(params.Account), id), c.Key, body, commonParams, card)
	} else if params.Customer != nil {
		err = c.B.Call("DELETE", stripe.FormatURLPath("/customers/%s/sources/%s", stripe.StringValue(params.Customer), id), c.Key, body, commonParams, card)
	} else if params.Recipient != nil {
		err = c.B.Call("DELETE", stripe.FormatURLPath("/recipients/%s/cards/%s", stripe.StringValue(params.Recipient), id), c.Key, body, commonParams, card)
	} else {
		err = errors.New("Invalid card params: either account, customer or recipient need to be set")
	}

	return card, err
}

// List returns a list of cards.
// For more details see https://stripe.com/docs/api#list_cards.
func List(params *stripe.CardListParams) *Iter {
	return getC().List(params)
}

func (c Client) List(params *stripe.CardListParams) *Iter {
	body := &form.Values{}
	var lp *stripe.ListParams
	var p *stripe.Params

	if params != nil {
		form.AppendTo(body, params)
		lp = &params.ListParams
		p = params.ToParams()
	}

	return &Iter{stripe.GetIter(lp, body, func(b *form.Values) ([]interface{}, stripe.ListMeta, error) {
		list := &stripe.CardList{}
		var err error

		if params == nil {
			return nil, list.ListMeta, errors.New("params should not be nil")
		}

		if params.Account != nil {
			err = c.B.Call("GET", stripe.FormatURLPath("/accounts/%s/external_accounts?object=card", stripe.StringValue(params.Account)), c.Key, b, p, list)
		} else if params.Customer != nil {
			err = c.B.Call("GET", stripe.FormatURLPath("/customers/%s/sources?object=card", stripe.StringValue(params.Customer)), c.Key, b, p, list)
		} else if params.Recipient != nil {
			err = c.B.Call("GET", stripe.FormatURLPath("/recipients/%s/cards", stripe.StringValue(params.Recipient)), c.Key, b, p, list)
		} else {
			err = errors.New("Invalid card params: either account, customer or recipient need to be set")
		}

		ret := make([]interface{}, len(list.Data))
		for i, v := range list.Data {
			ret[i] = v
		}

		return ret, list.ListMeta, err
	})}
}

// Iter is an iterator for lists of Cards.
// The embedded Iter carries methods with it;
// see its documentation for details.
type Iter struct {
	*stripe.Iter
}

// Card returns the most recent Card
// visited by a call to Next.
func (i *Iter) Card() *stripe.Card {
	return i.Current().(*stripe.Card)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
