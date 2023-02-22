/**
 * Created by jmather-c on 2/17/23.
 */

public with sharing class QuoteCouponAssociationTriggerHandler {
    // allows us to bypass the trigger in tests
    @TestVisible
    private static Boolean TRIGGER_ENABLED = true;

    // left non-static and non-final so we can tweak and test in non-multi-curr envs
    @TestVisible
    private String ASSOC_FIELD_CURRENCY_ISO_CODE = constants.FIELD_CURRENCY_ISO_CODE;
    @TestVisible
    private String QUOTE_FIELD_CURRENCY_ISO_CODE = constants.FIELD_CURRENCY_ISO_CODE;
    @TestVisible
    private String COUPON_FIELD_CURRENCY_ISO_CODE = constants.FIELD_CURRENCY_ISO_CODE;
    @TestVisible
    private static final String LABEL_QUOTE = SBQQ__Quote__c.SObjectType.getDescribe().getLabel();
    @TestVisible
    private static final String LABEL_QUOTE_COUPON = Quote_Stripe_Coupon__c.SObjectType.getDescribe().getLabel();
    @TestVisible
    private static final String LABEL_QUOTE_STRIPE_COUPON_ASSOC = Quote_Stripe_Coupon_Association__c.SObjectType.getDescribe().getLabel();
    @TestVisible
    private static final String LABEL_QUOTE_LINE_STRIPE_COUPON_ASSOC = Quote_Line_Stripe_Coupon_Association__c.SObjectType.getDescribe().getLabel();
    @TestVisible
    private static final String FORMAT_DEPENDENT_ERROR_MESSAGE = 'The Currency of the {0} ({1}) does not match the Currency of the {2} ({3}).';
    private static final String FORMAT_ASSOC_ERROR_MESSAGE = 'The Currency of the {0} ({1}) does not match the Currency of the {2} ({3}) or the {4} ({5}).';

    public void process(List<Quote_Stripe_Coupon_Association__c> quoteAssocs) {
        if (TRIGGER_ENABLED == false) {
            return;
        }

        Set<Id> couponIds = pickIds(quoteAssocs, Quote_Stripe_Coupon_Association__c.Quote_Stripe_Coupon__c);
        Set<Id> quoteIds = pickIds(quoteAssocs, Quote_Stripe_Coupon_Association__c.Quote__c);

        Map<Id, String> couponCurrencies = getCouponIsoCodeMap(couponIds);
        Map<Id, String> quoteCurrencies = getQuoteIsoCodeMap(quoteIds);

        for (Quote_Stripe_Coupon_Association__c assoc : quoteAssocs) {
            String couponCurrency = couponCurrencies.get(assoc.Quote_Stripe_Coupon__c);
            String quoteCurrency = quoteCurrencies.get(assoc.Quote__c);
            String assocCurrency = (String) assoc.get(ASSOC_FIELD_CURRENCY_ISO_CODE);

            // this would only be null if the coupon did not end up being queried because it is not an Amount Off coupon
            if (couponCurrency == null) {
                continue;
            }

            validateCurrency(assocCurrency, quoteCurrency, couponCurrency, assoc, LABEL_QUOTE_STRIPE_COUPON_ASSOC);
        }
    }

    public void process(List<Quote_Line_Stripe_Coupon_Association__c> quoteLineAssocs) {
        if (TRIGGER_ENABLED == false) {
            return;
        }

        Set<Id> couponIds = pickIds(quoteLineAssocs, Quote_Line_Stripe_Coupon_Association__c.Quote_Stripe_Coupon__c);
        Set<Id> quoteLineIds = pickIds(quoteLineAssocs, Quote_Line_Stripe_Coupon_Association__c.Quote_Line__c);

        Map<Id, String> couponCurrencies = getCouponIsoCodeMap(couponIds);
        Map<Id, String> quoteLineCurrencies = getQuoteLineIsoCodeMap(quoteLineIds);

        for (Quote_Line_Stripe_Coupon_Association__c assoc : quoteLineAssocs) {
            String couponCurrency = couponCurrencies.get(assoc.Quote_Stripe_Coupon__c);
            String quoteCurrency = quoteLineCurrencies.get(assoc.Quote_Line__c);
            String assocCurrency = (String) assoc.get(ASSOC_FIELD_CURRENCY_ISO_CODE);

            // this would only be null if the coupon did not end up being queried because it is not an Amount Off coupon
            if (couponCurrency == null) {
                continue;
            }

            validateCurrency(assocCurrency, quoteCurrency, couponCurrency, assoc, LABEL_QUOTE_LINE_STRIPE_COUPON_ASSOC);
        }
    }

    @TestVisible
    private void validateCurrency(String assocCurrency, String quoteCurrency, String couponCurrency, SObject obj, String objLabel) {
        Boolean check1 = assocCurrency == couponCurrency;
        Boolean check2 = assocCurrency == quoteCurrency;
        Boolean check3 = couponCurrency == quoteCurrency;

        System.debug('Checks: ' + check1 + ', ' + check2 + ', ' + check3);

        if (check1 && check2 && check3) {
            return;
        }

        if (check1 == false && check2 == false) {
            String errorMsg = String.format(FORMAT_ASSOC_ERROR_MESSAGE,
                    new String[] { objLabel, assocCurrency, LABEL_QUOTE, quoteCurrency, LABEL_QUOTE_COUPON, couponCurrency});
            obj.addError(ASSOC_FIELD_CURRENCY_ISO_CODE, errorMsg);
        }

        if (check3 == false) {
            String errorMsg = String.format(FORMAT_DEPENDENT_ERROR_MESSAGE,
                    new String[] { LABEL_QUOTE, quoteCurrency, LABEL_QUOTE_COUPON, couponCurrency });
            obj.addError(errorMsg);
        }
    }

    @TestVisible
    private Map<Id, String> getCouponIsoCodeMap(Set<Id> couponIds) {
        String ids = utilities.idsToQueryString(couponIds);

        String query = 'SELECT Id, ' + COUPON_FIELD_CURRENCY_ISO_CODE + ' FROM Quote_Stripe_Coupon__c WHERE Id IN ' + ids;
        // make the validation not apply to % off coupons.
        query += ' AND Amount_Off__c > 0';
        List<SObject> coupons = Database.query(query);

        Map<Id, String> result = new Map<Id, String>();

        for (SObject coupon : coupons) {
            result.put(coupon.Id, (String) coupon.get(COUPON_FIELD_CURRENCY_ISO_CODE));
        }

        return result;
    }

    @TestVisible
    private Map<Id, String> getQuoteIsoCodeMap(Set<Id> quoteIds) {
        String ids = utilities.idsToQueryString(quoteIds);
        String query = 'SELECT Id, ' + QUOTE_FIELD_CURRENCY_ISO_CODE + ' FROM SBQQ__Quote__c WHERE Id IN ' + ids;
        List<SObject> quotes = Database.query(query);

        Map<Id, String> result = new Map<Id, String>();

        for (SObject quote : quotes) {
            result.put(quote.Id, (String) quote.get(QUOTE_FIELD_CURRENCY_ISO_CODE));
        }

        return result;
    }

    @TestVisible
    private Map<Id, String> getQuoteLineIsoCodeMap(Set<Id> quoteLineIds) {
        String ids = utilities.idsToQueryString(quoteLineIds);
        String query = 'SELECT Id, SBQQ__Quote__r.' + QUOTE_FIELD_CURRENCY_ISO_CODE + ' FROM SBQQ__QuoteLine__c WHERE Id IN ' + ids;
        List<SObject> quoteLines = Database.query(query);

        Map<Id, String> result = new Map<Id, String>();

        for (SObject quoteLine : quoteLines) {
            SObject quote = quoteLine.getSObject('SBQQ__Quote__r');
            result.put(quoteLine.Id, (String) quote.get(QUOTE_FIELD_CURRENCY_ISO_CODE));
        }

        return result;
    }

    @TestVisible
    private Set<Id> pickIds(List<SObject> objs, SObjectField field) {
        return pickIds(objs, field.getDescribe().getName());
    }

    private Set<Id> pickIds(List<SObject> objs, String field) {
        Set<Id> result = new Set<Id>();

        for (SObject obj : objs) {
            result.add((String) obj.get(field));
        }

        return result;
    }
}