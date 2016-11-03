package stripe

import (
	"encoding/json"
	"testing"
)

func TestApplicationUnmarshal(t *testing.T) {
	applicationData := map[string]interface{}{
		"id":   "ca_1234",
		"name": "My Application Name",
	}

	bytes, err := json.Marshal(&applicationData)
	if err != nil {
		t.Error(err)
	}

	var application Application
	err = json.Unmarshal(bytes, &application)
	if err != nil {
		t.Error(err)
	}

	if application.ID != "ca_1234" {
		t.Errorf("Problem deserializing application, got ID %v", application.ID)
	}

	if application.Name != "My Application Name" {
		t.Errorf("Problem deserializing application, got name %v", application.Name)
	}
}
