//
//
// File generated from our OpenAPI spec
//
//

package stripe

// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
type TestHelpersConfirmationTokenPaymentMethodDataACSSDebitParams struct {
	// Customer's bank account number.
	AccountNumber *string `form:"account_number"`
	// Institution number of the customer's bank.
	InstitutionNumber *string `form:"institution_number"`
	// Transit number of the customer's bank.
	TransitNumber *string `form:"transit_number"`
}

// If this is an `affirm` PaymentMethod, this hash contains details about the Affirm payment method.
type TestHelpersConfirmationTokenPaymentMethodDataAffirmParams struct{}

// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataAfterpayClearpayParams struct{}

// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataAlipayParams struct{}

// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
type TestHelpersConfirmationTokenPaymentMethodDataAUBECSDebitParams struct {
	// The account number for the bank account.
	AccountNumber *string `form:"account_number"`
	// Bank-State-Branch number of the bank account.
	BSBNumber *string `form:"bsb_number"`
}

// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
type TestHelpersConfirmationTokenPaymentMethodDataBACSDebitParams struct {
	// Account number of the bank account that the funds will be debited from.
	AccountNumber *string `form:"account_number"`
	// Sort code of the bank account. (e.g., `10-20-30`)
	SortCode *string `form:"sort_code"`
}

// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
type TestHelpersConfirmationTokenPaymentMethodDataBancontactParams struct{}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParams struct {
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
type TestHelpersConfirmationTokenPaymentMethodDataBLIKParams struct{}

// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
type TestHelpersConfirmationTokenPaymentMethodDataBoletoParams struct {
	// The tax ID of the customer (CPF for individual consumers or CNPJ for businesses consumers)
	TaxID *string `form:"tax_id"`
}

// If this is a `cashapp` PaymentMethod, this hash contains details about the Cash App Pay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataCashAppParams struct{}

// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
type TestHelpersConfirmationTokenPaymentMethodDataCustomerBalanceParams struct{}

// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
type TestHelpersConfirmationTokenPaymentMethodDataEPSParams struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
type TestHelpersConfirmationTokenPaymentMethodDataFPXParams struct {
	// Account holder type for FPX transaction
	AccountHolderType *string `form:"account_holder_type"`
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataGiropayParams struct{}

// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataGrabpayParams struct{}

// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
type TestHelpersConfirmationTokenPaymentMethodDataIDEALParams struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
type TestHelpersConfirmationTokenPaymentMethodDataInteracPresentParams struct{}

// Customer's date of birth
type TestHelpersConfirmationTokenPaymentMethodDataKlarnaDOBParams struct {
	// The day of birth, between 1 and 31.
	Day *int64 `form:"day"`
	// The month of birth, between 1 and 12.
	Month *int64 `form:"month"`
	// The four-digit year of birth.
	Year *int64 `form:"year"`
}

// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
type TestHelpersConfirmationTokenPaymentMethodDataKlarnaParams struct {
	// Customer's date of birth
	DOB *TestHelpersConfirmationTokenPaymentMethodDataKlarnaDOBParams `form:"dob"`
}

// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
type TestHelpersConfirmationTokenPaymentMethodDataKonbiniParams struct{}

// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
type TestHelpersConfirmationTokenPaymentMethodDataLinkParams struct{}

// If this is a `mobilepay` PaymentMethod, this hash contains details about the MobilePay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataMobilepayParams struct{}

// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
type TestHelpersConfirmationTokenPaymentMethodDataOXXOParams struct{}

// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
type TestHelpersConfirmationTokenPaymentMethodDataP24Params struct {
	// The customer's bank.
	Bank *string `form:"bank"`
}

// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
type TestHelpersConfirmationTokenPaymentMethodDataPayNowParams struct{}

// If this is a `paypal` PaymentMethod, this hash contains details about the PayPal payment method.
type TestHelpersConfirmationTokenPaymentMethodDataPaypalParams struct{}

// If this is a `pix` PaymentMethod, this hash contains details about the Pix payment method.
type TestHelpersConfirmationTokenPaymentMethodDataPixParams struct{}

// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataPromptPayParams struct{}

// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
type TestHelpersConfirmationTokenPaymentMethodDataRadarOptionsParams struct {
	// A [Radar Session](https://stripe.com/docs/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session *string `form:"session"`
}

// If this is a `Revolut Pay` PaymentMethod, this hash contains details about the Revolut Pay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataRevolutPayParams struct{}

// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
type TestHelpersConfirmationTokenPaymentMethodDataSEPADebitParams struct {
	// IBAN of the bank account.
	IBAN *string `form:"iban"`
}

// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
type TestHelpersConfirmationTokenPaymentMethodDataSofortParams struct {
	// Two-letter ISO code representing the country the bank account is located in.
	Country *string `form:"country"`
}

// If this is a `swish` PaymentMethod, this hash contains details about the Swish payment method.
type TestHelpersConfirmationTokenPaymentMethodDataSwishParams struct{}

// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
type TestHelpersConfirmationTokenPaymentMethodDataUSBankAccountParams struct {
	// Account holder type: individual or company.
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
type TestHelpersConfirmationTokenPaymentMethodDataWeChatPayParams struct{}

// If this is a `zip` PaymentMethod, this hash contains details about the Zip payment method.
type TestHelpersConfirmationTokenPaymentMethodDataZipParams struct{}

// If provided, this hash will be used to create a PaymentMethod.
type TestHelpersConfirmationTokenPaymentMethodDataParams struct {
	// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
	ACSSDebit *TestHelpersConfirmationTokenPaymentMethodDataACSSDebitParams `form:"acss_debit"`
	// If this is an `affirm` PaymentMethod, this hash contains details about the Affirm payment method.
	Affirm *TestHelpersConfirmationTokenPaymentMethodDataAffirmParams `form:"affirm"`
	// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
	AfterpayClearpay *TestHelpersConfirmationTokenPaymentMethodDataAfterpayClearpayParams `form:"afterpay_clearpay"`
	// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
	Alipay *TestHelpersConfirmationTokenPaymentMethodDataAlipayParams `form:"alipay"`
	// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
	AUBECSDebit *TestHelpersConfirmationTokenPaymentMethodDataAUBECSDebitParams `form:"au_becs_debit"`
	// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
	BACSDebit *TestHelpersConfirmationTokenPaymentMethodDataBACSDebitParams `form:"bacs_debit"`
	// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
	Bancontact *TestHelpersConfirmationTokenPaymentMethodDataBancontactParams `form:"bancontact"`
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails *TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParams `form:"billing_details"`
	// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
	BLIK *TestHelpersConfirmationTokenPaymentMethodDataBLIKParams `form:"blik"`
	// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
	Boleto *TestHelpersConfirmationTokenPaymentMethodDataBoletoParams `form:"boleto"`
	// If this is a `cashapp` PaymentMethod, this hash contains details about the Cash App Pay payment method.
	CashApp *TestHelpersConfirmationTokenPaymentMethodDataCashAppParams `form:"cashapp"`
	// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
	CustomerBalance *TestHelpersConfirmationTokenPaymentMethodDataCustomerBalanceParams `form:"customer_balance"`
	// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
	EPS *TestHelpersConfirmationTokenPaymentMethodDataEPSParams `form:"eps"`
	// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
	FPX *TestHelpersConfirmationTokenPaymentMethodDataFPXParams `form:"fpx"`
	// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
	Giropay *TestHelpersConfirmationTokenPaymentMethodDataGiropayParams `form:"giropay"`
	// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
	Grabpay *TestHelpersConfirmationTokenPaymentMethodDataGrabpayParams `form:"grabpay"`
	// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
	IDEAL *TestHelpersConfirmationTokenPaymentMethodDataIDEALParams `form:"ideal"`
	// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
	InteracPresent *TestHelpersConfirmationTokenPaymentMethodDataInteracPresentParams `form:"interac_present"`
	// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
	Klarna *TestHelpersConfirmationTokenPaymentMethodDataKlarnaParams `form:"klarna"`
	// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
	Konbini *TestHelpersConfirmationTokenPaymentMethodDataKonbiniParams `form:"konbini"`
	// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
	Link *TestHelpersConfirmationTokenPaymentMethodDataLinkParams `form:"link"`
	// Set of [key-value pairs](https://stripe.com/docs/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata"`
	// If this is a `mobilepay` PaymentMethod, this hash contains details about the MobilePay payment method.
	Mobilepay *TestHelpersConfirmationTokenPaymentMethodDataMobilepayParams `form:"mobilepay"`
	// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
	OXXO *TestHelpersConfirmationTokenPaymentMethodDataOXXOParams `form:"oxxo"`
	// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
	P24 *TestHelpersConfirmationTokenPaymentMethodDataP24Params `form:"p24"`
	// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
	PayNow *TestHelpersConfirmationTokenPaymentMethodDataPayNowParams `form:"paynow"`
	// If this is a `paypal` PaymentMethod, this hash contains details about the PayPal payment method.
	Paypal *TestHelpersConfirmationTokenPaymentMethodDataPaypalParams `form:"paypal"`
	// If this is a `pix` PaymentMethod, this hash contains details about the Pix payment method.
	Pix *TestHelpersConfirmationTokenPaymentMethodDataPixParams `form:"pix"`
	// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
	PromptPay *TestHelpersConfirmationTokenPaymentMethodDataPromptPayParams `form:"promptpay"`
	// Options to configure Radar. See [Radar Session](https://stripe.com/docs/radar/radar-session) for more information.
	RadarOptions *TestHelpersConfirmationTokenPaymentMethodDataRadarOptionsParams `form:"radar_options"`
	// If this is a `Revolut Pay` PaymentMethod, this hash contains details about the Revolut Pay payment method.
	RevolutPay *TestHelpersConfirmationTokenPaymentMethodDataRevolutPayParams `form:"revolut_pay"`
	// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
	SEPADebit *TestHelpersConfirmationTokenPaymentMethodDataSEPADebitParams `form:"sepa_debit"`
	// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
	Sofort *TestHelpersConfirmationTokenPaymentMethodDataSofortParams `form:"sofort"`
	// If this is a `swish` PaymentMethod, this hash contains details about the Swish payment method.
	Swish *TestHelpersConfirmationTokenPaymentMethodDataSwishParams `form:"swish"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type"`
	// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
	USBankAccount *TestHelpersConfirmationTokenPaymentMethodDataUSBankAccountParams `form:"us_bank_account"`
	// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
	WeChatPay *TestHelpersConfirmationTokenPaymentMethodDataWeChatPayParams `form:"wechat_pay"`
	// If this is a `zip` PaymentMethod, this hash contains details about the Zip payment method.
	Zip *TestHelpersConfirmationTokenPaymentMethodDataZipParams `form:"zip"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TestHelpersConfirmationTokenPaymentMethodDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// Shipping information for this ConfirmationToken.
type TestHelpersConfirmationTokenShippingParams struct {
	// Shipping address
	Address *AddressParams `form:"address"`
	// Recipient name.
	Name *string `form:"name"`
	// Recipient phone (including extension)
	Phone *string `form:"phone"`
}

// Creates a test mode Confirmation Token server side for your integration tests.
type TestHelpersConfirmationTokenParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand"`
	// ID of an existing PaymentMethod.
	PaymentMethod *string `form:"payment_method"`
	// If provided, this hash will be used to create a PaymentMethod.
	PaymentMethodData *TestHelpersConfirmationTokenPaymentMethodDataParams `form:"payment_method_data"`
	// Return URL used to confirm the Intent.
	ReturnURL *string `form:"return_url"`
	// Indicates that you intend to make future payments with this ConfirmationToken's payment method.
	//
	// The presence of this property will [attach the payment method](https://stripe.com/docs/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete.
	SetupFutureUsage *string `form:"setup_future_usage"`
	// Shipping information for this ConfirmationToken.
	Shipping *TestHelpersConfirmationTokenShippingParams `form:"shipping"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersConfirmationTokenParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
