---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/2249
is_stripe_api_change: true
released_in_version: 84.4.0-beta.1
---

* Add support for new resource `FinancialConnectionsAuthorization`
* Add support for `Get` method on resource `FinancialConnectionsAuthorization`
* Add support for `DetachPayment` method on resource `Invoice`
* Remove support for `Cancel`, `ListLineItems`, and `Reopen` methods on resource `Order`
* Remove support for `AttachCadence` method on resource `Subscription`
* Add support for `AdditionalFiles` and `Site` on `AccountSettingsPaypayPaymentsParams` and `AccountSettingsPaypayPayments`
* Remove support for `Capital` on `AccountSettings`
* Add support for new value `pl_nip` on enums `CheckoutSessionCollectedInformationTaxIds.Type`, `OrderTaxDetailsTaxId.Type`, and `QuotePreviewInvoiceCustomerTaxIds.Type`
* Add support for new value `capital.financing_summary.line_of_credit_update` on enum `Event.Type`
* Add support for `Authorization` and `StatusDetails` on `FinancialConnectionsAccount`
* Add support for `RelinkOptions` on `FinancialConnectionsSessionParams` and `FinancialConnectionsSession`
* Add support for `RelinkResult` on `FinancialConnectionsSession`
* Remove support for `BillingCadence` on `InvoiceCreatePreviewParams`, `SubscriptionParams`, and `Subscription`
* Remove support for `BillingCadenceDetails` on `InvoiceParent` and `QuotePreviewInvoiceParent`
* Remove support for value `billing_cadence_details` from enums `InvoiceParent.Type` and `QuotePreviewInvoiceParent.Type`
* Add support for `CarRentalData`, `FlightData`, and `LodgingData` on `PaymentIntentPaymentDetails`
* Add support for new values `ae_bank_account`, `ag_bank_account`, `bh_bank_account`, `gm_bank_account`, `hk_bank_account`, `kh_bank_account`, `lc_bank_account`, `mc_bank_account`, `mg_bank_account`, `my_bank_account`, `qa_bank_account`, `rw_bank_account`, `th_bank_account`, `tt_bank_account`, and `vn_bank_account` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
* Add support for `AlternativeReference` on `V2CoreVaultGbBankAccount`, `V2CoreVaultUsBankAccount`, and `V2MoneyManagementPayoutMethod`
* Add support for `AccountHolderAddress` and `AccountHolderName` on `V2MoneyManagementFinancialAddressCredentialsUsBankAccount`
* Add support for `Fingerprint` on `V2MoneyManagementPayoutMethodCard`
* Add support for snapshot event `EventTypeInvoicePaymentDetached` with resource `InvoicePayment`
* Add support for error code `request_blocked` on `QuotePreviewInvoiceLastFinalizationError`
* Add support for error codes `blocked_payout_method` and `unsupported_payout_method` on `BlockedByStripeError`
* Add support for error code `invalid_payout_method_data` on `InvalidPayoutMethodError`
* Add support for error code `limit_payout_method` on `QuotaExceededError`
