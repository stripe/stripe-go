/**
 * Created by jmather-c on 8/10/23.
 */

import {LightningElement, api, track} from 'lwc';

/**
 * @enum {string}
 * @readonly
 */
const ConnectionStatus = {
    failed: 'failed',
    fresh: 'fresh',
    connected: 'connected',
    disconnected: 'disconnected',
    loading: 'loading',
}

const Strings = {
    connected: 'Connected',
    not_connected: 'Not Connected',
    authorize: 'Authorize',
    reauthorize: 'Reauthorize',
    success: 'success',
    offline: 'offline',
};

const Icons = {
    true: 'utility:success',
    false: 'utility:offline',
}

const AssistiveText = {
    true: 'success',
    false: 'offline',
}

const ActionLabels = {
    true: 'Reauthorize',
    false: 'Authorize',
}

const NoticeClasses = {
    true: 'slds-notify slds-notify_alert slds-theme_alert-texture slds-theme_success',
    false: 'slds-notify slds-notify_alert slds-theme_alert-texture slds-theme_offline',
}

export default class SystemConnectionsStepItem extends LightningElement {
    @api system;
    @api status;

    triggerConnect() {
        let evt = new CustomEvent("connect", {
            detail: {
                system: this.system,
            }
        });
        this.dispatchEvent(evt);
    }

    get _isConnected() {
        return this.status === ConnectionStatus.connected || this.status === ConnectionStatus.fresh;
    }

    get _status() {
        const label = this._isConnected ? Strings.connected : Strings.not_connected;
        return `${this.system}: ${label}`;
    }

    get _icon() {
        return Icons[this._isConnected];
    }

    get _noticeClasses() {
        return NoticeClasses[this._isConnected];
    }

    get _assistiveText() {
        return AssistiveText[this._isConnected];
    }

    get _actionLabel() {
        const label = ActionLabels[this._isConnected];
        return `${label} ${this.system}`;
    }
}