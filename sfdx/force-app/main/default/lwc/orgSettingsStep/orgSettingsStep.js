import { LightningElement } from 'lwc';

const SUB_PAGES = {
    PAGE_LAYOUTS: 'PageLayouts',
    LIGHTNING_RECORD_PAGES: 'LightningPages',
};

export default class OrgSettingsStep extends LightningElement {
    targetedObjects = [
        'Account',
        'Order',
        'PricebookEntry',
        'Product2',
    ];

    openPageLayouts(e) {
        this.gotoSetupObjectManagerPage(e.target.dataset.object, SUB_PAGES.PAGE_LAYOUTS);
    }

    openLightningRecordPages(e) {
        this.gotoSetupObjectManagerPage(e.target.dataset.object, SUB_PAGES.LIGHTNING_RECORD_PAGES);
    }

    /**
     * Returns the relative url to a page within Setup.
     *
     * @param {string} object
     * @param {string} subPage
     */
    gotoSetupObjectManagerPage(object, subPage) {
        window.open(`/lightning/setup/ObjectManager/${object}/${subPage}/view`, '_blank');
    }
}

