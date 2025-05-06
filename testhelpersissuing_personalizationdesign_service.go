//
//
// File generated from our OpenAPI spec
//
//

package stripe

import (
	"context"
	"net/http"
)

// v1TestHelpersIssuingPersonalizationDesignService is used to invoke /v1/issuing/personalization_designs APIs.
type v1TestHelpersIssuingPersonalizationDesignService struct {
	B   Backend
	Key string
}

// Updates the status of the specified testmode personalization design object to active.
func (c v1TestHelpersIssuingPersonalizationDesignService) Activate(ctx context.Context, id string, params *TestHelpersIssuingPersonalizationDesignActivateParams) (*IssuingPersonalizationDesign, error) {
	if params == nil {
		params = &TestHelpersIssuingPersonalizationDesignActivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/issuing/personalization_designs/%s/activate", id)
	personalizationdesign := &IssuingPersonalizationDesign{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

// Updates the status of the specified testmode personalization design object to inactive.
func (c v1TestHelpersIssuingPersonalizationDesignService) Deactivate(ctx context.Context, id string, params *TestHelpersIssuingPersonalizationDesignDeactivateParams) (*IssuingPersonalizationDesign, error) {
	if params == nil {
		params = &TestHelpersIssuingPersonalizationDesignDeactivateParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/issuing/personalization_designs/%s/deactivate", id)
	personalizationdesign := &IssuingPersonalizationDesign{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, personalizationdesign)
	return personalizationdesign, err
}

// Updates the status of the specified testmode personalization design object to rejected.
func (c v1TestHelpersIssuingPersonalizationDesignService) Reject(ctx context.Context, id string, params *TestHelpersIssuingPersonalizationDesignRejectParams) (*IssuingPersonalizationDesign, error) {
	if params == nil {
		params = &TestHelpersIssuingPersonalizationDesignRejectParams{}
	}
	params.Context = ctx
	path := FormatURLPath(
		"/v1/test_helpers/issuing/personalization_designs/%s/reject", id)
	personalizationdesign := &IssuingPersonalizationDesign{}
	err := c.B.Call(http.MethodPost, path, c.Key, params, personalizationdesign)
	return personalizationdesign, err
}
