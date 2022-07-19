/*** Stripe Toolbelt v2.0.1 ***/
public with sharing class constants {

    public string ORG_DOMAIN { get; set; }


    // Namespaces
    //fetches the namespace of the package in case we need to prefix it any where in apex for CRUD or accessing anything that will be namespaced in a package
    private static FINAL List<String> RAW_NAMESPACE = String.valueOf(constants.class).split('[.]', 2);
    public static FINAL String NAMESPACE = RAW_NAMESPACE.size() > 1 ? RAW_NAMESPACE[0] : 'c';

    // if empty we are outside of a package
    public static FINAL String NAMESPACE_API = NAMESPACE == 'c' ? '' : NAMESPACE + '__';

    // Sandbox identifcation
    // TODO ORG_ID generation should move to where the instance type stuff is managed and should be passed to sentry
    private static FINAL Organization ORG = [SELECT Id FROM Organization LIMIT 1];
    public static FINAL String ORG_ID = ORG.Id;
    private static FINAL List<Sentry_Environment.InstanceType> SANDBOX_INSTANCE_TYPES = new List<Sentry_Environment.InstanceType>{
        Sentry_Environment.InstanceType.SANDBOX,
        Sentry_Environment.InstanceType.SCRATCH_ORG
    };
    public static FINAL Boolean IS_SANDBOX = SANDBOX_INSTANCE_TYPES.contains(Sentry_Environment.getInstanceType());

    // oAuth callout urls
    public static FINAL String SALESFORCE_INSTANCE_URI = System.URL.getSalesforceBaseURL().toExternalForm();
    public static FINAL String SALESFORCE_BASE_URI = 'https://' + (IS_SANDBOX ? 'test' : 'login') + '.salesforce.com';
    public static FINAL String SALESFORCE_OAUTH_URI = SALESFORCE_BASE_URI + '/services/oauth2';

    //Stripe Ruby URIS
    public static FINAL String RUBY_SERVICE_BASE_URI = 'https://stripe-force.herokuapp.com';

    //Setup Data Record Name
    public static FINAL String SETUP_DATA_RECORD_NAME = 'SetupData';
    public static FINAL String PACKAGED_PERMISSION_SET_NAME = 'Stripe Connector Integration User';

    // this seemingly-useless code is used for the bootstrap process in the setup.page
    public String getNamespace() {
        return constants.NAMESPACE;
    }

    public constants() {
        // empty
    }
}
