/**
 * Created by jmather-c on 6/1/23.
 */

class Debugger {
    static debug = false;

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
        if (this.debug) {
            target.apply(console, this.sanitize(args));
        }
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