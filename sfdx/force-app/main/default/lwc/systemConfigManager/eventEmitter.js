/**
 * Created by jmather-c on 8/31/23.
 */

/**
 *
 * @param {string[]} events
 * @return {Object<String, Array>}
 */
function initializeEvents(events) {
    const eventContainer = {};
    for (const event_name in events) {
        eventContainer[event_name] = [];
    }
    return eventContainer;
}

export default class EventEmitter {
    _events = {};
    _lastEvent = {};

    /**
     *
     * @param {string[]} events
     * @param {boolean} [fireLastOnRegister]
     */
    constructor(events, fireLastOnRegister) {
        this._events = initializeEvents(events);
        this._fireLastOnRegister = fireLastOnRegister || false;
    }

    /**
     *
     * @param {string} event_name
     * @param {Function<Event>} callback
     */
    on(event_name, callback) {
        if (!this._events[event_name]) {
            throw new Error(`Event ${event_name} is not supported`);
        }
        this._events[event_name].push(callback);
        if (this._fireLastOnRegister && this._lastEvent[event_name]) {
            callback(this._createEvent(event_name, this._lastEvent[event_name]));
        }
    }

    /**
     *
     * @param {string} event_name
     * @param {Function<Event>} callback
     */
    off(event_name, callback) {
        if (!this._events[event_name]) {
            throw new Error(`Event ${event_name} is not supported`);
        }
        this._events[event_name] = this._events[event_name].filter(cb => cb !== callback);
    }

    /**
     *
     * @param {string} event_name
     * @param {Object} detail
     * @protected
     * @emits {Event<{event_name}>}
     */
    _fireEvent(event_name, detail) {
        if (!this._events[event_name]) {
            throw new Error(`Event ${event_name} is not supported`);
        }
        this._lastEvent[event_name] = detail;
        this._events[event_name].forEach(cb => cb(this._createEvent(event_name, detail)));
    }

    /**
     *
     * @param {string} event_name
     * @param {Object} detail
     * @return {CustomEvent<{event_name}>}
     * @protected
     */
    _createEvent(event_name, detail) {
        return new CustomEvent(event_name, { detail: JSON.parse(JSON.stringify(detail)) });
    }
}