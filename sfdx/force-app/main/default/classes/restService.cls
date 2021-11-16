@RestResource(urlMapping='/batchApi')
global with sharing class restService {
    public class SetupIsNotCompleteException extends Exception {}
    public class NotConnectedToStripeException extends Exception {}
    public class RequestAuthenticationFailedException extends Exception {}
    @HttpPost
    global static void restServiceEvents() {
        RestRequest restRequest = RestContext.request;
        RestResponse response = RestContext.response;
        try { 
            //may not need to grab this from the header 
            /* String reqOrgId = restRequest.headers.get('Salesforce-Account-Id'); 
            String reqGlobalKey = restRequest.headers.get('Salesforce-Key'); 
            String reqPackageGuidKey = restRequest.headers.get('organization_key'); 

            List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Salesforce_Connected__c, Stripe_Connected__c, Global_Key__c, API_Key__c
                                                                FROM Setup_Connection_Data__mdt
                                                                WITH SECURITY_ENFORCED];    
            if(!setupConfigList.isEmpty()) {
                Setup_Connection_Data__mdt setupData = setupConfigList[0];
                if(setupData.Salesforce_Connected__c == true && setupData.Stripe_Connected__c == true) {
                    String orgId = constants.ORG_ID;
                    //String globalKey = setupData.Global_Key__c; //Arnold comment out when in package 
                    String globalKey = 'ABC123';
                    String packageGuidKey = setupData.API_Key__c;
                    if((orgId == reqOrgId) && (globalKey == reqGlobalKey) && (packageGuidKey == reqPackageGuidKey)) { */
                        //Validation complete do stuff now
                        Map<String, Object> parsedReqBody = (Map<String, Object>)JSON.deserializeUntyped(restRequest.requestBody.toString());
                        System.debug('parsedReqBody');
                        System.debug(parsedReqBody);
                        List<Object> orderIds = (List<Object>)parsedReqBody.get('order_ids');
                        List<Object> product2Ids = (List<Object>)parsedReqBody.get('product2_ids'); 
                        List<Object> orderFields = (List<Object>)parsedReqBody.get('order_fields');
                        List<Object> product2Fields = (List<Object>)parsedReqBody.get('product2_fields');  

                        //do more stuff before this
                        response.addHeader('Content-Type', 'application/json');
                        response.statusCode = 200;
                        Map<String, Object> resBody = new Map<String, Object>();
                        Map<String, List<Object>> orderDetails = new Map<String,List<Object>>();
                        orderDetails.put('ORDER_UID',new List<Object>());
                        resBody.put('orders',orderDetails);
                        response.responseBody = Blob.valueOf(JSON.serialize(resBody));//serialized list of order query data goes here
                            
                   /*  } else {
                        throw new RequestAuthenticationFailedException('Authentication header values do not match');
                    }
                } else {
                    throw new NotConnectedToStripeException('This org is not connected to Strip, contact your Salesforce administrator for further assistance.');
                }
            } else {
                throw new SetupIsNotCompleteException('Setup is not complete, contact your Salesforce administrator for further assistance.');
            } */
        } catch(Exception e) {
            errorLogger.create('Rest Service Event', e.getMessage(), String.valueOf(e.getLineNumber()), restRequest.requestBody.toString());
            response.statusCode = 500;
            Map<String, Object> resBody = new Map<String, Object>();
            resBody.put('error', 'APEX_ERROR');
            resBody.put('message', e.getLineNumber() + ': ' + e.getMessage());
            response.responseBody = Blob.valueOf(JSON.serialize(resBody));
            return;
        }
    }
}
