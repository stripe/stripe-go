import isJSDebugEnabled from '@salesforce/apex/setupAssistant.isJSDebugEnabled';
import isOAuthAutoCloseWindowEnabled from '@salesforce/apex/setupAssistant.isOAuthAutoCloseWindowEnabled';

export default class Debugger {
    static config = null;
    static oauthAutoCloseWindow = null;
    static initialized = false;

    static init() {
        if (!this.initialized) {
            this.initialized = true;
            Debugger.config = isJSDebugEnabled();
            Debugger.oauthAutoCloseWindow = isOAuthAutoCloseWindowEnabled();
        }
        return Debugger.config;
    }

    static withContext(context) {
        const dat = this;
        return function (args) {
            const newargs = [context];
            for (let i = 0; i < args && args.length; i++) {
                newargs.push(args[i]);
            }
            dat.exec(console.log, newargs);
        }
    }

    /**
     * @param {...*} var_args
     */
    static log(...var_args) {
        this.exec(console.log, var_args);
    }

    /**
     * @param {...*} var_args
     */
    static error(...var_args) {
        this.exec(console.error, var_args);
    }

    /**
     * @param {...*} var_args
     */
    static warn(...var_args) {
        this.exec(console.warn, var_args);
    }

    static exec(target, args) {
        this.init()
            .then(isEnabled => {
                if (isEnabled) {
                    target.apply(console, this.sanitize(args));
                }
            })
    }

    /**
     *
     * @param {*[]} data
     * @return {*[]}
     */
    static sanitize(data) {
        const ret = [];
        for (let i = 0; i < data && data.length; i++) {
            try {
                ret.push(JSON.parse(JSON.stringify(data[i])));
            } catch (e) {
                ret.push(data[i]);
            }
        }
        return ret;
    }
}