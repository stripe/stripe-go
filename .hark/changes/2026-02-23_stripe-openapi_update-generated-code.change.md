---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/2276
is_stripe_api_change: true
released_in_version: 84.5.0-beta.1
---

* Add support for `SmartDisputes` on `AccountSettingsParams`, `AccountSettings`, `V2CoreAccountConfigurationMerchantParams`, and `V2CoreAccountConfigurationMerchant`
* Add support for `EmailCustomersOnSuccessfulPayment` on `AccountSettingsPaymentsParams` and `AccountSettingsPayments`
* Add support for `ManagedPayments` on `CheckoutSessionParams`, `CheckoutSession`, `PaymentIntent`, `SetupIntent`, and `Subscription`
* Add support for new value `lk_vat` on enums `CheckoutSessionCollectedInformationTaxIds.Type`, `OrderTaxDetailsTaxId.Type`, and `QuotePreviewInvoiceCustomerTaxIds.Type`
* Add support for new value `pay_by_bank` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
* Add support for new values `bt_bank_account`, `cr_bank_account`, `do_bank_account`, `gt_bank_account`, `md_bank_account`, `mk_bank_account`, `mo_bank_account`, `mz_bank_account`, `pe_bank_account`, `pk_bank_account`, `tw_bank_account`, and `uz_bank_account` on enum `V2CoreAccountConfigurationRecipientDefaultOutboundDestination.Type`
* Add support for `Purpose` on `V2MoneyManagementOutboundPaymentParams` and `V2MoneyManagementOutboundPayment`
* Add support for `BranchNumber` and `SwiftCode` on `V2MoneyManagementPayoutMethodBankAccount`
* Add support for error codes `storer_capability_missing` and `storer_capability_not_active` on `QuotePreviewInvoiceLastFinalizationError`
