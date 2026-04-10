//
//
// File generated from our OpenAPI spec
//
//

package stripe

// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
type TestHelpersConfirmationTokenPaymentMethodDataACSSDebitParams struct {
	// Customer's bank account number.
	AccountNumber *string `form:"account_number" json:"account_number"`
	// Institution number of the customer's bank.
	InstitutionNumber *string `form:"institution_number" json:"institution_number"`
	// Transit number of the customer's bank.
	TransitNumber *string `form:"transit_number" json:"transit_number"`
}

// If this is an `affirm` PaymentMethod, this hash contains details about the Affirm payment method.
type TestHelpersConfirmationTokenPaymentMethodDataAffirmParams struct{}

// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataAfterpayClearpayParams struct{}

// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataAlipayParams struct{}

// If this is a Alma PaymentMethod, this hash contains details about the Alma payment method.
type TestHelpersConfirmationTokenPaymentMethodDataAlmaParams struct{}

// If this is a AmazonPay PaymentMethod, this hash contains details about the AmazonPay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataAmazonPayParams struct{}

// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
type TestHelpersConfirmationTokenPaymentMethodDataAUBECSDebitParams struct {
	// The account number for the bank account.
	AccountNumber *string `form:"account_number" json:"account_number"`
	// Bank-State-Branch number of the bank account.
	BSBNumber *string `form:"bsb_number" json:"bsb_number"`
}

// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
type TestHelpersConfirmationTokenPaymentMethodDataBACSDebitParams struct {
	// Account number of the bank account that the funds will be debited from.
	AccountNumber *string `form:"account_number" json:"account_number,omitempty"`
	// Sort code of the bank account. (e.g., `10-20-30`)
	SortCode *string `form:"sort_code" json:"sort_code,omitempty"`
}

// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
type TestHelpersConfirmationTokenPaymentMethodDataBancontactParams struct{}

// If this is a `billie` PaymentMethod, this hash contains details about the Billie payment method.
type TestHelpersConfirmationTokenPaymentMethodDataBillieParams struct{}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Email address.
	Email *string `form:"email" json:"email,omitempty"`
	// Full name.
	Name *string `form:"name" json:"name,omitempty"`
	// Billing phone number (including extension).
	Phone *string `form:"phone" json:"phone,omitempty"`
	// Taxpayer identification number. Used only for transactions between LATAM buyers and non-LATAM sellers.
	TaxID       *string                                                                       `form:"tax_id" json:"tax_id,omitempty"`
	UnsetFields []TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParamsUnsetField `form:"-" json:"-"`
}

// TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParamsUnsetField is the list of fields that can be cleared/unset on TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParams.
type TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParamsUnsetField string

const (
	TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParamsUnsetFieldAddress TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParamsUnsetField = "address"
	TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParamsUnsetFieldEmail   TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParamsUnsetField = "email"
	TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParamsUnsetFieldName    TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParamsUnsetField = "name"
	TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParamsUnsetFieldPhone   TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParamsUnsetField = "phone"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParams) AddUnsetField(field TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
type TestHelpersConfirmationTokenPaymentMethodDataBLIKParams struct{}

// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
type TestHelpersConfirmationTokenPaymentMethodDataBoletoParams struct {
	// The tax ID of the customer (CPF for individual consumers or CNPJ for businesses consumers)
	TaxID *string `form:"tax_id" json:"tax_id"`
}

// If this is a `cashapp` PaymentMethod, this hash contains details about the Cash App Pay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataCashAppParams struct{}

// If this is a Crypto PaymentMethod, this hash contains details about the Crypto payment method.
type TestHelpersConfirmationTokenPaymentMethodDataCryptoParams struct{}

// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
type TestHelpersConfirmationTokenPaymentMethodDataCustomerBalanceParams struct{}

// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
type TestHelpersConfirmationTokenPaymentMethodDataEPSParams struct {
	// The customer's bank.
	Bank *string `form:"bank" json:"bank,omitempty"`
}

// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
type TestHelpersConfirmationTokenPaymentMethodDataFPXParams struct {
	// Account holder type for FPX transaction
	AccountHolderType *string `form:"account_holder_type" json:"account_holder_type,omitempty"`
	// The customer's bank.
	Bank *string `form:"bank" json:"bank"`
}

// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataGiropayParams struct{}

// If this is a Gopay PaymentMethod, this hash contains details about the Gopay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataGopayParams struct{}

// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataGrabpayParams struct{}

// If this is an `IdBankTransfer` PaymentMethod, this hash contains details about the IdBankTransfer payment method.
type TestHelpersConfirmationTokenPaymentMethodDataIDBankTransferParams struct {
	// Bank where the account is held.
	Bank *string `form:"bank" json:"bank,omitempty"`
}

// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
type TestHelpersConfirmationTokenPaymentMethodDataIDEALParams struct {
	// The customer's bank. Only use this parameter for existing customers. Don't use it for new customers.
	Bank *string `form:"bank" json:"bank,omitempty"`
}

// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
type TestHelpersConfirmationTokenPaymentMethodDataInteracPresentParams struct{}

// If this is a `kakao_pay` PaymentMethod, this hash contains details about the Kakao Pay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataKakaoPayParams struct{}

// Customer's date of birth
type TestHelpersConfirmationTokenPaymentMethodDataKlarnaDOBParams struct {
	// The day of birth, between 1 and 31.
	Day *int64 `form:"day" json:"day"`
	// The month of birth, between 1 and 12.
	Month *int64 `form:"month" json:"month"`
	// The four-digit year of birth.
	Year *int64 `form:"year" json:"year"`
}

// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
type TestHelpersConfirmationTokenPaymentMethodDataKlarnaParams struct {
	// Customer's date of birth
	DOB *TestHelpersConfirmationTokenPaymentMethodDataKlarnaDOBParams `form:"dob" json:"dob,omitempty"`
}

// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
type TestHelpersConfirmationTokenPaymentMethodDataKonbiniParams struct{}

// If this is a `kr_card` PaymentMethod, this hash contains details about the Korean Card payment method.
type TestHelpersConfirmationTokenPaymentMethodDataKrCardParams struct{}

// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
type TestHelpersConfirmationTokenPaymentMethodDataLinkParams struct{}

// If this is a MB WAY PaymentMethod, this hash contains details about the MB WAY payment method.
type TestHelpersConfirmationTokenPaymentMethodDataMbWayParams struct{}

// If this is a `mobilepay` PaymentMethod, this hash contains details about the MobilePay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataMobilepayParams struct{}

// If this is a `multibanco` PaymentMethod, this hash contains details about the Multibanco payment method.
type TestHelpersConfirmationTokenPaymentMethodDataMultibancoParams struct{}

// If this is a `naver_pay` PaymentMethod, this hash contains details about the Naver Pay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataNaverPayParams struct {
	// Whether to use Naver Pay points or a card to fund this transaction. If not provided, this defaults to `card`.
	Funding *string `form:"funding" json:"funding,omitempty"`
}

// If this is an nz_bank_account PaymentMethod, this hash contains details about the nz_bank_account payment method.
type TestHelpersConfirmationTokenPaymentMethodDataNzBankAccountParams struct {
	// The name on the bank account. Only required if the account holder name is different from the name of the authorized signatory collected in the PaymentMethod's billing details.
	AccountHolderName *string `form:"account_holder_name" json:"account_holder_name,omitempty"`
	// The account number for the bank account.
	AccountNumber *string `form:"account_number" json:"account_number"`
	// The numeric code for the bank account's bank.
	BankCode *string `form:"bank_code" json:"bank_code"`
	// The numeric code for the bank account's bank branch.
	BranchCode *string `form:"branch_code" json:"branch_code"`
	Reference  *string `form:"reference" json:"reference,omitempty"`
	// The suffix of the bank account number.
	Suffix *string `form:"suffix" json:"suffix"`
}

// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
type TestHelpersConfirmationTokenPaymentMethodDataOXXOParams struct{}

// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
type TestHelpersConfirmationTokenPaymentMethodDataP24Params struct {
	// The customer's bank.
	Bank *string `form:"bank" json:"bank,omitempty"`
}

// If this is a `pay_by_bank` PaymentMethod, this hash contains details about the PayByBank payment method.
type TestHelpersConfirmationTokenPaymentMethodDataPayByBankParams struct{}

// If this is a `payco` PaymentMethod, this hash contains details about the PAYCO payment method.
type TestHelpersConfirmationTokenPaymentMethodDataPaycoParams struct{}

// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
type TestHelpersConfirmationTokenPaymentMethodDataPayNowParams struct{}

// If this is a `paypal` PaymentMethod, this hash contains details about the PayPal payment method.
type TestHelpersConfirmationTokenPaymentMethodDataPaypalParams struct{}

// If this is a `paypay` PaymentMethod, this hash contains details about the PayPay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataPaypayParams struct{}

// If this is a `payto` PaymentMethod, this hash contains details about the PayTo payment method.
type TestHelpersConfirmationTokenPaymentMethodDataPaytoParams struct {
	// The account number for the bank account.
	AccountNumber *string `form:"account_number" json:"account_number,omitempty"`
	// Bank-State-Branch number of the bank account.
	BSBNumber *string `form:"bsb_number" json:"bsb_number,omitempty"`
	// The PayID alias for the bank account.
	PayID *string `form:"pay_id" json:"pay_id,omitempty"`
}

// If this is a `pix` PaymentMethod, this hash contains details about the Pix payment method.
type TestHelpersConfirmationTokenPaymentMethodDataPixParams struct{}

// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataPromptPayParams struct{}

// If this is a `qris` PaymentMethod, this hash contains details about the QRIS payment method.
type TestHelpersConfirmationTokenPaymentMethodDataQrisParams struct{}

// Options to configure Radar. See [Radar Session](https://docs.stripe.com/radar/radar-session) for more information.
type TestHelpersConfirmationTokenPaymentMethodDataRadarOptionsParams struct {
	// A [Radar Session](https://docs.stripe.com/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session *string `form:"session" json:"session,omitempty"`
}

// Customer's date of birth
type TestHelpersConfirmationTokenPaymentMethodDataRechnungDOBParams struct {
	// The day of birth, between 1 and 31.
	Day *int64 `form:"day" json:"day"`
	// The month of birth, between 1 and 12.
	Month *int64 `form:"month" json:"month"`
	// The four-digit year of birth.
	Year *int64 `form:"year" json:"year"`
}

// If this is a `rechnung` PaymentMethod, this hash contains details about the Rechnung payment method.
type TestHelpersConfirmationTokenPaymentMethodDataRechnungParams struct {
	// Customer's date of birth
	DOB *TestHelpersConfirmationTokenPaymentMethodDataRechnungDOBParams `form:"dob" json:"dob"`
}

// If this is a `revolut_pay` PaymentMethod, this hash contains details about the Revolut Pay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataRevolutPayParams struct{}

// If this is a `samsung_pay` PaymentMethod, this hash contains details about the SamsungPay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataSamsungPayParams struct{}

// If this is a `satispay` PaymentMethod, this hash contains details about the Satispay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataSatispayParams struct{}

// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
type TestHelpersConfirmationTokenPaymentMethodDataSEPADebitParams struct {
	// IBAN of the bank account.
	IBAN *string `form:"iban" json:"iban"`
}

// If this is a Shopeepay PaymentMethod, this hash contains details about the Shopeepay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataShopeepayParams struct{}

// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
type TestHelpersConfirmationTokenPaymentMethodDataSofortParams struct {
	// Two-letter ISO code representing the country the bank account is located in.
	Country *string `form:"country" json:"country"`
}

// This hash contains details about the Stripe balance payment method.
type TestHelpersConfirmationTokenPaymentMethodDataStripeBalanceParams struct {
	// The connected account ID whose Stripe balance to use as the source of payment
	Account *string `form:"account" json:"account,omitempty"`
}

// If this is a `swish` PaymentMethod, this hash contains details about the Swish payment method.
type TestHelpersConfirmationTokenPaymentMethodDataSwishParams struct{}

// If this is a TWINT PaymentMethod, this hash contains details about the TWINT payment method.
type TestHelpersConfirmationTokenPaymentMethodDataTWINTParams struct{}

// Configuration options for setting up an eMandate
type TestHelpersConfirmationTokenPaymentMethodDataUpiMandateOptionsParams struct {
	// Amount to be charged for future payments.
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.
	AmountType *string `form:"amount_type" json:"amount_type,omitempty"`
	// A description of the mandate or subscription that is meant to be displayed to the customer.
	Description *string `form:"description" json:"description,omitempty"`
	// End date of the mandate or subscription.
	EndDate *int64 `form:"end_date" json:"end_date,omitempty"`
}

// If this is a `upi` PaymentMethod, this hash contains details about the UPI payment method.
type TestHelpersConfirmationTokenPaymentMethodDataUpiParams struct {
	// Configuration options for setting up an eMandate
	MandateOptions *TestHelpersConfirmationTokenPaymentMethodDataUpiMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
}

// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
type TestHelpersConfirmationTokenPaymentMethodDataUSBankAccountParams struct {
	// Account holder type: individual or company.
	AccountHolderType *string `form:"account_holder_type" json:"account_holder_type,omitempty"`
	// Account number of the bank account.
	AccountNumber *string `form:"account_number" json:"account_number,omitempty"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType *string `form:"account_type" json:"account_type,omitempty"`
	// The ID of a Financial Connections Account to use as a payment method.
	FinancialConnectionsAccount *string `form:"financial_connections_account" json:"financial_connections_account,omitempty"`
	// Routing number of the bank account.
	RoutingNumber *string `form:"routing_number" json:"routing_number,omitempty"`
}

// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
type TestHelpersConfirmationTokenPaymentMethodDataWeChatPayParams struct{}

// If this is a `zip` PaymentMethod, this hash contains details about the Zip payment method.
type TestHelpersConfirmationTokenPaymentMethodDataZipParams struct{}

// If provided, this hash will be used to create a PaymentMethod.
type TestHelpersConfirmationTokenPaymentMethodDataParams struct {
	// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
	ACSSDebit *TestHelpersConfirmationTokenPaymentMethodDataACSSDebitParams `form:"acss_debit" json:"acss_debit,omitempty"`
	// If this is an `affirm` PaymentMethod, this hash contains details about the Affirm payment method.
	Affirm *TestHelpersConfirmationTokenPaymentMethodDataAffirmParams `form:"affirm" json:"affirm,omitempty"`
	// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
	AfterpayClearpay *TestHelpersConfirmationTokenPaymentMethodDataAfterpayClearpayParams `form:"afterpay_clearpay" json:"afterpay_clearpay,omitempty"`
	// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
	Alipay *TestHelpersConfirmationTokenPaymentMethodDataAlipayParams `form:"alipay" json:"alipay,omitempty"`
	// This field indicates whether this payment method can be shown again to its customer in a checkout flow. Stripe products such as Checkout and Elements use this field to determine whether a payment method can be shown as a saved payment method in a checkout flow. The field defaults to `unspecified`.
	AllowRedisplay *string `form:"allow_redisplay" json:"allow_redisplay,omitempty"`
	// If this is a Alma PaymentMethod, this hash contains details about the Alma payment method.
	Alma *TestHelpersConfirmationTokenPaymentMethodDataAlmaParams `form:"alma" json:"alma,omitempty"`
	// If this is a AmazonPay PaymentMethod, this hash contains details about the AmazonPay payment method.
	AmazonPay *TestHelpersConfirmationTokenPaymentMethodDataAmazonPayParams `form:"amazon_pay" json:"amazon_pay,omitempty"`
	// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
	AUBECSDebit *TestHelpersConfirmationTokenPaymentMethodDataAUBECSDebitParams `form:"au_becs_debit" json:"au_becs_debit,omitempty"`
	// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
	BACSDebit *TestHelpersConfirmationTokenPaymentMethodDataBACSDebitParams `form:"bacs_debit" json:"bacs_debit,omitempty"`
	// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
	Bancontact *TestHelpersConfirmationTokenPaymentMethodDataBancontactParams `form:"bancontact" json:"bancontact,omitempty"`
	// If this is a `billie` PaymentMethod, this hash contains details about the Billie payment method.
	Billie *TestHelpersConfirmationTokenPaymentMethodDataBillieParams `form:"billie" json:"billie,omitempty"`
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails *TestHelpersConfirmationTokenPaymentMethodDataBillingDetailsParams `form:"billing_details" json:"billing_details,omitempty"`
	// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
	BLIK *TestHelpersConfirmationTokenPaymentMethodDataBLIKParams `form:"blik" json:"blik,omitempty"`
	// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
	Boleto *TestHelpersConfirmationTokenPaymentMethodDataBoletoParams `form:"boleto" json:"boleto,omitempty"`
	// If this is a `cashapp` PaymentMethod, this hash contains details about the Cash App Pay payment method.
	CashApp *TestHelpersConfirmationTokenPaymentMethodDataCashAppParams `form:"cashapp" json:"cashapp,omitempty"`
	// If this is a Crypto PaymentMethod, this hash contains details about the Crypto payment method.
	Crypto *TestHelpersConfirmationTokenPaymentMethodDataCryptoParams `form:"crypto" json:"crypto,omitempty"`
	// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
	CustomerBalance *TestHelpersConfirmationTokenPaymentMethodDataCustomerBalanceParams `form:"customer_balance" json:"customer_balance,omitempty"`
	// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
	EPS *TestHelpersConfirmationTokenPaymentMethodDataEPSParams `form:"eps" json:"eps,omitempty"`
	// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
	FPX *TestHelpersConfirmationTokenPaymentMethodDataFPXParams `form:"fpx" json:"fpx,omitempty"`
	// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
	Giropay *TestHelpersConfirmationTokenPaymentMethodDataGiropayParams `form:"giropay" json:"giropay,omitempty"`
	// If this is a Gopay PaymentMethod, this hash contains details about the Gopay payment method.
	Gopay *TestHelpersConfirmationTokenPaymentMethodDataGopayParams `form:"gopay" json:"gopay,omitempty"`
	// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
	Grabpay *TestHelpersConfirmationTokenPaymentMethodDataGrabpayParams `form:"grabpay" json:"grabpay,omitempty"`
	// If this is an `IdBankTransfer` PaymentMethod, this hash contains details about the IdBankTransfer payment method.
	IDBankTransfer *TestHelpersConfirmationTokenPaymentMethodDataIDBankTransferParams `form:"id_bank_transfer" json:"id_bank_transfer,omitempty"`
	// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
	IDEAL *TestHelpersConfirmationTokenPaymentMethodDataIDEALParams `form:"ideal" json:"ideal,omitempty"`
	// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
	InteracPresent *TestHelpersConfirmationTokenPaymentMethodDataInteracPresentParams `form:"interac_present" json:"interac_present,omitempty"`
	// If this is a `kakao_pay` PaymentMethod, this hash contains details about the Kakao Pay payment method.
	KakaoPay *TestHelpersConfirmationTokenPaymentMethodDataKakaoPayParams `form:"kakao_pay" json:"kakao_pay,omitempty"`
	// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
	Klarna *TestHelpersConfirmationTokenPaymentMethodDataKlarnaParams `form:"klarna" json:"klarna,omitempty"`
	// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
	Konbini *TestHelpersConfirmationTokenPaymentMethodDataKonbiniParams `form:"konbini" json:"konbini,omitempty"`
	// If this is a `kr_card` PaymentMethod, this hash contains details about the Korean Card payment method.
	KrCard *TestHelpersConfirmationTokenPaymentMethodDataKrCardParams `form:"kr_card" json:"kr_card,omitempty"`
	// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
	Link *TestHelpersConfirmationTokenPaymentMethodDataLinkParams `form:"link" json:"link,omitempty"`
	// If this is a MB WAY PaymentMethod, this hash contains details about the MB WAY payment method.
	MbWay *TestHelpersConfirmationTokenPaymentMethodDataMbWayParams `form:"mb_way" json:"mb_way,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// If this is a `mobilepay` PaymentMethod, this hash contains details about the MobilePay payment method.
	Mobilepay *TestHelpersConfirmationTokenPaymentMethodDataMobilepayParams `form:"mobilepay" json:"mobilepay,omitempty"`
	// If this is a `multibanco` PaymentMethod, this hash contains details about the Multibanco payment method.
	Multibanco *TestHelpersConfirmationTokenPaymentMethodDataMultibancoParams `form:"multibanco" json:"multibanco,omitempty"`
	// If this is a `naver_pay` PaymentMethod, this hash contains details about the Naver Pay payment method.
	NaverPay *TestHelpersConfirmationTokenPaymentMethodDataNaverPayParams `form:"naver_pay" json:"naver_pay,omitempty"`
	// If this is an nz_bank_account PaymentMethod, this hash contains details about the nz_bank_account payment method.
	NzBankAccount *TestHelpersConfirmationTokenPaymentMethodDataNzBankAccountParams `form:"nz_bank_account" json:"nz_bank_account,omitempty"`
	// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
	OXXO *TestHelpersConfirmationTokenPaymentMethodDataOXXOParams `form:"oxxo" json:"oxxo,omitempty"`
	// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
	P24 *TestHelpersConfirmationTokenPaymentMethodDataP24Params `form:"p24" json:"p24,omitempty"`
	// If this is a `pay_by_bank` PaymentMethod, this hash contains details about the PayByBank payment method.
	PayByBank *TestHelpersConfirmationTokenPaymentMethodDataPayByBankParams `form:"pay_by_bank" json:"pay_by_bank,omitempty"`
	// If this is a `payco` PaymentMethod, this hash contains details about the PAYCO payment method.
	Payco *TestHelpersConfirmationTokenPaymentMethodDataPaycoParams `form:"payco" json:"payco,omitempty"`
	// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
	PayNow *TestHelpersConfirmationTokenPaymentMethodDataPayNowParams `form:"paynow" json:"paynow,omitempty"`
	// If this is a `paypal` PaymentMethod, this hash contains details about the PayPal payment method.
	Paypal *TestHelpersConfirmationTokenPaymentMethodDataPaypalParams `form:"paypal" json:"paypal,omitempty"`
	// If this is a `paypay` PaymentMethod, this hash contains details about the PayPay payment method.
	Paypay *TestHelpersConfirmationTokenPaymentMethodDataPaypayParams `form:"paypay" json:"paypay,omitempty"`
	// If this is a `payto` PaymentMethod, this hash contains details about the PayTo payment method.
	Payto *TestHelpersConfirmationTokenPaymentMethodDataPaytoParams `form:"payto" json:"payto,omitempty"`
	// If this is a `pix` PaymentMethod, this hash contains details about the Pix payment method.
	Pix *TestHelpersConfirmationTokenPaymentMethodDataPixParams `form:"pix" json:"pix,omitempty"`
	// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
	PromptPay *TestHelpersConfirmationTokenPaymentMethodDataPromptPayParams `form:"promptpay" json:"promptpay,omitempty"`
	// If this is a `qris` PaymentMethod, this hash contains details about the QRIS payment method.
	Qris *TestHelpersConfirmationTokenPaymentMethodDataQrisParams `form:"qris" json:"qris,omitempty"`
	// Options to configure Radar. See [Radar Session](https://docs.stripe.com/radar/radar-session) for more information.
	RadarOptions *TestHelpersConfirmationTokenPaymentMethodDataRadarOptionsParams `form:"radar_options" json:"radar_options,omitempty"`
	// If this is a `rechnung` PaymentMethod, this hash contains details about the Rechnung payment method.
	Rechnung *TestHelpersConfirmationTokenPaymentMethodDataRechnungParams `form:"rechnung" json:"rechnung,omitempty"`
	// If this is a `revolut_pay` PaymentMethod, this hash contains details about the Revolut Pay payment method.
	RevolutPay *TestHelpersConfirmationTokenPaymentMethodDataRevolutPayParams `form:"revolut_pay" json:"revolut_pay,omitempty"`
	// If this is a `samsung_pay` PaymentMethod, this hash contains details about the SamsungPay payment method.
	SamsungPay *TestHelpersConfirmationTokenPaymentMethodDataSamsungPayParams `form:"samsung_pay" json:"samsung_pay,omitempty"`
	// If this is a `satispay` PaymentMethod, this hash contains details about the Satispay payment method.
	Satispay *TestHelpersConfirmationTokenPaymentMethodDataSatispayParams `form:"satispay" json:"satispay,omitempty"`
	// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
	SEPADebit *TestHelpersConfirmationTokenPaymentMethodDataSEPADebitParams `form:"sepa_debit" json:"sepa_debit,omitempty"`
	// ID of the SharedPaymentGrantedToken used to confirm this PaymentIntent.
	SharedPaymentGrantedToken *string `form:"shared_payment_granted_token" json:"shared_payment_granted_token,omitempty"`
	// If this is a Shopeepay PaymentMethod, this hash contains details about the Shopeepay payment method.
	Shopeepay *TestHelpersConfirmationTokenPaymentMethodDataShopeepayParams `form:"shopeepay" json:"shopeepay,omitempty"`
	// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
	Sofort *TestHelpersConfirmationTokenPaymentMethodDataSofortParams `form:"sofort" json:"sofort,omitempty"`
	// This hash contains details about the Stripe balance payment method.
	StripeBalance *TestHelpersConfirmationTokenPaymentMethodDataStripeBalanceParams `form:"stripe_balance" json:"stripe_balance,omitempty"`
	// If this is a `swish` PaymentMethod, this hash contains details about the Swish payment method.
	Swish *TestHelpersConfirmationTokenPaymentMethodDataSwishParams `form:"swish" json:"swish,omitempty"`
	// If this is a TWINT PaymentMethod, this hash contains details about the TWINT payment method.
	TWINT *TestHelpersConfirmationTokenPaymentMethodDataTWINTParams `form:"twint" json:"twint,omitempty"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type" json:"type"`
	// If this is a `upi` PaymentMethod, this hash contains details about the UPI payment method.
	Upi *TestHelpersConfirmationTokenPaymentMethodDataUpiParams `form:"upi" json:"upi,omitempty"`
	// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
	USBankAccount *TestHelpersConfirmationTokenPaymentMethodDataUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
	// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
	WeChatPay *TestHelpersConfirmationTokenPaymentMethodDataWeChatPayParams `form:"wechat_pay" json:"wechat_pay,omitempty"`
	// If this is a `zip` PaymentMethod, this hash contains details about the Zip payment method.
	Zip *TestHelpersConfirmationTokenPaymentMethodDataZipParams `form:"zip" json:"zip,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TestHelpersConfirmationTokenPaymentMethodDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The selected installment plan to use for this payment attempt.
// This parameter can only be provided during confirmation.
type TestHelpersConfirmationTokenPaymentMethodOptionsCardInstallmentsPlanParams struct {
	// For `fixed_count` installment plans, this is required. It represents the number of installment payments your customer will make to their credit card.
	Count *int64 `form:"count" json:"count,omitempty"`
	// For `fixed_count` installment plans, this is required. It represents the interval between installment payments your customer will make to their credit card.
	// One of `month`.
	Interval *string `form:"interval" json:"interval,omitempty"`
	// Type of installment plan, one of `fixed_count`, `bonus`, or `revolving`.
	Type *string `form:"type" json:"type"`
}

// Installment configuration for payments confirmed using this ConfirmationToken.
type TestHelpersConfirmationTokenPaymentMethodOptionsCardInstallmentsParams struct {
	// The selected installment plan to use for this payment attempt.
	// This parameter can only be provided during confirmation.
	Plan *TestHelpersConfirmationTokenPaymentMethodOptionsCardInstallmentsPlanParams `form:"plan" json:"plan"`
}

// Configuration for any card payments confirmed using this ConfirmationToken.
type TestHelpersConfirmationTokenPaymentMethodOptionsCardParams struct {
	// Installment configuration for payments confirmed using this ConfirmationToken.
	Installments *TestHelpersConfirmationTokenPaymentMethodOptionsCardInstallmentsParams `form:"installments" json:"installments,omitempty"`
}

// Payment-method-specific configuration for this ConfirmationToken.
type TestHelpersConfirmationTokenPaymentMethodOptionsParams struct {
	// Configuration for any card payments confirmed using this ConfirmationToken.
	Card *TestHelpersConfirmationTokenPaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
}

// Shipping information for this ConfirmationToken.
type TestHelpersConfirmationTokenShippingParams struct {
	// Shipping address
	Address *AddressParams `form:"address" json:"address"`
	// Recipient name.
	Name *string `form:"name" json:"name"`
	// Recipient phone (including extension)
	Phone       *string                                                `form:"phone" json:"phone,omitempty"`
	UnsetFields []TestHelpersConfirmationTokenShippingParamsUnsetField `form:"-" json:"-"`
}

// TestHelpersConfirmationTokenShippingParamsUnsetField is the list of fields that can be cleared/unset on TestHelpersConfirmationTokenShippingParams.
type TestHelpersConfirmationTokenShippingParamsUnsetField string

const (
	TestHelpersConfirmationTokenShippingParamsUnsetFieldPhone TestHelpersConfirmationTokenShippingParamsUnsetField = "phone"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *TestHelpersConfirmationTokenShippingParams) AddUnsetField(field TestHelpersConfirmationTokenShippingParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Creates a test mode Confirmation Token server side for your integration tests.
type TestHelpersConfirmationTokenParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// ID of an existing PaymentMethod.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// If provided, this hash will be used to create a PaymentMethod.
	PaymentMethodData *TestHelpersConfirmationTokenPaymentMethodDataParams `form:"payment_method_data" json:"payment_method_data,omitempty"`
	// Payment-method-specific configuration for this ConfirmationToken.
	PaymentMethodOptions *TestHelpersConfirmationTokenPaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
	// Return URL used to confirm the Intent.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
	// Indicates that you intend to make future payments with this ConfirmationToken's payment method.
	//
	// The presence of this property will [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete.
	SetupFutureUsage *string `form:"setup_future_usage" json:"setup_future_usage,omitempty"`
	// Shipping information for this ConfirmationToken.
	Shipping *TestHelpersConfirmationTokenShippingParams `form:"shipping" json:"shipping,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersConfirmationTokenParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}

// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataACSSDebitParams struct {
	// Customer's bank account number.
	AccountNumber *string `form:"account_number" json:"account_number"`
	// Institution number of the customer's bank.
	InstitutionNumber *string `form:"institution_number" json:"institution_number"`
	// Transit number of the customer's bank.
	TransitNumber *string `form:"transit_number" json:"transit_number"`
}

