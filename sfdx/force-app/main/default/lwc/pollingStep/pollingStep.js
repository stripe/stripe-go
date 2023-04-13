/**
 * Created by jmather-c on 3/28/23.
 */

import activatePolling from '@salesforce/apex/setupAssistant.activatePolling';
import getSyncPreferences from '@salesforce/apex/setupAssistant.getSyncPreferences';
import { LightningElement, api, track } from 'lwc';
import LightningConfirm from 'lightning/confirm';

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

    async connectedCallback() {
        await this.initPageStateData();
    }

    async initPageStateData() {
        const syncPrefs = await this.fetchSyncPreferences();

        if (syncPrefs === null) {
            return;
        }

        // console.log(syncPrefs);

        this.pollingEnabled = syncPrefs.polling_enabled;
        this.lastSynced = syncPrefs.last_synced;
        this.stripeAccountId = syncPrefs.stripe_account_id;
        this.configurationHash = syncPrefs.configurationHash;
        this.isConfigEnabled = syncPrefs.enabled;

        const syncStartDate = syncPrefs.sync_start_date;

        if (syncStartDate === "0") {
            this.syncStartDate = new Date().toISOString();
        } else if (syncStartDate !== null || syncStartDate !== "" || syncStartDate !== undefined) {
            this.syncStartDate = new Date(syncStartDate * 1000).toISOString();
        }
    }

    /**
     *
     * @returns {Promise<ResponseData|null>}
     */
    async fetchSyncPreferences() {
        const syncPreferences = await getSyncPreferences();
        /** @type {ResponsePayload} */
        const responseData = JSON.parse(syncPreferences);
        if(!responseData.results.isConnected) {
            return null;
        }
        if(responseData.error) {
            this.showToast(responseData.error, 'error', 'sticky');
            return null;
        }

        return responseData.results;
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
        const syncPrefs = await this.fetchSyncPreferences();

        if (syncPrefs === null) {
            return false;
        }

        const updatedSyncPreferences = await activatePolling({
            syncStartDate: new Date(this.syncStartDate).getTime() / 1000,
            isConfigEnabled: this.isConfigEnabled,
            configurationHash: this.configurationHash,
        });

        const savedSyncPreferencesResponseData = JSON.parse(updatedSyncPreferences);
        if (savedSyncPreferencesResponseData.isSuccess) {
            console.log('save result', JSON.parse(JSON.stringify(savedSyncPreferencesResponseData.results)));
            this.configurationHash = savedSyncPreferencesResponseData.results.configurationHash;
            this.showToast('Changes were successfully saved', 'success');
            this.dispatchEvent(new CustomEvent('savecomplete', {
                detail: {
                    saveSuccess: true,
                    configurationHash: this.configurationHash
                }
            }));

            return true;
        } else if (savedSyncPreferencesResponseData.error) {
            this.showToast(savedSyncPreferencesResponseData.error, 'error', 'sticky');
        }

        this.dispatchEvent(new CustomEvent('savecomplete', {
            detail: {
                saveSuccess: false,
                configurationHash: this.configurationHash
            }
        }));

        return false;
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
        // console.log('syncPrefs', this.syncPreferences);
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