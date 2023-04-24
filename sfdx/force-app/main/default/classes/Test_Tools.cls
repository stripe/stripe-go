/**
 * Created by jmather-c on 3/9/23.
 */

@IsTest
public with sharing class Test_Tools {
    public static String PROFILE_CHATTER_FREE_USER = 'Chatter Free User';
    public User createUser(String profile, String uniqueId) {
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
}