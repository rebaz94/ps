// Code generated by go-swagger; DO NOT EDIT.

package o_auth_applications

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDReader is a Reader for the DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenID structure.
type DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /organizations/{organization}/oauth-applications/{application_id}/tokens/{token_id}] delete_organizations_organization_oauth-applications_application_id_tokens_token_id", response, response.Code())
	}
}

// NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent creates a DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent with default headers values
func NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent() *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent {
	return &DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent{}
}

/*
DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent describes a response with status code 204, with default header values.

Deletes an OAuth application's OAuth token
*/
type DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent struct {
}

// IsSuccess returns true when this delete organizations organization oauth applications application Id tokens token Id no content response has a 2xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete organizations organization oauth applications application Id tokens token Id no content response has a 3xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organizations organization oauth applications application Id tokens token Id no content response has a 4xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organizations organization oauth applications application Id tokens token Id no content response has a 5xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organizations organization oauth applications application Id tokens token Id no content response a status code equal to that given
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete organizations organization oauth applications application Id tokens token Id no content response
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent) Code() int {
	return 204
}

func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/oauth-applications/{application_id}/tokens/{token_id}][%d] deleteOrganizationsOrganizationOauthApplicationsApplicationIdTokensTokenIdNoContent ", 204)
}

func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/oauth-applications/{application_id}/tokens/{token_id}][%d] deleteOrganizationsOrganizationOauthApplicationsApplicationIdTokensTokenIdNoContent ", 204)
}

func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized creates a DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized with default headers values
func NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized() *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized {
	return &DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized{}
}

/*
DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized struct {
}

// IsSuccess returns true when this delete organizations organization oauth applications application Id tokens token Id unauthorized response has a 2xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete organizations organization oauth applications application Id tokens token Id unauthorized response has a 3xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organizations organization oauth applications application Id tokens token Id unauthorized response has a 4xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete organizations organization oauth applications application Id tokens token Id unauthorized response has a 5xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organizations organization oauth applications application Id tokens token Id unauthorized response a status code equal to that given
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete organizations organization oauth applications application Id tokens token Id unauthorized response
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized) Code() int {
	return 401
}

func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/oauth-applications/{application_id}/tokens/{token_id}][%d] deleteOrganizationsOrganizationOauthApplicationsApplicationIdTokensTokenIdUnauthorized ", 401)
}

func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/oauth-applications/{application_id}/tokens/{token_id}][%d] deleteOrganizationsOrganizationOauthApplicationsApplicationIdTokensTokenIdUnauthorized ", 401)
}

func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden creates a DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden with default headers values
func NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden() *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden {
	return &DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden{}
}

/*
DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden struct {
}

// IsSuccess returns true when this delete organizations organization oauth applications application Id tokens token Id forbidden response has a 2xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete organizations organization oauth applications application Id tokens token Id forbidden response has a 3xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organizations organization oauth applications application Id tokens token Id forbidden response has a 4xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete organizations organization oauth applications application Id tokens token Id forbidden response has a 5xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organizations organization oauth applications application Id tokens token Id forbidden response a status code equal to that given
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete organizations organization oauth applications application Id tokens token Id forbidden response
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden) Code() int {
	return 403
}

func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/oauth-applications/{application_id}/tokens/{token_id}][%d] deleteOrganizationsOrganizationOauthApplicationsApplicationIdTokensTokenIdForbidden ", 403)
}

func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/oauth-applications/{application_id}/tokens/{token_id}][%d] deleteOrganizationsOrganizationOauthApplicationsApplicationIdTokensTokenIdForbidden ", 403)
}

func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound creates a DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound with default headers values
func NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound() *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound {
	return &DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound{}
}

/*
DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound struct {
}

// IsSuccess returns true when this delete organizations organization oauth applications application Id tokens token Id not found response has a 2xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete organizations organization oauth applications application Id tokens token Id not found response has a 3xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organizations organization oauth applications application Id tokens token Id not found response has a 4xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete organizations organization oauth applications application Id tokens token Id not found response has a 5xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organizations organization oauth applications application Id tokens token Id not found response a status code equal to that given
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete organizations organization oauth applications application Id tokens token Id not found response
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound) Code() int {
	return 404
}

func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/oauth-applications/{application_id}/tokens/{token_id}][%d] deleteOrganizationsOrganizationOauthApplicationsApplicationIdTokensTokenIdNotFound ", 404)
}

func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/oauth-applications/{application_id}/tokens/{token_id}][%d] deleteOrganizationsOrganizationOauthApplicationsApplicationIdTokensTokenIdNotFound ", 404)
}

func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError creates a DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError with default headers values
func NewDeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError() *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError {
	return &DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError{}
}

/*
DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError struct {
}

// IsSuccess returns true when this delete organizations organization oauth applications application Id tokens token Id internal server error response has a 2xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete organizations organization oauth applications application Id tokens token Id internal server error response has a 3xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organizations organization oauth applications application Id tokens token Id internal server error response has a 4xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organizations organization oauth applications application Id tokens token Id internal server error response has a 5xx status code
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete organizations organization oauth applications application Id tokens token Id internal server error response a status code equal to that given
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete organizations organization oauth applications application Id tokens token Id internal server error response
func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError) Code() int {
	return 500
}

func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/oauth-applications/{application_id}/tokens/{token_id}][%d] deleteOrganizationsOrganizationOauthApplicationsApplicationIdTokensTokenIdInternalServerError ", 500)
}

func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/oauth-applications/{application_id}/tokens/{token_id}][%d] deleteOrganizationsOrganizationOauthApplicationsApplicationIdTokensTokenIdInternalServerError ", 500)
}

func (o *DeleteOrganizationsOrganizationOauthApplicationsApplicationIDTokensTokenIDInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
