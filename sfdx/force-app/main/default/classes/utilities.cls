public with sharing class utilities {
    public class SetupIsNotCompleteException extends Exception {}

    public static sObject getSobjectByMap(Map<String, Object> sourceMap, String sObjectName) {
        return getSobjectByMap(sourceMap, sObjectName, constants.NAMESPACE_API);
    }

    /*
        Takes in a map and converts it into the specified sObject type.
        This function will strip away any keys that don't match a field on the sObject
        so you don't have to worry about it being dirty.
    */
    public static sObject getSobjectByMap(Map<String, Object> sourceMap, String sObjectName, String namespace) {
        if (!Schema.getGlobalDescribe().containsKey(sObjectName)) {
            if (Schema.getGlobalDescribe().containsKey(namespace + sObjectName)) {
                sObjectName = namespace + sObjectName;
            } else {
                return null;
            }
        }

        Map<String, Schema.SObjectField> fieldMap = Schema.getGlobalDescribe().get(sObjectName).getDescribe().fields.getMap();
        Map<String, Object> objMap = new Map<String, Object>();

        for (String field : sourceMap.keySet()) {
            String value = String.valueOf(sourceMap.get(field));

            //only retain the value if it exists on our object
            if (!fieldMap.containsKey(field)) {
                if (fieldMap.containsKey(namespace + field)) {
                    field = namespace + field;
                } else {
                    continue;
                }
            }

            objMap.put(field, String.isEmpty(value) ? null : value);
        }

        return (sObject)JSON.deserialize(JSON.serialize(objMap), Type.forName(sObjectName));
    }

    public static HttpResponse makeCallout(String endPoint, String method){
        return makeCallout(endPoint, method, null, null);
    }
 
    public static HttpResponse makeCallout(String endPoint, String method, String body){
        return makeCallout(endPoint, method, body, null);
    }
 
    public static HttpResponse makeCallout(String endPoint, String method, Map<String, String> headers){
        return makeCallout(endPoint, method, null, headers);
    }
 
    public static HttpResponse makeCallout(String endPoint, String method, String body, Map<String, String> headers){
        Http h = new Http();
        HttpRequest req = new HttpRequest();
 
        req.setEndpoint(endPoint);
        req.setMethod(method);
        req.setTimeout(60000);
 
        if(body != null) {
            //these headers are required for the API to respond properly 
            req.setHeader('Content-Type', 'application/json');
            req.setHeader('Accept', 'application/json');
            req.setBody(body);
        }
 
        if(headers != null) {
            for(String header : headers.keySet()) {
                req.setHeader(header, headers.get(header));
            }
        } else {
            List<Setup_Connection_Data__mdt> setupConfigList = [SELECT Salesforce_Connected__c, Stripe_Connected__c, Global_Key__c, API_Key__c
                                                                FROM Setup_Connection_Data__mdt
                                                                WITH SECURITY_ENFORCED];    
            if(!setupConfigList.isEmpty()) {
                Setup_Connection_Data__mdt setupData = setupConfigList[0];
                String orgId = constants.ORG_ID;
                //String globalKey = setupData.Global_Key__c; 
                //String packageID = setupAssistant.getPackageVersion();
                String globalKey = 'dee37489-9f31-91d5';
                String packageGuidKey = setupData.API_Key__c;
                req.setHeader('Salesforce-Account-Id', orgId);
                req.setHeader('Salesforce-Key', globalKey);
                //req.setHeader('Package-Id', packageID);
            } else {
                throw new SetupIsNotCompleteException('Setup is not complete, contact your Salesforce administrator for further assistance.');
            }
        }
 
        HttpResponse res = h.send(req);
 
        return res;
    }
 
 
}
