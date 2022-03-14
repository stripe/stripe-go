/**
 * Created by jmather on 2019-08-12.
 */

@IsTest
public with sharing class Sentry_ExceptionFactoryTest {
    class BasicException extends Exception {}
    class ExtendedException extends Sentry_Exception {}

    @IsTest
    static void TestMakingExceptions() {
        Sentry_Exception ex;

        Test.startTest();

        Exception sentryException = (Exception) Sentry_ExceptionFactory.build();
        Exception sentryExceptionFromString = Sentry_ExceptionFactory.build('Sentry_Exception');
        Exception basic = Sentry_ExceptionFactory.build(BasicException.class);
        Exception extended = Sentry_ExceptionFactory.build(ExtendedException.class);
        try {
            Sentry_ExceptionFactory.build(Sentry_ExceptionFactoryTest.class);
        } catch (Sentry_Exception e) {
            ex = e;
        }

        Test.stopTest();

        System.assert(sentryException instanceof Sentry_Exception);
        System.assert(sentryExceptionFromString instanceof Sentry_Exception);
        System.assert(basic instanceof BasicException);
        System.assert(extended instanceof ExtendedException);
        System.assertEquals('Sentry_ExceptionFactoryTest', ex.context.get('exceptionClass'));
    }
}