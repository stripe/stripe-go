/**
 * Created by jmather on 2019-08-08.
 */

global with sharing class Sentry_ExceptionFactory {
    global static Sentry_Exception build() {
        return (Sentry_Exception) build(Sentry_Exception.class);
    }

    global static Exception build(String exceptionClass) {
        Type exType = Type.forName(exceptionClass);
        return build(exType);
    }

    global static Exception build(Type exType) {
        try {
            Sentry_Exception ex = (Sentry_Exception) exType.newInstance();
            try {
                SObject foo = null;
                Id fooId = foo.Id;
            } catch (Exception e) {
                String stack = e.getStackTraceString();
                List<String> lines = stack.split('\n');
                lines.remove(0);
                lines.remove(0);

                ex.setStackTrace(lines);
                return ex;
            }
        } catch (Exception e) {}

        try {
            Exception ex = (Exception) exType.newInstance();
            return ex;
        } catch (Exception e) {}

        Sentry_Exception ex = (Sentry_Exception) Sentry_ExceptionFactory.build(Sentry_Exception.class);
        ex.setMessage('Could not instantiate given class as an exception!');
        ex.context.put('exceptionClass', exType.getName());
        throw ex;
    }
}