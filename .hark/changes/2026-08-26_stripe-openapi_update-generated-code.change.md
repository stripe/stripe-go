---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/2413
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.5.0-beta.1
---

* Add support for new resources `V2CoreApprovalRequest`, `V2SignalsAccountActivity`, `V2SignalsAccountEvaluation`, and `V2SignalsAccountSignal`
* Add support for `Get` and `List` methods on resource `V2SignalsAccountSignal`
* Add support for `Get` and `New` methods on resource `V2SignalsAccountEvaluation`
* Add support for `Del`, `Get`, and `New` methods on resource `V2SignalsAccountActivity`
* Add support for `Cancel`, `Get`, `List`, and `Update` methods on resource `V2CoreApprovalRequest`
* Add support for `Disable` method on resource `V2MoneyManagementPayoutMethod`
* Add support for `DisableStripeUserAuthentication` on `AccountSessionComponentsPaymentMethodSettingsFeaturesParams`
* ⚠️ Remove support for `PaymentMethodTypes` on `PaymentIntentConfirmParams`, `PaymentIntentParams`, and `SetupIntentParams`
* ⚠️ Change type of `ProductCatalogTrialOffer.Price` and `ProductCatalogTrialOfferEndBehaviorTransition.Price` from `$Price` to `deletable($Price)`
* Add support for `Billie` on `QuotePreviewInvoicePaymentSettingsPaymentMethodOptions`
* Add support for new value `billie` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
* Add support for `PayoutMethods` on `V2CoreAccountDefaultsParams` and `V2CoreAccountDefaults`
* Add support for `Restricted` on `V2CoreVaultGbBankAccount` and `V2CoreVaultUsBankAccount`
* Add support for `EnabledDeliverySchemes` on `V2MoneyManagementPayoutMethodBankAccount`
* ⚠️ Remove support for `EnabledDeliveryOptions` on `V2MoneyManagementPayoutMethodBankAccount`
* Add support for new value `disabled` on enum `V2MoneyManagementPayoutMethodUsageStatus.Payments`
* Add support for new value `disabled` on enum `V2MoneyManagementPayoutMethodUsageStatus.Transfers`
* Add support for event notifications `V2CoreApprovalRequestApprovedEvent`, `V2CoreApprovalRequestCanceledEvent`, `V2CoreApprovalRequestCreatedEvent`, `V2CoreApprovalRequestExpiredEvent`, `V2CoreApprovalRequestFailedEvent`, `V2CoreApprovalRequestRejectedEvent`, and `V2CoreApprovalRequestSucceededEvent` with related object `V2CoreApprovalRequest`
* Add support for event notification `V2SignalsAccountEvaluationCompleteEvent` with related object `V2SignalsAccountEvaluation`
* Add support for error codes `authentication_failure`, `capability_not_active`, `expired_payment_method`, `incorrect_postal_code`, `invalid_canceled_subscription_fields`, and `payment_method_restricted` on `QuotePreviewInvoiceLastFinalizationError`
* Add support for error code `default_payout_method_cannot_be_disabled` on `CannotProceedError`
