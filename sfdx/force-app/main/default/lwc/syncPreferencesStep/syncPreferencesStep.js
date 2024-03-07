import { saveSyncPreferences, getSyncPreferences, getMulticurrencySelectionOptions, getFilterSettings, saveFilterSettings } from './configMapperConverter';
import manualTranslation from '@salesforce/apex/setupAssistant.manualTranslation';
import syncAllRecords from '@salesforce/apex/setupAssistant.syncAllRecords';
import { LightningElement, api, track} from 'lwc';
import { getErrorMessage } from 'c/utils'
import LightningConfirm from 'lightning/confirm';
import Debugger from "c/debugger";
import {ConfigData, ConfigManager, ConfigStates} from "c/systemConfigManager";

const POLLING_FIRST_ENABLE_ALERT_TITLE = 'Are you sure?';
const POLLING_FIRST_ENABLE_ALERT_TEMPLATE = `You are activating live syncing for your integration, meaning we will begin creating Stripe subscriptions for all Salesforce orders activated on or after SYNC_START_DATE. You can pause this in the future in the 'Sync Preferences' menu.`
const POLLING_DISABLE_ALERT_TITLE = 'Confirm Disable Syncing';
const POLLING_DISABLE_ALERT = 'Disabling syncing will halt all order processing. Are you sure you want to continue?';

export default class SyncPreferencesStep extends LightningElement {
    isDev = false;
    @api setupComplete = false;
    @track stripeConnectRecord;
    @track stripeAccountId;
    @track lastSynced;
    @track defaultCurrency;
    @track syncRecordRetention;
    @track syncStartDate;
    @track apiPercentageLimit;
    @track cpqTermUnit;
    @track cpqProratePrecision;
    @track isMultiCurrencyEnabled;
    @track isPreferencesUpdated = false;
    @track totalApiCalls; 
    @track currencyOptions;
    @track isCpqInstalled;
    @track orderFilter;
    @track accountFilter;
    @track productFilter;
    @track priceBookFilter;
    @track isSandbox;
    @track syncProductsDisabled = false; 
    @track syncPricebooksDisabled = false;
    @track pollingEnabled = false;
    @track pollingEnabledInitialValue = false;
    @track hiddenSyncPrefsFields = [];
    @track intervalOptions = [
        {
            label: 'Month',
            value: 'month'
        },
        {
            label: 'Day',
            value: 'day'
        }
    ];
    @track prorateOptions = [
        {
            label: 'Month',
            value: 'month'
        },
        {
            label: 'Monthly + Daily',
            value: 'month+day'
        },
    ];
    accountFilterError = '';
    orderFilterError = '';
    productFilterError = '';
    priceBookFilterError = '';

    // DEV NOTE: Manual Translation additions: refactor or remove these //
    manualTranslationValue = '';
    manualTranslationProcessing = false;

    get canDevReset() {
        return this.isDev && this.hasSynced;
    }

    devReset() {
        if (this.isDev === false) {
            return;
        }

        this.pollingEnabled = false;
        this.lastSynced = null;
    }

    updateManualTranslationValue(event) {
        this.manualTranslationValue = event.currentTarget.value;
    }

    async translateRecordId() {
        // prevent double-clicks
        if (this.manualTranslationProcessing) {
            return;
        }

        this.manualTranslationProcessing = true;

        let inputField = this.template.querySelector('[data-id="manualTranslationId"]');
        if (inputField === null) {
            this.manualTranslationProcessing = false;
            return;
        }

        // check that somehow we didn't get called with an invalid input
        if (inputField.checkValidity() === false) {
            inputField.reportValidity();
            this.manualTranslationProcessing = false;
            return;
        }

        const translateRecord = await manualTranslation({
            translationRecordID: this.manualTranslationValue
        });
        const responseData = JSON.parse(translateRecord);

        this.manualTranslationProcessing = false;

        if(responseData.error) {
            this.showToast(responseData.error, 'error', 'sticky');
            return;
        }
        if(!responseData.results.isRecordValid) {
            this.showToast(responseData.results.errorMessage, 'error', 'sticky');
            return;
        }

        this.showToast('Record "' + this.manualTranslationValue + '" successfully queued for translation', 'success');
        this.manualTranslationValue = '';
        inputField.focus();

        return;
    }

