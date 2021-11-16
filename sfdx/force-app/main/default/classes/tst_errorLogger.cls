@isTest
public with sharing class tst_errorLogger {
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

        List<Error_Log__c> errorLogs = [SELECT Id FROM Error_Log__c];

        System.assertEquals(0, errorLogs.size());
    }
}
