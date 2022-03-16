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

      /*gets us a map of all the fields and only gets lookup relationships that we care about*/
    public static Set<Object> getListOfFieldsByObject(String objectApiName) {
        Set<Object> objectFieldMapList = new Set<Object>(); 
        //check if we have access to the object 
        if (Schema.describeSObjects(new List<String>{ objectApiName })[0].isAccessible()) {
            //get all fields on object
            Map<String, Schema.SObjectType> objectSchemaMap = Schema.getGlobalDescribe();
            Schema.SObjectType objectSchema = objectSchemaMap.get(objectApiName);
            Map<String, Schema.SObjectField> objectApiFieldMap = objectSchema.getDescribe().fields.getMap();
            Boolean isObjectCustom = (Boolean)objectSchema.getDescribe().isCustom();
            for (String objectFieldName : objectApiFieldMap.keySet()) {
                 Map<String, Object> objectFieldMap = new Map<String,Object>(); 
                //check we have access to the field
                if(objectApiFieldMap.get(objectFieldName).getDescribe().isAccessible()) {
                    //get field type 
                    String typeOfField = String.valueOf(objectApiFieldMap.get(objectFieldName).getDescribe().getType()).toLowercase();
                    //check if field is a custom field or not
                    Boolean isFieldCustom = (Boolean)objectApiFieldMap.get(objectFieldName).getDescribe().isCustom();
                    //check if its a lookup field
                    if(typeOfField == 'reference') {
                        //get list of objects that the field looks up to 
                        List<Object> objectsFieldReferences = objectApiFieldMap.get(objectFieldName).getDescribe().getReferenceTo();
                        //iterate over objects the field looks up to and only add to list if its a field we care about
                        for(Object referencedObject : objectsFieldReferences) {
                            //make sure the field only looks up to 1 object before adding it to map 
                            if(objectsFieldReferences.size() == 1) {
                                String stringReferncedObjecect = String.valueOf(referencedObject);
                                String relationshipName = objectApiFieldMap.get(objectFieldName).getDescribe().getRelationshipName();

                                //use this list to target lookup fields based on the object the look up to 
                                List<String> listOfImportantObjectsApiNames = new List<String> {
                                    'Product2', 
                                    'Contract',
                                    'Order',
                                    'OrderItem',
                                    'PricebookEntry',
                                    'SBQQ__QuoteLine__c',
                                    'Account',
                                    'Contact',
                                    'SBQQ__Subscription__c',
                                    'SBQQ__Quote__c',
                                    'Opportunity',
                                    'Pricebook2',
                                    'ContractLineItem',
                                    'OpportunityLineItem',
                                    'ProductItem'
                                };
                                if(listOfImportantObjectsApiNames.contains(stringReferncedObjecect)) {
                                    objectFieldMap.put('label',objectApiFieldMap.get(objectFieldName).getDescribe().getLabel());
                                    objectFieldMap.put('value',objectFieldName);
                                    objectFieldMap.put('isFieldCustom',isFieldCustom);
                                    objectFieldMap.put('isObjectCustom',isObjectCustom);
                                    objectFieldMap.put('type',typeOfField);
                                    objectFieldMap.put('object',stringReferncedObjecect);
                                    objectFieldMap.put('relationshipName',relationshipName);
                                } 
                            } else {
                                //in this case we now the field looks up to more than one object and we will not support traversals 
                                objectFieldMap.put('label',objectApiFieldMap.get(objectFieldName).getDescribe().getLabel());
                                objectFieldMap.put('value',objectFieldName);
                                objectFieldMap.put('isFieldCustom',isFieldCustom);
                                objectFieldMap.put('isObjectCustom',isObjectCustom);
                                objectFieldMap.put('type','polymorphic');
                            }
                        }
                    } else {
                        //in this case we know the field is not a look up so we add to field map
                        objectFieldMap.put('label',objectApiFieldMap.get(objectFieldName).getDescribe().getLabel());
                        objectFieldMap.put('value',objectFieldName);
                        objectFieldMap.put('type',typeOfField);
                        objectFieldMap.put('isFieldCustom',isFieldCustom);
                        objectFieldMap.put('isObjectCustom',isObjectCustom);
                    }
                }
                //add field map to list to be returned 
                if(objectFieldMap != null && !objectFieldMap.isEmpty())objectFieldMapList.add(objectFieldMap);
            }
        }
        objectFieldMapList.remove(null);
        return objectFieldMapList;
    }

    public static Boolean isCpqEnabled() {
        //checks if the namespace of the CPQ package exists (all CPQ package namespaces start with 'SBQQ'
        // if a non-admin installs this package, this check would fail
        // we are using a wildcard because there are variants of the CPQ package which have a unique namespace (SBQQW)
        Integer sbqqPackageCount = [SELECT COUNT() FROM PackageLicense WHERE NamespacePrefix LIKE 'SBQQ%'];
        return sbqqPackageCount > 0;
    }

    /* gets the organization type (sandbox or production) we are using a custom setting object so this record
    will not be copied over to newly generated sandboxes instead a new record will be created for each instance*/
    @AuraEnabled
    public static String setOrgType() {
        responseData rd = new responseData();
        try {
            Organization_Type__c orgTypeSetting = Organization_Type__c.getOrgDefaults();
            //only set value if the record doesn't exist yet
            if (orgTypeSetting.Id == null) {
                orgTypeSetting.isSandbox__c = (Boolean)constants.IS_SANDBOX;
                upsert orgTypeSetting;
            } 
            /*sets custom setting record that controls weather sending errors to sentry is live or not*/
            Sentry_Active_Config__c activateSentrySetting = Sentry_Active_Config__c.getOrgDefaults();
            if (activateSentrySetting.Id == null) {
                activateSentrySetting.IsDebug__c = false;
                activateSentrySetting.Environment_Name__c = (String)constants.ORG_ID;
                activateSentrySetting.IsIssueCreationDisabled__c = false;
                activateSentrySetting.Sentry_Config__c = 'Default';
                upsert activateSentrySetting;
            } 
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd.getJsonString();
    }

    public static List<Setup_Connection_Data__mdt> getStripeConnectionKey() {
        return [SELECT Global_Key__c
                FROM Setup_Connection_Data__mdt
                LIMIT 1];  
    }

    public static HttpResponse makeCallout(String endPoint, String method){
        return makeCallout(endPoint, method, null);
    }

    public static HttpResponse makeCallout(String endPoint, String method, String body){
        Http httpMethod = new Http();
        HttpRequest request = new HttpRequest();
 
        Integer requestTimeoutLength = 500;
        request.setEndpoint(endPoint);
        request.setMethod(method);
        request.setTimeout(requestTimeoutLength);
 
        if(body != null) {
            //these headers are required for the API to respond properly 
            request.setHeader('Content-Type', 'application/json');
            request.setHeader('Accept', 'application/json');
            request.setBody(body);
        }
     
        /* these standard headers will be add to the request if none are provided
            the ruby service expects these headers for all requests. */
        List<Setup_Connection_Data__mdt> setupConfigList = getStripeConnectionKey();  

        //lets us create custom metadata record if we are in test context
        if(Test.isRunningTest()) {  
            setupConfigList = setupConfigMetadata;
        }
        
        if(!setupConfigList.isEmpty()) {
            Setup_Connection_Data__mdt setupData = setupConfigList[0];
            request.setHeader('Salesforce-Account-Id', (String)constants.ORG_ID);
            request.setHeader('Salesforce-Key', (String)setupData.Global_Key__c);
            request.setHeader('Salesforce-Type', String.valueOf(Sentry_Environment.getInstanceType()));
            request.setHeader('Salesforce-Package-Namespace', constants.NAMESPACE);
            if(!String.isEmpty(constants.NAMESPACE_API)) {
                request.setHeader('Package-Id', (String)setupAssistant.getPackageVersion());
            }
        } else {
            // TODO log.info missing metadata could be the issue
            throw new SetupIsNotCompleteException('Setup is not complete, contact your Salesforce administrator for further assistance.');
        }
    
        HttpResponse response = httpMethod.send(request);
        return response;
    }

    //Creates a test metadata record in test context
    @testVisible static List<Setup_Connection_Data__mdt> setupConfigMetadata { 
        get {
            if (setupConfigMetadata == null)
                setupConfigMetadata = getStripeConnectionKey(); 
            return setupConfigMetadata;
        } set; }
 
}
