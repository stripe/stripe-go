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
	PaymentMethodTypeAfterpayClearpay PaymentMethodType = "afterpay_clearpay"
	PaymentMethodTypeAlipay           PaymentMethodType = "alipay"
	PaymentMethodTypeAUBECSDebit      PaymentMethodType = "au_becs_debit"
	PaymentMethodTypeBACSDebit        PaymentMethodType = "bacs_debit"
	PaymentMethodTypeBancontact       PaymentMethodType = "bancontact"
	PaymentMethodTypeBoleto           PaymentMethodType = "boleto"
	PaymentMethodTypeCard             PaymentMethodType = "card"
	PaymentMethodTypeCardPresent      PaymentMethodType = "card_present"
	PaymentMethodTypeEPS              PaymentMethodType = "eps"
	PaymentMethodTypeFPX              PaymentMethodType = "fpx"
	PaymentMethodTypeGiropay          PaymentMethodType = "giropay"
	PaymentMethodTypeGrabpay          PaymentMethodType = "grabpay"
	PaymentMethodTypeIdeal            PaymentMethodType = "ideal"
	PaymentMethodTypeInteracPresent   PaymentMethodType = "interac_present"
	PaymentMethodTypeKlarna           PaymentMethodType = "klarna"
	PaymentMethodTypeOXXO             PaymentMethodType = "oxxo"
	PaymentMethodTypeP24              PaymentMethodType = "p24"
	PaymentMethodTypeSepaDebit        PaymentMethodType = "sepa_debit"
	PaymentMethodTypeSofort           PaymentMethodType = "sofort"
	PaymentMethodTypeWechatPay        PaymentMethodType = "wechat_pay"
)

// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
type PaymentMethodACSSDebitParams struct {
	AccountNumber     *string `form:"account_number"`
	InstitutionNumber *string `form:"institution_number"`
	TransitNumber     *string `form:"transit_number"`
}

// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
type PaymentMethodAfterpayClearpayParams struct{}

// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
type PaymentMethodAlipayParams struct{}

// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
type PaymentMethodAUBECSDebitParams struct {
	AccountNumber *string `form:"account_number"`
	BSBNumber     *string `form:"bsb_number"`
}

// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
type PaymentMethodBACSDebitParams struct {
	AccountNumber *string `form:"account_number"`
	SortCode      *string `form:"sort_code"`
}

// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
type PaymentMethodBancontactParams struct{}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type BillingDetailsParams struct {
	Address *AddressParams `form:"address"`
	Email   *string        `form:"email"`
	Name    *string        `form:"name"`
	Phone   *string        `form:"phone"`
}

// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
type PaymentMethodBoletoParams struct {
	TaxID *string `form:"tax_id"`
}

// If this is a `card` PaymentMethod, this hash contains the user's card details. For backwards compatibility, you can alternatively provide a Stripe token (e.g., for Apple Pay, Amex Express Checkout, or legacy Checkout) into the card hash with format `card: {token: "tok_visa"}`. When providing a card number, you must meet the requirements for [PCI compliance](https://stripe.com/docs/security#validating-pci-compliance). We strongly recommend using Stripe.js instead of interacting with this API directly.
type PaymentMethodCardParams struct {
	CVC      *string `form:"cvc"`
	ExpMonth *string `form:"exp_month"`
	ExpYear  *string `form:"exp_year"`
	Number   *string `form:"number"`
	Token    *string `form:"token"`
}

// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
type PaymentMethodEPSParams struct {
	Bank *string `form:"bank"`
}

// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
type PaymentMethodFPXParams struct {
	AccountHolderType *string `form:"account_holder_type"`
	Bank              *string `form:"bank"`
}

// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
type PaymentMethodGiropayParams struct{}

// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
type PaymentMethodGrabpayParams struct{}

// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
type PaymentMethodIdealParams struct {
	Bank *string `form:"bank"`
}

// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
type PaymentMethodInteracPresentParams struct{}

// Customer's date of birth
type PaymentMethodKlarnaDOBParams struct {
	Day   *int64 `form:"day"`
	Month *int64 `form:"month"`
	Year  *int64 `form:"year"`
}

// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
type PaymentMethodKlarnaParams struct {
	DOB *PaymentMethodKlarnaDOBParams `form:"dob"`
}

// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
type PaymentMethodOXXOParams struct{}

// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
type PaymentMethodP24Params struct {
	Bank                *string `form:"bank"`
	TOSShownAndAccepted *bool   `form:"tos_shown_and_accepted"`
}

// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
type PaymentMethodSepaDebitParams struct {
	Iban *string `form:"iban"`
}

// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
type PaymentMethodSofortParams struct {
	Country *string `form:"country"`
}

// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
type PaymentMethodWechatPayParams struct{}

// Creates a PaymentMethod object. Read the [Stripe.js reference](https://stripe.com/docs/stripe-js/reference#stripe-create-payment-method) to learn how to create PaymentMethods via Stripe.js.
//
// Instead of creating a PaymentMethod directly, we recommend using the [PaymentIntents API to accept a payment immediately or the <a href="/docs/payments/save-and-reuse">SetupIntent](https://stripe.com/docs/payments/accept-a-payment) API to collect payment method details ahead of a future payment.
type PaymentMethodParams struct {
	Params           `form:"*"`
	ACSSDebit        *PaymentMethodACSSDebitParams        `form:"acss_debit"`
	AfterpayClearpay *PaymentMethodAfterpayClearpayParams `form:"afterpay_clearpay"`
	Alipay           *PaymentMethodAlipayParams           `form:"alipay"`
	AUBECSDebit      *PaymentMethodAUBECSDebitParams      `form:"au_becs_debit"`
	BACSDebit        *PaymentMethodBACSDebitParams        `form:"bacs_debit"`
	Bancontact       *PaymentMethodBancontactParams       `form:"bancontact"`
	BillingDetails   *BillingDetailsParams                `form:"billing_details"`
	Boleto           *PaymentMethodBoletoParams           `form:"boleto"`
	Card             *PaymentMethodCardParams             `form:"card"`
	EPS              *PaymentMethodEPSParams              `form:"eps"`
	FPX              *PaymentMethodFPXParams              `form:"fpx"`
	Giropay          *PaymentMethodGiropayParams          `form:"giropay"`
	Grabpay          *PaymentMethodGrabpayParams          `form:"grabpay"`
	Ideal            *PaymentMethodIdealParams            `form:"ideal"`
	InteracPresent   *PaymentMethodInteracPresentParams   `form:"interac_present"`
	Klarna           *PaymentMethodKlarnaParams           `form:"klarna"`
	OXXO             *PaymentMethodOXXOParams             `form:"oxxo"`
	P24              *PaymentMethodP24Params              `form:"p24"`
	SepaDebit        *PaymentMethodSepaDebitParams        `form:"sepa_debit"`
	Sofort           *PaymentMethodSofortParams           `form:"sofort"`
	Type             *string                              `form:"type"`
	WechatPay        *PaymentMethodWechatPayParams        `form:"wechat_pay"`
	// The following parameters are used when cloning a PaymentMethod to the connected account
	Customer      *string `form:"customer"`
	PaymentMethod *string `form:"payment_method"`
}

// Returns a list of PaymentMethods. For listing a customer's payment methods, you should use [List a Customer's PaymentMethods](https://stripe.com/docs/api/payment_methods/customer_list)
type PaymentMethodListParams struct {
	ListParams `form:"*"`
	Customer   *string `form:"customer"`
	Type       *string `form:"type"`
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
	Params   `form:"*"`
	Customer *string `form:"customer"`
}

