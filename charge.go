//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Assessments from Stripe. If set, the value is `fraudulent`.
type ChargeFraudStripeReport string

// List of values that ChargeFraudStripeReport can take
const (
	ChargeFraudStripeReportFraudulent ChargeFraudStripeReport = "fraudulent"
)

// Assessments reported by you. If set, possible values of are `safe` and `fraudulent`.
type ChargeFraudUserReport string

// List of values that ChargeFraudUserReport can take
const (
	ChargeFraudUserReportFraudulent ChargeFraudUserReport = "fraudulent"
	ChargeFraudUserReportSafe       ChargeFraudUserReport = "safe"
)

// For authenticated transactions: how the customer was authenticated by
// the issuing bank.
type ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow can take
const (
	ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlowChallenge    ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "challenge"
	ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlowFrictionless ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow = "frictionless"
)

// Indicates the outcome of 3D Secure authentication.
type ChargePaymentMethodDetailsCardThreeDSecureResult string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureResult can take
const (
	ChargePaymentMethodDetailsCardThreeDSecureResultAttemptAcknowledged ChargePaymentMethodDetailsCardThreeDSecureResult = "attempt_acknowledged"
	ChargePaymentMethodDetailsCardThreeDSecureResultAuthenticated       ChargePaymentMethodDetailsCardThreeDSecureResult = "authenticated"
	ChargePaymentMethodDetailsCardThreeDSecureResultFailed              ChargePaymentMethodDetailsCardThreeDSecureResult = "failed"
	ChargePaymentMethodDetailsCardThreeDSecureResultNotSupported        ChargePaymentMethodDetailsCardThreeDSecureResult = "not_supported"
	ChargePaymentMethodDetailsCardThreeDSecureResultProcessingError     ChargePaymentMethodDetailsCardThreeDSecureResult = "processing_error"
)

// Additional information about why 3D Secure succeeded or failed based
// on the `result`.
type ChargePaymentMethodDetailsCardThreeDSecureResultReason string

// List of values that ChargePaymentMethodDetailsCardThreeDSecureResultReason can take
const (
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonAbandoned           ChargePaymentMethodDetailsCardThreeDSecureResultReason = "abandoned"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonBypassed            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "bypassed"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonCanceled            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "canceled"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonCardNotEnrolled     ChargePaymentMethodDetailsCardThreeDSecureResultReason = "card_not_enrolled"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonNetworkNotSupported ChargePaymentMethodDetailsCardThreeDSecureResultReason = "network_not_supported"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonProtocolError       ChargePaymentMethodDetailsCardThreeDSecureResultReason = "protocol_error"
	ChargePaymentMethodDetailsCardThreeDSecureResultReasonRejected            ChargePaymentMethodDetailsCardThreeDSecureResultReason = "rejected"
)

// The type of account being debited or credited
type ChargePaymentMethodDetailsCardPresentReceiptAccountType string

// List of values that ChargePaymentMethodDetailsCardPresentReceiptAccountType can take
const (
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeChecking ChargePaymentMethodDetailsCardPresentReceiptAccountType = "checking"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeCredit   ChargePaymentMethodDetailsCardPresentReceiptAccountType = "credit"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypePrepaid  ChargePaymentMethodDetailsCardPresentReceiptAccountType = "prepaid"
	ChargePaymentMethodDetailsCardPresentReceiptAccountTypeUnknown  ChargePaymentMethodDetailsCardPresentReceiptAccountType = "unknown"
)

// The Klarna payment method used for this transaction.
// Can be one of `pay_later`, `pay_now`, `pay_with_financing`, or `pay_in_installments`
type ChargePaymentMethodDetailsKlarnaPaymentMethodCategory string

// List of values that ChargePaymentMethodDetailsKlarnaPaymentMethodCategory can take
const (
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayLater          ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_later"
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayNow            ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_now"
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayWithFinancing  ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_with_financing"
	ChargePaymentMethodDetailsKlarnaPaymentMethodCategoryPayInInstallments ChargePaymentMethodDetailsKlarnaPaymentMethodCategory = "pay_in_installments"
)

