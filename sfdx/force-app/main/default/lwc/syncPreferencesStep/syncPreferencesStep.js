import saveSyncPreferences from '@salesforce/apex/setupAssistant.saveSyncPreferences';
import getSyncPreferences from '@salesforce/apex/setupAssistant.getSyncPreferences';
import getMulticurrencySelectionOptions from '@salesforce/apex/setupAssistant.getMulticurrencySelectionOptions';
import getFilterSettings from '@salesforce/apex/setupAssistant.getFilterSettings';
import saveFilterSettings from '@salesforce/apex/setupAssistant.saveFilterSettings';
import { LightningElement, api, track} from 'lwc';

export default class SyncPreferencesStep extends LightningElement {
    @api setupComplete = false; 
    @track stripeConnectRecord;
    @track stripeAccountId;
    @track lastSynced;
    @track defaultCurrency;
    @track syncRecordRetention;
    @track syncStartDate;
    @track apiPercentageLimit;
    @track cpqTermUnit;
    @track isMultiCurrencyEnabled;
    @track isPreferencesUpdated = false;
    @track totalApiCalls; 
    @track currencyOptions;
    @track isCpqInstalled;
    @track orderFilter;
    @track accountFilter;
    @track productFilter;
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
    accountFilterError = '';
    orderFilterError = '';
    productFilterError = '';
 
    @api async connectedCallback() {
        try {
            const syncPreferences = await getSyncPreferences();
            this.data = JSON.parse(syncPreferences);
            if(this.data.isSuccess && this.data.results.isConnected) {
                this.stripeAccountId = this.data.results.stripe_account_id;
                this.lastSynced = new Date(this.data.results.last_synced * 1000).toLocaleString(undefined, {year:'numeric', month:'numeric', day: '2-digit', hour: 'numeric', minute:'2-digit', timeZoneName:'short'});
                this.defaultCurrency = this.data.results.default_currency;
                this.syncRecordRetention = this.data.results.sync_record_retention;
                this.syncStartDate = new Date(this.data.results.sync_start_date * 1000).toISOString();
                this.apiPercentageLimit = this.data.results.api_percentage_limit;
                this.cpqTermUnit = this.data.results.cpq_term_unit;
                this.isCpqInstalled = this.data.results.isCpqInstalled;

                const multiCurrencyCheck = await getMulticurrencySelectionOptions();
                this.data = JSON.parse(multiCurrencyCheck);
                if(this.data.isSuccess) {
                    this.totalApiCalls = this.data.results.orgMaxApiLimit;
                    this.isMultiCurrencyEnabled = this.data.results.isMultiCurrencyEnabled;
                    if (this.isMultiCurrencyEnabled) {
                        this.currencyOptions = this.data.results.supportedISOCodes;
                    }
                } else if(this.data.error) { 
                    this.showToast(this.data.error, 'error', 'sticky');
                }

                const filterSettings = await getFilterSettings();
                const filterSettingsResponseData = JSON.parse(filterSettings);
                if(filterSettingsResponseData.isSuccess) {
                    this.orderFilter = filterSettingsResponseData.results.Order;
                    this.accountFilter = filterSettingsResponseData.results.Account;
                    this.productFilter = filterSettingsResponseData.results.Product2;
                } else if(filterSettingsResponseData.error) { 
                    this.showToast(filterSettingsResponseData.error, 'error', 'sticky');
                }
            } else if(this.data.error) { 
                this.showToast(this.data.error, 'error', 'sticky');
            }
        } catch (error) {
            this.showToast(error.message, 'error', 'sticky');
        } finally {
            this.loading = false;
        }
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

    valueChange(savedValue, updatedValue) {
        if(savedValue !== updatedValue) {
            this.dispatchEvent(new CustomEvent('valuechange'));
            return updatedValue;
        }
        return savedValue;
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

    get accountFilterClassList() {
        return (this.accountFilterError.length > 0? 'slds-has-error' : ''); 
    }

    get orderFilterClassList() {
        return (this.orderFilterError.length > 0 ? 'slds-has-error' : ''); 
    }

    get productFilterClassList() {
        return (this.productFilterError.length > 0 ? 'slds-has-error' : ''); 
    }

    @api async saveModifiedSyncPreferences() {
        let saveSuccess = false;
        try {
            if((this.apiPercentageLimit < 100 && this.apiPercentageLimit > 0) && (this.syncRecordRetention < 1000000 && this.syncRecordRetention > 100)) {
                const updatedSyncPreferences = await saveSyncPreferences({
                    cpqTermUnit: this.cpqTermUnit,
                    defaultCurrency: this.defaultCurrency,
                    syncRecordRetention: this.syncRecordRetention,
                    syncStartDate: (new Date(this.syncStartDate).getTime() / 1000),
                    apiPercentageLimit: this.apiPercentageLimit
                });
                const savedSyncPreferencesResponseData =  JSON.parse(updatedSyncPreferences);
                if(savedSyncPreferencesResponseData.isSuccess) {
                    // Reset object filter validation messages
                    this.accountFilterError = '';
                    this.productFilterError = '';
                    this.orderFilterError = '';
                    const updatedFilterSettings = await saveFilterSettings({
                        productFilter: this.productFilter,
                        orderFilter: this.orderFilter,
                        accountFiter: this.accountFilter
                    });
                    const filterResponseData =  JSON.parse(updatedFilterSettings);
                    if (filterResponseData.isSuccess ) {
                        if (filterResponseData.results.isFiltersSaved) {
                            saveSuccess = true
                            this.showToast('Changes were successfully saved', 'success');
                        } else {
                            if(filterResponseData.results.isValidationError) {
                                const listOfExceptions = filterResponseData.results.ValidationErrors;
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
                                        }  
                                    }
                                return;
                            }
                        }
                    } else {
                        this.showToast(filterResponseData.error, 'error', 'sticky');
                    } 
                    
                } else {
                    this.showToast(savedSyncPreferencesResponseData.error, 'error', 'sticky');
                } 
            }  

        } catch (error) {
            this.showToast(error.message, 'error', 'sticky');
        } finally {
            this.dispatchEvent(new CustomEvent('savecomplete', {
                detail: {
                    saveSuccess: saveSuccess
                }
            }));
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