    /***
     *
     * @returns {boolean}
     */
    get disableManualTranslationId() {
        return this.manualTranslationProcessing;
    }

    /**
     *
     * @returns {boolean}
     */
    get disableManualTranslation() {
        // we check the local values to ensure this is re-ran whenever they change.
        if (this.manualTranslationValue.length === 0 || this.manualTranslationProcessing) {
            return true;
        }

        // we defer to the input to tell us if it's valid or not.
        let inputField = this.template.querySelector('[data-id="manualTranslationId"]');

        if (inputField === null) {
            return true;
        }

        return inputField.checkValidity() === false;
    }

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
        Debugger.log('SyncPreferencesStep', 'got config', config);
        if (config === null || config.isConnected === false) {
            Debugger.log('SyncPreferencesStep', 'Not yet connected...');
            return;
        }
        try {
            const syncPreferences = await getSyncPreferences();
            const responseData = JSON.parse(syncPreferences);
            if(!responseData.results.isConnected) {
                return;
            }
            if(responseData.error) {
                this.showToast(responseData.error, 'error', 'sticky');
                return;
            }
            this.stripeAccountId = responseData.results.stripe_account_id;
            if (responseData.results.last_synced) {
                this.lastSynced = new Date(responseData.results.last_synced * 1000).toLocaleString(undefined, {year:'numeric', month:'numeric', day: '2-digit', hour: 'numeric', minute:'2-digit', timeZoneName:'short'});
            }
            this.isM = responseData.results.default_currency;
            this.defaultCurrency = responseData.results.default_currency;
            this.syncRecordRetention = responseData.results.sync_record_retention;
            if (responseData.results.sync_start_date && responseData.results.sync_start_date !== "0") {
                this.syncStartDate = new Date(responseData.results.sync_start_date * 1000).toISOString();
            }
            this.apiPercentageLimit = responseData.results.api_percentage_limit || null;
            this.cpqTermUnit = responseData.results.cpq_term_unit || null;
            this.cpqProratePrecision = responseData.results.cpq_prorate_precision || 'month';
            this.isCpqInstalled = responseData.results.isCpqInstalled;
            this.isSandbox = responseData.results.isSandbox;
            this.pollingEnabled = responseData.results.polling_enabled;
            this.pollingEnabledInitialValue = responseData.results.polling_enabled;
            this.configurationHash = responseData.results.configurationHash;
            this.hiddenSyncPrefsFields = responseData.results.hiddenSyncPrefsFields;

            const multiCurrencyCheck = await getMulticurrencySelectionOptions();
            const multiCurrencyResponseData = JSON.parse(multiCurrencyCheck);
            if(multiCurrencyResponseData.error) {
                this.showToast(multiCurrencyResponseData.error, 'error', 'sticky');
                return;
            }
            this.totalApiCalls = multiCurrencyResponseData.results.orgMaxApiLimit;
            this.isMultiCurrencyEnabled = multiCurrencyResponseData.results.isMultiCurrencyEnabled;
            if (this.isMultiCurrencyEnabled) {
                this.currencyOptions = multiCurrencyResponseData.results.supportedISOCodes;
            }

            const filterSettings = await getFilterSettings();
            const filterSettingsResponseData = JSON.parse(filterSettings);

            if(!filterSettingsResponseData.isSuccess && filterSettingsResponseData.error) {
                this.showToast(filterSettingsResponseData.error, 'error', 'sticky');
            }

            this.orderFilter = filterSettingsResponseData.results.order_filter || '';
            this.accountFilter = filterSettingsResponseData.results.account_filter || '';
            this.productFilter = filterSettingsResponseData.results.product_filter || '';
            this.priceBookFilter = filterSettingsResponseData.results.pricebook_entry_filter || '';
        } catch (error) {
            let errorMessage = getErrorMessage(error);
            this.showToast(errorMessage, 'error', 'sticky');
        } finally {
            this.loading = false;
        }
    }

    get pollingEnabledChanged() {
        return this.pollingEnabled !== this.pollingEnabledInitialValue;
    }

    async updatePollingEnabled(event) {
        this.pollingEnabled = !this.pollingEnabled;
        this.dispatchEvent(new CustomEvent('valuechange'));
    }

    updateDefaultCurrency(event) {
        let updatedValue = this.valueChange(this.defaultCurrency, event.detail.value)
        this.defaultCurrency = updatedValue;
    }

    updateSyncRecordRetention(event) {
        let updatedValue = this.valueChange(this.syncRecordRetention, event.target.value)
        this.syncRecordRetention = updatedValue;
    }

    updateSyncStartDate(event) {
        let updatedValue = this.valueChange(this.syncStartDate, event.target.value)
        this.syncStartDate = updatedValue;
    }

    updateApiPercentageLimit(event) {
        let updatedValue = this.valueChange(this.apiPercentageLimit, event.target.value)
        this.apiPercentageLimit = updatedValue;
    }

    updateCpqTermInterval(event) {
        let updatedValue = this.valueChange(this.cpqTermUnit, event.detail.value)
        this.cpqTermUnit = updatedValue;
    }

    updateCpqProratePrecision(event) {
        let updatedValue = this.valueChange(this.cpqProratePrecision, event.detail.value)
        this.cpqProratePrecision = updatedValue;
    }

    updateAccountFilter(event) {
        let updatedValue = this.valueChange(this.accountFilter, event.detail.value)
        this.accountFilter = updatedValue;
    }

    updateProductFilter(event) {
        let updatedValue = this.valueChange(this.productFilter, event.detail.value)
        this.productFilter = updatedValue;
    }

    updateOrderFilter(event) {
        let updatedValue = this.valueChange(this.orderFilter, event.detail.value)
        this.orderFilter = updatedValue;
    }

    updatePriceBookFilter(event) {
        let updatedValue = this.valueChange(this.priceBookFilter, event.detail.value)
        this.priceBookFilter = updatedValue;
    }

    valueChange(savedValue, updatedValue) {
        if(savedValue !== updatedValue) {
            this.dispatchEvent(new CustomEvent('valuechange'));
            return updatedValue;
        }
        return savedValue;
    }

    showProductSyncModal() {
        this.template.querySelector('.stripe-modal_product-sync').show();
    }

    showPricebookSyncModal() {
        this.template.querySelector('.stripe-modal_pricebook-sync').show();
    }

    hideProductSyncModal() {
        this.template.querySelector('.stripe-modal_product-sync').hide();
    }

    hidePricebookSyncModal() {
        this.template.querySelector('.stripe-modal_pricebook-sync').hide();
    }

    syncProducts() {
        this.syncAllRecords('Product2');
        this.hideProductSyncModal();
    }

    syncPricebooks() {
        this.syncAllRecords('PricebookEntry');
        this.hidePricebookSyncModal();
    }


    get apiLimitCount() {
        if (this.apiPercentageLimit > 0) {
            return Math.floor(this.totalApiCalls * (this.apiPercentageLimit / 100));
        } else {
            return 0;
        }
    }

    get hasAccountFilterError() {
        return this.accountFilterError.length > 0;
    }
    
    get hasOrderFilterError() {
        return this.orderFilterError.length > 0;
    }
    
    get hasProductFilterError  () {
        return this.productFilterError.length > 0;
    }

    get hasPriceBookFilterError  () {
        return this.priceBookFilterError.length > 0;
    }

    get accountFilterClassList() {
        return (this.accountFilterError.length > 0? 'slds-has-error' : ''); 
    }

    get orderFilterClassList() {
        return (this.orderFilterError.length > 0 ? 'slds-has-error' : ''); 
    }

    get productFilterClassList() {
        return (this.productFilterError.length > 0 ? 'slds-has-error' : ''); 
    }

    get priceBookFilterClassList() {
        return (this.priceBookFilterError.length > 0 ? 'slds-has-error' : ''); 
    }

    @api async saveModifiedSyncPreferences() {
        Debugger.log('saveModifiedSyncPreferences', 'enter', this.pollingEnabledChanged, this.hasSynced, this.pollingEnabled);

        let saveSuccess = false;
        try {
            // confirm they wish to pause polling
            if (this.pollingEnabledChanged && this.hasSynced && this.pollingEnabled === false) {
                const result = await LightningConfirm.open({
                    message: POLLING_DISABLE_ALERT,
                    theme: 'texture',
                    label: POLLING_DISABLE_ALERT_TITLE, // this is the header text
                });

                if (result === false) {
                    Debugger.log('saveModifiedSyncPreferences', 'user cancelled disabling polling');
                    this.dispatchEvent(new CustomEvent('savecomplete', {
                        detail: {
                            saveSuccess: saveSuccess,
                            configurationHash: this.configurationHash
                        }
                    }));
                    return;
                }
            }

            // confirm they wish to initialize polling
            if (this.pollingEnabledChanged && this.hasSynced === false && this.pollingEnabled) {
                const startDateEle = this.template.querySelector('[data-id="syncStartDate"]');
                if (startDateEle.checkValidity() === false) {
                    Debugger.log('saveModifiedSyncPreferences', 'invalid sync start date');
                    startDateEle.reportValidity();
                    this.showToast("A Sync Start Date must be set.", 'error', 'sticky');
                    return;
                }
                const formattedDate = new Date(this.syncStartDate).toLocaleDateString();
                const alertMsg = POLLING_FIRST_ENABLE_ALERT_TEMPLATE.replace('SYNC_START_DATE', formattedDate);

                const result = await LightningConfirm.open({
                    message: alertMsg,
                    theme: 'texture',
                    label: POLLING_FIRST_ENABLE_ALERT_TITLE, // this is the header text
                });

                if (result === false) {
                    Debugger.log('saveModifiedSyncPreferences', 'user cancelled enabling polling');
                    this.dispatchEvent(new CustomEvent('savecomplete', {
                        detail: {
                            saveSuccess: saveSuccess,
                            configurationHash: this.configurationHash
                        }
                    }));
                    return;
                }

                // hack to make the sync start date input disable
                this.lastSynced = 'Scheduled';
            }

            const validApiPercentageLimit = this.apiPercentageLimit === null || (this.apiPercentageLimit > 0 && this.apiPercentageLimit < 100);
            const validSyncRecordRetention = this.syncRecordRetention < 1000000 && this.syncRecordRetention > 100;
            Debugger.log('saveModifiedSyncPreferences', 'before call check', {validApiPercentageLimit, validSyncRecordRetention});
            if (validApiPercentageLimit && validSyncRecordRetention) {
                const toPersist = {
                    pollingEnabled: this.pollingEnabled,
                    defaultCurrency: this.defaultCurrency,
                    syncRecordRetention: this.syncRecordRetention,
                    syncStartDate: (new Date(this.syncStartDate).getTime() / 1000),
                    apiPercentageLimit: this.apiPercentageLimit,
                    cpqTermUnit: this.cpqTermUnit,
                    cpqProratePrecision: this.cpqProratePrecision,
                    configurationHash: this.configurationHash,
                };
                const updatedSyncPreferences = await saveSyncPreferences(toPersist);
                const savedSyncPreferencesResponseData =  JSON.parse(updatedSyncPreferences);
                if(savedSyncPreferencesResponseData.isSuccess) {
                    Debugger.log('saveModifiedSyncPreferences', 'sync preferences saved')
                    this.configurationHash = savedSyncPreferencesResponseData.results.configurationHash;
                    this.pollingEnabledInitialValue = this.pollingEnabled;
                    // Reset object filter validation messages
                    this.accountFilterError = '';
                    this.productFilterError = '';
                    this.orderFilterError = '';
                    this.priceBookFilterError = '';
                    const updatedFilterSettings = await saveFilterSettings({
                        productFilter: this.productFilter,
                        orderFilter: this.orderFilter,
                        accountFilter: this.accountFilter,
                        priceBookFilter: this.priceBookFilter
                    });
                    const filterResponseData =  JSON.parse(updatedFilterSettings);
                    if (filterResponseData.isSuccess ) {
                        if (filterResponseData.results.isFiltersSaved) {
                            Debugger.log('saveModifiedSyncPreferences', 'filters saved');
                            this.configurationHash = filterResponseData.results.configurationHash;
                            saveSuccess = true
                            this.showToast('Changes were successfully saved', 'success');
                        } else {
                            if(filterResponseData.results.isValidationError) {
                                const listOfExceptions = filterResponseData.results.ValidationErrors;
                                Debugger.log('saveModifiedSyncPreferences', 'filters validation error', listOfExceptions);
                                for (let i = 0; i < listOfExceptions.length; i++) {
                                    let objectWithError = listOfExceptions[i].Object;
                                    let exceptionThrown = listOfExceptions[i].Error;
                                    if (objectWithError === 'Account') {
                                        // Set account validation message
                                        this.accountFilterError = exceptionThrown;
                                    } else if (objectWithError === 'Product2') {
                                        // Set product validation message
                                        this.productFilterError = exceptionThrown;
                                    } else if (objectWithError === 'Order') {
                                        // Set order validation message
                                        this.orderFilterError = exceptionThrown;
                                    }  else if (objectWithError === 'PricebookEntry') {
                                        // Set pricebook validation message
                                        this.priceBookFilterError = exceptionThrown;
                                    }
                                }
                            } else {
                                Debugger.log('saveModifiedSyncPreferences', 'filters save error', filterResponseData.error);
                                this.showToast(filterResponseData.error, 'error', 'sticky');
                            } 
                        }
                    } 
                } else {
                    Debugger.log('saveModifiedSyncPreferences', 'sync preferences save error', savedSyncPreferencesResponseData.error);
                    this.showToast(savedSyncPreferencesResponseData.error, 'error', 'sticky');
                }
            }
        } catch (error) {
            Debugger.log('saveModifiedSyncPreferences', 'error', error);
            let errorMessage = getErrorMessage(error);
            this.showToast(errorMessage, 'error', 'sticky');
        } finally {
            Debugger.log('saveModifiedSyncPreferences', 'save complete');
            this.dispatchEvent(new CustomEvent('savecomplete', {
                detail: {
                    saveSuccess: saveSuccess,
                    configurationHash: this.configurationHash
                }
            }));
        }
    }

    @api updateConfigHash(configHash) {
        this.configurationHash = configHash;
    }

    get hasSynced() {
        return this.lastSynced !== null && this.lastSynced !== undefined;
    }

    async syncAllRecords(objectType) {
        try {
            const syncAllRecordsResponse = await syncAllRecords({
                objectType: objectType
            });
            const responseData =  JSON.parse(syncAllRecordsResponse);
            if (responseData.error) {
                this.showToast(responseData.error, 'error', 'sticky');
                return
            }

            if (!responseData.results.syncAllRecordsDispatched) {
                this.showToast('There was an error syncing records, please try again.', 'error', 'sticky');
                return
            } 

            this.showToast('All ' + objectType + ' records have been dispatched to be synced.', 'success');
            if (objectType === 'PricebookEntry') {
                //disable Pricebook Entry button 
                this.syncPricebooksDisabled = true;
                return
            }
            this.syncProductsDisabled = true;
        
        } catch (error) {
            let errorMessage = getErrorMessage(error);
            this.showToast(errorMessage, 'error', 'sticky');
        } 
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