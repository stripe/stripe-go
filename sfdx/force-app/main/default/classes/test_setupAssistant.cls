@isTest
public with sharing class test_setupAssistant {
    @isTest
    static public void testSetupAssistant_getSetupData() {
        setupAssistant.saveData(new Map<String, Object>(), false);
        insertTestConnectedRecord();

        String response = setupAssistant.getSetupData();

        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assert(resultsMap.containsKey('setupData'), '`data` not found');
    }

    @isTest
    static public void testUtilities_isCpqEnabled() {
        Boolean isCpqEnabled = utilities.isCpqEnabled();
        System.assertNotEquals(null, isCpqEnabled);
    }

    @isTest
    static public void testUtilities_setOrgType() {
        String response = utilities.setOrgType();

        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));
        System.assertEquals(true, responseMap.get('isSuccess'));

        Organization_Type__c orgTypeSetting = Organization_Type__c.getOrgDefaults();
        System.assertNotEquals(orgTypeSetting.isSandbox__c, null);
    }


    @isTest
    static public void testSetupAssistant_validateConnectionStatus() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());
        String connectedCallbackResponse = setupAssistant.validateConnectionStatus(true, '');
        Test.stopTest();

        Map<String, Object> connectedCallbackResponseMap = (Map<String, Object>)JSON.deserializeUntyped(connectedCallbackResponse);
        System.assert((Boolean)connectedCallbackResponseMap.get('isSuccess'), connectedCallbackResponseMap.get('error'));

        Map<String, Object> connectecCallbackResultsMap = (Map<String, Object>)connectedCallbackResponseMap.get('results');
        System.assertEquals('stripeConnected', connectecCallbackResultsMap.get('isStripeConnected'));
        System.assertEquals('salesforceConnected', connectecCallbackResultsMap.get('isSalesforceConnected'));
    }

    @isTest
    static public void testSetupAssistant_connected() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());
        String response = setupAssistant.validateConnectionStatus(true, '');
        Test.stopTest();

        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));
        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals('stripeConnected', resultsMap.get('isStripeConnected'));
        System.assertEquals('salesforceConnected', resultsMap.get('isSalesforceConnected'));
    }

    @isTest
    static public void testSetupAssistant_stripeNotConnected() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeNotConnectedMock());
        String stripeResponse = setupAssistant.validateConnectionStatus(false, 'stripe');
        Test.stopTest();

        Map<String, Object> stripeResponseMap = (Map<String, Object>)JSON.deserializeUntyped(stripeResponse);

        System.assert((Boolean)stripeResponseMap.get('isSuccess'), stripeResponseMap.get('error'));

        Map<String, Object> stripeResultsMap = (Map<String, Object>)stripeResponseMap.get('results');
        System.assertEquals('stripeDisconnected', stripeResultsMap.get('isStripeConnected'));
    }

    @isTest
    static public void testSetupAssistant_salesforceNotConnected() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeNotConnectedMock());
        String salesforceResponse = setupAssistant.validateConnectionStatus(false, 'salesforce');
        Test.stopTest();

        Map<String, Object> salesforceResponseMap = (Map<String, Object>)JSON.deserializeUntyped(salesforceResponse);

        System.assert((Boolean)salesforceResponseMap.get('isSuccess'), salesforceResponseMap.get('error'));

        Map<String, Object> salesforceResultsMap = (Map<String, Object>)salesforceResponseMap.get('results');
        System.assertEquals('salesforceDisconnected', salesforceResultsMap.get('isSalesforceConnected'));
    }

    /*
    this is the case there is no 'connected' custom setting record stored
    in this case we generate an api key -> send it the ruby post-install endpoint -> then create the record
    */
    @isTest
    static public void testSetupAssistant_generateApiKey() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);

        scheduleSyncRecordDeletion.abortStripeCronJobs();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());
        String response = setupAssistant.validateConnectionStatus(true, '');
        Test.stopTest();

        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));
        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');

        Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();
        System.assertNotEquals(stripeConnectRec.Id, null);
        System.assertNotEquals(stripeConnectRec.API_Key__c, null);
    }

    @isTest
    static public void testSetupAssistant_getPicklistValuesForMapperConnectedCallback() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        String response = setupAssistant.getPicklistValuesForMapper(true, '');
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        Map<String, Object> fieldListByObjectMap = (Map<String, Object>)resultsMap.get('fieldListByObjectMap');
        List<Object> Order = (List<Object>)fieldListByObjectMap.get('Order');

        System.assertNotEquals(resultsMap, null);
        System.assertNotEquals(fieldListByObjectMap, null);
        System.assertNotEquals(Order, null);
    }

    @isTest
    static public void testSetupAssistant_getPicklistValuesForMapper() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        String response = setupAssistant.getPicklistValuesForMapper(false, 'Account');
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');

        String ObjectApiName = (String)resultsMap.get('ObjectApiName');
        List<Object> listOfObjectFields = (List<Object>)resultsMap.get('listOfObjectFields');
        System.assertEquals(ObjectApiName, 'Account');
        System.assertNotEquals(listOfObjectFields, null);
        System.assertNotEquals(resultsMap, null);
    }

    @isTest
    static public void testSetupAssistant_saveMappingConfigurations() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeSuccessMock());
        String response = setupAssistant.saveMappingConfigurations('test');
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals((Boolean)resultsMap.get('isConfigSaved'), true);
    }

    @isTest
    static public void testSetupAssistant_failSaveMappingConfigurations() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectFailMock());
        String response = setupAssistant.saveMappingConfigurations('test');
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assertEquals((Boolean)responseMap.get('isSuccess'), true);

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals((Boolean)resultsMap.get('isConfigSaved'), false);
    }

    @isTest
    static public void testSetupAssistant_getMappingConfigurations() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeSuccessMock());
        String response = setupAssistant.getMappingConfigurations();
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        Map<String, Object> allMappingConfigurations = (Map<String, Object>)resultsMap.get('allMappingConfigurations');
        Map<String, Object> default_mappings = (Map<String, Object>)allMappingConfigurations.get('default_mappings');
        Map<String, Object> field_defaults = (Map<String, Object>)allMappingConfigurations.get('field_defaults');
        Map<String, Object> field_mappings = (Map<String, Object>)allMappingConfigurations.get('field_mappings');
        Map<String, Object> required_mappings = (Map<String, Object>)allMappingConfigurations.get('required_mappings');

        System.assertNotEquals(allMappingConfigurations, null);
        System.assertNotEquals(default_mappings, null);
        System.assertNotEquals(field_defaults, null);
        System.assertNotEquals(field_mappings, null);
        System.assertNotEquals(required_mappings, null);
    }

    @isTest
    static public void testSetupAssistant_failGetMappingConfigurations() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectFailMock());
        String response = setupAssistant.getMappingConfigurations();
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assertEquals((Boolean)responseMap.get('isSuccess'), true);
    }

    @isTest
    static public void testSetupAssistant_getFormattedStripeObjectFields() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeSuccessMock());
        String response = setupAssistant.getFormattedStripeObjectFields();
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        List<Object> testCustomerFields =  (List<Object>)resultsMap.get('formattedStripeCustomerFields');
        List<Object> testProductFields =  (List<Object>)resultsMap.get('formattedStripeProductItemFields');
        List<Object> testSubscriptionFields =  (List<Object>)resultsMap.get('formattedStripeSubscriptionFields');
        List<Object> testSubscriptionItemFields =  (List<Object>)resultsMap.get('formattedStripeSubscriptionItemFields');
        List<Object> testPriceFields =  (List<Object>)resultsMap.get('formattedStripePriceFields');

        System.assertEquals(5, testCustomerFields.size());
        System.assertEquals(5, testProductFields.size());
        System.assertEquals(5, testSubscriptionFields.size());
        System.assertEquals(5, testSubscriptionItemFields.size());
        System.assertEquals(5, testPriceFields.size());
    }


    @isTest
    static public void testSetupAssistant_getSyncPreferences() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());
        String response = setupAssistant.getSyncPreferences();
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();
        System.debug(responseMap);
        System.assert((Boolean)responseMap.get('isSuccess'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals('1234', (String)resultsMap.get('stripe_account_id'));
        System.assertEquals('test', (String)resultsMap.get('last_synced'));
        System.assertEquals('USD', (String)resultsMap.get('default_currency'));
        System.assertEquals('2 days', (String)resultsMap.get('sync_record_retention'));
        System.assertEquals('yesterday', (String)resultsMap.get('sync_start_date'));
        System.assertEquals('1000515', (String)resultsMap.get('api_percentage_limit'));
        System.assertEquals('monthly', (String)resultsMap.get('cpq_term_unit'));
    }

 @isTest
    static public void testSetupAssistant_getMulticurrencySelectionOptions() {
        Test.startTest();
        String response = setupAssistant.getMulticurrencySelectionOptions();
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assert((Boolean)responseMap.get('isSuccess'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
         Boolean multiCurrencyEnabled = UserInfo.isMultiCurrencyOrganization();

        if (multiCurrencyEnabled) {
            System.assertEquals(true, (Boolean)resultsMap.get('isMultiCurrencyEnabled'));
        } else {
            System.assertEquals(false, (Boolean)resultsMap.get('isMultiCurrencyEnabled'));
        }
    }

    @isTest
    static public void testSetupAssistant_manualRetry() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Sync_Record__c testSyncRecord = new Sync_Record__c (
            Primary_Record_ID__c = '1234',
            Primary_Object_Type__c = 'Order',
            Secondary_Record_ID__c = '4321',
            Secondary_Object_Type__c = 'OrderItem',
            Resolution_Status__c = 'Error',
            Compound_ID__c = '12344321'
        );
        insert testSyncRecord;

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeSuccessMock());
        String response = setupAssistant.manualRetry(testSyncRecord.Id);
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

    }

    @isTest
    static public void testSetupAssistant_saveSyncPreferences() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());
        String response = setupAssistant.saveSyncPreferences('TEST','10000', '1644854972', '95', 'monthly');
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));
    }

    @isTest
    static public void testSetupAssistant_getFilterSettings() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());
        String response = setupAssistant.getFilterSettings();
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();
        System.debug(responseMap);
        System.assert((Boolean)responseMap.get('isSuccess'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');

        System.assertEquals('testOrderFilter', (String)resultsMap.get('Order'));
        System.assertEquals('testAccountFilter', (String)resultsMap.get('Account'));
        System.assertEquals('testProduct2Filter', (String)resultsMap.get('Product2'));
    }

    @isTest
    static public void testSetupAssistant_failSaveFilterSettings() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());
        String response = setupAssistant.saveFilterSettings('Id != null', 'Ids != null', '');
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals(false, (Boolean)resultsMap.get('isFiltersSaved'));
        System.assertEquals(true, (Boolean)resultsMap.get('isValidationError'));
    }

    @isTest
    static public void testSetupAssistant_successfulSaveFilterSettings() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());
        String response = setupAssistant.saveFilterSettings('Id != null', 'Id != null', 'Id != null');
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals(true, (Boolean)resultsMap.get('isFiltersSaved'));
        System.assertEquals(false, (Boolean)resultsMap.get('isValidationError'));
    }

    @isTest
    static public void testSetupAssistant_syncAllRecords() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeSuccessMock());
        String response = setupAssistant.syncAllRecords('Product2');
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals(true, (Boolean)resultsMap.get('syncAllRecordsDispatched'));
    }

    public static Setup_Connection_Data__mdt getTestStripeConnectionKey() {
        Setup_Connection_Data__mdt testSetupData = new Setup_Connection_Data__mdt();
        Map<String,Object> fields =  (Map<String,Object>) JSON.deserializeUntyped(JSON.serializePretty(testSetupData));

        fields.put(Setup_Connection_Data__mdt.Label.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.DeveloperName.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.Global_Key__c.getDescribe().getName(), 'testGlobalGUID');
        testSetupData = (Setup_Connection_Data__mdt) JSON.deserialize(JSON.serialize(fields), Setup_Connection_Data__mdt.class);

        return testSetupData;
    }

    public static void setTestGlobalKey(Setup_Connection_Data__mdt testSetupData) {
        List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Global_Key__c
                                                            FROM Setup_Connection_Data__mdt];
        if(!setupConfigList.isEmpty()) {
            utilities.setupConfigMetadata[0] = testSetupData;
            setupAssistant.setupConfigMetadata[0] = testSetupData;
        } else {
            utilities.setupConfigMetadata.add(testSetupData);
            setupAssistant.setupConfigMetadata.add(testSetupData);
        }
    }

    public static void insertTestConnectedRecord() {
        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c();
        stripeConnectRec.Stripe_Connected__c = true;
        stripeConnectRec.Salesforce_Connected__c = true;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        insert stripeConnectRec;
    }

    @isTest
    static public void testSetupAssistant_checkUserPermissions() {

        Test.startTest();
        String response = setupAssistant.checkUserPermissions();
        Test.stopTest();

        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');

        Map<String, Object> permissionIssueMap = (Map<String, Object>)resultsMap.get('permissionIssueMap');

        Boolean isIntegrationPermSetAssigned = [SELECT count() FROM PermissionSetAssignment
                                                WHERE AssigneeId = :Userinfo.getUserId()
                                                AND PermissionSet.Name = 'Stripe_Connector_Integration_User'
                                                WITH SECURITY_ENFORCED
                                                LIMIT 1] > 0;
        System.assertEquals(isIntegrationPermSetAssigned, (Boolean)permissionIssueMap.get('isPermSetAssigned'));
        if((Boolean)permissionIssueMap.get('isObjectPermissionMissing')) {
            System.assertNotEquals(null, permissionIssueMap.get('missingObjectPermissionList'));
        }
    }

    private class stripeConnectionSuccessMock implements HttpCalloutMock {
        public HttpResponse respond(HttpRequest req) {
            HttpResponse res = new HttpResponse();

            System.assertEquals(req.getHeader('Salesforce-Account-Id'), (String)constants.ORG_ID);
            System.assertEquals(req.getHeader('Salesforce-Type'), String.valueOf(Sentry_Environment.getInstanceType()));
            System.assertEquals(req.getHeader('Salesforce-Package-Namespace'), constants.NAMESPACE);
            if(!String.isEmpty(constants.NAMESPACE_API)) {
                System.assertEquals(req.getHeader('Package-Id'), (String)utilities.getPackageVersion());
            }

            if(req.getEndpoint().endsWith('configuration')) {
                Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();
                System.assertEquals(req.getHeader('Salesforce-Key'), (String)stripeConnectRec.API_Key__c);
                Map<String,Object> responseBody = new Map<String,Object>();
                responseBody.put('connection_status', (Object)new Map<String,Object>{
                    'salesforce' => true,
                    'stripe' => true,
                    'last_synced' => 'test',
                    'stripe_account_id' => 1234
                });
                responseBody.put('settings', (Object)new Map<String,Object>{
                    'default_currency' => 'USD',
                    'sync_record_retention' => '2 days',
                    'sync_start_date' => 'yesterday',
                    'api_percentage_limit' => 1000515,
                    'cpq_term_unit' => 'monthly',
                    'filters' => new Map<String,Object> {
                        'Product2' => 'testProduct2Filter',
                        'Order' => 'testOrderFilter',
                        'Account' => 'testAccountFilter'
                    }
                });

                res.setBody(Json.serialize(responseBody));
                res.setStatusCode(200);
            }
            if(req.getEndpoint().endsWith('post-install')) {
                List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Global_Key__c
                                                                    FROM Setup_Connection_Data__mdt];
                if(!setupConfigList.isEmpty()) {
                    System.assertEquals(req.getHeader('Salesforce-Key'), setupConfigList[0].Global_Key__c);
                } else {
                    System.assertEquals(req.getHeader('Salesforce-Key'), 'testGlobalGUID');
                }

                res.setStatusCode(200);
            }
            return res;
        }
    }

        private class stripeSuccessMock implements HttpCalloutMock {
            public HttpResponse respond(HttpRequest req) {
                HttpResponse res = new HttpResponse();
                System.assertEquals(req.getHeader('Salesforce-Account-Id'), (String)constants.ORG_ID);
                System.assertEquals(req.getHeader('Salesforce-Type'), String.valueOf(Sentry_Environment.getInstanceType()));
                System.assertEquals(req.getHeader('Salesforce-Package-Namespace'), constants.NAMESPACE);
                if(!String.isEmpty(constants.NAMESPACE_API)) {
                    System.assertEquals(req.getHeader('Package-Id'), (String)utilities.getPackageVersion());
                }
                Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();


                if (req.getEndpoint().endsWith('configuration')) {
                    System.assertEquals(req.getHeader('Salesforce-Key'), (String)stripeConnectRec.API_Key__c);
                    Map<String,Object> fieldMap = new Map<String,Object>();
                    Map<String,Object> testFieldMap = new Map<String,Object>();
                    testFieldMap.put('testFieldMapping', 'test');
                    fieldMap.put('field_mappings', (Object)testFieldMap);

                    Map<String,Object> defaultMap = new Map<String,Object>();
                    Map<String,Object> testdefaultMap = new Map<String,Object>();
                    testdefaultMap.put('testDefaultMap', 'test');
                    defaultMap.put('default_mappings',(Object)testdefaultMap);

                    Map<String,Object> requiredMapping = new Map<String,Object>();
                    Map<String,Object> testRequiredMapping = new Map<String,Object>();
                    testRequiredMapping.put('testRequiredMapping', 'test');
                    requiredMapping.put('required_mappings', (Object)testRequiredMapping);

                    Map<String,Object> fieldDefaultMap = new Map<String,Object>();
                    Map<String,Object> testFieldDefaultMap = new Map<String,Object>();
                    testFieldDefaultMap.put('testFieldDefaultMap', 'test');
                    fieldDefaultMap.put('field_defaults', (Object)testFieldDefaultMap);

                    Map<String,Object> responseBody = new Map<String,Object>();
                    responseBody.put('field_mappings', (Object)fieldMap);
                    responseBody.put('default_mappings', (Object)defaultMap);
                    responseBody.put('required_mappings', (Object)requiredMapping);
                    responseBody.put('field_defaults', (Object)fieldDefaultMap);

                    res.setBody(Json.serialize(responseBody));
                    res.setStatusCode(200);
                } else if(req.getEndpoint().endsWith('translate')) {
                    System.assertEquals(req.getHeader('Salesforce-Key'), (String)stripeConnectRec.API_Key__c);
                    Map<String,Object> responseObject = new Map<String,Object>();

                    Map<String,Object> responseBody = new Map<String,Object>();
                    responseBody.put('retry_status', (Object)responseObject);

                    res.setBody(Json.serialize(responseBody));
                    res.setStatusCode(200);
                } else if(req.getEndpoint().endsWith('openapi.json')) {
                    System.assertEquals(req.getHeader('Salesforce-Key'), '');
                    List<Object> testFormattedFieldsList = new List<Object> {
                        'testField1',
                        'testField2',
                        'testField3',
                        'testField4',
                        'testField5'
                    };
                    Map<String,Object> responseBody = new Map<String,Object>();
                    responseBody.put('formattedStripeCustomerFields', (Object)testFormattedFieldsList);
                    responseBody.put('formattedStripeProductItemFields', (Object)testFormattedFieldsList);
                    responseBody.put('formattedStripeSubscriptionFields', (Object)testFormattedFieldsList);
                    responseBody.put('formattedStripeSubscriptionItemFields', (Object)testFormattedFieldsList);
                    responseBody.put('formattedStripePriceFields', (Object)testFormattedFieldsList);
                    res.setBody(Json.serialize(responseBody));
                    res.setStatusCode(200);
                } else if(req.getEndpoint().endsWith('translate_all')) {
                    res.setStatusCode(200);
                }

                return res;
            }
        }

        private class stripeNotConnectedMock implements HttpCalloutMock {
            public HttpResponse respond(HttpRequest req) {
                HttpResponse res = new HttpResponse();
                if(req.getEndpoint().endsWith('configuration')) {

                    Map<String,Object> responseBody = new Map<String,Object>();
                    responseBody.put('connection_status', (Object)new Map<String,Object>{
                        'salesforce' => false,
                        'stripe' => false,
                        'last_synced' => 'test',
                        'stripe_account_id' => 1234
                    });
                    responseBody.put('settings', (Object)new Map<String,Object>{
                        'default_currency' => 'USD',
                        'sync_record_retention' => '2 days',
                        'sync_start_date' => 'yesterday',
                        'api_percentage_limit' => 1000515,
                        'cpq_term_unit' => 'monthly',
                        'filters' => new Map<String,Object> {
                            'Product2' => 'testProduct2Filter',
                            'Order' => 'testOrderFilter',
                            'Account' => 'testAccountFilter'
                        }
                    });

                    res.setBody(Json.serialize(responseBody));
                    res.setStatusCode(200);
                }
                return res;
            }
        }

        private class stripeConnectFailMock implements HttpCalloutMock {
            public HttpResponse respond(HttpRequest req) {
                HttpResponse res = new HttpResponse();

                Map<String,Object> responseBody = new Map<String,Object>();
                res.setBody(Json.serialize(responseBody));
                res.setStatusCode(500);
                return res;
            }
        }
}


