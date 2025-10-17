//
//
// File generated from our OpenAPI spec
//
//

// Package blocklistentry provides the /v1/identity/blocklist_entries APIs
package blocklistentry

import (
	"net/http"

	stripe "github.com/stripe/stripe-go/v83"
	"github.com/stripe/stripe-go/v83/form"
)

// Client is used to invoke /v1/identity/blocklist_entries APIs.
// Deprecated: Use [stripe.Client] instead. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
type Client struct {
	B   stripe.Backend
	Key string
}

// Creates a BlocklistEntry object from a verification report.
//
// A blocklist entry prevents future identity verifications that match the same identity information.
// You can create blocklist entries from verification reports that contain document extracted data
// or a selfie.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#add-a-blocklist-entry)
func New(params *stripe.IdentityBlocklistEntryParams) (*stripe.IdentityBlocklistEntry, error) {
	return getC().New(params)
}

// Creates a BlocklistEntry object from a verification report.
//
// A blocklist entry prevents future identity verifications that match the same identity information.
// You can create blocklist entries from verification reports that contain document extracted data
// or a selfie.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#add-a-blocklist-entry)
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) New(params *stripe.IdentityBlocklistEntryParams) (*stripe.IdentityBlocklistEntry, error) {
	blocklistentry := &stripe.IdentityBlocklistEntry{}
	err := c.B.Call(
		http.MethodPost, "/v1/identity/blocklist_entries", c.Key, params, blocklistentry)
	return blocklistentry, err
}

// Retrieves a BlocklistEntry object by its identifier.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#block-list)
func Get(id string, params *stripe.IdentityBlocklistEntryParams) (*stripe.IdentityBlocklistEntry, error) {
	return getC().Get(id, params)
}

// Retrieves a BlocklistEntry object by its identifier.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#block-list)
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Get(id string, params *stripe.IdentityBlocklistEntryParams) (*stripe.IdentityBlocklistEntry, error) {
	path := stripe.FormatURLPath("/v1/identity/blocklist_entries/%s", id)
	blocklistentry := &stripe.IdentityBlocklistEntry{}
	err := c.B.Call(http.MethodGet, path, c.Key, params, blocklistentry)
	return blocklistentry, err
}

// Disables a BlocklistEntry object.
//
// After a BlocklistEntry is disabled, it will no longer block future verifications that match
// the same information. This action is irreversible. To re-enable it, a new BlocklistEntry
// must be created using the same verification report.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#disable-a-blocklist-entry)
func Disable(id string, params *stripe.IdentityBlocklistEntryDisableParams) (*stripe.IdentityBlocklistEntry, error) {
	return getC().Disable(id, params)
}

// Disables a BlocklistEntry object.
//
// After a BlocklistEntry is disabled, it will no longer block future verifications that match
// the same information. This action is irreversible. To re-enable it, a new BlocklistEntry
// must be created using the same verification report.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#disable-a-blocklist-entry)
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) Disable(id string, params *stripe.IdentityBlocklistEntryDisableParams) (*stripe.IdentityBlocklistEntry, error) {
	path := stripe.FormatURLPath("/v1/identity/blocklist_entries/%s/disable", id)
	blocklistentry := &stripe.IdentityBlocklistEntry{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, blocklistentry)
	return blocklistentry, err
}

// Returns a list of BlocklistEntry objects associated with your account.
//
// The blocklist entries are returned sorted by creation date, with the most recently created
// entries appearing first.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#block-list)
func List(params *stripe.IdentityBlocklistEntryListParams) *Iter {
	return getC().List(params)
}

// Returns a list of BlocklistEntry objects associated with your account.
//
// The blocklist entries are returned sorted by creation date, with the most recently created
// entries appearing first.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#block-list)
//
// Deprecated: Client methods are deprecated. This should be accessed instead through [stripe.Client]. See the [migration guide] for more info.
//
// [migration guide]: https://github.com/stripe/stripe-go/wiki/Migration-guide-for-Stripe-Client
func (c Client) List(listParams *stripe.IdentityBlocklistEntryListParams) *Iter {
	return &Iter{
		Iter: stripe.GetIter(listParams, func(p *stripe.Params, b *form.Values) ([]interface{}, stripe.ListContainer, error) {
			list := &stripe.IdentityBlocklistEntryList{}
			err := c.B.CallRaw(http.MethodGet, "/v1/identity/blocklist_entries", c.Key, []byte(b.Encode()), p, list)

			ret := make([]interface{}, len(list.Data))
			for i, v := range list.Data {
				ret[i] = v
			}

			return ret, list, err
		}),
	}
}

// Iter is an iterator for identity blocklist entries.
type Iter struct {
	*stripe.Iter
}

// IdentityBlocklistEntry returns the identity blocklist entry which the iterator is currently pointing to.
func (i *Iter) IdentityBlocklistEntry() *stripe.IdentityBlocklistEntry {
	return i.Current().(*stripe.IdentityBlocklistEntry)
}

// IdentityBlocklistEntryList returns the current list object which the iterator is
// currently using. List objects will change as new API calls are made to
// continue pagination.
func (i *Iter) IdentityBlocklistEntryList() *stripe.IdentityBlocklistEntryList {
	return i.List().(*stripe.IdentityBlocklistEntryList)
}

func getC() Client {
	return Client{stripe.GetBackend(stripe.APIBackend), stripe.Key}
}
