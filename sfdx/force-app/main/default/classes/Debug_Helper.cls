public with sharing class Debug_Helper {
    // This should always be null going to production.
    // We leave this configurable for developers to expose debug logs during development.
    @TestVisible
    private Boolean enableDebug = false;

    @TestVisible
    private OutputService outputService = new DefaultOutputService();
    @TestVisible
    private static final Integer lengthLimit = 100;

    // for testability
    public interface OutputService { void debug(String str); }
    public class DefaultOutputService implements OutputService { public void debug(String str) { System.debug(str); } }
    public class DebugHelperException extends Sentry_Exception {}

    // empty constructor to use the default above to make it easy to flip for debugging.
    public Debug_Helper() {
        enableDebug = utilities.isApexDebuggingEnabled();
    }

    public Debug_Helper(Boolean enableDebug) {
        this.enableDebug = enableDebug;
    }

    public Debug_Helper(Boolean enableDebug, OutputService outputService) {
        this.enableDebug = enableDebug;
        if (outputService == null) {
            DebugHelperException e = (DebugHelperException) Sentry_ExceptionFactory.build(DebugHelperException.class);
            e.setMessage('Instantiated Debug_Helper with null ouputService');
            throw e;
        } else {
            this.outputService = outputService;
        }
    }

    public void debug(String source, Object data) {
        if (data instanceof Sentry_Exception) {
            Map<String, Object> exceptionMap = new Map<String, Object>();
            exceptionMap.put('message', ((Sentry_Exception)data).getMessage());
            exceptionMap.put('stackTrace', ((Sentry_Exception)data).getCustomStackTraceAsString());
            exceptionMap.put('cause', ((Sentry_Exception)data).getCause());
            exceptionMap.put('lineNumber', ((Sentry_Exception)data).getLineNumber());
            exceptionMap.put('type', ((Sentry_Exception)data).getTypeName());
            exceptionMap.put('context', ((Sentry_Exception)data).context);
            debug(source, exceptionMap);
            return;
        }
        if (data instanceof Exception) {
            Map<String, Object> exceptionMap = new Map<String, Object>();
            exceptionMap.put('message', ((Exception)data).getMessage());
            exceptionMap.put('stackTrace', ((Exception)data).getStackTraceString());
            exceptionMap.put('cause', ((Exception)data).getCause());
            exceptionMap.put('lineNumber', ((Exception)data).getLineNumber());
            exceptionMap.put('type', ((Exception)data).getTypeName());
            debug(source, exceptionMap);
            return;
        }
        if (data == null) {
            debug(source, 'null');
        } else {
            // apparently currently the toString() on Integer is throwing exceptions?
            debug(source, String.valueOf(data));
        }
    }

    public void debug(String source, String message) {
        if (enableDebug == false) {
            return;
        }

        if (message == null) {
            outputService.debug('[' + source + '] ' + '//null//');
            return;
        }

        if (message.length() < lengthLimit) {
            outputService.debug('[' + source + '] ' + message);
            return;
        }

        while (message.length() > 0) {
            if (message.length() > lengthLimit) {
                String nextMsg = message.substring(0, lengthLimit);
                message = message.substring(lengthLimit);
                outputService.debug('[' + source + '] ' + nextMsg.trim());
            } else {
                outputService.debug('[' + source + '] ' + message.trim());
                message = '';
            }
        }
    }
}