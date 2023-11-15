@IsTest
public with sharing class test_utilities {
    @IsTest
    static public void testPackageVersion() {
        Test.startTest();
        String packageVersion = utilities.getPackageVersion();
        Test.stopTest();

        System.assert(packageVersion != null);
    }

    @IsTest
    static public void testPackageVersionString() {
        Test.startTest();
        String packageVersion = utilities.getPackageVersionString();
        Test.stopTest();

        System.assert(packageVersion != null);
    }

    @IsTest
    static void testGetSobjectByMap() {
        Map<String, Object> q = new Map<String, Object> {
            'Account__c' => '123'
        };

        Test.startTest();
        SBQQ__Quote__c quote = (SBQQ__Quote__c) utilities.getSobjectByMap(q, 'Quote__c', 'SBQQ__');
        Test.stopTest();

        System.assertNotEquals(null, quote);
        System.assertEquals('123', quote.SBQQ__Account__c);
    }

    @IsTest
    static void testGetConnectionKey() {
        Exception ex;
        Setup_Connection_Data__mdt testSetupData = test_setupAssistant.getTestStripeConnectionKey();
        test_setupAssistant.setTestGlobalKey(testSetupData);

        String key = utilities.getStripeConnectionKey();

        utilities.setupConfigMetadata.clear();

        try {
            utilities.getStripeConnectionKey();
        } catch (Exception e) {
            ex = e;
        }

        System.assertNotEquals(null, ex, 'Got an exception');
        System.assertEquals(testSetupData.Global_Key__c, key, 'Got the right key');
    }

    @IsTest
    static void testGetOAuthSigningKey() {
        Setup_Connection_Data__mdt testSetupData = test_setupAssistant.getTestStripeConnectionKey();
        test_setupAssistant.setTestGlobalKey(testSetupData);

        Blob key = utilities.getOAuthStateSigningKey();

        utilities.setupConfigMetadata.clear();

        Blob failed = utilities.getOAuthStateSigningKey();

        System.assertEquals(null, failed, 'returns null with no record set');
        System.assertNotEquals(null, key, 'Returns a value with record set');
    }

    @IsTest
    static void testIsDevModeEnabled() {
        Setup_Connection_Data__mdt testSetupData = test_setupAssistant.getTestStripeConnectionKey();
        test_setupAssistant.setTestGlobalKey(testSetupData);

        testSetupData.Enable_JS_Debug__c = true;

        Boolean yep1 = utilities.isDeveloperModeEnabled();
        Boolean yepJs = utilities.isJSDebuggingEnabled();
        Boolean nopeApex2 = utilities.isApexDebuggingEnabled();

        testSetupData.Enable_Apex_Debug__c = true;

        Boolean yepAll = utilities.isDeveloperModeEnabled();
        Boolean yepJs2 = utilities.isJSDebuggingEnabled();
        Boolean yepApex = utilities.isApexDebuggingEnabled();

        utilities.setupConfigMetadata.clear();

        Boolean nopeAll = utilities.isDeveloperModeEnabled();
        Boolean nopeJs = utilities.isJSDebuggingEnabled();
        Boolean nopeApex = utilities.isApexDebuggingEnabled();

        System.assertEquals(false, nopeAll, 'Nope all');
        System.assertEquals(false, nopeJs, 'Nope JS');
        System.assertEquals(false, nopeApex, 'Nope Apex');
        System.assertEquals(true, yep1, 'Yep all');
        System.assertEquals(true, yepJs, 'Yep JS');
        System.assertEquals(false, nopeApex2, 'Nope Apex 2');
        System.assertEquals(true, yepAll, 'Yep all 2');
        System.assertEquals(true, yepJs2, 'Yep JS 2');
        System.assertEquals(true, yepApex, 'Yep Apex');
    }
}
