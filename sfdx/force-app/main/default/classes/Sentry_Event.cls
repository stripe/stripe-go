/**
 * @description       :
 * @author            : jmather
 * @group             :
 * @last modified on  : 04-28-2021
 * @last modified by  : Jheel
 * Modifications Log
 * Ver   Date         Author    Modification
 * 1.0   08-10-2019   jmather   Initial Version
 * 1.1   04-28-2021   Jheel     Added logic to send proper detailed data to Sentry with Breadcrumbs as well
**/
public with sharing class Sentry_Event {

    public String event_id;

    public String timestamp;

    public String level = 'error';

    public String platform = 'salesforce';

    public Map<String, String> sdk = new Map<String, String> { 'name' => 'sentry.apex', 'version' => '1.1' };

    public String logger;

    public String server_name;

    public String release;

    public String message;

    public String environment;

    public String[] modules;

    public Map<String, Object> request = Sentry_Context.Basic();

    public Map<String, Object> extra = Sentry_Context.Basic();

    public Map<String, Object> user = Sentry_Context.User();

    public Map<String, Object> contexts = Sentry_Context.OsSystem();

    public Map<String, Object> tags = Sentry_Context.Basic();

    public String[] fingerprint;

    public Map<String, Object> breadcrumbs = Sentry_Context.Breadcrumbs();

    public Map<String, Object> exceptionValues = new Map<String, Object> { 'values' => new List<Map<String, Object>>() };

    public Map<String, Object> breadcrumbsValues = new Map<String, Object> { 'values' => new List<Map<String, Object>>() };

    public String messageFormatted;

    public String[] messageParams;

    public String stacktrace;

    public Sentry_Event() {
        initialize();
    }

    public Sentry_Event(Sentry_Exception ex) {
        this((Exception) ex);
        System.debug('custom' + ex.getCustomStackTraceAsString());
    }

    public Sentry_Event(Exception ex) {
        this(null, ex, null);
    }

    // Custom Method with custom issueTitle and extra messages
    public Sentry_Event(String issueTitle, Exception ex, List<Sentry_LogMessage> lExtraMessages) {
        initialize();

        Sentry_Log.logSentry('Got exception: ' + ex);
        Sentry_Log.logSentry('Got stack: ' + ex.getStackTraceString());

        Map<String, Object> exceptionData = Sentry_Context.create();
        Map<String, Object> exceptionDataStack = Sentry_Context.create();

        // For setting custom title for the issue which will be created
        if(String.isNotBlank(issueTitle)){
            exceptionData.put('type', issueTitle);
        } else {
            // Default title will be the type of exception occured.
            exceptionData.put('type', ex.getTypeName());
        }
        exceptionData.put('message', JSON.serialize(ex.getMessage()));
        List<Map<String, Object>> frames = new List<Map<String, Object>>();

        List<String> lines = ex.getStackTraceString().split('\n');
        if (ex instanceof Sentry_Exception) {
            lines = ((Sentry_Exception) ex).getCustomStackTrace();
        }

        for (Integer i = lines.size() - 1; i >= 0; i--) {
            String line = lines[i];
            Pattern p = Pattern.compile('Class\\.([^\\.]+)\\.([^:]+): line ([0-9]+), column ([0-9]+)');
            Matcher m = p.matcher(line);
            // Sample: Class.SampleClass.execute: line 36, column 1
            if (m.find()) {
                frames.add(new Map<String, Object> {
                        'class' => m.group(1),
                        'filename' => m.group(1) + '.cls',
                        'function' => m.group(1) + '.' + m.group(2) + ', line ' + m.group(3),
                        'lineno' => Integer.valueOf(m.group(3)),
                        'column' => Integer.valueOf(m.group(4)),
                        'in_app'=> true,
                        'vars' => new Map<String, Object> {
                            'column' => m.group(4),
                            'line' => Integer.valueOf(m.group(3)),
                            'class'=> m.group(1),
                            'function' => m.group(2)
                        }
                });
            } else {
                p = Pattern.compile('Class\\.([^\\.]+): line ([0-9]+), column ([0-9]+)');
                m = p.matcher(line);

                if (m.find()) {
                    System.debug('@@m:--'+m);
                    frames.add(new Map<String, Object> {
                            'class' => m.group(1),
                            'filename' => m.group(1) + '.cls',
                            'function' => m.group(1) /*+ '.' + m.group(2) */+ ', line ' + m.group(2),
                            'lineno' => Integer.valueOf(m.group(2)),
                            'column' => Integer.valueOf(m.group(3)),
                            'in_app'=> true,
                            'vars' => new Map<String, Object> {
                                'column' => m.group(3),
                                'line' => Integer.valueOf(m.group(2)),
                                'class'=> m.group(1)
                            }
                    });
                }
            }
        }

        exceptionDataStack.put('frames', frames);
        exceptionData.put('stacktrace', exceptionDataStack);

        // add "value" in exceptionData
        exceptionData.put('value', ex.getTypeName()+' '+ex.getMessage());
        exceptionData.put('cause', ex.getCause());
        exceptionData.put('line_number', ex.getLineNumber());

        addException(exceptionData);
        message = ex.getMessage();

        // For loop for adding extra messages in the issue
        if(lExtraMessages != null && !lExtraMessages.isEmpty()) {
            for(Sentry_LogMessage objSLogMessage: lExtraMessages) {
                Map<String, Object> extraData = Sentry_Context.create();
                extraData.put('type', objSLogMessage.type);
                extraData.put('category', objSLogMessage.category);
                extraData.put('level', objSLogMessage.level);
                extraData.put('timestamp', objSLogMessage.ts.formatGmt('yyyy-MM-dd') + 'T' + objSLogMessage.ts.formatGmt('HH:mm:ss') + 'Z');
                extraData.put('message', objSLogMessage.message);

                // Adding extra messages
                addbreadcrumbs(extraData);
            }
        }
    }

    private void addException(Map<String, Object> exData) {
        ((List<Map<String, Object>>)exceptionValues.get('values')).add(exData);
    }

     private void addbreadcrumbs(Map<String, Object> exData) {
        ((List<Map<String, Object>>)breadcrumbsValues.get('values')).add(exData);
    }

    private void initialize() {
        Blob b = Crypto.generateAesKey(128);
        String h = EncodingUtil.convertToHex(b);
        event_id = h.substring(0,8)+ '-' + h.substring(8,12) + '-' + h.substring(12,16) + '-' + h.substring(16,20) + '-' + h.substring(20);

        timestamp = Datetime.now().formatGmt('yyyy-MM-dd') + 'T' + Datetime.now().formatGmt('HH:mm:ss') + 'Z';
        level = Sentry_LogMessage.LEVEL_ERROR;
        environment = Sentry_Config.getEnvironmentName();
    }

    public Map<String, Object> toMessage() {
        Map<String, Object> msg = Sentry_Context.create();
        msg.put('event_id', event_id);
        msg.put('environment', environment);
        msg.put('user', user);
        msg.put('timestamp', timestamp);
        msg.put('sdk', sdk);
        msg.put('platform', platform);
        msg.put('level', level);
        msg.put('exception', exceptionValues);
        msg.put('breadcrumbs', breadcrumbsValues);
        msg.put('contexts', contexts);
        return msg;
    }
}