import { LightningElement, api, wire } from 'lwc';
import { NavigationMixin, CurrentPageReference } from 'lightning/navigation';

export default class ViewAllSyncRecords extends NavigationMixin(LightningElement) {
	@wire(CurrentPageReference)
	pageRef;

	@api invoke() {
		const syncRecordAPIName = this.pageRef.attributes.objectApiName;

		this[NavigationMixin.Navigate]({
			type: 'standard__objectPage',
			attributes: {
					objectApiName: syncRecordAPIName,
					actionName: 'list'
			}
		});
	}
}