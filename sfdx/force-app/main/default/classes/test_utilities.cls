@IsTest
public with sharing class test_utilities {
    static Test_Tools testTools = new Test_Tools();

    @IsTest
    static void testDmlManager() {
        String profile = Test_Tools.PROFILE_CHATTER_FREE_USER;
        String uniqueId = 'dmlUnprivileged';
        User unprivUser = testTools.insertUser(profile, uniqueId);

        Exception unknownAccessModifier;
        Exception unknownAccessObject;
        Exception unknownAccessField;
        Exception unprivilegedObject;
        Exception unprivilegedField;
        try {
            Map<String, Map<String, Set<String>>> failedAccess = new Map<String, Map<String, Set<String>>> {
                    'Account' => new Map<String, Set<String>> {
                        'Foo' => new Set<String> {'bar'}
                    }
            };
            dmlManager.checkAccess(failedAccess);
        } catch (Exception e) {
            unknownAccessModifier = e;
        }

        try {
            Map<String, Map<String, Set<String>>> failedAccess = new Map<String, Map<String, Set<String>>> {
                    'Foo' => new Map<String, Set<String>> {
                        'Delete' => new Set<String> {}
                    }
            };
            dmlManager.checkAccess(failedAccess);
        } catch (Exception e) {
            unknownAccessObject = e;
        }

        try {
            Map<String, Map<String, Set<String>>> failedAccess = new Map<String, Map<String, Set<String>>> {
                    'Account' => new Map<String, Set<String>> {
                        'Read' => new Set<String> {'bar'}
                    }
            };
            dmlManager.checkAccess(failedAccess);
        } catch (Exception e) {
            unknownAccessField = e;
        }

        try {
            Map<String, Map<String, Set<String>>> failedAccess = new Map<String, Map<String, Set<String>>> {
                    'Account' => new Map<String, Set<String>> {
                        'Delete' => new Set<String> {}
                    }
            };
            System.runAs(unprivUser) {
                dmlManager.checkAccess(failedAccess);
            }
        } catch (Exception e) {
            unprivilegedObject = e;
        }

        try {
            Map<String, Map<String, Set<String>>> failedAccess = new Map<String, Map<String, Set<String>>> {
                    'User' => new Map<String, Set<String>> {
                        'Read' => new Set<String> {'AccountId'}
                    }
            };

            System.runAs(unprivUser) {
                dmlManager.checkAccess(failedAccess);
            }
        } catch (Exception e) {
            unprivilegedField = e;
        }

        Map<String, Map<String, Set<String>>> access = new Map<String, Map<String, Set<String>>> {
                'Account' => new Map<String, Set<String>> {
                        'Create' => new Set<String> {'Name'},
                        'Read' => new Set<String> {'Name'},
                        'Update' => new Set<String> {'Name'},
                        'Delete' => new Set<String> {}
                }
        };

        dmlManager.checkAccess(access);

        System.assertNotEquals(null, unknownAccessModifier, 'Got an exception for unknown access modifier');
        System.assertNotEquals(null, unknownAccessObject, 'Got an exception for unknown access object');
        System.assertNotEquals(null, unknownAccessField, 'Got an exception for unknown access field');
        System.assertNotEquals(null, unprivilegedObject, 'Got an exception for unprivileged object');
        System.assertNotEquals(null, unprivilegedField, 'Got an exception for unprivileged field');
    }

    @IsTest
    static public void getPrimaryStripeAccount() {
        Stripe_Account__c stripeAccount = new Stripe_Account__c();
        String stipeAcctId = 'acct_1234567890';
        stripeAccount.Stripe_ID__c = stipeAcctId;
        stripeAccount.Is_Live_Mode__c = false;
        stripeAccount.External_ID__c = stipeAcctId + '|' + stripeAccount.Is_Live_Mode__c;
        stripeAccount.Name = stipeAcctId;
        stripeAccount.Is_Primary__c = true;
        insert(stripeAccount);

        Test.startTest();
        Stripe_Account__c result = utilities.getPrimaryStripeAccount();
        Test.stopTest();

        System.assertNotEquals(null, result);
        System.assertEquals(result.Stripe_ID__c, stipeAcctId);
    }

    @IsTest
    static public void idsToQueryString() {
        Set<Id> ids = new Set<Id> {
            '001000000000000AAA',
            '001000000000000BBB',
            '001000000000000CCC'
        };

        String query = utilities.idsToQueryString(ids);
        System.assertEquals('(\'001000000000000AAA\',\'001000000000000BBB\',\'001000000000000CCC\')', query);
    }
    
    @IsTest
    static public void isUsingV2Middleware() {
        Setup_Connection_Data__mdt testSetupData = test_setupAssistant.getTestStripeConnectionKey();
        testSetupData.Middleware_Infrastructure_Version__c = 'v2';
        test_setupAssistant.setTestGlobalKey(testSetupData);
        test_setupAssistant.insertTestConnectedRecord();
        Test.setMock(HttpCalloutMock.class, new test_setupAssistant.UnifiedConfigMock());

        Test.startTest();
        Boolean isV2 = utilities.isUsingV2Middleware();
        Test.stopTest();

        System.assert(isV2, true);
    }

    @IsTest
    static public void isCPQEnabled() {
        Test.startTest();
        Boolean isCpqEnabled = utilities.isCpqEnabled();
        Test.stopTest();

        System.assert(isCpqEnabled, true);
    }


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
