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
	"github.com/go-openapi/swag"
)

// NewGetOrganizationsOrganizationDatabasesDatabaseBranchesParams creates a new GetOrganizationsOrganizationDatabasesDatabaseBranchesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationsOrganizationDatabasesDatabaseBranchesParams() *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams {
	return &GetOrganizationsOrganizationDatabasesDatabaseBranchesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationsOrganizationDatabasesDatabaseBranchesParamsWithTimeout creates a new GetOrganizationsOrganizationDatabasesDatabaseBranchesParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationsOrganizationDatabasesDatabaseBranchesParamsWithTimeout(timeout time.Duration) *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams {
	return &GetOrganizationsOrganizationDatabasesDatabaseBranchesParams{
		timeout: timeout,
	}
}

// NewGetOrganizationsOrganizationDatabasesDatabaseBranchesParamsWithContext creates a new GetOrganizationsOrganizationDatabasesDatabaseBranchesParams object
// with the ability to set a context for a request.
func NewGetOrganizationsOrganizationDatabasesDatabaseBranchesParamsWithContext(ctx context.Context) *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams {
	return &GetOrganizationsOrganizationDatabasesDatabaseBranchesParams{
		Context: ctx,
	}
}

// NewGetOrganizationsOrganizationDatabasesDatabaseBranchesParamsWithHTTPClient creates a new GetOrganizationsOrganizationDatabasesDatabaseBranchesParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationsOrganizationDatabasesDatabaseBranchesParamsWithHTTPClient(client *http.Client) *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams {
	return &GetOrganizationsOrganizationDatabasesDatabaseBranchesParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationsOrganizationDatabasesDatabaseBranchesParams contains all the parameters to send to the API endpoint

	for the get organizations organization databases database branches operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationsOrganizationDatabasesDatabaseBranchesParams struct {

	/* Database.

	   The name of the database the branch belongs to
	*/
	Database string

	/* Organization.

	   The name of the organization the branch belongs to
	*/
	Organization string

	/* Page.

	   If provided, specifies the page offset of returned results

	   Default: 1
	*/
	Page *float64

	/* PerPage.

	   If provided, specifies the number of returned results

	   Default: 25
	*/
	PerPage *float64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organizations organization databases database branches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) WithDefaults() *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organizations organization databases database branches params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) SetDefaults() {
	var (
		pageDefault = float64(1)

		perPageDefault = float64(25)
	)

	val := GetOrganizationsOrganizationDatabasesDatabaseBranchesParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get organizations organization databases database branches params
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) WithTimeout(timeout time.Duration) *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organizations organization databases database branches params
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organizations organization databases database branches params
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) WithContext(ctx context.Context) *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organizations organization databases database branches params
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organizations organization databases database branches params
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) WithHTTPClient(client *http.Client) *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organizations organization databases database branches params
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatabase adds the database to the get organizations organization databases database branches params
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) WithDatabase(database string) *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams {
	o.SetDatabase(database)
	return o
}

// SetDatabase adds the database to the get organizations organization databases database branches params
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) SetDatabase(database string) {
	o.Database = database
}

// WithOrganization adds the organization to the get organizations organization databases database branches params
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) WithOrganization(organization string) *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get organizations organization databases database branches params
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPage adds the page to the get organizations organization databases database branches params
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) WithPage(page *float64) *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get organizations organization databases database branches params
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) SetPage(page *float64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get organizations organization databases database branches params
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) WithPerPage(perPage *float64) *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organizations organization databases database branches params
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) SetPerPage(perPage *float64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationsOrganizationDatabasesDatabaseBranchesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Page != nil {

		// query param page
		var qrPage float64

		if o.Page != nil {
			qrPage = *o.Page
		}
		qPage := swag.FormatFloat64(qrPage)
		if qPage != "" {

			if err := r.SetQueryParam("page", qPage); err != nil {
				return err
			}
		}
	}

	if o.PerPage != nil {

		// query param per_page
		var qrPerPage float64

		if o.PerPage != nil {
			qrPerPage = *o.PerPage
		}
		qPerPage := swag.FormatFloat64(qrPerPage)
		if qPerPage != "" {

			if err := r.SetQueryParam("per_page", qPerPage); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
