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
	Params       `form:"*"`
	ExpiresAt    *int64  `form:"expires_at"`
	ExpiresAtNow *bool   `form:"-"` // See custom AppendTo
	File         *string `form:"file"`
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
	Expired      *bool             `form:"expired"`
	File         *string           `form:"file"`
}

// To share the contents of a `File` object with non-Stripe users, you can
// create a `FileLink`. `FileLink`s contain a URL that can be used to
// retrieve the contents of the file without authentication.
type FileLink struct {
	APIResource
	Created   int64             `json:"created"`
	Expired   bool              `json:"expired"`
	ExpiresAt int64             `json:"expires_at"`
	File      *File             `json:"file"`
	ID        string            `json:"id"`
	Livemode  bool              `json:"livemode"`
	Metadata  map[string]string `json:"metadata"`
	Object    string            `json:"object"`
	URL       string            `json:"url"`
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