// If this is an `affirm` PaymentMethod, this hash contains details about the Affirm payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataAffirmParams struct{}

// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataAfterpayClearpayParams struct{}

// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataAlipayParams struct{}

// If this is a Alma PaymentMethod, this hash contains details about the Alma payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataAlmaParams struct{}

// If this is a AmazonPay PaymentMethod, this hash contains details about the AmazonPay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataAmazonPayParams struct{}

// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
type TestHelpersConfirmationTokenCreatePaymentMethodDataAUBECSDebitParams struct {
	// The account number for the bank account.
	AccountNumber *string `form:"account_number" json:"account_number"`
	// Bank-State-Branch number of the bank account.
	BSBNumber *string `form:"bsb_number" json:"bsb_number"`
}

// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
type TestHelpersConfirmationTokenCreatePaymentMethodDataBACSDebitParams struct {
	// Account number of the bank account that the funds will be debited from.
	AccountNumber *string `form:"account_number" json:"account_number,omitempty"`
	// Sort code of the bank account. (e.g., `10-20-30`)
	SortCode *string `form:"sort_code" json:"sort_code,omitempty"`
}

// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataBancontactParams struct{}

// If this is a `billie` PaymentMethod, this hash contains details about the Billie payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataBillieParams struct{}

// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
type TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParams struct {
	// Billing address.
	Address *AddressParams `form:"address" json:"address,omitempty"`
	// Email address.
	Email *string `form:"email" json:"email,omitempty"`
	// Full name.
	Name *string `form:"name" json:"name,omitempty"`
	// Billing phone number (including extension).
	Phone *string `form:"phone" json:"phone,omitempty"`
	// Taxpayer identification number. Used only for transactions between LATAM buyers and non-LATAM sellers.
	TaxID       *string                                                                             `form:"tax_id" json:"tax_id,omitempty"`
	UnsetFields []TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParamsUnsetField `form:"-" json:"-"`
}

// TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParamsUnsetField is the list of fields that can be cleared/unset on TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParams.
type TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParamsUnsetField string

const (
	TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParamsUnsetFieldAddress TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParamsUnsetField = "address"
	TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParamsUnsetFieldEmail   TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParamsUnsetField = "email"
	TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParamsUnsetFieldName    TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParamsUnsetField = "name"
	TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParamsUnsetFieldPhone   TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParamsUnsetField = "phone"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParams) AddUnsetField(field TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataBLIKParams struct{}

// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataBoletoParams struct {
	// The tax ID of the customer (CPF for individual consumers or CNPJ for businesses consumers)
	TaxID *string `form:"tax_id" json:"tax_id"`
}

// If this is a `cashapp` PaymentMethod, this hash contains details about the Cash App Pay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataCashAppParams struct{}

// If this is a Crypto PaymentMethod, this hash contains details about the Crypto payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataCryptoParams struct{}

// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataCustomerBalanceParams struct{}

// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataEPSParams struct {
	// The customer's bank.
	Bank *string `form:"bank" json:"bank,omitempty"`
}

// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataFPXParams struct {
	// Account holder type for FPX transaction
	AccountHolderType *string `form:"account_holder_type" json:"account_holder_type,omitempty"`
	// The customer's bank.
	Bank *string `form:"bank" json:"bank"`
}

// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataGiropayParams struct{}

// If this is a Gopay PaymentMethod, this hash contains details about the Gopay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataGopayParams struct{}

// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataGrabpayParams struct{}

// If this is an `IdBankTransfer` PaymentMethod, this hash contains details about the IdBankTransfer payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataIDBankTransferParams struct {
	// Bank where the account is held.
	Bank *string `form:"bank" json:"bank,omitempty"`
}

// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataIDEALParams struct {
	// The customer's bank. Only use this parameter for existing customers. Don't use it for new customers.
	Bank *string `form:"bank" json:"bank,omitempty"`
}

// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataInteracPresentParams struct{}

// If this is a `kakao_pay` PaymentMethod, this hash contains details about the Kakao Pay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataKakaoPayParams struct{}

// Customer's date of birth
type TestHelpersConfirmationTokenCreatePaymentMethodDataKlarnaDOBParams struct {
	// The day of birth, between 1 and 31.
	Day *int64 `form:"day" json:"day"`
	// The month of birth, between 1 and 12.
	Month *int64 `form:"month" json:"month"`
	// The four-digit year of birth.
	Year *int64 `form:"year" json:"year"`
}

// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataKlarnaParams struct {
	// Customer's date of birth
	DOB *TestHelpersConfirmationTokenCreatePaymentMethodDataKlarnaDOBParams `form:"dob" json:"dob,omitempty"`
}

// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataKonbiniParams struct{}

// If this is a `kr_card` PaymentMethod, this hash contains details about the Korean Card payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataKrCardParams struct{}

// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataLinkParams struct{}

// If this is a MB WAY PaymentMethod, this hash contains details about the MB WAY payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataMbWayParams struct{}

// If this is a `mobilepay` PaymentMethod, this hash contains details about the MobilePay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataMobilepayParams struct{}

// If this is a `multibanco` PaymentMethod, this hash contains details about the Multibanco payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataMultibancoParams struct{}

// If this is a `naver_pay` PaymentMethod, this hash contains details about the Naver Pay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataNaverPayParams struct {
	// Whether to use Naver Pay points or a card to fund this transaction. If not provided, this defaults to `card`.
	Funding *string `form:"funding" json:"funding,omitempty"`
}

// If this is an nz_bank_account PaymentMethod, this hash contains details about the nz_bank_account payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataNzBankAccountParams struct {
	// The name on the bank account. Only required if the account holder name is different from the name of the authorized signatory collected in the PaymentMethod's billing details.
	AccountHolderName *string `form:"account_holder_name" json:"account_holder_name,omitempty"`
	// The account number for the bank account.
	AccountNumber *string `form:"account_number" json:"account_number"`
	// The numeric code for the bank account's bank.
	BankCode *string `form:"bank_code" json:"bank_code"`
	// The numeric code for the bank account's bank branch.
	BranchCode *string `form:"branch_code" json:"branch_code"`
	Reference  *string `form:"reference" json:"reference,omitempty"`
	// The suffix of the bank account number.
	Suffix *string `form:"suffix" json:"suffix"`
}

// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataOXXOParams struct{}

// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataP24Params struct {
	// The customer's bank.
	Bank *string `form:"bank" json:"bank,omitempty"`
}

// If this is a `pay_by_bank` PaymentMethod, this hash contains details about the PayByBank payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataPayByBankParams struct{}

// If this is a `payco` PaymentMethod, this hash contains details about the PAYCO payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataPaycoParams struct{}

// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataPayNowParams struct{}

// If this is a `paypal` PaymentMethod, this hash contains details about the PayPal payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataPaypalParams struct{}

// If this is a `paypay` PaymentMethod, this hash contains details about the PayPay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataPaypayParams struct{}

// If this is a `payto` PaymentMethod, this hash contains details about the PayTo payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataPaytoParams struct {
	// The account number for the bank account.
	AccountNumber *string `form:"account_number" json:"account_number,omitempty"`
	// Bank-State-Branch number of the bank account.
	BSBNumber *string `form:"bsb_number" json:"bsb_number,omitempty"`
	// The PayID alias for the bank account.
	PayID *string `form:"pay_id" json:"pay_id,omitempty"`
}

// If this is a `pix` PaymentMethod, this hash contains details about the Pix payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataPixParams struct{}

// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataPromptPayParams struct{}

// If this is a `qris` PaymentMethod, this hash contains details about the QRIS payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataQrisParams struct{}

// Options to configure Radar. See [Radar Session](https://docs.stripe.com/radar/radar-session) for more information.
type TestHelpersConfirmationTokenCreatePaymentMethodDataRadarOptionsParams struct {
	// A [Radar Session](https://docs.stripe.com/radar/radar-session) is a snapshot of the browser metadata and device details that help Radar make more accurate predictions on your payments.
	Session *string `form:"session" json:"session,omitempty"`
}

// Customer's date of birth
type TestHelpersConfirmationTokenCreatePaymentMethodDataRechnungDOBParams struct {
	// The day of birth, between 1 and 31.
	Day *int64 `form:"day" json:"day"`
	// The month of birth, between 1 and 12.
	Month *int64 `form:"month" json:"month"`
	// The four-digit year of birth.
	Year *int64 `form:"year" json:"year"`
}

// If this is a `rechnung` PaymentMethod, this hash contains details about the Rechnung payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataRechnungParams struct {
	// Customer's date of birth
	DOB *TestHelpersConfirmationTokenCreatePaymentMethodDataRechnungDOBParams `form:"dob" json:"dob"`
}

// If this is a `revolut_pay` PaymentMethod, this hash contains details about the Revolut Pay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataRevolutPayParams struct{}

// If this is a `samsung_pay` PaymentMethod, this hash contains details about the SamsungPay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataSamsungPayParams struct{}

// If this is a `satispay` PaymentMethod, this hash contains details about the Satispay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataSatispayParams struct{}

// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
type TestHelpersConfirmationTokenCreatePaymentMethodDataSEPADebitParams struct {
	// IBAN of the bank account.
	IBAN *string `form:"iban" json:"iban"`
}

// If this is a Shopeepay PaymentMethod, this hash contains details about the Shopeepay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataShopeepayParams struct{}

// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataSofortParams struct {
	// Two-letter ISO code representing the country the bank account is located in.
	Country *string `form:"country" json:"country"`
}

// This hash contains details about the Stripe balance payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataStripeBalanceParams struct {
	// The connected account ID whose Stripe balance to use as the source of payment
	Account *string `form:"account" json:"account,omitempty"`
}

// If this is a `swish` PaymentMethod, this hash contains details about the Swish payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataSwishParams struct{}

// If this is a TWINT PaymentMethod, this hash contains details about the TWINT payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataTWINTParams struct{}

// Configuration options for setting up an eMandate
type TestHelpersConfirmationTokenCreatePaymentMethodDataUpiMandateOptionsParams struct {
	// Amount to be charged for future payments.
	Amount *int64 `form:"amount" json:"amount,omitempty"`
	// One of `fixed` or `maximum`. If `fixed`, the `amount` param refers to the exact amount to be charged in future payments. If `maximum`, the amount charged can be up to the value passed for the `amount` param.
	AmountType *string `form:"amount_type" json:"amount_type,omitempty"`
	// A description of the mandate or subscription that is meant to be displayed to the customer.
	Description *string `form:"description" json:"description,omitempty"`
	// End date of the mandate or subscription.
	EndDate *int64 `form:"end_date" json:"end_date,omitempty"`
}

// If this is a `upi` PaymentMethod, this hash contains details about the UPI payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataUpiParams struct {
	// Configuration options for setting up an eMandate
	MandateOptions *TestHelpersConfirmationTokenCreatePaymentMethodDataUpiMandateOptionsParams `form:"mandate_options" json:"mandate_options,omitempty"`
}

// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataUSBankAccountParams struct {
	// Account holder type: individual or company.
	AccountHolderType *string `form:"account_holder_type" json:"account_holder_type,omitempty"`
	// Account number of the bank account.
	AccountNumber *string `form:"account_number" json:"account_number,omitempty"`
	// Account type: checkings or savings. Defaults to checking if omitted.
	AccountType *string `form:"account_type" json:"account_type,omitempty"`
	// The ID of a Financial Connections Account to use as a payment method.
	FinancialConnectionsAccount *string `form:"financial_connections_account" json:"financial_connections_account,omitempty"`
	// Routing number of the bank account.
	RoutingNumber *string `form:"routing_number" json:"routing_number,omitempty"`
}

// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataWeChatPayParams struct{}

// If this is a `zip` PaymentMethod, this hash contains details about the Zip payment method.
type TestHelpersConfirmationTokenCreatePaymentMethodDataZipParams struct{}

