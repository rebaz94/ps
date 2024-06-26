// Code generated by go-swagger; DO NOT EDIT.

package database_branch_passwords

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new database branch passwords API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for database branch passwords API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsID(params *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDNoContent, error)

	GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswords(params *GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsOK, error)

	GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsID(params *GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDOK, error)

	PatchOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsID(params *PatchOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDOK, error)

	PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswords(params *PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsCreated, error)

	PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDRenew(params *PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDRenewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDRenewOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsID deletes a password

### Authorization
A service token or OAuth token must have at least one of the following access or scopes in order to use this API endpoint:

**Service Token Accesses**

	`delete_production_branch_password`, `delete_branch_password`

**OAuth Scopes**

	| Resource | Scopes |

| :------- | :---------- |
| Organization | `manage_passwords`, `manage_production_branch_passwords` |
| Database | `manage_passwords`, `manage_production_branch_passwords` |
| Branch | `manage_passwords` |
*/
func (a *Client) DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsID(params *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_organizations_organization_databases_database_branches_branch_passwords_id",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches/{branch}/passwords/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_organizations_organization_databases_database_branches_branch_passwords_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswords lists passwords

### Authorization
A service token or OAuth token must have at least one of the following access or scopes in order to use this API endpoint:

**Service Token Accesses**

	`read_branch`, `delete_branch`, `create_branch`, `connect_production_branch`, `connect_branch`

**OAuth Scopes**

	| Resource | Scopes |

| :------- | :---------- |
| Organization | `manage_passwords`, `manage_production_branch_passwords` |
| Database | `manage_passwords`, `manage_production_branch_passwords` |
| Branch | `manage_passwords` |
*/
func (a *Client) GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswords(params *GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_organizations_organization_databases_database_branches_branch_passwords",
		Method:             "GET",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches/{branch}/passwords",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_organizations_organization_databases_database_branches_branch_passwords: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsID gets a password

### Authorization
A service token or OAuth token must have at least one of the following access or scopes in order to use this API endpoint:

**Service Token Accesses**

	`read_branch`, `delete_branch`, `create_branch`, `connect_production_branch`, `connect_branch`

**OAuth Scopes**

	| Resource | Scopes |

| :------- | :---------- |
| Organization | `manage_passwords`, `manage_production_branch_passwords` |
| Database | `manage_passwords`, `manage_production_branch_passwords` |
| Branch | `manage_passwords` |
*/
func (a *Client) GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsID(params *GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_organizations_organization_databases_database_branches_branch_passwords_id",
		Method:             "GET",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches/{branch}/passwords/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_organizations_organization_databases_database_branches_branch_passwords_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PatchOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsID updates a password

### Authorization
A service token or OAuth token must have at least one of the following access or scopes in order to use this API endpoint:

**Service Token Accesses**

	`connect_production_branch`, `connect_branch`

**OAuth Scopes**

	| Resource | Scopes |

| :------- | :---------- |
| Organization | `manage_passwords`, `manage_production_branch_passwords` |
| Database | `manage_passwords`, `manage_production_branch_passwords` |
| Branch | `manage_passwords` |
*/
func (a *Client) PatchOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsID(params *PatchOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PatchOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPatchOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "patch_organizations_organization_databases_database_branches_branch_passwords_id",
		Method:             "PATCH",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches/{branch}/passwords/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PatchOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PatchOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for patch_organizations_organization_databases_database_branches_branch_passwords_id: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswords creates a password

### Authorization
A service token or OAuth token must have at least one of the following access or scopes in order to use this API endpoint:

**Service Token Accesses**

	`connect_production_branch`, `connect_branch`

**OAuth Scopes**

	| Resource | Scopes |

| :------- | :---------- |
| Organization | `manage_passwords`, `manage_production_branch_passwords` |
| Database | `manage_passwords`, `manage_production_branch_passwords` |
| Branch | `manage_passwords` |
*/
func (a *Client) PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswords(params *PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_organizations_organization_databases_database_branches_branch_passwords",
		Method:             "POST",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches/{branch}/passwords",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_organizations_organization_databases_database_branches_branch_passwords: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDRenew renews a password

### Authorization
A service token or OAuth token must have at least one of the following access or scopes in order to use this API endpoint:

**Service Token Accesses**

	`connect_production_branch`, `connect_branch`

**OAuth Scopes**

	| Resource | Scopes |

| :------- | :---------- |
| Organization | `manage_passwords`, `manage_production_branch_passwords` |
| Database | `manage_passwords`, `manage_production_branch_passwords` |
| Branch | `manage_passwords` |
*/
func (a *Client) PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDRenew(params *PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDRenewParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDRenewOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDRenewParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_organizations_organization_databases_database_branches_branch_passwords_id_renew",
		Method:             "POST",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches/{branch}/passwords/{id}/renew",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDRenewReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PostOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDRenewOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_organizations_organization_databases_database_branches_branch_passwords_id_renew: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
