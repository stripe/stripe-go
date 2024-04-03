@IsTest
public with sharing class Test_Tools {
    public static final String PROFILE_CHATTER_FREE_USER = 'Chatter Free User';
    
    // Include this for backwards compatibility
    // Deleting this prevents us from rolling our new QA packages
    public User createUser(String profile, String uniqueId) {
        return insertUser(profile, uniqueId);
    }

    public User insertUser(String profile, String uniqueId) {
        Profile p = [SELECT Id FROM Profile WHERE Name = :profile];
        String alias = uniqueId.substring(0, 8);
        User u = new User(
                FirstName = 'Test',
                LastName = 'User',
                Username = uniqueId + '@example.org',
                Email = uniqueId + '@example.org',
                Alias = alias,
                CommunityNickname = alias,
                TimeZoneSidKey = 'America/Los_Angeles',
                LocaleSidKey = 'en_US',
                EmailEncodingKey = 'UTF-8',
                LanguageLocaleKey = 'en_US',
                ProfileId = p.Id,
                IsActive = true);
        insert u;
        return u;
    }

    public Sync_Record__c createSyncRecord(String uniqueId) {
        Sync_Record__c testSyncRecord = new Sync_Record__c (
                Primary_Record_ID__c = '1234' + uniqueId,
                Primary_Object_Type__c = 'Order',
                Secondary_Record_ID__c = '4321' + uniqueId,
                Secondary_Object_Type__c = 'OrderItem',
                Resolution_Status__c = 'Error',
                Compound_ID__c = '1234' + uniqueId + '4321' + uniqueId
        );
        return testSyncRecord;
    }
}