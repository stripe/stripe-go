/**
 * @description       : 
 * @author            : jmather
 * @group             : 
 * @last modified on  : 02-11-2022
 * @last modified by  : Stripe
 * Modifications Log 
 * Ver   Date         Author    Modification
 * 1.0   08-10-2019   jmather   Initial Version
 * 1.1   02-11-2022   Stripe    Moved mock callout into this class
**/


@IsTest
public with sharing class SentryTest {
    private class MyException extends Sentry_Exception {}
    private class SYS_CustomException extends Exception {}
    @IsTest
    static void Record() {
        Exception ex;

        Test.setMock(HttpCalloutMock.class, new Sentry_ApiMock());

        Sentry_Active_Config__c config = Sentry_Active_Config__c.getOrgDefaults();
        config.IsIssueCreationDisabled__c = false;
        config.IsDebug__c = true;
        config.Sentry_Config__c = 'Test';
        config.Environment_Name__c = 'Test';
        insert config;

        Test.startTest();

        try {
           try{    
                MyException testException = (MyException) Sentry_ExceptionFactory.build(MyException.class);
                testException.setMessage('Something broke.');
                throw testException;
            } catch (Sentry_Exception sentryTestException) {
                Sentry.record(sentryTestException);
                throw sentryTestException;
            }
        } catch (Exception e) {
            ex = e;
        }

        Test.stopTest();

        System.assertNotEquals(null, ex);
    }
    
    @IsTest
    static void RecordwithCustomException() 
    {
        Exception ex;

        Test.setMock(HttpCalloutMock.class, new Sentry_ApiMock());

        Sentry_Active_Config__c config = Sentry_Active_Config__c.getOrgDefaults();
        config.IsIssueCreationDisabled__c = false;
        config.IsDebug__c = true;
        config.Sentry_Config__c = 'Test';
        config.Environment_Name__c = 'Test';
        insert config;

        Test.startTest();
        List<Sentry_LogMessage> lExtraMessages = new List<Sentry_LogMessage>();
     	lExtraMessages.add(new Sentry_LogMessage(Sentry_LogMessage.LEVEL_INFO, Sentry_LogMessage.LEVEL_DEBUG, Sentry_LogMessage.LEVEL_DEBUG,'Testing debug 1', new Map<String, String>()));   
        Sentry.record('Custom Issue Title',new SYS_CustomException('Custom Exception'), lExtraMessages);
        Test.stopTest();

        System.assertEquals(null, ex);
    }
    
    @IsTest
    static void RecordCustomExceptionWithoutSendingToSentry() 
    {
        Exception ex;

        Test.setMock(HttpCalloutMock.class, new Sentry_ApiMock());

        Sentry_Active_Config__c config = Sentry_Active_Config__c.getOrgDefaults();
        config.IsIssueCreationDisabled__c = true;
        config.IsDebug__c = true;
        config.Sentry_Config__c = 'Test';
        config.Environment_Name__c = 'Test';
        insert config;

        Test.startTest();
        List<Sentry_LogMessage> lExtraMessages = new List<Sentry_LogMessage>();
     	lExtraMessages.add(new Sentry_LogMessage(Sentry_LogMessage.LEVEL_INFO, Sentry_LogMessage.LEVEL_DEBUG, Sentry_LogMessage.LEVEL_DEBUG,'Testing debug 1', new Map<String, String>()));   
        Sentry.record('Custom Issue Title',new SYS_CustomException('Custom Exception'), lExtraMessages);
        Test.stopTest();

        System.assertEquals(null, ex);
    }
    
    private class Sentry_ApiMock implements HttpCalloutMock {
        public HttpResponse respond(HttpRequest param1) {
            HttpResponse response = new HttpResponse();
            response.setStatusCode(200);
            return response;
        }
    }
}