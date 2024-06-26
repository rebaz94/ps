// Code generated by go-swagger; DO NOT EDIT.

package database_branches

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

// NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams creates a new DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams() *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams {
	return &DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParamsWithTimeout creates a new DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams object
// with the ability to set a timeout on a request.
func NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParamsWithTimeout(timeout time.Duration) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams {
	return &DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams{
		timeout: timeout,
	}
}

// NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParamsWithContext creates a new DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams object
// with the ability to set a context for a request.
func NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParamsWithContext(ctx context.Context) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams {
	return &DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams{
		Context: ctx,
	}
}

// NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParamsWithHTTPClient creates a new DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParamsWithHTTPClient(client *http.Client) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams {
	return &DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams{
		HTTPClient: client,
	}
}

/*
DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams contains all the parameters to send to the API endpoint

	for the delete organizations organization databases database branches name safe migrations operation.

	Typically these are written to a http.Request.
*/
type DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams struct {

	/* Database.

	   The name of the database the branch belongs to
	*/
	Database string

	/* Name.

	   The name of the branch
	*/
	Name string

	/* Organization.

	   The name of the organization the branch belongs to
	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete organizations organization databases database branches name safe migrations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams) WithDefaults() *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete organizations organization databases database branches name safe migrations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete organizations organization databases database branches name safe migrations params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams) WithTimeout(timeout time.Duration) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organizations organization databases database branches name safe migrations params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organizations organization databases database branches name safe migrations params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams) WithContext(ctx context.Context) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organizations organization databases database branches name safe migrations params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organizations organization databases database branches name safe migrations params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams) WithHTTPClient(client *http.Client) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organizations organization databases database branches name safe migrations params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatabase adds the database to the delete organizations organization databases database branches name safe migrations params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams) WithDatabase(database string) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams {
	o.SetDatabase(database)
	return o
}

// SetDatabase adds the database to the delete organizations organization databases database branches name safe migrations params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams) SetDatabase(database string) {
	o.Database = database
}

// WithName adds the name to the delete organizations organization databases database branches name safe migrations params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams) WithName(name string) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete organizations organization databases database branches name safe migrations params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams) SetName(name string) {
	o.Name = name
}

// WithOrganization adds the organization to the delete organizations organization databases database branches name safe migrations params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams) WithOrganization(organization string) *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the delete organizations organization databases database branches name safe migrations params
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationsOrganizationDatabasesDatabaseBranchesNameSafeMigrationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param database
	if err := r.SetPathParam("database", o.Database); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
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
