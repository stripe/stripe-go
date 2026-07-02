//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The blockchain network for this wallet
type CryptoCustomerConsumerWalletNetwork string

// List of values that CryptoCustomerConsumerWalletNetwork can take
const (
	CryptoCustomerConsumerWalletNetworkAptos      CryptoCustomerConsumerWalletNetwork = "aptos"
	CryptoCustomerConsumerWalletNetworkAvalanche  CryptoCustomerConsumerWalletNetwork = "avalanche"
	CryptoCustomerConsumerWalletNetworkBase       CryptoCustomerConsumerWalletNetwork = "base"
	CryptoCustomerConsumerWalletNetworkBitcoin    CryptoCustomerConsumerWalletNetwork = "bitcoin"
	CryptoCustomerConsumerWalletNetworkEthereum   CryptoCustomerConsumerWalletNetwork = "ethereum"
	CryptoCustomerConsumerWalletNetworkOptimism   CryptoCustomerConsumerWalletNetwork = "optimism"
	CryptoCustomerConsumerWalletNetworkPolygon    CryptoCustomerConsumerWalletNetwork = "polygon"
	CryptoCustomerConsumerWalletNetworkSolana     CryptoCustomerConsumerWalletNetwork = "solana"
	CryptoCustomerConsumerWalletNetworkStellar    CryptoCustomerConsumerWalletNetwork = "stellar"
	CryptoCustomerConsumerWalletNetworkWorldchain CryptoCustomerConsumerWalletNetwork = "worldchain"
)

// Lists the Consumer Wallets for a Crypto Customer.
type CryptoCustomerConsumerWalletListParams struct {
	ListParams `form:"*"`
	ID         *string `form:"-"` // Included in URL
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoCustomerConsumerWalletListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// A consumer wallet represents a cryptocurrency wallet address associated with a Crypto Customer.
type CryptoCustomerConsumerWallet struct {
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// The blockchain network for this wallet
	Network CryptoCustomerConsumerWalletNetwork `json:"network"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Whether ownership of this wallet has been verified
	VerifiedOwnership bool `json:"verified_ownership"`
	// The wallet address
	WalletAddress string `json:"wallet_address"`
}

// CryptoCustomerConsumerWalletList is a list of CustomerConsumerWallets as retrieved from a list endpoint.
type CryptoCustomerConsumerWalletList struct {
	APIResource
	ListMeta
	Data []*CryptoCustomerConsumerWallet `json:"data"`
}
