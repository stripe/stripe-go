---
title: Update generated code
pr_url: https://github.com/stripe/stripe-go/pull/2235
is_stripe_api_change: true
released_in_version: 84.1.0
---

* Add support for new resources `V2CoreAccountLink`, `V2CoreAccountPersonToken`, `V2CoreAccountPerson`, `V2CoreAccountToken`, and `V2CoreAccount`
* Add support for `Get` and `New` methods on resources `V2CoreAccountPersonToken` and `V2CoreAccountToken`
* Add support for `New` method on resource `V2CoreAccountLink`
* Add support for `Close`, `Get`, `List`, `New`, and `Update` methods on resource `V2CoreAccount`
* Add support for `Del`, `Get`, `List`, `New`, and `Update` methods on resource `V2CoreAccountPerson`
* Add support for `CustomerAccount` on `BillingCreditBalanceSummaryParams`, `BillingCreditBalanceSummary`, `BillingCreditBalanceTransactionListParams`, `BillingCreditGrantListParams`, `BillingCreditGrantParams`, `BillingCreditGrant`, `BillingPortalSessionParams`, `BillingPortalSession`, `CashBalance`, `CheckoutSessionListParams`, `CheckoutSessionParams`, `CheckoutSession`, `ConfirmationTokenPaymentMethodPreview`, `CreditNoteListParams`, `CreditNote`, `CustomerBalanceTransaction`, `CustomerCashBalanceTransaction`, `CustomerSessionParams`, `CustomerSession`, `Customer`, `Discount`, `FinancialConnectionsAccountAccountHolder`, `FinancialConnectionsAccountListAccountHolderParams`, `FinancialConnectionsSessionAccountHolderParams`, `FinancialConnectionsSessionAccountHolder`, `InvoiceCreatePreviewParams`, `InvoiceItemListParams`, `InvoiceItemParams`, `InvoiceItem`, `InvoiceListParams`, `InvoiceParams`, `Invoice`, `PaymentIntentListParams`, `PaymentIntentParams`, `PaymentIntent`, `PaymentMethodAttachParams`, `PaymentMethodListParams`, `PaymentMethod`, `PromotionCodeListParams`, `PromotionCodeParams`, `PromotionCode`, `QuoteListParams`, `QuoteParams`, `Quote`, `SetupAttempt`, `SetupIntentListParams`, `SetupIntentParams`, `SetupIntent`, `SubscriptionListParams`, `SubscriptionParams`, `SubscriptionScheduleListParams`, `SubscriptionScheduleParams`, `SubscriptionSchedule`, `Subscription`, `TaxIdListOwnerParams`, `TaxIdOwnerParams`, `TaxIdOwner`, and `TaxId`
* Add support for `Metadata` on `CheckoutSessionLineItemParams` and `LineItem`
* Add support for `PaytoPayments` on `AccountCapabilitiesParams` and `AccountCapabilities`
* Add support for `Signer` on `AccountDocumentsProofOfRegistrationParams` and `AccountDocumentsProofOfUltimateBeneficialOwnershipParams`
* Add support for `BillingCycleAnchor` on `BillingPortalConfigurationFeaturesSubscriptionUpdateParams` and `BillingPortalConfigurationFeaturesSubscriptionUpdate`
* Add support for `Payto` on `ChargePaymentMethodDetails`, `CheckoutSessionPaymentMethodOptionsParams`, `CheckoutSessionPaymentMethodOptions`, `ConfirmationTokenPaymentMethodDataParams`, `ConfirmationTokenPaymentMethodPreview`, `InvoicePaymentSettingsPaymentMethodOptionsParams`, `InvoicePaymentSettingsPaymentMethodOptions`, `MandatePaymentMethodDetails`, `PaymentAttemptRecordPaymentMethodDetails`, `PaymentIntentConfirmPaymentMethodDataParams`, `PaymentIntentConfirmPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodDataParams`, `PaymentIntentPaymentMethodOptionsParams`, `PaymentIntentPaymentMethodOptions`, `PaymentMethodConfigurationParams`, `PaymentMethodConfiguration`, `PaymentMethodParams`, `PaymentMethod`, `PaymentRecordPaymentMethodDetails`, `SetupAttemptPaymentMethodDetails`, `SetupIntentConfirmPaymentMethodDataParams`, `SetupIntentConfirmPaymentMethodOptionsParams`, `SetupIntentPaymentMethodDataParams`, `SetupIntentPaymentMethodOptionsParams`, `SetupIntentPaymentMethodOptions`, `SubscriptionPaymentSettingsPaymentMethodOptionsParams`, and `SubscriptionPaymentSettingsPaymentMethodOptions`
* Add support for `ExpectedDebitDate` on `ChargePaymentMethodDetailsAcssDebit`, `ChargePaymentMethodDetailsAuBecsDebit`, `ChargePaymentMethodDetailsBacsDebit`, `ChargePaymentMethodDetailsNzBankAccount`, `ChargePaymentMethodDetailsSepaDebit`, `ChargePaymentMethodDetailsUsBankAccount`, `PaymentAttemptRecordPaymentMethodDetailsAcssDebit`, `PaymentAttemptRecordPaymentMethodDetailsAuBecsDebit`, `PaymentAttemptRecordPaymentMethodDetailsBacsDebit`, `PaymentAttemptRecordPaymentMethodDetailsNzBankAccount`, `PaymentAttemptRecordPaymentMethodDetailsSepaDebit`, `PaymentAttemptRecordPaymentMethodDetailsUsBankAccount`, `PaymentRecordPaymentMethodDetailsAcssDebit`, `PaymentRecordPaymentMethodDetailsAuBecsDebit`, `PaymentRecordPaymentMethodDetailsBacsDebit`, `PaymentRecordPaymentMethodDetailsNzBankAccount`, `PaymentRecordPaymentMethodDetailsSepaDebit`, and `PaymentRecordPaymentMethodDetailsUsBankAccount`
* Add support for `LineItems` on `CheckoutSessionParams`
* Add support for new value `mollie` on enums `ConfirmationTokenPaymentMethodPreviewIdeal.Bank`, `PaymentAttemptRecordPaymentMethodDetailsIdeal.Bank`, and `PaymentRecordPaymentMethodDetailsIdeal.Bank`
* Add support for new value `MLLENL2A` on enums `ConfirmationTokenPaymentMethodPreviewIdeal.BIC`, `PaymentAttemptRecordPaymentMethodDetailsIdeal.BIC`, and `PaymentRecordPaymentMethodDetailsIdeal.BIC`
* Add support for new value `payto` on enums `ConfirmationTokenPaymentMethodPreview.Type` and `PaymentMethod.Type`
* Add support for `Invoice` on `CustomerCustomerBalanceTransactionListParams`
* Add support for `RelatedCustomerAccount` on `IdentityVerificationSessionListParams`, `IdentityVerificationSessionParams`, and `IdentityVerificationSession`
* Add support for new value `payto` on enums `InvoicePaymentSettings.PaymentMethodTypes` and `SubscriptionPaymentSettings.PaymentMethodTypes`
* Add support for `Subtotal` on `InvoiceLineItem`
* Add support for `AuthorizationCode`, `Description`, `IIN`, `Installments`, `Issuer`, `NetworkAdviceCode`, `NetworkDeclineCode`, and `StoredCredentialUsage` on `PaymentAttemptRecordPaymentMethodDetailsCard` and `PaymentRecordPaymentMethodDetailsCard`
* Add support for new value `payto` on enums `PaymentIntent.ExcludedPaymentMethodTypes` and `SetupIntent.ExcludedPaymentMethodTypes`
* Add support for new value `payto` on enum `PaymentLink.PaymentMethodTypes`
* Add support for `AllowRedisplay` on `PaymentMethodListParams`
* Add support for `ReportedBy` on `PaymentRecord`
* Add support for `Changes` on `V2CoreEvent`
* Add support for error code `account_token_required_for_v2_account` on `Error`, `InvoiceLastFinalizationError`, `PaymentIntentLastPaymentError`, `SetupAttemptSetupError`, `SetupIntentLastSetupError`, and `StripeError`
