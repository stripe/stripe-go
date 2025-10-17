//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"

	"github.com/stripe/stripe-go/v83/form"
)

// v1IdentityBlocklistEntryService is used to invoke /v1/identity/blocklist_entries APIs.
type v1IdentityBlocklistEntryService struct {
	B   Backend
	Key string
}

// Creates a BlocklistEntry object from a verification report.
//
// A blocklist entry prevents future identity verifications that match the same identity information.
// You can create blocklist entries from verification reports that contain document extracted data
// or a selfie.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#add-a-blocklist-entry)
func (c v1IdentityBlocklistEntryService) Create(ctx context.Context, params *IdentityBlocklistEntryCreateParams) (*IdentityBlocklistEntry, error) {
	if params == nil {
		params = &IdentityBlocklistEntryCreateParams{}
	}
	params.Context = ctx
	blocklistentry := &IdentityBlocklistEntry{}
	err := c.B.Call(
		http.MethodPost, "/v1/identity/blocklist_entries", c.Key, params, blocklistentry)
	return blocklistentry, err
}

// Retrieves a BlocklistEntry object by its identifier.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#block-list)
func (c v1IdentityBlocklistEntryService) Retrieve(ctx context.Context, id string, params *IdentityBlocklistEntryRetrieveParams) (*IdentityBlocklistEntry, error) {
	if params == nil {
		params = &IdentityBlocklistEntryRetrieveParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/identity/blocklist_entries/%s", id)
	blocklistentry := &IdentityBlocklistEntry{}
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
func (c v1IdentityBlocklistEntryService) Disable(ctx context.Context, id string, params *IdentityBlocklistEntryDisableParams) (*IdentityBlocklistEntry, error) {
	if params == nil {
		params = &IdentityBlocklistEntryDisableParams{}
	}
	params.Context = ctx
	path := FormatURLPath("/v1/identity/blocklist_entries/%s/disable", id)
	blocklistentry := &IdentityBlocklistEntry{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, blocklistentry)
	return blocklistentry, err
}

// Returns a list of BlocklistEntry objects associated with your account.
//
// The blocklist entries are returned sorted by creation date, with the most recently created
// entries appearing first.
//
// Related guide: [Identity Verification Blocklist](https://docs.stripe.com/docs/identity/review-tools#block-list)
func (c v1IdentityBlocklistEntryService) List(ctx context.Context, listParams *IdentityBlocklistEntryListParams) Seq2[*IdentityBlocklistEntry, error] {
	if listParams == nil {
		listParams = &IdentityBlocklistEntryListParams{}
	}
	listParams.Context = ctx
	return newV1List(listParams, func(p *Params, b *form.Values) (*v1Page[*IdentityBlocklistEntry], error) {
		list := &v1Page[*IdentityBlocklistEntry]{}
		if p == nil {
			p = &Params{}
		}
		p.Context = ctx
		err := c.B.CallRaw(http.MethodGet, "/v1/identity/blocklist_entries", c.Key, []byte(b.Encode()), p, list)
		return list, err
	}).All()
}
