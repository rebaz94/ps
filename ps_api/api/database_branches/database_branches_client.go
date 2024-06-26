// Code generated by go-swagger; DO NOT EDIT.

package database_branches

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new database branches API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for database branches API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteOrganizationsOrganizationDatabasesDatabaseBranchesName(params *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameNoContent, error)

	DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrations(params *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsOK, error)

	GetOrganizationsOrganizationDatabasesDatabaseBranches(params *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationsOrganizationDatabasesDatabaseBranchesOK, error)

	GetOrganizationsOrganizationDatabasesDatabaseBranchesName(params *GetOrganizationsOrganizationDatabasesDatabaseBranchesNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationsOrganizationDatabasesDatabaseBranchesNameOK, error)

	GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchema(params *GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaOK, error)

	GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaLint(params *GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaLintParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaLintOK, error)

	PostOrganizationsOrganizationDatabasesDatabaseBranches(params *PostOrganizationsOrganizationDatabasesDatabaseBranchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOrganizationsOrganizationDatabasesDatabaseBranchesCreated, error)

	PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemote(params *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteOK, error)

	PostOrganizationsOrganizationDatabasesDatabaseBranchesNamePromote(params *PostOrganizationsOrganizationDatabasesDatabaseBranchesNamePromoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOrganizationsOrganizationDatabasesDatabaseBranchesNamePromoteOK, error)

	PostOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrations(params *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
	DeleteOrganizationsOrganizationDatabasesDatabaseBranchesName deletes a branch

### Authorization
A service token or OAuth token must have at least one of the following access or scopes in order to use this API endpoint:

**Service Token Accesses**

	`delete_branch`

**OAuth Scopes**

	| Resource | Scopes |

| :------- | :---------- |
| Organization | `delete_branches`, `delete_production_branches` |
| Database | `delete_branches`, `delete_production_branches` |
| Branch | `delete_branch` |
*/
func (a *Client) DeleteOrganizationsOrganizationDatabasesDatabaseBranchesName(params *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_organizations_organization_databases_database_branches_name",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_organizations_organization_databases_database_branches_name: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrations disables safe migrations for a branch
*/
func (a *Client) DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrations(params *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "delete_organizations_organization_databases_database_branches_name_safe-migrations",
		Method:             "DELETE",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches/{name}/safe-migrations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsReader{formats: a.formats},
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
	success, ok := result.(*DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for delete_organizations_organization_databases_database_branches_name_safe-migrations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetOrganizationsOrganizationDatabasesDatabaseBranches lists branches

### Authorization
A service token or OAuth token must have at least one of the following access or scopes in order to use this API endpoint:

**Service Token Accesses**

	`read_branch`, `delete_branch`, `create_branch`, `connect_production_branch`, `connect_branch`

**OAuth Scopes**

	| Resource | Scopes |

| :------- | :---------- |
| Organization | `read_branches` |
| Database | `read_branches` |
| Branch | `read_branch` |
*/
func (a *Client) GetOrganizationsOrganizationDatabasesDatabaseBranches(params *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationsOrganizationDatabasesDatabaseBranchesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationsOrganizationDatabasesDatabaseBranchesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_organizations_organization_databases_database_branches",
		Method:             "GET",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationsOrganizationDatabasesDatabaseBranchesReader{formats: a.formats},
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
	success, ok := result.(*GetOrganizationsOrganizationDatabasesDatabaseBranchesOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_organizations_organization_databases_database_branches: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetOrganizationsOrganizationDatabasesDatabaseBranchesName gets a branch

### Authorization
A service token or OAuth token must have at least one of the following access or scopes in order to use this API endpoint:

**Service Token Accesses**

	`read_branch`, `delete_branch`, `create_branch`, `connect_production_branch`, `connect_branch`

**OAuth Scopes**

	| Resource | Scopes |

| :------- | :---------- |
| Organization | `read_branches` |
| Database | `read_branches` |
| Branch | `read_branch` |
*/
func (a *Client) GetOrganizationsOrganizationDatabasesDatabaseBranchesName(params *GetOrganizationsOrganizationDatabasesDatabaseBranchesNameParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationsOrganizationDatabasesDatabaseBranchesNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationsOrganizationDatabasesDatabaseBranchesNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_organizations_organization_databases_database_branches_name",
		Method:             "GET",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationsOrganizationDatabasesDatabaseBranchesNameReader{formats: a.formats},
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
	success, ok := result.(*GetOrganizationsOrganizationDatabasesDatabaseBranchesNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_organizations_organization_databases_database_branches_name: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchema gets a branch schema

### Authorization
A service token or OAuth token must have at least one of the following access or scopes in order to use this API endpoint:

**Service Token Accesses**

	`read_branch`, `delete_branch`, `create_branch`, `connect_production_branch`, `connect_branch`

**OAuth Scopes**

	| Resource | Scopes |

| :------- | :---------- |
| Organization | `read_branches` |
| Database | `read_branches` |
| Branch | `read_branch` |
*/
func (a *Client) GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchema(params *GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_organizations_organization_databases_database_branches_name_schema",
		Method:             "GET",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches/{name}/schema",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaReader{formats: a.formats},
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
	success, ok := result.(*GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_organizations_organization_databases_database_branches_name_schema: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaLint lints a branch schema

### Authorization
A service token or OAuth token must have at least one of the following access or scopes in order to use this API endpoint:

**Service Token Accesses**

	`read_branch`, `delete_branch`, `create_branch`, `connect_production_branch`, `connect_branch`

**OAuth Scopes**

	| Resource | Scopes |

| :------- | :---------- |
| Organization | `read_branches` |
| Database | `read_branches` |
| Branch | `read_branch` |
*/
func (a *Client) GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaLint(params *GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaLintParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaLintOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaLintParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "get_organizations_organization_databases_database_branches_name_schema_lint",
		Method:             "GET",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches/{name}/schema/lint",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaLintReader{formats: a.formats},
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
	success, ok := result.(*GetOrganizationsOrganizationDatabasesDatabaseBranchesNameSchemaLintOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for get_organizations_organization_databases_database_branches_name_schema_lint: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PostOrganizationsOrganizationDatabasesDatabaseBranches creates a branch

### Authorization
A service token or OAuth token must have at least one of the following access or scopes in order to use this API endpoint:

**Service Token Accesses**

	`create_branch`, `restore_production_branch_backup`, `restore_backup`

**OAuth Scopes**

	| Resource | Scopes |

| :------- | :---------- |
| Organization | `write_branches`, `restore_production_branch_backups`, `restore_backups` |
| Database | `write_branches`, `restore_production_branch_backups`, `restore_backups` |
| Branch | `restore_backups` |
*/
func (a *Client) PostOrganizationsOrganizationDatabasesDatabaseBranches(params *PostOrganizationsOrganizationDatabasesDatabaseBranchesParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOrganizationsOrganizationDatabasesDatabaseBranchesCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOrganizationsOrganizationDatabasesDatabaseBranchesParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_organizations_organization_databases_database_branches",
		Method:             "POST",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOrganizationsOrganizationDatabasesDatabaseBranchesReader{formats: a.formats},
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
	success, ok := result.(*PostOrganizationsOrganizationDatabasesDatabaseBranchesCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_organizations_organization_databases_database_branches: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemote demotes a branch

	Demotes a branch from production to development

### Authorization
A   OAuth token must have at least one of the following   scopes in order to use this API endpoint:

**OAuth Scopes**

	| Resource | Scopes |

| :------- | :---------- |
| Organization | `demote_branches` |
| Database | `demote_branches` |
*/
func (a *Client) PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemote(params *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_organizations_organization_databases_database_branches_name_demote",
		Method:             "POST",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches/{name}/demote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteReader{formats: a.formats},
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
	success, ok := result.(*PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_organizations_organization_databases_database_branches_name_demote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
	PostOrganizationsOrganizationDatabasesDatabaseBranchesNamePromote promotes a branch

	Promotes a branch from development to production

### Authorization
A service token or OAuth token must have at least one of the following access or scopes in order to use this API endpoint:

**Service Token Accesses**

	`connect_production_branch`

**OAuth Scopes**

	| Resource | Scopes |

| :------- | :---------- |
| Organization | `promote_branches` |
| Database | `promote_branches` |
*/
func (a *Client) PostOrganizationsOrganizationDatabasesDatabaseBranchesNamePromote(params *PostOrganizationsOrganizationDatabasesDatabaseBranchesNamePromoteParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOrganizationsOrganizationDatabasesDatabaseBranchesNamePromoteOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOrganizationsOrganizationDatabasesDatabaseBranchesNamePromoteParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_organizations_organization_databases_database_branches_name_promote",
		Method:             "POST",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches/{name}/promote",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOrganizationsOrganizationDatabasesDatabaseBranchesNamePromoteReader{formats: a.formats},
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
	success, ok := result.(*PostOrganizationsOrganizationDatabasesDatabaseBranchesNamePromoteOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_organizations_organization_databases_database_branches_name_promote: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
PostOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrations enables safe migrations for a branch
*/
func (a *Client) PostOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrations(params *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*PostOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "post_organizations_organization_databases_database_branches_name_safe-migrations",
		Method:             "POST",
		PathPattern:        "/organizations/{organization}/databases/{database}/branches/{name}/safe-migrations",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &PostOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsReader{formats: a.formats},
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
	success, ok := result.(*PostOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for post_organizations_organization_databases_database_branches_name_safe-migrations: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
