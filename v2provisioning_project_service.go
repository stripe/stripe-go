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

// v2ProvisioningProjectService is used to invoke project related APIs.
type v2ProvisioningProjectService struct {
	B   Backend
	Key string
}

// Creates a new project.
func (c v2ProvisioningProjectService) Create(ctx context.Context, params *V2ProvisioningProjectCreateParams) (*V2ProvisioningProject, error) {
	if params == nil {
		params = &V2ProvisioningProjectCreateParams{}
	}
	params.Context = ctx
	project := &V2ProvisioningProject{}
	err := c.B.Call(
		http.MethodPost, "/v2/provisioning/projects", c.Key, params, project)
	return project, err
}
