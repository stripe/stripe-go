//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "time"

// Closed Enum. The status of the network token.
type V2CoreVaultNetworkTokenStatus string

// List of values that V2CoreVaultNetworkTokenStatus can take
const (
	V2CoreVaultNetworkTokenStatusActive      V2CoreVaultNetworkTokenStatus = "active"
	V2CoreVaultNetworkTokenStatusDeactivated V2CoreVaultNetworkTokenStatus = "deactivated"
	V2CoreVaultNetworkTokenStatusSuspended   V2CoreVaultNetworkTokenStatus = "suspended"
)

// This field is unset in create and retrieve responses. It is populated only after a successful generate_cryptogram request.
type V2CoreVaultNetworkTokenCryptogram struct {
	// The electronic commerce indicator associated with the cryptogram.
	Eci string `json:"eci,omitempty"`
	// The cryptogram type.
	Type string `json:"type"`
	// The cryptogram value.
	Value string `json:"value"`
}

// A NetworkToken object represents a network token provisioned for a card.
type V2CoreVaultNetworkToken struct {
	APIResource
	// Created timestamp.
	Created time.Time `json:"created"`
	// This field is unset in create and retrieve responses. It is populated only after a successful generate_cryptogram request.
	Cryptogram *V2CoreVaultNetworkTokenCryptogram `json:"cryptogram,omitempty"`
	// The month the network token expires.
	ExpMonth string `json:"exp_month,omitempty"`
	// The year the network token expires.
	ExpYear string `json:"exp_year,omitempty"`
	// ID of the NetworkToken object.
	ID string `json:"id"`
	// Whether the object exists in live mode or in test mode.
	Livemode bool `json:"livemode"`
	// The network token number.
	Number string `json:"number,omitempty"`
	// String representing the object's type. Objects of the same type share the same value of the object field.
	Object string `json:"object"`
	// Closed Enum. The status of the network token.
	Status V2CoreVaultNetworkTokenStatus `json:"status"`
}
