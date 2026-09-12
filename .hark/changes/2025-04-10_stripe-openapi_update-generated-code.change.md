---
title: Update generated code for beta
pr_url: https://github.com/stripe/stripe-go/pull/2019
is_stripe_api_change: true
released_in_version: 82.1.0-beta.2
---

### Breaking changes
* Change type of `V2MoneyManagementReceivedDebit.StatusTransitions` from `an object` to `nullable(an object)`
* Remove support for values `bank_accounts.local_uk`, `bank_accounts.wire_uk`, `cards_uk`, and `crypto_wallets_v2` from enum `EventsV2CoreAccountIncludingConfigurationRecipientCapabilityStatusUpdatedEvent.UpdatedCapability`

### Additions
* Add support for new resources `Privacy.RedactionJobRootObjects`, `Privacy.RedactionJobValidationError`, and `Privacy.RedactionJob`
* Add support for `Cancel`, `Get`, `List`, `New`, `Run`, `Update`, and `Validate` methods on resource `RedactionJob`
* Add support for `Get` and `List` methods on resource `RedactionJobValidationError`
* Add support for `MinorityOwnedBusinessDesignation` on `AccountBusinessProfileParams` and `AccountBusinessProfile`
* Add support for `ExportTaxTransactions` and `PaymentDisputes` on `AccountSessionComponentsParams`
* Add support for new value `tax_id_prohibited` on enums `InvoiceLastFinalizationError.Code`, `PaymentIntentLastPaymentError.Code`, `QuotePreviewInvoiceLastFinalizationError.Code`, `SetupAttemptSetupError.Code`, `SetupIntentLastSetupError.Code`, and `StripeError.Code`
* Add support for new value `verification_legal_entity_structure_mismatch` on enums `BankAccountFutureRequirementsErrors.Code` and `BankAccountRequirementsErrors.Code`
* Add support for new value `fixed_term_loan` on enum `CapitalFinancingOffer.Type`
* Add support for `WalletOptions` on `CheckoutSessionParams` and `CheckoutSession`
* Add support for new values `privacy.redaction_job.canceled`, `privacy.redaction_job.created`, `privacy.redaction_job.ready`, `privacy.redaction_job.succeeded`, and `privacy.redaction_job.validation_error` on enum `Event.Type`
* Add support for `Klarna` on `PaymentMethodDomain`
* Add support for `In` on `TaxRegistrationCountryOptionsParams` and `TaxRegistrationCountryOptions`
