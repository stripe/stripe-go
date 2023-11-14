/**
 * Created by jmather-c on 6/1/23.
 */
import isJSDebugEnabled from '@salesforce/apex/setupAssistant.isJSDebugEnabled';
import isOAuthAutoCloseWindowEnabled from '@salesforce/apex/setupAssistant.isOAuthAutoCloseWindowEnabled';

class Debugger {
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
        return function (var_args) {
            const newargs = [context];
            for (let i = 0; i < arguments.length; i++) {
                newargs.push(arguments[i]);
            }
            dat.exec(console.log, newargs);
        }
    }

    /**
     * @param {...*} var_args
     */
    static log(var_args) {
        this.exec(console.log, arguments);
    }

    /**
     * @param {...*} var_args
     */
    static error(var_args) {
        this.exec(console.error, arguments);
    }

    /**
     * @param {...*} var_args
     */
    static warn(var_args) {
        this.exec(console.warn, arguments);
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
        for (let i = 0; i < data.length; i++) {
            try {
                if (data[i] instanceof Proxy) {
                    ret.push(JSON.parse(JSON.stringify(data[i])));
                } else {
                    ret.push(data[i]);
                }
            } catch (e) {
                ret.push(data[i]);
            }
        }
        return ret;
    }
}

export {Debugger}