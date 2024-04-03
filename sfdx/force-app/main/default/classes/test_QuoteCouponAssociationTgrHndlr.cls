// the name is shortened because of 40 char class name limit
@IsTest
public with sharing class test_QuoteCouponAssociationTgrHndlr {

    // we use this to save multiple queries to look things up
    private static Boolean isCpqInstalled = utilities.isCpqEnabled();
    private static String defaultIsoCurrency = UserInfo.getDefaultCurrency();
    private static String alternateIsoCurrency = utilities.getCurrencyIsoCodePreferNonDefault();
    private static Boolean isTrueMultiCurrencyEnvironment = UserInfo.isMultiCurrencyOrganization() && alternateIsoCurrency != defaultIsoCurrency;
    private static Boolean isNotTrueMultiCurrencyEnvironment = isTrueMultiCurrencyEnvironment == false;

    // this combines both the individual processing tests and ONLY runs when multi-curr
    // with two currencies is enabled.
    @IsTest
    static void testProcessQuoteAndQuoteLineCouponAssociationsAsASafeGuard() {
        if(!isCpqInstalled || isNotTrueMultiCurrencyEnvironment) {
            return;
        }

        QuoteCouponAssociationTriggerHandler th = new QuoteCouponAssociationTriggerHandler();

        SBQQ__Quote__c quote = test_updateOrderCouponTrigger.createQuote(alternateIsoCurrency);

        Quote_Stripe_Coupon__c coupon = test_updateOrderCouponTrigger.createStripeCoupon(alternateIsoCurrency);
        coupon.Percent_Off__c = null;
        coupon.Amount_Off__c = 1;
        update coupon;

        SBQQ__QuoteLine__c quoteLine = [SELECT Id FROM SBQQ__QuoteLine__c WHERE SBQQ__Quote__c = :quote.Id];

        Quote_Stripe_Coupon_Association__c quoteAssoc =  new Quote_Stripe_Coupon_Association__c();
        quoteAssoc.Quote__c = quote.Id;
        quoteAssoc.Quote_Stripe_Coupon__c = coupon.Id;
        quoteAssoc.put(th.ASSOC_FIELD_CURRENCY_ISO_CODE, alternateIsoCurrency);

        Quote_Line_Stripe_Coupon_Association__c quoteLineAssoc =  new Quote_Line_Stripe_Coupon_Association__c();
        quoteLineAssoc.Quote_Line__c = quoteLine.Id;
        quoteLineAssoc.Quote_Stripe_Coupon__c = coupon.Id;
        quoteLineAssoc.put(th.ASSOC_FIELD_CURRENCY_ISO_CODE, alternateIsoCurrency);

        List<Quote_Stripe_Coupon_Association__c> quoteAssocList = new List<Quote_Stripe_Coupon_Association__c> { quoteAssoc };
        List<Quote_Line_Stripe_Coupon_Association__c> quoteLineAssocList = new List<Quote_Line_Stripe_Coupon_Association__c> { quoteLineAssoc };

        Test.startTest();

        th.process(quoteAssocList);
        th.process(quoteLineAssocList);

        Test.stopTest();

        System.assert(quoteAssoc.getErrors().isEmpty());
        System.assert(quoteLineAssoc.getErrors().isEmpty());

        coupon.put(th.COUPON_FIELD_CURRENCY_ISO_CODE, defaultIsoCurrency);
        update coupon;

        th.process(quoteAssocList);
        th.process(quoteLineAssocList);

        System.assertEquals(1, quoteAssoc.getErrors().size(), 'Association should have an error');
        System.assertEquals(1, quoteLineAssoc.getErrors().size(), 'Association should have an error');
    }

    @IsTest
    static void testProcessQuoteCouponAssociationsWithAmountOff() {
        if(!isCpqInstalled) {
            return;
        }

        QuoteCouponAssociationTriggerHandler th = getNewTriggerHandler();

        SBQQ__Quote__c quote = test_updateOrderCouponTrigger.createQuote();
        Quote_Stripe_Coupon__c coupon = test_updateOrderCouponTrigger.createStripeCoupon();
        coupon.Percent_Off__c = null;
        coupon.Amount_Off__c = 1;
        update coupon;

        Quote_Stripe_Coupon_Association__c quoteAssoc =  new Quote_Stripe_Coupon_Association__c();
        quoteAssoc.Quote__c = quote.Id;
        quoteAssoc.Quote_Stripe_Coupon__c = coupon.Id;

        String altCurrencyToUse = alternateIsoCurrency;

        if (isTrueMultiCurrencyEnvironment) {
            quoteAssoc.put(th.ASSOC_FIELD_CURRENCY_ISO_CODE, defaultIsoCurrency);
        } else {
            // we pick a random currency since we're faking the constraints anyway
            altCurrencyToUse = 'EUR';

            quoteAssoc = setUpAlternateCurrencyFields(quoteAssoc, true, th, quote, coupon);
        }

        List<Quote_Stripe_Coupon_Association__c> quoteAssocList = new List<Quote_Stripe_Coupon_Association__c> { quoteAssoc };

        Test.startTest();

        th.process(quoteAssocList);

        System.debug(quoteAssoc.getErrors());
        System.assert(quoteAssoc.getErrors().isEmpty());

        coupon.put(th.COUPON_FIELD_CURRENCY_ISO_CODE, altCurrencyToUse);
        update coupon;

        th.process(quoteAssocList);

        Test.stopTest();

        System.assertEquals(1, quoteAssoc.getErrors().size(), 'Association should have an error');
    }

    @IsTest
    static void testProcessQuoteCouponAssociationsWithPercentOff() {
        if(!isCpqInstalled) {
            return;
        }

        QuoteCouponAssociationTriggerHandler th = getNewTriggerHandler();

        SBQQ__Quote__c quote = test_updateOrderCouponTrigger.createQuote();
        Quote_Stripe_Coupon__c coupon = test_updateOrderCouponTrigger.createStripeCoupon();
        update coupon;

        Quote_Stripe_Coupon_Association__c quoteAssoc =  new Quote_Stripe_Coupon_Association__c();
        quoteAssoc.Quote__c = quote.Id;
        quoteAssoc.Quote_Stripe_Coupon__c = coupon.Id;

        String altCurrencyToUse = alternateIsoCurrency;

        if (isTrueMultiCurrencyEnvironment) {
            quoteAssoc.put(th.ASSOC_FIELD_CURRENCY_ISO_CODE, defaultIsoCurrency);
        } else {
            // we pick a random currency since we're faking the constraints anyway
            altCurrencyToUse = 'EUR';

            quoteAssoc = setUpAlternateCurrencyFields(quoteAssoc, false, th, quote, coupon);
        }

        List<Quote_Stripe_Coupon_Association__c> quoteAssocList = new List<Quote_Stripe_Coupon_Association__c> { quoteAssoc };

        Test.startTest();

        th.process(quoteAssocList);

        System.assert(quoteAssoc.getErrors().isEmpty());

        coupon.put(th.COUPON_FIELD_CURRENCY_ISO_CODE, altCurrencyToUse);
        update coupon;

        th.process(quoteAssocList);

        Test.stopTest();

        System.assertEquals(0, quoteAssoc.getErrors().size(), 'Association should still not have an error');
    }

    @IsTest
    static void testProcessQuoteLineCouponAssociationsWithAmountOff() {
        Boolean isCpqInstalled = utilities.isCpqEnabled();
        if(!isCpqInstalled) {
            return;
        }

        QuoteCouponAssociationTriggerHandler th = getNewTriggerHandler();

        SBQQ__Quote__c quote = test_updateOrderCouponTrigger.createQuote();

        Quote_Stripe_Coupon__c coupon = test_updateOrderCouponTrigger.createStripeCoupon();
        coupon.Percent_Off__c = null;
        coupon.Amount_Off__c = 1;
        update coupon;

        SBQQ__QuoteLine__c quoteLine = [SELECT Id FROM SBQQ__QuoteLine__c WHERE SBQQ__Quote__c = :quote.Id];

        Quote_Line_Stripe_Coupon_Association__c quoteLineAssoc =  new Quote_Line_Stripe_Coupon_Association__c();
        quoteLineAssoc.Quote_Line__c = quoteLine.Id;
        quoteLineAssoc.Quote_Stripe_Coupon__c = coupon.Id;

        String altCurrencyToUse = alternateIsoCurrency;

        if (isTrueMultiCurrencyEnvironment) {
            quoteLineAssoc.put(th.ASSOC_FIELD_CURRENCY_ISO_CODE, defaultIsoCurrency);
        } else {
            // we pick a random currency since we're faking the constraints anyway
            altCurrencyToUse = 'EUR';

            quoteLineAssoc = setUpAlternateCurrencyFields(quoteLineAssoc, true, th, quote, coupon);
        }

        List<Quote_Line_Stripe_Coupon_Association__c> quoteLineAssocList = new List<Quote_Line_Stripe_Coupon_Association__c> { quoteLineAssoc };

        Test.startTest();

        th.process(quoteLineAssocList);

        System.assert(quoteLineAssoc.getErrors().isEmpty());

        coupon.put(th.COUPON_FIELD_CURRENCY_ISO_CODE, altCurrencyToUse);
        update coupon;

        th.process(quoteLineAssocList);

        Test.stopTest();

        System.assertEquals(1, quoteLineAssoc.getErrors().size(), 'Association should have an error');
    }

    @IsTest
    static void testProcessQuoteLineCouponAssociationsWithPercentOff() {
        Boolean isCpqInstalled = utilities.isCpqEnabled();
        if(!isCpqInstalled) {
            return;
        }

        QuoteCouponAssociationTriggerHandler th = getNewTriggerHandler();

        SBQQ__Quote__c quote = test_updateOrderCouponTrigger.createQuote();
        Quote_Stripe_Coupon__c coupon = test_updateOrderCouponTrigger.createStripeCoupon();
        SBQQ__QuoteLine__c quoteLine = [SELECT Id FROM SBQQ__QuoteLine__c WHERE SBQQ__Quote__c = :quote.Id];

        Quote_Line_Stripe_Coupon_Association__c quoteLineAssoc =  new Quote_Line_Stripe_Coupon_Association__c();
        quoteLineAssoc.Quote_Line__c = quoteLine.Id;
        quoteLineAssoc.Quote_Stripe_Coupon__c = coupon.Id;

        String altCurrencyToUse = alternateIsoCurrency;

        if (isTrueMultiCurrencyEnvironment) {
            quoteLineAssoc.put(th.ASSOC_FIELD_CURRENCY_ISO_CODE, defaultIsoCurrency);
        } else {
            // we pick a random currency since we're faking the constraints anyway
            altCurrencyToUse = 'EUR';

            quoteLineAssoc = setUpAlternateCurrencyFields(quoteLineAssoc, false, th, quote, coupon);
        }

        List<Quote_Line_Stripe_Coupon_Association__c> quoteLineAssocList = new List<Quote_Line_Stripe_Coupon_Association__c> { quoteLineAssoc };

        Test.startTest();

        th.process(quoteLineAssocList);

        System.assert(quoteLineAssoc.getErrors().isEmpty());

        coupon.put(th.COUPON_FIELD_CURRENCY_ISO_CODE, altCurrencyToUse);
        update coupon;

        th.process(quoteLineAssocList);

        Test.stopTest();

        System.assertEquals(0, quoteLineAssoc.getErrors().size(), 'Association should still not have an error');
    }

    @IsTest
    static void testValidateCurrency() {
        QuoteCouponAssociationTriggerHandler th = getNewTriggerHandler();

        // the validation is root object independent, so we will just use the quote coupon association
        Quote_Stripe_Coupon_Association__c assoc1 = new Quote_Stripe_Coupon_Association__c();
        Quote_Stripe_Coupon_Association__c assoc2 = new Quote_Stripe_Coupon_Association__c();
        Quote_Stripe_Coupon_Association__c assoc3 = new Quote_Stripe_Coupon_Association__c();
        Quote_Stripe_Coupon_Association__c assoc4 = new Quote_Stripe_Coupon_Association__c();
        String objLabel = 'MyObj';

        Test.startTest();

        th.validateCurrency('USD', 'USD', 'USD', assoc1, objLabel);
        th.validateCurrency('EUR', 'USD', 'USD', assoc2, objLabel);
        th.validateCurrency('USD', 'EUR', 'USD', assoc3, objLabel);
        th.validateCurrency('USD', 'USD', 'EUR', assoc4, objLabel);

        Test.stopTest();

        assertHasNoErrors(assoc1);

        Database.Error err = assertHasOneError(assoc2);
        System.assert(err.getFields().contains(th.ASSOC_FIELD_CURRENCY_ISO_CODE));
        System.assert(err.getMessage().contains(objLabel + ' (EUR)'));
        System.assert(err.getMessage().contains(QuoteCouponAssociationTriggerHandler.LABEL_QUOTE_COUPON + ' (USD)'));
        System.assert(err.getMessage().contains(QuoteCouponAssociationTriggerHandler.LABEL_QUOTE + ' (USD)'));

        err = assertHasOneError(assoc3);
        System.assert(err.getFields().contains(th.ASSOC_FIELD_CURRENCY_ISO_CODE) == false);
        System.assert(err.getMessage().contains(objLabel + ' (EUR)') == false);
        System.assert(err.getMessage().contains(objLabel + ' (USD)') == false);
        System.assert(err.getMessage().contains(QuoteCouponAssociationTriggerHandler.LABEL_QUOTE_COUPON + ' (USD)'));
        System.assert(err.getMessage().contains(QuoteCouponAssociationTriggerHandler.LABEL_QUOTE + ' (EUR)'));

        err = assertHasOneError(assoc4);
        System.assert(err.getFields().contains(th.ASSOC_FIELD_CURRENCY_ISO_CODE) == false);
        System.assert(err.getMessage().contains(objLabel + ' (EUR)') == false);
        System.assert(err.getMessage().contains(objLabel + ' (USD)') == false);
        System.assert(err.getMessage().contains(QuoteCouponAssociationTriggerHandler.LABEL_QUOTE_COUPON + ' (EUR)'));
        System.assert(err.getMessage().contains(QuoteCouponAssociationTriggerHandler.LABEL_QUOTE + ' (USD)'));
    }

    @IsTest
    static void testIsoCodeMaps() {
        Boolean isCpqInstalled = utilities.isCpqEnabled();
        if(!isCpqInstalled) {
            return;
        }

        QuoteCouponAssociationTriggerHandler th = getNewTriggerHandler();

        SBQQ__Quote__c quote = test_updateOrderCouponTrigger.createQuote();
        Quote_Stripe_Coupon__c coupon1 = test_updateOrderCouponTrigger.createStripeCoupon();
        Quote_Stripe_Coupon__c coupon2 = test_updateOrderCouponTrigger.createStripeCoupon();

        // set coupons to amount only
        coupon1.Percent_Off__c = null;
        coupon1.Amount_Off__c = 1;
        coupon2.Percent_Off__c = null;
        coupon2.Amount_Off__c = 1;
        update new List<SObject> { coupon1, coupon2 };

        test_updateOrderCouponTrigger.createStripeCouponQuoteAssociation(quote.Id, coupon1.Id);
        SBQQ__QuoteLine__c quoteLine = [SELECT Id FROM SBQQ__QuoteLine__c WHERE SBQQ__Quote__c = :quote.Id];
        test_updateOrderCouponTrigger.createStripeCouponQuoteLineAssociation(quoteLine.Id, coupon2.Id);

        // update the quote data, and play a bunch of games to do so
        List<String> fields = new List<String>{
                'Id',
                'SBQQ__Account__c',
                'SBQQ__Status__c',
                'SBQQ__StartDate__c',
                'SBQQ__PricebookId__c',
                'SBQQ__Opportunity2__c',
                'SBQQ__SubscriptionTerm__c',
                'SBQQ__ContractingMethod__c',
                'SBQQ__Primary__c'
        };

        List<String> multiCurrFields = new List<String> { constants.FIELD_CURRENCY_ISO_CODE };

        if (th.QUOTE_FIELD_CURRENCY_ISO_CODE != constants.FIELD_CURRENCY_ISO_CODE) {
            fields.add(th.QUOTE_FIELD_CURRENCY_ISO_CODE);
            multiCurrFields.clear();
        }

        List<SBQQ__Quote__c> quotes = utilities.execMultiCurrLookupQuery(
                SBQQ__Quote__c.SObjectType,
                'Id',
                new Set<Id> { quote.Id },
                fields,
                multiCurrFields
        );

        quote = quotes.get(0);

        Test.startTest();

        // we use this to flip from Name to Name__c for the override on non-multi-curr envs
        Map<Id, String> couponCodes = th.getCouponIsoCodeMap(new Set<Id> { coupon1.Id, coupon2.Id });
        Map<Id, String> quoteCodes = th.getQuoteIsoCodeMap(new Set<Id> { quote.Id });
        Map<Id, String> quoteLineCodes = th.getQuoteLineIsoCodeMap(new Set<Id> { quoteLine.Id });

        Test.stopTest();

        // these checks ref th.FIELD_CURRENCY_ISO_CODE so they will work with multicurr disabled.
        System.assertEquals(2, couponCodes.size());
        System.debug(couponCodes);
        System.assertEquals(coupon1.get(th.COUPON_FIELD_CURRENCY_ISO_CODE), couponCodes.get(coupon1.Id));
        System.assertEquals(coupon2.get(th.COUPON_FIELD_CURRENCY_ISO_CODE), couponCodes.get(coupon2.Id));

        System.debug(quoteCodes);
        System.assertEquals(1, quoteCodes.size());
        System.assertEquals(quote.get(th.QUOTE_FIELD_CURRENCY_ISO_CODE), quoteCodes.get(quote.Id));

        System.debug(quoteLineCodes);
        System.assertEquals(1, quoteLineCodes.size());
        System.assertEquals(quote.get(th.QUOTE_FIELD_CURRENCY_ISO_CODE), quoteLineCodes.get(quoteLine.Id));
    }

    @IsTest
    static void testPick() {
        QuoteCouponAssociationTriggerHandler th = getNewTriggerHandler();

        Quote_Stripe_Coupon__c coupon1 = test_updateOrderCouponTrigger.createStripeCoupon();
        Quote_Stripe_Coupon__c coupon2 = test_updateOrderCouponTrigger.createStripeCoupon();
        List<SObject> objs = new List<SObject> { coupon1, coupon2 };

        Test.startTest();

        Set<Id> ids = th.pickIds(objs, Quote_Stripe_Coupon__c.Id);

        Test.stopTest();

        System.assert(ids.contains(coupon1.Id));
        System.assert(ids.contains(coupon2.Id));
    }

    static Quote_Line_Stripe_Coupon_Association__c setUpAlternateCurrencyFields(
            Quote_Line_Stripe_Coupon_Association__c quoteLineAssoc,
            Boolean match,
            QuoteCouponAssociationTriggerHandler th,
            SBQQ__Quote__c quote,
            Quote_Stripe_Coupon__c coupon) {
        // we're going to use the name as the 'currency', so we have to insert to populate it
        // and we're disabling the trigger from processing here because we're not /testing/ it here.
        QuoteCouponAssociationTriggerHandler.TRIGGER_ENABLED = false;
        insert quoteLineAssoc;
        QuoteCouponAssociationTriggerHandler.TRIGGER_ENABLED = true;

        // we need to pull again to get the name
        quoteLineAssoc = [SELECT Id, Name, Quote_Line__c, Quote_Stripe_Coupon__c FROM Quote_Line_Stripe_Coupon_Association__c WHERE Id = :quoteLineAssoc.Id];

        // we set the pointer fields to the name of the coupon
        // we're manually setting the values wrong to ensure it /should/ fail if it were actually checked
        // but it won't be checked because the coupon is a percent off.
        String suffix = match ? '' : ' No Match';
        quote.put(th.QUOTE_FIELD_CURRENCY_ISO_CODE, quoteLineAssoc.get(th.ASSOC_FIELD_CURRENCY_ISO_CODE) + suffix);
        coupon.put(th.COUPON_FIELD_CURRENCY_ISO_CODE, quoteLineAssoc.get(th.ASSOC_FIELD_CURRENCY_ISO_CODE) + suffix);

        update new List<SObject> { quote, coupon };

        return quoteLineAssoc;
    }

    static Quote_Stripe_Coupon_Association__c setUpAlternateCurrencyFields(
            Quote_Stripe_Coupon_Association__c quoteAssoc,
            Boolean match,
            QuoteCouponAssociationTriggerHandler th,
            SBQQ__Quote__c quote,
            Quote_Stripe_Coupon__c coupon) {
        // we're going to use the name as the 'currency', so we have to insert to populate it
        // and we're disabling the trigger from processing here because we're not /testing/ it here.
        QuoteCouponAssociationTriggerHandler.TRIGGER_ENABLED = false;
        insert quoteAssoc;
        QuoteCouponAssociationTriggerHandler.TRIGGER_ENABLED = true;

        // we need to pull again to get the name
        quoteAssoc = [SELECT Id, Name, Quote__c, Quote_Stripe_Coupon__c FROM Quote_Stripe_Coupon_Association__c WHERE Id = :quoteAssoc.Id];

        // we set the pointer fields to the name of the coupon
        // we're manually setting the values wrong to ensure it /should/ fail if it were actually checked
        // but it won't be checked because the coupon is a percent off.
        String suffix = match ? '' : ' No Match';
        quote.put(th.QUOTE_FIELD_CURRENCY_ISO_CODE, quoteAssoc.get(th.ASSOC_FIELD_CURRENCY_ISO_CODE) + suffix);
        coupon.put(th.COUPON_FIELD_CURRENCY_ISO_CODE, quoteAssoc.get(th.ASSOC_FIELD_CURRENCY_ISO_CODE) + suffix);

        update new List<SObject> { quote, coupon };

        return quoteAssoc;
    }

    static void assertHasNoErrors(SObject obj) {
        List<Database.Error> errors = obj.getErrors();
        System.assertEquals(0, errors.size());
    }

    static Database.Error assertHasOneError(SObject obj) {
        List<Database.Error> errors = obj.getErrors();
        System.assertEquals(1, errors.size());
        System.debug('Err: ' + errors.get(0).getMessage());
        return errors.get(0);
    }

    static QuoteCouponAssociationTriggerHandler getNewTriggerHandler() {
        QuoteCouponAssociationTriggerHandler th = new QuoteCouponAssociationTriggerHandler();

        if (isNotTrueMultiCurrencyEnvironment) {
            th.ASSOC_FIELD_CURRENCY_ISO_CODE = 'Name';
            th.QUOTE_FIELD_CURRENCY_ISO_CODE = 'SBQQ__Notes__c';
            th.COUPON_FIELD_CURRENCY_ISO_CODE = 'Name__c';
        }

        return th;
    }
}