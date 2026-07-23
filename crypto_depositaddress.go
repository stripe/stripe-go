//
//
// File generated from our OpenAPI spec
//
//

package stripe

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
	ListParams      `form:"*"`
	Address         *string `form:"address" json:"address,omitempty"`
	Customer        *string `form:"customer" json:"customer,omitempty"`
	CustomerAccount *string `form:"customer_account" json:"customer_account,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand  []*string `form:"expand" json:"expand,omitempty"`
	Network *string   `form:"network" json:"network,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoDepositAddressListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Creates a new crypto deposit address for the authenticated merchant on the specified network.
// The returned address can be used across multiple PaymentIntents.
type CryptoDepositAddressParams struct {
	Params   `form:"*"`
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand   []*string         `form:"expand" json:"expand,omitempty"`
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	Network  *string           `form:"network" json:"network,omitempty"`
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
	Params   `form:"*"`
	Customer *string `form:"customer" json:"customer,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand   []*string         `form:"expand" json:"expand,omitempty"`
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	Network  *string           `form:"network" json:"network"`
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

type CryptoDepositAddressSupportedToken struct {
	// The on-chain contract address for the supported token currency on this specific network.
	TokenContractAddress string `json:"token_contract_address"`
	// The supported token currency. Supported token currencies include: `usdc`.
	TokenCurrency CryptoDepositAddressSupportedTokenTokenCurrency `json:"token_currency"`
}

// A crypto deposit address is a blockchain address that can be used by a merchant for deposit mode crypto payments.
type CryptoDepositAddress struct {
	APIResource
	Address  string `json:"address"`
	Created  int64  `json:"created"`
	Customer string `json:"customer,omitempty"`
	// Unique identifier for the object.
	ID       string `json:"id"`
	Livemode bool   `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string           `json:"metadata"`
	Network  CryptoDepositAddressNetwork `json:"network"`
	// String representing the object's type. Objects of the same type share the same value.
	Object          string                                `json:"object"`
	SupportedTokens []*CryptoDepositAddressSupportedToken `json:"supported_tokens"`
}

// CryptoDepositAddressList is a list of DepositAddresses as retrieved from a list endpoint.
type CryptoDepositAddressList struct {
	APIResource
	ListMeta
	Data []*CryptoDepositAddress `json:"data"`
}
