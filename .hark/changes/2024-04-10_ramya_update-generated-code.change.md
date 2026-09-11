---
title: Update generated code
pr_link: https://github.com/stripe/stripe-go/pull/1841
released_in_version: 78.0.0
---

* This release changes the pinned API version to `2024-04-10`. Please read the [API Changelog](https://docs.stripe.com/changelog/2024-04-10) and carefully review the API changes before upgrading.

### ⚠️ Breaking changes

 * When no `x-stripe-should-retry` header is set in the response, the library now retries all requests with `status >= 500`, not just non-POST methods.
 * Change the type on the status of TerminalReader object from string to enum with values of `TerminalReaderStatusOffline` and `TerminalReaderStatusOnline`
 * Rename `Features` to `MarketingFeatures` on `ProductCreateOptions`, `ProductUpdateOptions`, and `Product`.

#### ⚠️ Removal of enum values, properties and events that are no longer part of the publicly documented Stripe API
* Remove `SubscriptionPause` from `BillingPortalConfigurationFeatures ` and `BillingPortalConfigurationFeaturesParams ` as the feature to pause subscription on the portal has been deprecated.
* Remove deprecated values for the `BalanceTransactionType` enum by removing the below constants
    * `BalanceTransactionTypeObligationInbound`
    * `BalanceTransactionTypeObligationPayout`
    * `BalanceTransactionTypeObligationPayoutFailure`
    * `BalanceTransactionTypeObligationReversalOutbound`
 * Remove deprecated value for the `ClimateSupplierRemovalPathway` enum by removing the constant `ClimateSupplierRemovalPathwayVarious`
 * Remove deprecated events types
    * `EventTypeInvoiceItemUpdated`
    * `EventTypeOrderCreated`
    * `EventTypeRecipientCreated`
    * `EventTypeRecipientDeleted`
    * `EventTypeRecipientUpdated`
    * `EventTypeSKUCreated`
    * `EventTypeSKUDeleted`
 * Remove the field `RequestIncrementalAuthorization` on the `PaymentIntentPaymentMethodOptionsCardPresentParams` struct - this was shipped by mistake
 * Remove support for `id_bank_transfer`, `multibanco, netbanking`, `pay_by_bank`, and `upi` on `PaymentMethodConfiguration`. TODO - List the affected types and constants
 * Remove deprecated value for the `SetupIntentPaymentMethodOptionsCardRequestThreeDSecure` enum by removing the constant `SetupIntentPaymentMethodOptionsCardRequestThreeDSecureChallengeOnly`
 * Remove deprecated value for the `TaxRateTaxType` enum by removing the constant `TaxRateTaxTypeServiceTax`
 * Remove `PaymentIntentPaymentMethodData*Params` in favor of reusing existing `PaymentMethodData*Params` for all the payment method types.
    * Remove  `PaymentIntentPaymentMethodDataBLIKParams` in favor of `PaymentMethodDataBLIKParams`
    * Remove  `PaymentIntentPaymentMethodDataCashAppParams` in favor of `PaymentMethodDataCashAppParams`
    * Remove  `PaymentIntentPaymentMethodDataCustomerBalanceParams` in favor of `PaymentMethodDataCustomerBalanceParams`
    * Remove  `PaymentIntentPaymentMethodDataKonbiniParams` in favor of `PaymentMethodDataKonbiniParams`
    * Remove  `PaymentIntentPaymentMethodDataLinkParams` in favor of `PaymentMethodDataLinkParams`
    * Remove  `PaymentIntentPaymentMethodDataPayNowParams` in favor of `PaymentMethodDataPayNowParams`
    * Remove  `PaymentIntentPaymentMethodDataPaypalParams` in favor of `PaymentMethodDataPaypalParams`
    * Remove  `PaymentIntentPaymentMethodDataPixParams` in favor of `PaymentMethodDataPixParams`
    * Remove  `PaymentIntentPaymentMethodDataPromptPayParams` in favor of `PaymentMethodDataPromptPayParams`
    * Remove  `PaymentIntentPaymentMethodDataRevolutPayParams` in favor of `PaymentMethodDataRevolutPayParams`
    * Remove  `PaymentIntentPaymentMethodDataUSBankAccounParams` in favor of `PaymentMethodDataUSBankAccounParams`
    * Remove  `PaymentIntentPaymentMethodDataZipParams` in favor of `PaymentMethodDataZipParams`
 * Remove the legacy field `InvoiceRenderingOptionsParams` in `Invoice`, `InvoiceParams`. Use `InvoiceRenderingParams` instead.
