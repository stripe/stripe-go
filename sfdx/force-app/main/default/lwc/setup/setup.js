import { LightningElement, track } from 'lwc';
import { getErrorMessage } from 'c/utils'
import { Manager, ServiceEvents } from "c/serviceManager";
import Debugger from 'c/debugger';
const DebugLog = Debugger.withContext('SetupWizard');

import companyLogo from '@salesforce/resourceUrl/companyLogo';
import checkUserPermissions from '@salesforce/apex/setupAssistant.checkUserPermissions';
import getSetupData from '@salesforce/apex/setupAssistant.getSetupData';
import getPackageVersion from '@salesforce/apex/utilities.getPackageVersion';
import setOrgType from '@salesforce/apex/utilities.setOrgType';
import saveData from '@salesforce/apex/setupAssistant.saveData';
import illustrations from '@salesforce/resourceUrl/illustrations';
import handleGeneratePackageKeyForV2 from '@salesforce/apex/setupAssistant.handleGeneratePackageKeyForV2';

export default class FirstTimeSetup extends LightningElement {
    systemConnectionsIllustration = illustrations + '/stripe_illustration_systemConnections.svg';
    dataMappingIllustration = illustrations + '/stripe_illustration_dataMapping.svg';
    syncPreferencesIllustration = illustrations + '/stripe_illustration_syncPreferences.svg';
    integrationUserIllustration = illustrations + '/stripe_setup_authorizeWebhooks.svg';

    formattedPackageVersion = '';
    @track contentShown = false;
    @track logoUrl = companyLogo;
    @track activeStepIndex = 0;
    @track activeSectionIndex = undefined;
    @track showIntro = false;
    @track setupComplete = false;
    @track setupStarted = false;
    @track nextDisabled = true;
    @track saveDisabled = true;
    @track loading = true;
    @track contentLoading = false;
    @track changesCanceled = false;
    @track currentTab = '';
    @track tabToNavTo = '';
    @track stepName;
    @track steps = [
        {
            title: 'Connect Your Stripe Account',
            name: 'C-STRIPE-ACCOUNT-SETUP-STEP',
            orderIndex: 1,
            isComplete: false,
            isActive: false,
            component: 'c-stripe-account-setup-step',
        },
        {
            title: 'Connect Your Salesforce Account',
            name: 'C-INTEGRATION-USER-SETUP-STEP',
            orderIndex: 2,
            isComplete: false,
            isActive: false,
            component: 'c-integration-user-setup-step',
        },
        {
            title: 'Define Data Mapping',
            name: 'C-DATA-MAPPING-STEP',
            orderIndex: 3,
            isComplete: false,
            isActive: false,
            component: 'c-data-mapping-step',
        },
        {
            title: 'Manage Integration',
            name: 'C-SYNC-PREFERENCES-STEP',
            orderIndex: 4,
            isComplete: false,
            isActive: false,
            component: 'c-sync-preferences-step',
        },
        {
            title: 'Activate Syncing',
            name: 'C-POLLING-STEP',
            orderIndex: 5,
            isComplete: false,
            isActive: false,
            component: 'c-polling-step',
        },
        {
            title: 'Update Salesforce Pages',
            name: 'C-ORG-SETTINGS-STEP',
            orderIndex: 6,
            isComplete: false,
            isActive: false,
            component: 'c-org-settings-step',
        },
    ];

    setupStepRefs = {
        stripeAccount: 0,
        integrationUser: 1,
        dataMapping: 2,
        syncPreferences: 3,
        polling: 4,
        orgSettings: 5,
    };

    generalStepRefs = {
        stripeAccount: 0,
        integrationUser: 1,
        dataMapping: 2,
        syncPreferences: 3,
    };

    @track missingPermissions = {};
    @track hasMissingPermissions = false;

    @track _showSetupLanding = true;
    get showSetupLanding() {
        return this.setupComplete ? false : this._showSetupLanding;
    }

    set showSetupLanding(value) {
        this._showSetupLanding = value;
    }

    get showSettingsLanding() {
        return this.activeSectionIndex == undefined;
    }

    get backDisabled() {
        return this.activeStepIndex === 0;
    }

    get isFinalStep() {
        return this.activeStepIndex === (this.steps.length - 1);
    }

    get nextLabel() {
        return this.isFinalStep ? 'Finish' : 'Next';
    }

    get sidebarClasslist() {
        return 'slds-col slds-size_1-of-4' + (this.setupComplete ? ' stripe-sidebar_navigation' : ' stripe-sidebar_progress');
    }

    get stepContentClasslist() {
        return this.showSetupLanding ? 'slds-hide' : '';
    }

    get settingsContentClasslist() {
        return this.showSettingsLanding ? 'slds-hide' : '';
    }

