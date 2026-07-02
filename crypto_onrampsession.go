//
//
// File generated from our OpenAPI spec
//
//

package stripe

// If a platform wants to lock the currencies an session will support, they can add supported currencies to this array. If left null, the experience will allow selection of all supported destination currencies.
type CryptoOnrampSessionTransactionDetailsDestinationCurrency string

// List of values that CryptoOnrampSessionTransactionDetailsDestinationCurrency can take
const (
	CryptoOnrampSessionTransactionDetailsDestinationCurrencyAvax  CryptoOnrampSessionTransactionDetailsDestinationCurrency = "avax"
	CryptoOnrampSessionTransactionDetailsDestinationCurrencyBtc   CryptoOnrampSessionTransactionDetailsDestinationCurrency = "btc"
	CryptoOnrampSessionTransactionDetailsDestinationCurrencyEth   CryptoOnrampSessionTransactionDetailsDestinationCurrency = "eth"
	CryptoOnrampSessionTransactionDetailsDestinationCurrencyMatic CryptoOnrampSessionTransactionDetailsDestinationCurrency = "matic"
	CryptoOnrampSessionTransactionDetailsDestinationCurrencySol   CryptoOnrampSessionTransactionDetailsDestinationCurrency = "sol"
	CryptoOnrampSessionTransactionDetailsDestinationCurrencyUsdc  CryptoOnrampSessionTransactionDetailsDestinationCurrency = "usdc"
	CryptoOnrampSessionTransactionDetailsDestinationCurrencyWld   CryptoOnrampSessionTransactionDetailsDestinationCurrency = "wld"
	CryptoOnrampSessionTransactionDetailsDestinationCurrencyXlm   CryptoOnrampSessionTransactionDetailsDestinationCurrency = "xlm"
)

// The specific crypto network the `destination_currency` is settled on. If `destination_networks` is set, it must be a value in that array.
type CryptoOnrampSessionTransactionDetailsDestinationNetwork string

// List of values that CryptoOnrampSessionTransactionDetailsDestinationNetwork can take
const (
	CryptoOnrampSessionTransactionDetailsDestinationNetworkAvalanche  CryptoOnrampSessionTransactionDetailsDestinationNetwork = "avalanche"
	CryptoOnrampSessionTransactionDetailsDestinationNetworkBase       CryptoOnrampSessionTransactionDetailsDestinationNetwork = "base"
	CryptoOnrampSessionTransactionDetailsDestinationNetworkBitcoin    CryptoOnrampSessionTransactionDetailsDestinationNetwork = "bitcoin"
	CryptoOnrampSessionTransactionDetailsDestinationNetworkEthereum   CryptoOnrampSessionTransactionDetailsDestinationNetwork = "ethereum"
	CryptoOnrampSessionTransactionDetailsDestinationNetworkOptimism   CryptoOnrampSessionTransactionDetailsDestinationNetwork = "optimism"
	CryptoOnrampSessionTransactionDetailsDestinationNetworkPolygon    CryptoOnrampSessionTransactionDetailsDestinationNetwork = "polygon"
	CryptoOnrampSessionTransactionDetailsDestinationNetworkSolana     CryptoOnrampSessionTransactionDetailsDestinationNetwork = "solana"
	CryptoOnrampSessionTransactionDetailsDestinationNetworkStellar    CryptoOnrampSessionTransactionDetailsDestinationNetwork = "stellar"
	CryptoOnrampSessionTransactionDetailsDestinationNetworkWorldchain CryptoOnrampSessionTransactionDetailsDestinationNetwork = "worldchain"
)

// Speed at which the cryptocurrency is delivered to the wallet
// One of:
//
//	`instant` (default): crypto is delivered when payment is confirmed
//	`standard`: crypto is delivered when payment settles
type CryptoOnrampSessionTransactionDetailsSettlementSpeed string

