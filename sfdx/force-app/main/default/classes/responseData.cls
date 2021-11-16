/*** Strike Response Data v2.0.0 ***/
public with sharing class responseData {
    public Boolean isSuccess;
    public Map<String, Object> results;
    public String error;

    public responseData() {
        this.isSuccess = true;
        this.results = new Map<String, Object>();
    }

    public void addError(Exception e) {
        this.error = e.getMessage();
        this.isSuccess = false;
    }

    public void addError(String error){
        this.error = error;
        this.isSuccess = false;
    }

    public void put(String key, Object val){
        results.put(key, val);
    }

    public String getJsonString() {
        String jsonString = JSON.serialize(this);
        return jsonString.replaceAll('"' + constants.NAMESPACE_API + '([a-zA-Z0-9_]+?__(c|r))":', '"$1":');
    }
}
