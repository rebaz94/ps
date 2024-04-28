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

// NewGetOrganizationsOrganizationDatabasesNameParams creates a new GetOrganizationsOrganizationDatabasesNameParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationsOrganizationDatabasesNameParams() *GetOrganizationsOrganizationDatabasesNameParams {
	return &GetOrganizationsOrganizationDatabasesNameParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationsOrganizationDatabasesNameParamsWithTimeout creates a new GetOrganizationsOrganizationDatabasesNameParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationsOrganizationDatabasesNameParamsWithTimeout(timeout time.Duration) *GetOrganizationsOrganizationDatabasesNameParams {
	return &GetOrganizationsOrganizationDatabasesNameParams{
		timeout: timeout,
	}
}

// NewGetOrganizationsOrganizationDatabasesNameParamsWithContext creates a new GetOrganizationsOrganizationDatabasesNameParams object
// with the ability to set a context for a request.
func NewGetOrganizationsOrganizationDatabasesNameParamsWithContext(ctx context.Context) *GetOrganizationsOrganizationDatabasesNameParams {
	return &GetOrganizationsOrganizationDatabasesNameParams{
		Context: ctx,
	}
}

// NewGetOrganizationsOrganizationDatabasesNameParamsWithHTTPClient creates a new GetOrganizationsOrganizationDatabasesNameParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationsOrganizationDatabasesNameParamsWithHTTPClient(client *http.Client) *GetOrganizationsOrganizationDatabasesNameParams {
	return &GetOrganizationsOrganizationDatabasesNameParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationsOrganizationDatabasesNameParams contains all the parameters to send to the API endpoint

	for the get organizations organization databases name operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationsOrganizationDatabasesNameParams struct {

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

// WithDefaults hydrates default values in the get organizations organization databases name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsOrganizationDatabasesNameParams) WithDefaults() *GetOrganizationsOrganizationDatabasesNameParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organizations organization databases name params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsOrganizationDatabasesNameParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organizations organization databases name params
func (o *GetOrganizationsOrganizationDatabasesNameParams) WithTimeout(timeout time.Duration) *GetOrganizationsOrganizationDatabasesNameParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organizations organization databases name params
func (o *GetOrganizationsOrganizationDatabasesNameParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organizations organization databases name params
func (o *GetOrganizationsOrganizationDatabasesNameParams) WithContext(ctx context.Context) *GetOrganizationsOrganizationDatabasesNameParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organizations organization databases name params
func (o *GetOrganizationsOrganizationDatabasesNameParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organizations organization databases name params
func (o *GetOrganizationsOrganizationDatabasesNameParams) WithHTTPClient(client *http.Client) *GetOrganizationsOrganizationDatabasesNameParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organizations organization databases name params
func (o *GetOrganizationsOrganizationDatabasesNameParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithName adds the name to the get organizations organization databases name params
func (o *GetOrganizationsOrganizationDatabasesNameParams) WithName(name string) *GetOrganizationsOrganizationDatabasesNameParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the get organizations organization databases name params
func (o *GetOrganizationsOrganizationDatabasesNameParams) SetName(name string) {
	o.Name = name
}

// WithOrganization adds the organization to the get organizations organization databases name params
func (o *GetOrganizationsOrganizationDatabasesNameParams) WithOrganization(organization string) *GetOrganizationsOrganizationDatabasesNameParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get organizations organization databases name params
func (o *GetOrganizationsOrganizationDatabasesNameParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationsOrganizationDatabasesNameParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
