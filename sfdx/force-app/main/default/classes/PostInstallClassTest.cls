/**
 * Created by rpena on 3/12/2024.
 */

@IsTest
public with sharing class PostInstallClassTest {

    @IsTest
    static void testInstallScript() {
        Setup_Connection_Data__mdt testSetupData = test_setupAssistant.getTestStripeConnectionKey();
        test_setupAssistant.setTestGlobalKey(testSetupData);
        test_setupAssistant.insertTestConnectedRecord();
        Test.setMock(HttpCalloutMock.class, new test_setupAssistant.UnifiedConfigMock());

        PostInstallClass postinstall = new PostInstallClass();

        Test.startTest();
        Test.testInstall(postinstall, new Version(1,0), true);
        Test.stopTest();

        List<Stripe_Account__c> stripeAccts = [SELECT Id, Stripe_ID__c, Is_Live_Mode__c FROM Stripe_Account__c];
        System.assertNotEquals(0, stripeAccts.size(), 'No Stripe Account found');
    }

    @IsTest
    static void testInstallScriptPackage() {
        Setup_Connection_Data__mdt testSetupData = test_setupAssistant.getTestStripeConnectionKey();
        testSetupData.Middleware_Infrastructure_Version__c = 'v2';
        test_setupAssistant.setTestGlobalKey(testSetupData);
        test_setupAssistant.insertTestConnectedRecord();
       
        PostInstallClass postinstall = new PostInstallClass();
        Test.startTest();
        Test.testInstall(postinstall, new Version(3,0), true);
        Test.stopTest();

        List<Stripe_Account__c> stripeAccts = [SELECT Id, Stripe_ID__c, Is_Live_Mode__c FROM Stripe_Account__c];
        System.assertNotEquals(0, stripeAccts.size(), 'No Stripe Account found');
    }
}