    get landingButtonLabel() {
        return this.setupStarted ? 'Continue' : 'Get Started';
    }

    enableNext() {
        this.nextDisabled = false;
    }

    enableSave() {
        this.saveDisabled = false;
    }

    enableContentLoading() {
        this.contentLoading = true;
    }

    disableContentLoading() {
        this.contentLoading = false;
    }

    completeSave(event) {
        if (event.detail.saveSuccess && event.detail.configurationHash) {
            this.template.querySelector('c-data-mapping-step').updateConfigHash(event.detail.configurationHash);
            this.template.querySelector('c-sync-preferences-step').updateConfigHash(event.detail.configurationHash);

            // only present during setup.
            const maybeStep = this.template.querySelector('c-polling-step');
            if (maybeStep) {
                maybeStep.updateConfigHash(event.detail.configurationHash);
            }
        }

        this.contentLoading = false;
        if(event.detail.saveSuccess && this.setupComplete === true) {
            this.saveDisabled = true;
        } 
        if (event.detail.saveSuccess && this.setupComplete === false) {
            this.nextNavigate();
        }
    }

    completeStep(event) {
        DebugLog('stepComplete', event.detail.step);
        DebugLog('activeStep', this.activeStepIndex);
        this.steps[this.activeStepIndex].isComplete = true;
        if (event.detail.step === 'integration_user_connection' && this.activeStepIndex === this.setupStepRefs.integrationUser) {
            this.nextDisabled = false;
        }
        if (event.detail.step === 'stripe_account_connection' && this.activeStepIndex === this.setupStepRefs.stripeAccount) {
            this.nextDisabled = false;
        }
    }

    async systemsConnected(event) {
        DebugLog('systemsConnected', event.detail);
        if (event.detail.isConnected && this.setupComplete === false && event.detail.isFirstRun === false) {
            DebugLog('Refreshing data dependent items.');
            await handleGeneratePackageKeyForV2({});
            //this.template.querySelector('c-data-mapping-step').getPicklistValuesForMapper(true, '');
            this.template.querySelector('c-sync-preferences-step').connectedCallback();
            this.template.querySelector('c-polling-step').connectedCallback();
            this.nextDisabled = false;
        }
    }

    disconnectedCallback() {
        if (this.boundSystemsConnected) {
            Manager.off(ServiceEvents.core_functionality_established, this.boundSystemsConnected);
        }
    }

    async connectedCallback() {
        if (this.boundSystemsConnected === undefined) {
            this.boundSystemsConnected = this.systemsConnected.bind(this);
            Manager.on(ServiceEvents.core_functionality_established, this.boundSystemsConnected);
        }
        try {
            const userPermissionCheck = await checkUserPermissions();
            const userPermissionResponseData = JSON.parse(userPermissionCheck);
            if (userPermissionResponseData.error) {
                this.showSetupToast(userPermissionResponseData.error, 'error', 'sticky');
                return;
            }

            const permissionIssueMap = userPermissionResponseData.results.permissionIssueMap;

            if (!permissionIssueMap.isPermSetAssigned || permissionIssueMap.isSystemPermissionMissing || permissionIssueMap.isObjectPermissionMissing) {
                this.missingPermissions = permissionIssueMap;
                this.hasMissingPermissions = true;
                return;
            }

            // TODO this should return an object that can be assigned rather than mutating stuff directly
            await this.fetchSetupData();

            const setOrganizationType = await setOrgType();
            const orgTypeResponseData =  JSON.parse(setOrganizationType);

            if (orgTypeResponseData.error) {
                this.showSetupToast(orgTypeResponseData.error, 'error', 'sticky');
                return;
            }

            const rawPackageVersion = await getPackageVersion();
            const packageVersion = JSON.parse(rawPackageVersion).results;
            this.formattedPackageVersion = packageVersion.major + '.' + packageVersion.minor;
        } catch (error) {
            let errorMessage = getErrorMessage(error);
            this.showSetupToast(errorMessage, 'error', 'sticky');
        } finally {
            this.loading = false;

            // skip the initial landing screen by grabbing the first nav item and "clicking"
            if (this.setupComplete && this.activeSectionIndex === undefined) {
                const comp = this;
                setTimeout(() => {
                    const orgSettings = comp.template.querySelector('li[data-section-name="orgSettings"]');
                    orgSettings.click();
                }, 0);
            }
        }
    }

