---
title: Update generated code for private-preview
pr_link: https://github.com/stripe/stripe-go/pull/2155
is_stripe_api_change: true
released_in_version: 83.1.0-alpha.3
---

* Add support for new resource `PaymentMethodBalance`
* Add support for `CheckBalance` method on resource `PaymentMethod`
* Add support for `Benefits` on `Card`, `ChargePaymentMethodDetailsCard`, `ConfirmationTokenPaymentMethodPreviewCard`, and `PaymentMethodCard`
* Add support for `Benefit` on `PaymentIntentConfirmPaymentDetailsParams`, `PaymentIntentPaymentDetailsParams`, and `PaymentIntentPaymentDetails`
* Add support for `SetupDetails` on `SetupIntentConfirmParams`, `SetupIntentParams`, and `SetupIntent`
* Add support for new value `card_creator` on enum `V2CoreAccount.AppliedConfigurations`
* Add support for `CardCreator` on `V2CoreAccountConfigurationParams`, `V2CoreAccountConfiguration`, `V2CoreAccountIdentityAttestationsTermsOfServiceParams`, and `V2CoreAccountIdentityAttestationsTermsOfService`
* Add support for new values `commercial.celtic.charge_card`, `commercial.celtic.spend_card`, `commercial.cross_river_bank.charge_card`, `commercial.cross_river_bank.spend_card`, `commercial.stripe.charge_card`, and `commercial.stripe.prepaid_card` on enum `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for new value `card_creator` on enum `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Configuration`
* Add support for thin events `V2CoreAccountIncludingConfigurationCardCreatorCapabilityStatusUpdatedEvent` and `V2CoreAccountIncludingConfigurationCardCreatorUpdatedEvent` with related object `V2CoreAccount`
* Remove support for thin events `V1CustomerDiscountCreatedEvent`, `V1CustomerDiscountDeletedEvent`, and `V1CustomerDiscountUpdatedEvent` with related object `Discount`
