//
//
// File generated from our OpenAPI spec
//
//

package stripe

// URL of the icon for the business. The image will be square and at least 128px x 128px.
type V2NetworkBusinessProfileBrandingIcon struct {
	// The URL of the image in its original size.
	Original string `json:"original"`
}

// URL of the logo for the business. The image will be at least 128px x 128px.
type V2NetworkBusinessProfileBrandingLogo struct {
	// The URL of the image in its original size.
	Original string `json:"original"`
}

// Branding data for the business.
type V2NetworkBusinessProfileBranding struct {
	// URL of the icon for the business. The image will be square and at least 128px x 128px.
	Icon *V2NetworkBusinessProfileBrandingIcon `json:"icon,omitempty"`
	// URL of the logo for the business. The image will be at least 128px x 128px.
	Logo *V2NetworkBusinessProfileBrandingLogo `json:"logo,omitempty"`
	// A CSS hex color value representing the primary branding color for this business.
	PrimaryColor string `json:"primary_color,omitempty"`
	// A CSS hex color value representing the secondary branding color for this business.
	SecondaryColor string `json:"secondary_color,omitempty"`
}

// The Stripe business profile represents a business' public identity on the Stripe network.
type V2NetworkBusinessProfile struct {
	APIResource
	// Branding data for the business.
	Branding *V2NetworkBusinessProfileBranding `json:"branding,omitempty"`
	// The description of the business.
	Description string `json:"description,omitempty"`
	// The display name of the Stripe business profile.
	DisplayName string `json:"display_name"`
	// The ID of the Stripe business profile; also known as the Network ID. This is the ID used to identify the business on the Stripe network.
	ID string `json:"id"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// The URL of the business.
	URL string `json:"url,omitempty"`
	// The username of the Stripe business profile.
	Username string `json:"username"`
}
