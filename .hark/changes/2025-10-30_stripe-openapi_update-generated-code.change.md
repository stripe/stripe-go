---
title: Update generated code for private-preview
pr_url: https://github.com/stripe/stripe-go/pull/2198
is_stripe_api_change: true
released_in_version: 83.2.0-alpha.2
---

* Add support for `PaymentMethodPreview` on `DelegatedCheckoutRequestedSession`
* Add support for `OrderID` on `DelegatedCheckoutRequestedSessionOrderDetails`
* Add support for `Lead` on `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercialParams`, `V2CoreAccountConfigurationCardCreatorCapabilitiesCommercial`, `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialParams`, and `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercial`
* Add support for `GlobalAccountHolder` on `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercialParams` and `V2CoreAccountIdentityAttestationsTermsOfServiceCardCreatorCommercial`
* Add support for new value `commercial.lead.prepaid_card` on enum `V2CoreAccountRequirementsEntryImpactRestrictsCapability.Capability`
* Add support for new value `commercial.lead.prepaid_card` on enum `EventsV2CoreAccountIncludingConfigurationCardCreatorCapabilityStatusUpdatedEvent.UpdatedCapability`
