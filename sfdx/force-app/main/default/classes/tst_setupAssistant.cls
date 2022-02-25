@isTest
public with sharing class tst_setupAssistant {
    @isTest
    static public void testSetupAssistant() {
        Test.startTest();
        setupAssistant con = new setupAssistant();
        Test.stopTest();

        System.assertEquals(constants.NAMESPACE, con.namespace);
    }

    @isTest
    static public void testSetupAssistant_getSetupData() {
        setupAssistant.saveData(new Map<String, Object>());

        String response = setupAssistant.getSetupData();

        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assert(resultsMap.containsKey('setupData'), '`data` not found');
    } 

    @isTest
    static public void testUtilities_isCpqEnabled() {

        String response = utilities.isCpqEnabled();
        
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));
        System.assertEquals(true, responseMap.get('isSuccess'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals(true, resultsMap.get('isCpqInstalled'));
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
        Setup_Connection_Data__mdt testSetupData = new Setup_Connection_Data__mdt();
        Map<String,Object> fields =  (Map<String,Object>) JSON.deserializeUntyped(JSON.serializePretty(testSetupData));
        fields.put(Setup_Connection_Data__mdt.Label.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.DeveloperName.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.Global_Key__c.getDescribe().getName(), 'testGlobalGUID');
        testSetupData = (Setup_Connection_Data__mdt) JSON.deserialize(JSON.serialize(fields), Setup_Connection_Data__mdt.class);
        
        List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Global_Key__c
                                                            FROM Setup_Connection_Data__mdt];  
        if(!setupConfigList.isEmpty()) {
            utilities.setupConfigMetadata[0] = testSetupData;
        } else {
            utilities.setupConfigMetadata.add(testSetupData);  
        }

        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c(); 
        stripeConnectRec.Stripe_Connected__c = false;
        stripeConnectRec.Salesforce_Connected__c = false;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        upsert stripeConnectRec;

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());  
        String connectedCallbackResponse = setupAssistant.validateConnectionStatus(true);
        String response = setupAssistant.validateConnectionStatus(false);
        Test.stopTest();
        
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Map<String, Object> connectedCallbackResponseMap = (Map<String, Object>)JSON.deserializeUntyped(connectedCallbackResponse);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));
        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        Map<String, Object> connectecCallbackResultsMap = (Map<String, Object>)connectedCallbackResponseMap.get('results');
        System.assertEquals('fresh', resultsMap.get('isConnected'));
        System.assertEquals(false, connectecCallbackResultsMap.get('isConnected'));
    }

    @isTest
    static public void testSetupAssistant_connected() {
        Setup_Connection_Data__mdt testSetupData = new Setup_Connection_Data__mdt();
        Map<String,Object> fields =  (Map<String,Object>) JSON.deserializeUntyped(JSON.serializePretty(testSetupData));
        fields.put(Setup_Connection_Data__mdt.Label.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.DeveloperName.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.Global_Key__c.getDescribe().getName(), 'testGlobalGUID');
        testSetupData = (Setup_Connection_Data__mdt) JSON.deserialize(JSON.serialize(fields), Setup_Connection_Data__mdt.class);
        
        List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Global_Key__c
                                                            FROM Setup_Connection_Data__mdt];  
        if(!setupConfigList.isEmpty()) {
            utilities.setupConfigMetadata[0] = testSetupData;
        } else {
            utilities.setupConfigMetadata.add(testSetupData);  
        }

        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c(); 
        stripeConnectRec.Stripe_Connected__c = true;
        stripeConnectRec.Salesforce_Connected__c = true;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        upsert stripeConnectRec;

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());  
        String response = setupAssistant.validateConnectionStatus(true);
        Test.stopTest();

        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));
        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals(true, (Boolean)resultsMap.get('isConnected'));
    }

    @isTest
    static public void testSetupAssistant_failConnected() {
        Setup_Connection_Data__mdt testSetupData = new Setup_Connection_Data__mdt();
        Map<String,Object> fields =  (Map<String,Object>) JSON.deserializeUntyped(JSON.serializePretty(testSetupData));
        fields.put(Setup_Connection_Data__mdt.Label.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.DeveloperName.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.Global_Key__c.getDescribe().getName(), 'testGlobalGUID');
        testSetupData = (Setup_Connection_Data__mdt) JSON.deserialize(JSON.serialize(fields), Setup_Connection_Data__mdt.class);
        
        List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Global_Key__c
                                                            FROM Setup_Connection_Data__mdt];  
        if(!setupConfigList.isEmpty()) {
            utilities.setupConfigMetadata[0] = testSetupData;
        } else {
            utilities.setupConfigMetadata.add(testSetupData);  
        }

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectFailMock());  
        String response = setupAssistant.validateConnectionStatus(false);
        Test.stopTest();

        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'));
        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals(false, resultsMap.get('isConnected'));
    }

    @isTest
    static public void testSetupAssistant_Disconnected() {
        Setup_Connection_Data__mdt testSetupData = new Setup_Connection_Data__mdt();
        Map<String,Object> fields =  (Map<String,Object>) JSON.deserializeUntyped(JSON.serializePretty(testSetupData));
        fields.put(Setup_Connection_Data__mdt.Label.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.DeveloperName.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.Global_Key__c.getDescribe().getName(), 'testGlobalGUID');
        testSetupData = (Setup_Connection_Data__mdt) JSON.deserialize(JSON.serialize(fields), Setup_Connection_Data__mdt.class);
        
        List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Global_Key__c
                                                            FROM Setup_Connection_Data__mdt];  
        if(!setupConfigList.isEmpty()) {
            setupAssistant.setupConfigMetadata[0] = testSetupData;
        } else {
            setupAssistant.setupConfigMetadata.add(testSetupData);  
        }
        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectionSuccessMock());  
        String connectedCallbackResponse = setupAssistant.validateConnectionStatus(true);
        String response = setupAssistant.validateConnectionStatus(false);
        Test.stopTest();
        
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Map<String, Object> connectedCallbackResponseMap = (Map<String, Object>)JSON.deserializeUntyped(connectedCallbackResponse);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));
        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        Map<String, Object> connectecCallbackResultsMap = (Map<String, Object>)connectedCallbackResponseMap.get('results');

    }

    @isTest
    static public void testSetupAssistant_getPicklistValuesForMapperConnectedCallback() {
        Setup_Connection_Data__mdt testSetupData = new Setup_Connection_Data__mdt();
        Map<String,Object> fields =  (Map<String,Object>) JSON.deserializeUntyped(JSON.serializePretty(testSetupData));
        fields.put(Setup_Connection_Data__mdt.Label.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.DeveloperName.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.Global_Key__c.getDescribe().getName(), 'testGlobalGUID');
        testSetupData = (Setup_Connection_Data__mdt) JSON.deserialize(JSON.serialize(fields), Setup_Connection_Data__mdt.class);
        
        List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Global_Key__c
                                                            FROM Setup_Connection_Data__mdt];  
        if(!setupConfigList.isEmpty()) {
            utilities.setupConfigMetadata[0] = testSetupData;
        } else {
            utilities.setupConfigMetadata.add(testSetupData);  
        }

        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c(); 
        stripeConnectRec.Stripe_Connected__c = true;
        stripeConnectRec.Salesforce_Connected__c = true;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        upsert stripeConnectRec;

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
        Setup_Connection_Data__mdt testSetupData = new Setup_Connection_Data__mdt();
        Map<String,Object> fields =  (Map<String,Object>) JSON.deserializeUntyped(JSON.serializePretty(testSetupData));
        fields.put(Setup_Connection_Data__mdt.Label.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.DeveloperName.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.Global_Key__c.getDescribe().getName(), 'testGlobalGUID');
        testSetupData = (Setup_Connection_Data__mdt) JSON.deserialize(JSON.serialize(fields), Setup_Connection_Data__mdt.class);
        
        List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Global_Key__c
                                                            FROM Setup_Connection_Data__mdt];  
        if(!setupConfigList.isEmpty()) {
            utilities.setupConfigMetadata[0] = testSetupData;
        } else {
            utilities.setupConfigMetadata.add(testSetupData);  
        }

        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c(); 
        stripeConnectRec.Stripe_Connected__c = true;
        stripeConnectRec.Salesforce_Connected__c = true;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        upsert stripeConnectRec;

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
        Setup_Connection_Data__mdt testSetupData = new Setup_Connection_Data__mdt();
        Map<String,Object> fields =  (Map<String,Object>) JSON.deserializeUntyped(JSON.serializePretty(testSetupData));
        fields.put(Setup_Connection_Data__mdt.Label.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.DeveloperName.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.Global_Key__c.getDescribe().getName(), 'testGlobalGUID');
        testSetupData = (Setup_Connection_Data__mdt) JSON.deserialize(JSON.serialize(fields), Setup_Connection_Data__mdt.class);
        
        List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Global_Key__c
                                                            FROM Setup_Connection_Data__mdt];  
        if(!setupConfigList.isEmpty()) {
            utilities.setupConfigMetadata[0] = testSetupData;
        } else {
            utilities.setupConfigMetadata.add(testSetupData);  
        }

        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c(); 
        stripeConnectRec.Stripe_Connected__c = true;
        stripeConnectRec.Salesforce_Connected__c = true;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        upsert stripeConnectRec;

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeMappingSuccessMock());  
        String response = setupAssistant.saveMappingConfigurations('test');
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));
        
        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals((Boolean)resultsMap.get('isConfigSaved'), true); 
    }

    @isTest
    static public void testSetupAssistant_failSaveMappingConfigurations() {
        Setup_Connection_Data__mdt testSetupData = new Setup_Connection_Data__mdt();
        Map<String,Object> fields =  (Map<String,Object>) JSON.deserializeUntyped(JSON.serializePretty(testSetupData));
        fields.put(Setup_Connection_Data__mdt.Label.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.DeveloperName.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.Global_Key__c.getDescribe().getName(), 'testGlobalGUID');
        testSetupData = (Setup_Connection_Data__mdt) JSON.deserialize(JSON.serialize(fields), Setup_Connection_Data__mdt.class);
        
        List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Global_Key__c
                                                            FROM Setup_Connection_Data__mdt];  
        if(!setupConfigList.isEmpty()) {
            utilities.setupConfigMetadata[0] = testSetupData;
        } else {
            utilities.setupConfigMetadata.add(testSetupData);  
        }

        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c(); 
        stripeConnectRec.Stripe_Connected__c = true;
        stripeConnectRec.Salesforce_Connected__c = true;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        upsert stripeConnectRec;

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectFailMock());  
        String response = setupAssistant.saveMappingConfigurations('test');
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();
        
        System.assertEquals((Boolean)responseMap.get('isSuccess'), false);

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals((Boolean)resultsMap.get('isConfigSaved'), false); 
    }

    @isTest
    static public void testSetupAssistant_getMappingConfigurations() {
        Setup_Connection_Data__mdt testSetupData = new Setup_Connection_Data__mdt();
        Map<String,Object> fields =  (Map<String,Object>) JSON.deserializeUntyped(JSON.serializePretty(testSetupData));
        fields.put(Setup_Connection_Data__mdt.Label.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.DeveloperName.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.Global_Key__c.getDescribe().getName(), 'testGlobalGUID');
        testSetupData = (Setup_Connection_Data__mdt) JSON.deserialize(JSON.serialize(fields), Setup_Connection_Data__mdt.class);
        
        List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Global_Key__c
                                                            FROM Setup_Connection_Data__mdt];  
        if(!setupConfigList.isEmpty()) {
            utilities.setupConfigMetadata[0] = testSetupData;
        } else {
            utilities.setupConfigMetadata.add(testSetupData);  
        }

        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c(); 
        stripeConnectRec.Stripe_Connected__c = true;
        stripeConnectRec.Salesforce_Connected__c = true;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        upsert stripeConnectRec;

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeMappingSuccessMock());  
        String response = setupAssistant.getMappingConfigurations();
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        Map<String, Object> allMappingConfigurations = (Map<String, Object>)resultsMap.get('allMappingConfigurations');
        Map<String, Object> default_mappings = (Map<String, Object>)allMappingConfigurations.get('default_mapping');
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
        Setup_Connection_Data__mdt testSetupData = new Setup_Connection_Data__mdt();
        Map<String,Object> fields =  (Map<String,Object>) JSON.deserializeUntyped(JSON.serializePretty(testSetupData));
        fields.put(Setup_Connection_Data__mdt.Label.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.DeveloperName.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.Global_Key__c.getDescribe().getName(), 'testGlobalGUID');
        testSetupData = (Setup_Connection_Data__mdt) JSON.deserialize(JSON.serialize(fields), Setup_Connection_Data__mdt.class);
        
        List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Global_Key__c
                                                            FROM Setup_Connection_Data__mdt];  
        if(!setupConfigList.isEmpty()) {
            utilities.setupConfigMetadata[0] = testSetupData;
        } else {
            utilities.setupConfigMetadata.add(testSetupData);  
        }

        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c(); 
        stripeConnectRec.Stripe_Connected__c = true;
        stripeConnectRec.Salesforce_Connected__c = true;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        upsert stripeConnectRec;

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeConnectFailMock());  
        String response = setupAssistant.getMappingConfigurations();
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assertEquals((Boolean)responseMap.get('isSuccess'), false);
    }


    @isTest
    static public void testSetupAssistant_manualRetry() {
        Setup_Connection_Data__mdt testSetupData = new Setup_Connection_Data__mdt();
        Map<String,Object> fields =  (Map<String,Object>) JSON.deserializeUntyped(JSON.serializePretty(testSetupData));
        fields.put(Setup_Connection_Data__mdt.Label.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.DeveloperName.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.Global_Key__c.getDescribe().getName(), 'testGlobalGUID');
        testSetupData = (Setup_Connection_Data__mdt) JSON.deserialize(JSON.serialize(fields), Setup_Connection_Data__mdt.class);
        
        List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Global_Key__c
                                                            FROM Setup_Connection_Data__mdt];  
        if(!setupConfigList.isEmpty()) {
            utilities.setupConfigMetadata[0] = testSetupData;
        } else {
            utilities.setupConfigMetadata.add(testSetupData);  
        }

        Stripe_Connection__c stripeConnectRec = new Stripe_Connection__c(); 
        stripeConnectRec.Stripe_Connected__c = true;
        stripeConnectRec.Salesforce_Connected__c = true;
        stripeConnectRec.API_Key__c = 'testPackageGUID';
        upsert stripeConnectRec;

        Test.startTest();
        Test.setMock(HttpCalloutMock.class, new stripeRetrySuccessMock());  
        String response = setupAssistant.manualRetry('test');
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        Test.stopTest();

        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));

    }

    private class stripeConnectionSuccessMock implements HttpCalloutMock {
        public HttpResponse respond(HttpRequest req) {
            HttpResponse res = new HttpResponse();
            if(req.getEndpoint().endsWith('configuration')) {
                Map<String,Object> responseObject = new Map<String,Object>();
                responseObject.put('salesforce', true);
                responseObject.put('stripe', true);
    
                Map<String,Object> responseBody = new Map<String,Object>();
                responseBody.put('connection_status', (Object)responseObject);
    
                res.setBody(Json.serialize(responseBody));
                res.setStatusCode(200);
            } 
            return res;
        }
    }

    private class stripeRetrySuccessMock implements HttpCalloutMock {
        public HttpResponse respond(HttpRequest req) {
            HttpResponse res = new HttpResponse();
            if(req.getEndpoint().endsWith('translate')) {
                Map<String,Object> responseObject = new Map<String,Object>();
                //responseObject.put('salesforce', true);
    
                Map<String,Object> responseBody = new Map<String,Object>();
                responseBody.put('retry_status', (Object)responseObject);
    
                res.setBody(Json.serialize(responseBody));
                res.setStatusCode(200);
            } 
            return res;
        }
    }

        private class stripeMappingSuccessMock implements HttpCalloutMock {
            public HttpResponse respond(HttpRequest req) {
                HttpResponse res = new HttpResponse();
                if (req.getEndpoint().endsWith('configuration')) {
                    Map<String,Object> fieldMap = new Map<String,Object>();
                    Map<String,Object> testFieldMap = new Map<String,Object>();
                    testFieldMap.put('testFieldMapping', 'test');
                    fieldMap.put('field_mappings', (Object)testFieldMap);

                    Map<String,Object> defaultMap = new Map<String,Object>();
                    Map<String,Object> testdefaultMap = new Map<String,Object>();
                    testdefaultMap.put('testDefaultMap', 'test');
                    defaultMap.put('default_mapping',(Object)testdefaultMap);

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
                    responseBody.put('default_mapping', (Object)defaultMap);
                    responseBody.put('required_mappings', (Object)requiredMapping);
                    responseBody.put('field_defaults', (Object)fieldDefaultMap);
        
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


