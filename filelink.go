//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"github.com/stripe/stripe-go/v72/form"
)

// Retrieves the file link with the given ID.
type FileLinkParams struct {
	Params `form:"*"`
	// A future timestamp after which the link will no longer be usable, or `now` to expire the link immediately.
	ExpiresAt    *int64 `form:"expires_at"`
	ExpiresAtNow *bool  `form:"-"` // See custom AppendTo
	// The ID of the file. The file's `purpose` must be one of the following: `business_icon`, `business_logo`, `customer_signature`, `dispute_evidence`, `finance_report_run`, `identity_document_downloadable`, `pci_document`, `selfie`, `sigma_scheduled_query`, or `tax_document_user_upload`.
	File *string `form:"file"`
}

// AppendTo implements custom encoding logic for FileLinkParams.
func (f *FileLinkParams) AppendTo(body *form.Values, keyParts []string) {
	if BoolValue(f.ExpiresAtNow) {
		body.Add(form.FormatKey(append(keyParts, "expires_at")), "now")
	}
}

// Returns a list of file links.
type FileLinkListParams struct {
	ListParams   `form:"*"`
	Created      *int64            `form:"created"`
	CreatedRange *RangeQueryParams `form:"created"`
	// Filter links by their expiration status. By default, all links are returned.
	Expired *bool `form:"expired"`
	// Only return links for the given file.
	File *string `form:"file"`
}

// To share the contents of a `File` object with non-Stripe users, you can
// create a `FileLink`. `FileLink`s contain a URL that can be used to
// retrieve the contents of the file without authentication.
type FileLink struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Whether this link is already expired.
	Expired bool `json:"expired"`
	// Time at which the link expires.
	ExpiresAt int64 `json:"expires_at"`
	// The file object this link points to.
	File *File `json:"file"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The publicly accessible URL to download the file.
	URL string `json:"url"`
}

// UnmarshalJSON handles deserialization of a FileLink.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (f *FileLink) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		f.ID = id
		return nil
	}

	type fileLink FileLink
	var v fileLink
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*f = FileLink(v)
	return nil
}

// FileLinkList is a list of FileLinks as retrieved from a list endpoint.
type FileLinkList struct {
	APIResource
	ListMeta
	Data []*FileLink `json:"data"`
}
