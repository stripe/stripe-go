@isTest
public with sharing class test_errorLogger {
    @isTest
    public static void testCreate() {
        Debug_Helper_Test.TestOutputService debugOutput = new Debug_Helper_Test.TestOutputService();
        errorLogger.debugger = new Debug_Helper(true, debugOutput);

        String dump = 'dump test';

        Test.startTest();
        try {
            Object parsedJson = JSON.deserializeUntyped('bad json');
        } catch (Exception e) {
            errorLogger.create('Error', 'error', 'error', dump);
            errorLogger.create('jsonParseError', e);
        }
        Test.stopTest();

        List<String> logs = debugOutput.lines;

        // this test is stupid, but we change namespaces that I can't test so I'm just going to test the first 3 lines
        System.assert(logs.size() > 3, 'Logs have a lot of entries');

        String line1 = logs[0];
        String line2 = logs[1];
        String line3 = logs[2];

        System.assertEquals('[errorLogger] {action_taken=Error, dump=dump test, line_number=error, message=error}', line1);
        System.assertEquals('[jsonParseError] {cause=null, lineNumber=11, message=Unexpected character (\'b\' (code 98)): expected a valid value (nu', line2);
        System.assertEquals('[jsonParseError] mber, String, array, object, \'true\', \'false\' or \'null\') at input location [1,2], stackTrace=(System', line3);
    }
}
