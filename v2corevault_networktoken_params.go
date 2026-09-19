//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Optional owner contact details used only when a network requires them for raw-card tokenization.
type V2CoreVaultNetworkTokenCardOwnerDetailsParams struct {
	// Cardholder email address.
	Email *string `form:"email" json:"email,omitempty"`
	// Cardholder phone number in international format, for example +15555550123.
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// Raw card values used to provision the network token.
type V2CoreVaultNetworkTokenCardParams struct {
	// The two-digit number representing the card's expiration month.
	ExpMonth *string `form:"exp_month" json:"exp_month"`
	// The four-digit number representing the card's expiration year.
	ExpYear *string `form:"exp_year" json:"exp_year"`
	// The card number.
	Number *string `form:"number" json:"number"`
	// The optional origin attestation for the card.
	Origin *string `form:"origin" json:"origin,omitempty"`
	// Optional owner contact details used only when a network requires them for raw-card tokenization.
	OwnerDetails *V2CoreVaultNetworkTokenCardOwnerDetailsParams `form:"owner_details" json:"owner_details,omitempty"`
}

// Create or Return a Network Token Using Raw Card Data.
type V2CoreVaultNetworkTokenParams struct {
	Params `form:"*"`
	// Raw card values used to provision the network token.
	Card *V2CoreVaultNetworkTokenCardParams `form:"card" json:"card,omitempty"`
	// Private preview supports card only.
	Type *string `form:"type" json:"type,omitempty"`
}

// The existing Stripe card reference to provision or resolve.
type V2CoreVaultNetworkTokenCreateFromCredentialCardParams struct {
	// The optional origin attestation for the referenced card.
	Origin *string `form:"origin" json:"origin,omitempty"`
	// A supported v2 Card ID or v1 PaymentMethod ID of type card.
	Reference *string `form:"reference" json:"reference"`
}

// Creates or returns a Network Token from an existing card reference.
type V2CoreVaultNetworkTokenCreateFromCredentialParams struct {
	Params `form:"*"`
	// The existing Stripe card reference to provision or resolve.
	Card *V2CoreVaultNetworkTokenCreateFromCredentialCardParams `form:"card" json:"card,omitempty"`
	// Private preview supports card only.
	Type *string `form:"type" json:"type"`
}

// Every successful call generates a new cryptogram, and retrying can generate another cryptogram.
// The cryptogram is returned only in this response and is never persisted.
type V2CoreVaultNetworkTokenGenerateCryptogramParams struct {
	Params `form:"*"`
	// The cryptogram type. When omitted, token_cryptogram is used.
	Type *string `form:"type" json:"type,omitempty"`
}

// Optional owner contact details used only when a network requires them for raw-card tokenization.
type V2CoreVaultNetworkTokenCreateCardOwnerDetailsParams struct {
	// Cardholder email address.
	Email *string `form:"email" json:"email,omitempty"`
	// Cardholder phone number in international format, for example +15555550123.
	Phone *string `form:"phone" json:"phone,omitempty"`
}

// Raw card values used to provision the network token.
type V2CoreVaultNetworkTokenCreateCardParams struct {
	// The two-digit number representing the card's expiration month.
	ExpMonth *string `form:"exp_month" json:"exp_month"`
	// The four-digit number representing the card's expiration year.
	ExpYear *string `form:"exp_year" json:"exp_year"`
	// The card number.
	Number *string `form:"number" json:"number"`
	// The optional origin attestation for the card.
	Origin *string `form:"origin" json:"origin,omitempty"`
	// Optional owner contact details used only when a network requires them for raw-card tokenization.
	OwnerDetails *V2CoreVaultNetworkTokenCreateCardOwnerDetailsParams `form:"owner_details" json:"owner_details,omitempty"`
}

// Create or Return a Network Token Using Raw Card Data.
type V2CoreVaultNetworkTokenCreateParams struct {
	Params `form:"*"`
	// Raw card values used to provision the network token.
	Card *V2CoreVaultNetworkTokenCreateCardParams `form:"card" json:"card,omitempty"`
	// Private preview supports card only.
	Type *string `form:"type" json:"type"`
}

// Retrieves an existing network token.
type V2CoreVaultNetworkTokenRetrieveParams struct {
	Params `form:"*"`
}
