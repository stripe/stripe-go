/**
 * Created by jmather-c on 1/4/23.
 */

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
        this.outputService = outputService;

        if (this.outputService == null) {
            DebugHelperException e = (DebugHelperException) Sentry_ExceptionFactory.build(DebugHelperException.class);
            e.setMessage('Instantiated Debug_Helper with null ouputService');
            throw e;
        }
    }

    public void debug(String source, Object data) {
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
            debug(source, data.toString());
        }
    }

    public void debug(String source, String message) {
        if (enableDebug == false || String.isBlank(source) || String.isBlank(message)) {
            return;
        }
        
        if (message.length() < lengthLimit) {
            this.outputService.debug('[' + source + '] ' + message);
            return;
        }

        while (String.isNotBlank(message)) {
            if (message.length() > lengthLimit) {
                String nextMsg = message.substring(0, lengthLimit);
                message = message.substring(lengthLimit);
                this.outputService.debug('[' + source + '] ' + nextMsg.trim());
            } else {
                this.outputService.debug('[' + source + '] ' + message.trim());
                message = '';
            }
        }
    }
}