// If provided, this hash will be used to create a PaymentMethod.
type TestHelpersConfirmationTokenCreatePaymentMethodDataParams struct {
	// If this is an `acss_debit` PaymentMethod, this hash contains details about the ACSS Debit payment method.
	ACSSDebit *TestHelpersConfirmationTokenCreatePaymentMethodDataACSSDebitParams `form:"acss_debit" json:"acss_debit,omitempty"`
	// If this is an `affirm` PaymentMethod, this hash contains details about the Affirm payment method.
	Affirm *TestHelpersConfirmationTokenCreatePaymentMethodDataAffirmParams `form:"affirm" json:"affirm,omitempty"`
	// If this is an `AfterpayClearpay` PaymentMethod, this hash contains details about the AfterpayClearpay payment method.
	AfterpayClearpay *TestHelpersConfirmationTokenCreatePaymentMethodDataAfterpayClearpayParams `form:"afterpay_clearpay" json:"afterpay_clearpay,omitempty"`
	// If this is an `Alipay` PaymentMethod, this hash contains details about the Alipay payment method.
	Alipay *TestHelpersConfirmationTokenCreatePaymentMethodDataAlipayParams `form:"alipay" json:"alipay,omitempty"`
	// This field indicates whether this payment method can be shown again to its customer in a checkout flow. Stripe products such as Checkout and Elements use this field to determine whether a payment method can be shown as a saved payment method in a checkout flow. The field defaults to `unspecified`.
	AllowRedisplay *string `form:"allow_redisplay" json:"allow_redisplay,omitempty"`
	// If this is a Alma PaymentMethod, this hash contains details about the Alma payment method.
	Alma *TestHelpersConfirmationTokenCreatePaymentMethodDataAlmaParams `form:"alma" json:"alma,omitempty"`
	// If this is a AmazonPay PaymentMethod, this hash contains details about the AmazonPay payment method.
	AmazonPay *TestHelpersConfirmationTokenCreatePaymentMethodDataAmazonPayParams `form:"amazon_pay" json:"amazon_pay,omitempty"`
	// If this is an `au_becs_debit` PaymentMethod, this hash contains details about the bank account.
	AUBECSDebit *TestHelpersConfirmationTokenCreatePaymentMethodDataAUBECSDebitParams `form:"au_becs_debit" json:"au_becs_debit,omitempty"`
	// If this is a `bacs_debit` PaymentMethod, this hash contains details about the Bacs Direct Debit bank account.
	BACSDebit *TestHelpersConfirmationTokenCreatePaymentMethodDataBACSDebitParams `form:"bacs_debit" json:"bacs_debit,omitempty"`
	// If this is a `bancontact` PaymentMethod, this hash contains details about the Bancontact payment method.
	Bancontact *TestHelpersConfirmationTokenCreatePaymentMethodDataBancontactParams `form:"bancontact" json:"bancontact,omitempty"`
	// If this is a `billie` PaymentMethod, this hash contains details about the Billie payment method.
	Billie *TestHelpersConfirmationTokenCreatePaymentMethodDataBillieParams `form:"billie" json:"billie,omitempty"`
	// Billing information associated with the PaymentMethod that may be used or required by particular types of payment methods.
	BillingDetails *TestHelpersConfirmationTokenCreatePaymentMethodDataBillingDetailsParams `form:"billing_details" json:"billing_details,omitempty"`
	// If this is a `blik` PaymentMethod, this hash contains details about the BLIK payment method.
	BLIK *TestHelpersConfirmationTokenCreatePaymentMethodDataBLIKParams `form:"blik" json:"blik,omitempty"`
	// If this is a `boleto` PaymentMethod, this hash contains details about the Boleto payment method.
	Boleto *TestHelpersConfirmationTokenCreatePaymentMethodDataBoletoParams `form:"boleto" json:"boleto,omitempty"`
	// If this is a `cashapp` PaymentMethod, this hash contains details about the Cash App Pay payment method.
	CashApp *TestHelpersConfirmationTokenCreatePaymentMethodDataCashAppParams `form:"cashapp" json:"cashapp,omitempty"`
	// If this is a Crypto PaymentMethod, this hash contains details about the Crypto payment method.
	Crypto *TestHelpersConfirmationTokenCreatePaymentMethodDataCryptoParams `form:"crypto" json:"crypto,omitempty"`
	// If this is a `customer_balance` PaymentMethod, this hash contains details about the CustomerBalance payment method.
	CustomerBalance *TestHelpersConfirmationTokenCreatePaymentMethodDataCustomerBalanceParams `form:"customer_balance" json:"customer_balance,omitempty"`
	// If this is an `eps` PaymentMethod, this hash contains details about the EPS payment method.
	EPS *TestHelpersConfirmationTokenCreatePaymentMethodDataEPSParams `form:"eps" json:"eps,omitempty"`
	// If this is an `fpx` PaymentMethod, this hash contains details about the FPX payment method.
	FPX *TestHelpersConfirmationTokenCreatePaymentMethodDataFPXParams `form:"fpx" json:"fpx,omitempty"`
	// If this is a `giropay` PaymentMethod, this hash contains details about the Giropay payment method.
	Giropay *TestHelpersConfirmationTokenCreatePaymentMethodDataGiropayParams `form:"giropay" json:"giropay,omitempty"`
	// If this is a Gopay PaymentMethod, this hash contains details about the Gopay payment method.
	Gopay *TestHelpersConfirmationTokenCreatePaymentMethodDataGopayParams `form:"gopay" json:"gopay,omitempty"`
	// If this is a `grabpay` PaymentMethod, this hash contains details about the GrabPay payment method.
	Grabpay *TestHelpersConfirmationTokenCreatePaymentMethodDataGrabpayParams `form:"grabpay" json:"grabpay,omitempty"`
	// If this is an `IdBankTransfer` PaymentMethod, this hash contains details about the IdBankTransfer payment method.
	IDBankTransfer *TestHelpersConfirmationTokenCreatePaymentMethodDataIDBankTransferParams `form:"id_bank_transfer" json:"id_bank_transfer,omitempty"`
	// If this is an `ideal` PaymentMethod, this hash contains details about the iDEAL payment method.
	IDEAL *TestHelpersConfirmationTokenCreatePaymentMethodDataIDEALParams `form:"ideal" json:"ideal,omitempty"`
	// If this is an `interac_present` PaymentMethod, this hash contains details about the Interac Present payment method.
	InteracPresent *TestHelpersConfirmationTokenCreatePaymentMethodDataInteracPresentParams `form:"interac_present" json:"interac_present,omitempty"`
	// If this is a `kakao_pay` PaymentMethod, this hash contains details about the Kakao Pay payment method.
	KakaoPay *TestHelpersConfirmationTokenCreatePaymentMethodDataKakaoPayParams `form:"kakao_pay" json:"kakao_pay,omitempty"`
	// If this is a `klarna` PaymentMethod, this hash contains details about the Klarna payment method.
	Klarna *TestHelpersConfirmationTokenCreatePaymentMethodDataKlarnaParams `form:"klarna" json:"klarna,omitempty"`
	// If this is a `konbini` PaymentMethod, this hash contains details about the Konbini payment method.
	Konbini *TestHelpersConfirmationTokenCreatePaymentMethodDataKonbiniParams `form:"konbini" json:"konbini,omitempty"`
	// If this is a `kr_card` PaymentMethod, this hash contains details about the Korean Card payment method.
	KrCard *TestHelpersConfirmationTokenCreatePaymentMethodDataKrCardParams `form:"kr_card" json:"kr_card,omitempty"`
	// If this is an `Link` PaymentMethod, this hash contains details about the Link payment method.
	Link *TestHelpersConfirmationTokenCreatePaymentMethodDataLinkParams `form:"link" json:"link,omitempty"`
	// If this is a MB WAY PaymentMethod, this hash contains details about the MB WAY payment method.
	MbWay *TestHelpersConfirmationTokenCreatePaymentMethodDataMbWayParams `form:"mb_way" json:"mb_way,omitempty"`
	// Set of [key-value pairs](https://docs.stripe.com/api/metadata) that you can attach to an object. This can be useful for storing additional information about the object in a structured format. Individual keys can be unset by posting an empty value to them. All keys can be unset by posting an empty value to `metadata`.
	Metadata map[string]string `form:"metadata" json:"metadata,omitempty"`
	// If this is a `mobilepay` PaymentMethod, this hash contains details about the MobilePay payment method.
	Mobilepay *TestHelpersConfirmationTokenCreatePaymentMethodDataMobilepayParams `form:"mobilepay" json:"mobilepay,omitempty"`
	// If this is a `multibanco` PaymentMethod, this hash contains details about the Multibanco payment method.
	Multibanco *TestHelpersConfirmationTokenCreatePaymentMethodDataMultibancoParams `form:"multibanco" json:"multibanco,omitempty"`
	// If this is a `naver_pay` PaymentMethod, this hash contains details about the Naver Pay payment method.
	NaverPay *TestHelpersConfirmationTokenCreatePaymentMethodDataNaverPayParams `form:"naver_pay" json:"naver_pay,omitempty"`
	// If this is an nz_bank_account PaymentMethod, this hash contains details about the nz_bank_account payment method.
	NzBankAccount *TestHelpersConfirmationTokenCreatePaymentMethodDataNzBankAccountParams `form:"nz_bank_account" json:"nz_bank_account,omitempty"`
	// If this is an `oxxo` PaymentMethod, this hash contains details about the OXXO payment method.
	OXXO *TestHelpersConfirmationTokenCreatePaymentMethodDataOXXOParams `form:"oxxo" json:"oxxo,omitempty"`
	// If this is a `p24` PaymentMethod, this hash contains details about the P24 payment method.
	P24 *TestHelpersConfirmationTokenCreatePaymentMethodDataP24Params `form:"p24" json:"p24,omitempty"`
	// If this is a `pay_by_bank` PaymentMethod, this hash contains details about the PayByBank payment method.
	PayByBank *TestHelpersConfirmationTokenCreatePaymentMethodDataPayByBankParams `form:"pay_by_bank" json:"pay_by_bank,omitempty"`
	// If this is a `payco` PaymentMethod, this hash contains details about the PAYCO payment method.
	Payco *TestHelpersConfirmationTokenCreatePaymentMethodDataPaycoParams `form:"payco" json:"payco,omitempty"`
	// If this is a `paynow` PaymentMethod, this hash contains details about the PayNow payment method.
	PayNow *TestHelpersConfirmationTokenCreatePaymentMethodDataPayNowParams `form:"paynow" json:"paynow,omitempty"`
	// If this is a `paypal` PaymentMethod, this hash contains details about the PayPal payment method.
	Paypal *TestHelpersConfirmationTokenCreatePaymentMethodDataPaypalParams `form:"paypal" json:"paypal,omitempty"`
	// If this is a `paypay` PaymentMethod, this hash contains details about the PayPay payment method.
	Paypay *TestHelpersConfirmationTokenCreatePaymentMethodDataPaypayParams `form:"paypay" json:"paypay,omitempty"`
	// If this is a `payto` PaymentMethod, this hash contains details about the PayTo payment method.
	Payto *TestHelpersConfirmationTokenCreatePaymentMethodDataPaytoParams `form:"payto" json:"payto,omitempty"`
	// If this is a `pix` PaymentMethod, this hash contains details about the Pix payment method.
	Pix *TestHelpersConfirmationTokenCreatePaymentMethodDataPixParams `form:"pix" json:"pix,omitempty"`
	// If this is a `promptpay` PaymentMethod, this hash contains details about the PromptPay payment method.
	PromptPay *TestHelpersConfirmationTokenCreatePaymentMethodDataPromptPayParams `form:"promptpay" json:"promptpay,omitempty"`
	// If this is a `qris` PaymentMethod, this hash contains details about the QRIS payment method.
	Qris *TestHelpersConfirmationTokenCreatePaymentMethodDataQrisParams `form:"qris" json:"qris,omitempty"`
	// Options to configure Radar. See [Radar Session](https://docs.stripe.com/radar/radar-session) for more information.
	RadarOptions *TestHelpersConfirmationTokenCreatePaymentMethodDataRadarOptionsParams `form:"radar_options" json:"radar_options,omitempty"`
	// If this is a `rechnung` PaymentMethod, this hash contains details about the Rechnung payment method.
	Rechnung *TestHelpersConfirmationTokenCreatePaymentMethodDataRechnungParams `form:"rechnung" json:"rechnung,omitempty"`
	// If this is a `revolut_pay` PaymentMethod, this hash contains details about the Revolut Pay payment method.
	RevolutPay *TestHelpersConfirmationTokenCreatePaymentMethodDataRevolutPayParams `form:"revolut_pay" json:"revolut_pay,omitempty"`
	// If this is a `samsung_pay` PaymentMethod, this hash contains details about the SamsungPay payment method.
	SamsungPay *TestHelpersConfirmationTokenCreatePaymentMethodDataSamsungPayParams `form:"samsung_pay" json:"samsung_pay,omitempty"`
	// If this is a `satispay` PaymentMethod, this hash contains details about the Satispay payment method.
	Satispay *TestHelpersConfirmationTokenCreatePaymentMethodDataSatispayParams `form:"satispay" json:"satispay,omitempty"`
	// If this is a `sepa_debit` PaymentMethod, this hash contains details about the SEPA debit bank account.
	SEPADebit *TestHelpersConfirmationTokenCreatePaymentMethodDataSEPADebitParams `form:"sepa_debit" json:"sepa_debit,omitempty"`
	// ID of the SharedPaymentGrantedToken used to confirm this PaymentIntent.
	SharedPaymentGrantedToken *string `form:"shared_payment_granted_token" json:"shared_payment_granted_token,omitempty"`
	// If this is a Shopeepay PaymentMethod, this hash contains details about the Shopeepay payment method.
	Shopeepay *TestHelpersConfirmationTokenCreatePaymentMethodDataShopeepayParams `form:"shopeepay" json:"shopeepay,omitempty"`
	// If this is a `sofort` PaymentMethod, this hash contains details about the SOFORT payment method.
	Sofort *TestHelpersConfirmationTokenCreatePaymentMethodDataSofortParams `form:"sofort" json:"sofort,omitempty"`
	// This hash contains details about the Stripe balance payment method.
	StripeBalance *TestHelpersConfirmationTokenCreatePaymentMethodDataStripeBalanceParams `form:"stripe_balance" json:"stripe_balance,omitempty"`
	// If this is a `swish` PaymentMethod, this hash contains details about the Swish payment method.
	Swish *TestHelpersConfirmationTokenCreatePaymentMethodDataSwishParams `form:"swish" json:"swish,omitempty"`
	// If this is a TWINT PaymentMethod, this hash contains details about the TWINT payment method.
	TWINT *TestHelpersConfirmationTokenCreatePaymentMethodDataTWINTParams `form:"twint" json:"twint,omitempty"`
	// The type of the PaymentMethod. An additional hash is included on the PaymentMethod with a name matching this value. It contains additional information specific to the PaymentMethod type.
	Type *string `form:"type" json:"type"`
	// If this is a `upi` PaymentMethod, this hash contains details about the UPI payment method.
	Upi *TestHelpersConfirmationTokenCreatePaymentMethodDataUpiParams `form:"upi" json:"upi,omitempty"`
	// If this is an `us_bank_account` PaymentMethod, this hash contains details about the US bank account payment method.
	USBankAccount *TestHelpersConfirmationTokenCreatePaymentMethodDataUSBankAccountParams `form:"us_bank_account" json:"us_bank_account,omitempty"`
	// If this is an `wechat_pay` PaymentMethod, this hash contains details about the wechat_pay payment method.
	WeChatPay *TestHelpersConfirmationTokenCreatePaymentMethodDataWeChatPayParams `form:"wechat_pay" json:"wechat_pay,omitempty"`
	// If this is a `zip` PaymentMethod, this hash contains details about the Zip payment method.
	Zip *TestHelpersConfirmationTokenCreatePaymentMethodDataZipParams `form:"zip" json:"zip,omitempty"`
}

