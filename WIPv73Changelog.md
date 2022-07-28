# Changelog

Default API version changed to "2022-07-28".

## Added

- `CheckoutSessionSetupIntentDataParams.Metadata`
- Invoice `UpcomingLines` method
- Add `ShippingCost` and `ShippingDetails` properties to `CheckoutSession` resource.
- Add `CheckoutSessionShippingCostTax` and `CheckoutSessionShippingCost` classes
- Add `IssuingCardCancellationReasonDesignRejected` constant to `IssuingCardCancellationReason`.
- Add `Validate` field to `Customer` resource.
- Add `Validate` field to `PaymentSourceParams`.
- Add `SetupAttemptPaymentMethodDetailsCardThreeDSecureResultExempted` constant in `SetupAttemptPaymentMethodDetailsCardThreeDSecureResult`.
- Add `SKUPackageDimensionsParams` and `SKUPackageDimensions`.
- Add `SourceACHCreditTransfer`, `SourceACHDebit`, `SourceACSSDebit`, `SourceAlipay`, `SourceAUBECSDebit`, `SourceBancontact`, `SourceCard`, `SourceCardPresent`, `SourceEPS`, `SourceGiropay`, `SourceIDEAL`,  `SourceKlarna`, `SourceMultibanco`, `SourceP24`, `SourceRedirect`, `SourceSEPACreditTransfer`, `SourceSEPADebit`, `SourceSofort`, `SourceThreeDSecure`, and `SourceWechat`. Additionally, matching fields in `Source` struct - `ACHCreditTransfer`, `ACHDebit`, `ACSSDebit`, `Alipay`, `AUBECSDebit`, `Bancontact`, `Card`, `CardPresent`, `EPS`, `Giropay`, `IDEAL`, `Klarna`, `Multibanco`, `P24`, `SEPACreditTransfer`, `SEPADebit`, `Sofort`, `ThreeDSecure`, and `Wechat`
- Add `SourceTransactionACHCreditTransfer`, `SourceTransactionCHFCreditTransfer`, `SourceTransactionGBPCreditTransfer`, `SourceTransactionPaperCheck`, and `SourceTransactionSEPACreditTransfer` and matching `SourceTransaction.ACHCreditTransfer`, `SourceTransaction.CHFCreditTransfer`, `SourceTransaction.GBPCreditTransfer`, `SourceTransaction.PaperCheck`, and `SourceTransaction.SEPACreditTransfer`.
- Add `Subscription.DeleteDiscount` methods.
- Add `SubscriptionItemUsageRecordSummariesParams`
- Add `UsageRecordSummary` `UsageRecordSummaries`, and `UsageRecordSummaryList` methods in `SubscriptionItem`
- Add `SubscriptionSchedulePhaseBillingCycleAnchor`, `SubscriptionSchedulePhaseBillingCycleAnchorAutomatic`, and `SubscriptionSchedulePhaseBillingCycleAnchorPhaseStart`
- Add `SubscriptionSchedulePhaseInvoiceSettings` and `SubscriptionSchedulePhaseInvoiceSettingsParams `
- `TerminalLocation` `UnmarshalJSON` - make `TerminalLocation` expandable

## Changed

- Renamed files
  | Old name | New name |
  | --- | --- |
  | fee.go | applicationfee.go |
  | fee/client.go | applicationfee/client.go |
  | sub.go | subscription.go |
  | sub/client.go | subscription/client.go |
  | subitem.go | subscriptionitem.go |
  | subitem/client.go | subscriptionitem/client.go |
  | subschedule.go | subscriptionschedule.go |
  | subschedule/client.go | subscriptionschedule/client.go |
  | reversal.go | transferreversal.go |
  | reversal/client.go | transferreversal/client.go |

