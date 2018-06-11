package stripe

import (
	"encoding/json"
	"testing"

	assert "github.com/stretchr/testify/require"
	"github.com/stripe/stripe-go/form"
)

func TestAccountExternalAccountParams_AppendTo(t *testing.T) {
	{
		params := &AccountExternalAccountParams{}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"bank_account"}, body.Get("object"))
	}

	{
		params := &AccountExternalAccountParams{Token: String("tok_123")}
		body := &form.Values{}

		// 0-length keyParts are not allowed, so call AppendTo directly (as
		// opposed to through the form package) and inject a realistic set
		params.AppendTo(body, []string{"external_account"})

		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"tok_123"}, body.Get("external_account"))
	}
}

func TestAccountUnmarshal(t *testing.T) {
	accountData := map[string]interface{}{
		"id": "acct_123",
		"external_accounts": map[string]interface{}{
			"object":   "list",
			"has_more": true,
			"data": []map[string]interface{}{
				{
					"object": "bank_account",
					"id":     "ba_123",
				},
				{
					"object": "card",
					"id":     "card_123",
				},
			},
		},
		"legal_entity": map[string]interface{}{
			"additional_owners": []map[string]interface{}{
				{
					"address": map[string]interface{}{
						"city":        "city1",
						"country":     "US",
						"line1":       "line1",
						"line2":       "line2",
						"postal_code": "90210",
						"state":       "CA",
					},
					"dob": map[string]interface{}{
						"day":   1,
						"month": 1,
						"year":  1970,
					},
					"first_name":                  "First",
					"last_name":                   "Last",
					"maiden_name":                 "Maiden",
					"personal_id_number_provided": true,
					"verification": map[string]interface{}{
						"details":       "details",
						"details_code":  "failed_other",
						"document":      "file_123",
						"document_back": "file_234",
						"status":        "pending",
					},
				},
				{
					"address": map[string]interface{}{
						"city": "city2",
					},
					"dob": map[string]interface{}{
						"day":   1,
						"month": 1,
						"year":  1960,
					},
					"first_name": "First2",
				},
			},
			"address": map[string]interface{}{
				"city":        "city-3",
				"country":     "US",
				"line1":       "line1-3",
				"line2":       "line2-3",
				"postal_code": "91111",
				"state":       "CA",
			},
			"business_tax_id_provided": true,
			"dob": map[string]interface{}{
				"day":   1,
				"month": 1,
				"year":  1980,
			},
			"first_name":  "John",
			"last_name":   "Doe",
			"maiden_name": "Maiden",
			"personal_address": map[string]interface{}{
				"city":        "personal_city1",
				"country":     "US",
				"line1":       "personal_line1",
				"line2":       "personal_line2",
				"postal_code": "90210",
				"state":       "CA",
			},
			"personal_id_number_provided": true,
			"ssn_last_4_provided":         true,
			"type":                        "company",
			"verification": map[string]interface{}{
				"details":       "details",
				"details_code":  "failed_other",
				"document":      "file_345",
				"document_back": "file_456",
				"status":        "unverified",
			},
		},
		"metadata": map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
		"payout_schedule": map[string]interface{}{
			"delay_days": 2,
			"interval":   "weekly",
		},
		"tos_acceptance": map[string]interface{}{
			"date":       1528573382,
			"ip":         "127.0.0.1",
			"user_agent": "user agent",
		},
		"type": "custom",
		"verification": map[string]interface{}{
			"disabled_reason": "fields_needed",
			"due_by":          1528573382,
			"fields_needed": []interface{}{
				"legal_entity.verification.document",
				"legal_entity.business_name",
			},
		},
	}

	bytes, err := json.Marshal(&accountData)
	assert.NoError(t, err)

	var account Account
	err = json.Unmarshal(bytes, &account)
	assert.NoError(t, err)

	assert.Equal(t, "acct_123", account.ID)

	assert.Equal(t, "value1", account.Metadata["key1"])
	assert.Equal(t, "value2", account.Metadata["key2"])

	assert.Equal(t, int64(2), account.PayoutSchedule.DelayDays)
	assert.Equal(t, PayoutIntervalWeekly, account.PayoutSchedule.Interval)

	assert.Equal(t, AccountTypeCustom, account.Type)

	assert.Equal(t, int64(1528573382), account.TOSAcceptance.Date)
	assert.Equal(t, "127.0.0.1", account.TOSAcceptance.IP)
	assert.Equal(t, "user agent", account.TOSAcceptance.UserAgent)

	assert.Equal(t, IdentityVerificationDisabledReasonFieldsNeeded, account.Verification.DisabledReason)
	assert.Equal(t, int64(1528573382), account.Verification.DueBy)
	assert.Equal(t, 2, len(account.Verification.FieldsNeeded))
	assert.Equal(t, "legal_entity.verification.document", account.Verification.FieldsNeeded[0])
	assert.Equal(t, "legal_entity.business_name", account.Verification.FieldsNeeded[1])

	// Assert ExternalAccounts are fully deserialized
	assert.Equal(t, true, account.ExternalAccounts.HasMore)

	assert.Equal(t, 2, len(account.ExternalAccounts.Data))
	assert.Equal(t, "ba_123", account.ExternalAccounts.Data[0].ID)
	assert.Equal(t, "card_123", account.ExternalAccounts.Data[1].ID)

	// Ensure LegalEntity is fully deserialized
	assert.NotNil(t, account.LegalEntity)

	assert.Equal(t, 2, len(account.LegalEntity.AdditionalOwners))
	assert.Equal(t, "city1", account.LegalEntity.AdditionalOwners[0].Address.City)
	assert.Equal(t, "US", account.LegalEntity.AdditionalOwners[0].Address.Country)
	assert.Equal(t, "line1", account.LegalEntity.AdditionalOwners[0].Address.Line1)
	assert.Equal(t, "line2", account.LegalEntity.AdditionalOwners[0].Address.Line2)
	assert.Equal(t, "90210", account.LegalEntity.AdditionalOwners[0].Address.PostalCode)
	assert.Equal(t, "CA", account.LegalEntity.AdditionalOwners[0].Address.State)
	assert.Equal(t, int64(1), account.LegalEntity.AdditionalOwners[0].DOB.Day)
	assert.Equal(t, int64(1), account.LegalEntity.AdditionalOwners[0].DOB.Month)
	assert.Equal(t, int64(1970), account.LegalEntity.AdditionalOwners[0].DOB.Year)
	assert.Equal(t, "First", account.LegalEntity.AdditionalOwners[0].FirstName)
	assert.Equal(t, "Last", account.LegalEntity.AdditionalOwners[0].LastName)
	assert.Equal(t, "Maiden", account.LegalEntity.AdditionalOwners[0].MaidenName)
	assert.True(t, account.LegalEntity.AdditionalOwners[0].PersonalIDNumberProvided)
	assert.Equal(t, "details", account.LegalEntity.AdditionalOwners[0].Verification.Details)
	assert.Equal(t, IdentityVerificationDetailsCodeFailedOther, account.LegalEntity.AdditionalOwners[0].Verification.DetailsCode)
	assert.Equal(t, "file_123", account.LegalEntity.AdditionalOwners[0].Verification.Document.ID)
	assert.Equal(t, "file_234", account.LegalEntity.AdditionalOwners[0].Verification.DocumentBack.ID)
	assert.Equal(t, IdentityVerificationStatusPending, account.LegalEntity.AdditionalOwners[0].Verification.Status)

	assert.Equal(t, "city2", account.LegalEntity.AdditionalOwners[1].Address.City)
	assert.Equal(t, "First2", account.LegalEntity.AdditionalOwners[1].FirstName)

	assert.Equal(t, "city1", account.LegalEntity.AdditionalOwners[0].Address.City)
	assert.Equal(t, "US", account.LegalEntity.AdditionalOwners[0].Address.Country)
	assert.Equal(t, "line1", account.LegalEntity.AdditionalOwners[0].Address.Line1)
	assert.Equal(t, "line2", account.LegalEntity.AdditionalOwners[0].Address.Line2)
	assert.Equal(t, "90210", account.LegalEntity.AdditionalOwners[0].Address.PostalCode)
	assert.Equal(t, "CA", account.LegalEntity.AdditionalOwners[0].Address.State)

	assert.Equal(t, "city-3", account.LegalEntity.Address.City)
	assert.Equal(t, "US", account.LegalEntity.Address.Country)
	assert.Equal(t, "line1-3", account.LegalEntity.Address.Line1)
	assert.Equal(t, "line2-3", account.LegalEntity.Address.Line2)
	assert.Equal(t, "91111", account.LegalEntity.Address.PostalCode)
	assert.Equal(t, "CA", account.LegalEntity.Address.State)

	assert.True(t, account.LegalEntity.BusinessTaxIDProvided)
	assert.Equal(t, int64(1), account.LegalEntity.DOB.Day)
	assert.Equal(t, int64(1), account.LegalEntity.DOB.Month)
	assert.Equal(t, int64(1980), account.LegalEntity.DOB.Year)
	assert.True(t, account.LegalEntity.PersonalIDNumberProvided)
	assert.Equal(t, "John", account.LegalEntity.FirstName)
	assert.Equal(t, "Doe", account.LegalEntity.LastName)
	assert.Equal(t, "Maiden", account.LegalEntity.MaidenName)

	assert.Equal(t, "personal_city1", account.LegalEntity.PersonalAddress.City)
	assert.Equal(t, "US", account.LegalEntity.PersonalAddress.Country)
	assert.Equal(t, "personal_line1", account.LegalEntity.PersonalAddress.Line1)
	assert.Equal(t, "personal_line2", account.LegalEntity.PersonalAddress.Line2)
	assert.Equal(t, "90210", account.LegalEntity.PersonalAddress.PostalCode)
	assert.Equal(t, "CA", account.LegalEntity.PersonalAddress.State)

	assert.True(t, account.LegalEntity.PersonalIDNumberProvided)
	assert.True(t, account.LegalEntity.SSNLast4Provided)
	assert.Equal(t, LegalEntityTypeCompany, account.LegalEntity.Type)

	assert.Equal(t, IdentityVerificationDetailsCodeFailedOther, account.LegalEntity.Verification.DetailsCode)
	assert.Equal(t, "file_345", account.LegalEntity.Verification.Document.ID)
	assert.Equal(t, "file_456", account.LegalEntity.Verification.DocumentBack.ID)
	assert.Equal(t, IdentityVerificationStatusUnverified, account.LegalEntity.Verification.Status)
}

func TestPayoutScheduleParams_AppendTo(t *testing.T) {
	{
		params := &PayoutScheduleParams{DelayDaysMinimum: Bool(true)}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"minimum"}, body.Get("delay_days"))
	}
}