// The type of transaction-specific details of the payment method used in the payment, one of `ach_credit_transfer`, `ach_debit`, `acss_debit`, `alipay`, `au_becs_debit`, `bancontact`, `card`, `card_present`, `eps`, `giropay`, `ideal`, `klarna`, `multibanco`, `p24`, `sepa_debit`, `sofort`, `stripe_account`, or `wechat`.
// An additional hash is included on `payment_method_details` with a name matching this value.
// It contains information specific to the payment method.
type ChargePaymentMethodDetailsType string

// List of values that ChargePaymentMethodDetailsType can take
const (
	ChargePaymentMethodDetailsTypeAchCreditTransfer ChargePaymentMethodDetailsType = "ach_credit_transfer"
	ChargePaymentMethodDetailsTypeAchDebit          ChargePaymentMethodDetailsType = "ach_debit"
	ChargePaymentMethodDetailsTypeAcssDebit         ChargePaymentMethodDetailsType = "acss_debit"
	ChargePaymentMethodDetailsTypeAlipay            ChargePaymentMethodDetailsType = "alipay"
	ChargePaymentMethodDetailsTypeAUBECSDebit       ChargePaymentMethodDetailsType = "au_becs_debit"
	ChargePaymentMethodDetailsTypeBACSDebit         ChargePaymentMethodDetailsType = "bacs_debit"
	ChargePaymentMethodDetailsTypeBancontact        ChargePaymentMethodDetailsType = "bancontact"
	ChargePaymentMethodDetailsTypeCard              ChargePaymentMethodDetailsType = "card"
	ChargePaymentMethodDetailsTypeCardPresent       ChargePaymentMethodDetailsType = "card_present"
	ChargePaymentMethodDetailsTypeEps               ChargePaymentMethodDetailsType = "eps"
	ChargePaymentMethodDetailsTypeFPX               ChargePaymentMethodDetailsType = "fpx"
	ChargePaymentMethodDetailsTypeGiropay           ChargePaymentMethodDetailsType = "giropay"
	ChargePaymentMethodDetailsTypeGrabpay           ChargePaymentMethodDetailsType = "grabpay"
	ChargePaymentMethodDetailsTypeIdeal             ChargePaymentMethodDetailsType = "ideal"
	ChargePaymentMethodDetailsTypeInteracPresent    ChargePaymentMethodDetailsType = "interac_present"
	ChargePaymentMethodDetailsTypeKlarna            ChargePaymentMethodDetailsType = "klarna"
	ChargePaymentMethodDetailsTypeMultibanco        ChargePaymentMethodDetailsType = "multibanco"
	ChargePaymentMethodDetailsTypeP24               ChargePaymentMethodDetailsType = "p24"
	ChargePaymentMethodDetailsTypeSepaDebit         ChargePaymentMethodDetailsType = "sepa_debit"
	ChargePaymentMethodDetailsTypeSofort            ChargePaymentMethodDetailsType = "sofort"
	ChargePaymentMethodDetailsTypeStripeAccount     ChargePaymentMethodDetailsType = "stripe_account"
	ChargePaymentMethodDetailsTypeWechat            ChargePaymentMethodDetailsType = "wechat"
)

// Returns a list of charges you've previously created. The charges are returned in sorted order, with the most recent charges appearing first.
type ChargeListParams struct {
	ListParams    `form:"*"`
	Created       *int64            `form:"created"`
	CreatedRange  *RangeQueryParams `form:"created"`
	Customer      *string           `form:"customer"`
	PaymentIntent *string           `form:"payment_intent"`
	TransferGroup *string           `form:"transfer_group"`
}
type DestinationParams struct {
	Account *string `form:"account"`
	Amount  *int64  `form:"amount"`
}

// An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://stripe.com/docs/connect/destination-charges) for details.
type ChargeTransferDataParams struct {
	Amount *int64 `form:"amount"`
	// This parameter can only be used on Charge creation.
	Destination *string `form:"destination"`
}
type ChargeLevel3LineItemsParams struct {
	DiscountAmount     *int64  `form:"discount_amount"`
	ProductCode        *string `form:"product_code"`
	ProductDescription *string `form:"product_description"`
	Quantity           *int64  `form:"quantity"`
	TaxAmount          *int64  `form:"tax_amount"`
	UnitCost           *int64  `form:"unit_cost"`
}
type ChargeLevel3Params struct {
	CustomerReference  *string                        `form:"customer_reference"`
	LineItems          []*ChargeLevel3LineItemsParams `form:"line_items"`
	MerchantReference  *string                        `form:"merchant_reference"`
	ShippingAddressZip *string                        `form:"shipping_address_zip"`
	ShippingAmount     *int64                         `form:"shipping_amount"`
	ShippingFromZip    *string                        `form:"shipping_from_zip"`
}

