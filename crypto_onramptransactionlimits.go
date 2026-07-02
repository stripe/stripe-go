//
//
// File generated from our OpenAPI spec
//
//

package stripe

// Retrieves the remaining onramp limit for a crypto customer.
type CryptoOnrampTransactionLimitsParams struct {
	Params `form:"*"`
	// The IP address of the customer requesting transaction limits. We support IPv4 and IPv6 addresses.
	CustomerIPAddress *string `form:"customer_ip_address" json:"customer_ip_address,omitempty"`
	// The destination blockchain network to use for limit calculations.
	DestinationNetwork *string `form:"destination_network" json:"destination_network,omitempty"`
	// The destination tag for the wallet address, if applicable for the network.
	DestinationTag *string `form:"destination_tag" json:"destination_tag,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The wallet address to use for destination-specific limit calculations.
	WalletAddress *string `form:"wallet_address" json:"wallet_address,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoOnrampTransactionLimitsParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves the remaining onramp limit for a crypto customer.
type CryptoOnrampTransactionLimitsRetrieveParams struct {
	Params `form:"*"`
	// The IP address of the customer requesting transaction limits. We support IPv4 and IPv6 addresses.
	CustomerIPAddress *string `form:"customer_ip_address" json:"customer_ip_address,omitempty"`
	// The destination blockchain network to use for limit calculations.
	DestinationNetwork *string `form:"destination_network" json:"destination_network,omitempty"`
	// The destination tag for the wallet address, if applicable for the network.
	DestinationTag *string `form:"destination_tag" json:"destination_tag,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The wallet address to use for destination-specific limit calculations.
	WalletAddress *string `form:"wallet_address" json:"wallet_address,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoOnrampTransactionLimitsRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The remaining onramp limit for the crypto customer, separated by currency, payment method, and settlement speed.
//
// Limits are shown for currencies that correspond to the regions where the customer previously transacted. If the customer has no prior transactions, we return limits for all supported currencies.
type CryptoOnrampTransactionLimitsLimits struct{}

// This object represents the limit for the remaining amount that the crypto customer can onramp.
type CryptoOnrampTransactionLimits struct {
	APIResource
	// The ID of the crypto customer.
	CryptoCustomerID string `json:"crypto_customer_id"`
	// The remaining onramp limit for the crypto customer, separated by currency, payment method, and settlement speed.
	//
	// Limits are shown for currencies that correspond to the regions where the customer previously transacted. If the customer has no prior transactions, we return limits for all supported currencies.
	Limits *CryptoOnrampTransactionLimitsLimits `json:"limits"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
}
