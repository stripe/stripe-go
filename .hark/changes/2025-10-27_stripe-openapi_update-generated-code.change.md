---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/2156
is_stripe_api_change: true
released_in_version: 83.2.0-beta.1
---

* Add support for `Update` method on resource `V2MoneyManagementFinancialAccount`
* Add support for `ConfirmMicrodeposits`, `List`, and `SendMicrodeposits` methods on resource `V2CoreVaultUSBankAccount`
* Add support for `List` method on resource `V2CoreVaultGBBankAccount`
* Add support for new value `verification_data_not_found` on enums `BankAccountFutureRequirementsErrors.Code` and `BankAccountRequirementsErrors.Code`
* Add support for `PaymentPortalURL` on `ChargePaymentMethodDetailsRechnung`, `PaymentAttemptRecordPaymentMethodDetailsRechnung`, and `PaymentRecordPaymentMethodDetailsRechnung`
* Add support for `TaxIDElement` on `CustomerSessionComponentsParams` and `CustomerSessionComponents`
* Add support for `StartingAfter` on `PaymentAttemptRecordListParams`
* Add support for new value `solana` on enums `PaymentAttemptRecordPaymentMethodDetailsCrypto.Network` and `PaymentRecordPaymentMethodDetailsCrypto.Network`
* Add support for `Reference` on `PaymentIntentAmountDetailsLineItemPaymentMethodOptionsKlarna`, `PaymentIntentAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`, `PaymentIntentCaptureAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`, `PaymentIntentConfirmAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`, and `PaymentIntentIncrementAuthorizationAmountDetailsLineItemsPaymentMethodOptionsKlarnaParams`
* Add support for `SubscriptionReference` on `PaymentIntentAmountDetailsLineItemPaymentMethodOptionsKlarna`
* Add support for `Closed` on `V2CoreAccountListParams` and `V2CoreAccount`
* Add support for new value `payment_method` on enum `V2CoreAccountConfigurationCustomerAutomaticIndirectTax.LocationSource`
* Add support for `USD` on `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrenciesParams` and `V2CoreAccountConfigurationStorerCapabilitiesHoldsCurrencies`
* Add support for new values `application_custom` and `application_express` on enum `V2CoreAccountDefaultsResponsibilities.FeesCollector`
* Add support for `RepresentativeDeclaration` on `V2CoreAccountIdentityAttestationsParams` and `V2CoreAccountIdentityAttestations`
* Add support for new value `holds_currencies.usd` on enum `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for `Verification` on `V2CoreVaultUsBankAccount`
* Add support for `V1ID` on `EventsV2MoneyManagementTransactionCreatedEvent`
* Remove support for thin event `V2BillingBillSettingUpdatedEvent` with related object `V2BillingBillSetting`
* Add support for error code `payment_intent_rate_limit_exceeded` on `QuotePreviewInvoiceLastFinalizationError`
* Add support for error codes `blocked_payout_method_crypto_wallet` and `unsupported_payout_method_crypto_wallet` on `BlockedByStripeError`
* Add support for error code `outbound_flow_from_closed_financial_account_unsupported` on `FeatureNotEnabledError`
* Add support for error code `limit_payout_method_crypto_wallet` on `QuotaExceededError`