// To charge a credit card or other payment source, you create a Charge object. If your API key is in test mode, the supplied payment source (e.g., card) won't actually be charged, although everything else will occur as if in live mode. (Stripe assumes that the charge would have completed successfully).
type ChargeParams struct {
	Params                    `form:"*"`
	Amount                    *int64                    `form:"amount"`
	ApplicationFee            *int64                    `form:"application_fee"`
	ApplicationFeeAmount      *int64                    `form:"application_fee_amount"`
	Capture                   *bool                     `form:"capture"`
	Currency                  *string                   `form:"currency"`
	Customer                  *string                   `form:"customer"`
	Description               *string                   `form:"description"`
	Destination               *DestinationParams        `form:"destination"`
	ExchangeRate              *float64                  `form:"exchange_rate"`
	FraudDetails              *FraudDetailsParams       `form:"fraud_details"`
	Level3                    *ChargeLevel3Params       `form:"level3"`
	OnBehalfOf                *string                   `form:"on_behalf_of"`
	ReceiptEmail              *string                   `form:"receipt_email"`
	Shipping                  *ShippingDetailsParams    `form:"shipping"`
	Source                    *SourceParams             `form:"*"` // SourceParams has custom encoding so brought to top level with "*"
	StatementDescriptor       *string                   `form:"statement_descriptor"`
	StatementDescriptorSuffix *string                   `form:"statement_descriptor_suffix"`
	TransferData              *ChargeTransferDataParams `form:"transfer_data"`
	TransferGroup             *string                   `form:"transfer_group"`
}

// SetSource adds valid sources to a ChargeParams object,
// returning an error for unsupported sources.
func (p *ChargeParams) SetSource(sp interface{}) error {
	source, err := SourceParamsFor(sp)
	p.Source = source
	return err
}

// A set of key-value pairs you can attach to a charge giving information about its riskiness. If you believe a charge is fraudulent, include a `user_report` key with a value of `fraudulent`. If you believe a charge is safe, include a `user_report` key with a value of `safe`. Stripe will use the information you send to improve our fraud detection algorithms.
type FraudDetailsParams struct {
	UserReport *string `form:"user_report"`
}

// Capture the payment of an existing, uncaptured, charge. This is the second half of the two-step payment flow, where first you [created a charge](https://stripe.com/docs/api#create_charge) with the capture option set to false.
//
// Uncaptured payments expire a set number of days after they are created ([7 by default](https://stripe.com/docs/charges/placing-a-hold)). If they are not captured by that point in time, they will be marked as refunded and will no longer be capturable.
type CaptureParams struct {
	Params                    `form:"*"`
	Amount                    *int64                    `form:"amount"`
	ApplicationFee            *int64                    `form:"application_fee"`
	ApplicationFeeAmount      *int64                    `form:"application_fee_amount"`
	ExchangeRate              *float64                  `form:"exchange_rate"`
	ReceiptEmail              *string                   `form:"receipt_email"`
	StatementDescriptor       *string                   `form:"statement_descriptor"`
	StatementDescriptorSuffix *string                   `form:"statement_descriptor_suffix"`
	TransferData              *ChargeTransferDataParams `form:"transfer_data"`
	TransferGroup             *string                   `form:"transfer_group"`
}

// Information on fraud assessments for the charge.
type FraudDetails struct {
	StripeReport ChargeFraudStripeReport `json:"stripe_report"`
	UserReport   ChargeFraudUserReport   `json:"user_report"`
}
type ChargeLevel3LineItem struct {
	DiscountAmount     int64  `json:"discount_amount"`
	ProductCode        string `json:"product_code"`
	ProductDescription string `json:"product_description"`
	Quantity           int64  `json:"quantity"`
	TaxAmount          int64  `json:"tax_amount"`
	UnitCost           int64  `json:"unit_cost"`
}
type ChargeLevel3 struct {
	CustomerReference  string                  `json:"customer_reference"`
	LineItems          []*ChargeLevel3LineItem `json:"line_items"`
	MerchantReference  string                  `json:"merchant_reference"`
	ShippingAddressZip string                  `json:"shipping_address_zip"`
	ShippingAmount     int64                   `json:"shipping_amount"`
	ShippingFromZip    string                  `json:"shipping_from_zip"`
}

