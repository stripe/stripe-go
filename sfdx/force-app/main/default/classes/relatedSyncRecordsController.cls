public with sharing class relatedSyncRecordsController {
    @AuraEnabled
    public static String getRelatedSyncRecords(Id recordId) {
        responseData rd = new responseData(); 
        Boolean hasSyncRecords = false;
        try {
            Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults(); 
            
            if(stripeConnectRec.Id != null && stripeConnectRec.Salesforce_Connected__c && stripeConnectRec.Stripe_Connected__c) {    
                //gets a list sync records to display in component
                List<Sync_Record__c> syncRecordList = [SELECT Id, Name, Primary_Record_ID__c, Primary_Object_Type__c, Secondary_Record_ID__c, Secondary_Object_Type__c, Resolution_Status__c, Resolution_Message__c, CreatedDate
                                                        FROM Sync_Record__c
                                                        WHERE (Primary_Record_ID__c = :recordId
                                                        OR Secondary_Record_ID__c = :recordId)
                                                        WITH SECURITY_ENFORCED
                                                        ORDER BY CreatedDate DESC];   

                if (!syncRecordList.isEmpty()) {
                    hasSyncRecords = true;  
                    rd.put('syncRecordList',syncRecordList); 
                } 
            }
        } catch(QueryException e) {  
            rd.put('missingPermSet',true);  
            return rd.getJsonString();
        } catch (Exception e) {
            rd.addError(e);
        }
        rd.put('hasSyncRecords',hasSyncRecords);
        return rd.getJsonString();
    }
}
