//
//
// File generated from our OpenAPI spec
//
//

package stripe

import "encoding/json"

// Card brand. Can be `amex`, `diners`, `discover`, `jcb`, `mastercard`, `unionpay`, `visa`, or `unknown`.
type PaymentMethodCardBrand string

// List of values that PaymentMethodCardBrand can take
const (
	PaymentMethodCardBrandAmex       PaymentMethodCardBrand = "amex"
	PaymentMethodCardBrandDiners     PaymentMethodCardBrand = "diners"
	PaymentMethodCardBrandDiscover   PaymentMethodCardBrand = "discover"
	PaymentMethodCardBrandJCB        PaymentMethodCardBrand = "jcb"
	PaymentMethodCardBrandMastercard PaymentMethodCardBrand = "mastercard"
	PaymentMethodCardBrandUnionpay   PaymentMethodCardBrand = "unionpay"
	PaymentMethodCardBrandUnknown    PaymentMethodCardBrand = "unknown"
	PaymentMethodCardBrandVisa       PaymentMethodCardBrand = "visa"
)

// All available networks for the card.
type PaymentMethodCardNetwork string

// List of values that PaymentMethodCardNetwork can take
const (
	PaymentMethodCardNetworkAmex       PaymentMethodCardNetwork = "amex"
	PaymentMethodCardNetworkDiners     PaymentMethodCardNetwork = "diners"
	PaymentMethodCardNetworkDiscover   PaymentMethodCardNetwork = "discover"
	PaymentMethodCardNetworkInterac    PaymentMethodCardNetwork = "interac"
	PaymentMethodCardNetworkJCB        PaymentMethodCardNetwork = "jcb"
	PaymentMethodCardNetworkMastercard PaymentMethodCardNetwork = "mastercard"
	PaymentMethodCardNetworkUnionpay   PaymentMethodCardNetwork = "unionpay"
	PaymentMethodCardNetworkUnknown    PaymentMethodCardNetwork = "unknown"
	PaymentMethodCardNetworkVisa       PaymentMethodCardNetwork = "visa"
)

// The type of the card wallet, one of `amex_express_checkout`, `apple_pay`, `google_pay`, `masterpass`, `samsung_pay`, or `visa_checkout`. An additional hash is included on the Wallet subhash with a name matching this value. It contains additional information specific to the card wallet type.
type PaymentMethodCardWalletType string

// List of values that PaymentMethodCardWalletType can take
const (
	PaymentMethodCardWalletTypeAmexExpressCheckout PaymentMethodCardWalletType = "amex_express_checkout"
	PaymentMethodCardWalletTypeApplePay            PaymentMethodCardWalletType = "apple_pay"
	PaymentMethodCardWalletTypeGooglePay           PaymentMethodCardWalletType = "google_pay"
	PaymentMethodCardWalletTypeMasterpass          PaymentMethodCardWalletType = "masterpass"
	PaymentMethodCardWalletTypeSamsungPay          PaymentMethodCardWalletType = "samsung_pay"
	PaymentMethodCardWalletTypeVisaCheckout        PaymentMethodCardWalletType = "visa_checkout"
)

// Account holder type, if provided. Can be one of `individual` or `company`.
type PaymentMethodFPXAccountHolderType string

// List of values that PaymentMethodFPXAccountHolderType can take
const (
	PaymentMethodFPXAccountHolderTypeCompany    PaymentMethodFPXAccountHolderType = "company"
	PaymentMethodFPXAccountHolderTypeIndividual PaymentMethodFPXAccountHolderType = "individual"
)

// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
type PaymentMethodType string

