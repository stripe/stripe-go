//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The blockchain network where this address can accept funds.
type CryptoDepositAddressNetwork string

// List of values that CryptoDepositAddressNetwork can take
const (
	CryptoDepositAddressNetworkBase   CryptoDepositAddressNetwork = "base"
	CryptoDepositAddressNetworkSolana CryptoDepositAddressNetwork = "solana"
	CryptoDepositAddressNetworkTempo  CryptoDepositAddressNetwork = "tempo"
)

// The supported token currency. Supported token currencies include: `usdc`.
type CryptoDepositAddressSupportedTokenTokenCurrency string

// List of values that CryptoDepositAddressSupportedTokenTokenCurrency can take
const (
	CryptoDepositAddressSupportedTokenTokenCurrencyUsdc CryptoDepositAddressSupportedTokenTokenCurrency = "usdc"
)

// Lists crypto deposit addresses for the authenticated merchant.
// Supports cursor-based pagination and optional filtering by customer, network, or on-chain address.
type CryptoDepositAddressListParams struct {
	ListParams `form:"*"`
	// Only return the deposit address matching this on-chain address.
	Address *string `form:"address" json:"address,omitempty"`
	// Only return deposit addresses scoped to this [Customer](https://docs.stripe.com/api/customers/object).
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Only return deposit addresses belonging to this customer account.
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Only return deposit addresses for this blockchain network.
	Network *string `form:"network" json:"network,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoDepositAddressListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Creates a new crypto deposit address for the authenticated merchant on the specified network.
// The returned address can be used across multiple PaymentIntents.
type CryptoDepositAddressParams struct {
	Params `form:"*"`
	// If set, this deposit address is scoped to a [Customer](https://docs.stripe.com/api/customers/object) and can only receive funds from that customer. Otherwise, this deposit address can receive funds from any customer.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The blockchain network to generate a deposit address for.
	Network *string `form:"network" json:"network,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoDepositAddressParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CryptoDepositAddressParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Creates a new crypto deposit address for the authenticated merchant on the specified network.
// The returned address can be used across multiple PaymentIntents.
type CryptoDepositAddressCreateParams struct {
	Params `form:"*"`
	// If set, this deposit address is scoped to a [Customer](https://docs.stripe.com/api/customers/object) and can only receive funds from that customer. Otherwise, this deposit address can receive funds from any customer.
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// The blockchain network to generate a deposit address for.
	Network *string `form:"network" json:"network"`
}

// AddExpand appends a new field to expand.
func (p *CryptoDepositAddressCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CryptoDepositAddressCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves the details of an existing crypto deposit address by ID.
type CryptoDepositAddressRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoDepositAddressRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// The tokens that can be sent to this deposit address on its network.
type CryptoDepositAddressSupportedToken struct {
	// The on-chain contract address for the supported token currency on this specific network.
	TokenContractAddress string `json:"token_contract_address"`
	// The supported token currency. Supported token currencies include: `usdc`.
	TokenCurrency CryptoDepositAddressSupportedTokenTokenCurrency `json:"token_currency"`
}

// A crypto deposit address is a blockchain address that can be used by a merchant for deposit mode crypto payments.
//
// Related guide: [Machine payments](https://docs.stripe.com/payments/machine)
type CryptoDepositAddress struct {
	APIResource
	// The on-chain address where funds can be received.
	Address string `json:"address"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// If set, this deposit address is scoped to a [Customer](https://docs.stripe.com/api/customers/object) and can only receive funds from that customer. Otherwise, this deposit address can receive funds from any customer.
	Customer string `json:"customer,omitempty"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// The blockchain network where this address can accept funds.
	Network CryptoDepositAddressNetwork `json:"network"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// The tokens that can be sent to this deposit address on its network.
	SupportedTokens []*CryptoDepositAddressSupportedToken `json:"supported_tokens"`
}

// CryptoDepositAddressList is a list of DepositAddresses as retrieved from a list endpoint.
type CryptoDepositAddressList struct {
	APIResource
	ListMeta
	Data []*CryptoDepositAddress `json:"data"`
}
