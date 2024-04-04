import {ConfigManager, ConfigStates, ConfigData} from 'c/systemConfigManager';
import { LightningElement, api } from 'lwc';
import LightningConfirm from 'lightning/confirm';
import Debugger from "c/debugger";

/** @typedef {Object} ResponseData
 * @property {boolean} polling_enabled
 * @property {number|null} last_synced
 * @property {string} stripe_account_id
 * @property {string|number} sync_start_date
 * @property {boolean} isConnected
 * @property {string} configurationHash
 * @property {string} default_currency
 * @property {string} sync_record_retention
 * @property {string} api_percentage_limit
 * @property {string} cpq_term_unit
 * @property {string} isCpqInstalled
 * @property {string} isSandbox
 * @property {string} enabled
 */

/** @typedef {Object} ResponsePayload
 * @property {ResponseData} results
 * @property {string} error
 */

/**
 * @typedef {Object} PageContent
 * @property {string} badgeLabel
 * @property {string} badgeStyle
 * @property {string} badgeIcon
 * @property {string} actionLabel
 * @property {string} actionBrand
 * @property {string} actionIcon
 */

/** @enum {string} */
const POLLING_STATUS = {
    DISABLED: 'disabled',
    INACTIVE: 'inactive',
    ACTIVE: 'active',
}

/** @type {{POLLING_STATUS:PageContent}} */
const contentByStatus = {
    [POLLING_STATUS.DISABLED]: {
        badgeLabel: 'Disabled',
        badgeStyle: 'warning',
        badgeIcon: 'utility:stop',
        actionLabel: 'Activate Polling',
        actionBrand: 'neutral',
        actionIcon: 'utility:sync',
    },
    [POLLING_STATUS.INACTIVE]: {
        badgeLabel: 'Inactive',
        badgeStyle: 'warning',
        badgeIcon: 'utility:pause_alt',
        actionLabel: 'Unpause Polling',
        actionBrand: 'neutral',
        actionIcon: '',
    },
    [POLLING_STATUS.ACTIVE]: {
        badgeLabel: 'Active',
        badgeStyle: '',
        badgeIcon: 'utility:play',
        actionLabel: 'Pause Polling',
        actionBrand: 'neutral',
        actionIcon: '',
    }
}

const ALERT_TEMPLATE = `You are activating live syncing for your integration, meaning we will begin creating Stripe subscriptions for all Salesforce orders activated on or after SYNC_START_DATE. You can pause this in the future in the 'Sync Preferences' menu.`

export default class PollingStep extends LightningElement {
    isDev = false;
    pollingEnabled = false;
    lastSynced = null;
    stripeAccountId = null;
    syncStartDate = null;
    configurationHash = null;
    isConfigEnabled = false;

    @api showButton = false;
    @api showTitle = false;

    _firstTranslationCallback = true;
    /**
     *
     * @param {SystemConfig<TranslationData>} event
     */
    _processTranslationConfigUpdate(event) {
        const values= event.detail;
        Debugger.log('SyncPreferencesStep', '_processTranslationConfigUpdate', values);

        // we will have literally just done the connectedCallback...
        if (this._firstTranslationCallback) {
            this._firstTranslationCallback = false;
            Debugger.log('SyncPreferencesStep', '_processTranslationConfigUpdate', 'skipping first run');
            return;
        }
        this.connectedCallback();
        Debugger.log('SyncPreferencesStep', '_processTranslationConfigUpdate', 'complete');
    }

    _initConfigManager() {
        Debugger.log('SyncPreferencesStep', '_initConfigManager');
        if (this._boundProcessTranslationConfigUpdate) {
            Debugger.log('SyncPreferencesStep', '_initConfigManager', 'already ran');
            return;
        }
        this._boundProcessTranslationConfigUpdate = this._processTranslationConfigUpdate.bind(this);
        ConfigManager.register(ConfigData.translation, ConfigStates.updated, this._boundProcessTranslationConfigUpdate);
        Debugger.log('SyncPreferencesStep', '_initConfigManager', 'registered')
    }

    _deinitConfigManager() {
        ConfigManager.unregister(ConfigData.translation, ConfigStates.updated, this._boundProcessTranslationConfigUpdate);
    }

    disconnectedCallback() {
        this._deinitConfigManager();
    }

    @api async connectedCallback() {
        this._initConfigManager();
        const config = await ConfigManager.getCachedTranslationData();
        if (config === null || config.isConnected === false) {
            Debugger.log('PollingStep', 'Not yet connected...');
            return;
        }

        await this.initPageStateData();
    }

    async initPageStateData() {
        const translationData = await ConfigManager.getCachedTranslationData();

        if (translationData === null) {
            return;
        }

        Debugger.log('pollingStep', 'initPageStateData', translationData);

        this.pollingEnabled = translationData.polling_enabled;
        this.lastSynced = translationData.last_synced;
        this.stripeAccountId = translationData.stripe_account_id;
        this.configurationHash = translationData.configurationHash;
        this.isConfigEnabled = translationData.enabled;

        const syncStartDate = translationData.sync_start_date;
        if (syncStartDate && syncStartDate !== "0") {
            this.syncStartDate = new Date(parseInt(syncStartDate, 10) * 1000).toISOString();
        } else {
            this.syncStartDate = new Date().toISOString();
        }
    }