// List of values that PaymentMethodType can take
const (
	PaymentMethodTypeACSSDebit        PaymentMethodType = "acss_debit"
	PaymentMethodTypeAffirm           PaymentMethodType = "affirm"
	PaymentMethodTypeAfterpayClearpay PaymentMethodType = "afterpay_clearpay"
	PaymentMethodTypeAlipay           PaymentMethodType = "alipay"
	PaymentMethodTypeAUBECSDebit      PaymentMethodType = "au_becs_debit"
	PaymentMethodTypeBACSDebit        PaymentMethodType = "bacs_debit"
	PaymentMethodTypeBancontact       PaymentMethodType = "bancontact"
	PaymentMethodTypeBLIK             PaymentMethodType = "blik"
	PaymentMethodTypeBoleto           PaymentMethodType = "boleto"
	PaymentMethodTypeCard             PaymentMethodType = "card"
	PaymentMethodTypeCardPresent      PaymentMethodType = "card_present"
	PaymentMethodTypeCustomerBalance  PaymentMethodType = "customer_balance"
	PaymentMethodTypeEPS              PaymentMethodType = "eps"
	PaymentMethodTypeFPX              PaymentMethodType = "fpx"
	PaymentMethodTypeGiropay          PaymentMethodType = "giropay"
	PaymentMethodTypeGrabpay          PaymentMethodType = "grabpay"
	PaymentMethodTypeIdeal            PaymentMethodType = "ideal"
	PaymentMethodTypeInteracPresent   PaymentMethodType = "interac_present"
	PaymentMethodTypeKlarna           PaymentMethodType = "klarna"
	PaymentMethodTypeKonbini          PaymentMethodType = "konbini"
	PaymentMethodTypeLink             PaymentMethodType = "link"
	PaymentMethodTypeOXXO             PaymentMethodType = "oxxo"
	PaymentMethodTypeP24              PaymentMethodType = "p24"
	PaymentMethodTypePayNow           PaymentMethodType = "paynow"
	PaymentMethodTypePromptPay        PaymentMethodType = "promptpay"
	PaymentMethodTypeSepaDebit        PaymentMethodType = "sepa_debit"
	PaymentMethodTypeSofort           PaymentMethodType = "sofort"
	PaymentMethodTypeUSBankAccount    PaymentMethodType = "us_bank_account"
	PaymentMethodTypeWechatPay        PaymentMethodType = "wechat_pay"
)

// Account holder type: individual or company.
type PaymentMethodUSBankAccountAccountHolderType string

// List of values that PaymentMethodUSBankAccountAccountHolderType can take
const (
	PaymentMethodUSBankAccountAccountHolderTypeCompany    PaymentMethodUSBankAccountAccountHolderType = "company"
	PaymentMethodUSBankAccountAccountHolderTypeIndividual PaymentMethodUSBankAccountAccountHolderType = "individual"
)

// Account type: checkings or savings. Defaults to checking if omitted.
type PaymentMethodUSBankAccountAccountType string

// List of values that PaymentMethodUSBankAccountAccountType can take
const (
	PaymentMethodUSBankAccountAccountTypeChecking PaymentMethodUSBankAccountAccountType = "checking"
	PaymentMethodUSBankAccountAccountTypeSavings  PaymentMethodUSBankAccountAccountType = "savings"
)

// All supported networks.
type PaymentMethodUSBankAccountNetworksSupported string

// List of values that PaymentMethodUSBankAccountNetworksSupported can take
const (
	PaymentMethodUSBankAccountNetworksSupportedAch            PaymentMethodUSBankAccountNetworksSupported = "ach"
	PaymentMethodUSBankAccountNetworksSupportedUSDomesticWire PaymentMethodUSBankAccountNetworksSupported = "us_domestic_wire"
)

// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
type PaymentMethodACSSDebitParams struct {
	// Customer's bank account number.
	AccountNumber *string `form:"account_number"`
	// Institution number of the customer's bank.
	InstitutionNumber *string `form:"institution_number"`
	// Transit number of the customer's bank.
	TransitNumber *string `form:"transit_number"`
}

// If this is an `affirm` PaymentMethod, this hash contains details about the Affirm payment method.
type PaymentMethodAffirmParams struct{}

// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
type PaymentMethodAfterpayClearpayParams struct{}

// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
type PaymentMethodAlipayParams struct{}

// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
type PaymentMethodAUBECSDebitParams struct {
	// The account number for the bank account.
	AccountNumber *string `form:"account_number"`
	// Bank-State-Branch number of the bank account.
	BSBNumber *string `form:"bsb_number"`
}

// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
type PaymentMethodBACSDebitParams struct {
	// Account number of the bank account that the funds will be debited from.
	AccountNumber *string `form:"account_number"`
	// Sort code of the bank account. (e.g., `10-20-30`)
	SortCode *string `form:"sort_code"`
}

// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
type PaymentMethodBancontactParams struct{}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type BillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address"`
	// Email address.
	Email *string `form:"email"`
	// Full name.
	Name *string `form:"name"`
	// Billing phone number (including extension).
	Phone *string `form:"phone"`
}

// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
type PaymentMethodBLIKParams struct{}

// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
type PaymentMethodBoletoParams struct {
	// The tax ID of the customer (CPF for individual consumers or CNPJ for businesses consumers)
	TaxID *string `form:"tax_id"`
}

// If this is a `card` PaymentMethod, this hash contains the user's card details. For backwards compatibility, you can alternatively provide a Stripe token (e.g., for Apple Pay, Amex Express Checkout, or legacy Checkout) into the card hash with format `card: {token: "tok_visa"}`. When providing a card number, you must meet the requirements for [PCI compliance](https://stripe.com/docs/security#validating-pci-compliance). We strongly recommend using Stripe.js instead of interacting with this API directly.
type PaymentMethodCardParams struct {
	// The card's CVC. It is highly recommended to always include this value.
	CVC *string `form:"cvc"`
	// Two-digit number representing the card's expiration month.
	ExpMonth *string `form:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear *string `form:"exp_year"`
	// The card number, as a string without any separators.
	Number *string `form:"number"`
	Token  *string `form:"token"`
}

// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
type PaymentMethodCustomerBalanceParams struct{}

// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
type PaymentMethodEPSParams struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
type PaymentMethodFPXParams struct {
	// Account holder type for FPX transaction
	AccountHolderType *string `form:"account_holder_type"`
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
type PaymentMethodGiropayParams struct{}

// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
type PaymentMethodGrabpayParams struct{}

// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
type PaymentMethodIdealParams struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
type PaymentMethodInteracPresentParams struct{}

// Customer's date of birth
type PaymentMethodKlarnaDOBParams struct {
	// The day of birth, between 1 and 31.
	Day *int64 `form:"day"`
	// The month of birth, between 1 and 12.
	Month *int64 `form:"month"`
	// The four-digit year of birth.
	Year *int64 `form:"year"`
}

// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
type PaymentMethodKlarnaParams struct {
	// Customer's date of birth
	DOB *PaymentMethodKlarnaDOBParams `form:"dob"`
}

// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
type PaymentMethodKonbiniParams struct{}

// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
type PaymentMethodLinkParams struct{}

// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
type PaymentMethodOXXOParams struct{}

// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
type PaymentMethodP24Params struct {
	// The customer's bank.
	Bank                *string `form:"bank"`
	TOSShownAndAccepted *bool   `form:"tos_shown_and_accepted"`
}

// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
type PaymentMethodPayNowParams struct{}

// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
type PaymentMethodPromptPayParams struct{}

// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
type PaymentMethodRadarOptionsParams struct {
	// A [Radar Session](https://stripe.com/docs/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session *string `form:"session"`
}

// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
type PaymentMethodSepaDebitParams struct {
	// IBAN of the bank account.
	Iban *string `form:"iban"`
}

// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
type PaymentMethodSofortParams struct {
	// Two-letter ISO code representing the country the bank account is located in.
	Country *string `form:"country"`
}

// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
type PaymentMethodUSBankAccountParams struct {
	// Bank account type.
	AccountHolderType *string `form:"account_holder_type"`
	// Account number of the bank account.
	AccountNumber *string `form:"account_number"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType *string `form:"account_type"`
	// The ID of a Financial Connections Account to use as a payment method.
	FinancialConnectionsAccount *string `form:"financial_connections_account"`
	// Routing number of the bank account.
	RoutingNumber *string `form:"routing_number"`
}

// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
type PaymentMethodWechatPayParams struct{}

