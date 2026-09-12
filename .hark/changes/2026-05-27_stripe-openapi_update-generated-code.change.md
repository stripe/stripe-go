---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/2350
is_breaking: true
is_stripe_api_change: true
released_in_version: 85.3.0-beta.1
---

* Add support for `Pause` method on resource `Subscription`
* Add support for `Get` method on resource `V2IamActivityLog`
* Add support for new value `mastercard` on enum `IssuingSettlement.Network`
* ⚠️ Change type of `ProductCatalogTrialOfferEndBehaviorTransition.Price` from `string` to `expandable($Price)`
* Add support for `AmountPaidOffStripe` on `QuotePreviewInvoice`
* Add support for new value `twint` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
* Add support for `Discountable` on `QuotePreviewSubscriptionSchedulePhaseAddInvoiceItem`
* Add support for `Bizum` and `Scalapay` on `SharedPaymentGrantedTokenPaymentMethodDetails`
* Add support for new values `bizum` and `scalapay` on enum `SharedPaymentGrantedTokenPaymentMethodDetails.Type`
* Add support for `PaymentBehavior` on `SubscriptionResumeParams`
* Add support for `StatusDetails` on `Subscription`
* Add support for new values `ao_bank_account`, `az_bank_account`, `bd_bank_account`, `bo_bank_account`, `br_bank_account`, `cl_bank_account`, `ga_bank_account`, `gh_bank_account`, `gi_bank_account`, `hn_bank_account`, `kr_bank_account`, `kz_bank_account`, `la_bank_account`, `ne_bank_account`, `ng_bank_account`, `ni_bank_account`, `py_bank_account`, `sa_bank_account`, `sm_bank_account`, and `uy_bank_account` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
* ⚠️ Change type of `V2MoneyManagementReceivedCreditBankTransferGbBankAccount.Network` from `literal('fps')` to `enum('chaps'|'fps')`
* Add support for error codes `payment_method_microdeposit_processing_error` and `siret_invalid` on `QuotePreviewInvoiceLastFinalizationError`
