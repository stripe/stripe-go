trigger QuoteLineStripeCouponAssociationTrigger on Quote_Line_Stripe_Coupon_Association__c (before insert, before update) {
    if (UserInfo.isMultiCurrencyOrganization()) {
        QuoteCouponAssociationTriggerHandler th = new QuoteCouponAssociationTriggerHandler();
        List<Quote_Line_Stripe_Coupon_Association__c> objs = (List<Quote_Line_Stripe_Coupon_Association__c>) Trigger.new;
        th.process(objs);
    }
}