// The ID of the Radar rule that matched the payment, if applicable.
type ChargeOutcomeRule struct {
	Action    string `json:"action"`
	ID        string `json:"id"`
	Predicate string `json:"predicate"`
}

// Details about whether the payment was accepted, and why. See [understanding declines](https://stripe.com/docs/declines) for details.
type ChargeOutcome struct {
	NetworkStatus string             `json:"network_status"`
	Reason        string             `json:"reason"`
	RiskLevel     string             `json:"risk_level"`
	RiskScore     int64              `json:"risk_score"`
	Rule          *ChargeOutcomeRule `json:"rule"`
	SellerMessage string             `json:"seller_message"`
	Type          string             `json:"type"`
}

// UnmarshalJSON handles deserialization of a ChargeOutcomeRule.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *ChargeOutcomeRule) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type chargeOutcomeRule ChargeOutcomeRule
	var v chargeOutcomeRule
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = ChargeOutcomeRule(v)
	return nil
}

type ChargePaymentMethodDetailsAchCreditTransfer struct {
	AccountNumber string `json:"account_number"`
	BankName      string `json:"bank_name"`
	RoutingNumber string `json:"routing_number"`
	SwiftCode     string `json:"swift_code"`
}
type ChargePaymentMethodDetailsAchDebit struct {
	AccountHolderType BankAccountAccountHolderType `json:"account_holder_type"`
	BankName          string                       `json:"bank_name"`
	Country           string                       `json:"country"`
	Fingerprint       string                       `json:"fingerprint"`
	Last4             string                       `json:"last4"`
	RoutingNumber     string                       `json:"routing_number"`
}
type ChargePaymentMethodDetailsAcssDebit struct {
	BankName          string `json:"bank_name"`
	Fingerprint       string `json:"fingerprint"`
	InstitutionNumber string `json:"institution_number"`
	Last4             string `json:"last4"`
	Mandate           string `json:"mandate"`
	TransitNumber     string `json:"transit_number"`
}
type ChargePaymentMethodDetailsAfterpayClearpay struct {
	Reference string `json:"reference"`
}
type ChargePaymentMethodDetailsAlipay struct {
	BuyerID       string `json:"buyer_id"`
	Fingerprint   string `json:"fingerprint"`
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsAUBECSDebit struct {
	BSBNumber   string `json:"bsb_number"`
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
	Mandate     string `json:"mandate"`
}
type ChargePaymentMethodDetailsBACSDebit struct {
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
	Mandate     string `json:"mandate"`
	SortCode    string `json:"sort_code"`
}
type ChargePaymentMethodDetailsBancontact struct {
	BankCode                  string         `json:"bank_code"`
	BankName                  string         `json:"bank_name"`
	Bic                       string         `json:"bic"`
	GeneratedSepaDebit        *PaymentMethod `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate       `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string         `json:"iban_last4"`
	PreferredLanguage         string         `json:"preferred_language"`
	VerifiedName              string         `json:"verified_name"`
}
type ChargePaymentMethodDetailsBoleto struct {
	TaxID string `json:"tax_id"`
}

// Check results by Card networks on Card address and CVC at time of payment.
type ChargePaymentMethodDetailsCardChecks struct {
	AddressLine1Check      CardVerification `json:"address_line1_check"`
	AddressPostalCodeCheck CardVerification `json:"address_postal_code_check"`
	CVCCheck               CardVerification `json:"cvc_check"`
}

// Installment details for this payment (Mexico only).
//
// For more information, see the [installments integration guide](https://stripe.com/docs/payments/installments).
type ChargePaymentMethodDetailsCardInstallments struct {
	Plan *PaymentIntentPaymentMethodOptionsCardInstallmentsPlan `json:"plan"`
}

// Populated if this transaction used 3D Secure authentication.
type ChargePaymentMethodDetailsCardThreeDSecure struct {
	AuthenticationFlow ChargePaymentMethodDetailsCardThreeDSecureAuthenticationFlow `json:"authentication_flow"`
	Result             ChargePaymentMethodDetailsCardThreeDSecureResult             `json:"result"`
	ResultReason       ChargePaymentMethodDetailsCardThreeDSecureResultReason       `json:"result_reason"`
	Version            string                                                       `json:"version"`
}
type ChargePaymentMethodDetailsCardWalletAmexExpressCheckout struct{}
type ChargePaymentMethodDetailsCardWalletApplePay struct{}
type ChargePaymentMethodDetailsCardWalletGooglePay struct{}
type ChargePaymentMethodDetailsCardWalletMasterpass struct {
	BillingAddress  *Address `json:"billing_address"`
	Email           string   `json:"email"`
	Name            string   `json:"name"`
	ShippingAddress *Address `json:"shipping_address"`
}
type ChargePaymentMethodDetailsCardWalletSamsungPay struct{}
type ChargePaymentMethodDetailsCardWalletVisaCheckout struct {
	BillingAddress  *Address `json:"billing_address"`
	Email           string   `json:"email"`
	Name            string   `json:"name"`
	ShippingAddress *Address `json:"shipping_address"`
}

// If this Card is part of a card wallet, this contains the details of the card wallet.
type ChargePaymentMethodDetailsCardWallet struct {
	AmexExpressCheckout *ChargePaymentMethodDetailsCardWalletAmexExpressCheckout `json:"amex_express_checkout"`
	ApplePay            *ChargePaymentMethodDetailsCardWalletApplePay            `json:"apple_pay"`
	DynamicLast4        string                                                   `json:"dynamic_last4"`
	GooglePay           *ChargePaymentMethodDetailsCardWalletGooglePay           `json:"google_pay"`
	Masterpass          *ChargePaymentMethodDetailsCardWalletMasterpass          `json:"masterpass"`
	SamsungPay          *ChargePaymentMethodDetailsCardWalletSamsungPay          `json:"samsung_pay"`
	Type                PaymentMethodCardWalletType                              `json:"type"`
	VisaCheckout        *ChargePaymentMethodDetailsCardWalletVisaCheckout        `json:"visa_checkout"`
}
type ChargePaymentMethodDetailsCard struct {
	Brand        PaymentMethodCardBrand                      `json:"brand"`
	Checks       *ChargePaymentMethodDetailsCardChecks       `json:"checks"`
	Country      string                                      `json:"country"`
	ExpMonth     uint64                                      `json:"exp_month"`
	ExpYear      uint64                                      `json:"exp_year"`
	Fingerprint  string                                      `json:"fingerprint"`
	Funding      CardFunding                                 `json:"funding"`
	Installments *ChargePaymentMethodDetailsCardInstallments `json:"installments"`
	Last4        string                                      `json:"last4"`
	MOTO         bool                                        `json:"moto"`
	Network      PaymentMethodCardNetwork                    `json:"network"`
	ThreeDSecure *ChargePaymentMethodDetailsCardThreeDSecure `json:"three_d_secure"`
	Wallet       *ChargePaymentMethodDetailsCardWallet       `json:"wallet"`

	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	Description string `json:"description"`
	IIN         string `json:"iin"`
	Issuer      string `json:"issuer"`
}

// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
type ChargePaymentMethodDetailsCardPresentReceipt struct {
	AccountType                  ChargePaymentMethodDetailsCardPresentReceiptAccountType `json:"account_type"`
	ApplicationCryptogram        string                                                  `json:"application_cryptogram"`
	ApplicationPreferredName     string                                                  `json:"application_preferred_name"`
	AuthorizationCode            string                                                  `json:"authorization_code"`
	AuthorizationResponseCode    string                                                  `json:"authorization_response_code"`
	CardholderVerificationMethod string                                                  `json:"cardholder_verification_method"`
	DedicatedFileName            string                                                  `json:"dedicated_file_name"`
	TerminalVerificationResults  string                                                  `json:"terminal_verification_results"`
	TransactionStatusInformation string                                                  `json:"transaction_status_information"`
}
type ChargePaymentMethodDetailsCardPresent struct {
	AmountAuthorized     int64                                         `json:"amount_authorized"`
	Brand                PaymentMethodCardBrand                        `json:"brand"`
	CardholderName       string                                        `json:"cardholder_name"`
	Country              string                                        `json:"country"`
	EmvAuthData          string                                        `json:"emv_auth_data"`
	ExpMonth             uint64                                        `json:"exp_month"`
	ExpYear              uint64                                        `json:"exp_year"`
	Fingerprint          string                                        `json:"fingerprint"`
	Funding              CardFunding                                   `json:"funding"`
	GeneratedCard        string                                        `json:"generated_card"`
	Last4                string                                        `json:"last4"`
	Network              PaymentMethodCardNetwork                      `json:"network"`
	OvercaptureSupported bool                                          `json:"overcapture_supported"`
	ReadMethod           string                                        `json:"read_method"`
	Receipt              *ChargePaymentMethodDetailsCardPresentReceipt `json:"receipt"`

	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	Description string `json:"description"`
	IIN         string `json:"iin"`
	Issuer      string `json:"issuer"`
}
type ChargePaymentMethodDetailsEps struct {
	Bank         string `json:"bank"`
	VerifiedName string `json:"verified_name"`
}
type ChargePaymentMethodDetailsFPX struct {
	AccountHolderType PaymentMethodFPXAccountHolderType `json:"account_holder_type"`
	Bank              string                            `json:"bank"`
	TransactionID     string                            `json:"transaction_id"`
}
type ChargePaymentMethodDetailsGiropay struct {
	BankCode     string `json:"bank_code"`
	BankName     string `json:"bank_name"`
	Bic          string `json:"bic"`
	VerifiedName string `json:"verified_name"`
}
type ChargePaymentMethodDetailsGrabpay struct {
	TransactionID string `json:"transaction_id"`
}
type ChargePaymentMethodDetailsIdeal struct {
	Bank                      string         `json:"bank"`
	Bic                       string         `json:"bic"`
	GeneratedSepaDebit        *PaymentMethod `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate       `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string         `json:"iban_last4"`
	VerifiedName              string         `json:"verified_name"`
}

// A collection of fields required to be displayed on receipts. Only required for EMV transactions.
type ChargePaymentMethodDetailsInteracPresentReceipt struct {
	AccountType                  string `json:"account_type"`
	ApplicationCryptogram        string `json:"application_cryptogram"`
	ApplicationPreferredName     string `json:"application_preferred_name"`
	AuthorizationCode            string `json:"authorization_code"`
	AuthorizationResponseCode    string `json:"authorization_response_code"`
	CardholderVerificationMethod string `json:"cardholder_verification_method"`
	DedicatedFileName            string `json:"dedicated_file_name"`
	TerminalVerificationResults  string `json:"terminal_verification_results"`
	TransactionStatusInformation string `json:"transaction_status_information"`
}
type ChargePaymentMethodDetailsInteracPresent struct {
	Brand            string                                           `json:"brand"`
	CardholderName   string                                           `json:"cardholder_name"`
	Country          string                                           `json:"country"`
	EmvAuthData      string                                           `json:"emv_auth_data"`
	ExpMonth         int64                                            `json:"exp_month"`
	ExpYear          int64                                            `json:"exp_year"`
	Fingerprint      string                                           `json:"fingerprint"`
	Funding          string                                           `json:"funding"`
	GeneratedCard    string                                           `json:"generated_card"`
	Last4            string                                           `json:"last4"`
	Network          string                                           `json:"network"`
	PreferredLocales []string                                         `json:"preferred_locales"`
	ReadMethod       string                                           `json:"read_method"`
	Receipt          *ChargePaymentMethodDetailsInteracPresentReceipt `json:"receipt"`

	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	Description string `json:"description"`
	IIN         string `json:"iin"`
	Issuer      string `json:"issuer"`
}
type ChargePaymentMethodDetailsKlarna struct {
	PaymentMethodCategory ChargePaymentMethodDetailsKlarnaPaymentMethodCategory `json:"payment_method_category"`
	PreferredLocale       string                                                `json:"preferred_locale"`
}
type ChargePaymentMethodDetailsMultibanco struct {
	Entity    string `json:"entity"`
	Reference string `json:"reference"`
}
type ChargePaymentMethodDetailsOXXO struct {
	Number string `json:"number"`
}
type ChargePaymentMethodDetailsP24 struct {
	Bank         string `json:"bank"`
	Reference    string `json:"reference"`
	VerifiedName string `json:"verified_name"`
}
type ChargePaymentMethodDetailsSepaCreditTransfer struct {
	BankName string `json:"bank_name"`
	Bic      string `json:"bic"`
	Iban     string `json:"iban"`
}
type ChargePaymentMethodDetailsSepaDebit struct {
	BankCode    string   `json:"bank_code"`
	BranchCode  string   `json:"branch_code"`
	Country     string   `json:"country"`
	Fingerprint string   `json:"fingerprint"`
	Last4       string   `json:"last4"`
	Mandate     *Mandate `json:"mandate"`
}
type ChargePaymentMethodDetailsSofort struct {
	BankCode                  string         `json:"bank_code"`
	BankName                  string         `json:"bank_name"`
	Bic                       string         `json:"bic"`
	Country                   string         `json:"country"`
	GeneratedSepaDebit        *PaymentMethod `json:"generated_sepa_debit"`
	GeneratedSepaDebitMandate *Mandate       `json:"generated_sepa_debit_mandate"`
	IbanLast4                 string         `json:"iban_last4"`
	PreferredLanguage         string         `json:"preferred_language"`
	VerifiedName              string         `json:"verified_name"`
}
type ChargePaymentMethodDetailsStripeAccount struct{}
type ChargePaymentMethodDetailsWechat struct{}
type ChargePaymentMethodDetailsWechatPay struct {
	Fingerprint   string `json:"fingerprint"`
	TransactionID string `json:"transaction_id"`
}

// Details about the payment method at the time of the transaction.
type ChargePaymentMethodDetails struct {
	AchCreditTransfer  *ChargePaymentMethodDetailsAchCreditTransfer  `json:"ach_credit_transfer"`
	AchDebit           *ChargePaymentMethodDetailsAchDebit           `json:"ach_debit"`
	AcssDebit          *ChargePaymentMethodDetailsAcssDebit          `json:"acss_debit"`
	AfterpayClearpay   *ChargePaymentMethodDetailsAfterpayClearpay   `json:"afterpay_clearpay"`
	Alipay             *ChargePaymentMethodDetailsAlipay             `json:"alipay"`
	AUBECSDebit        *ChargePaymentMethodDetailsAUBECSDebit        `json:"au_becs_debit"`
	BACSDebit          *ChargePaymentMethodDetailsBACSDebit          `json:"bacs_debit"`
	Bancontact         *ChargePaymentMethodDetailsBancontact         `json:"bancontact"`
	Boleto             *ChargePaymentMethodDetailsBoleto             `json:"boleto"`
	Card               *ChargePaymentMethodDetailsCard               `json:"card"`
	CardPresent        *ChargePaymentMethodDetailsCardPresent        `json:"card_present"`
	Eps                *ChargePaymentMethodDetailsEps                `json:"eps"`
	FPX                *ChargePaymentMethodDetailsFPX                `json:"fpx"`
	Giropay            *ChargePaymentMethodDetailsGiropay            `json:"giropay"`
	Grabpay            *ChargePaymentMethodDetailsGrabpay            `json:"grabpay"`
	Ideal              *ChargePaymentMethodDetailsIdeal              `json:"ideal"`
	InteracPresent     *ChargePaymentMethodDetailsInteracPresent     `json:"interac_present"`
	Klarna             *ChargePaymentMethodDetailsKlarna             `json:"klarna"`
	Multibanco         *ChargePaymentMethodDetailsMultibanco         `json:"multibanco"`
	OXXO               *ChargePaymentMethodDetailsOXXO               `json:"oxxo"`
	P24                *ChargePaymentMethodDetailsP24                `json:"p24"`
	SepaCreditTransfer *ChargePaymentMethodDetailsSepaCreditTransfer `json:"sepa_credit_transfer"`
	SepaDebit          *ChargePaymentMethodDetailsSepaDebit          `json:"sepa_debit"`
	Sofort             *ChargePaymentMethodDetailsSofort             `json:"sofort"`
	StripeAccount      *ChargePaymentMethodDetailsStripeAccount      `json:"stripe_account"`
	Type               ChargePaymentMethodDetailsType                `json:"type"`
	Wechat             *ChargePaymentMethodDetailsWechat             `json:"wechat"`
	WechatPay          *ChargePaymentMethodDetailsWechatPay          `json:"wechat_pay"`
}

// An optional dictionary including the account to automatically transfer to as part of a destination charge. [See the Connect documentation](https://stripe.com/docs/connect/destination-charges) for details.
type ChargeTransferData struct {
	Amount      int64    `json:"amount"`
	Destination *Account `json:"destination"`
}

// To charge a credit or a debit card, you create a `Charge` object. You can
// retrieve and refund individual charges as well as list all charges. Charges
// are identified by a unique, random ID.
//
// Related guide: [Accept a payment with the Charges API](https://stripe.com/docs/payments/accept-a-payment-charges).
type Charge struct {
	APIResource
	Amount                        int64                       `json:"amount"`
	AmountCaptured                int64                       `json:"amount_captured"`
	AmountRefunded                int64                       `json:"amount_refunded"`
	Application                   *Application                `json:"application"`
	ApplicationFee                *ApplicationFee             `json:"application_fee"`
	ApplicationFeeAmount          int64                       `json:"application_fee_amount"`
	AuthorizationCode             string                      `json:"authorization_code"`
	BalanceTransaction            *BalanceTransaction         `json:"balance_transaction"`
	BillingDetails                *BillingDetails             `json:"billing_details"`
	CalculatedStatementDescriptor string                      `json:"calculated_statement_descriptor"`
	Captured                      bool                        `json:"captured"`
	Created                       int64                       `json:"created"`
	Currency                      Currency                    `json:"currency"`
	Customer                      *Customer                   `json:"customer"`
	Description                   string                      `json:"description"`
	Destination                   *Account                    `json:"destination"`
	Dispute                       *Dispute                    `json:"dispute"`
	Disputed                      bool                        `json:"disputed"`
	FailureCode                   string                      `json:"failure_code"`
	FailureMessage                string                      `json:"failure_message"`
	FraudDetails                  *FraudDetails               `json:"fraud_details"`
	ID                            string                      `json:"id"`
	Invoice                       *Invoice                    `json:"invoice"`
	Level3                        ChargeLevel3                `json:"level3"`
	Livemode                      bool                        `json:"livemode"`
	Metadata                      map[string]string           `json:"metadata"`
	Object                        string                      `json:"object"`
	OnBehalfOf                    *Account                    `json:"on_behalf_of"`
	Order                         *Order                      `json:"order"`
	Outcome                       *ChargeOutcome              `json:"outcome"`
	Paid                          bool                        `json:"paid"`
	PaymentIntent                 *PaymentIntent              `json:"payment_intent"`
	PaymentMethod                 string                      `json:"payment_method"`
	PaymentMethodDetails          *ChargePaymentMethodDetails `json:"payment_method_details"`
	ReceiptEmail                  string                      `json:"receipt_email"`
	ReceiptNumber                 string                      `json:"receipt_number"`
	ReceiptURL                    string                      `json:"receipt_url"`
	Refunded                      bool                        `json:"refunded"`
	Refunds                       *RefundList                 `json:"refunds"`
	Review                        *Review                     `json:"review"`
	Shipping                      *ShippingDetails            `json:"shipping"`
	Source                        *PaymentSource              `json:"source"`
	SourceTransfer                *Transfer                   `json:"source_transfer"`
	StatementDescriptor           string                      `json:"statement_descriptor"`
	StatementDescriptorSuffix     string                      `json:"statement_descriptor_suffix"`
	Status                        string                      `json:"status"`
	Transfer                      *Transfer                   `json:"transfer"`
	TransferData                  *ChargeTransferData         `json:"transfer_data"`
	TransferGroup                 string                      `json:"transfer_group"`
}

// ChargeList is a list of Charges as retrieved from a list endpoint.
type ChargeList struct {
	APIResource
	ListMeta
	Data []*Charge `json:"data"`
}

// UnmarshalJSON handles deserialization of a Charge.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (c *Charge) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		c.ID = id
		return nil
	}

	type charge Charge
	var v charge
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*c = Charge(v)
	return nil
}