// Creates a PaymentMethod object. Read the [Stripe.js reference](https://stripe.com/docs/stripe-js/reference#stripe-create-payment-method) to learn how to create PaymentMethods via Stripe.js.
//
// Instead of creating a PaymentMethod directly, we recommend using the [PaymentIntents API to accept a payment immediately or the <a href="/docs/payments/save-and-reuse">SetupIntent](https://stripe.com/docs/payments/accept-a-payment) API to collect payment method details ahead of a future payment.
type PaymentMethodParams struct {
	Params `form:"*"`
	// This is a legacy parameter that will be removed in the future. It is a hash that does not accept any keys.
	ACSSDebit *PaymentMethodACSSDebitParams `form:"acss_debit"`
	// This is a legacy parameter that will be removed in the future. It is a hash that does not accept any keys.
	Affirm *PaymentMethodAffirmParams `form:"affirm"`
	// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
	AfterpayClearpay *PaymentMethodAfterpayClearpayParams `form:"afterpay_clearpay"`
	// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
	Alipay *PaymentMethodAlipayParams `form:"alipay"`
	// This is a legacy parameter that will be removed in the future. It is a hash that does not accept any keys.
	AUBECSDebit *PaymentMethodAUBECSDebitParams `form:"au_becs_debit"`
	// This is a legacy parameter that will be removed in the future. It is a hash that does not accept any keys.
	BACSDebit *PaymentMethodBACSDebitParams `form:"bacs_debit"`
	// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
	Bancontact *PaymentMethodBancontactParams `form:"bancontact"`
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails *BillingDetailsParams `form:"billing_details"`
	// This is a legacy parameter that will be removed in the future. It is a hash that does not accept any keys.
	BLIK *PaymentMethodBLIKParams `form:"blik"`
	// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
	Boleto *PaymentMethodBoletoParams `form:"boleto"`
	// If this is a `card` PaymentMethod, this hash contains the user's card details. For backwards compatibility, you can alternatively provide a Stripe token (e.g., for Apple Pay, Amex Express Checkout, or legacy Checkout) into the card hash with format `card: {token: "tok_visa"}`. When providing a card number, you must meet the requirements for [PCI compliance](https://stripe.com/docs/security#validating-pci-compliance). We strongly recommend using Stripe.js instead of interacting with this API directly.
	Card *PaymentMethodCardParams `form:"card"`
	// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
	CustomerBalance *PaymentMethodCustomerBalanceParams `form:"customer_balance"`
	// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
	EPS *PaymentMethodEPSParams `form:"eps"`
	// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
	FPX *PaymentMethodFPXParams `form:"fpx"`
	// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
	Giropay *PaymentMethodGiropayParams `form:"giropay"`
	// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
	Grabpay *PaymentMethodGrabpayParams `form:"grabpay"`
	// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
	Ideal *PaymentMethodIdealParams `form:"ideal"`
	// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
	InteracPresent *PaymentMethodInteracPresentParams `form:"interac_present"`
	// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
	Klarna *PaymentMethodKlarnaParams `form:"klarna"`
	// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
	Konbini *PaymentMethodKonbiniParams `form:"konbini"`
	// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
	Link *PaymentMethodLinkParams `form:"link"`
	// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
	OXXO *PaymentMethodOXXOParams `form:"oxxo"`
	// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
	P24 *PaymentMethodP24Params `form:"p24"`
	// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
	PayNow *PaymentMethodPayNowParams `form:"paynow"`
	// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
	PromptPay *PaymentMethodPromptPayParams `form:"promptpay"`
	// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
	RadarOptions *PaymentMethodRadarOptionsParams `form:"radar_options"`
	// This is a legacy parameter that will be removed in the future. It is a hash that does not accept any keys.
	SepaDebit *PaymentMethodSepaDebitParams `form:"sepa_debit"`
	// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
	Sofort *PaymentMethodSofortParams `form:"sofort"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type"`
	// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
	USBankAccount *PaymentMethodUSBankAccountParams `form:"us_bank_account"`
	// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
	WechatPay *PaymentMethodWechatPayParams `form:"wechat_pay"`
	// The following parameters are used when cloning a PaymentMethod to the connected account
	// The `Customer` to whom the original PaymentMethod is attached.
	Customer *string `form:"customer"`
	// The PaymentMethod to share.
	PaymentMethod *string `form:"payment_method"`
}

// Returns a list of PaymentMethods attached to the StripeAccount. For listing a customer's payment methods, you should use [List a Customer's PaymentMethods](https://stripe.com/docs/api/payment_methods/customer_list)
type PaymentMethodListParams struct {
	ListParams `form:"*"`
	// The ID of the customer whose PaymentMethods will be retrieved.
	Customer *string `form:"customer"`
	// A required filter on the list, based on the object `type` field.
	Type *string `form:"type"`
}

