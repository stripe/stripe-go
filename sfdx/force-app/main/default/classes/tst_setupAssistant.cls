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
    static public void testSetupAssistant_isCpqEnabled() {
        Test.startTest();
        String response = setupAssistant.isCpqEnabled();
        Test.stopTest();
        
        Map<String, Object> responseMap = (Map<String, Object>)JSON.deserializeUntyped(response);
        System.assert((Boolean)responseMap.get('isSuccess'), responseMap.get('error'));
        System.assertEquals(true, responseMap.get('isSuccess'));

        Map<String, Object> resultsMap = (Map<String, Object>)responseMap.get('results');
        System.assertEquals(true, resultsMap.get('isCpqInstalled'));
    }

    @isTest
    static public void testSetupAssistant_validateConnectionStatus() {
        Test.startTest();
        Setup_Connection_Data__mdt testSetupData = new Setup_Connection_Data__mdt();
        Map<String,Object> fields =  (Map<String,Object>) JSON.deserializeUntyped(JSON.serializePretty(testSetupData));
        fields.put(Setup_Connection_Data__mdt.Label.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.DeveloperName.getDescribe().getName(), 'Default');
        fields.put(Setup_Connection_Data__mdt.Salesforce_Connected__c.getDescribe().getName(), false);
        fields.put(Setup_Connection_Data__mdt.Stripe_Connected__c.getDescribe().getName(), false);
        fields.put(Setup_Connection_Data__mdt.Global_Key__c.getDescribe().getName(), 'testGlobalGUID');
        fields.put(Setup_Connection_Data__mdt.API_Key__c.getDescribe().getName(), 'testPackageGUID');
        testSetupData = (Setup_Connection_Data__mdt) JSON.deserialize(JSON.serialize(fields), Setup_Connection_Data__mdt.class);
        
        List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Salesforce_Connected__c, Stripe_Connected__c, Global_Key__c, API_Key__c
                                                            FROM Setup_Connection_Data__mdt
                                                            WITH SECURITY_ENFORCED];  
        if(!setupConfigList.isEmpty()) {
            setupAssistant.setupConfigMetadata[0] = testSetupData;
        } else {
            setupAssistant.setupConfigMetadata.add(testSetupData);  
        }

        Test.setMock(HttpCalloutMock.class, new stripeSuccessMock());  
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

    private class stripeSuccessMock implements HttpCalloutMock {
        public HttpResponse respond(HttpRequest req) {
            system.debug(req.getEndpoint());
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

}
