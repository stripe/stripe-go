//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stripe/stripe-go/v82/form"
)

// v1BankAccountService is used to invoke bankaccount related APIs.
type v1BankAccountService struct {
	B   Backend
	Key string
}

// Create creates a new bank account
func (c v1BankAccountService) Create(ctx context.Context, params *BankAccountCreateParams) (*BankAccount, error) {
	if params == nil {
		return nil, fmt.Errorf("params should not be nil")
	}

	var path string
	if (params.Account != nil && params.Customer != nil) || (params.Account == nil && params.Customer == nil) {
		return nil, fmt.Errorf("Invalid bank account params: exactly one of Account or Customer need to be set")
	} else if params.Account != nil {
		path = FormatURLPath("/v1/accounts/%s/external_accounts", StringValue(params.Account))
	} else if params.Customer != nil {
		path = FormatURLPath("/v1/customers/%s/sources", StringValue(params.Customer))
	}
	params.Context = ctx
	body := &form.Values{}

	// Note that we call this special append method instead of the standard one
	// from the form package. We should not use form's because doing so will
	// include some parameters that are undesirable here.
	if err := params.AppendToAsSourceOrExternalAccount(body); err != nil {
		return nil, err
	}

	// Because bank account creation uses the custom append above, we have to
	// make an explicit call using a form and CallRaw instead of the standard
	// Call (which takes a set of parameters).
	bankaccount := &BankAccount{}
	err := c.B.CallRaw(http.MethodPost, path, c.Key, []byte(body.Encode()), &params.Params, bankaccount)
	return bankaccount, err
}

// Get returns the details of a bank account.
func (c v1BankAccountService) Retrieve(ctx context.Context, id string, params *BankAccountRetrieveParams) (*BankAccount, error) {
	if params == nil {
		return nil, fmt.Errorf("params should not be nil")
	}
	var path string
	if (params.Account != nil && params.Customer != nil) || (params.Account == nil && params.Customer == nil) {
		return nil, fmt.Errorf("Invalid bank account params: exactly one of Account or Customer need to be set")
	} else if params.Account != nil {
		path = FormatURLPath("/v1/accounts/%s/external_accounts/%s", StringValue(params.Account), id)
	} else if params.Customer != nil {
		path = FormatURLPath("/v1/customers/%s/sources/%s", StringValue(params.Customer), id)
	}
	params.Context = ctx
	bankaccount := &BankAccount{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, bankaccount)
	return bankaccount, err
}

// Updates the metadata, account holder name, account holder type of a bank account belonging to
// a connected account and optionally sets it as the default for its currency. Other bank account
// details are not editable by design.
//
// You can only update bank accounts when [account.controller.requirement_collection is application, which includes <a href="/connect/custom-accounts">Custom accounts](https://docs.stripe.com/api/accounts/object#account_object-controller-requirement_collection).
//
// You can re-enable a disabled bank account by performing an update call without providing any
// arguments or changes.
func (c v1BankAccountService) Update(ctx context.Context, id string, params *BankAccountUpdateParams) (*BankAccount, error) {
	if params == nil {
		return nil, fmt.Errorf("params should not be nil")
	}
	var path string
	if (params.Account != nil && params.Customer != nil) || (params.Account == nil && params.Customer == nil) {
		return nil, fmt.Errorf("Invalid bank account params: exactly one of Account or Customer need to be set")
	} else if params.Account != nil {
		path = FormatURLPath("/v1/accounts/%s/external_accounts/%s", StringValue(params.Account), id)
	} else if params.Customer != nil {
		path = FormatURLPath("/v1/customers/%s/sources/%s", StringValue(params.Customer), id)
	}
	params.Context = ctx
	bankaccount := &BankAccount{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, bankaccount)
	return bankaccount, err
}

// Delete a specified external account for a given account.
func (c v1BankAccountService) Delete(ctx context.Context, id string, params *BankAccountDeleteParams) (*BankAccount, error) {
	if params == nil {
		return nil, fmt.Errorf("params should not be nil")
	}
	var path string
	if (params.Account != nil && params.Customer != nil) || (params.Account == nil && params.Customer == nil) {
		return nil, fmt.Errorf("Invalid bank account params: exactly one of Account or Customer need to be set")
	} else if params.Account != nil {
		path = FormatURLPath("/v1/accounts/%s/external_accounts/%s", StringValue(params.Account), id)
	} else if params.Customer != nil {
		path = FormatURLPath("/v1/customers/%s/sources/%s", StringValue(params.Customer), id)
	}
	params.Context = ctx
	bankaccount := &BankAccount{}
	err := c.B.Call(http.MethodDelete, path, c.Key, params, bankaccount)
	return bankaccount, err
}
func (c v1BankAccountService) List(ctx context.Context, listParams *BankAccountListParams) Seq2[*BankAccount, error] {
	var path string
	var outerErr error

	// There's no bank accounts list URL, so we use one sources or external
	// accounts. An override on BankAccountListParam's `AppendTo` will add the
	// filter `object=bank_account` to make sure that only bank accounts come
	// back with the response.
	if listParams == nil {
		outerErr = fmt.Errorf("params should not be nil")
	} else if (listParams.Account != nil && listParams.Customer != nil) || (listParams.Account == nil && listParams.Customer == nil) {
		outerErr = fmt.Errorf("Invalid bank account params: exactly one of Account or Customer need to be set")
	} else if listParams.Account != nil {
		path = FormatURLPath("/v1/accounts/%s/external_accounts",
			StringValue(listParams.Account))
	} else if listParams.Customer != nil {
		path = FormatURLPath("/v1/customers/%s/sources",
			StringValue(listParams.Customer))
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) ([]*BankAccount, ListContainer, error) {
		list := &BankAccountList{}

		if outerErr != nil {
			return nil, nil, outerErr
		}

		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, path, c.Key, []byte(b.Encode()), p, list)
		return list.Data, list, err
	}).All()
}
