//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Profile icon image.
type ProfileBrandingIcon struct {
	// The original image.
	Original string `json:"original"`
}

// Profile logo image.
type ProfileBrandingLogo struct {
	// The original image.
	Original string `json:"original"`
}

// Branding information for the Stripe profile.
type ProfileBranding struct {
	// Profile icon image.
	Icon *ProfileBrandingIcon `json:"icon"`
	// Profile logo image.
	Logo *ProfileBrandingLogo `json:"logo"`
	// The primary brand color for the profile.
	PrimaryColor string `json:"primary_color"`
	// The secondary brand color for the profile.
	SecondaryColor string `json:"secondary_color"`
}

// A Stripe profile
type Profile struct {
	// Branding information for the Stripe profile.
	Branding *ProfileBranding `json:"branding"`
	// A description of the business or entity represented by the Stripe profile.
	Description string `json:"description"`
	// The display name shown for the Stripe profile.
	DisplayName string `json:"display_name"`
	// Unique identifier for the Stripe profile.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The external website URL associated with the Stripe profile.
	URL string `json:"url"`
	// The unique username for the Stripe profile.
	Username string `json:"username"`
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
