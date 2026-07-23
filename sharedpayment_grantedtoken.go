//
//
// File generated from our OpenAPI spec
//
//

package stripe

// The reason why the SharedPaymentGrantedToken has been deactivated.
type SharedPaymentGrantedTokenDeactivatedReason string

// List of values that SharedPaymentGrantedTokenDeactivatedReason can take
const (
	SharedPaymentGrantedTokenDeactivatedReasonConsumed SharedPaymentGrantedTokenDeactivatedReason = "consumed"
	SharedPaymentGrantedTokenDeactivatedReasonExpired  SharedPaymentGrantedTokenDeactivatedReason = "expired"
	SharedPaymentGrantedTokenDeactivatedReasonResolved SharedPaymentGrantedTokenDeactivatedReason = "resolved"
	SharedPaymentGrantedTokenDeactivatedReasonRevoked  SharedPaymentGrantedTokenDeactivatedReason = "revoked"
)

// The type of the card wallet, one of `amex_express_checkout`, `apple_pay`, `google_pay`, `masterpass`, `samsung_pay`, `visa_checkout`, or `link`. An additional hash is included on the Wallet subhash with a name matching this value. It contains additional information specific to the card wallet type.
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletTypeAmexExpressCheckout SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType = "amex_express_checkout"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletTypeApplePay            SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType = "apple_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletTypeGooglePay           SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType = "google_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletTypeLink                SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType = "link"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletTypeMasterpass          SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType = "masterpass"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletTypeSamsungPay          SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType = "samsung_pay"
	SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletTypeVisaCheckout        SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType = "visa_checkout"
)

// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
type SharedPaymentGrantedTokenPaymentMethodDetailsType string

// List of values that SharedPaymentGrantedTokenPaymentMethodDetailsType can take
const (
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeAffirm  SharedPaymentGrantedTokenPaymentMethodDetailsType = "affirm"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeCard    SharedPaymentGrantedTokenPaymentMethodDetailsType = "card"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeKlarna  SharedPaymentGrantedTokenPaymentMethodDetailsType = "klarna"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeLink    SharedPaymentGrantedTokenPaymentMethodDetailsType = "link"
	SharedPaymentGrantedTokenPaymentMethodDetailsTypeShopPay SharedPaymentGrantedTokenPaymentMethodDetailsType = "shop_pay"
)

// The recurring interval at which the shared payment token's amount usage restrictions reset.
type SharedPaymentGrantedTokenUsageLimitsRecurringInterval string

// List of values that SharedPaymentGrantedTokenUsageLimitsRecurringInterval can take
const (
	SharedPaymentGrantedTokenUsageLimitsRecurringIntervalMonth SharedPaymentGrantedTokenUsageLimitsRecurringInterval = "month"
	SharedPaymentGrantedTokenUsageLimitsRecurringIntervalWeek  SharedPaymentGrantedTokenUsageLimitsRecurringInterval = "week"
	SharedPaymentGrantedTokenUsageLimitsRecurringIntervalYear  SharedPaymentGrantedTokenUsageLimitsRecurringInterval = "year"
)

