package stripe

import (
	"encoding/json"
	"testing"

	"github.com/max-cape/stripe-go-test/form"
	assert "github.com/stretchr/testify/require"
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

func TestAccount_Unmarshal(t *testing.T) {
	accountData := map[string]interface{}{
		"id": "acct_123",
		"business_profile": map[string]interface{}{
			"mcc": "123",
		},
		"business_type": "company",
		"capabilities": map[string]interface{}{
			"card_payments": "active",
			"transfers":     "inactive",
		},
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
		"metadata": map[string]interface{}{
			"key1": "value1",
			"key2": "value2",
		},
		"object": "account",
		"requirements": map[string]interface{}{
			"current_deadline": 1234567890,
			"currently_due": []interface{}{
				"tos_acceptance.date",
				"tos_acceptance.ip",
			},
			"disabled_reason": "rejected.fraud",
			"errors": []map[string]interface{}{
				{
					"code":        "invalid_value_other",
					"reason":      "This value is invalid",
					"requirement": "tos_acceptance.date",
				},
			},
			"eventually_due": []interface{}{
				"relationship.representative",
			},
			"past_due": []interface{}{},
		},
		"settings": map[string]interface{}{
			"branding": map[string]interface{}{
				"icon": "file_123",
				"logo": "file_234",
			},
			"card_payments": map[string]interface{}{
				"decline_on": map[string]interface{}{
					"avs_failure": true,
					"cvc_failure": false,
				},
				"statement_descriptor_prefix": "prefix",
			},
			"payments": map[string]interface{}{
				"statement_descriptor": "descriptor",
			},
			"payouts": map[string]interface{}{
				"debit_negative_balances": true,
				"schedule": map[string]interface{}{
					"delay_days": 2,
					"interval":   "weekly",
				},
				"statement_descriptor_prefix": "prefix",
			},
		},
		"tos_acceptance": map[string]interface{}{
			"date":              1528573382,
			"ip":                "127.0.0.1",
			"user_agent":        "user agent",
			"service_agreement": "recipient",
		},
		"type": "custom",
	}

	bytes, err := json.Marshal(&accountData)
	assert.NoError(t, err)

	var account Account
	err = json.Unmarshal(bytes, &account)
	assert.NoError(t, err)

	assert.Equal(t, "acct_123", account.ID)

	assert.Equal(t, "123", account.BusinessProfile.MCC)

	assert.Equal(t, AccountCapabilityStatusActive, account.Capabilities.CardPayments)
	assert.Equal(t, AccountCapabilityStatus(""), account.Capabilities.LegacyPayments)
	assert.Equal(t, AccountCapabilityStatusInactive, account.Capabilities.Transfers)

	assert.Equal(t, AccountBusinessTypeCompany, account.BusinessType)

	// Assert ExternalAccounts are fully deserialized
	assert.Equal(t, true, account.ExternalAccounts.HasMore)
	assert.Equal(t, 2, len(account.ExternalAccounts.Data))
	assert.Equal(t, "ba_123", account.ExternalAccounts.Data[0].ID)
	assert.Equal(t, "card_123", account.ExternalAccounts.Data[1].ID)

	assert.Equal(t, "value1", account.Metadata["key1"])
	assert.Equal(t, "value2", account.Metadata["key2"])

	assert.Equal(t, int64(1234567890), account.Requirements.CurrentDeadline)
	assert.Equal(t, 2, len(account.Requirements.CurrentlyDue))
	assert.Equal(t, AccountRequirementsDisabledReasonRejectedFraud, account.Requirements.DisabledReason)
	assert.Equal(t, 1, len(account.Requirements.Errors))
	assert.Equal(t, "invalid_value_other", account.Requirements.Errors[0].Code)
	assert.Equal(t, 1, len(account.Requirements.EventuallyDue))
	assert.Equal(t, 0, len(account.Requirements.PastDue))

	assert.Equal(t, "file_123", account.Settings.Branding.Icon.ID)
	assert.Equal(t, "file_234", account.Settings.Branding.Logo.ID)
	assert.Equal(t, true, account.Settings.CardPayments.DeclineOn.AVSFailure)
	assert.Equal(t, false, account.Settings.CardPayments.DeclineOn.CVCFailure)
	assert.Equal(t, "prefix", account.Settings.CardPayments.StatementDescriptorPrefix)
	assert.Equal(t, "descriptor", account.Settings.Payments.StatementDescriptor)
	assert.Equal(t, true, account.Settings.Payouts.DebitNegativeBalances)
	assert.Equal(t, int64(2), account.Settings.Payouts.Schedule.DelayDays)
	assert.Equal(t, AccountSettingsPayoutsScheduleIntervalWeekly, account.Settings.Payouts.Schedule.Interval)

	assert.Equal(t, int64(1528573382), account.TOSAcceptance.Date)
	assert.Equal(t, "127.0.0.1", account.TOSAcceptance.IP)
	assert.Equal(t, "user agent", account.TOSAcceptance.UserAgent)
	assert.Equal(t, AccountTOSAcceptanceServiceAgreementRecipient, account.TOSAcceptance.ServiceAgreement)

	assert.Equal(t, AccountTypeCustom, account.Type)
}

func TestAccount_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON string
	{
		var v Account
		err := json.Unmarshal([]byte(`"acct_123"`), &v)
		assert.NoError(t, err)
		assert.Equal(t, "acct_123", v.ID)
	}

	// Unmarshals from a JSON object
	{
		v := Account{ID: "acct_123"}
		data, err := json.Marshal(&v)
		assert.NoError(t, err)

		err = json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, "acct_123", v.ID)
	}
}

func TestExternalAccount_UnmarshalJSON(t *testing.T) {
	// Unmarshals from a JSON object
	{
		// We build the JSON object manually here because it's key that the
		// `object` field is included so that the source knows what type to
		// decode
		data := []byte(`{"id":"ba_123", "object":"bank_account"}`)

		var v AccountExternalAccount
		err := json.Unmarshal(data, &v)
		assert.NoError(t, err)
		assert.Equal(t, AccountExternalAccountTypeBankAccount, v.Type)

		// The external account has a field for each possible type, so the
		// bank account is located one level down
		assert.Equal(t, "ba_123", v.BankAccount.ID)
	}
}

func TestPayoutScheduleParams_AppendTo(t *testing.T) {
	{
		params := &AccountSettingsPayoutsScheduleParams{DelayDaysMinimum: Bool(true)}
		body := &form.Values{}
		form.AppendTo(body, params)
		t.Logf("body = %+v", body)
		assert.Equal(t, []string{"minimum"}, body.Get("delay_days"))
	}
}