// Attaches a PaymentMethod object to a Customer.
//
// To attach a new PaymentMethod to a customer for future payments, we recommend you use a [SetupIntent](https://stripe.com/docs/api/setup_intents)
// or a PaymentIntent with [setup_future_usage](https://stripe.com/docs/api/payment_intents/create#create_payment_intent-setup_future_usage).
// These approaches will perform any necessary steps to ensure that the PaymentMethod can be used in a future payment. Using the
// /v1/payment_methods/:id/attach endpoint does not ensure that future payments can be made with the attached PaymentMethod.
// See [Optimizing cards for future payments](https://stripe.com/docs/payments/payment-intents#future-usage) for more information about setting up future payments.
//
// To use this PaymentMethod as the default for invoice or subscription payments,
// set [invoice_settings.default_payment_method](https://stripe.com/docs/api/customers/update#update_customer-invoice_settings-default_payment_method),
// on the Customer to the PaymentMethod's ID.
type PaymentMethodAttachParams struct {
	Params `form:"*"`
	// The ID of the customer to which to attach the PaymentMethod.
	Customer *string `form:"customer"`
}

// Detaches a PaymentMethod object from a Customer. After a PaymentMethod is detached, it can no longer be used for a payment or re-attached to a Customer.
type PaymentMethodDetachParams struct {
	Params `form:"*"`
}
type PaymentMethodACSSDebit struct {
	// Name of the bank associated with the bank account.
	BankName string `json:"bank_name"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Institution number of the bank account.
	InstitutionNumber string `json:"institution_number"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// Transit number of the bank account.
	TransitNumber string `json:"transit_number"`
}
type PaymentMethodAffirm struct{}
type PaymentMethodAfterpayClearpay struct{}
type PaymentMethodAlipay struct{}
type PaymentMethodAUBECSDebit struct {
	// Six-digit number identifying bank and branch associated with this bank account.
	BSBNumber string `json:"bsb_number"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
}
type PaymentMethodBACSDebit struct {
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// Sort code of the bank account. (e.g., `10-20-30`)
	SortCode string `json:"sort_code"`
}
type PaymentMethodBancontact struct{}
type BillingDetails struct {
	// Billing address.
	Address *Address `json:"address"`
	// Email address.
	Email string `json:"email"`
	// Full name.
	Name string `json:"name"`
	// Billing phone number (including extension).
	Phone string `json:"phone"`
}
type PaymentMethodBLIK struct{}
type PaymentMethodBoleto struct {
	// Uniquely identifies the customer tax id (CNPJ or CPF)
	TaxID string `json:"tax_id"`
}

// Checks on Card address and CVC if provided.
type PaymentMethodCardChecks struct {
	// If a address line1 was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressLine1Check CardVerification `json:"address_line1_check"`
	// If a address postal code was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	AddressPostalCodeCheck CardVerification `json:"address_postal_code_check"`
	// If a CVC was provided, results of the check, one of `pass`, `fail`, `unavailable`, or `unchecked`.
	CVCCheck CardVerification `json:"cvc_check"`
}

// Contains information about card networks that can be used to process the payment.
type PaymentMethodCardNetworks struct {
	// All available networks for the card.
	Available []PaymentMethodCardNetwork `json:"available"`
	// The preferred network for the card.
	Preferred PaymentMethodCardNetwork `json:"preferred"`
}

// Contains details on how this Card maybe be used for 3D Secure authentication.
type PaymentMethodCardThreeDSecureUsage struct {
	// Whether 3D Secure is supported on this card.
	Supported bool `json:"supported"`
}
type PaymentMethodCardWalletAmexExpressCheckout struct{}
type PaymentMethodCardWalletApplePay struct{}
type PaymentMethodCardWalletGooglePay struct{}
type PaymentMethodCardWalletMasterpass struct {
	// Owner's verified billing address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	BillingAddress *Address `json:"billing_address"`
	// Owner's verified email. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Email string `json:"email"`
	// Owner's verified full name. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	Name string `json:"name"`
	// Owner's verified shipping address. Values are verified or provided by the wallet directly (if supported) at the time of authorization or settlement. They cannot be set or mutated.
	ShippingAddress *Address `json:"shipping_address"`
}
type PaymentMethodCardWalletSamsungPay struct{}
type PaymentMethodCardWalletVisaCheckout struct {
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
type PaymentMethodCardWallet struct {
	AmexExpressCheckout *PaymentMethodCardWalletAmexExpressCheckout `json:"amex_express_checkout"`
	ApplePay            *PaymentMethodCardWalletApplePay            `json:"apple_pay"`
	// (For tokenized numbers only.) The last four digits of the device account number.
	DynamicLast4 string                             `json:"dynamic_last4"`
	GooglePay    *PaymentMethodCardWalletGooglePay  `json:"google_pay"`
	Masterpass   *PaymentMethodCardWalletMasterpass `json:"masterpass"`
	SamsungPay   *PaymentMethodCardWalletSamsungPay `json:"samsung_pay"`
	// The type of the card wallet, one of `amex_express_checkout`, `apple_pay`, `google_pay`, `masterpass`, `samsung_pay`, or `visa_checkout`. An additional hash is included on the Wallet subhash with a name matching this value. It contains additional information specific to the card wallet type.
	Type         PaymentMethodCardWalletType          `json:"type"`
	VisaCheckout *PaymentMethodCardWalletVisaCheckout `json:"visa_checkout"`
}
type PaymentMethodCard struct {
	// Card brand. Can be `amex`, `diners`, `discover`, `jcb`, `mastercard`, `unionpay`, `visa`, or `unknown`.
	Brand PaymentMethodCardBrand `json:"brand"`
	// Checks on Card address and CVC if provided.
	Checks *PaymentMethodCardChecks `json:"checks"`
	// Two-letter ISO code representing the country of the card. You could use this attribute to get a sense of the international breakdown of cards you've collected.
	Country string `json:"country"`
	// Two-digit number representing the card's expiration month.
	ExpMonth uint64 `json:"exp_month"`
	// Four-digit number representing the card's expiration year.
	ExpYear uint64 `json:"exp_year"`
	// Uniquely identifies this particular card number. You can use this attribute to check whether two customers who've signed up with you are using the same card number, for example. For payment methods that tokenize card information (Apple Pay, Google Pay), the tokenized number might be provided instead of the underlying card number.
	//
	// *Starting May 1, 2021, card fingerprint in India for Connect will change to allow two fingerprints for the same card --- one for India and one for the rest of the world.*
	Fingerprint string `json:"fingerprint"`
	// Card funding type. Can be `credit`, `debit`, `prepaid`, or `unknown`.
	Funding CardFunding `json:"funding"`
	// The last four digits of the card.
	Last4 string `json:"last4"`
	// Contains information about card networks that can be used to process the payment.
	Networks *PaymentMethodCardNetworks `json:"networks"`
	// Contains details on how this Card maybe be used for 3D Secure authentication.
	ThreeDSecureUsage *PaymentMethodCardThreeDSecureUsage `json:"three_d_secure_usage"`
	// If this Card is part of a card wallet, this contains the details of the card wallet.
	Wallet *PaymentMethodCardWallet `json:"wallet"`
	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	// A high-level description of the type of cards issued in this range. (For internal use only and not typically available in standard API requests.)
	Description string `json:"description"`
	// Issuer identification number of the card. (For internal use only and not typically available in standard API requests.)
	IIN string `json:"iin"`
	// The name of the card's issuing bank. (For internal use only and not typically available in standard API requests.)
	Issuer string `json:"issuer"`
}
type PaymentMethodCardPresent struct{}
type PaymentMethodCustomerBalance struct{}
type PaymentMethodEPS struct {
	// The customer's bank. Should be one of `arzte_und_apotheker_bank`, `austrian_anadi_bank_ag`, `bank_austria`, `bankhaus_carl_spangler`, `bankhaus_schelhammer_und_schattera_ag`, `bawag_psk_ag`, `bks_bank_ag`, `brull_kallmus_bank_ag`, `btv_vier_lander_bank`, `capital_bank_grawe_gruppe_ag`, `dolomitenbank`, `easybank_ag`, `erste_bank_und_sparkassen`, `hypo_alpeadriabank_international_ag`, `hypo_noe_lb_fur_niederosterreich_u_wien`, `hypo_oberosterreich_salzburg_steiermark`, `hypo_tirol_bank_ag`, `hypo_vorarlberg_bank_ag`, `hypo_bank_burgenland_aktiengesellschaft`, `marchfelder_bank`, `oberbank_ag`, `raiffeisen_bankengruppe_osterreich`, `schoellerbank_ag`, `sparda_bank_wien`, `volksbank_gruppe`, `volkskreditbank_ag`, or `vr_bank_braunau`.
	Bank string `json:"bank"`
}
type PaymentMethodFPX struct {
	// Account holder type, if provided. Can be one of `individual` or `company`.
	AccountHolderType PaymentMethodFPXAccountHolderType `json:"account_holder_type"`
	// The customer's bank, if provided. Can be one of `affin_bank`, `agrobank`, `alliance_bank`, `ambank`, `bank_islam`, `bank_muamalat`, `bank_rakyat`, `bsn`, `cimb`, `hong_leong_bank`, `hsbc`, `kfh`, `maybank2u`, `ocbc`, `public_bank`, `rhb`, `standard_chartered`, `uob`, `deutsche_bank`, `maybank2e`, or `pb_enterprise`.
	Bank          string `json:"bank"`
	TransactionID string `json:"transaction_id"`
}
type PaymentMethodGiropay struct{}
type PaymentMethodGrabpay struct{}
type PaymentMethodIdeal struct {
	// The customer's bank, if provided. Can be one of `abn_amro`, `asn_bank`, `bunq`, `handelsbanken`, `ing`, `knab`, `moneyou`, `rabobank`, `regiobank`, `revolut`, `sns_bank`, `triodos_bank`, or `van_lanschot`.
	Bank string `json:"bank"`
	// The Bank Identifier Code of the customer's bank, if the bank was provided.
	Bic string `json:"bic"`
}
type PaymentMethodInteracPresent struct{}

// The customer's date of birth, if provided.
type PaymentMethodKlarnaDOB struct {
	// The day of birth, between 1 and 31.
	Day int64 `json:"day"`
	// The month of birth, between 1 and 12.
	Month int64 `json:"month"`
	// The four-digit year of birth.
	Year int64 `json:"year"`
}
type PaymentMethodKlarna struct {
	// The customer's date of birth, if provided.
	DOB *PaymentMethodKlarnaDOB `json:"dob"`
}
type PaymentMethodKonbini struct{}
type PaymentMethodLink struct {
	// Account owner's email address.
	Email string `json:"email"`
	// Token used for persistent Link logins.
	PersistentToken string `json:"persistent_token"`
}
type PaymentMethodOXXO struct{}
type PaymentMethodP24 struct {
	// The customer's bank, if provided.
	Bank string `json:"bank"`
}
type PaymentMethodPayNow struct{}
type PaymentMethodPromptPay struct{}

// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
type PaymentMethodRadarOptions struct {
	// A [Radar Session](https://stripe.com/docs/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session string `json:"session"`
}

// Information about the object that generated this PaymentMethod.
type PaymentMethodSepaDebitGeneratedFrom struct {
	// The ID of the Charge that generated this PaymentMethod, if any.
	Charge *Charge `json:"charge"`
	// The ID of the SetupAttempt that generated this PaymentMethod, if any.
	SetupAttempt *SetupAttempt `json:"setup_attempt"`
}
type PaymentMethodSepaDebit struct {
	// Bank code of bank associated with the bank account.
	BankCode string `json:"bank_code"`
	// Branch code of bank associated with the bank account.
	BranchCode string `json:"branch_code"`
	// Two-letter ISO code representing the country the bank account is located in.
	Country string `json:"country"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Information about the object that generated this PaymentMethod.
	GeneratedFrom *PaymentMethodSepaDebitGeneratedFrom `json:"generated_from"`
	// Last four characters of the IBAN.
	Last4 string `json:"last4"`
}
type PaymentMethodSofort struct {
	// Two-letter ISO code representing the country the bank account is located in.
	Country string `json:"country"`
}