    async fetchSetupData() {
        DebugLog('fetchSetupData', 'Fetching setup data');
        this.steps[this.activeStepIndex].isActive = true;
        try {
            const setupData = await getSetupData();
            const responseData = JSON.parse(setupData);
            DebugLog('fetchSetupData', 'Setup data response', responseData);

            if (responseData.error) {
                DebugLog('fetchSetupData', 'Error fetching setup data', responseData.error);
                this.showSetupToast(responseData.error, 'error', 'sticky');
                return;
            }

            this.setupComplete = responseData.results.setupData.isSetupComplete__c;
            if (this.setupComplete) {
                return;
            }

            let completedSteps = JSON.parse(responseData.results.setupData.Steps_Completed__c);
            DebugLog('fetchSetupData', 'Completed steps', completedSteps);

            if (!responseData.results.isConnected && (completedSteps != null && Object.keys(completedSteps).length <= 0)) {
                DebugLog('fetchSetupData', 'Not connected, no completed steps');
                // also in what scenario is length less than 0?
                return;
            }

            this.nextDisabled = false;
            this.setupStarted = true;
            this.steps[this.activeStepIndex].isActive = false;

            // this seems bad... but so much of this is bad... so... whatever? lol. It works.
            for (const step in completedSteps) {
                DebugLog('fetchSetupData', 'Processing step', step);
                this.steps[this.activeStepIndex].isComplete = true;

                if (this.activeStepIndex < this.steps.length - 1) {
                    this.activeStepIndex++;
                }
            }
            this.steps[this.activeStepIndex].isActive = true;

             // This is causing a known bug where if a user auths Salesforce but doesn't
            // press "Next", the next time the user loads the app, it'll show "Reauthorize" but Next is not enabled. 
            // This is a known bug and will be fixed in a future release since commenting this out will cause other downstream affects.
            if (this.activeStepIndex < 2) {
                this.nextDisabled = true;
            }

            DebugLog('fetchSetupData', 'Final state', { activeStepIndex: this.activeStepIndex, steps: this.steps });
        } catch (error) {
            let errorMessage = getErrorMessage(error);
            DebugLog('fetchSetupData', 'Error fetching setup data', { errorMessage, error });
            this.showSetupToast(errorMessage, 'error', 'sticky');
        }
    }

    async save() {
        this.contentLoading = true;
        if(this.activeSectionIndex == this.generalStepRefs.dataMapping) {
            return this.template.querySelector('c-data-mapping-step').saveDataMappings();
        } else if(this.activeSectionIndex == this.generalStepRefs.syncPreferences) {
            return this.template.querySelector('c-sync-preferences-step').saveModifiedSyncPreferences();
        }
    }

    next(e) {
        DebugLog('got next', e);
        this.contentLoading = true;
        this.stepName = this.steps[this.activeStepIndex].name;
        DebugLog('processing step', this.stepName);
        if(this.stepName === 'C-DATA-MAPPING-STEP') {
            return this.template.querySelector('c-data-mapping-step').saveDataMappings();
        } else if(this.stepName === 'C-SYNC-PREFERENCES-STEP') {
            return this.template.querySelector('c-sync-preferences-step').saveModifiedSyncPreferences();
        } else if(this.stepName === 'C-POLLING-STEP') {
            if (e.currentTarget.dataset && e.currentTarget.dataset.activate === "now") {
                return this.template.querySelector('c-polling-step').activatePolling();
            }            
        }
        return this.nextNavigate();
    }

    async nextNavigate() {
        try {
            if(this.activeStepIndex <= (this.steps.length - 1)) {
                let stepNumber = this.activeStepIndex + 1;
                let isLastStep = false;
                if (stepNumber === (this.steps.length)) {
                    isLastStep = true
                    this.steps[this.activeStepIndex].isActive = false;
                    this.steps[this.activeStepIndex].isComplete = true;
                    this.nextDisabled = true;
                    this.setupComplete = true;
                    this.showSetupToast('Setup completed', 'success');
                }

                const saveSetupData = await saveData({
                    newSetupDataRec: {
                        Steps_Completed__c: JSON.stringify({
                            [this.stepName]: stepNumber
                        })
                    },
                    isSetupComplete: isLastStep
                });
                this.data = JSON.parse(saveSetupData);
                if(this.data.isSuccess) {
                    this.setupData = this.data.setupData;
                } else {
                    this.showSetupToast(this.data.error, 'error', 'sticky');
                }
            }
        } catch (error) {
            let errorMessage = getErrorMessage(error);
            this.showSetupToast(errorMessage, 'error', 'sticky');
        } finally {
            if(this.activeStepIndex < (this.steps.length - 1)) {
                this.showNextStep();
            }
            this.contentLoading = false;
        }
    }

    back() {
        this.contentLoading = true;
        if(this.activeStepIndex > 0) {
            const lastActiveStep = this.template.querySelector('c-step[data-index="' + this.activeStepIndex + '"]');
            this.activeStepIndex--;
            const newActiveStep = this.template.querySelector('c-step[data-index="' + this.activeStepIndex + '"]');
            if(lastActiveStep && newActiveStep) {
                this.steps[lastActiveStep.dataset.index].isActive = false;
                // this.steps[lastActiveStep.dataset.index].isComplete = false;
                this.steps[newActiveStep.dataset.index].isActive = true;
                // this.steps[newActiveStep.dataset.index].isComplete = false;
                if (newActiveStep.dataset.index < 2) {
                    this.nextDisabled = false;
                }
                lastActiveStep.classList.add('slds-hide');
                newActiveStep.classList.remove('slds-hide');
            }
        }
        this.contentLoading = false;
    }

