@isTest
public with sharing class test_utilities {
    @isTest
    static public void testPackageVersion() {
        Test.startTest();
        String packageVersion = utilities.getPackageVersion();
        Test.stopTest();

        System.assert(packageVersion != null);
    }
}


