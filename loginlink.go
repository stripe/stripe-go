//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"encoding/json"
	"time"
)

// Creates a login link for a connected account to access the Express Dashboard.
//
// You can only create login links for accounts that use the [Express Dashboard](https://stripe.com/connect/express-dashboard) and are connected to your platform.
type LoginLinkParams struct {
	Params  `form:"*"`
	Account *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
}

// AddExpand appends a new field to expand.
func (p *LoginLinkParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Login Links are single-use URLs for a connected account to access the Express Dashboard. The connected account's [account.controller.stripe_dashboard.type](https://stripe.com/api/accounts/object#account_object-controller-stripe_dashboard-type) must be `express` to have access to the Express Dashboard.
type LoginLink struct {
	APIResource
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created time.Time `json:"created"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The URL for the login link.
	URL string `json:"url"`
}

// UnmarshalJSON handles deserialization of a LoginLink.
// This custom unmarshaling is needed to handle the time fields correctly.
func (l *LoginLink) UnmarshalJSON(data []byte) error {
	type loginLink LoginLink
	v := struct {
		Created int64 `json:"created"`
		*loginLink
	}{
		loginLink: (*loginLink)(l),
	}
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	l.Created = time.Unix(v.Created, 0)
	return nil
}

// MarshalJSON handles serialization of a LoginLink.
// This custom marshaling is needed to handle the time fields correctly.
func (l LoginLink) MarshalJSON() ([]byte, error) {
	type loginLink LoginLink
	v := struct {
		Created int64 `json:"created"`
		loginLink
	}{
		loginLink: (loginLink)(l),
		Created:   l.Created.Unix(),
	}
	b, err := json.Marshal(v)
	if err != nil {
		return nil, err
	}
	return b, err
}
