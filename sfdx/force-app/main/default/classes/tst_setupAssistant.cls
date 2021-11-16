@isTest
public with sharing class tst_setupAssistant {
    @isTest
    static public void testSetupAssistant() {
        Test.startTest();
        setupAssistant con = new setupAssistant();
        Test.stopTest();

        System.assertEquals(constants.NAMESPACE, con.namespace);
    }

}