// Contains information about US bank account networks that can be used.
type PaymentMethodUSBankAccountNetworks struct {
	// The preferred network.
	Preferred string `json:"preferred"`
	// All supported networks.
	Supported []PaymentMethodUSBankAccountNetworksSupported `json:"supported"`
}
type PaymentMethodUSBankAccount struct {
	// Account holder type: individual or company.
	AccountHolderType PaymentMethodUSBankAccountAccountHolderType `json:"account_holder_type"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType PaymentMethodUSBankAccountAccountType `json:"account_type"`
	// The name of the bank.
	BankName string `json:"bank_name"`
	// The ID of the Financial Connections Account used to create the payment method.
	FinancialConnectionsAccount string `json:"financial_connections_account"`
	// Uniquely identifies this particular bank account. You can use this attribute to check whether two bank accounts are the same.
	Fingerprint string `json:"fingerprint"`
	// Last four digits of the bank account number.
	Last4 string `json:"last4"`
	// Contains information about US bank account networks that can be used.
	Networks *PaymentMethodUSBankAccountNetworks `json:"networks"`
	// Routing number of the bank account.
	RoutingNumber string `json:"routing_number"`
}
type PaymentMethodWechatPay struct{}

// PaymentMethod objects represent your customer's payment instruments.
// You can use them with [PaymentIntents](https://stripe.com/docs/payments/payment-intents) to collect payments or save them to
// Customer objects to store instrument details for future payments.
//
// Related guides: [Payment Methods](https://stripe.com/docs/payments/payment-methods) and [More Payment Scenarios](https://stripe.com/docs/payments/more-payment-scenarios).
type PaymentMethod struct {
	APIResource
	ACSSDebit        *PaymentMethodACSSDebit        `json:"acss_debit"`
	Affirm           *PaymentMethodAffirm           `json:"affirm"`
	AfterpayClearpay *PaymentMethodAfterpayClearpay `json:"afterpay_clearpay"`
	Alipay           *PaymentMethodAlipay           `json:"alipay"`
	AUBECSDebit      *PaymentMethodAUBECSDebit      `json:"au_becs_debit"`
	BACSDebit        *PaymentMethodBACSDebit        `json:"bacs_debit"`
	Bancontact       *PaymentMethodBancontact       `json:"bancontact"`
	BillingDetails   *BillingDetails                `json:"billing_details"`
	BLIK             *PaymentMethodBLIK             `json:"blik"`
	Boleto           *PaymentMethodBoleto           `json:"boleto"`
	Card             *PaymentMethodCard             `json:"card"`
	CardPresent      *PaymentMethodCardPresent      `json:"card_present"`
	// Time at which the object was created. Measured in seconds since the Unix epoch.
	Created int64 `json:"created"`
	// The ID of the Customer to which this PaymentMethod is saved. This will not be set when the PaymentMethod has not been saved to a Customer.
	Customer        *Customer                     `json:"customer"`
	CustomerBalance *PaymentMethodCustomerBalance `json:"customer_balance"`
	EPS             *PaymentMethodEPS             `json:"eps"`
	FPX             *PaymentMethodFPX             `json:"fpx"`
	Giropay         *PaymentMethodGiropay         `json:"giropay"`
	Grabpay         *PaymentMethodGrabpay         `json:"grabpay"`
	// Unique identifier for the object.
	ID             string                       `json:"id"`
	Ideal          *PaymentMethodIdeal          `json:"ideal"`
	InteracPresent *PaymentMethodInteracPresent `json:"interac_present"`
	Klarna         *PaymentMethodKlarna         `json:"klarna"`
	Konbini        *PaymentMethodKonbini        `json:"konbini"`
	Link           *PaymentMethodLink           `json:"link"`
	// Has the value `true` if the object exists in live mode or the value `false` if the object exists in test mode.
	Livemode bool `json:"livemode"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format.
	Metadata map[string]string `json:"metadata"`
	// String representing the object's type. Objects of the same type share the same value.
	Object    string                  `json:"object"`
	OXXO      *PaymentMethodOXXO      `json:"oxxo"`
	P24       *PaymentMethodP24       `json:"p24"`
	PayNow    *PaymentMethodPayNow    `json:"paynow"`
	PromptPay *PaymentMethodPromptPay `json:"promptpay"`
	// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
	RadarOptions *PaymentMethodRadarOptions `json:"radar_options"`
	SepaDebit    *PaymentMethodSepaDebit    `json:"sepa_debit"`
	Sofort       *PaymentMethodSofort       `json:"sofort"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type          PaymentMethodType           `json:"type"`
	USBankAccount *PaymentMethodUSBankAccount `json:"us_bank_account"`
	WechatPay     *PaymentMethodWechatPay     `json:"wechat_pay"`
}

// PaymentMethodList is a list of PaymentMethods as retrieved from a list endpoint.
type PaymentMethodList struct {
	APIResource
	ListMeta
	Data []*PaymentMethod `json:"data"`
}

// UnmarshalJSON handles deserialization of a PaymentMethod.
// This custom unmarshaling is needed because the resulting
// property may be an id or the full struct if it was expanded.
func (p *PaymentMethod) UnmarshalJSON(data []byte) error {
	if id, ok := ParseID(data); ok {
		p.ID = id
		return nil
	}

	type paymentMethod PaymentMethod
	var v paymentMethod
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}

	*p = PaymentMethod(v)
	return nil
}