// AddMetadata adds a new key-value pair to the Metadata.
func (p *TestHelpersConfirmationTokenCreatePaymentMethodDataParams) AddMetadata(key string, value string) {
	if p.Metadata == nil {
		p.Metadata = make(map[string]string)
	}

	p.Metadata[key] = value
}

// The selected installment plan to use for this payment attempt.
// This parameter can only be provided during confirmation.
type TestHelpersConfirmationTokenCreatePaymentMethodOptionsCardInstallmentsPlanParams struct {
	// For `fixed_count` installment plans, this is required. It represents the number of installment payments your customer will make to their credit card.
	Count *int64 `form:"count" json:"count,omitempty"`
	// For `fixed_count` installment plans, this is required. It represents the interval between installment payments your customer will make to their credit card.
	// One of `month`.
	Interval *string `form:"interval" json:"interval,omitempty"`
	// Type of installment plan, one of `fixed_count`, `bonus`, or `revolving`.
	Type *string `form:"type" json:"type"`
}

// Installment configuration for payments confirmed using this ConfirmationToken.
type TestHelpersConfirmationTokenCreatePaymentMethodOptionsCardInstallmentsParams struct {
	// The selected installment plan to use for this payment attempt.
	// This parameter can only be provided during confirmation.
	Plan *TestHelpersConfirmationTokenCreatePaymentMethodOptionsCardInstallmentsPlanParams `form:"plan" json:"plan"`
}

// Configuration for any card payments confirmed using this ConfirmationToken.
type TestHelpersConfirmationTokenCreatePaymentMethodOptionsCardParams struct {
	// Installment configuration for payments confirmed using this ConfirmationToken.
	Installments *TestHelpersConfirmationTokenCreatePaymentMethodOptionsCardInstallmentsParams `form:"installments" json:"installments,omitempty"`
}

// Payment-method-specific configuration for this ConfirmationToken.
type TestHelpersConfirmationTokenCreatePaymentMethodOptionsParams struct {
	// Configuration for any card payments confirmed using this ConfirmationToken.
	Card *TestHelpersConfirmationTokenCreatePaymentMethodOptionsCardParams `form:"card" json:"card,omitempty"`
}

// Shipping information for this ConfirmationToken.
type TestHelpersConfirmationTokenCreateShippingParams struct {
	// Shipping address
	Address *AddressParams `form:"address" json:"address"`
	// Recipient name.
	Name *string `form:"name" json:"name"`
	// Recipient phone (including extension)
	Phone       *string                                                      `form:"phone" json:"phone,omitempty"`
	UnsetFields []TestHelpersConfirmationTokenCreateShippingParamsUnsetField `form:"-" json:"-"`
}

// TestHelpersConfirmationTokenCreateShippingParamsUnsetField is the list of fields that can be cleared/unset on TestHelpersConfirmationTokenCreateShippingParams.
type TestHelpersConfirmationTokenCreateShippingParamsUnsetField string

const (
	TestHelpersConfirmationTokenCreateShippingParamsUnsetFieldPhone TestHelpersConfirmationTokenCreateShippingParamsUnsetField = "phone"
)

// AddUnsetField adds a field to the list of fields to clear/unset on this params object.
func (p *TestHelpersConfirmationTokenCreateShippingParams) AddUnsetField(field TestHelpersConfirmationTokenCreateShippingParamsUnsetField) {
	p.UnsetFields = append(p.UnsetFields, field)
}

// Creates a test mode Confirmation Token server side for your integration tests.
type TestHelpersConfirmationTokenCreateParams struct {
	Params `form:"*"`
	// Specifies which fields in the response should be expanded.
	Expand []*string `form:"expand" json:"expand,omitempty"`
	// ID of an existing PaymentMethod.
	PaymentMethod *string `form:"payment_method" json:"payment_method,omitempty"`
	// If provided, this hash will be used to create a PaymentMethod.
	PaymentMethodData *TestHelpersConfirmationTokenCreatePaymentMethodDataParams `form:"payment_method_data" json:"payment_method_data,omitempty"`
	// Payment-method-specific configuration for this ConfirmationToken.
	PaymentMethodOptions *TestHelpersConfirmationTokenCreatePaymentMethodOptionsParams `form:"payment_method_options" json:"payment_method_options,omitempty"`
	// Return URL used to confirm the Intent.
	ReturnURL *string `form:"return_url" json:"return_url,omitempty"`
	// Indicates that you intend to make future payments with this ConfirmationToken's payment method.
	//
	// The presence of this property will [attach the payment method](https://docs.stripe.com/payments/save-during-payment) to the PaymentIntent's Customer, if present, after the PaymentIntent is confirmed and any required actions from the user are complete.
	SetupFutureUsage *string `form:"setup_future_usage" json:"setup_future_usage,omitempty"`
	// Shipping information for this ConfirmationToken.
	Shipping *TestHelpersConfirmationTokenCreateShippingParams `form:"shipping" json:"shipping,omitempty"`
}

// AddExpand appends a new field to expand.
func (p *TestHelpersConfirmationTokenCreateParams) AddExpand(f string) {
	p.Expand = append(p.Expand, &f)
}
