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

// NewPostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams creates a new PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams() *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams {
	return &PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParamsWithTimeout creates a new PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams object
// with the ability to set a timeout on a request.
func NewPostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParamsWithTimeout(timeout time.Duration) *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams {
	return &PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams{
		timeout: timeout,
	}
}

// NewPostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParamsWithContext creates a new PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams object
// with the ability to set a context for a request.
func NewPostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParamsWithContext(ctx context.Context) *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams {
	return &PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams{
		Context: ctx,
	}
}

// NewPostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParamsWithHTTPClient creates a new PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams object
// with the ability to set a custom HTTPClient for a request.
func NewPostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParamsWithHTTPClient(client *http.Client) *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams {
	return &PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams{
		HTTPClient: client,
	}
}

/*
PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams contains all the parameters to send to the API endpoint

	for the post organizations organization databases database deploy requests number cancel operation.

	Typically these are written to a http.Request.
*/
type PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams struct {

	/* Database.

	   The name of the deploy request's database
	*/
	Database string

	/* Number.

	   The number of the deploy request
	*/
	Number string

	/* Organization.

	   The name of the deploy request's organization
	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the post organizations organization databases database deploy requests number cancel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams) WithDefaults() *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the post organizations organization databases database deploy requests number cancel params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the post organizations organization databases database deploy requests number cancel params
func (o *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams) WithTimeout(timeout time.Duration) *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post organizations organization databases database deploy requests number cancel params
func (o *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post organizations organization databases database deploy requests number cancel params
func (o *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams) WithContext(ctx context.Context) *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post organizations organization databases database deploy requests number cancel params
func (o *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post organizations organization databases database deploy requests number cancel params
func (o *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams) WithHTTPClient(client *http.Client) *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post organizations organization databases database deploy requests number cancel params
func (o *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatabase adds the database to the post organizations organization databases database deploy requests number cancel params
func (o *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams) WithDatabase(database string) *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams {
	o.SetDatabase(database)
	return o
}

// SetDatabase adds the database to the post organizations organization databases database deploy requests number cancel params
func (o *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams) SetDatabase(database string) {
	o.Database = database
}

// WithNumber adds the number to the post organizations organization databases database deploy requests number cancel params
func (o *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams) WithNumber(number string) *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the post organizations organization databases database deploy requests number cancel params
func (o *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams) SetNumber(number string) {
	o.Number = number
}

// WithOrganization adds the organization to the post organizations organization databases database deploy requests number cancel params
func (o *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams) WithOrganization(organization string) *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the post organizations organization databases database deploy requests number cancel params
func (o *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *PostOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberCancelParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param database
	if err := r.SetPathParam("database", o.Database); err != nil {
		return err
	}

	// path param number
	if err := r.SetPathParam("number", o.Number); err != nil {
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
