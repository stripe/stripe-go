/*before insert to ensure we are always below our set threshold of records stored*/
trigger syncRecordRetention on Sync_Record__c (before insert) { 

    try {
        Stripe_Connection__c stripeConnectRec = Stripe_Connection__c.getOrgDefaults();    
         if(stripeConnectRec.Id != null && stripeConnectRec.Sync_Record_Retention__c != null && 
            stripeConnectRec.Salesforce_Connected__c && stripeConnectRec.Stripe_Connected__c) {     
            Integer allSavedSyncRecordsCount = [SELECT COUNT()
                                                FROM Sync_Record__c
                                                WITH SECURITY_ENFORCED]; 
            
            /*if there are records found and the records to be inserted added to the records in the database 
            is greater than the threshold set we will delete all resolved and success records before inserting new ones*/
            if((allSavedSyncRecordsCount + Trigger.new.size()) > Integer.valueOf(stripeConnectRec.Sync_Record_Retention__c)) {  
                List<Sync_Record__c> syncRecordsToDelete = [SELECT Id, Resolution_Status__c, CreatedDate
                                                            FROM Sync_Record__c
                                                            WHERE (Resolution_Status__c = 'Resolved'
                                                            OR Resolution_Status__c = 'Success')
                                                            WITH SECURITY_ENFORCED
                                                            ORDER BY CreatedDate ASC]; 

                if(!syncRecordsToDelete.isEmpty()) {
                    delete syncRecordsToDelete; 
                }
            }
        }
    } catch (Exception e) {
        errorLogger.create('syncRecordRetentionTrigger', '', string.valueOf(e.getLineNumber()), e.getMessage());
    }
}