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
	"github.com/go-openapi/swag"
)

// NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams creates a new GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams() *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams {
	return &GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParamsWithTimeout creates a new GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParamsWithTimeout(timeout time.Duration) *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams {
	return &GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams{
		timeout: timeout,
	}
}

// NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParamsWithContext creates a new GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams object
// with the ability to set a context for a request.
func NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParamsWithContext(ctx context.Context) *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams {
	return &GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams{
		Context: ctx,
	}
}

// NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParamsWithHTTPClient creates a new GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParamsWithHTTPClient(client *http.Client) *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams {
	return &GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams contains all the parameters to send to the API endpoint

	for the get organizations organization databases database deploy requests number operations operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams struct {

	/* Database.

	   The name of the database the deploy request belongs to
	*/
	Database string

	/* Number.

	   The number of the deploy request
	*/
	Number string

	/* Organization.

	   The name of the organization the deploy request belongs to
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

// WithDefaults hydrates default values in the get organizations organization databases database deploy requests number operations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) WithDefaults() *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organizations organization databases database deploy requests number operations params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) SetDefaults() {
	var (
		pageDefault = float64(1)

		perPageDefault = float64(25)
	)

	val := GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) WithTimeout(timeout time.Duration) *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) WithContext(ctx context.Context) *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) WithHTTPClient(client *http.Client) *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDatabase adds the database to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) WithDatabase(database string) *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams {
	o.SetDatabase(database)
	return o
}

// SetDatabase adds the database to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) SetDatabase(database string) {
	o.Database = database
}

// WithNumber adds the number to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) WithNumber(number string) *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams {
	o.SetNumber(number)
	return o
}

// SetNumber adds the number to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) SetNumber(number string) {
	o.Number = number
}

// WithOrganization adds the organization to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) WithOrganization(organization string) *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithPage adds the page to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) WithPage(page *float64) *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) SetPage(page *float64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) WithPerPage(perPage *float64) *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get organizations organization databases database deploy requests number operations params
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) SetPerPage(perPage *float64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberOperationsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