// Detaches a PaymentMethod object from a Customer.
type PaymentMethodDetachParams struct {
	Params `form:"*"`
}
type PaymentMethodACSSDebit struct {
	BankName          string `json:"bank_name"`
	Fingerprint       string `json:"fingerprint"`
	InstitutionNumber string `json:"institution_number"`
	Last4             string `json:"last4"`
	TransitNumber     string `json:"transit_number"`
}
type PaymentMethodAfterpayClearpay struct{}
type PaymentMethodAlipay struct{}
type PaymentMethodAUBECSDebit struct {
	BSBNumber   string `json:"bsb_number"`
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
}
type PaymentMethodBACSDebit struct {
	Fingerprint string `json:"fingerprint"`
	Last4       string `json:"last4"`
	SortCode    string `json:"sort_code"`
}
type PaymentMethodBancontact struct{}
type BillingDetails struct {
	Address *Address `json:"address"`
	Email   string   `json:"email"`
	Name    string   `json:"name"`
	Phone   string   `json:"phone"`
}
type PaymentMethodBoleto struct {
	TaxID string `json:"tax_id"`
}

// Checks on Card address and CVC if provided.
type PaymentMethodCardChecks struct {
	AddressLine1Check      CardVerification `json:"address_line1_check"`
	AddressPostalCodeCheck CardVerification `json:"address_postal_code_check"`
	CVCCheck               CardVerification `json:"cvc_check"`
}

// Contains information about card networks that can be used to process the payment.
type PaymentMethodCardNetworks struct {
	Available []PaymentMethodCardNetwork `json:"available"`
	Preferred PaymentMethodCardNetwork   `json:"preferred"`
}

// Contains details on how this Card maybe be used for 3D Secure authentication.
type PaymentMethodCardThreeDSecureUsage struct {
	Supported bool `json:"supported"`
}
type PaymentMethodCardWalletAmexExpressCheckout struct{}
type PaymentMethodCardWalletApplePay struct{}
type PaymentMethodCardWalletGooglePay struct{}
type PaymentMethodCardWalletMasterpass struct {
	BillingAddress  *Address `json:"billing_address"`
	Email           string   `json:"email"`
	Name            string   `json:"name"`
	ShippingAddress *Address `json:"shipping_address"`
}
type PaymentMethodCardWalletSamsungPay struct{}
type PaymentMethodCardWalletVisaCheckout struct {
	BillingAddress  *Address `json:"billing_address"`
	Email           string   `json:"email"`
	Name            string   `json:"name"`
	ShippingAddress *Address `json:"shipping_address"`
}

// If this Card is part of a card wallet, this contains the details of the card wallet.
type PaymentMethodCardWallet struct {
	AmexExpressCheckout *PaymentMethodCardWalletAmexExpressCheckout `json:"amex_express_checkout"`
	ApplePay            *PaymentMethodCardWalletApplePay            `json:"apple_pay"`
	DynamicLast4        string                                      `json:"dynamic_last4"`
	GooglePay           *PaymentMethodCardWalletGooglePay           `json:"google_pay"`
	Masterpass          *PaymentMethodCardWalletMasterpass          `json:"masterpass"`
	SamsungPay          *PaymentMethodCardWalletSamsungPay          `json:"samsung_pay"`
	Type                PaymentMethodCardWalletType                 `json:"type"`
	VisaCheckout        *PaymentMethodCardWalletVisaCheckout        `json:"visa_checkout"`
}
type PaymentMethodCard struct {
	Brand             PaymentMethodCardBrand              `json:"brand"`
	Checks            *PaymentMethodCardChecks            `json:"checks"`
	Country           string                              `json:"country"`
	ExpMonth          uint64                              `json:"exp_month"`
	ExpYear           uint64                              `json:"exp_year"`
	Fingerprint       string                              `json:"fingerprint"`
	Funding           CardFunding                         `json:"funding"`
	Last4             string                              `json:"last4"`
	Networks          *PaymentMethodCardNetworks          `json:"networks"`
	ThreeDSecureUsage *PaymentMethodCardThreeDSecureUsage `json:"three_d_secure_usage"`
	Wallet            *PaymentMethodCardWallet            `json:"wallet"`
	// Please note that the fields below are for internal use only and are not returned
	// as part of standard API requests.
	Description string `json:"description"`
	IIN         string `json:"iin"`
	Issuer      string `json:"issuer"`
}
type PaymentMethodCardPresent struct{}
type PaymentMethodEPS struct {
	Bank string `json:"bank"`
}
type PaymentMethodFPX struct {
	AccountHolderType PaymentMethodFPXAccountHolderType `json:"account_holder_type"`
	Bank              string                            `json:"bank"`
	TransactionID     string                            `json:"transaction_id"`
}
type PaymentMethodGiropay struct{}
type PaymentMethodGrabpay struct{}
type PaymentMethodIdeal struct {
	Bank string `json:"bank"`
	Bic  string `json:"bic"`
}
type PaymentMethodInteracPresent struct{}

