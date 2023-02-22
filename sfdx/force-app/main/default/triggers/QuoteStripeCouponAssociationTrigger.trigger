/**
 * Created by jmather-c on 2/17/23.
 */

trigger QuoteStripeCouponAssociationTrigger on Quote_Stripe_Coupon_Association__c (before insert, before update) {
    if (UserInfo.isMultiCurrencyOrganization()) {
        QuoteCouponAssociationTriggerHandler th = new QuoteCouponAssociationTriggerHandler();
        List<Quote_Stripe_Coupon_Association__c> objs = (List<Quote_Stripe_Coupon_Association__c>) Trigger.new;
        th.process(objs);
    }
}