    /**
     *
     * @param {{target:{value:string}}} event
     */
    updateSyncStartDate(event) {
        this.syncStartDate = event.target.value;
    }

    /**
     *
     * @returns {PageContent}
     */
    get content() {
        return contentByStatus[this.pollingStatus];
    }

    get badgeClass() {
        const style = this.content.badgeStyle;
        return (style !== '' && style !== null) ? `slds-theme_${style}` : '';
    }
    get badgeLabel() { return 'Status: ' + this.content.badgeLabel }
    get badgeIcon() { return this.content.badgeIcon }
    get actionLabel() { return this.content.actionLabel }
    get actionBrand() { return this.content.actionBrand }
    get actionIcon() { return this.content.actionIcon }

    /**
     *
     * @param {Event} event
     */
    async changePollingStatus(event) {
        switch(this.pollingStatus) {
            case POLLING_STATUS.DISABLED:
                await this.activatePolling();
                return;
            case POLLING_STATUS.INACTIVE:
                this.pollingStatus = POLLING_STATUS.ACTIVE;
                return;
            case POLLING_STATUS.ACTIVE:
                this.pollingStatus = POLLING_STATUS.INACTIVE;
        }
    }

    /**
     *
     * @returns {Promise<boolean>}
     */
    async save() {
        const translationData = await ConfigManager.getCachedTranslationData();

        // basic sanity check to ensure we're in a good state...
        if (translationData === null) {
            return false;
        }

        const payload = {
            polling_enabled: this.pollingEnabled,
            sync_start_date: new Date(this.syncStartDate).getTime() / 1000,
        }

        let saveSuccess = false;
        try {
            const result = await ConfigManager.saveTranslationConfig(payload);
            this.configurationHash = result.configurationHash;
            this.showToast('Changes were successfully saved', 'success');
            saveSuccess = true;
        } catch (err) {
            this.showToast(err.message, 'error', 'sticky');
        }

        this.dispatchEvent(new CustomEvent('savecomplete', {
            detail: {
                saveSuccess,
                configurationHash: this.configurationHash
            }
        }));

        return saveSuccess;
    }

    @api async activatePolling() {
        // const inputField = this.template.querySelector('[data-id="sync-start-date"]');
        // if (inputField.checkValidity() === false) {
        //     inputField.reportValidity();
        //     return;
        // }

        const formattedDate = new Date(this.syncStartDate).toLocaleDateString();
        const alertMsg = ALERT_TEMPLATE.replace('SYNC_START_DATE', formattedDate);

        const result = await LightningConfirm.open({
            message: alertMsg,
            theme: 'texture',
            label: 'Are you sure?', // this is the header text
        });

        if (result) {
            const isSaved = await this.save();
            if (isSaved) {
                this.pollingStatus = POLLING_STATUS.ACTIVE;
            }
        } else {
            this.dispatchEvent(new CustomEvent('savecomplete', {
                detail: {
                    saveSuccess: false,
                    configurationHash: this.configurationHash
                }
            }));
        }
    }

    get canDevReset() {
        return this.isDev && this.hasSynced;
    }

    devReset() {
        if (this.isDev === false) {
            return;
        }

        this.pollingStatus = POLLING_STATUS.DISABLED;
    }

    /**
     *
     * @param {POLLING_STATUS} value
     */
    set pollingStatus(value) {
        switch (value) {
            case POLLING_STATUS.DISABLED:
                this.pollingEnabled = false;
                this.lastSynced = null;
                return;
            case POLLING_STATUS.INACTIVE:
                this.pollingEnabled = false;
                this.lastSynced = 1;
                return;
            case POLLING_STATUS.ACTIVE:
                this.pollingEnabled = true;
                this.lastSynced = 1;
                return;
        }
    }

    /**
     *
     * @returns {POLLING_STATUS}
     */
    get pollingStatus() {
        Debugger.log('syncPrefs', this.syncPreferences);
        if (this.hasSynced === false) {
            return POLLING_STATUS.DISABLED;
        }

        return (this.isPolling) ? POLLING_STATUS.ACTIVE : POLLING_STATUS.INACTIVE;
    }


    /**
     *
     * @returns {boolean}
     */
    get isPolling() {
        return this.pollingEnabled || false;
    }

    /**
     *
     * @returns {boolean}
     */
    get hasSynced() {
        return this.lastSynced !== null && this.lastSynced !== undefined;
    }

    @api updateConfigHash(configHash) {
        this.configurationHash = configHash;
    }

    showToast(message, variant, mode) {
        this.dispatchEvent(new CustomEvent('showtoast', {
            bubbles: true,
            composed: true,
            detail: {
                toast: {
                    message: message,
                    variant: variant ? variant : 'info',
                    mode: mode ? mode : 'dismissible'
                }
            }
        }));
    }
}