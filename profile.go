//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

type ProfileBrandingIcon struct {
	Original string `json:"original"`
}
type ProfileBrandingLogo struct {
	Original string `json:"original"`
}
type ProfileBranding struct {
	Icon           *ProfileBrandingIcon `json:"icon"`
	Logo           *ProfileBrandingLogo `json:"logo"`
	PrimaryColor   string               `json:"primary_color"`
	SecondaryColor string               `json:"secondary_color"`
}

// A Stripe profile
type Profile struct {
	Branding    *ProfileBranding `json:"branding"`
	Description string           `json:"description"`
	DisplayName string           `json:"display_name"`
	ID          string           `json:"id"`
	Livemode    bool             `json:"livemode"`
	Object      string           `json:"object"`
	URL         string           `json:"url"`
	Username    string           `json:"username"`
}

// UnmarshalJSON handles deserialization of a Profile.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *Profile) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type profile Profile
	var v profile
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = Profile(v)
	return nil
}
