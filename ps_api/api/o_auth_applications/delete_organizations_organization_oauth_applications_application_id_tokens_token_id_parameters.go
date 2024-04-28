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

// NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams creates a new DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams() *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	return &DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParamsWithTimeout creates a new DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams object
// with the ability to set a timeout on a request.
func NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParamsWithTimeout(timeout time.Duration) *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	return &DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams{
		timeout: timeout,
	}
}

// NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParamsWithContext creates a new DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams object
// with the ability to set a context for a request.
func NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParamsWithContext(ctx context.Context) *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	return &DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams{
		Context: ctx,
	}
}

// NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParamsWithHTTPClient creates a new DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams object
// with the ability to set a custom HTTPClient for a request.
func NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParamsWithHTTPClient(client *http.Client) *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	return &DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams{
		HTTPClient: client,
	}
}

/*
DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams contains all the parameters to send to the API endpoint

	for the delete organizations organization oauth applications application id tokens token id operation.

	Typically these are written to a http.Request.
*/
type DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams struct {

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

// WithDefaults hydrates default values in the delete organizations organization oauth applications application id tokens token id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WithDefaults() *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the delete organizations organization oauth applications application id tokens token id params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the delete organizations organization oauth applications application id tokens token id params
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WithTimeout(timeout time.Duration) *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete organizations organization oauth applications application id tokens token id params
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete organizations organization oauth applications application id tokens token id params
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WithContext(ctx context.Context) *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete organizations organization oauth applications application id tokens token id params
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete organizations organization oauth applications application id tokens token id params
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WithHTTPClient(client *http.Client) *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete organizations organization oauth applications application id tokens token id params
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationID adds the applicationID to the delete organizations organization oauth applications application id tokens token id params
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WithApplicationID(applicationID string) *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	o.SetApplicationID(applicationID)
	return o
}

// SetApplicationID adds the applicationId to the delete organizations organization oauth applications application id tokens token id params
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) SetApplicationID(applicationID string) {
	o.ApplicationID = applicationID
}

// WithOrganization adds the organization to the delete organizations organization oauth applications application id tokens token id params
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WithOrganization(organization string) *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	o.SetOrganization(organization)
	return o
}

// SetOrganization adds the organization to the delete organizations organization oauth applications application id tokens token id params
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) SetOrganization(organization string) {
	o.Organization = organization
}

// WithTokenID adds the tokenID to the delete organizations organization oauth applications application id tokens token id params
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WithTokenID(tokenID string) *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams {
	o.SetTokenID(tokenID)
	return o
}

// SetTokenID adds the tokenId to the delete organizations organization oauth applications application id tokens token id params
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) SetTokenID(tokenID string) {
	o.TokenID = tokenID
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
