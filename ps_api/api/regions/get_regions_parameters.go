// Code generated by go-swagger; DO NOT EDIT.

package regions

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

// NewGetRegionsParams creates a new GetRegionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetRegionsParams() *GetRegionsParams {
	return &GetRegionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetRegionsParamsWithTimeout creates a new GetRegionsParams object
// with the ability to set a timeout on a request.
func NewGetRegionsParamsWithTimeout(timeout time.Duration) *GetRegionsParams {
	return &GetRegionsParams{
		timeout: timeout,
	}
}

// NewGetRegionsParamsWithContext creates a new GetRegionsParams object
// with the ability to set a context for a request.
func NewGetRegionsParamsWithContext(ctx context.Context) *GetRegionsParams {
	return &GetRegionsParams{
		Context: ctx,
	}
}

// NewGetRegionsParamsWithHTTPClient creates a new GetRegionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetRegionsParamsWithHTTPClient(client *http.Client) *GetRegionsParams {
	return &GetRegionsParams{
		HTTPClient: client,
	}
}

/*
GetRegionsParams contains all the parameters to send to the API endpoint

	for the get regions operation.

	Typically these are written to a http.Request.
*/
type GetRegionsParams struct {

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

// WithDefaults hydrates default values in the get regions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegionsParams) WithDefaults() *GetRegionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get regions params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetRegionsParams) SetDefaults() {
	var (
		pageDefault = float64(1)

		perPageDefault = float64(25)
	)

	val := GetRegionsParams{
		Page:    &pageDefault,
		PerPage: &perPageDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the get regions params
func (o *GetRegionsParams) WithTimeout(timeout time.Duration) *GetRegionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get regions params
func (o *GetRegionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get regions params
func (o *GetRegionsParams) WithContext(ctx context.Context) *GetRegionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get regions params
func (o *GetRegionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get regions params
func (o *GetRegionsParams) WithHTTPClient(client *http.Client) *GetRegionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get regions params
func (o *GetRegionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPage adds the page to the get regions params
func (o *GetRegionsParams) WithPage(page *float64) *GetRegionsParams {
	o.SetPage(page)
	return o
}

// SetPage adds the page to the get regions params
func (o *GetRegionsParams) SetPage(page *float64) {
	o.Page = page
}

// WithPerPage adds the perPage to the get regions params
func (o *GetRegionsParams) WithPerPage(perPage *float64) *GetRegionsParams {
	o.SetPerPage(perPage)
	return o
}

// SetPerPage adds the perPage to the get regions params
func (o *GetRegionsParams) SetPerPage(perPage *float64) {
	o.PerPage = perPage
}

// WriteToRequest writes these params to a swagger request
func (o *GetRegionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
