---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/1957
is_stripe_api_change: true
released_in_version: 81.2.0
---

* Add support for `NetworkAdviceCode` and `NetworkDeclineCode` on `ChargeOutcome`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
* Add support for new values `payout_minimum_balance_hold` and `payout_minimum_balance_release` on enum `BalanceTransactionType`
* Add support for `CreditsApplicationInvoiceVoided` on `BillingCreditBalanceTransactionCredit`
* Change type of `BillingCreditBalanceTransactionCreditType` from `literal('credits_granted')` to `enum('credits_application_invoice_voided'|'credits_granted')`
* Add support for `AllowRedisplay` on `Card` and `Source`
* Add support for `RegulatedStatus` on `Card`, `ChargePaymentMethodDetailsCard`, `ConfirmationTokenPaymentMethodPreviewCard`, and `PaymentMethodCard`
* Add support for `Funding` on `ChargePaymentMethodDetailsAmazonPay` and `ChargePaymentMethodDetailsRevolutPay`
* Add support for `NetworkTransactionID` on `ChargePaymentMethodDetailsCard`
* Add support for `ReferencePrefix` on `CheckoutSessionPaymentMethodOptionsBacsDebitMandateOptionsParams`, `CheckoutSessionPaymentMethodOptionsBacsDebitMandateOptions`, `CheckoutSessionPaymentMethodOptionsSepaDebitMandateOptionsParams`, `CheckoutSessionPaymentMethodOptionsSepaDebitMandateOptions`, `PaymentIntentConfirmPaymentMethodOptionsBacsDebitMandateOptionsParams`, `PaymentIntentConfirmPaymentMethodOptionsSepaDebitMandateOptionsParams`, `PaymentIntentPaymentMethodOptionsBacsDebitMandateOptionsParams`, `PaymentIntentPaymentMethodOptionsBacsDebitMandateOptions`, `PaymentIntentPaymentMethodOptionsSepaDebitMandateOptionsParams`, `PaymentIntentPaymentMethodOptionsSepaDebitMandateOptions`, `SetupIntentConfirmPaymentMethodOptionsBacsDebitMandateOptionsParams`, `SetupIntentConfirmPaymentMethodOptionsSepaDebitMandateOptionsParams`, `SetupIntentPaymentMethodOptionsBacsDebitMandateOptionsParams`, `SetupIntentPaymentMethodOptionsBacsDebitMandateOptions`, `SetupIntentPaymentMethodOptionsSepaDebitMandateOptionsParams`, and `SetupIntentPaymentMethodOptionsSepaDebitMandateOptions`
* Add support for new values `al_tin`, `am_tin`, `ao_tin`, `ba_tin`, `bb_tin`, `bs_tin`, `cd_nif`, `gn_nif`, `kh_tin`, `me_pib`, `mk_vat`, `mr_nif`, `np_pan`, `sn_ninea`, `sr_fin`, `tj_tin`, `ug_tin`, `zm_tin`, and `zw_tin` on enums `CheckoutSessionCustomerDetailsTaxIdsType`, `InvoiceCustomerTaxIdsType`, `TaxCalculationCustomerDetailsTaxIdsType`, `TaxIdType`, and `TaxTransactionCustomerDetailsTaxIdsType`
* Add support for `VisaCompliance` on `DisputeEvidenceDetailsEnhancedEligibility`, `DisputeEvidenceEnhancedEvidenceParams`, and `DisputeEvidenceEnhancedEvidence`
* Add support for new value `request_signature` on enum `ForwardingRequestReplacements`
* Add support for `AccountHolderAddress` and `BankAddress` on `FundingInstructionsBankTransferFinancialAddressesIban`, `FundingInstructionsBankTransferFinancialAddressesSortCode`, `FundingInstructionsBankTransferFinancialAddressesSpei`, `FundingInstructionsBankTransferFinancialAddressesZengin`, `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesIban`, `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesSortCode`, `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesSpei`, and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesZengin`
* Add support for `AccountHolderName` on `FundingInstructionsBankTransferFinancialAddressesSpei` and `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressesSpei`
* Add support for `DisabledReason` on `InvoiceAutomaticTax`, `SubscriptionAutomaticTax`, `SubscriptionScheduleDefaultSettingsAutomaticTax`, and `SubscriptionSchedulePhasesAutomaticTax`
* Add support for `TaxID` on `IssuingAuthorizationMerchantData` and `IssuingTransactionMerchantData`
* Add support for `TrialPeriodDays` on `PaymentLinkSubscriptionDataParams`
* Add support for `Al`, `Am`, `Ao`, `Ba`, `Bb`, `Bs`, `Cd`, `Gn`, `Kh`, `Me`, `Mk`, `Mr`, `Np`, `Pe`, `Sn`, `Sr`, `Tj`, `Ug`, `Uy`, `Zm`, and `Zw` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
