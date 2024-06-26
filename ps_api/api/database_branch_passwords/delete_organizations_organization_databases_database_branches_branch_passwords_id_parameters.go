// Code generated by go-swagger; DO NOT EDIT.

package database_branch_passwords

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams creates a new DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams() *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams {
	return &DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParamsWithTimeout creates a new DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams object
// with the ability to set a timeout on a request.
func NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParamsWithTimeout(timeout time.Duration) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams {
	return &DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams{
		timeout: timeout,
	}
}

// NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParamsWithContext creates a new DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams object
// with the ability to set a context for a request.
func NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParamsWithContext(ctx context.Context) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams {
	return &DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams{
		Context: ctx,
	}
}

// NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParamsWithHTTPClient creates a new DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParamsWithHTTPClient(client *http.Client) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams {
	return &DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams{
		HTTPClient: client,
	}
}

/*
DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams contains all the parameters to send to the API endpoint

	for the delete organizations organization databases database branches branch passwords id operation.

	Typically these are written to a http.Request.
*/
type DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams struct {

	/* Branch.

	   The name of the branch the password belongs to
	*/
	Branch string

	/* Database.

	   The name of the database the password belongs to
	*/
	Database string

	/* ID.

	   The ID of the password
	*/
	ID string

	/* Organization.

	   The name of the organization the password belongs to
	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete organizations organization databases database branches branch passwords id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) WithDefaults() *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete organizations organization databases database branches branch passwords id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete organizations organization databases database branches branch passwords id params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) WithTimeout(timeout time.Duration) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organizations organization databases database branches branch passwords id params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organizations organization databases database branches branch passwords id params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) WithContext(ctx context.Context) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organizations organization databases database branches branch passwords id params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organizations organization databases database branches branch passwords id params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) WithHTTPClient(client *http.Client) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organizations organization databases database branches branch passwords id params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBranch adds the branch to the delete organizations organization databases database branches branch passwords id params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) WithBranch(branch string) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams {
	o.SetBranch(branch)
	return o
}

// SetBranch adds the branch to the delete organizations organization databases database branches branch passwords id params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) SetBranch(branch string) {
	o.Branch = branch
}

// WithDatabase adds the database to the delete organizations organization databases database branches branch passwords id params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) WithDatabase(database string) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams {
	o.SetDatabase(database)
	return o
}

// SetDatabase adds the database to the delete organizations organization databases database branches branch passwords id params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) SetDatabase(database string) {
	o.Database = database
}

// WithID adds the id to the delete organizations organization databases database branches branch passwords id params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) WithID(id string) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete organizations organization databases database branches branch passwords id params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) SetID(id string) {
	o.ID = id
}

// WithOrganization adds the organization to the delete organizations organization databases database branches branch passwords id params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) WithOrganization(organization string) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the delete organizations organization databases database branches branch passwords id params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesBranchPasswordsIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param branch
	if err := r.SetPathParam("branch", o.Branch); err != nil {
		return err
	}

	// path param database
	if err := r.SetPathParam("database", o.Database); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
