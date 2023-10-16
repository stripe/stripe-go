/**
 * Created by jmather-c on 4/27/23.
 */

public with sharing class ConfigExportVFController {
    ConfigPayloadHelper payloadHelper = new ConfigPayloadHelper();
    Rest_Service_Helper restServiceHelper = new Rest_Service_Helper();
    List<Object> relatedObjectList = new List<Object>();
    List<Sync_Record__c> syncRecords = new List<Sync_Record__c>();

    public ConfigExportVFController() {
        List<String> syncRecordIds = new List<Id>();

        // if there is a sync record id on the page, export this in the config
        String syncRecordId = ApexPages.currentPage().getParameters().get('recordId');
        if (syncRecordId != null) {
            syncRecordIds.add(syncRecordId);
        }
        Set<Object> accessErrors = new Set<Object>();
        Map<String, Set<Id>> lookupObjIds = new Map<String, Set<Id>>();
        Map<String, List<Sync_Record__c>> primaryRecMap = new Map<String, List<Sync_Record__c>>();
        Map<String, List<Sync_Record__c>> secondaryRecMap = new Map<String, List<Sync_Record__c>>();
        Map<String, Map<String, Object>> responseObjects = new Map<String, Map<String, Object>>();
        
        Set<String> syncRecordFields = restServiceHelper.getAccessibleFields('Sync_Record__c', accessErrors);
        
        // for every sync record, get the primary and secondary objects
        List<SObject> records = getRecordData('Sync_Record__c', syncRecordFields, syncRecordIds);
        for (SObject record : records) {
            Sync_Record__c syncRecord = (Sync_Record__c)record;
            syncRecords.add(syncRecord);
            if (syncRecord.Primary_Record_ID__c != null) {
                if (lookupObjIds.containsKey(syncRecord.Primary_Object_Type__c) == false) {
                    lookupObjIds.put(syncRecord.Primary_Object_Type__c, new Set<Id>());
                }
                lookupObjIds.get(syncRecord.Primary_Object_Type__c).add(syncRecord.Primary_Record_ID__c);

                if (primaryRecMap.containsKey(syncRecord.Primary_Record_ID__c) == false) {
                    primaryRecMap.put(syncRecord.Primary_Record_ID__c, new List<SObject>());
                }
                primaryRecMap.get(syncRecord.Primary_Record_ID__c).add(syncRecord);
            }
            if (syncRecord.Secondary_Record_ID__c != null) {
                if (lookupObjIds.containsKey(syncRecord.Secondary_Object_Type__c) == false) {
                    lookupObjIds.put(syncRecord.Secondary_Object_Type__c, new Set<Id>());
                }
                lookupObjIds.get(syncRecord.Secondary_Object_Type__c).add(syncRecord.Secondary_Record_ID__c);

                if (secondaryRecMap.containsKey(syncRecord.Secondary_Record_ID__c) == false) {
                    secondaryRecMap.put(syncRecord.Secondary_Record_ID__c, new List<SObject>());
                }
                secondaryRecMap.get(syncRecord.Secondary_Record_ID__c).add(syncRecord);
            }
        }

        for (String objName : lookupObjIds.keySet()) {
            Set<String> fields = restServiceHelper.getAccessibleFields(objName, accessErrors);
            List<SObject> objs = getRecordData(objName, fields, new List<Id>(lookupObjIds.get(objName)));
            for (SObject obj : objs) {
                if (primaryRecMap.containsKey(obj.Id)) {
                    for (Sync_Record__c syncRec : primaryRecMap.get(obj.Id)) {
                        if (responseObjects.containsKey(syncRec.Id) == false) {
                            responseObjects.put(syncRec.Id, new Map<String, Object>());
                        }
                        responseObjects.get(syncRec.Id).put('primary', obj);
                    }
                }
                if (secondaryRecMap.containsKey(obj.Id)) {
                    for (Sync_Record__c syncRec : secondaryRecMap.get(obj.Id)) {
                        if (responseObjects.containsKey(syncRec.Id) == false) {
                            responseObjects.put(syncRec.Id, new Map<String, Object>());
                        }
                        responseObjects.get(syncRec.Id).put('secondary', obj);
                    }
                }
            }
        }

        relatedObjectList = new List<Object>();
        for (SObject record : records) {
            Sync_Record__c syncRecord = (Sync_Record__c)record;
            Map<String, Object> relatedObjects = new Map<String, Object>();
            if (responseObjects.containsKey(syncRecord.Id)) {
                relatedObjects = responseObjects.get(syncRecord.Id);
            }
            relatedObjectList.add(new Map<String, Object>{
                    'syncRecord' => syncRecord,
                    'relatedObjects' => relatedObjects
            });
        }
    }

    public List<SObject> getRecordData(String objName, set<String> fields, List<String> ids) {
        List<SObject> records = new List<SObject>();
        if (ids.size() > 0) {
            String query = 'SELECT ' + String.join(new List<String>(fields), ',') + ' FROM ' + objName + ' WHERE Id IN :ids';
            return Database.query(query);
        }
        return records;
    }

    public String getConfig() {
        responseData resp = getSystemConfig();
        // only if there is a sync record id on the page, export this in the config
        if (this.syncRecords.size() > 0) {
            resp.put('syncRecordData', relatedObjectList);
        }
        return resp.getResultsJsonString();
    }

    // gets the account configuration from ruby service
    public responseData getSystemConfig() {
        Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();
        Boolean isConnected = utilities.isConnected(stripeConnectRec);
        responseData rd = new responseData();

        if (isConnected == false) {
            rd.put('error', 'Unable to connect to remote system.');
            return rd;
        }

        try {
            //construct call out to rubys configuration endpoint
            String route = constants.RUBY_SERVICE_BASE_URI + constants.RUBY_SERVICE_CONFIG_URI;
            HttpResponse response =  utilities.makeCallout(route,'GET');

            //if the call out is successful pull out all mapping objects and add to list
            Map<String, Object> responseBody;
            if(response.getStatusCode() == 200) {
                responseBody = (Map<String,Object>)JSON.deserializeUntyped(response.getBody());
                payloadHelper.extractSettingsData(responseBody, rd, false);
                payloadHelper.extractFieldMappingData(responseBody, rd, false);
                payloadHelper.extractFilterSettingsData(responseBody, rd);
            } else {
                errorLogger.create('getConfiguration', String.valueOf(response.getStatusCode()), (String)response.getStatus(), 'Failed to get mapping configuration from ruby service.');
            }
        } catch (Exception e) {
            rd.addError(e);
        }
        return rd;
    }
}