// List of values that CryptoOnrampSessionTransactionDetailsSettlementSpeed can take
const (
	CryptoOnrampSessionTransactionDetailsSettlementSpeedInstant  CryptoOnrampSessionTransactionDetailsSettlementSpeed = "instant"
	CryptoOnrampSessionTransactionDetailsSettlementSpeedStandard CryptoOnrampSessionTransactionDetailsSettlementSpeed = "standard"
)

// A fiat currency code
type CryptoOnrampSessionTransactionDetailsSourceCurrency string

// List of values that CryptoOnrampSessionTransactionDetailsSourceCurrency can take
const (
	CryptoOnrampSessionTransactionDetailsSourceCurrencyEUR CryptoOnrampSessionTransactionDetailsSourceCurrency = "eur"
	CryptoOnrampSessionTransactionDetailsSourceCurrencyGBP CryptoOnrampSessionTransactionDetailsSourceCurrency = "gbp"
	CryptoOnrampSessionTransactionDetailsSourceCurrencyUSD CryptoOnrampSessionTransactionDetailsSourceCurrency = "usd"
)

// Returns a list of onramp sessions that match the filter criteria. The onramp sessions are returned in sorted order, with the most recent onramp sessions appearing first.
type CryptoOnrampSessionListParams struct {
	ListParams `form:"*"`
	// Only return onramp sessions that were created during the given date interval.
	Created *int64 `form:"created" json:"created,omitempty"`
	// Only return onramp sessions that were created during the given date interval.
	CreatedRange *RangeQueryParams `form:"created" json:"-"`
	// The destination cryptocurrency to filter by.
	DestinationCurrency *string `form:"destination_currency" json:"destination_currency,omitempty"`
	// The destination blockchain network to filter by.
	DestinationNetwork *string `form:"destination_network" json:"destination_network,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// The status of the Onramp Session. One of = `{initialized, rejected, requires_payment, fulfillment_processing, fulfillment_complete}`
	Status *string `form:"status" json:"status,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoOnrampSessionListParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Pre-populate some of the required KYC information for the user if you've already collected it within your application.
//
// Related guide: [Using the API](https://docs.stripe.com/crypto/using-the-api#how-to-pre-populate-customer-information)
type CryptoOnrampSessionKycDetailsParams struct{}

// The end customer's crypto wallet destination tag (for each network) to use for this transaction. This only applies for tag-based assets such as XLM.
//
// * When left null, the user enters their wallet in the onramp UI.
type CryptoOnrampSessionWalletAddressesDestinationTagsParams struct {
	Stellar *string `form:"stellar" json:"stellar,omitempty"`
}

// The end customer's crypto wallet address (for each network) to use for this transaction.
//
// * When left null, the user enters their wallet in the onramp UI.
// * When set, the platform must set either `destination_networks` or `destination_network` and we perform address validation. Users can still select a different wallet in the onramp UI.
type CryptoOnrampSessionWalletAddressesParams struct {
	// The end customer's crypto wallet destination tag (for each network) to use for this transaction. This only applies for tag-based assets such as XLM.
	//
	// * When left null, the user enters their wallet in the onramp UI.
	DestinationTags *CryptoOnrampSessionWalletAddressesDestinationTagsParams `form:"destination_tags" json:"destination_tags,omitempty"`
}

// Creates a CryptoOnrampSession object.
//
// After the CryptoOnrampSession is created, display the onramp session modal using the client_secret.
//
// Related guide: [Set up an onramp integration](https://docs.stripe.com/docs/crypto/integrate-the-onramp)
type CryptoOnrampSessionParams struct {
	Params `form:"*"`
	// The IP address of the customer the platform intends to onramp.
	//
	// If the user's IP is in a region we can't support, we return an `HTTP 400` with an appropriate error code.
	//
	// We support IPv4 and IPv6 addresses. Geographic supportability is checked again later in the onramp flow, which provides a way to hide the onramp option from ineligible users for a better user experience.
	CustomerIPAddress *string `form:"customer_ip_address" json:"customer_ip_address,omitempty"`
	// The default amount of crypto to exchange into.
	//
	// * When left null, a default value is computed if `source_amount`, `destination_currency`, and `destination_network` are set.
	// * When set, both `destination_currency` and `destination_network` must also be set. All cryptocurrencies are supported to their full precisions (for example, 18 decimal places for `eth`). We validate and generate an error if the amount exceeds the supported precision based on the exchange currency. Setting `source_amount` is mutually exclusive with setting `destination_amount` (only one or the other is supported). Users can update the amount in the onramp UI.
	DestinationAmount *string `form:"destination_amount" json:"destination_amount,omitempty"`
	// The list of destination cryptocurrencies a user can choose from.
	//
	// * When left null, all supported cryptocurrencies are shown in the onramp UI subject to `destination_networks` if set.
	// * When set, it must be a non-empty array where all values in the array are valid cryptocurrencies. You can use it to lock users to a specific cryptocurrency by passing a single value array. Users **cannot** override this parameter.
	DestinationCurrencies []*string `form:"destination_currencies" json:"destination_currencies,omitempty"`
	// The default destination cryptocurrency.
	//
	// * When left null, the first value of `destination_currencies` is selected.
	// * When set, if `destination_currencies` is also set, the value of `destination_currency` must be present in that array. To lock a `destination_currency`, specify that value as the single value for `destination_currencies`. Users can select a different cryptocurrency in the onramp UI subject to `destination_currencies` if set.
	DestinationCurrency *string `form:"destination_currency" json:"destination_currency,omitempty"`
	// The default destination crypto network.
	//
	// * When left null, the first value of `destination_networks` is selected.
	// * When set, if `destination_networks` is also set, the value of `destination_network` must be present in that array. To lock a `destination_network`, specify that value as the single value for `destination_networks`. Users can select a different network in the onramp UI subject to `destination_networks` if set.
	DestinationNetwork *string `form:"destination_network" json:"destination_network,omitempty"`
	// The list of destination crypto networks user can choose from.
	//
	// * When left null, all supported crypto networks are shown in the onramp UI.
	// * When set, it must be a non-empty array where values in the array are each a valid crypto network. It can be used to lock users to a specific network by passing a single value array. Users **cannot** override this parameter.
	DestinationNetworks []*string `form:"destination_networks" json:"destination_networks,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Pre-populate some of the required KYC information for the user if you've already collected it within your application.
	//
	// Related guide: [Using the API](https://docs.stripe.com/crypto/using-the-api#how-to-pre-populate-customer-information)
	KycDetails *CryptoOnrampSessionKycDetailsParams `form:"kyc_details" json:"kyc_details,omitempty"`
	// Whether or not to lock the suggested wallet address. If destination tags are provided, this will also lock the destination tags.
	LockWalletAddress *bool `form:"lock_wallet_address" json:"lock_wallet_address,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Speed at which the cryptocurrency is delivered to the wallet
	// One of:
	//  `instant` (default): crypto is delivered when payment is confirmed
	//  `standard`: crypto is delivered when payment settles
	SettlementSpeed *string `form:"settlement_speed" json:"settlement_speed,omitempty"`
	// The default amount of fiat (in decimal) to exchange into crypto.
	//
	// * When left null, a default value is computed if `destination_amount` is set.
	// * When set, setting `source_amount` is mutually exclusive with setting `destination_amount` (only one or the other is supported). We don't support fractional pennies. If fractional minor units of a currency are passed in, it generates an error. Users can update the value in the onramp UI.
	SourceAmount *string `form:"source_amount" json:"source_amount,omitempty"`
	// The default source fiat currency for the onramp session.
	//
	// * When left null, a default currency is selected based on user locale.
	// * When set, it must be one of the fiat currencies supported by onramp. Users can still select a different currency in the onramp UI.
	SourceCurrency *string `form:"source_currency" json:"source_currency,omitempty"`
	// The end customer's crypto wallet address (for each network) to use for this transaction.
	//
	// * When left null, the user enters their wallet in the onramp UI.
	// * When set, the platform must set either `destination_networks` or `destination_network` and we perform address validation. Users can still select a different wallet in the onramp UI.
	WalletAddresses *CryptoOnrampSessionWalletAddressesParams `form:"wallet_addresses" json:"wallet_addresses,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoOnrampSessionParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CryptoOnrampSessionParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// If this is a Mandate accepted offline, this hash contains details about the offline acceptance.
type CryptoOnrampSessionCheckoutMandateDataCustomerAcceptanceOfflineParams struct{}

// If this is a Mandate accepted online, this hash contains details about the online acceptance.
type CryptoOnrampSessionCheckoutMandateDataCustomerAcceptanceOnlineParams struct {
	// The IP address from which the Mandate was accepted by the customer.
	IPAddress *string `form:"ip_address" json:"ip_address"`
	// The user agent of the browser from which the Mandate was accepted by the customer.
	UserAgent *string `form:"user_agent" json:"user_agent"`
}

// This hash contains details about the customer acceptance of the Mandate.
type CryptoOnrampSessionCheckoutMandateDataCustomerAcceptanceParams struct {
	// The time at which the customer accepted the Mandate.
	AcceptedAt *int64 `form:"accepted_at" json:"accepted_at,omitempty"`
	// If this is a Mandate accepted offline, this hash contains details about the offline acceptance.
	Offline *CryptoOnrampSessionCheckoutMandateDataCustomerAcceptanceOfflineParams `form:"offline" json:"offline,omitempty"`
	// If this is a Mandate accepted online, this hash contains details about the online acceptance.
	Online *CryptoOnrampSessionCheckoutMandateDataCustomerAcceptanceOnlineParams `form:"online" json:"online,omitempty"`
	// The type of customer acceptance information included with the Mandate. One of `online` or `offline`.
	Type *string `form:"type" json:"type"`
}

// This hash contains details about the mandate to create
type CryptoOnrampSessionCheckoutMandateDataParams struct {
	// This hash contains details about the customer acceptance of the Mandate.
	CustomerAcceptance *CryptoOnrampSessionCheckoutMandateDataCustomerAcceptanceParams `form:"customer_acceptance" json:"customer_acceptance"`
}

// Completes a headless CryptoOnrampSession.
//
// This method will attempt to confirm the payment and execute the quote to deliver the crypto to the customer.
type CryptoOnrampSessionCheckoutParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// This hash contains details about the mandate to create
	MandateData *CryptoOnrampSessionCheckoutMandateDataParams `form:"mandate_data" json:"mandate_data,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoOnrampSessionCheckoutParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Refreshes an executable quote for a CryptoOnrampSession.
type CryptoOnrampSessionQuoteParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoOnrampSessionQuoteParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Pre-populate some of the required KYC information for the user if you've already collected it within your application.
//
// Related guide: [Using the API](https://docs.stripe.com/crypto/using-the-api#how-to-pre-populate-customer-information)
type CryptoOnrampSessionCreateKycDetailsParams struct{}

// The end customer's crypto wallet destination tag (for each network) to use for this transaction. This only applies for tag-based assets such as XLM.
//
// * When left null, the user enters their wallet in the onramp UI.
type CryptoOnrampSessionCreateWalletAddressesDestinationTagsParams struct {
	Stellar *string `form:"stellar" json:"stellar,omitempty"`
}

// The end customer's crypto wallet address (for each network) to use for this transaction.
//
// * When left null, the user enters their wallet in the onramp UI.
// * When set, the platform must set either `destination_networks` or `destination_network` and we perform address validation. Users can still select a different wallet in the onramp UI.
type CryptoOnrampSessionCreateWalletAddressesParams struct {
	// The end customer's crypto wallet destination tag (for each network) to use for this transaction. This only applies for tag-based assets such as XLM.
	//
	// * When left null, the user enters their wallet in the onramp UI.
	DestinationTags *CryptoOnrampSessionCreateWalletAddressesDestinationTagsParams `form:"destination_tags" json:"destination_tags,omitempty"`
}

// Creates a CryptoOnrampSession object.
//
// After the CryptoOnrampSession is created, display the onramp session modal using the client_secret.
//
// Related guide: [Set up an onramp integration](https://docs.stripe.com/docs/crypto/integrate-the-onramp)
type CryptoOnrampSessionCreateParams struct {
	Params `form:"*"`
	// The IP address of the customer the platform intends to onramp.
	//
	// If the user's IP is in a region we can't support, we return an `HTTP 400` with an appropriate error code.
	//
	// We support IPv4 and IPv6 addresses. Geographic supportability is checked again later in the onramp flow, which provides a way to hide the onramp option from ineligible users for a better user experience.
	CustomerIPAddress *string `form:"customer_ip_address" json:"customer_ip_address,omitempty"`
	// The default amount of crypto to exchange into.
	//
	// * When left null, a default value is computed if `source_amount`, `destination_currency`, and `destination_network` are set.
	// * When set, both `destination_currency` and `destination_network` must also be set. All cryptocurrencies are supported to their full precisions (for example, 18 decimal places for `eth`). We validate and generate an error if the amount exceeds the supported precision based on the exchange currency. Setting `source_amount` is mutually exclusive with setting `destination_amount` (only one or the other is supported). Users can update the amount in the onramp UI.
	DestinationAmount *string `form:"destination_amount" json:"destination_amount,omitempty"`
	// The list of destination cryptocurrencies a user can choose from.
	//
	// * When left null, all supported cryptocurrencies are shown in the onramp UI subject to `destination_networks` if set.
	// * When set, it must be a non-empty array where all values in the array are valid cryptocurrencies. You can use it to lock users to a specific cryptocurrency by passing a single value array. Users **cannot** override this parameter.
	DestinationCurrencies []*string `form:"destination_currencies" json:"destination_currencies,omitempty"`
	// The default destination cryptocurrency.
	//
	// * When left null, the first value of `destination_currencies` is selected.
	// * When set, if `destination_currencies` is also set, the value of `destination_currency` must be present in that array. To lock a `destination_currency`, specify that value as the single value for `destination_currencies`. Users can select a different cryptocurrency in the onramp UI subject to `destination_currencies` if set.
	DestinationCurrency *string `form:"destination_currency" json:"destination_currency,omitempty"`
	// The default destination crypto network.
	//
	// * When left null, the first value of `destination_networks` is selected.
	// * When set, if `destination_networks` is also set, the value of `destination_network` must be present in that array. To lock a `destination_network`, specify that value as the single value for `destination_networks`. Users can select a different network in the onramp UI subject to `destination_networks` if set.
	DestinationNetwork *string `form:"destination_network" json:"destination_network,omitempty"`
	// The list of destination crypto networks user can choose from.
	//
	// * When left null, all supported crypto networks are shown in the onramp UI.
	// * When set, it must be a non-empty array where values in the array are each a valid crypto network. It can be used to lock users to a specific network by passing a single value array. Users **cannot** override this parameter.
	DestinationNetworks []*string `form:"destination_networks" json:"destination_networks,omitempty"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// Pre-populate some of the required KYC information for the user if you've already collected it within your application.
	//
	// Related guide: [Using the API](https://docs.stripe.com/crypto/using-the-api#how-to-pre-populate-customer-information)
	KycDetails *CryptoOnrampSessionCreateKycDetailsParams `form:"kyc_details" json:"kyc_details,omitempty"`
	// Whether or not to lock the suggested wallet address. If destination tags are provided, this will also lock the destination tags.
	LockWalletAddress *bool `form:"lock_wallet_address" json:"lock_wallet_address,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// Speed at which the cryptocurrency is delivered to the wallet
	// One of:
	//  `instant` (default): crypto is delivered when payment is confirmed
	//  `standard`: crypto is delivered when payment settles
	SettlementSpeed *string `form:"settlement_speed" json:"settlement_speed,omitempty"`
	// The default amount of fiat (in decimal) to exchange into crypto.
	//
	// * When left null, a default value is computed if `destination_amount` is set.
	// * When set, setting `source_amount` is mutually exclusive with setting `destination_amount` (only one or the other is supported). We don't support fractional pennies. If fractional minor units of a currency are passed in, it generates an error. Users can update the value in the onramp UI.
	SourceAmount *string `form:"source_amount" json:"source_amount,omitempty"`
	// The default source fiat currency for the onramp session.
	//
	// * When left null, a default currency is selected based on user locale.
	// * When set, it must be one of the fiat currencies supported by onramp. Users can still select a different currency in the onramp UI.
	SourceCurrency *string `form:"source_currency" json:"source_currency,omitempty"`
	// The end customer's crypto wallet address (for each network) to use for this transaction.
	//
	// * When left null, the user enters their wallet in the onramp UI.
	// * When set, the platform must set either `destination_networks` or `destination_network` and we perform address validation. Users can still select a different wallet in the onramp UI.
	WalletAddresses *CryptoOnrampSessionCreateWalletAddressesParams `form:"wallet_addresses" json:"wallet_addresses,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoOnrampSessionCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *CryptoOnrampSessionCreateParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Retrieves the details of a CryptoOnrampSession that was previously created.
type CryptoOnrampSessionRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *CryptoOnrampSessionRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Details about the fees associated with this transaction
type CryptoOnrampSessionTransactionDetailsFees struct {
	// The cost associated with moving crypto from Stripe to the end consumer's wallet. e.g: for ETH, this is called 'gas fee', for BTC this is a 'miner's fee'.
	NetworkFeeAmount string `json:"network_fee_amount"`
	// Fee for processing the transaction.
	TransactionFeeAmount string `json:"transaction_fee_amount"`
}

// The end customer's crypto wallet destination tag (for each network) to use for this transaction.
type CryptoOnrampSessionTransactionDetailsWalletAddressesDestinationTags struct {
	// A stellar destination tag
	Stellar string `json:"stellar"`
}

// The end customer's crypto wallet address (for each network) to use for this transaction.
type CryptoOnrampSessionTransactionDetailsWalletAddresses struct {
	// An avalanche address
	Avalanche string `json:"avalanche"`
	// A base address
	BaseNetwork string `json:"base_network"`
	// A bitcoin address
	Bitcoin string `json:"bitcoin"`
	// The end customer's crypto wallet destination tag (for each network) to use for this transaction.
	DestinationTags *CryptoOnrampSessionTransactionDetailsWalletAddressesDestinationTags `json:"destination_tags"`
	// An ethereum address
	Ethereum string `json:"ethereum"`
	// An optimism address
	Optimism string `json:"optimism"`
	// A polygon address
	Polygon string `json:"polygon"`
	// A solana address
	Solana string `json:"solana"`
	// A stellar address
	Stellar string `json:"stellar"`
	// A worldchain address
	Worldchain string `json:"worldchain"`
}
type CryptoOnrampSessionTransactionDetails struct {
	// The amount of crypto the customer will get deposited into their wallet
	DestinationAmount string `json:"destination_amount"`
	// If a platform wants to lock the currencies an session will support, they can add supported currencies to this array. If left null, the experience will allow selection of all supported destination currencies.
	DestinationCurrencies []CryptoOnrampSessionTransactionDetailsDestinationCurrency `json:"destination_currencies"`
	// The selected `destination_currency` to convert the `source` to. This should be a crypto currency code. If `destination_currencies` is set, it must be a value in that array.
	DestinationCurrency CryptoOnrampSessionTransactionDetailsDestinationCurrency `json:"destination_currency"`
	// The specific crypto network the `destination_currency` is settled on. If `destination_networks` is set, it must be a value in that array.
	DestinationNetwork CryptoOnrampSessionTransactionDetailsDestinationNetwork `json:"destination_network"`
	// If a platform wants to lock the supported networks, they can do so through this array. If left null, the experience will allow selection of all supported networks.
	DestinationNetworks []CryptoOnrampSessionTransactionDetailsDestinationNetwork `json:"destination_networks"`
	// Details about the fees associated with this transaction
	Fees *CryptoOnrampSessionTransactionDetailsFees `json:"fees"`
	// Whether or not to lock the suggested wallet address.
	LockWalletAddress bool `json:"lock_wallet_address"`
	// Speed at which the cryptocurrency is delivered to the wallet
	// One of:
	//  `instant` (default): crypto is delivered when payment is confirmed
	//  `standard`: crypto is delivered when payment settles
	SettlementSpeed CryptoOnrampSessionTransactionDetailsSettlementSpeed `json:"settlement_speed"`
	// The amount of fiat we intend to onramp - excluding fees
	SourceAmount string `json:"source_amount"`
	// A fiat currency code
	SourceCurrency CryptoOnrampSessionTransactionDetailsSourceCurrency `json:"source_currency"`
	// The on-chain transaction hash (also referred to as transaction ID or tx_hash) of the transaction that was sent to the customer's wallet. The format varies by chain (e.g. `0xc257...1a95` on Ethereum, `5UB1...v3xZ` on Solana, or `a1b2...bf00` on Bitcoin). This will only be set if the session reaches `status=fulfillment_complete` and we've transferred the crypto successfully to the external wallet.
	TransactionID string `json:"transaction_id"`
	// The consumer's wallet address (where crypto will be sent to)
	WalletAddress string `json:"wallet_address"`
	// The end customer's crypto wallet address (for each network) to use for this transaction.
	WalletAddresses *CryptoOnrampSessionTransactionDetailsWalletAddresses `json:"wallet_addresses"`
}

// A Crypto Onramp Session represents your customer's session as they purchase cryptocurrency through Stripe. Once payment is successful, Stripe will fulfill the delivery of cryptocurrency to your user's wallet and contain a reference to the crypto transaction ID.
//
// You can create an onramp session on your server and embed the widget on your frontend. Alternatively, you can redirect your users to the standalone hosted onramp.
//
// Related guide: [Integrate the onramp](https://docs.stripe.com/crypto/integrate-the-onramp)
type CryptoOnrampSession struct {
	APIResource
	// A client secret that can be used to drive a single session using our embedded widget.
	//
	// Related guide: [Set up an onramp integration](https://docs.stripe.com/crypto/integrate-the-onramp)
	ClientSecret string `json:"client_secret"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// Has the value `true` if any user kyc details were provided during the creation of the onramp session. Otherwise, has the value `false`.
	KycDetailsProvided bool `json:"kyc_details_provided"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Redirect your users to the URL for a prebuilt frontend integration of the crypto onramp on the standalone hosted onramp.
	//
	// Related guide: [Mint a session with a redirect url](https://docs.stripe.com/crypto/standalone-hosted-onramp#mint-a-session-with-a-redirect-url)
	RedirectURL string `json:"redirect_url"`
	// The status of the Onramp Session. One of = `{initialized, rejected, requires_payment, fulfillment_processing, fulfillment_complete}`
	Status             string                                 `json:"status"`
	TransactionDetails *CryptoOnrampSessionTransactionDetails `json:"transaction_details"`
}

// CryptoOnrampSessionList is a list of OnrampSessions as retrieved from a list endpoint.
type CryptoOnrampSessionList struct {
	APIResource
	ListMeta
	Data []*CryptoOnrampSession `json:"data"`
}