    sectionNavigate(event) {
        if(this.saveDisabled) {
            let targetSection = event.currentTarget;
            if(targetSection.dataset.index !== this.activeSectionIndex) {
                const currentActive = this.template.querySelector('.stripe-navigation-item_active');
                if(currentActive) {
                    // Clear previous active section
                    currentActive.classList.remove('stripe-navigation-item_active');
                    const previousSectionContent = this.template.querySelector('c-step[data-index="' + this.activeSectionIndex + '"]');
                    if(previousSectionContent) {
                        previousSectionContent.classList.add('slds-hide');
                    }
                }
                // Make selected section active
                targetSection.classList.add('stripe-navigation-item_active');
                this.activeSectionIndex = targetSection.dataset.index;
                const targetSectionContent = this.template.querySelector('c-step[data-index="' + this.activeSectionIndex + '"]');
                if(targetSectionContent) {
                    targetSectionContent.classList.remove('slds-hide');
                }
            }
        } else {
            this.currentTab = event.currentTarget;
            this.tabToNavTo = this.currentTab.dataset.index;
            this.showCancelModal();
        }
    }

    runTabConnectedCallback() {
        this.saveDisabled = true;
        this.changesCanceled = true;
        if(this.activeSectionIndex == this.generalStepRefs.syncPreferences) {
            this.template.querySelector('c-sync-preferences-step').connectedCallback();
            this.hideCancelModal();
            return;
        }

        this.template.querySelector('c-data-mapping-step').connectedCallback();
        this.hideCancelModal();
    }

    showCancelModal() {
        this.template.querySelector('.stripe-modal_confirm-cancel').show();
    }

    hideCancelModal() {
        this.template.querySelector('.stripe-modal_confirm-cancel').hide();
        if(this.changesCanceled) {
            this.showSetupToast('Your changes have been reverted', 'success');
            this.changesCanceled = false;

            //navigate to selcted step
            const currentActive = this.template.querySelector('.stripe-navigation-item_active');
            const previousSectionContent = this.template.querySelector('c-step[data-index="' + this.activeSectionIndex + '"]');
            if(previousSectionContent) {
                currentActive.classList.remove('stripe-navigation-item_active');
                previousSectionContent.classList.add('slds-hide');
            }

            this.currentTab.classList.add('stripe-navigation-item_active');
             this.activeSectionIndex = this.tabToNavTo;
            const targetSectionContent = this.template.querySelector('c-step[data-index="' + this.activeSectionIndex + '"]');
            if(targetSectionContent) {
                targetSectionContent.classList.remove('slds-hide');
            }
        }

    }

    showNextStep() {
        const lastActiveStep = this.template.querySelector('c-step[data-index="' + this.activeStepIndex + '"]');
        this.activeStepIndex++;
        const newActiveStep = this.template.querySelector('c-step[data-index="' + this.activeStepIndex + '"]');
        if (lastActiveStep && newActiveStep) {
            DebugLog('showNextStep', 'moving to next step');
            const lastActiveStepIndex = parseInt(lastActiveStep.dataset.index, 10);
            const newActiveStepIndex = parseInt(newActiveStep.dataset.index, 10);
            const lastStepObj = this.steps[lastActiveStepIndex];
            const nextStepObj = this.steps[newActiveStepIndex];
            DebugLog('showNextStep', 'last', lastStepObj);
            DebugLog('showNextStep', 'next', nextStepObj);
            lastStepObj.isActive = false;
            lastStepObj.isComplete = true;
            nextStepObj.isActive = true;

            if (newActiveStepIndex < 2) {
                this.nextDisabled = nextStepObj.isComplete === false;
            }
            
            lastActiveStep.classList.add('slds-hide');
            newActiveStep.classList.remove('slds-hide');
        }
    }

    continueSetup() {
        const activeStep = this.template.querySelector('c-step[data-index="' + this.activeStepIndex + '"]');
        if(activeStep) {
            activeStep.classList.remove('slds-hide');
            this.contentShown = true;
            this.showSetupLanding = false;
        }
    }

    showSetupToast(message, variant, mode) {
        this.template.querySelector('c-toast').show(message, variant, mode);
    }

    showToast(event) {
        event.stopPropagation();
        let toast = event.detail.toast;
        this.template.querySelector('c-toast').show(toast.message, toast.variant, toast.mode);
    }
}