/**
 * Created by jmather-c on 1/4/23.
 */

@IsTest
public with sharing class Debug_Helper_Test {
    public class TestOutputService implements Debug_Helper.OutputService {
        public List<String> lines = new List<String>();

        public TestOutputService() {}

        public void debug(String str) {
            lines.add(str);
        }
    }

    @IsTest
    static void VerifyDisabledWorks() {
        TestOutputService output = new TestOutputService();
        Debug_Helper helper = new Debug_Helper(false, output);

        helper.debug('a', 'b');

        System.assertEquals(true, output.lines.isEmpty(), 'Recorded no output');
    }

    @IsTest
    static void VerifyEnabledWorks() {
        TestOutputService output = new TestOutputService();
        Debug_Helper helper = new Debug_Helper(true, output);

        helper.debug('a', 'b');

        System.assertEquals(false, output.lines.isEmpty(), 'Recorded output');
        System.assertEquals('[a] b', output.lines[0], 'Recorded correct output');
    }

    @IsTest
    static void VerifyObjectOutputWorks() {
        TestOutputService output = new TestOutputService();
        Debug_Helper helper = new Debug_Helper(true, output);

        Map<String, String> testObj = new Map<String, String> { 'a' => 'b' };

        helper.debug('here', testObj);

        System.assertEquals(false, output.lines.isEmpty(), 'Recorded output');
        System.assertEquals('[here] {a=b}', output.lines[0], 'Recorded correct output');
    }

    @IsTest
    static void VerifyLongStringBreakdownWorks() {
        TestOutputService output = new TestOutputService();
        Debug_Helper helper = new Debug_Helper(true, output);

        String shortStr = 'long string breakdown ';
        Integer repeats = (Debug_Helper.lengthLimit / shortStr.length()) + 1;
        String fullStr = shortStr.repeat(repeats);

        helper.debug('a', fullStr);

        System.assertEquals(false, output.lines.isEmpty(), 'Recorded output');
        System.assertEquals(2, output.lines.size(), 'Recorded correct number of lines');
    }

    @IsTest
    static void VerifyNullThrowsException() {
        Exception e = null;

        try {
            Debug_Helper nullHelper = new Debug_Helper(false, null);
        } catch (Exception ex) {
            e = ex;
        }

        System.assert(e instanceof Debug_Helper.DebugHelperException);
        System.assertEquals('Instantiated Debug_Helper with null ouputService', e.getMessage());
    }
}