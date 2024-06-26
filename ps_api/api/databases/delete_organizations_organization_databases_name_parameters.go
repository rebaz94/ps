// Code generated by go-swagger; DO NOT EDIT.

package databases

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

// NewDeleteOrganizationsOrganizationDatabasesNameParams creates a new DeleteOrganizationsOrganizationDatabasesNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrganizationsOrganizationDatabasesNameParams() *DeleteOrganizationsOrganizationDatabasesNameParams {
	return &DeleteOrganizationsOrganizationDatabasesNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationsOrganizationDatabasesNameParamsWithTimeout creates a new DeleteOrganizationsOrganizationDatabasesNameParams object
// with the ability to set a timeout on a request.
func NewDeleteOrganizationsOrganizationDatabasesNameParamsWithTimeout(timeout time.Duration) *DeleteOrganizationsOrganizationDatabasesNameParams {
	return &DeleteOrganizationsOrganizationDatabasesNameParams{
		timeout: timeout,
	}
}

// NewDeleteOrganizationsOrganizationDatabasesNameParamsWithContext creates a new DeleteOrganizationsOrganizationDatabasesNameParams object
// with the ability to set a context for a request.
func NewDeleteOrganizationsOrganizationDatabasesNameParamsWithContext(ctx context.Context) *DeleteOrganizationsOrganizationDatabasesNameParams {
	return &DeleteOrganizationsOrganizationDatabasesNameParams{
		Context: ctx,
	}
}

// NewDeleteOrganizationsOrganizationDatabasesNameParamsWithHTTPClient creates a new DeleteOrganizationsOrganizationDatabasesNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrganizationsOrganizationDatabasesNameParamsWithHTTPClient(client *http.Client) *DeleteOrganizationsOrganizationDatabasesNameParams {
	return &DeleteOrganizationsOrganizationDatabasesNameParams{
		HTTPClient: client,
	}
}

/*
DeleteOrganizationsOrganizationDatabasesNameParams contains all the parameters to send to the API endpoint

	for the delete organizations organization databases name operation.

	Typically these are written to a http.Request.
*/
type DeleteOrganizationsOrganizationDatabasesNameParams struct {

	/* Name.

	   The name of the database
	*/
	Name string

	/* Organization.

	   The name of the organization the database belongs to
	*/
	Organization string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the delete organizations organization databases name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationsOrganizationDatabasesNameParams) WithDefaults() *DeleteOrganizationsOrganizationDatabasesNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete organizations organization databases name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationsOrganizationDatabasesNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete organizations organization databases name params
func (o *DeleteOrganizationsOrganizationDatabasesNameParams) WithTimeout(timeout time.Duration) *DeleteOrganizationsOrganizationDatabasesNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organizations organization databases name params
func (o *DeleteOrganizationsOrganizationDatabasesNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organizations organization databases name params
func (o *DeleteOrganizationsOrganizationDatabasesNameParams) WithContext(ctx context.Context) *DeleteOrganizationsOrganizationDatabasesNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organizations organization databases name params
func (o *DeleteOrganizationsOrganizationDatabasesNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organizations organization databases name params
func (o *DeleteOrganizationsOrganizationDatabasesNameParams) WithHTTPClient(client *http.Client) *DeleteOrganizationsOrganizationDatabasesNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organizations organization databases name params
func (o *DeleteOrganizationsOrganizationDatabasesNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the delete organizations organization databases name params
func (o *DeleteOrganizationsOrganizationDatabasesNameParams) WithName(name string) *DeleteOrganizationsOrganizationDatabasesNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the delete organizations organization databases name params
func (o *DeleteOrganizationsOrganizationDatabasesNameParams) SetName(name string) {
	o.Name = name
}

// WithOrganization adds the organization to the delete organizations organization databases name params
func (o *DeleteOrganizationsOrganizationDatabasesNameParams) WithOrganization(organization string) *DeleteOrganizationsOrganizationDatabasesNameParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the delete organizations organization databases name params
func (o *DeleteOrganizationsOrganizationDatabasesNameParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationsOrganizationDatabasesNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
