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

// NewPostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams creates a new PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams() *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams {
	return &PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParamsWithTimeout creates a new PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams object
// with the ability to set a timeout on a request.
func NewPostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParamsWithTimeout(timeout time.Duration) *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams {
	return &PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams{
		timeout: timeout,
	}
}

// NewPostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParamsWithContext creates a new PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams object
// with the ability to set a context for a request.
func NewPostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParamsWithContext(ctx context.Context) *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams {
	return &PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams{
		Context: ctx,
	}
}

// NewPostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParamsWithHTTPClient creates a new PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParamsWithHTTPClient(client *http.Client) *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams {
	return &PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams{
		HTTPClient: client,
	}
}

/*
PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams contains all the parameters to send to the API endpoint

	for the post organizations organization databases database branches name demote operation.

	Typically these are written to a http.Request.
*/
type PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams struct {

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

// WithDefaults hydrates default values in the post organizations organization databases database branches name demote params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams) WithDefaults() *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post organizations organization databases database branches name demote params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post organizations organization databases database branches name demote params
func (o *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams) WithTimeout(timeout time.Duration) *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post organizations organization databases database branches name demote params
func (o *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post organizations organization databases database branches name demote params
func (o *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams) WithContext(ctx context.Context) *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post organizations organization databases database branches name demote params
func (o *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post organizations organization databases database branches name demote params
func (o *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams) WithHTTPClient(client *http.Client) *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post organizations organization databases database branches name demote params
func (o *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatabase adds the database to the post organizations organization databases database branches name demote params
func (o *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams) WithDatabase(database string) *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams {
	o.SetDatabase(database)
	return o
}

// SetDatabase adds the database to the post organizations organization databases database branches name demote params
func (o *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams) SetDatabase(database string) {
	o.Database = database
}

// WithName adds the name to the post organizations organization databases database branches name demote params
func (o *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams) WithName(name string) *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the post organizations organization databases database branches name demote params
func (o *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams) SetName(name string) {
	o.Name = name
}

// WithOrganization adds the organization to the post organizations organization databases database branches name demote params
func (o *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams) WithOrganization(organization string) *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the post organizations organization databases database branches name demote params
func (o *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *PostOrganizationsOrganizationDatabasesDatabaseBranchesNameDemoteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
