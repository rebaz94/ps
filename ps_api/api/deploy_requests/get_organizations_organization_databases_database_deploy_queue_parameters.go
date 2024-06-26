// Code generated by go-swagger; DO NOT EDIT.

package deploy_requests

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

// NewGetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams creates a new GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams() *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams {
	return &GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationsOrganizationDatabasesDatabaseDeployQueueParamsWithTimeout creates a new GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationsOrganizationDatabasesDatabaseDeployQueueParamsWithTimeout(timeout time.Duration) *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams {
	return &GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams{
		timeout: timeout,
	}
}

// NewGetOrganizationsOrganizationDatabasesDatabaseDeployQueueParamsWithContext creates a new GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams object
// with the ability to set a context for a request.
func NewGetOrganizationsOrganizationDatabasesDatabaseDeployQueueParamsWithContext(ctx context.Context) *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams {
	return &GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams{
		Context: ctx,
	}
}

// NewGetOrganizationsOrganizationDatabasesDatabaseDeployQueueParamsWithHTTPClient creates a new GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationsOrganizationDatabasesDatabaseDeployQueueParamsWithHTTPClient(client *http.Client) *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams {
	return &GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams contains all the parameters to send to the API endpoint

	for the get organizations organization databases database deploy queue operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams struct {

	/* Database.

	   The name of the deploy request's database
	*/
	Database string

	/* Organization.

	   The name of the deploy request's organization
	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organizations organization databases database deploy queue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams) WithDefaults() *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organizations organization databases database deploy queue params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organizations organization databases database deploy queue params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams) WithTimeout(timeout time.Duration) *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organizations organization databases database deploy queue params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organizations organization databases database deploy queue params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams) WithContext(ctx context.Context) *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organizations organization databases database deploy queue params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organizations organization databases database deploy queue params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams) WithHTTPClient(client *http.Client) *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organizations organization databases database deploy queue params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatabase adds the database to the get organizations organization databases database deploy queue params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams) WithDatabase(database string) *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams {
	o.SetDatabase(database)
	return o
}

// SetDatabase adds the database to the get organizations organization databases database deploy queue params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams) SetDatabase(database string) {
	o.Database = database
}

// WithOrganization adds the organization to the get organizations organization databases database deploy queue params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams) WithOrganization(organization string) *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get organizations organization databases database deploy queue params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployQueueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param database
	if err := r.SetPathParam("database", o.Database); err != nil {
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
