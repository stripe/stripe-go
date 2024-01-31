/**
 * Created by jmather-c on 8/30/23.
 */
import validateSharedState from '@salesforce/apex/setupAssistant.validateSharedState';
import Debugger from "c/debugger";
import {EventEmitter} from "./eventEmitter";

const DebugLog = Debugger.withContext('PostMessageListener');
const SERVICE_DELIMINATOR = '|';
const STATE_DELIMINATOR = '-';

/** @enum {string} */
export const ListenerEvents = {
    error: 'error',
    response_received: 'response_received',
    state_validated: 'state_validated',
    connection_successful: 'connection_successful',
    complete: 'complete',
}

class PostMessageListener extends EventEmitter {
    _isInitialized = false;
    _isListening = false;

    constructor() {
        super(ListenerEvents);
        DebugLog('New PostMessageListener instance created');
        this._initializeMessageListener();
    }

    /**
     *
     * @param {ListenerEvents} event_name
     * @param {Function<Event>} callback
     */
    on(event_name, callback) { super.on(event_name, callback); }

    /**
     *
     * @param {ListenerEvents} event_name
     * @param {Function<Event>} callback
     */
    off(event_name, callback) { super.off(event_name, callback); }

    /**
     *
     * @param {string} origin
     */
    listenFor(origin) {
        DebugLog('Listening for post messages from ' + origin);
        this.updateOrigin(origin);
        this._isListening = true;
    }

    detach() {
        DebugLog('Detaching post message listener');
        this._detachMessageListener();
    }

    /**
     *
     * @param {string} origin
     */
    updateOrigin(origin) {
        DebugLog('Updating origin to ' + origin);
        this._origin = origin;
    }

    /**
     *
     * @param {string} event_name
     * @param {Object} detail
     * @protected
     * @emits {Event}
     */
    _fireEvent(event_name, detail) { super._fireEvent(event_name, detail); }

    _initializeMessageListener() {
        DebugLog('_initializeMessageListener', 'Initializing message listener');
        if (this._isInitialized) {
            DebugLog('_initializeMessageListener', 'Already initialized');
            return;
        }

        if (this._boundPostMessageListener === undefined) {
            DebugLog('_initializeMessageListener', 'Binding post message listener');
            this._boundPostMessageListener = this._listener.bind(this);
        }

        DebugLog('_initializeMessageListener', 'Attaching message listener');
        window.addEventListener("message", this._boundPostMessageListener);
        this._isInitialized = true;
    }

    _detachMessageListener() {
        DebugLog('_detachMessageListener', 'Detaching message listener');
        if (this._isInitialized === false) {
            DebugLog('_detachMessageListener', 'Message listener already detached');
            return;
        }

        DebugLog('_detachMessageListener', 'Removing message listener');
        window.removeEventListener('message', this._boundPostMessageListener);
        this._isInitialized = false;
    }

    _listener(event) {
        DebugLog('_listener', 'Got event: ' + event);
        DebugLog('_listener', 'Got event data: ', {data: event.data, origin: event.origin});

        if(this._isListening === false) {
            DebugLog('_listener', 'Not listening');
            return;
        }

        this._isListening = false;

        if (event.origin !== this._origin) {
            this._fireEvent(ListenerEvents.error, new Error('bad_message_origin'));
            DebugLog('_listener', `Bad post message origin. Expected ${this._origin} but got ${event.origin}`);
            return;
        }

        const pieces = event.data.split(STATE_DELIMINATOR);
        const result = pieces.splice(0, 1)[0];
        const service = pieces.splice(0, 1)[0];
        const state = pieces.join('-');
        this._fireEvent(ListenerEvents.response_received, {
            raw: event.data,
            result,
            service,
            state,
        });

        DebugLog({result, service, state});

        const postValidation = (validatedState) => {
            DebugLog('_listener', 'postValidation', 'validatedState', validatedState);
            this._fireEvent(ListenerEvents.state_validated, validatedState);

            // We assume a service looks like <system>_<id>
            // We do this instead of splitting on '_' in case the id contains an underscore (like stripe's acct_<id>).
            const [system, id] = service.split(SERVICE_DELIMINATOR);
            const details = {
                service,
                system,
                id,
                result,
                state: validatedState,
            }
            DebugLog('_listener', 'postValidation', 'details', details);

            if (result === 'connection_successful') {
                DebugLog('_listener', 'postValidation', 'firing connection_successful');
                this._fireEvent(ListenerEvents.connection_successful, details);
            } else {
                DebugLog('_listener', 'postValidation', 'firing error');
                this._fireEvent(ListenerEvents.error, details);
            }

            DebugLog('_listener', 'postValidation', 'firing complete');
            this._fireEvent(ListenerEvents.complete, details);
        };

        return validateSharedState({ service: service, state })
            .then(postValidation)
            .catch((error) => {
                DebugLog('_listener', 'validateSharedState', 'error', error);
                this._fireEvent(ListenerEvents.error, error);
            });
    }
}

export const MessageListener = new PostMessageListener();