// Retrieves an existing SharedPaymentGrantedToken object
type SharedPaymentGrantedTokenParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *SharedPaymentGrantedTokenParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Retrieves an existing SharedPaymentGrantedToken object
type SharedPaymentGrantedTokenRetrieveParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *SharedPaymentGrantedTokenRetrieveParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// Details about the agent that issued this SharedPaymentGrantedToken.
type SharedPaymentGrantedTokenAgentDetails struct {
	// The Stripe Profile ID of the agent that issued this SharedPaymentGrantedToken.
	NetworkBusinessProfile string `json:"network_business_profile"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsAffirm struct{}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type SharedPaymentGrantedTokenPaymentMethodDetailsBillingDetails struct {
	// Billing address.
	Address *Address `json:"address"`
	// Email address.
	Email string `json:"email"`
	// Full name.
	Name string `json:"name"`
	// Billing phone number (including extension).
	Phone string `json:"phone"`
	// Taxpayer identification number. Used only for transactions between LATAM buyers and non-LATAM sellers.
	TaxID string `json:"tax_id"`
}

// Checks on Card address and CVC if provided.
type SharedPaymentGrantedTokenPaymentMethodDetailsCardChecks struct {
	// If a address line1 was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressLine1Check string `json:"address_line1_check"`
	// If a address postal code was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressPostalCodeCheck string `json:"address_postal_code_check"`
	// If a CVC was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	CVCCheck string `json:"cvc_check"`
}

// Contains information about card networks that can be used to process the payment.
type SharedPaymentGrantedTokenPaymentMethodDetailsCardNetworks struct {
	// All networks available for selection via [payment_method_options.card.network](https://docs.stripe.com/api/payment_intents/confirm#confirm_payment_intent-payment_method_options-card-network).
	Available []string `json:"available"`
	// The preferred network for co-branded cards. Can be `cartes_bancaires`, `mastercard`, `visa` or `invalid_preference` if requested network is not valid for the card.
	Preferred string `json:"preferred"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletAmexExpressCheckout struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletApplePay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletGooglePay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletLink struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletMasterpass struct {
	// Owner's verified billing address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	BillingAddress *Address `json:"billing_address"`
	// Owner's verified email. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Email string `json:"email"`
	// Owner's verified full name. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Name string `json:"name"`
	// Owner's verified shipping address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	ShippingAddress *Address `json:"shipping_address"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletSamsungPay struct{}
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletVisaCheckout struct {
	// Owner's verified billing address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	BillingAddress *Address `json:"billing_address"`
	// Owner's verified email. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Email string `json:"email"`
	// Owner's verified full name. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Name string `json:"name"`
	// Owner's verified shipping address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	ShippingAddress *Address `json:"shipping_address"`
}

// If this Card is part of a card wallet, this contains the details of the card wallet.
type SharedPaymentGrantedTokenPaymentMethodDetailsCardWallet struct {
	AmexExpressCheckout *SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletAmexExpressCheckout `json:"amex_express_checkout,omitempty"`
	ApplePay            *SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletApplePay            `json:"apple_pay,omitempty"`
	// (For tokenized numbers only.) The last four digits of the device account number.
	DynamicLast4 string                                                             `json:"dynamic_last4"`
	GooglePay    *SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletGooglePay  `json:"google_pay,omitempty"`
	Link         *SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletLink       `json:"link,omitempty"`
	Masterpass   *SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletMasterpass `json:"masterpass,omitempty"`
	SamsungPay   *SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletSamsungPay `json:"samsung_pay,omitempty"`
	// The type of the card wallet, one of `amex_express_checkout`, `apple_pay`, `google_pay`, `masterpass`, `samsung_pay`, `visa_checkout`, or `link`. An additional hash is included on the Wallet subhash with a name matching this value. It contains additional information specific to the card wallet type.
	Type         SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletType          `json:"type"`
	VisaCheckout *SharedPaymentGrantedTokenPaymentMethodDetailsCardWalletVisaCheckout `json:"visa_checkout,omitempty"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsCard struct {
	// Card brand. Can be `amex`, `cartes_bancaires`, `diners`, `discover`, `eftpos_au`, `jcb`, `link`, `mastercard`, `unionpay`, `visa` or `unknown`.
	Brand string `json:"brand"`
	// Checks on Card address and CVC if provided.
	Checks *SharedPaymentGrantedTokenPaymentMethodDetailsCardChecks `json:"checks,omitempty"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// A high-level description of the type of cards issued in this range. (For internal use only and not typically available in standard API requests.)
	Description string `json:"description,omitempty"`
	// The brand to use when displaying the card, this accounts for customer's brand choice on dual-branded cards. Can be `american_express`, `cartes_bancaires`, `diners_club`, `discover`, `eftpos_australia`, `interac`, `jcb`, `mastercard`, `union_pay`, `visa`, or `other` and may contain more values in the future.
	DisplayBrand string `json:"display_brand"`
	// Two-digit number representing the card's expiration month.
	ExpMonth int64 `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear int64 `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who've signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.
	//
	// *As of May 1, 2021, card fingerprint in India for Connect changed to allow two fingerprints for the same card---one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint,omitempty"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding string `json:"funding"`
	// Issuer identification number of the card. (For internal use only and not typically available in standard API requests.)
	IIN string `json:"iin,omitempty"`
	// The name of the card's issuing bank. (For internal use only and not typically available in standard API requests.)
	Issuer string `json:"issuer,omitempty"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// Contains information about card networks that can be used to process the payment.
	Networks *SharedPaymentGrantedTokenPaymentMethodDetailsCardNetworks `json:"networks"`
	// If this Card is part of a card wallet, this contains the details of the card wallet.
	Wallet *SharedPaymentGrantedTokenPaymentMethodDetailsCardWallet `json:"wallet"`
}

// The customer's date of birth, if provided.
type SharedPaymentGrantedTokenPaymentMethodDetailsKlarnaDOB struct {
	// The day of birth, between 1 and 31.
	Day int64 `json:"day"`
	// The month of birth, between 1 and 12.
	Month int64 `json:"month"`
	// The four-digit year of birth.
	Year int64 `json:"year"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsKlarna struct {
	// The customer's date of birth, if provided.
	DOB *SharedPaymentGrantedTokenPaymentMethodDetailsKlarnaDOB `json:"dob,omitempty"`
}
type SharedPaymentGrantedTokenPaymentMethodDetailsLink struct {
	// Account owner's email address.
	Email string `json:"email"`
	// [Deprecated] This is a legacy parameter that no longer has any function.
	// Deprecated:
	PersistentToken string `json:"persistent_token,omitempty"`
}

// Details of the PaymentMethod that was shared via this token.
type SharedPaymentGrantedTokenPaymentMethodDetails struct {
	Affirm *SharedPaymentGrantedTokenPaymentMethodDetailsAffirm `json:"affirm,omitempty"`
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails *SharedPaymentGrantedTokenPaymentMethodDetailsBillingDetails `json:"billing_details"`
	Card           *SharedPaymentGrantedTokenPaymentMethodDetailsCard           `json:"card,omitempty"`
	Klarna         *SharedPaymentGrantedTokenPaymentMethodDetailsKlarna         `json:"klarna,omitempty"`
	Link           *SharedPaymentGrantedTokenPaymentMethodDetailsLink           `json:"link,omitempty"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type SharedPaymentGrantedTokenPaymentMethodDetailsType `json:"type"`
}

// Bot risk insight.
type SharedPaymentGrantedTokenRiskDetailsInsightsBot struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight.
	Score float64 `json:"score"`
}

// Card issuer decline risk insight.
type SharedPaymentGrantedTokenRiskDetailsInsightsCardIssuerDecline struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight.
	Score float64 `json:"score"`
}

// Card testing risk insight.
type SharedPaymentGrantedTokenRiskDetailsInsightsCardTesting struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight.
	Score float64 `json:"score"`
}

// Fraudulent dispute risk insight.
type SharedPaymentGrantedTokenRiskDetailsInsightsFraudulentDispute struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight.
	Score int64 `json:"score"`
}

// Stolen card risk insight.
type SharedPaymentGrantedTokenRiskDetailsInsightsStolenCard struct {
	// Recommended action for this insight.
	RecommendedAction string `json:"recommended_action"`
	// Risk score for this insight.
	Score int64 `json:"score"`
}

// Risk insights for this token, including scores and recommended actions for each risk type.
type SharedPaymentGrantedTokenRiskDetailsInsights struct {
	// Bot risk insight.
	Bot *SharedPaymentGrantedTokenRiskDetailsInsightsBot `json:"bot,omitempty"`
	// Card issuer decline risk insight.
	CardIssuerDecline *SharedPaymentGrantedTokenRiskDetailsInsightsCardIssuerDecline `json:"card_issuer_decline,omitempty"`
	// Card testing risk insight.
	CardTesting *SharedPaymentGrantedTokenRiskDetailsInsightsCardTesting `json:"card_testing,omitempty"`
	// Fraudulent dispute risk insight.
	FraudulentDispute *SharedPaymentGrantedTokenRiskDetailsInsightsFraudulentDispute `json:"fraudulent_dispute"`
	// Stolen card risk insight.
	StolenCard *SharedPaymentGrantedTokenRiskDetailsInsightsStolenCard `json:"stolen_card,omitempty"`
}

// Risk details of the SharedPaymentGrantedToken.
type SharedPaymentGrantedTokenRiskDetails struct {
	// Risk insights for this token, including scores and recommended actions for each risk type.
	Insights *SharedPaymentGrantedTokenRiskDetailsInsights `json:"insights"`
}

// The total amount captured using this SharedPaymentToken.
type SharedPaymentGrantedTokenUsageDetailsAmountCaptured struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Integer value of the amount in the smallest currency unit.
	Value int64 `json:"value"`
}

// Some details about how the SharedPaymentGrantedToken has been used already.
type SharedPaymentGrantedTokenUsageDetails struct {
	// The total amount captured using this SharedPaymentToken.
	AmountCaptured *SharedPaymentGrantedTokenUsageDetailsAmountCaptured `json:"amount_captured"`
}

// Limits on how this SharedPaymentGrantedToken can be used.
type SharedPaymentGrantedTokenUsageLimits struct {
	// Three-letter [ISO currency code](https://www.iso.org/iso-4217-currency-codes.html), in lowercase. Must be a [supported currency](https://stripe.com/docs/currencies).
	Currency Currency `json:"currency"`
	// Time at which this SharedPaymentToken expires and can no longer be used to confirm a PaymentIntent.
	ExpiresAt int64 `json:"expires_at"`
	// Max amount that can be captured using this SharedPaymentToken.
	MaxAmount int64 `json:"max_amount"`
	// The recurring interval at which the shared payment token's amount usage restrictions reset.
	RecurringInterval SharedPaymentGrantedTokenUsageLimitsRecurringInterval `json:"recurring_interval,omitempty"`
}

// SharedPaymentGrantedToken is the view-only resource of a SharedPaymentIssuedToken, which is a limited-use reference to a PaymentMethod.
// When another Stripe merchant shares a SharedPaymentIssuedToken with you, you can view attributes of the shared token using the SharedPaymentGrantedToken API, and use it with a PaymentIntent.
type SharedPaymentGrantedToken struct {
	APIResource
	// Details about the agent that issued this SharedPaymentGrantedToken.
	AgentDetails *SharedPaymentGrantedTokenAgentDetails `json:"agent_details"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// Time at which this SharedPaymentGrantedToken expires and can no longer be used to confirm a PaymentIntent.
	DeactivatedAt int64 `json:"deactivated_at"`
	// The reason why the SharedPaymentGrantedToken has been deactivated.
	DeactivatedReason SharedPaymentGrantedTokenDeactivatedReason `json:"deactivated_reason"`
	// Unique identifier for the object.
	ID string `json:"id"`
	// If the object exists in live mode, the value is `true`. If the object exists in test mode, the value is `false`.
	Livemode bool `json:"livemode"`
	// String representing the object's type. Objects of the same type share the same value.
	Object string `json:"object"`
	// Details of the PaymentMethod that was shared via this token.
	PaymentMethodDetails *SharedPaymentGrantedTokenPaymentMethodDetails `json:"payment_method_details"`
	// Risk details of the SharedPaymentGrantedToken.
	RiskDetails *SharedPaymentGrantedTokenRiskDetails `json:"risk_details,omitempty"`
	// Metadata about the SharedPaymentGrantedToken.
	SharedMetadata map[string]string `json:"shared_metadata"`
	// Some details about how the SharedPaymentGrantedToken has been used already.
	UsageDetails *SharedPaymentGrantedTokenUsageDetails `json:"usage_details"`
	// Limits on how this SharedPaymentGrantedToken can be used.
	UsageLimits *SharedPaymentGrantedTokenUsageLimits `json:"usage_limits"`
}
