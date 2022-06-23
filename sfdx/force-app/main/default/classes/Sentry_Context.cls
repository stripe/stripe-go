/**
 * Created by Jacob Mather <jmather@jmather.com> on 2019-03-25.
 */

public abstract class Sentry_Context {
    public static Map<String, Object> create() {
        return new Map<String, Object>();
    }

    public static Map<String, Object> Basic() {
        return new Map<String, Object>();
    }

    public static Map<String, Object> Breadcrumbs() {
        Map<String, Object> data = create();
        data.put('values', create());
        return data;
    }

    public static Map<String, Object> OsSystem() {
        return new Map<String, Object> {
                'os' => Server(),
                'runtime' => Runtime(),
                'org' => Org()
        };
    }

    public static Map<String, String> Org() {
        return new Map<String, String> {
                'instanceLabel' => Sentry_Environment.getInstanceLabel(),
                'instanceType' => String.valueOf(Sentry_Environment.getInstanceType())
        };
    }

    public static Map<String, String> Server() {
        return new Map<String, String> {
                'name' => 'Salesforce',
                // TODO should include salesforce release version here
                'version' => '1'
        };
    }

    public static Map<String, String> Runtime() {
        return new Map<String, String> {
            'type' => 'runtime',
            // TODO this returns a JSON string, which looks messy in sentry
            'PackageVersion' => utilities.getPackageVersion(),
            // TODO would be great to add CPQ version as well
            'CpqInstalled' => String.valueOf(utilities.isCpqEnabled())
        };
    }

    public static Map<String, String> User() {
        return new Map<String, String> {
            'id' => UserInfo.getUserId()
        };
    }
}