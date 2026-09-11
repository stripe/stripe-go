---
title: Update generated code for beta
pr_link: https://github.com/stripe/stripe-go/pull/2363
is_breaking: true
is_stripe_api_change: true
released_in_version: 86.2.0-beta.1
---

* Add support for `Redaction` on `Card`, `Charge`, `CheckoutSession`, `Customer`, `IssuingAuthorization`, `IssuingCard`, `IssuingCardholder`, `IssuingDispute`, `IssuingTransaction`, `PaymentIntent`, `PaymentMethod`, `SetupIntent`, `Source`, and `Token`
* Add support for `DisclaimerVariant` on `CapitalFinancingOffer` and `CapitalFinancingSummaryDetails`
* Add support for `Active` on `FinancialConnectionsAccountStatusDetails` and `FinancialConnectionsAuthorizationStatusDetails`
* Add support for new value `institution_requirement` on enum `FinancialConnectionsAccountStatusDetailsInactive.Cause`
* Change type of `FinancialConnectionsSessionLimitsParams.Accounts` from `longInteger` to `emptyable(longInteger)`
* Add support for `Pause` on `InvoiceCreatePreviewSubscriptionDetailsParams`
* Add support for new value `satispay` on enum `QuotePreviewInvoicePaymentSettings.PaymentMethodTypes`
* Add support for `ReleaseDetails` on `ReserveHold`
* Add support for `BuyerID` on `SharedPaymentGrantedTokenPaymentMethodDetailsBizum` and `SharedPaymentGrantedTokenPaymentMethodDetailsBlik`
* Add support for `Fingerprint` on `SharedPaymentGrantedTokenPaymentMethodDetailsPix`
* Add support for new value `money_manager` on enums `EventsV2CoreAccountLinkReturnedEvent.Configurations`, `V2CoreAccountLinkUseCaseAccountOnboarding.Configurations`, and `V2CoreAccountLinkUseCaseAccountUpdate.Configurations`
* ⚠️ Add support for new value `money_manager` on enum `V2CoreAccount.AppliedConfigurations`
* ⚠️ Remove support for value `storer` from enum `V2CoreAccount.AppliedConfigurations`
* Add support for `MoneyManager` on `V2CoreAccountConfigurationParams`, `V2CoreAccountConfiguration`, `V2CoreAccountIdentityAttestationsTermsOfServiceParams`, `V2CoreAccountIdentityAttestationsTermsOfService`, and `V2CoreAccountTokenIdentityAttestationsTermsOfServiceParams`
* ⚠️ Remove support for `Storer` on `V2CoreAccountConfigurationParams`, `V2CoreAccountConfiguration`, `V2CoreAccountIdentityAttestationsTermsOfServiceParams`, `V2CoreAccountIdentityAttestationsTermsOfService`, and `V2CoreAccountTokenIdentityAttestationsTermsOfServiceParams`
* Add support for new values `business_storage.inbound.eur`, `business_storage.inbound.gbp`, `business_storage.inbound.usd`, `business_storage.outbound.eur`, `business_storage.outbound.gbp`, `business_storage.outbound.usd`, `received_credits.bank_accounts`, and `received_debits.bank_accounts` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Capability` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for new value `money_manager` on enums `V2CoreAccountFutureRequirementsEntryImpactRestrictsCapability.Configuration` and `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Configuration`
* ⚠️ Remove support for `MaximumRps` on `V2CoreBatchJobParams` and `V2CoreBatchJob`
* Add support for `BIC` on `V2MoneyManagementFinancialAddressCredentialsUsBankAccount`
* ⚠️ Remove support for `SwiftCode` on `V2MoneyManagementFinancialAddressCredentialsUsBankAccount`
* Add support for `Processing` on `V2MoneyManagementOutboundPaymentStatusDetails` and `V2MoneyManagementOutboundTransferStatusDetails`
* Add support for new values `fx_rate_drift_exceeded_after_review`, `payout_method_amount_limit_exceeded`, and `review_rejected` on enum `V2MoneyManagementOutboundPaymentStatusDetailsFailed.Reason`
* Add support for new values `fx_rate_drift_exceeded_after_review` and `review_rejected` on enum `V2MoneyManagementOutboundTransferStatusDetailsFailed.Reason`
* Add support for `AccountHolderName` on `V2MoneyManagementReceivedCreditBankTransferUsBankAccount`
* Add support for new value `capability_inactive` on enum `V2MoneyManagementReceivedDebitStatusDetailsFailed.Reason`
* Add support for `Statuses` on `V2MoneyManagementFinancialAccountListParams`
* ⚠️ Remove support for `Status` on `V2MoneyManagementFinancialAccountListParams`
* Add support for event notifications `V2CoreAccountIncludingConfigurationMoneyManagerCapabilityStatusUpdatedEvent` and `V2CoreAccountIncludingConfigurationMoneyManagerUpdatedEvent` with related object `V2CoreAccount`
* Add support for event notification `V2MoneyManagementOutboundPaymentUnderReviewEvent` with related object `V2MoneyManagementOutboundPayment`
* Add support for event notification `V2MoneyManagementOutboundTransferUnderReviewEvent` with related object `V2MoneyManagementOutboundTransfer`
* ⚠️ Remove support for event notifications `V2CoreAccountIncludingConfigurationStorerCapabilityStatusUpdatedEvent` and `V2CoreAccountIncludingConfigurationStorerUpdatedEvent` with related object `V2CoreAccount`
* Add support for error codes `anomalous_money_movement_request`, `failed_tax_calculation`, `financial_account_balance_does_not_support_currency`, `financial_account_capability_not_enabled`, and `financial_account_capability_restricted` on `QuotePreviewInvoiceLastFinalizationError`
* Add support for error code `default_us_bank_account_cannot_be_archived` on `CannotProceedError`
