// Code generated by go-swagger; DO NOT EDIT.

package o_auth_applications

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

// NewGetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams creates a new GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams() *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	return &GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParamsWithTimeout creates a new GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams object
// with the ability to set a timeout on a request.
func NewGetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParamsWithTimeout(timeout time.Duration) *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	return &GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams{
		timeout: timeout,
	}
}

// NewGetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParamsWithContext creates a new GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams object
// with the ability to set a context for a request.
func NewGetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParamsWithContext(ctx context.Context) *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	return &GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams{
		Context: ctx,
	}
}

// NewGetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParamsWithHTTPClient creates a new GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParamsWithHTTPClient(client *http.Client) *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	return &GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams{
		HTTPClient: client,
	}
}

/*
GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams contains all the parameters to send to the API endpoint

	for the get organizations organization oauth applications application id tokens token id operation.

	Typically these are written to a http.Request.
*/
type GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams struct {

	/* ApplicationID.

	   The ID of the OAuth application
	*/
	ApplicationID string

	/* Organization.

	   The name of the organization the OAuth application belongs to
	*/
	Organization string

	/* TokenID.

	   The ID of the OAuth application token
	*/
	TokenID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get organizations organization oauth applications application id tokens token id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WithDefaults() *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get organizations organization oauth applications application id tokens token id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get organizations organization oauth applications application id tokens token id params
func (o *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WithTimeout(timeout time.Duration) *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get organizations organization oauth applications application id tokens token id params
func (o *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get organizations organization oauth applications application id tokens token id params
func (o *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WithContext(ctx context.Context) *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get organizations organization oauth applications application id tokens token id params
func (o *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get organizations organization oauth applications application id tokens token id params
func (o *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WithHTTPClient(client *http.Client) *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get organizations organization oauth applications application id tokens token id params
func (o *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the get organizations organization oauth applications application id tokens token id params
func (o *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WithApplicationID(applicationID string) *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the get organizations organization oauth applications application id tokens token id params
func (o *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) SetApplicationID(applicationID string) {
	o.ApplicationID = applicationID
}

// WithOrganization adds the organization to the get organizations organization oauth applications application id tokens token id params
func (o *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WithOrganization(organization string) *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the get organizations organization oauth applications application id tokens token id params
func (o *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithTokenID adds the tokenID to the get organizations organization oauth applications application id tokens token id params
func (o *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WithTokenID(tokenID string) *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	o.SetTokenID(tokenID)
	return o
}

// SetTokenID adds the tokenId to the get organizations organization oauth applications application id tokens token id params
func (o *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) SetTokenID(tokenID string) {
	o.TokenID = tokenID
}

// WriteToRequest writes these params to a swagger request
func (o *GetOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application_id
	if err := r.SetPathParam("application_id", o.ApplicationID); err != nil {
		return err
	}

	// path param organization
	if err := r.SetPathParam("organization", o.Organization); err != nil {
		return err
	}

	// path param token_id
	if err := r.SetPathParam("token_id", o.TokenID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}