- Renamed structs, fields, enums, and methods to be consistent with the library's naming conventions and with the other Stripe SDKs.

  - `Ach` to `ACH`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Account | `AccountCapabilitiesUSBankAccountAchPayments` | `AccountCapabilitiesUSBankAccountACHPayments` |
    | Account | `AccountCapabilitiesUSBankAccountAchPaymentsActive` | `AccountCapabilitiesUSBankAccountACHPaymentsActive` |
    | Account | `AccountCapabilitiesUSBankAccountAchPaymentsInactive` | `AccountCapabilitiesUSBankAccountACHPaymentsInactive` |
    | Account | `AccountCapabilitiesUSBankAccountAchPaymentsPending` | `AccountCapabilitiesUSBankAccountACHPaymentsPending` |
    | Account | `AccountCapabilitiesUSBankAccountAchPaymentsParams` | `AccountCapabilitiesUSBankAccountACHPaymentsParams` |
    | Account | `AccountCapabilitiesParams.USBankAccountAchPayments` | `AccountCapabilitiesParams.USBankAccountACHPayments` |
    | Account | `AccountCapabilities.USBankAccountAchPayments` | `AccountCapabilities.USBankAccountACHPayments` |
    | Charge | `ChargePaymentMethodDetailsTypeAchCreditTransfer` | `ChargePaymentMethodDetailsTypeACHCreditTransfer` |
    | Charge | `ChargePaymentMethodDetailsTypeAchDebit` | `ChargePaymentMethodDetailsTypeACHDebit` |
    | Invoice | `InvoicePaymentSettingsPaymentMethodTypeAchCreditTransfer` | `InvoicePaymentSettingsPaymentMethodTypeACHhCreditTransfer` |
    | Invoice | `InvoicePaymentSettingsPaymentMethodTypeAchDebit` | `InvoicePaymentSettingsPaymentMethodTypeACHDebit` |
    | PaymentMethod | `PaymentMethodUSBankAccountNetworksSupportedAch` | `PaymentMethodUSBankAccountNetworksSupportedACH` |
    | Subscription | `SubscriptionPaymentSettingsPaymentMethodTypeAchCreditTransfer` | `SubscriptionPaymentSettingsPaymentMethodTypeACHCreditTransfer` |
    | Subscription | `SubscriptionPaymentSettingsPaymentMethodTypeAchDebit` | `SubscriptionPaymentSettingsPaymentMethodTypeACHDebit` |
    | Treasury CreditReversal | `TreasuryCreditReversalNetworkAch` | `TreasuryCreditReversalNetworkACH` |
    | Treasury DebitReversal | `TreasuryDebitReversalNetworkAch` | `TreasuryDebitReversalNetworkACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountActiveFeatureInboundTransfersAch` | `TreasuryFinancialAccountActiveFeatureInboundTransfersACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountActiveFeatureOutboundPaymentsAch` | `TreasuryFinancialAccountActiveFeatureOutboundPaymentsACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountActiveFeatureOutboundTransfersAch` | `TreasuryFinancialAccountActiveFeatureOutboundTransfersACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountFinancialAddressSupportedNetworkAch` | `TreasuryFinancialAccountFinancialAddressSupportedNetworkACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountPendingFeatureInboundTransfersAch` | `TreasuryFinancialAccountPendingFeatureInboundTransfersACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountPendingFeatureOutboundPaymentsAch` | `TreasuryFinancialAccountPendingFeatureOutboundPaymentsACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountPendingFeatureOutboundTransfersAch` | `TreasuryFinancialAccountPendingFeatureOutboundTransfersACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountRestrictedFeatureInboundTransfersAch` | `TreasuryFinancialAccountRestrictedFeatureInboundTransfersACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountRestrictedFeatureOutboundPaymentsAch` | `TreasuryFinancialAccountRestrictedFeatureOutboundPaymentsACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountRestrictedFeatureOutboundTransfersAch` | `TreasuryFinancialAccountRestrictedFeatureOutboundTransfersACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountFeaturesInboundTransfersParams.Ach` | `TreasuryFinancialAccountFeaturesInboundTransfersParams.ACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountFeaturesInboundTransfersParams.Ach` | `TreasuryFinancialAccountFeaturesInboundTransfersParams.ACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountFeaturesInboundTransfersAchParams` | `TreasuryFinancialAccountFeaturesInboundTransfersACHParams` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountFeaturesOutboundPaymentsParams.Ach` | `TreasuryFinancialAccountFeaturesOutboundPaymentsParams.ACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchParams` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHParams` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountFeaturesOutboundTransfersParams.Ach` | `TreasuryFinancialAccountFeaturesOutboundTransfersParams.ACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountFeaturesOutboundTransfersAchParams` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHParams` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountUpdateFeaturesInboundTransfersParams.Ach` | `TreasuryFinancialAccountUpdateFeaturesInboundTransfersParams.ACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountUpdateFeaturesInboundTransfersAchParams` | `TreasuryFinancialAccountUpdateFeaturesInboundTransfersACHParams` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountUpdateFeaturesOutboundPaymentsParams.Ach` | `TreasuryFinancialAccountUpdateFeaturesOutboundPaymentsParams.ACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountUpdateFeaturesOutboundPaymentsAchParams` | `TreasuryFinancialAccountUpdateFeaturesOutboundPaymentsACHParams` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountUpdateFeaturesOutboundTransfersParams.Ach` | `TreasuryFinancialAccountUpdateFeaturesOutboundTransfersParams.ACH` |
    | Treasury FinancialAccount | `TreasuryFinancialAccountUpdateFeaturesOutboundTransfersAchParams` | `TreasuryFinancialAccountUpdateFeaturesOutboundTransfersACHParams` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatus` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatus` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusActive` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusActive` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusPending` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusPending` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusRestricted` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusRestricted` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailCode` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailCode` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailCodeActivating` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailCodeActivating` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailCodeCapabilityNotRequested` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailCodeCapabilityNotRequested` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailCodeFinancialAccountClosed` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailCodeFinancialAccountClosed` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailCodeRejectedOther` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailCodeRejectedOther` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailCodeRejectedUnsupportedBusiness` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailCodeRejectedUnsupportedBusiness` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailCodeRequirementsPastDue` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailCodeRequirementsPastDue` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailCodeRequirementsPendingVerification` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailCodeRequirementsPendingVerification` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailCodeRestrictedByPlatform` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailCodeRestrictedByPlatform` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailCodeRestrictedOther` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailCodeRestrictedOther` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailResolution` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailResolution` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailResolutionContactStripe` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailResolutionContactStripe` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailResolutionProvideInformation` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailResolutionProvideInformation` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailResolutionRemoveRestriction` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailResolutionRemoveRestriction` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailRestriction` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailRestriction` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailRestrictionInboundFlows` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailRestrictionInboundFlows` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailRestrictionOutboundFlows` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailRestrictionOutboundFlows` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetailRestrictionOutboundFlows` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetailRestrictionOutboundFlows` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatus` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatus` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusActive` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusActive` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusPending` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusPending` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusRestricted` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusRestricted` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailCode` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailCode` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailCode` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailCodeActivating` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailCode` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailCodeCapabilityNotRequested` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailCode` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailCodeFinancialAccountClosed` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailCode` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailCodeRejectedOther` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailCode` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailCodeRejectedUnsupportedBusiness` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailCode` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailCodeRequirementsPastDue` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailCode` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailCodeRequirementsPendingVerification` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailCode` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailCodeRestrictedByPlatform` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailCode` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailCodeRestrictedOther` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailResolution` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailResolution` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailResolutionContactStripe` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailResolutionContactStripe` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailResolutionProvideInformation` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailResolutionProvideInformation` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailResolutionRemoveRestriction` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailResolutionRemoveRestriction` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailRestriction` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailRestriction` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailRestrictionInboundFlows` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailRestrictionInboundFlows` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetailRestrictionOutboundFlows` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetailRestrictionOutboundFlows` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatus` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatus` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusActive` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusActive` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusPending` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusPending` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusRestricted` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusRestricted` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailCode` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailCode` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailCodeActivating` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailCodeActivating` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailCodeCapabilityNotRequested` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailCodeCapabilityNotRequested` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailCodeFinancialAccountClosed` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailCodeFinancialAccountClosed` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailCodeRejectedOther` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailCodeRejectedOther` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailCodeRejectedUnsupportedBusiness` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailCodeRejectedUnsupportedBusiness` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailCodeRequirementsPastDue` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailCodeRequirementsPastDue` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailCodeRequirementsPendingVerification` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailCodeRequirementsPendingVerification` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailCodeRestrictedByPlatform` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailCodeRestrictedByPlatform` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailCodeRestrictedOther` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailCodeRestrictedOther` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailResolution` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailResolution` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailResolutionContactStripe` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailResolutionContactStripe` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailResolutionProvideInformation` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailResolutionProvideInformation` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailResolutionRemoveRestriction` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailResolutionRemoveRestriction` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailRestriction` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailRestriction` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailRestrictionInboundFlows` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailRestrictionInboundFlows` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetailRestrictionOutboundFlows` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetailRestrictionOutboundFlows` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAchStatusDetail` | `TreasuryFinancialAccountFeaturesInboundTransfersACHStatusDetail` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfers.Ach` | `TreasuryFinancialAccountFeaturesInboundTransfers.ACH` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesInboundTransfersAch` | `TreasuryFinancialAccountFeaturesInboundTransfersACH` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAchStatusDetail` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACHStatusDetail` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPaymentsAch` | `TreasuryFinancialAccountFeaturesOutboundPaymentsACH` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundPayments.Ach` | `TreasuryFinancialAccountFeaturesOutboundPayments.ACH` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAch` | `TreasuryFinancialAccountFeaturesOutboundTransfersACH` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfersAchStatusDetail` | `TreasuryFinancialAccountFeaturesOutboundTransfersACHStatusDetail` |
    | Treasury FinancialAccountFeatures | `TreasuryFinancialAccountFeaturesOutboundTransfers.Ach` | `TreasuryFinancialAccountFeaturesOutboundTransfers.ACH` |
    | Treasury InboundTransfer | `TreasuryInboundTransferOriginPaymentMethodDetailsUSBankAccountNetworkAch` | `TreasuryInboundTransferOriginPaymentMethodDetailsUSBankAccountNetworkACH` |
    | Treasury OutboundPayment | `TreasuryOutboundPaymentOriginPaymentMethodDetailsUSBankAccountNetworkAch` | `TreasuryOutboundPaymentOriginPaymentMethodDetailsUSBankAccountNetworkACH` |
    | Treasury OutboundTransfer | `TreasuryOutboundTransferOriginPaymentMethodDetailsUSBankAccountNetworkAch` | `TreasuryOutboundTransferOriginPaymentMethodDetailsUSBankAccountNetworkACH` |
    | Treasury ReceivedCredit | `TreasuryReceivedCreditNetworkAch` | `TreasuryReceivedCreditNetworkACH` |
    | Treasury ReceivedCredit | `TreasuryReceivedCreditNetworkAch` | `TreasuryReceivedCreditNetworkACH` |

  - `Acss` to `ACSS`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Charge | `ChargePaymentMethodDetails.AcssDebit` | `ChargePaymentMethodDetails.ACSSDebit` |
    | Charge | `ChargePaymentMethodDetailsAcssDebit` | `ChargePaymentMethodDetailsACSSDebit` |
    | Charge | `ChargePaymentMethodDetailsTypeAcssDebit` | `ChargePaymentMethodDetailsTypeACSSDebit` |

  - `Amex` to `AmEx`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Charge | `ChargePaymentMethodDetailsCardWallet.AmexExpressCheckout` | `ChargePaymentMethodDetailsCardWallet.AmExExpressCheckout` |
    | Charge | `ChargePaymentMethodDetailsCardWalletAmexExpressCheckout` | `ChargePaymentMethodDetailsCardWalletAmExExpressCheckout` |
    | PaymentMethod | `PaymentMethodCardBrandAmex` | `PaymentMethodCardBrandAmEx` |
    | PaymentMethod | `PaymentMethodCardWalletTypeAmexExpressCheckout` | `PaymentMethodCardWalletTypeAmExExpressCheckout` |
    | PaymentMethod | `PaymentMethodCardWalletAmexExpressCheckout` | `PaymentMethodCardWalletAmExExpressCheckout` |
    | SetupIntent | `SetupIntentPaymentMethodOptionsCardNetwork.SetupIntentPaymentMethodOptionsCardNetworkAmex` | `SetupIntentPaymentMethodOptionsCardNetwork.SetupIntentPaymentMethodOptionsCardNetworkAmEx` |


  - `Bic` to `BIC`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Charge | `ChargePaymentMethodDetailsBancontact.Bic` | `ChargePaymentMethodDetailsBancontact.BIC` |
    | Charge | `ChargePaymentMethodDetailsGiropay.Bic` | `ChargePaymentMethodDetailsGiropay.BIC` |
    | Charge | `ChargePaymentMethodDetailsIDEAL.Bic` | `ChargePaymentMethodDetailsIDEAL.BIC` |
    | Charge | `ChargePaymentMethodDetailsSEPACreditTransfer.Bic` | `ChargePaymentMethodDetailsSEPACreditTransfer.BIC` |
    | Charge | `ChargePaymentMethodDetailsSofort.Bic` | `ChargePaymentMethodDetailsSofort.BIC` |
    | Charge | `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressIban.Bic` | `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressIBAN.BIC` |
    | FundingInstructions | `FundingInstructionsBankTransferFinancialAddressIBAN.Bic` | `FundingInstructionsBankTransferFinancialAddressIBAN.BIC` |
    | PaymentIntent | `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressIBAN.Bic` | `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressIBAN.BIC` |
    | PaymentMethod | `PaymentMethodIdeal.Bic` | `PaymentMethodIDEAL.BIC` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetailsBancontact.Bic` | `SetupAttemptPaymentMethodDetailsBancontact.BIC` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetailsIdeal.Bic` | `SetupAttemptPaymentMethodDetailsIDEAL.BIC` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetailsSofort.Bic` | `SetupAttemptPaymentMethodDetailsSofort.BIC` |

  - `Eps` to `EPS`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Charge | `ChargePaymentMethodDetailsTypeEps` | `ChargePaymentMethodDetailsTypeEPS` |

  - `FEDEX` to `FedEx`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Issuing Card | `IssuingCardShippingCarrierFEDEX` | `IssuingCardShippingCarrierFedEx` |

  - `Iban` to `IBAN`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | FundingInstructions | `FundingInstructionsBankTransferFinancialAddress.Iban` | `FundingInstructionsBankTransferFinancialAddress.IBAN` |
    | FundingInstructions | `FundingInstructionsBankTransferFinancialAddressIban` | `FundingInstructionsBankTransferFinancialAddressIBAN` |
    | FundingInstructions | `FundingInstructionsBankTransferFinancialAddressIban.Iban` | `FundingInstructionsBankTransferFinancialAddressIBAN.IBAN` |
    | FundingInstructions | `FundingInstructionsBankTransferFinancialAddressTypeIban` | `FundingInstructionsBankTransferFinancialAddressTypeIBAN` |
    | PaymentIntent | `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressTypeIban` | `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressTypeIBAN` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeIban` | `PaymentIntentPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeIBAN` |
    | PaymentIntent | `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddress.Iban` | `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddress.IBAN` |
    | PaymentIntent | `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressIban` | `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressIBAN` |
    | PaymentIntent | `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressIban.Iban` | `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressIBAN.IBAN` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetailsBancontact.IbanLast4` | `SetupAttemptPaymentMethodDetailsBancontact.IBANLast4` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetailsIdeal.IbanLast4` | `SetupAttemptPaymentMethodDetailsIDEAL.IBANLast4` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetailsSofort.IbanLast4` | `SetupAttemptPaymentMethodDetailsSofort.IBANLast4` |
    | SetupIntent | `SetupIntentPaymentMethodDataSepaDebitParams.Iban` | `SetupIntentPaymentMethodDataSEPADebitParams.IBAN` |
    | SetupIntent | `SetupIntentConfirmPaymentMethodDataSepaDebitParams.Iban` | `SetupIntentConfirmPaymentMethodDataSEPADebitParams.IBAN` |

  - `Ideal` to `IDEAL`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Account | `AccountCapabilitiesIdealPaymentsParams` | `AccountCapabilitiesIDEALPaymentsParams` |
    | Account | `AccountCapabilitiesParams.IdealPayments` | `AccountCapabilitiesParams.IDEALPayments` |
    | Account | `AccountCapabilities.IdealPayments` | `AccountCapabilities.IDEALPayments` |
    | Charge | `ChargePaymentMethodDetailsTypeIdeal` | `ChargePaymentMethodDetailsTypeIDEAL` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptions.Ideal` | `CheckoutSessionPaymentMethodOptions.IDEAL` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsIdeal` | `CheckoutSessionPaymentMethodOptionsIDEAL` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsIdealSetupFutureUsage` | `CheckoutSessionPaymentMethodOptionsIDEALSetupFutureUsage` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsIdealSetupFutureUsageNone` | `CheckoutSessionPaymentMethodOptionsIDEALSetupFutureUsageNone` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsIdealParams` | `CheckoutSessionPaymentMethodOptionsIDEALParams` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsParams.Ideal` | `CheckoutSessionPaymentMethodOptionsParams.IDEAL` |
    | Invoice | `InvoicePaymentSettingsPaymentMethodTypeIdeal` | `InvoicePaymentSettingsPaymentMethodTypeIDEAL` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptions.Ideal` | `PaymentIntentPaymentMethodOptions.IDEAL` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsIdeal` | `PaymentIntentPaymentMethodOptionsIDEAL` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsIdealSetupFutureUsage` | `PaymentIntentPaymentMethodOptionsIDEALSetupFutureUsage` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsIdealSetupFutureUsageNone` | `PaymentIntentPaymentMethodOptionsIDEALSetupFutureUsageNone` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsIdealSetupFutureUsageOffSession` | `PaymentIntentPaymentMethodOptionsIDEALSetupFutureUsageOffSession` |
    | PaymentIntent | `PaymentIntentPaymentMethodDataParams.Ideal` | `PaymentIntentPaymentMethodDataParams.IDEAL` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsParams.Ideal` | `PaymentIntentPaymentMethodOptionsParams.IDEAL` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsIdealParams` | `PaymentIntentPaymentMethodOptionsIDEALParams` |
    | PaymentLink | `PaymentLinkPaymentMethodTypeIdeal` | `PaymentLinkPaymentMethodTypeIDEAL` |
    | PaymentMethod | `PaymentMethodTypeIdeal` | `PaymentMethodTypeIDEAL` |
    | PaymentMethod | `PaymentMethodParams.Ideal` | `PaymentMethodParams.IDEAL` |
    | PaymentMethod | `PaymentMethod.Ideal` | `PaymentMethod.IDEAL` |
    | PaymentMethod | `PaymentMethodIdeal` | `PaymentMethodIDEAL` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetailsIdeal` | `SetupAttemptPaymentMethodDetailsIDEAL` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetails.Ideal` | `SetupAttemptPaymentMethodDetails.IDEAL` |
    | SetupIntent | `SetupIntentPaymentMethodDataIdealParams` | `SetupIntentPaymentMethodDataIDEALParams` |
    | SetupIntent | `SetupIntentPaymentMethodDataParams.Ideal` | `SetupIntentPaymentMethodDataParams.IDEAL` |
    | SetupIntent | `SetupIntentConfirmPaymentMethodDataParams.Ideal` | `SetupIntentConfirmPaymentMethodDataParams.IDEAL` |
    | Subscription | `SubscriptionPaymentSettingsPaymentMethodTypeIdeal` | `SubscriptionPaymentSettingsPaymentMethodTypeIDEAL` |

  - `Sepa` to `SEPA`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Charge | `ChargePaymentMethodDetailsTypeSepaDebit` | `ChargePaymentMethodDetailsTypeSEPAaDebit` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptions.SepaDebit` | `CheckoutSessionPaymentMethodOptions.SEPADebit` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsSepaDebit` | `CheckoutSessionPaymentMethodOptionsSEPADebit` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsage` | `CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsage` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsageNone` | `CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsageNone` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsageOffSession` | `CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsageOffSession` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsSepaDebitSetupFutureUsageOnSession` | `CheckoutSessionPaymentMethodOptionsSEPADebitSetupFutureUsageOnSession` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsSepaDebitParams` | `CheckoutSessionPaymentMethodOptionsSEPADebitParams` |
    | CheckoutSessions | `CheckoutSessionPaymentMethodOptionsParams.SepaDebit` | `CheckoutSessionPaymentMethodOptionsParams.SEPADebit` |
    | FundingInstructions | `FundingInstructionsBankTransferFinancialAddressSupportedNetworkSepa` | `FundingInstructionsBankTransferFinancialAddressSupportedNetworkSEPA` |
    | Invoice | `InvoicePaymentSettingsPaymentMethodTypeSepaCreditTransfer` | `InvoicePaymentSettingsPaymentMethodTypeSEPACreditTransfer` |
    | Invoice | `InvoicePaymentSettingsPaymentMethodTypeSepaDebit` | `InvoicePaymentSettingsPaymentMethodTypeSEPADebit` |
    | Mandate | `MandatePaymentMethodDetails.SepaDebit` | `MandatePaymentMethodDetails.SEPADebit` |
    | Mandate | `MandatePaymentMethodDetailsSepaDebit` | `MandatePaymentMethodDetailsSEPADebit` |
    | PaymentIntent | `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressSupportedNetworkSepa` | `PaymentIntentNextActionDisplayBankTransferInstructionsFinancialAddressSupportedNetworkSEPA` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeSepa` | `PaymentIntentPaymentMethodOptionsCustomerBalanceBankTransferRequestedAddressTypeSEPA` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsSepaDebitSetupFutureUsage` | `PaymentIntentPaymentMethodOptionsSepaDebitSetupFutureUsage` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsSepaDebitSetupFutureUsageNone` | `PaymentIntentPaymentMethodOptionsSEPADebitSetupFutureUsageNone` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsSepaDebitSetupFutureUsageNone` | `PaymentIntentPaymentMethodOptionsSEPADebitSetupFutureUsageNone` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsSepaDebitSetupFutureUsageOffSession` | `PaymentIntentPaymentMethodOptionsSepaDebitSetupFutureUsageOffSession` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsSepaDebitSetupFutureUsageOnSession` | `PaymentIntentPaymentMethodOptionsSEPADebitSetupFutureUsageOnSession` |
    | PaymentIntent | `PaymentIntentPaymentMethodDataParams.SepaDebit` | `PaymentIntentPaymentMethodDataParams.SEPADebit` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsSepaDebitMandateOptionsParams` | `PaymentIntentPaymentMethodOptionsSEPADebitMandateOptionsParams` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsParams.SepaDebit` | `PaymentIntentPaymentMethodOptionsParams.SEPADebit` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsSepaDebitParams` | `PaymentIntentPaymentMethodOptionsSEPADebitParams` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptions.SepaDebit` | `PaymentIntentPaymentMethodOptions.SEPADebit` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsSepaDebit` | `PaymentIntentPaymentMethodOptionsSEPADebit` |
    | PaymentIntent | `PaymentIntentPaymentMethodOptionsSepaDebitMandateOptions` | `PaymentIntentPaymentMethodOptionsSEPADebitMandateOptions` |
    | PaymentLink | `PaymentLinkPaymentMethodTypeSepaDebit` | `PaymentLinkPaymentMethodTypeSEPADebit` |
    | PaymentMethod | `PaymentMethodTypeSepaDebit` | `PaymentMethodTypeSEPADebit` |
    | PaymentMethod | `PaymentMethodParams.SepaDebit` | `PaymentMethodParams.SEPADebit` |
    | PaymentMethod | `PaymentMethod.SepaDebit` | `PaymentMethod.SEPADebit` |
    | PaymentMethod | `PaymentMethodSepaDebit` | `PaymentMethodSEPADebit` |
    | PaymentMethod | `PaymentMethodSepaDebitGeneratedFrom` | `PaymentMethodSEPADebitGeneratedFrom` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetailsBancontact.GeneratedSepaDebit` | `SetupAttemptPaymentMethodDetailsBancontact.GeneratedSEPADebit` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetailsBancontact.GeneratedSepaDebitMandate` | `SetupAttemptPaymentMethodDetailsBancontact.GeneratedSEPADebitMandate` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetailsIdeal.GeneratedSepaDebit` | `SetupAttemptPaymentMethodDetailsIDEAL.GeneratedSEPADebit` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetailsIdeal.GeneratedSepaDebitMandate` | `SetupAttemptPaymentMethodDetailsIDEAL.GeneratedSEPADebitMandate` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetailsSepaDebit` | `SetupAttemptPaymentMethodDetailsSEPADebit` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetailsSofort.GeneratedSepaDebit` | `SetupAttemptPaymentMethodDetailsSofort.GeneratedSEPADebit` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetailsSofort.GeneratedSepaDebitMandate` | `SetupAttemptPaymentMethodDetailsSofort.GeneratedSEPADebitMandate` |
    | SetupAttempt | `SetupAttemptPaymentMethodDetails.SepaDebit` | `SetupAttemptPaymentMethodDetails.SEPADebit` |
    | SetupIntent | `SetupIntentPaymentMethodDataSepaDebitParams` | `SetupIntentPaymentMethodDataSEPADebitParams` |
    | SetupIntent | `SetupIntentPaymentMethodDataParams.SepaDebit` | `SetupIntentPaymentMethodDataParams.SEPADebit` |
    | SetupIntent | `SetupIntentPaymentMethodOptionsSepaDebitMandateOptionsParams` | `SetupIntentPaymentMethodOptionsSEPADebitMandateOptionsParams` |
    | SetupIntent | `SetupIntentPaymentMethodOptionsUSBankAccountParams.SepaDebit` | `SetupIntentPaymentMethodOptionsUSBankAccountParams.SEPADebit` |
    | SetupIntent | `SetupIntentConfirmPaymentMethodDataSepaDebitParams` | `SetupIntentConfirmPaymentMethodDataSEPADebitParams` |
    | SetupIntent | `SetupIntentConfirmPaymentMethodDataParams.SepaDebit` | `SetupIntentConfirmPaymentMethodDataParams.SEPADebit` |
    | SetupIntent | `SetupIntentPaymentMethodOptionsSepaDebitMandateOptions` | `SetupIntentPaymentMethodOptionsSEPADebitMandateOptions` |
    | SetupIntent | `SetupIntentPaymentMethodOptionsSepaDebit` | `SetupIntentPaymentMethodOptionsSEPADebit` |
    | SetupIntent | `SetupIntentPaymentMethodOptionsUSBankAccount.SepaDebit` | `SetupIntentPaymentMethodOptionsUSBankAccount.SEPADebit` |
    | Subscription | `SubscriptionPaymentSettingsPaymentMethodTypeSepaCreditTransfer` | `SubscriptionPaymentSettingsPaymentMethodTypeSEPACreditTransfer` |
    | Subscription | `SubscriptionPaymentSettingsPaymentMethodTypeSepaDebit` | `SubscriptionPaymentSettingsPaymentMethodTypeSEPADebit` |

  - `ExternalAccount` to `AccountExternalAccount`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Account | `ExternalAccount` | `AccountExternalAccount` |
    | Account | `ExternalAccountType` | `AccountExternalAccountType` |
    | Account | `ExternalAccountTypeBankAccount` | `AccountExternalAccountTypeBankAccount` |
    | Account | `ExternalAccountTypeCard` | `AccountExternalAccountTypeCard` |
    | Account | `AccountExternalAccountParams` | `AccountExternalAccountParams` |
    | Account | `ExternalAccountList` | `AccountExternalAccountList` | <!-- ---- R, S -->

  - `InvoiceLine` to `InvoiceLineItem`
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Invoice | `InvoiceLine` | `InvoiceLineItem` |
    | Invoice | `InvoiceLineList` | `InvoiceLineItemList` |
    | Invoice | `InvoiceLineType` | `InvoiceLineItemType` |
    | Invoice | `InvoiceLineTypeInvoiceItem` | `InvoiceLineItemTypeInvoiceItem` |
    | Invoice | `InvoiceLineTypeSubscription` | `InvoiceLinItemeTypeSubscription` |
    | Invoice | `InvoiceLineDiscountAmount` | `InvoiceLineItemDiscountAmount` |
    | Invoice | `InvoiceLineProrationDetails` | `InvoiceLineItemProrationDetails` |
    | Invoice | `InvoiceLineProrationDetailsCreditedItems` | `InvoiceLineItemProrationDetailsCreditedItems` | <!-- ---- R, S -->

  - `Person` structs/enums to use `Person` prefix
    | Old name | New name |
    | --- | --- |
    | `VerificationDocumentDetailsCode` | `PersonVerificationDocumentDetailsCode` |
    | `VerificationDocumentDetailsCodeDocumentCorrupt` | `PersonVerificationDocumentDetailsCodeDocumentCorrupt` |
    | `VerificationDocumentDetailsCodeDocumentFailedCopy` | `PersonVerificationDocumentDetailsCodeDocumentFailedCopy` |
    | `VerificationDocumentDetailsCodeDocumentFailedGreyscale` | `PersonVerificationDocumentDetailsCodeDocumentFailedGreyscale` |
    | `VerificationDocumentDetailsCodeDocumentFailedOther` | `PersonVerificationDocumentDetailsCodeDocumentFailedOther` |
    | `VerificationDocumentDetailsCodeDocumentFailedTestMode` | `PersonVerificationDocumentDetailsCodeDocumentFailedTestMode` |
    | `VerificationDocumentDetailsCodeDocumentFraudulent` | `PersonVerificationDocumentDetailsCodeDocumentFraudulent` |
    | `VerificationDocumentDetailsCodeDocumentIDTypeNotSupported` | `PersonVerificationDocumentDetailsCodeDocumentIDTypeNotSupported` |
    | `VerificationDocumentDetailsCodeDocumentIDCountryNotSupported` | `PersonVerificationDocumentDetailsCodeDocumentIDCountryNotSupported` |
    | `VerificationDocumentDetailsCodeDocumentManipulated` | `PersonVerificationDocumentDetailsCodeDocumentManipulated` |
    | `VerificationDocumentDetailsCodeDocumentMissingBack` | `PersonVerificationDocumentDetailsCodeDocumentMissingBack` |
    | `VerificationDocumentDetailsCodeDocumentMissingFront` | `PersonVerificationDocumentDetailsCodeDocumentMissingFront` |
    | `VerificationDocumentDetailsCodeDocumentNotReadable` | `PersonVerificationDocumentDetailsCodeDocumentNotReadable` |
    | `VerificationDocumentDetailsCodeDocumentNotUploaded` | `PersonVerificationDocumentDetailsCodeDocumentNotUploaded` |
    | `VerificationDocumentDetailsCodeDocumentTooLarge` | `PersonVerificationDocumentDetailsCodeDocumentTooLarge` |
    | `IdentityVerificationStatus` | `PersonVerificationStatus` |
    | `IdentityVerificationStatusPending` | `PersonVerificationStatusPending` |
    | `IdentityVerificationStatusUnverified` | `PersonVerificationStatusUnverified` |
    | `IdentityVerificationStatusVerified` | `PersonVerificationStatusVerified` |
    | `Relationship` | `PersonRelationship` |
    | `RelationshipParams` | `PersonRelationshipParams` |
    | `RelationshipListParams` | `PersonListRelationshipParams` |
    | `DOBParams` | `PersonDOBParams` |
    | `DocumentsParams` | `PersonDocumentsParams` |
    | `DocumentsCompanyAuthorizationParams` | `PersonDocumentsCompanyAuthorizationParams` |
    | `DocumentsPassportParams` | `PersonDocumentsPassportParams` |
    | `DocumentsVisaParams` | `PersonDocumentsVisaParams` |
    | `Requirements` | `PersonRequirements` |

  - Other
    | Resource | Old name | New name |
    | --- | --- | --- |
    | Account | `PayoutInterval` | `AccountSettingsPayoutsScheduleInterval` |
    | Account | `PayoutIntervalDaily` | `AccountSettingsPayoutsScheduleIntervalDaily` |
    | Account | `PayoutIntervalManual` | `AccountSettingsPayoutsScheduleIntervalManual` |
    | Account | `PayoutIntervalMonthly` | `AccountSettingsPayoutsScheduleIntervalMonthly` |
    | Account | `PayoutIntervalWeekly` | `AccountSettingsPayoutsScheduleIntervalWeekly` |
    | Account | `AccountDocumentsParams.CompanyMemorandumOfAssocation` | `AccountDocumentsParams.CompanyMemorandumOfAssociation` |
    | Account | `AccountDeclineSettingsParams` | `AccountSettingsCardPaymentsDeclineOnParams` |
    | Account | `PayoutScheduleParams` | `AccountSettingsPayoutsScheduleParams` |
    | Account | `AccountPayoutSchedule` | `AccountSettingsPayoutsSchedule` |
    | Balance | `Amount.Value` | `Amount.Amount` |
    | Balance | `BalanceDetails` | `BalanceIssuing` |
    | BalanceTransaction | `BalanceTransactionSourceTypeReversal` | `BalanceTransactionSourceTypeTransferReversal` (this is a fix, the previous value of the enum was incorrect) |
    | BalanceTransaction | `BalanceTransactionTypeIssuingAuthorizationDispute` | `BalanceTransactionTypeIssuingDispute` |
    | BalanceTransaction | `BalanceTransactionTypeIssuingAuthorizationTransaction` | `BalanceTransactionTypeIssuingTransaction` |
    | BalanceTransaction | `BalanceTransactionFee` | `BalanceTransactionFeeDetail` |
    | BalanceTransaction | `BalanceTransactionSource.Reversal` | `BalanceTransactionSource.TransferReversal` |
    | Card | `CardBrandAmex` | `CardBrandAmericanExpress` |
    | Card | `TokenizationMethodAndroidPay` | `CardTokenizationMethodAndroidPay` |
    | Card | `TokenizationMethodApplePay` | `CardTokenizationMethodApplePay` |
    | Charge | `DestinationParams` | `ChargeDestinationParams` |
    | Charge | `ChargeLevel3LineItemsParams` | `ChargeLevel3LineItemParams` |
    | Charge | `FraudDetailsParams` | `ChargeFraudDetailsParams` |
    | Charge | `CaptureParams` | `ChargeCaptureParams` |
    | Charge | `ChargeTransferDataParams` | `ChargeCaptureTransferDataParams` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDs` | `CheckoutSessionCustomerDetailsTaxID` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsType` | `CheckoutSessionCustomerDetailsTaxIDType` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeAETRN` | `CheckoutSessionCustomerDetailsTaxIDTypeAETRN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeAUABN` | `CheckoutSessionCustomerDetailsTaxIDTypeAUABN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeAUARN` | `CheckoutSessionCustomerDetailsTaxIDTypeAUARN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeBGUIC` | `CheckoutSessionCustomerDetailsTaxIDTypeBGUIC` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeBRCNPJ` | `CheckoutSessionCustomerDetailsTaxIDTypeBRCNPJ` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeBRCPF` | `CheckoutSessionCustomerDetailsTaxIDTypeBRCPF` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCABN` | `CheckoutSessionCustomerDetailsTaxIDTypeCABN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCAGSTHST` | `CheckoutSessionCustomerDetailsTaxIDTypeCAGSTHST` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCAPSTBC` | `CheckoutSessionCustomerDetailsTaxIDTypeCAPSTBC` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCAPSTMB` | `CheckoutSessionCustomerDetailsTaxIDTypeCAPSTMB` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCAPSTSK` | `CheckoutSessionCustomerDetailsTaxIDTypeCAPSTSK` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCAQST` | `CheckoutSessionCustomerDetailsTaxIDTypeCAQST` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCHVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeCHVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeCLTIN` | `CheckoutSessionCustomerDetailsTaxIDTypeCLTIN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeESCIF` | `CheckoutSessionCustomerDetailsTaxIDTypeESCIF` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeEUOSSVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeEUOSSVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeEUVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeEUVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeGBVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeGBVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeGEVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeGEVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeHKBR` | `CheckoutSessionCustomerDetailsTaxIDTypeHKBR` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeHUTIN` | `CheckoutSessionCustomerDetailsTaxIDTypeHUTIN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeIDNPWP` | `CheckoutSessionCustomerDetailsTaxIDTypeIDNPWP` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeILVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeILVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeINGST` | `CheckoutSessionCustomerDetailsTaxIDTypeINGST` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeISVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeISVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeJPCN` | `CheckoutSessionCustomerDetailsTaxIDTypeJPCN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeJPRN` | `CheckoutSessionCustomerDetailsTaxIDTypeJPRN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeKRBRN` | `CheckoutSessionCustomerDetailsTaxIDTypeKRBRN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeLIUID` | `CheckoutSessionCustomerDetailsTaxIDTypeLIUID` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeMXRFC` | `CheckoutSessionCustomerDetailsTaxIDTypeMXRFC` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeMYFRP` | `CheckoutSessionCustomerDetailsTaxIDTypeMYFRP` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeMYITN` | `CheckoutSessionCustomerDetailsTaxIDTypeMYITN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeMYSST` | `CheckoutSessionCustomerDetailsTaxIDTypeMYSST` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeNOVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeNOVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeNZGST` | `CheckoutSessionCustomerDetailsTaxIDTypeNZGST` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeRUINN` | `CheckoutSessionCustomerDetailsTaxIDTypeRUINN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeRUKPP` | `CheckoutSessionCustomerDetailsTaxIDTypeRUKPP` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeSAVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeSAVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeSGGST` | `CheckoutSessionCustomerDetailsTaxIDTypeSGGST` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeSGUEN` | `CheckoutSessionCustomerDetailsTaxIDTypeSGUEN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeSITIN` | `CheckoutSessionCustomerDetailsTaxIDTypeSITIN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeTHVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeTHVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeTWVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeTWVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeUAVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeUAVAT` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeUnknown` | `CheckoutSessionCustomerDetailsTaxIDTypeUnknown` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeUSEIN` | `CheckoutSessionCustomerDetailsTaxIDTypeUSEIN` |
    | CheckoutSessions | `CheckoutSessionCustomerDetailsTaxIDsTypeZAVAT` | `CheckoutSessionCustomerDetailsTaxIDTypeZAVAT` |
    | CheckoutSessions | `CheckoutSessionSubscriptionDataItemsParams` | `CheckoutSessionSubscriptionDataItemParams` |
    | CreditNote | `CreditNoteVoidParams` | `CreditNoteVoidCreditNoteParams` |
    | CreditNote | `CreditNoteLineItemListPreviewParams` | `CreditNotePreviewLinesParams` |
    | CreditNote | `CreditNoteLineItemListParams` | `CreditNoteListLinesParams` |
    | Customer | `CustomerInvoiceCustomField` | `CustomerInvoiceSettingsCustomField`|
    | Customer | `CustomerInvoiceCustomFieldParams` | `Customer dInvoiceSettingsCustomFieldParams`|
    | Customer | `CustomerShippingDetails` | `CustomerShipping` |
    | Customer | `CustomerShippingDetailsParams` | `CustomerShippingParams` |
    | Dispute | `EvidenceDetails` | `DisputeEvidenceDetails` |
    | FeeRefund | `FeeRefundListParams.ApplicationFee` | `FeeRefundListParams.ID` |
    | Invoice | `InvoiceUpcomingCustomerDetailsShippingParams` | `InvoiceUpcomingLinesCustomerDetailsShippingParams` |
    | Invoice | `InvoiceUpcomingCustomerDetailsTaxParams` | `InvoiceUpcomingLinesCustomerDetailsTaxParams` |
    | Invoice | `InvoiceUpcomingCustomerDetailsTaxIDParams` | `InvoiceUpcomingLinesCustomerDetailsTaxIDParams` |
    | Invoice | `InvoiceUpcomingCustomerDetailsParams` | `InvoiceUpcomingLinesCustomerDetailsParams` |
    | Invoice | `InvoiceSendParams` | `InvoiceSendInvoiceParams` |
    | Invoice | `InvoiceVoidParams` | `InvoiceVoidInvoiceParams` |
    | Invoice | `InvoiceLineList` | `InvoiceLineItemList` |
    | Invoice | `InvoiceLineListParams` | `InvoiceListLinesParams` |
    | Invoice | `InvoiceLineListParams.ID` | `InvoiceListLinesParams.Invoice` |
    | Invoice | `InvoiceDiscountAmount` | `InvoiceTotalDiscountAmount` |
    | Invoice | `InvoiceTaxAmount` | `InvoiceTotalTaxAmount` |
    | Invoice | `Invoice.ThreasholdReason` | `Invoice.ThresholdReason` (closes #1244) |
    | Invoice | `GetNext` | `Upcoming` |
    | Issuing | `IssuingAuthorizationWalletType` | `IssuingAuthorizationWallet` |
    | Issuing | `IssuingAuthorizationWalletTypeApplePay` | `IssuingAuthorizationWalletApplePay` |
    | Issuing | `IssuingAuthorizationWalletTypeGooglePay` | `IssuingAuthorizationWalletGooglePay` |
    | Issuing | `IssuingAuthorizationWalletTypeSamsungPlay` | `IssuingAuthorizationWalletSamsungPlay` |
    | PaymentSource | `PaymentSourceTypeObject` | `PaymentSourceTypeSource` |
    | PaymentSource | `SourceParams` | `PaymentSourceSourceParams` |
    | PaymentSource | `SourceList` | `PaymentSourceList` |
    | PaymentSource | `SourceListParams` | `PaymentSourceListParams` |
    | PaymentSource | `CustomerSourceParams` | `PaymentSourceParams` |
    | PaymentSource | `SourceVerifyParams` | `PaymentSourceVerifyParams` |
    | PaymentSource | `PaymentSource.SourceObject` | `PaymentSource.Source` |
    | Price | `PriceRecurringListParams` | `PriceListRecurringParams` |
    | Product | `PackageDimensions` | `ProductPackageDimensions` |
    | Product | `PackageDimensionsParams` | `ProductPackageDimensionsParams` |
    | Radar | `RadarValueListItem.RadarValueList` | `RadarValueListItem.ValueList` |
    | Radar | `RadarValueListItemListParams.RadarValueList` | `RadarValueListItemListParams.ValueList` |
    | Radar | `RadarValueListItemParams.RadarValueList` | `RadarValueListItemParams.ValueList` |
    | Radar.EarlyFraudWarning | `RadarEarlyFraudWarningList.Values` | `RadarEarlyFraudWarningList.Data` |
    | Reporting | `ReportRun` | `ReportingReportRun ` |
    | Reporting | `ReportRunList` | `ReportingReportRunList ` |
    | Reporting | `ReportRunListParams` | `ReportingReportRun ` |
    | Reporting | `ReportRunParameters` | `ReportingReportRunParameters` |
    | Reporting | `ReportRunParametersParams` | `ReportingReportRunParametersParams` |
    | Reporting | `ReportRunParams` | `ReportingReportRunParams` |
    | Reporting | `ReportRunStatus` | `ReportingReportRunStatus` |
    | Reporting | `ReportRunStatusFailed` | `ReportingReportRunStatusFailed` |
    | Reporting | `ReportRunStatusPending` | `ReportingReportRunStatusPending` |
    | Reporting | `ReportRunStatusSucceeded` | `ReportingReportRunStatusSucceeded` |
    | Reporting | `ReportType` | `ReportingReportType` |
    | Reporting | `ReportTypeListParams` | `ReportingReportTypeListParams` |
    | Reporting | `ReportTypeListParams` | `ReportingReportTypeListParams` |
    | Reporting | `ReportTypeParams` | `ReportingReportTypeParams` |
    | Reporting | `ReportTypeParams` | `ReportingReportTypeParams` |
    | Reversal | `reversal` | `transferreversal` |
    | Reversal | `ReversalParams` | `TransferReversalParams` |
    | Reversal | `ReversalListParams` | `TransferReversalListParams` |
    | Reversal | `Reversal` | `TransferReversal` |
    | Reversal | `ReversalList` | `TransferReversalList` |
    | Sku | `InventoryParams` | `SKUInventoryParams` |
    | Sku | `Inventory` | `SKUInventory` |
    | Source | `SourceCodeVerificationFlowStatus` | `SourceCodeVerificationStatus` |
    | Source | `SourceCodeVerificationFlowStatusFailed` | `SourceCodeVerificationStatusFailed` |
    | Source | `SourceCodeVerificationFlowStatusPending` | `SourceCodeVerificationStatusPending` |
    | Source | `SourceCodeVerificationFlowStatusSucceeded` | `SourceCodeVerificationStatusSucceeded` |
    | Source | `SourceRefundAttributesMethod` | `SourceReceiverRefundAttributesMethod` |
    | Source | `SourceRefundAttributesMethodEmail` | `SourceReceiverRefundAttributesMethodEmail` |
    | Source | `SourceRefundAttributesMethodManual` | `SourceReceiverRefundAttributesMethodManual` |
    | Source | `SourceRefundAttributesStatus` | `SourceReceiverRefundAttributesStatus` |
    | Source | `SourceRefundAttributesStatusAvailable` | `SourceReceiverRefundAttributesStatusAvailable` |
    | Source | `SourceRefundAttributesStatusMissing` | `SourceReceiverRefundAttributesStatusMissing` |
    | Source | `SourceRefundAttributesStatusRequested` | `SourceReceiverRefundAttributesStatusRequested` |
    | Source | `SourceRedirectFlowFailureReason` | `SourceRedirectFailureReason` |
    | Source | `SourceRedirectFlowFailureReasonDeclined` | `SourceRedirectFailureReasonDeclined` |
    | Source | `SourceRedirectFlowFailureReasonProcessingError` | `SourceRedirectFailureReasonProcessingError` |
    | Source | `SourceRedirectFlowFailureReasonUserAbort` | `SourceRedirectFailureReasonUserAbort` |
    | Source | `SourceRedirectFlowStatus` | `SourceRedirectStatus` |
    | Source | `SourceRedirectFlowStatusFailed` | `SourceRedirectStatusFailed` |
    | Source | `SourceRedirectFlowStatusNotRequired` | `SourceRedirectStatusNotRequired` |
    | Source | `SourceRedirectFlowStatusPending` | `SourceRedirectStatusPending` |
    | Source | `SourceRedirectFlowStatusSucceeded` | `SourceRedirectStatusSucceeded` |
    | Source | `SourceObjectDetachParams` | `SourceDetachParams` |
    | Source | `SourceObjectParams` | `SourceParams` |
    | Source | `SourceOrderItemsParams` | `ourceSourceOrderItemParams` |
    | Source | `SourceOrderParams` | `SourceSourceOrderParams` |
    | Source | `RedirectParams` | `SourceRedirectParams` |
    | Source | `CodeVerificationFlow` | `SourceCodeVerification` |
    | Source | `SourceSourceOrderItems` | `SourceSourceOrderItem` |
    | SubscriptionSchedule | `SubscriptionSchedulePhaseBillingCycleAnchor` | `SubscriptionScheduleDefaultSettingsBillingCycleAnchor` |
    | SubscriptionSchedule | `SubscriptionSchedulePhaseBillingCycleAnchorAutomatic` | `SubscriptionScheduleDefaultSettingsBillingCycleAnchorAutomatic` |
    | SubscriptionSchedule | `SubscriptionSchedulePhaseBillingCycleAnchorPhaseStart` | `SubscriptionScheduleDefaultSettingsBillingCycleAnchorPhaseStart` |
    | SubscriptionSchedule | `SubscriptionScheduleInvoiceSettingsParams` | `SubscriptionScheduleDefaultSettingsInvoiceSettingsParams` |
    | SubscriptionSchedule | `SubscriptionScheduleInvoiceSettings` | `SubscriptionScheduleDefaultSettingsInvoiceSettings` |
    | TaxRate | `TaxRateTaxTypeJct` | `TaxRateTaxTypeJCT` |
    | Token | `PIIParams` | `TokenPIIParams` |

- Replace `AccountAddressParams` with `AccountCompanyAddressKanaParams` and `AccountCompanyAddressKanjiParams`
- Change type:
  | Resource | Target | Old type | New type | Rationale |
  | --- | --- | --- | --- | --- |
  | Account | `AccountCapabilities.AffirmPayments` | `AccountCapabilitiesAffirmPayments` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCapabilities.BankTransferPayments` | `AccountCapabilitiesBankTransferPayments` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCapabilities.KonbiniPayments` | `AccountCapabilitiesKonbiniPayments` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCapabilities.LinkPayments` | `AccountCapabilitiesLinkPayments` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCapabilities.PayNowPayments` | `AccountCapabilitiesPayNowPayments` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCapabilities.PromptPayPayments` | `AccountCapabilitiesPromptPayPayments` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCapabilities.Treasury` | `AccountCapabilitiesTreasury` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCapabilities.USBankAccountAchPayments` | `AccountCapabilitiesUSBankAccountAchPayments` | `AccountCapabilityStatus` | Be consistent about using `AccountCapabilityStatus` for all of `AccountCapabilities` (all the same enum values) |
  | Account | `AccountCompanyParams.Address` | `*AccountAddressParams` | `*AddressParams` | `*AccountAddressParams` had `Town` field that is not applicable to `AccountCompanyParams.Address` |
  | Account | `AccountCompanyParams.AddressKana` | `*AccountAddressParams` | `*AccountCompanyAddressKanaParams` | `AccountAddressParams` has been split into separate `AccountCompanyAddressKanaParams` and `AccountCompanyAddressKanjiParams` |
  | Account | `AccountCompanyParams.AddressKanji` | `*AccountAddressParams` | `*AccountCompanyAddressKanjiParams` | `AccountAddressParams` has been split into separate `AccountCompanyAddressKanaParams` and `AccountCompanyAddressKanjiParams` |
  | Account | `AccountSettingsCardIssuingParams.TOSAcceptance` | `*AccountTOSAcceptanceParams` | `*AccountSettingsCardIssuingTOSAcceptanceParams` | `AccountTOSAcceptanceParams` has `ServiceAgreement` field that is not applicable to `AccountSettingsCardIssuingParams.TOSAcceptance` |
  | ApplicationFee | `ApplicationFee.Application` | `string` | `*Application` | `Application` is expandable |
  | BalanceTransaction | `BalanceTransactionSource.IssuingTransaction` | `*IssuingAuthorization` | `*IssuingTransaction` | `ApplicationFee.Application` is expandable |
  | BillingPortalConfiguration | `BillingPortalConfiguration.Application` | `string` | `*Application` | `Application` is expandable |
  | Card | `Card.AddressLine1Check` | `CardVerification` | `CardAddressLine1Check` | Same values. `CardVerification` was replaced with separate definitions for `AddressLine1Check`, `AddressZipCheck`, and `CVCCheck` |
  | Card | `Card.AddressZipCheck` | `CardVerification` | `CardAddressZipCheck` | Same values. `CardVerification` was replaced with separate definitions for `AddressLine1Check`, `AddressZipCheck`, and `CVCCheck` |
  | Card | `Card.CVCCheck` | `CardVerification` | `CardCVCCheck` | Same values. `CardVerification` was replaced with separate definitions for `AddressLine1Check`, `AddressZipCheck`, and `CVCCheck` |
  | Card | `Card.ExpMonth` | `uint8` | `int64` | Be consistent with `int` types in library |
  | Card | `Card.ExpYear` | `uint16` | `int64` | Be consistent with `int` types in library |
  | Charge | `Charge.BillingDetails` | `*BillingDetails` | `ChargeBillingDetails` | Unshared with PaymentMethod |
  | Charge | `Charge.Level3` | `ChargeLevel3` | `*ChargeLevel3` | `ChargeLevel3` is expandable |
  | Charge | `ChargePaymentMethodDetailsCardChecks.AddressLine1Check` | `CardVerification` | `ChargePaymentMethodDetailsCardChecksAddressLine1Check` | Same values. `CardVerification` was replaced with resource- and field-specific enums (with the same values) |
  | Charge | `ChargePaymentMethodDetailsCardChecks.AddressPostalCodeCheck` | `CardVerification` | `ChargePaymentMethodDetailsCardChecksAddressPostalCodeCheck` | Same values. `CardVerification` was replaced with resource- and field-specific enums (with the same values) |
  | Charge | `ChargePaymentMethodDetailsCardChecks.CVCCheck` | `CardVerification` | `ChargePaymentMethodDetailsCardChecksCVCCheck` | Same values. `CardVerification` was replaced with resource- and field-specific enums (with the same values) |
  | Customer | `Customer.Address` | `Address` | `*Address` | `Address` is nullable |
  | Customer | `CustomerParams.Source` | `*SourceParams` | `*string` | `customer.source` should be the ID of a source |
  | Discount | `Discount.CheckoutSession` | `*CheckoutSession` | `string` | `discount.checkout_session` is not expandable |
  | Discount | `Discount.Customer` | `string` | `*Customer` | `discount.customer` is expandable |
  | Invoice | `InvoiceCustomerTaxID.Type` | `TaxIDType` | `*TaxIDType` | `invoice.customer_tax_ids.type` is nullable |
  | Invoice | `Invoice.Type` | `*CollectionMethod` | `CollectionMethod` | `invoice.collection_method` is not nullable |
  | Invoice | `Invoice.CustomerName` | `*string` | `string` | `invoice.customer_name` is not nullable |
  | Invoice | `Invoice.CustomerPhone` | `*string` | `string` | `invoice.customer_phone` is not nullable |
  | Invoice | `Invoice.CustomerTaxExempt` | `CustomerTaxExempt` | `*CustomerTaxExempt` | `invoice.customer_tax_exempt` is nullable |
  | Invoice | `Invoice.StatusTransitions` | `InvoiceStatusTransitions` | `*InvoiceStatusTransitions` | `invoice.status_transitions` is nullable |
  | Issuing Card | `IssuingCardShippingParams.Name` | `string` | `*string` | Params fields should be pointers |
  | Issuing Dispute | `IssuingDispute.Status` | `*IssuingDisputeStatus` | `IssuingDisputeStatus` | `issuing_dispute.status` is not nullable |
  | Mandate | `MandatePaymentMethodDetails.Type` | `PaymentMethodType` | `MandatePaymentMethodDetailsType` | Unshare enum definition, same types |
  | PaymentIntent | `PaymentIntentPaymentMethodOptionsCard.Network` | `PaymentMethodCardNetwork` | `PaymentIntentPaymentMethodOptionsCardNetwork` | Unshare enum definition, same types (+ `cartes_bancaires` added to both) |
  | PaymentIntent | `PaymentIntentMandateDataCustomerAcceptanceParams.AcceptedAt` | `int64` | `*int64` | Params fields should be pointers |
  | PaymentIntent | `PaymentIntentMandateDataCustomerAcceptanceParams.Type` | `MandateCustomerAcceptanceType` | `*string` | Params do not use enums |
  | PaymentIntent | `PaymentIntentPaymentMethodDataParams.BillingDetails` | `*BillingDetailsParams` | `*PaymentIntentPaymentMethodDataBillingDetailsParams` | Unshare struct definition, same fields |
  | PaymentIntent | `PaymentIntent.Currency` | `string` | `Currency` | Use shared enum |
  | PaymentIntent | `PaymentIntent.ShippingDetails` | `ShippingDetails` | `*ShippingDetails` | `payment_intent.shipping` is nullable |
  | PaymentMethod | `PaymentMethodCardChecks.AddressLine1Check` | `CardVerification` | `PaymentMethodCardChecksAddressLine1Check` | Same values. `CardVerification` was replaced with separate definitions for `AddressLine1Check`, `AddressPostalCodeCheck`, and `CVCCheck` |
  | PaymentMethod | `PaymentMethodCardChecks.AddressPostalCodeCheck` | `CardVerification` | `PaymentMethodCardChecksAddressPostalCodeCheck` | Same values. `CardVerification` was replaced with separate definitions for `AddressLine1Check`, `AddressPostalCodeCheck`, and `CVCCheck` |
  | PaymentMethod | `PaymentMethodCardChecks.CVCCheck` | `CardVerification` | `PaymentMethodCardChecksCVCCheck` | Same values. `CardVerification` was replaced with separate definitions for `AddressLine1Check`, `AddressPostalCodeCheck`, and `CVCCheck` |
  | PaymentMethod | `BillingDetailsParams.ExpMonth` (now `PaymentMethodBillingDetailsParams.ExpMonth`) | `*string` | `*int64` | Previous type was incorrect |
  | PaymentMethod | `BillingDetailsParams.ExpYear` (now `PaymentMethodBillingDetailsParams.ExpYear`) | `*string` | `*int64` | Previous type was incorrect |
  | PaymentMethod | `PaymentMethodCard.ExpMonth` | `uint64` | `int64` | Be consistent with `int` types in library |
  | PaymentMethod | `PaymentMethodCard.ExpYear` | `uint64` | `int64` | Be consistent with `int` types in library |
  | Payout | `Payout.Description` | `*string` | `string` | `payout.description` is not nullable |
  | Person | `Person.Address` | `*AccountAddress` | `*Address` | `*AccountAddress` had `Town` field that is not applicable to `Person.Address` |
  | Person | `Person.AddressKana` | `*AccountAddress` | `*PersonAddressKana` | `AccountAddress` has been split into separate `PersonAddressKana` and `PersonAddressKanji` |
  | Person | `Person.AddressKanji` | `*AccountAddress` | `*PersonAddressKanji` | `AccountAddress` has been split into separate `PersonAddressKana` and `PersonAddressKanji` |
  | Person | `PersonParams.Address` | `*AccountAddressParams` | `*AddressParams` | `*AccountAddressParams` had `Town` field that is not applicable to `PersonParams.Address` |
  | Person | `PersonParams.AddressKana` | `*AccountAddressParams` | `*PersonAddressKanaParams` | `AccountAddressParams` has been split into separate `PersonAddressKanaParams` and `PersonAddressKanjiParams` |
  | Person | `PersonParams.AddressKanji` | `*AccountAddressParams` | `*PersonAddressKanjiParams` | `AccountAddressParams` has been split into separate `PersonAddressKanaParams` and `PersonAddressKanjiParams` |
  | Plan | `Plan.AggregateUsage` | `string` | `PlanAggregateUsage` | Use enum that was already defined but unused |
  | Plan | `Plan.TiersMode` | `string` | `PlanTiersMode` | Use enum that was already defined but unused |
  | Refund | `Refund.SourceTransferReversal` | `*Reversal` | `*TransferReversal` | Rename `Reversal` to `TransferReversal` for consistency with other Stripe client libraries |
  | Refund | `Refund.TransferReversal` | `*Reversal` | `*TransferReversal` | Rename `Reversal` to `TransferReversal` for consistency with other Stripe client libraries |
  | SetupIntent | `SetupIntentPaymentMethodOptionsACSSDebit.Currency` | `string` | `*SetupIntentPaymentMethodOptionsACSSDebitCurrency` | `SetupIntentPaymentMethodOptionsACSSDebitCurrency` is an alias to `string` type |
  | Source | `Source.Receiver` | `*ReceiverFlow` | `*SourceReceiver` | Rename `ReceiverFlow` to `SourceReceiver` for consistency with other Stripe client libraries |
  | Source | `Source.Redirect` | `*RedirectFlow` | `*SourceRedirect` | Rename `ReceiverFlow` to `SourceReceiver` for consistency with other Stripe client libraries |
  | Subscription | `SubscriptionListParams.Created` | `int64` | `*int64` | Params fields should be pointers |
  | Subscription | `SubscriptionListParams.Customer` | `string` | `*string` | Params fields should be pointers |
  | Subscription | `SubscriptionListParams.Plan` | `string` | `*string` | Params fields should be pointers |
  | Subscription | `SubscriptionListParams.Price` | `string` | `*string` | Params fields should be pointers |
  | Subscription | `SubscriptionListParams.Status` | `string` | `*string` | Params fields should be pointers |
  | Subscription | `Subscription.PauseCollection` | `SubscriptionPauseCollection` | `*SubscriptionPauseCollection` | `PauseCollection` is nullable |
  | Subscription | `Subscription.PendingInvoiceItemInterval` | `SubscriptionPendingInvoiceItemInterval` | `*SubscriptionPendingInvoiceItemInterval` | `PendingInvoiceItemInterval` is nullable |
  | SubscriptionItem | `SubscriptionItem.BillingThresholds` | `SubscriptionItemBillingThresholds` | `*SubscriptionItemBillingThresholds` | `BillingThresholds` is nullable |
  | SubscriptionSchedule | `SubscriptionScheduleListParams.CanceledAt` | `int64` | `*int64` | Params fields should be pointers |
  | SubscriptionSchedule | `SubscriptionScheduleListParams.CompletedAt` | `int64` | `*int64` | Params fields should be pointers |
  | SubscriptionSchedule | `SubscriptionScheduleListParams.Created` | `int64` | `*int64` | Params fields should be pointers |
  | SubscriptionSchedule | `SubscriptionScheduleListParams.Customer` | `string` | `*string` | Params fields should be pointers |
  | SubscriptionSchedule | `SubscriptionScheduleListParams.ReleasedAt` | `int64` | `*int64` | Params fields should be pointers |
  | SubscriptionSchedule | `SubscriptionScheduleDefaultSettings.CollectionMethod` | `SubscriptionCollectionMethod` | `*SubscriptionCollectionMethod` | `CollectionMethod` is nullable |
  | Terminal ConnectionToken | `TerminalConnectionTokenParams.Location` | `string` | `*string` | Params fields should be pointers |
  | Terminal Location | `TerminalLocationParams.Account` | `*AccountAddressParams` | `*AddressParams` | `AccountAddressParams` had extra `Town` field |
  | Terminal Location | `AccountAddressParams.Account` | `*AccountAddressParams` | `*Address` | `AccountAddressParams` had extra `Town` field, and also was the Params struct and not resource struct |
  | Terminal Reader | `TerminalReader.Location` | `string` | `*TerminalLocation` | `terminal_reader.location` is expandable |
  | Topup | `TopupParams.Source` | `*SourceParams` | `*string` | `topup.source` should be the ID of a source |
  | Topup | `Transfer.Destination` | `*TransferDestination` | `*Account` | `transfer.destination` should be (expandable) Account (previously, `UnmarshalJSON` would write the Account into `TransferDestination.Account`, but this should already be handled by `Account`'s `UnmarshalJSON`) |
  | Transfer | `Transfer.SourceTransaction` | `*BalanceTransactionSource` | `*Charge` | `transfer.source_transaction` should be (expandable) Charge |

- Moved `BalanceTransaction` iterator from `balance.go` to `balancetransaction.go`
- Fixed `BalanceTransactionSource` `UnmarshalJSON` for when `BalanceTransactionSource.Type == "transfer_reversal"` (previously, we were checking if `Type == "reversal"`, which was always false)
- For BankAccount and Card client methods, check that exactly one of `params.Account` and `params.Customer` is set (previously they could both be set, but only one would be used, and it was different between BankAccount and Card)
- Replace `CardVerification` with field-specific enums (with the same values)
- Move `Del` from `discount/client.go` to `customer/client.go` and rename to `DeleteDiscount`
- Move `DelSub` from `discount/client.go` to `subscription/client.go` and rename to `DeleteDiscount`
- Add separate parameter struct for CreditNote `ListPreviewLines` (renamed to `PreviewLines`) method (`[CreditNoteLineItemListPreviewParams -> CreditNotePreviewParams].Lines` `CreditNoteLineParams` -> `CreditNotePreviewLineParams`)
- Replace `FeeRefundParams.ApplicationFee` with `FeeRefundParams.Fee` and `FeeRefundParams.ID`
- Add separate parameter struct for Invoice `GetNext` (renamed to `Upcoming`) method (`InvoiceUpcomingParams`, and nested params `InvoiceUpcomingLinesInvoiceItemPriceDataParams`, `InvoiceUpcomingLinesInvoiceItemDiscountParams`, `InvoiceUpcomingLinesDiscountParams`, `InvoiceUpcomingLinesInvoiceItemPeriodParams`). `Upcoming`-only fields `Coupon`, `CustomerDetails`, `InvoiceItems`, `Subscription`, `SubscriptionBillingCycleAnchor`, `Schedule`, `SubscriptionBillingCycleAnchor`, `SubscriptionBillingCycleAnchorNow`, `SubscriptionBillingCycleAnchorUnchanged`, `SubscriptionCancelAt`, `SubscriptionCancelAtPeriodEnd`, `SubscriptionCancelNow`, `SubscriptionDefaultTaxRates`, `SubscriptionItems`, `SubscriptionProrationBehavior`, `SubscriptionProrationDate`, `SubscriptionStartDate`, `SubscriptionTrialEnd`, `SubscriptionTrialEndNow`, and `SubscriptionTrialFromPlan` are removed from `InvoiceParams`.
- Add separate structs for `BillingDetails` and `BillingDetailsParams`: `PaymentMethodBillingDetails`, `PaymentMethodBillingDetailsParams`
- Add separate structs for `PaymentMethodCardNetwork`: `PaymentMethodCardNetworksAvailable`, `PaymentMethodCardNetworksPreferred`

<!-- ---- up to R -->

<!-- ---- after S -->

## Deprecated

- SKU

## Removed

- `UnmarshalJSON` for resources that are not expandable: `BillingPortalSession`, `Capability`, `CheckoutSession`, `FileLink`, `InvoiceItem`, `LineItem`, `Person`, `WebhookEndpoint`
- `AccountRejectReason` (was only referenced in `account/client_test.go`, actual `AccountRejectParams.Reason` is `*string`)
- `AccountParams.RequestedCapabilities` (use Capabilities instead: https://stripe.com/docs/connect/account-capabilities)

  ```go
  // Before
  params := &stripe.AccountParams{
    RequestedCapabilities: ["card_payments"],
  }

  // After
  params := &stripe.AccountParams{
    Capabilities: &stripe.AccountCapabilitiesParams{
        CardPayments: &stripe.AccountCapabilitiesCardPaymentsParams{
            Requested: stripe.Bool(true),
        },
    },
  }
  ```

- ``AccountSettingsParams.Dashboard` and `AccountSettingsDashboardParams` (Note: `Dashboard` are still available on `AccountSettings`, but it's not available as parameters for any of the methods)
- `AccountCompany.RegistrationNumber` (Note: `RegistrationNumber` is still available on `AccountCompanyParams`, but is not returned in the response)
- `BalanceTransactionStatus`. It was meant to be an enum, but none of the enum values were defined, so it was just an alias for string.
- `CardParams.AccountType`. `AccountType` does not exist on any client method for Card. It does on BankAccount, which is similar.
- `id` param from CheckoutSessions `ListLineItems`. Use `CheckoutSessionListLineItemsParams.Session` instead.

  ```go
  // Before
  params := &stripe.CheckoutSessionListLineItemsParams{}
  i := ListLineItems("cs_123", params)

  // After
  params := &stripe.CheckoutSessionListLineItemsParams{Session: stripe.String("cs_123")}
  i := ListLineItems(params)
  ```

- `CheckoutSessionLineItemPriceDataRecurringParams.AggregateUsage`, `CheckoutSessionLineItemPriceDataRecurringParams.TrialPeriodDays`, and `CheckoutSessionLineItemPriceDataRecurringParams.UsageType`
- `CheckoutSessionPaymentIntentDataParams.Params`, `CheckoutSessionSetupIntentDataParams.Params`, `CheckoutSessionSubscriptionDataParams.Params`. `Params` should only be embedded in root method struct, and has extraneous fields not applicable to child/sub structs.
- `CheckoutSessionTotalDetailsBreakdownTax.TaxRate`. Use `CheckoutSessionTotalDetailsBreakdownTax.Rate`
- `CheckoutSessionTotalDetailsBreakdownTax.Deleted`
- `CustomerParams.Token`
- `Discount` `APIResource` embed
- `FilePurposeFoundersStockDocument` (`"founders_stock_document"` option for `File.Purpose`)
- `InvoiceParams.Paid`. Use `invoice.status` to check for status. `invoice.status` is a read-only field.
- `InvoiceParams.SubscriptionPlan` and `InvoiceParams.SubscriptionQuantity` (note: these would have been on `InvoiceUpcomingParams`)
- `InvoiceListLinesParams.Customer` and `InvoiceListLinesParams.Subscription` (these are not available for Invoice `ListLines`, but are available for `List`)
- `IssuingAuthorizationRequestHistoryViolatedAuthorizationControlEntity` and `IssuingAuthorizationRequestHistoryViolatedAuthorizationControlName` (unused enums)
- `IssuingCardSpendingControlsParams.SpendingLimitsCurrency`. `issuing_card` has `currency`, and `issuing_card.spending_controls.spending_limits.amount` will use that currency
- `IssuingDisputeEvidenceServiceNotAsDescribed.ProductDescription`, `IssuingDisputeEvidenceServiceNotAsDescribed.ProductType`, `IssuingDisputeEvidenceServiceNotAsDescribedParams.ProductDescription`, `IssuingDisputeEvidenceServiceNotAsDescribedParams.ProductType`, and `IssuingDisputeEvidenceServiceNotAsDescribedProductType`. `issuing_dispute.evidence.service_not_as_described` does not have `product_description` or `product_type`. `issuing_dispute.evidence.canceled` does.
- `LineItemTax.TaxRate`. Use `LineItemTax.Rate`
- `LineItem.Deleted`
- `LoginLink.RedirectURL`
- `PaymentIntentOffSession` (unused enum)
- `PaymentIntentConfirmParams.PaymentMethodTypes`
- `PaymentMethodFPX.TransactionID`
- `Payout.BankAccount` and `Payout.Card` (These fields were never populated, use `PayoutDestination.BankAccount` and `PayoutDestination.Card` instead)
- `PlanParams.ProductID`. Use `PlanParams.Product.ID`

  ```go
  // Before
  params := &stripe.PlanParams{
    ProductID: stripe.String("prod_12345abc"),
  }

  // After
  params := &stripe.PlanParams{
    Product: &stripe.PlanProductParams{
      ID:                  stripe.String("plan_id"),
    },
  }
  ```
- Remove `Shipping` and `ShippingRate` properties from `CheckoutSession` resource. Please use `ShippingCost` and `ShippingDetails` properties instead.
- Remove `DefaultCurrency` property from `Customer` resource. Please use `Currency` property instead.

- `Updated` and `UpdatedBy` from `RadarValueList`
- `Name` from `RadarValueListItem`
- `ReviewReasonType` type from `Review` resource. Use `ReviewReason` instead
- `SetupIntentCancellationReasonFailedInvoice` and `SetupIntentCancellationReasonFraudulent` values from `SetupIntentCancellationReason`
- `SigmaScheduledQueryRun.Query`. The field was invalid
- `SKUParams.Description` and `SKU.Description`
- `SourceMandateAcceptanceStatus`, `SourceMandateAcceptanceStatusAccepted`, `SourceMandateAcceptanceStatusRefused`, `SourceMandateNotificationMethod`, `SourceMandateNotificationMethodEmail`, `SourceMandateNotificationMethodManual`, and `SourceMandateNotificationMethodNone`
- `Source.TypeData` and SourceParams and replace with payment method-specific fields (AUBECSDebit, Bancontact, Card, CardPresent, EPS, Giropay, IDEAL, Klarna, Multibanco, P24, SEPACreditTransfer, SEPADebit, Sofort, ThreeDSecure, Wechat) and `Source.AppendTo` method
- `SourceTransaction.CustomerData`. The field was deprecated
- `SourceTransaction.TypeData` and `SourceTransaction.UnmarshalJSON`. Use payment specific fields - `ACHCreditTransfer`, `CHFCreditTransfer`, `GBPCreditTransfer`, `PaperCheck`, and `SEPACreditTransfer`
- `SubscriptionPaymentBehavior`, `SubscriptionPaymentBehaviorAllowIncomplete`, `SubscriptionPaymentBehaviorErrorIfIncomplete`, and `SubscriptionPaymentBehaviorPendingIfIncomplete`
- `SubscriptionProrationBehavior`, `SubscriptionProrationBehaviorAlwaysInvoice`, `SubscriptionProrationBehaviorCreateProrations`, and `SubscriptionProrationBehaviorNone`
- `SubscriptionStatusAll`
- `SubscriptionParams.Card`, `SubscriptionParams.Plan`, and `SubscriptionParams.Quantity`
- `Subscription.Plan` and `Subscription.Quantity`
- `SubscriptionItemParams.ID`. The field was deprecated
- `SubscriptionSchedulePhaseAddInvoiceItemPriceDataRecurringParams` and `SubscriptionSchedulePhaseAddInvoiceItemPriceDataParams`
- `Del` TaxRate
- `TerminalReaderGetParams`. Use `TerminalReaderParams`
- `TerminalReaderList.Location` and `TerminalReaderList.Status` (Not available for the list, but is available for individual `TerminalReader`s in `TerminalReaderList.Data`)
- `Token.Email` and `TokenParams.Email`
- `WebhookEndpointListParams.Created` and `WebhookEndpointListParams.CreatedRange` (use `StartingAfter` from `ListParams`)
- `WebhookEndpoint.Connected`
