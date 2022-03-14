/**
 * Created by jmather on 2019-08-12.
 */
@IsTest
public with sharing class Sentry_LogTest {
    @IsTest
    static void TestLogging() {
        Map<String, String> context = new Map<String, String> { 'foo' => 'bar' };
        Sentry_Exception ex = Sentry_ExceptionFactory.build();

        Test.startTest();

        Sentry_Log.logSentry('foo');
        Sentry_Log.logSentry('foo', context);
        Sentry_Log.logError('foo');
        Sentry_Log.logError('foo', context);
        Sentry_Log.logWarn('foo');
        Sentry_Log.logWarn('foo', context);
        Sentry_Log.logInfo('foo');
        Sentry_Log.logInfo('foo', context);
        Sentry_Log.logDebug('foo');
        Sentry_Log.logDebug('foo', context);
        Sentry_Log.logTrace('foo');
        Sentry_Log.logTrace('foo', context);
        Sentry_Log.logException(ex);

        Test.stopTest();

        System.assertEquals(13, Sentry_Log.messages.size());
    }
}