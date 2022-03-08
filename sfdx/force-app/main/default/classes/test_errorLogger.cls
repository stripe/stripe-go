@isTest
public with sharing class test_errorLogger {
    @isTest
    public static void testCreate() {
        String dump = 'dump test';

        Test.startTest();
        try {
            Object parsedJson = JSON.deserializeUntyped('bad json');
        } catch (Exception e) {
            errorLogger.create('Error', 'error', 'error', dump);
        }
        Test.stopTest();

        List<Error_Log__c> ErrorLogs = [SELECT Id FROM Error_Log__c];

        System.assertEquals(0, ErrorLogs.size());
    }
}