// The customer's date of birth, if provided.
type PaymentMethodKlarnaDOB struct {
	Day   int64 `json:"day"`
	Month int64 `json:"month"`
	Year  int64 `json:"year"`
}
type PaymentMethodKlarna struct {
	DOB *PaymentMethodKlarnaDOB `json:"dob"`
}
type PaymentMethodOXXO struct{}
type PaymentMethodP24 struct {
	Bank string `json:"bank"`
}

// Information about the object that generated this PaymentMethod.
type PaymentMethodSepaDebitGeneratedFrom struct {
	Charge       *Charge       `json:"charge"`
	SetupAttempt *SetupAttempt `json:"setup_attempt"`
}
type PaymentMethodSepaDebit struct {
	BankCode      string                               `json:"bank_code"`
	BranchCode    string                               `json:"branch_code"`
	Country       string                               `json:"country"`
	Fingerprint   string                               `json:"fingerprint"`
	GeneratedFrom *PaymentMethodSepaDebitGeneratedFrom `json:"generated_from"`
	Last4         string                               `json:"last4"`
}
type PaymentMethodSofort struct {
	Country string `json:"country"`
}
type PaymentMethodWechatPay struct{}

// PaymentMethod objects represent your customer's payment instruments.
// They can be used with [PaymentIntents](https://stripe.com/docs/payments/payment-intents) to collect payments or saved to
// Customer objects to store instrument details for future payments.
//
// Related guides: [Payment Methods](https://stripe.com/docs/payments/payment-methods) and [More Payment Scenarios](https://stripe.com/docs/payments/more-payment-scenarios).
type PaymentMethod struct {
	APIResource
	ACSSDebit        *PaymentMethodACSSDebit        `json:"acss_debit"`
	AfterpayClearpay *PaymentMethodAfterpayClearpay `json:"afterpay_clearpay"`
	Alipay           *PaymentMethodAlipay           `json:"alipay"`
	AUBECSDebit      *PaymentMethodAUBECSDebit      `json:"au_becs_debit"`
	BACSDebit        *PaymentMethodBACSDebit        `json:"bacs_debit"`
	Bancontact       *PaymentMethodBancontact       `json:"bancontact"`
	BillingDetails   *BillingDetails                `json:"billing_details"`
	Boleto           *PaymentMethodBoleto           `json:"boleto"`
	Card             *PaymentMethodCard             `json:"card"`
	CardPresent      *PaymentMethodCardPresent      `json:"card_present"`
	Created          int64                          `json:"created"`
	Customer         *Customer                      `json:"customer"`
	EPS              *PaymentMethodEPS              `json:"eps"`
	FPX              *PaymentMethodFPX              `json:"fpx"`
	Giropay          *PaymentMethodGiropay          `json:"giropay"`
	Grabpay          *PaymentMethodGrabpay          `json:"grabpay"`
	ID               string                         `json:"id"`
	Ideal            *PaymentMethodIdeal            `json:"ideal"`
	InteracPresent   *PaymentMethodInteracPresent   `json:"interac_present"`
	Klarna           *PaymentMethodKlarna           `json:"klarna"`
	Livemode         bool                           `json:"livemode"`
	Metadata         map[string]string              `json:"metadata"`
	Object           string                         `json:"object"`
	OXXO             *PaymentMethodOXXO             `json:"oxxo"`
	P24              *PaymentMethodP24              `json:"p24"`
	SepaDebit        *PaymentMethodSepaDebit        `json:"sepa_debit"`
	Sofort           *PaymentMethodSofort           `json:"sofort"`
	Type             PaymentMethodType              `json:"type"`
	WechatPay        *PaymentMethodWechatPay        `json:"wechat_pay"`
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
