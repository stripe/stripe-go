@isTest
public with sharing class test_setupAssistant {
    static Debug_Helper debugger = new Debug_Helper(true);
    public static final String TEST_STRIPE_ACCOUNT_ID = 'acct_12345';

    @IsTest
    static void testMaintenanceUtilitiesResetServiceConnection() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        stripeConnectionSuccessMock mock = new stripeConnectionSuccessMock();
        Test.setMock(HttpCalloutMock.class, mock);
        Test.startTest();
        maintenanceUtilities.resetServiceConnection();
        Test.stopTest();

        System.assertEquals(1, mock.calls);
        Stripe_Connection__c stripeConnectRec = utilities.getStripeConnection();
        System.assertNotEquals(null, stripeConnectRec.API_Key__c);
    }

    @IsTest
    static void testMaintenanceUtilitiesResetEnvironment() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        testSetupData.Enable_Apex_Debug__c = true;
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();
        Setup_Data__c setupData = new Setup_Data__c(
                isSetupComplete__c = false,
                Steps_Completed__c = '{"foo": 1}'
        );
        insert setupData;

        Test.startTest();
        maintenanceUtilities.resetEnvironment();
        maintenanceUtilities.resetUserLandData();
        maintenanceUtilities.resetSetupData();
        maintenanceUtilities.resetConnection();
        Test.stopTest();

        Stripe_Connection__c stripeConnectRec = utilities.getStripeConnection();
        System.assertEquals(null, stripeConnectRec.Id);
        System.assertEquals(null, stripeConnectRec.API_Key__c);
        List<Setup_Data__c> setupDataCheck = [SELECT Id FROM Setup_Data__c];
        System.assertEquals(0, setupDataCheck.size());
        List<Stripe_Account__c> stripeAccountCheck = [SELECT Id FROM Stripe_Account__c];
        System.assertEquals(0, stripeAccountCheck.size());
    }

    @IsTest
    static void testSetStripeAccountAsDefault() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        AccountManagementMock mock = new AccountManagementMock();
        Test.setMock(HttpCalloutMock.class, mock);
        mock.nextCallFails = true;
        Exception ex;

        Test.startTest();
        Boolean checkFail;
        try {
            checkFail = setupAssistant.setStripeAccountAsDefault(TEST_STRIPE_ACCOUNT_ID, false);
        } catch (Exception e) {
            ex = e;
        }
        Boolean checkPass = setupAssistant.setStripeAccountAsDefault(TEST_STRIPE_ACCOUNT_ID, false);
        Test.stopTest();

        System.assertEquals(null, checkFail);
        System.assertNotEquals(null, ex);
        System.assertEquals(true, checkPass);
    }

    @IsTest
    static void testDeleteStripeAccount() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        AccountManagementMock mock = new AccountManagementMock();
        Test.setMock(HttpCalloutMock.class, mock);
        mock.nextCallFails = true;
        Exception ex;

        Test.startTest();
        Boolean checkFail;
        try {
            checkFail = setupAssistant.deleteStripeAccount(TEST_STRIPE_ACCOUNT_ID, false);
        } catch (Exception e) {
            ex = e;
        }
        Boolean checkPass = setupAssistant.deleteStripeAccount(TEST_STRIPE_ACCOUNT_ID, false);
        Test.stopTest();

        System.assertEquals(null, checkFail);
        System.assertNotEquals(null, ex);
        System.assertEquals(true, checkPass);
    }

    @IsTest
    static void testAddStripeAccount() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        StripeOAuthState state = StripeOAuthState.create();
        state.user_id = '1';
        state.csac_id = '1';
        state.stripe_account_id = 'acct_234567';
        state.stripe_account_livemode = 'test';
        String stateStr = state.toString();

        Test.startTest();
        Stripe_Account__c acct = setupAssistant.addStripeAccount(stateStr);
        Test.stopTest();

        System.assertNotEquals(null, acct);
        System.assertEquals(state.stripe_account_id, acct.Stripe_ID__c);
        System.assertEquals(false, acct.Is_Live_Mode__c);
        System.assertEquals(false, acct.Is_Primary__c);
        String extId = state.stripe_account_id + '|' + state.stripe_account_livemode;
        System.assertEquals(extId, acct.External_ID__c);
    }

    @IsTest
    static void testGetCoreData() {
        String response = setupAssistant.getCoreData();
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assertEquals(true, responseMap.get('isSuccess'));
        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals(true, resultsMap.containsKey('packageVersion'), 'Has package key');
        System.assertEquals(true, resultsMap.containsKey('isCpqInstalled'), 'Has CPQ key');
        System.assertEquals(true, resultsMap.containsKey('configDownloadUri'), 'Has config download key');
        System.assertEquals(true, resultsMap.containsKey('isMultiCurrencyEnabled'), 'Has multicurrency key');
        System.assertEquals(true, resultsMap.containsKey('isSandbox'), 'Has sandbox key');
        System.assertEquals(true, resultsMap.containsKey('fieldListByObjectMap'), 'Has fieldListByObjectMap key');
        System.assertEquals(true, resultsMap.containsKey('supportedISOCodes'), 'Has supportedISOCodes key');
        System.assertEquals(true, resultsMap.containsKey('defaultCurrency'), 'Has defaultCurrency key');
        System.assertEquals(true, resultsMap.containsKey('orgMaxApiLimit'), 'Has orgMaxApiLimit key');
    }

    @IsTest
    static void testGetTranslationConfig() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        testSetupData.Enable_Apex_Debug__c = true;
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());
        Test.startTest();

        String saveResp = setupAssistant.getTranslationConfig();

        Test.stopTest();

        debugger.debug('testSaveTranslationConfig', saveResp);
        Map<String, Object> saveRespMap = (Map<String, Object>)JSON.deserializeUntyped(saveResp);
        System.assertEquals(true, saveRespMap.get('isSuccess'));
        Map<String, Object> resultsMap = (Map<String, Object>)saveRespMap.get('results');
        System.assertEquals('1234', resultsMap.get('configurationHash'));
    }

    @IsTest
    static void testSaveTranslationConfig() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        testSetupData.Enable_Apex_Debug__c = true;
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.setMock(HttpCalloutMock.class, new stripeSuccessMock());
        Test.startTest();

        String config = '{"configurationHash": "1234", "pricebook_entry_filter":null,"account_filter":null,"order_filter":"Status = \'Activated\'","product_filter":null,"allMappingConfigurations":{"field_defaults":{"subscription_schedule":{},"subscription_phase":{},"subscription_item":{},"price_order_item":{},"customer":{"balance":"300"},"product":{},"coupon":{},"price":{}},"field_mappings":{"subscription_schedule":{},"subscription_phase":{},"subscription_item":{},"price_order_item":{},"customer":{},"product":{},"coupon":{},"price":{}}},"polling_enabled":false,"cpq_term_unit":"month","api_percentage_limit":"95","sync_start_date":null,"sync_record_retention":"10000"}';
        String saveResp = setupAssistant.saveTranslationConfig(config);

        Test.stopTest();

        debugger.debug('testSaveTranslationConfig', saveResp);
        Map<String, Object> saveRespMap = (Map<String, Object>)JSON.deserializeUntyped(saveResp);
        System.assertEquals(true, saveRespMap.get('isSuccess'));
        Map<String, Object> resultsMap = (Map<String, Object>)saveRespMap.get('results');
        System.assertEquals(true, resultsMap.get('isConfigSaved'));
        System.assertEquals(true, resultsMap.get('isFiltersSaved'));
        System.assertEquals(false, resultsMap.get('isValidationError'));
    }

    @IsTest
    static void testIsJSDebugEnabled() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        testSetupData.Enable_JS_Debug__c = false;
        setTestGlobalKey(testSetupData);

        Boolean isJSDebugEnabledFalse = setupAssistant.isJSDebugEnabled();
        testSetupData.Enable_JS_Debug__c = true;
        Boolean isJSDebugEnabledTrue = setupAssistant.isJSDebugEnabled();

        System.assertEquals(true, isJSDebugEnabledTrue, 'correctly relays true');
        System.assertEquals(false, isJSDebugEnabledFalse, 'correctly relays false');
    }

    @IsTest
    static void testIsOAuthAutoCloseWindowEnabled() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        testSetupData.OAuth_AutoClose_Window__c = false;
        setTestGlobalKey(testSetupData);

        Boolean isOAuthAutoCloseWindowEnabledFalse = setupAssistant.isOAuthAutoCloseWindowEnabled();
        testSetupData.OAuth_AutoClose_Window__c = true;
        Boolean isOAuthAutoCloseWindowEnabledTrue = setupAssistant.isOAuthAutoCloseWindowEnabled();

        System.assertEquals(true, isOAuthAutoCloseWindowEnabledTrue, 'correctly relays true');
        System.assertEquals(false, isOAuthAutoCloseWindowEnabledFalse, 'correctly relays false');
    }

    @IsTest
    static void testGetDownloadForSupportUrl() {
        Id recId = UserInfo.getUserId();
        String url = setupAssistant.getDownloadForSupportUrl(recId);
        System.assert(url.contains('recordId=' + recId), 'url contains recordId set to ' + recId);
    }

    @isTest
    static public void testSetupAssistant_getSetupData() {
        setupAssistant.saveData(new Map<String, Object>(), false);
        insertTestConnectedRecord();

        String response = setupAssistant.getSetupData();

        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals(true, resultsMap.get('isConnected'));
        System.assert(resultsMap.containsKey('setupData'), '`data` not found');

        Map<String, Object> setupDataMap = (Map<String, Object>)resultsMap.get('setupData');
        System.assertEquals(false, setupDataMap.get('isSetupComplete__c'));
    }

    @isTest
    static public void testSetupAssistant_getSetupDataWithNoDefault() {
        String response = setupAssistant.getSetupData();

        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assert(resultsMap.containsKey('setupData'), '`data` not found');

        Map<String, Object> setupDataMap = (Map<String, Object>)resultsMap.get('setupData');
        System.assert(setupDataMap.containsKey('Steps_Completed__c'), '`data` not found');
        System.assertEquals(false, setupDataMap.get('isSetupComplete__c'));
    }

    @isTest
    static public void testSetupAssistant_getSetupDataWithData() {
        setupAssistant.saveData(new Map<String, Object>(), false);
        Map<String, Object> data = new Map<String, Object> {
                'Steps_Completed__c' => '{"foo": 1}'
        };
        setupAssistant.saveData(data, true);
        insertTestConnectedRecord();

        String response = setupAssistant.getSetupData();

        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals(true, resultsMap.get('isConnected'));
        System.assert(resultsMap.containsKey('setupData'), '`data` not found');

        Map<String, Object> setupDataMap = (Map<String, Object>)resultsMap.get('setupData');
        System.assertEquals(true, setupDataMap.get('isSetupComplete__c'));
    }

    @isTest
    static public void testUtilities_isCpqEnabled() {
        Boolean isCpqEnabled = utilities.isCpqEnabled();
        // TODO should change null to explicit boolean check based on org ID
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
        String connectedCallbackResponse = setupAssistant.validateConnectionStatus(null);
        Test.stopTest();

        Map<String, Object> connectedCallbackResponseMap = (Map<String, Object>)JSON.deserializeUntyped(connectedCallbackResponse);
        System.assert((Boolean)connectedCallbackResponseMap.get('isSuccess'), connectedCallbackResponseMap.get('error'));

        Map<String, Object> connectecCallbackResultsMap = (Map<String, Object>)connectedCallbackResponseMap.get('results');
        System.assertEquals(setupAssistant.STATUS_CONNECTED, connectecCallbackResultsMap.get(setupAssistant.KEY_STRIPE));
        System.assertEquals(setupAssistant.STATUS_CONNECTED, connectecCallbackResultsMap.get(setupAssistant.KEY_SALESFORCE));
    }

    @isTest
    static public void testSetupAssistant_connected() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());
        String response = setupAssistant.validateConnectionStatus(null);
        Test.stopTest();

        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));
        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals(setupAssistant.STATUS_CONNECTED, resultsMap.get(setupAssistant.KEY_STRIPE));
        System.assertEquals(setupAssistant.STATUS_CONNECTED, resultsMap.get(setupAssistant.KEY_SALESFORCE));
    }

    @isTest
    static public void testSetupAssistant_stripeNotConnected() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeNotConnectedMock());
        String stripeResponse = setupAssistant.validateConnectionStatus('stripe');
        Test.stopTest();

        Map<String, Object> stripeResponseMap = (Map<String, Object>)JSON.deserializeUntyped(stripeResponse);

        System.assert((Boolean)stripeResponseMap.get('isSuccess'), stripeResponseMap.get('error'));

        Map<String, Object> stripeResultsMap = (Map<String, Object>)stripeResponseMap.get('results');
        System.assertEquals(setupAssistant.STATUS_DISCONNECTED, stripeResultsMap.get(setupAssistant.KEY_STRIPE));
    }

    @isTest
    static public void testSetupAssistant_salesforceNotConnected() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeNotConnectedMock());
        String salesforceResponse = setupAssistant.validateConnectionStatus('salesforce');
        Test.stopTest();

        Map<String, Object> salesforceResponseMap = (Map<String, Object>)JSON.deserializeUntyped(salesforceResponse);

        System.assert((Boolean)salesforceResponseMap.get('isSuccess'), salesforceResponseMap.get('error'));

        Map<String, Object> salesforceResultsMap = (Map<String, Object>)salesforceResponseMap.get('results');
        System.assertEquals(setupAssistant.STATUS_DISCONNECTED, salesforceResultsMap.get(setupAssistant.KEY_SALESFORCE));
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
        String response = setupAssistant.validateConnectionStatus(null);
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

        System.assertEquals((Boolean)responseMap.get('isSuccess'), false);
    }

    @IsTest
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
        String configurationHash = (String)resultsMap.get('configurationHash');

        System.assertNotEquals(allMappingConfigurations, null);
        System.assertNotEquals(default_mappings, null);
        System.assertNotEquals(field_defaults, null);
        System.assertNotEquals(field_mappings, null);
        System.assertNotEquals(required_mappings, null);
        System.assertEquals(configurationHash, '1234');
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
        OpenAPIMock callMock = new OpenAPIMock();
        Test.setMock(HttpCalloutMock.class, callMock);

        Test.startTest();
        String response = setupAssistant.getFormattedStripeObjectFields();
        System.assertEquals(1, callMock.calls);

        Test.stopTest();

        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        List<Object> testCustomerFields =  (List<Object>)resultsMap.get('formattedStripeCustomerFields');
        List<Object> testProductFields =  (List<Object>)resultsMap.get('formattedStripeProductItemFields');
        List<Object> testSubscriptionFields =  (List<Object>)resultsMap.get('formattedStripeSubscriptionFields');
        List<Object> testSubscriptionPhaseFields =  (List<Object>)resultsMap.get('formattedStripeSubscriptionSchedulePhaseFields'); 
        List<Object> testSubscriptionItemFields =  (List<Object>)resultsMap.get('formattedStripeSubscriptionItemFields');
        List<Object> testPriceFields =  (List<Object>)resultsMap.get('formattedStripePriceFields');

        Map<String, List<Integer>> assertions = new Map<String, List<Integer>>{
            'testCustomerFields' => new List<Integer>{6, testCustomerFields.size()},
            'testProductFields' => new List<Integer>{4, testProductFields.size()},
            'testSubscriptionFields' => new List<Integer>{7, testSubscriptionFields.size()},
            'testSubscriptionPhaseFields' => new List<Integer>{5, testSubscriptionPhaseFields.size()},
            'testSubscriptionItemFields' => new List<Integer>{1, testSubscriptionItemFields.size()},
            'testPriceFields' => new List<Integer>{3, testPriceFields.size()}
        };

        // I know this looks dumb but so does fixing these one at a time...
        Map<String, Integer> expect = new Map<String, Integer>{};
        Map<String, Integer> actual = new Map<String, Integer>{};
        for (String key : assertions.keySet()) {
            expect.put(key, assertions.get(key)[0]);
            actual.put(key, assertions.get(key)[1]);
        }
        System.assertEquals(expect, actual);
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
        System.assertEquals('1', (String)resultsMap.get('last_synced'));
        System.assertEquals('USD', (String)resultsMap.get('default_currency'));
        System.assertEquals('2 days', (String)resultsMap.get('sync_record_retention'));
        System.assertEquals('yesterday', (String)resultsMap.get('sync_start_date'));
        System.assertEquals('month', (String)resultsMap.get('cpq_prorate_precision'));
        System.assertEquals(true, (Boolean)resultsMap.get('enabled'));
        System.assertEquals('1234', (String)resultsMap.get('configurationHash'));
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
    static public void testSetupAssistant_manualTranslation() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();
        
        Account testAccountRecord = new Account (
            Name = 'TestAccount'
        );
        insert testAccountRecord;

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeSuccessMock());
        String response = setupAssistant.manualTranslation(string.valueOf(testAccountRecord.Id));
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));
    }

    @isTest
    static public void testSetupAssistant_manualTranslationFail() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();
        
        Opportunity testOpportunityRecord = new Opportunity (
            Name = 'testOpportunity',
            CloseDate = system.today() + 4,
            StageName = 'prospecting'
        );
        insert testOpportunityRecord;

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeSuccessMock());
        String unsupportedObjectResponse = setupAssistant.manualTranslation(string.valueOf(testOpportunityRecord.Id));
        Map<String, Object> unsupportedObjectResponseMap = (Map<String, Object>)JSON.deserializeUntyped(unsupportedObjectResponse);
        
        String fakeAccountRecordResponse = setupAssistant.manualTranslation('001DR00001MM9tsZZZ');
        Map<String, Object> fakeAccountRecordResponseMap = (Map<String, Object>)JSON.deserializeUntyped(fakeAccountRecordResponse);
        Test.stopTest();

        System.assert((Boolean)unsupportedObjectResponseMap.get('isSuccess'), unsupportedObjectResponseMap.get('error'));

        Map<String, Object> unsupportedObjectResultsMap = (Map<String, Object>)unsupportedObjectResponseMap.get('results');
        System.assertEquals('this is an id for a record on the Opportunity object which is not suppoted, full list of supported objects: (Account, Order, Pricebook2, Product2)', (String)unsupportedObjectResultsMap.get('errorMessage'));

        System.assert((Boolean)fakeAccountRecordResponseMap.get('isSuccess'), fakeAccountRecordResponseMap.get('error'));
        Map<String, Object> fakeAccountRecordResultsMap = (Map<String, Object>)fakeAccountRecordResponseMap.get('results');
        System.assertEquals('No Account objects found for this Salesforce record Id', (String)fakeAccountRecordResultsMap.get('errorMessage'));
    }

    @isTest
    static public void testSetupAssistant_saveSyncPreferences() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());
        String response = setupAssistant.saveSyncPreferences('TEST','10000', '1644854972', '95', 'monthly', 'month', true, '1234');
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));
    }

    @isTest
    static public void testSetupAssistant_activatePolling() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());
        String response = setupAssistant.activatePolling('1644854972', true, '1234');
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

        System.assertEquals('testOrderFilter', (String)resultsMap.get('order_filter'));
        System.assertEquals('testAccountFilter', (String)resultsMap.get('account_filter'));
        System.assertEquals('testProduct2Filter', (String)resultsMap.get('product_filter'));
        System.assertEquals('testPricebookEntryFilter', (String)resultsMap.get('pricebook_entry_filter'));
    }

    @isTest
    static public void testSetupAssistant_failSaveFilterSettings() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());
        String response = setupAssistant.saveFilterSettings('Id != null', 'Ids != null', '', '');
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
        String response = setupAssistant.saveFilterSettings('Id != null', 'Id != null', 'Id != null', 'Id != null');
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
        Test.stopTest();

        debugger.debug('testSetupAssistant_syncAllRecords', 'got test response: ' + response);
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);

        debugger.debug('testSetupAssistant_syncAllRecords', 'response map: ' + responseMap);
        System.assert((Boolean)responseMap.get('isSuccess'), 'Is success');
        System.assertEquals(null, responseMap.get('error'), 'No error provided');

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        debugger.debug('testSetupAssistant_syncAllRecords', 'results map: ' + resultsMap);
        System.assert((Boolean)resultsMap.get('isConnected'), 'Is connected');
        System.assertEquals(true, (Boolean)resultsMap.get('syncAllRecordsDispatched'), 'all records dispatched');
    }

    public static Setup_Connection_Data__mdt getTestStripeConnectionKey() {
        Setup_Connection_Data__mdt testSetupData = new Setup_Connection_Data__mdt();

        Map<String,Object> fields = (Map<String,Object>) JSON.deserializeUntyped(JSON.serializePretty(testSetupData));
        fields.put(Setup_Connection_Data__mdt.Label.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.DeveloperName.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.Global_Key__c.getDescribe().getName(), 'testGlobalGUID');
        fields.put(Setup_Connection_Data__mdt.Middleware_Infrastructure_Version__c.getDescribe().getName(), 'v1');
        fields.put(Setup_Connection_Data__mdt.OAuth_State_Signing_Key__c.getDescribe().getName(), EncodingUtil.convertToHex(Crypto.generateAesKey(256)));
        testSetupData = (Setup_Connection_Data__mdt) JSON.deserialize(JSON.serialize(fields), Setup_Connection_Data__mdt.class);

        return testSetupData;
    }

    public static void setTestGlobalKey(Setup_Connection_Data__mdt testSetupData) {
        if (utilities.setupConfigMetadata.isEmpty() == false) {
            utilities.setupConfigMetadata[0] = testSetupData;
        } else {
            utilities.setupConfigMetadata.add(testSetupData);
        }

        if (setupAssistant.setupConfigMetadata.isEmpty() == false) {
            setupAssistant.setupConfigMetadata[0] = testSetupData;
        } else {
            setupAssistant.setupConfigMetadata.add(testSetupData);
        }
    }

    public static void insertTestConnectedRecord() {
        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c();
        stripeConnectRec.Stripe_Connected__c = true;
        stripeConnectRec.Salesforce_Connected__c = true;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        stripeConnectRec.Sync_Record_Retention__c = '5000';
        insert stripeConnectRec;

        Stripe_Account__c stripeAccountRec = new Stripe_Account__c(
                Stripe_ID__c = TEST_STRIPE_ACCOUNT_ID,
                Connection_Status__c = constants.STRIPE_ACCOUNT_CONN_STATUS_CONNECTED,
                Is_Live_Mode__c = false,
                Is_Primary__c = true,
                External_ID__c =  TEST_STRIPE_ACCOUNT_ID + '|test'
        );
        insert stripeAccountRec;
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
        public Integer calls = 0;

        public HttpResponse respond(HttpRequest req) {
            this.calls++;
            HttpResponse res = new HttpResponse();

            System.assertEquals(req.getHeader('Salesforce-Account-Id'), (String)constants.ORG_ID);
            System.assertEquals(req.getHeader('Salesforce-Type'), String.valueOf(Sentry_Environment.getInstanceType()));
            System.assertEquals(req.getHeader('Salesforce-Package-Namespace'), constants.NAMESPACE);
            if(utilities.isPackagedEnvironment()) {
                System.assertEquals(req.getHeader('Salesforce-Package-Id'), (String)utilities.getPackageVersionString());
            }
            if(req.getEndpoint().endsWith('configuration')) {
                Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();
                System.assertEquals(req.getHeader('Salesforce-Key'), (String)stripeConnectRec.API_Key__c);
                Map<String,Object> responseBody = new Map<String,Object>{
                    'enabled' => true,
                    'configuration_hash' => '1234'
                };
                responseBody.put('connection_status', (Object)new Map<String,Object>{
                    'salesforce' => true,
                    'stripe' => true,
                    'last_synced' => '1',
                    'stripe_account_id' => 1234
                });
                responseBody.put('settings', (Object)new Map<String,Object>{
                    'default_currency' => 'USD',
                    'sync_record_retention' => '2 days',
                    'sync_start_date' => 'yesterday',
                    'api_percentage_limit' => 1000515,
                    'cpq_term_unit' => 'monthly',
                    'cpq_prorate_precision' => 'month',
                    'filters' => new Map<String,Object> {
                        'Product2' => 'testProduct2Filter',
                        'Order' => 'testOrderFilter',
                        'Account' => 'testAccountFilter',
                        'PricebookEntry' => 'testPricebookEntryFilter'
                    }
                });

                responseBody.put('hidden_sync_pref_fields', new List<String> {
                        'api_percentage_limit',
                        'cpq_term_unit'
                });

                res.setBody(Json.serialize(responseBody));
                res.setStatusCode(200);
            }
            if(req.getEndpoint().endsWith('post-install')) {
                if(!utilities.setupConfigMetadata.isEmpty()) {
                    System.assertEquals(req.getHeader('Salesforce-Key'), utilities.setupConfigMetadata[0].Global_Key__c);
                } else {
                    System.assertEquals(req.getHeader('Salesforce-Key'), 'testGlobalGUID');
                }

                res.setStatusCode(200);
            }
            return res;
        }
    }

    private class stripeSuccessMock implements HttpCalloutMock {
        public Integer calls = 0;
        public HttpResponse respond(HttpRequest req) {
            calls++;

            HttpResponse res = new HttpResponse();
            System.assertEquals(req.getHeader('Salesforce-Account-Id'), (String)constants.ORG_ID);
            System.assertEquals(req.getHeader('Salesforce-Type'), String.valueOf(Sentry_Environment.getInstanceType()));
            System.assertEquals(req.getHeader('Salesforce-Package-Namespace'), constants.NAMESPACE);
            if (utilities.isPackagedEnvironment()) {
                System.assertEquals(req.getHeader('Salesforce-Package-Id'), (String)utilities.getPackageVersionString());
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
                responseBody.put('configuration_hash', '1234');
                responseBody.put('field_mappings', (Object)fieldMap);
                responseBody.put('default_mappings', (Object)defaultMap);
                responseBody.put('required_mappings', (Object)requiredMapping);
                responseBody.put('field_defaults', (Object)fieldDefaultMap);

                res.setBody(JSON.serialize(responseBody));
                res.setStatusCode(200);
            } else if(req.getEndpoint().endsWith('translate')) {
                System.assertEquals(req.getHeader('Salesforce-Key'), (String)stripeConnectRec.API_Key__c);
                Map<String,Object> responseObject = new Map<String,Object>();

                Map<String,Object> responseBody = new Map<String,Object>();
                responseBody.put('retry_status', (Object)responseObject);

                res.setBody(JSON.serialize(responseBody));
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
                responseBody.put('formattedStripeSubscriptionSchedulePhaseFields', (Object)testFormattedFieldsList);
                responseBody.put('formattedStripePriceFields', (Object)testFormattedFieldsList);
                res.setBody(JSON.serialize(responseBody));
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

    private class OpenAPIMock implements HttpCalloutMock {
        public Integer calls = 0;
        public HttpResponse respond(HttpRequest req) {
            calls++;
            HttpResponse res = new HttpResponse();

            res.setBody('{"formattedStripeCustomerFields":[{"label":"Standard Mappings","name":"standard","description":"","fields":[{"name":"balance","value":"balance","description":"<p>An integer amount in cents (or local equivalent) that represents the customer\'s current balance, which affect the customer\'s future invoices. A negative amount represents a credit that decreases the amount due on an invoice; a positive amount increases the amount due on an invoice.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"coupon","value":"coupon","description":"<p>A coupon contains information about a percent-off or amount-off discount you\\nmight want to apply to a customer. Coupons may be applied to <a href=\\"https://stripe.com/docs/api#subscriptions\\">subscriptions</a>, <a href=\\"https://stripe.com/docs/api#invoices\\">invoices</a>,\\n<a href=\\"https://stripe.com/docs/api/checkout/sessions\\">checkout sessions</a>, <a href=\\"https://stripe.com/docs/api#quotes\\">quotes</a>, and more. Coupons do not work with conventional one-off <a href=\\"https://stripe.com/docs/api#create_charge\\">charges</a> or <a href=\\"https://stripe.com/docs/api/payment_intents\\">payment intents</a>.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"description","value":"description","description":"<p>An arbitrary string that you can _attach to a customer object. It is displayed alongside the customer in the dashboard.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"email","value":"email","description":"<p>Customer\'s email address. It\'s displayed alongside the customer in your dashboard and can be useful for searching and tracking. This may be up to <em>512 characters</em>.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"invoice prefix","value":"invoice_prefix","description":"<p>The prefix for the customer used to generate unique invoice numbers. Must be 3–12 uppercase letters or numbers.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"name","value":"name","description":"<p>The customer\'s full name or business name.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"next invoice sequence","value":"next_invoice_sequence","description":"<p>The sequence to be used on the customer\'s next invoice. Defaults to 1.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"payment method","value":"payment_method","description":"<p>ID of the payment method used in this charge.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"phone","value":"phone","description":"<p>The customer\'s phone number.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"promotion code","value":"promotion_code","description":"<p>The API ID of a promotion code to apply to the customer. The customer will have a discount applied on all recurring payments. Charges you create through the API will not have the discount.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"source","value":"source","description":"<p>The source object for errors returned on a request involving a source.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"tax exempt","value":"tax_exempt","description":"<p>The customer\'s tax exemption. One of <code>none</code>, <code>exempt</code>, or <code>reverse</code>.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":"","enum":["","exempt","none","reverse"]},{"name":"test clock","value":"test_clock","description":"<p>ID of the test clock to _attach to the customer.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Address","name":"address","description":"","fields":[{"name":"city","value":"address.city","description":"<p>City, district, suburb, town, or village.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"country","value":"address.country","description":"<p>Two-letter country code (<a href=\\"https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2\\">ISO 3166-1 alpha-2</a>).</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"line1","value":"address.line1","description":"<p>Address line 1 (e.g., street, PO Box, or company name).</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"line2","value":"address.line2","description":"<p>Address line 2 (e.g., apartment, suite, unit, or building).</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"postal code","value":"address.postal_code","description":"<p>ZIP or postal code.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"state","value":"address.state","description":"<p>State, county, province, or region.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Cash balance Settings","name":"cash_balance.settings","description":"","fields":[{"name":"reconciliation mode","value":"cash_balance.settings.reconciliation_mode","description":"<p>The configuration for how funds that land in the customer cash balance are reconciled.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Invoice settings","name":"invoice_settings","description":"","fields":[{"name":"default payment method","value":"invoice_settings.default_payment_method","description":"<p>ID of the default payment method for the invoice. It must belong to the customer associated with the invoice. If not set, defaults to the subscription\'s default payment method, if any, or to the default payment method in the customer\'s invoice settings.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"footer","value":"invoice_settings.footer","description":"<p>Footer displayed on the invoice.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Shipping","name":"shipping","description":"","fields":[{"name":"carrier","value":"shipping.carrier","description":"<p>The delivery service that shipped a physical product, such as Fedex, UPS, USPS, etc.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"name","value":"shipping.name","description":"<p>Recipient name.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"phone","value":"shipping.phone","description":"<p>Recipient phone (including extension).</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"tracking number","value":"shipping.tracking_number","description":"<p>The tracking number for a physical product, obtained from the delivery service. If multiple tracking numbers were generated for this purchase, please separate them with commas.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Shipping address","name":"shipping.address","description":"","fields":[{"name":"city","value":"shipping.address.city","description":"<p>City, district, suburb, town, or village.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"country","value":"shipping.address.country","description":"<p>Two-letter country code (<a href=\\"https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2\\">ISO 3166-1 alpha-2</a>).</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"line1","value":"shipping.address.line1","description":"<p>Address line 1 (e.g., street, PO Box, or company name).</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"line2","value":"shipping.address.line2","description":"<p>Address line 2 (e.g., apartment, suite, unit, or building).</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"postal code","value":"shipping.address.postal_code","description":"<p>ZIP or postal code.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"state","value":"shipping.address.state","description":"<p>State, county, province, or region.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]}],"formattedStripeProductItemFields":[{"label":"Standard Mappings","name":"standard","description":"","fields":[{"name":"active","value":"active","description":"<p>Whether the product is currently available for purchase. Defaults to <code>true</code>.</p>","type":"boolean","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"description","value":"description","description":"<p>The product\'s description, meant to be displayable to the customer. Use this field to optionally store a long form explanation of the product being sold for your own rendering purposes.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"name","value":"name","description":"<p>The product\'s name, meant to be displayable to the customer.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"shippable","value":"shippable","description":"<p>Whether this product is shipped (i.e., physical goods).</p>","type":"boolean","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"statement descriptor","value":"statement_descriptor","description":"<p>An arbitrary string to be displayed on your customer\'s credit card or bank statement. While most banks display this information consistently, some may display it incorrectly or not at all.</p>\\n<p>This may be up to 22 characters. The statement description may not include <code>&lt;</code>, <code>&gt;</code>, <code>\\\\</code>, <code>\\"</code>, <code>\'</code> characters, and will appear on your customer\'s statement in capital letters. Non-ASCII characters are automatically stripped.\\n It must contain at least one letter.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"tax code","value":"tax_code","description":"<p>A <a href=\\"https://stripe.com/docs/tax/tax-categories\\">tax code</a> ID.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"unit label","value":"unit_label","description":"<p>A label that represents units of this product in Stripe and on customers’ receipts and invoices. When set, this will be included in associated invoice line item descriptions.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"url","value":"url","description":"<p>A URL of a publicly-accessible webpage for this product.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Default price data","name":"default_price_data","description":"","fields":[{"name":"currency","value":"default_price_data.currency","description":"<p>Three-letter <a href=\\"https://www.iso.org/iso-4217-currency-codes.html\\">ISO currency code</a>, in lowercase. Must be a <a href=\\"https://stripe.com/docs/currencies\\">supported currency</a>.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"tax behavior","value":"default_price_data.tax_behavior","description":"<p>Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of <code>inclusive</code>, <code>exclusive</code>, or <code>unspecified</code>. Once specified as either <code>inclusive</code> or <code>exclusive</code>, it cannot be changed.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"unit amount","value":"default_price_data.unit_amount","description":"<p>The cost of each unit of product being credited.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"unit amount decimal","value":"default_price_data.unit_amount_decimal","description":"<p>Same as <code>unit_amount</code>, but contains a decimal value with at most 12 decimal places.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Default price data Recurring","name":"default_price_data.recurring","description":"","fields":[{"name":"interval","value":"default_price_data.recurring.interval","description":"<p>Interval (or event) to which the amount applies.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"interval count","value":"default_price_data.recurring.interval_count","description":"<p>The number of intervals between payments. For example, <code>interval=month</code> and <code>interval_count=3</code> indicates one payment every three months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks). This parameter is optional when <code>interval=sporadic</code>.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Package dimensions","name":"package_dimensions","description":"","fields":[{"name":"height","value":"package_dimensions.height","description":"<p>Height, in inches.</p>","type":"number","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"length","value":"package_dimensions.length","description":"<p>Length, in inches.</p>","type":"number","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"weight","value":"package_dimensions.weight","description":"<p>Weight, in ounces.</p>","type":"number","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"width","value":"package_dimensions.width","description":"<p>Width, in inches.</p>","type":"number","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]}],"formattedStripeSubscriptionFields":[{"label":"Standard Mappings","name":"standard","description":"","fields":[{"name":"end behavior","value":"end_behavior","description":"<p>Configures how the subscription schedule behaves when it ends. Possible values are <code>release</code> or <code>cancel</code> with the default being <code>release</code>. <code>release</code> will end the subscription schedule and keep the underlying subscription running.<code>cancel</code> will end the subscription schedule and cancel the underlying subscription.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":"","enum":["cancel","none","release","renew"]},{"name":"from subscription","value":"from_subscription","description":"<p>Migrate an existing subscription to be managed by a subscription schedule. If this parameter is set, a subscription schedule will be created using the subscription\'s item(s), set to auto-renew using the subscription\'s interval. When using this parameter, other parameters (such as phase values) cannot be set. To create a subscription schedule with other modifications, we recommend making two separate API calls.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"start date","value":"start_date","description":"<p>When the subscription schedule starts. We recommend using <code>now</code> so that it starts the subscription immediately. You can also use a Unix timestamp to backdate the subscription so that it starts on a past date, or set a future date for the subscription to start on.</p>","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Default settings","name":"default_settings","description":"","fields":[{"name":"application fee percent","value":"default_settings.application_fee_percent","description":"<p>This represents the percentage of the subscription invoice subtotal that will be transferred to the application owner\'s Stripe account.</p>","type":"number","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"billing cycle anchor","value":"default_settings.billing_cycle_anchor","description":"<p>Determines the date of the first full invoice, and, for plans with <code>month</code> or <code>year</code> intervals, the day of the month for subsequent invoices.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"collection method","value":"default_settings.collection_method","description":"<p>Either <code>charge_automatically</code>, or <code>send_invoice</code>. When charging automatically, Stripe will attempt to pay this invoice using the default source attached to the customer. When sending an invoice, Stripe will email this invoice to the customer with payment instructions.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"default payment method","value":"default_settings.default_payment_method","description":"<p>ID of the default payment method for the invoice. It must belong to the customer associated with the invoice. If not set, defaults to the subscription\'s default payment method, if any, or to the default payment method in the customer\'s invoice settings.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Default settings Billing thresholds","name":"billing_thresholds","description":"","fields":[{"name":"amount gte","value":"default_settings.billing_thresholds.amount_gte","description":"<p>The total invoice amount threshold boundary if it triggered the threshold invoice.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"reset billing cycle anchor","value":"default_settings.billing_thresholds.reset_billing_cycle_anchor","description":"<p>Indicates if the <code>billing_cycle_anchor</code> should be reset when a threshold is reached. If true, <code>billing_cycle_anchor</code> will be updated to the date/time the threshold was last reached; otherwise, the value will remain unchanged. This value may not be <code>true</code> if the subscription contains items with plans that have <code>aggregate_usage=last_ever</code>.</p>","type":"boolean","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Default settings Transfer data","name":"transfer_data","description":"","fields":[{"name":"amount percent","value":"default_settings.transfer_data.amount_percent","description":"<p>A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice subtotal that will be transferred to the destination account. By default, the entire amount will be transferred to the destination.</p>","type":"number","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"destination","value":"default_settings.transfer_data.destination","description":"<p>ID of an existing, connected Stripe account to transfer funds to if <code>transfer_data</code> was specified in the charge request.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Default settings Automatic tax","name":"default_settings.automatic_tax","description":"","fields":[{"name":"enabled","value":"default_settings.automatic_tax.enabled","description":"<p>Whether Stripe automatically computes tax on this invoice.</p>","type":"boolean","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Default settings Invoice settings","name":"default_settings.invoice_settings","description":"","fields":[{"name":"days until due","value":"default_settings.invoice_settings.days_until_due","description":"<p>Number of days within which a customer must pay invoices generated by this quote. This value will be <code>null</code> for quotes where <code>collection_method=charge_automatically</code>.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"template","value":"default_settings.invoice_settings.rendering.template","description":"Invoice rendering template id to use for this subscription\'s invoice","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"template version","value":"default_settings.invoice_settings.rendering.template_version","description":"Version of the rendering template that will be used. If this field is null, then the latest version of the template will be automatically used.","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Prebilling","name":"prebilling","description":"","fields":[{"name":"Iterations","value":"prebilling.iterations","description":"<p>This is used to determine the number of billing cycles to prebill.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]}],"formattedStripeSubscriptionItemFields":[{"label":"Standard Mappings","name":"standard","description":"","fields":[{"name":"quantity","value":"quantity","description":"<p>The quantity you\'d like to apply to the subscription item you\'re creating.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]}],"formattedStripeSubscriptionSchedulePhaseFields":[{"label":"Standard Mappings","name":"standard","description":"","fields":[{"name":"application fee percent","value":"application_fee_percent","description":"<p>This represents the percentage of the subscription invoice subtotal that will be transferred to the application owner\'s Stripe account.</p>","type":"number","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"billing cycle anchor","value":"billing_cycle_anchor","description":"<p>Determines the date of the first full invoice, and, for plans with <code>month</code> or <code>year</code> intervals, the day of the month for subsequent invoices.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":"","enum":["automatic","phase_start"]},{"name":"collection method","value":"collection_method","description":"<p>Either <code>charge_automatically</code>, or <code>send_invoice</code>. When charging automatically, Stripe will attempt to pay this invoice using the default source attached to the customer. When sending an invoice, Stripe will email this invoice to the customer with payment instructions.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":"","enum":["charge_automatically","send_invoice"]},{"name":"coupon","value":"coupon","description":"<p>A coupon contains information about a percent-off or amount-off discount you\\nmight want to apply to a customer. Coupons may be applied to <a href=\\"https://stripe.com/docs/api#subscriptions\\">subscriptions</a>, <a href=\\"https://stripe.com/docs/api#invoices\\">invoices</a>,\\n<a href=\\"https://stripe.com/docs/api/checkout/sessions\\">checkout sessions</a>, <a href=\\"https://stripe.com/docs/api#quotes\\">quotes</a>, and more. Coupons do not work with conventional one-off <a href=\\"https://stripe.com/docs/api#create_charge\\">charges</a> or <a href=\\"https://stripe.com/docs/api/payment_intents\\">payment intents</a>.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"default payment method","value":"default_payment_method","description":"<p>ID of the default payment method for the invoice. It must belong to the customer associated with the invoice. If not set, defaults to the subscription\'s default payment method, if any, or to the default payment method in the customer\'s invoice settings.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"end date","value":"end_date","description":"<p>End date of the mandate or subscription. If not provided, the mandate will be active until canceled. If provided, end date should be after start date.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"iterations","value":"iterations","description":"","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"proration behavior","value":"proration_behavior","description":"<p>Whether to create prorations when canceling subscriptions. Possible values are <code>none</code> and <code>create_prorations</code>.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":"","enum":["always_invoice","create_prorations","none"]},{"name":"trial","value":"trial","description":"","type":"boolean","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"trial end","value":"trial_end","description":"<p>If the subscription has a trial, the end of that trial.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Automatic tax","name":"automatic_tax","description":"","fields":[{"name":"enabled","value":"automatic_tax.enabled","description":"<p>Whether Stripe automatically computes tax on this invoice.</p>","type":"boolean","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Billing thresholds","name":"billing_thresholds","description":"","fields":[{"name":"enabled","value":"billing_thresholds.enabled","description":"<p>Whether Stripe automatically computes tax on this invoice.</p>","type":"boolean","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Invoice settings","name":"invoice_settings","description":"","fields":[{"name":"days until due","value":"invoice_settings.days_until_due","description":"<p>Number of days within which a customer must pay invoices generated by this quote. This value will be <code>null</code> for quotes where <code>collection_method=charge_automatically</code>.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Transfer data","name":"transfer_data","description":"","fields":[{"name":"amount percent","value":"transfer_data.amount_percent","description":"<p>A non-negative decimal between 0 and 100, with at most two decimal places. This represents the percentage of the subscription invoice subtotal that will be transferred to the destination account. By default, the entire amount will be transferred to the destination.</p>","type":"number","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"destination","value":"transfer_data.destination","description":"<p>ID of an existing, connected Stripe account to transfer funds to if <code>transfer_data</code> was specified in the charge request.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]}],"formattedStripePriceFields":[{"label":"Standard Mappings","name":"standard","description":"","fields":[{"name":"active","value":"active","description":"<p>Whether the price can be used for new purchases. Defaults to <code>true</code>.</p>","type":"boolean","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"billing scheme","value":"billing_scheme","description":"<p>Describes how to compute the price per period. Either <code>per_unit</code> or <code>tiered</code>. <code>per_unit</code> indicates that the fixed amount (specified in <code>unit_amount</code> or <code>unit_amount_decimal</code>) will be charged per unit in <code>quantity</code> (for prices with <code>usage_type=licensed</code>), or per unit of total usage (for prices with <code>usage_type=metered</code>). <code>tiered</code> indicates that the unit pricing will be computed using a tiering strategy as defined using the <code>tiers</code> and <code>tiers_mode</code> attributes.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":"","enum":["per_unit","tiered"]},{"name":"currency","value":"currency","description":"<p>Three-letter <a href=\\"https://www.iso.org/iso-4217-currency-codes.html\\">ISO currency code</a>, in lowercase. Must be a <a href=\\"https://stripe.com/docs/currencies\\">supported currency</a>.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"lookup key","value":"lookup_key","description":"<p>A lookup key used to retrieve prices dynamically from a static string. This may be up to 200 characters.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"nickname","value":"nickname","description":"<p>A brief description of the price, hidden from customers.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"tax behavior","value":"tax_behavior","description":"<p>Specifies whether the price is considered inclusive of taxes or exclusive of taxes. One of <code>inclusive</code>, <code>exclusive</code>, or <code>unspecified</code>. Once specified as either <code>inclusive</code> or <code>exclusive</code>, it cannot be changed.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":"","enum":["exclusive","inclusive","unspecified"]},{"name":"tiers mode","value":"tiers_mode","description":"<p>Defines if the tiering price should be <code>graduated</code> or <code>volume</code> based. In <code>volume</code>-based tiering, the maximum quantity within a period determines the per unit price, in <code>graduated</code> tiering pricing can successively change as the quantity grows.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":"","enum":["graduated","volume"]},{"name":"transfer lookup key","value":"transfer_lookup_key","description":"<p>If set to true, will atomically remove the lookup key from the existing price, and assign it to this price.</p>","type":"boolean","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"unit amount","value":"unit_amount","description":"<p>A positive integer in cents (or local equivalent) (or 0 for a free price) representing how much to charge.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"unit amount decimal","value":"unit_amount_decimal","description":"<p>Same as <code>unit_amount</code>, but accepts a decimal value in cents (or local equivalent) with at most 12 decimal places. Only one of <code>unit_amount</code> and <code>unit_amount_decimal</code> can be set.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Recurring","name":"recurring","description":"","fields":[{"name":"aggregate usage","value":"recurring.aggregate_usage","description":"<p>Specifies a usage aggregation strategy for plans of <code>usage_type=metered</code>. Allowed values are <code>sum</code> for summing up all usage during a period, <code>last_during_period</code> for using the last usage record reported within a period, <code>last_ever</code> for using the last usage record ever (across period bounds) or <code>max</code> which uses the usage record with the maximum reported usage during a period. Defaults to <code>sum</code>.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"interval","value":"recurring.interval","description":"<p>Interval (or event) to which the amount applies.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"interval count","value":"recurring.interval_count","description":"<p>The number of intervals between payments. For example, <code>interval=month</code> and <code>interval_count=3</code> indicates one payment every three months. Maximum of one year interval allowed (1 year, 12 months, or 52 weeks). This parameter is optional when <code>interval=sporadic</code>.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"usage type","value":"recurring.usage_type","description":"<p>Configures how the quantity per period should be determined. Can be either <code>metered</code> or <code>licensed</code>. <code>licensed</code> automatically bills the <code>quantity</code> set when adding it to a subscription. <code>metered</code> aggregates the total usage based on usage records. Defaults to <code>licensed</code>.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]},{"label":"Transform quantity","name":"transform_quantity","description":"","fields":[{"name":"divide by","value":"transform_quantity.divide_by","description":"<p>Divide usage by this number.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"round","value":"transform_quantity.round","description":"<p>After division, either round the result <code>up</code> or <code>down</code>.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]}],"formattedStripeCouponFields":[{"label":"Standard Mappings","name":"standard","description":"","fields":[{"name":"amount off","value":"amount_off","description":"<p>A positive integer representing the amount to subtract from an invoice total (required if <code>percent_off</code> is not passed).</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"duration","value":"duration","description":"<p>Specifies how long the discount will be in effect if used on a subscription. Can be <code>forever</code>, <code>once</code>, or <code>repeating</code>. Defaults to <code>once</code>.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":"","enum":["forever","once","repeating"]},{"name":"duration in months","value":"duration_in_months","description":"<p>Required only if <code>duration</code> is <code>repeating</code>, in which case it must be a positive integer that specifies the number of months the discount will be in effect.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"max redemptions","value":"max_redemptions","description":"<p>A positive integer specifying the number of times the coupon can be redeemed before it\'s no longer valid. For example, you might have a 50% off coupon that the first 20 readers of your blog can use.</p>","type":"integer","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"name","value":"name","description":"<p>Name of the coupon displayed to customers on, for instance invoices, or receipts. By default the <code>id</code> is shown if <code>name</code> is not set.</p>","type":"string","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""},{"name":"percent off","value":"percent_off","description":"<p>A positive float larger than 0, and smaller or equal to 100, that represents the discount the coupon will apply (required if <code>amount_off</code> is not passed).</p>","type":"number","defaultValue":"","requiredValue":"","hasOverride":false,"staticValue":false,"hasSfValue":false,"hasRequiredValue":false,"sfValue":"","sfValueType":""}]}]}');
            res.setStatusCode(200);
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

    public class UnifiedConfigMock implements HttpCalloutMock {
        public HttpResponse respond(HttpRequest req) {
            stripeConnectionSuccessMock mock1 = new stripeConnectionSuccessMock();
            stripeSuccessMock mock2 = new stripeSuccessMock();

            HttpResponse res1 = mock1.respond(req);
            HttpResponse res2 = mock2.respond(req);

            Map<String,Object> body1 = (Map<String,Object>) JSON.deserializeUntyped(res1.getBody());
            Map<String,Object> body2 = (Map<String,Object>) JSON.deserializeUntyped(res2.getBody());

            for (String key : body1.keySet()) {
                body2.put(key, body1.get(key));
            }

            res2.setBody(JSON.serialize(body2));

            return res2;
        }
    }

    public class ConnectionStatusesMock implements HttpCalloutMock {
        public Integer calls = 0;
        public HttpResponse respond(HttpRequest req) {
            calls += 1;
            HttpResponse res = new HttpResponse();
            Map<String, Object> services = new Map<String, Object>();
            String salesforceSvc = setupAssistant.KEY_SALESFORCE + setupAssistant.SERVICE_DELIMINATOR + UserInfo.getOrganizationId();
            String stripeAcct = TEST_STRIPE_ACCOUNT_ID;
            String stripeLivemode = 'test';
            String stripeSvc = setupAssistant.KEY_STRIPE + setupAssistant.SERVICE_DELIMINATOR + stripeAcct + setupAssistant.SERVICE_DELIMINATOR + stripeLivemode;

            services.put(salesforceSvc, true);
            services.put(stripeSvc, true);

            res.setBody(JSON.serialize(services));
            res.setStatusCode(200);

            return res;
        }
    }

    public class AccountManagementMock implements HttpCalloutMock {
        public Integer calls = 0;
        public Boolean nextCallFails = false;

        public HttpResponse respond(HttpRequest req) {
            calls += 1;

            if (nextCallFails) {
                nextCallFails = false;
                return respondWithFailure();
            }

            String method = req.getMethod();
            String endpoint = req.getEndpoint();
            Boolean isAccountsUri = endpoint.contains(constants.RUBY_SERVICE_ACCOUNTS_URI);
            Boolean isSetDefaultUri = endpoint.endsWith(constants.RUBY_SERVICE_ACCOUNTS_SET_DEFAULT_URI);

            if (method == 'POST' && isAccountsUri && isSetDefaultUri) {
                HttpResponse res = new HttpResponse();
                res.setBody(JSON.serialize(new Map<String, Object>()));
                res.setStatusCode(200);
                return res;
            }

            if (method == 'DELETE' && isAccountsUri && isSetDefaultUri == false) {
                HttpResponse res = new HttpResponse();
                res.setBody(JSON.serialize(new Map<String, Object>()));
                res.setStatusCode(200);
                return res;
            }

            return respondWithFailure();
        }

        public HttpResponse respondWithFailure() {
            HttpResponse res = new HttpResponse();
            res.setBody(JSON.serialize(new Map<String, Object>()));
            res.setStatusCode(500);
            return res;
        }
    }

    @IsTest
    static void validateSharedState() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);

        StripeOAuthState state = StripeOAuthState.create();
        String stateString = state.toString();

        Map<String, Object> restoredState = setupAssistant.validateSharedState(stateString);
        Map<String, Object> oldState = (Map<String, Object>) JSON.deserializeUntyped(state.toJSON());
        for (String key : oldState.keySet()) {
            System.assertEquals(oldState.get(key), restoredState.get(key));
        }
        for (String key : restoredState.keySet()) {
            System.assertEquals(oldState.get(key), restoredState.get(key));
        }
    }

    @IsTest
    static void getExportableConfigDownloadUrl() {
        // this basically just has to work... we can't actually test it...
        String url = setupAssistant.getExportableConfigDownloadUrl();
        System.assertNotEquals(null, url);
    }

    @IsTest
    static void revokeIntegrationUserAuthorization() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        // this is really just a placeholder...
        setupAssistant.revokeIntegrationUserAuthorization();
    }

    @IsTest
    static void getIntegrationUserAuthorizationUri() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        responseData uriUnforced = unpack(setupAssistant.getIntegrationUserAuthorizationUri(false));
        responseData uriForced = unpack(setupAssistant.getIntegrationUserAuthorizationUri(true));

        System.assert(uriUnforced.getString('authorization_uri').contains('prompt=login') == false);
        System.assert(uriUnforced.getString('authorization_uri').contains(uriUnforced.getString('message_origin_uri')));
        System.assert(uriForced.getString('authorization_uri').contains('prompt=login') == true);
        System.assert(uriForced.getString('authorization_uri').contains(uriForced.getString('message_origin_uri')));
    }

    @IsTest
    static void getStripeAuthorizationUri() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        responseData testMode = unpack(setupAssistant.getStripeAuthorizationUri(false));
        responseData liveMode = unpack(setupAssistant.getStripeAuthorizationUri(true));

        System.assert(testMode.getString('authorization_uri').contains('live') == false);
        System.assert(testMode.getString('authorization_uri').contains('test') == true);
        System.assert(testMode.getString('authorization_uri').contains(testMode.getString('message_origin_uri')));
        System.assert(liveMode.getString('authorization_uri').contains('test') == false);
        System.assert(liveMode.getString('authorization_uri').contains('live') == true);
        System.assert(liveMode.getString('authorization_uri').contains(liveMode.getString('message_origin_uri')));
    }

    @IsTest
    static void checkServiceConnectionStatuses() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        responseData statuses = unpack(setupAssistant.checkServiceConnectionStatuses());
        Map<String, Object> services = (Map<String, Object>) statuses.results.get('services');
        System.assert(services.get(setupAssistant.KEY_SALESFORCE_SERVICE) == setupAssistant.STATUS_CONNECTED);
        String stripeKey = 'stripe|' + TEST_STRIPE_ACCOUNT_ID + '|test';
        System.assert(services.get(stripeKey) == setupAssistant.STATUS_CONNECTED);
    }

    @IsTest
    static void testGetServiceConnectionStatuses() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        String salesforceSvc = setupAssistant.KEY_SALESFORCE + setupAssistant.SERVICE_DELIMINATOR + UserInfo.getOrganizationId();
        String stripeAcct = TEST_STRIPE_ACCOUNT_ID;
        String stripeLivemode = 'test';
        String stripeSvc = setupAssistant.KEY_STRIPE + setupAssistant.SERVICE_DELIMINATOR + stripeAcct + setupAssistant.SERVICE_DELIMINATOR + stripeLivemode;

        Map<String, Object> services = new Map<String, Object>();
        services.put(salesforceSvc, true);
        services.put(stripeSvc, true);
        HttpResponse resp = new HttpResponse();
        resp.setBody(JSON.serialize(services));
        resp.setStatusCode(200);

        Continuation con = (Continuation) setupAssistant.getServiceConnectionStatuses(stripeSvc);
        Test.setContinuationResponse('continuation-1', resp);
        setupAssistant cont = new setupAssistant();

        String result = (String) Test.invokeContinuationMethod(cont, con);

        debugger.debug('apex result', result);
        responseData statuses = unpack(result);
        services = (Map<String, Object>) statuses.results.get('services');
        debugger.debug('services', services);
        System.assertEquals(setupAssistant.STATUS_FRESH, (String)services.get(stripeSvc));
        System.assertEquals(setupAssistant.STATUS_CONNECTED, (String)services.get(salesforceSvc));
    }

    @IsTest
    static void testGetServiceConnectionStatusesSync() {
        Setup_Connection_Data__mdt testSetupData = getTestStripeConnectionKey();
        testSetupData.Enable_Apex_Debug__c = true;
        setTestGlobalKey(testSetupData);
        insertTestConnectedRecord();

        String salesforceSvc = setupAssistant.KEY_SALESFORCE + setupAssistant.SERVICE_DELIMINATOR + UserInfo.getOrganizationId();
        String stripeAcct = TEST_STRIPE_ACCOUNT_ID;
        String stripeLivemode = 'test';
        String stripeSvc = setupAssistant.KEY_STRIPE + setupAssistant.SERVICE_DELIMINATOR + stripeAcct + setupAssistant.SERVICE_DELIMINATOR + stripeLivemode;

        Test.setMock(HttpCalloutMock.class, new ConnectionStatusesMock());

        Test.startTest();

        String result = setupAssistant.getServiceConnectionStatusesSync(stripeSvc);

        Test.stopTest();

        debugger.debug('apex result', result);
        responseData statuses = unpack(result);
        Map<String, Object> services = (Map<String, Object>) statuses.results.get(setupAssistant.KEY_SERVICES);
        debugger.debug('services', services);
        System.assertEquals(setupAssistant.STATUS_FRESH, (String)services.get(stripeSvc));
        System.assertEquals(setupAssistant.STATUS_CONNECTED, (String)services.get(salesforceSvc));
    }


    private static responseData unpack(String data) {
        debugger.debug('unpacking', data);
        responseData rd = new responseData();
        Map<String, Object> res = (Map<String, Object>) JSON.deserializeUntyped(data);
        rd.results = (Map<String, Object>) res.get('results');
        rd.error = (String) res.get('error');
        rd.isSuccess = (Boolean) res.get('isSuccess');
        return rd;
    }
}