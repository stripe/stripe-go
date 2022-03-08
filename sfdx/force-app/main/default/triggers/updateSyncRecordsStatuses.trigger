trigger updateSyncRecordsStatuses on Sync_Record__c (after update) {

    try {
        //Use Lists to store the sync record and primary objced Id for comparison later if it is resolved 
        Set<String> newlyResolvedSyncRecordPrimaryObjectIdList = new  Set<String>();

        /*Iterate through all records that have been updated from Error or skipped to resolved or success and store the 
        Primary Object Id of each one in a list for limiting our query */
        for (Sync_Record__c syncRecord: Trigger.new) { 
            Sync_Record__c oldSyncRecord = trigger.oldMap.get(syncRecord.Id);
            if((String)syncRecord.Resolution_Status__c == 'Resolved' || (String)syncRecord.Resolution_Status__c == 'Success' 
               && (oldSyncRecord.Resolution_Status__c == 'Error' || oldSyncRecord.Resolution_Status__c == 'Skipped')) {
                newlyResolvedSyncRecordPrimaryObjectIdList.add(syncRecord.Primary_Record_ID__c);
            }
        }

        //Get all sync records that are unresolved
        List<Sync_Record__c> allUnresolvedSyncRecords = [SELECT Id, Primary_Record_ID__c, Resolution_Status__c
                                                        FROM Sync_Record__c
                                                        WHERE (Resolution_Status__c != 'Resolved'
                                                        AND Resolution_Status__c != 'Success')
                                                        AND Primary_Record_ID__c IN :newlyResolvedSyncRecordPrimaryObjectIdList
                                                        AND Id NOT IN :trigger.newMap.keySet()
                                                        WITH SECURITY_ENFORCED]; 
        

        //set all sync records with the same primary object Id to resolved and add to list for upsert
        for (Sync_Record__c recordToMarkResolved : allUnresolvedSyncRecords) {
            recordToMarkResolved.Resolution_Status__c = 'Resolved';
        }

        // update all sync records found
        update allUnresolvedSyncRecords;

    } catch (Exception e) {
        errorLogger.create('updateSyncRecordsStatusesTrigger', '', string.valueOf(e.getLineNumber()), e.getMessage());
    }
}