/*** Stripe Response Data v2.0.0 ***/
public with sharing class responseData {
    public Boolean isSuccess;
    public Map<String, Object> results;
    public String error;

    @TestVisible
    private Debug_Helper debugger = new Debug_Helper();

    public responseData() {
        this.isSuccess = true;
        this.results = new Map<String, Object>();
    }

    public void addError(Exception e) {
        if(!Test.isRunningTest()) {
            Sentry.record(e);
        } else {
            this.debugger.debug('responseData.addError(Exception)', e);
        }

        this.error = e.getMessage();
        this.isSuccess = false;
    }

    public void addError(String error){
        this.error = error;
        this.isSuccess = false;
    }

    public void put(String key, Object val){
        this.results.put(key, val);
    }

    public String getString(String key) {
        return String.valueOf(this.results.get(key));
    }

    public String getResultsJsonString() {
        if (this.error != null) {
            this.results.put('error', this.error);
        }

        return getSanitizedJsonString(this.results);
    }

    public String getJsonString() {
        Map<String, Object> response = new Map<String, Object> {
                'isSuccess' => this.isSuccess,
                'error' => this.error,
                'results' => this.results
        };
        return getSanitizedJsonString(response);
    }

    //removes extra underscores that are typically appended to the end of the namespace
    public String getSanitizedJsonString(Object obj) {
        String jsonString = JSON.serialize(obj);
        return jsonString.replaceAll('"' + constants.NAMESPACE_API + '([a-zA-Z0-9_]+?__(c|r))":', '"$1":');
    }
}
