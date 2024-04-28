// Code generated by go-swagger; DO NOT EDIT.

package databases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteOrganizationsOrganizationDatabasesNameReader is a Reader for the DeleteOrganizationsOrganizationDatabasesName structure.
type DeleteOrganizationsOrganizationDatabasesNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationsOrganizationDatabasesNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrganizationsOrganizationDatabasesNameNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDeleteOrganizationsOrganizationDatabasesNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteOrganizationsOrganizationDatabasesNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteOrganizationsOrganizationDatabasesNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteOrganizationsOrganizationDatabasesNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /organizations/{organization}/databases/{name}] delete_organizations_organization_databases_name", response, response.Code())
	}
}

// NewDeleteOrganizationsOrganizationDatabasesNameNoContent creates a DeleteOrganizationsOrganizationDatabasesNameNoContent with default headers values
func NewDeleteOrganizationsOrganizationDatabasesNameNoContent() *DeleteOrganizationsOrganizationDatabasesNameNoContent {
	return &DeleteOrganizationsOrganizationDatabasesNameNoContent{}
}

/*
DeleteOrganizationsOrganizationDatabasesNameNoContent describes a response with status code 204, with default header values.

Deletes a database
*/
type DeleteOrganizationsOrganizationDatabasesNameNoContent struct {
}

// IsSuccess returns true when this delete organizations organization databases name no content response has a 2xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete organizations organization databases name no content response has a 3xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organizations organization databases name no content response has a 4xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organizations organization databases name no content response has a 5xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organizations organization databases name no content response a status code equal to that given
func (o *DeleteOrganizationsOrganizationDatabasesNameNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete organizations organization databases name no content response
func (o *DeleteOrganizationsOrganizationDatabasesNameNoContent) Code() int {
	return 204
}

func (o *DeleteOrganizationsOrganizationDatabasesNameNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/databases/{name}][%d] deleteOrganizationsOrganizationDatabasesNameNoContent ", 204)
}

func (o *DeleteOrganizationsOrganizationDatabasesNameNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/databases/{name}][%d] deleteOrganizationsOrganizationDatabasesNameNoContent ", 204)
}

func (o *DeleteOrganizationsOrganizationDatabasesNameNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationsOrganizationDatabasesNameUnauthorized creates a DeleteOrganizationsOrganizationDatabasesNameUnauthorized with default headers values
func NewDeleteOrganizationsOrganizationDatabasesNameUnauthorized() *DeleteOrganizationsOrganizationDatabasesNameUnauthorized {
	return &DeleteOrganizationsOrganizationDatabasesNameUnauthorized{}
}

/*
DeleteOrganizationsOrganizationDatabasesNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type DeleteOrganizationsOrganizationDatabasesNameUnauthorized struct {
}

// IsSuccess returns true when this delete organizations organization databases name unauthorized response has a 2xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete organizations organization databases name unauthorized response has a 3xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organizations organization databases name unauthorized response has a 4xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete organizations organization databases name unauthorized response has a 5xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organizations organization databases name unauthorized response a status code equal to that given
func (o *DeleteOrganizationsOrganizationDatabasesNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the delete organizations organization databases name unauthorized response
func (o *DeleteOrganizationsOrganizationDatabasesNameUnauthorized) Code() int {
	return 401
}

func (o *DeleteOrganizationsOrganizationDatabasesNameUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/databases/{name}][%d] deleteOrganizationsOrganizationDatabasesNameUnauthorized ", 401)
}

func (o *DeleteOrganizationsOrganizationDatabasesNameUnauthorized) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/databases/{name}][%d] deleteOrganizationsOrganizationDatabasesNameUnauthorized ", 401)
}

func (o *DeleteOrganizationsOrganizationDatabasesNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationsOrganizationDatabasesNameForbidden creates a DeleteOrganizationsOrganizationDatabasesNameForbidden with default headers values
func NewDeleteOrganizationsOrganizationDatabasesNameForbidden() *DeleteOrganizationsOrganizationDatabasesNameForbidden {
	return &DeleteOrganizationsOrganizationDatabasesNameForbidden{}
}

/*
DeleteOrganizationsOrganizationDatabasesNameForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type DeleteOrganizationsOrganizationDatabasesNameForbidden struct {
}

// IsSuccess returns true when this delete organizations organization databases name forbidden response has a 2xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete organizations organization databases name forbidden response has a 3xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organizations organization databases name forbidden response has a 4xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete organizations organization databases name forbidden response has a 5xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organizations organization databases name forbidden response a status code equal to that given
func (o *DeleteOrganizationsOrganizationDatabasesNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the delete organizations organization databases name forbidden response
func (o *DeleteOrganizationsOrganizationDatabasesNameForbidden) Code() int {
	return 403
}

func (o *DeleteOrganizationsOrganizationDatabasesNameForbidden) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/databases/{name}][%d] deleteOrganizationsOrganizationDatabasesNameForbidden ", 403)
}

func (o *DeleteOrganizationsOrganizationDatabasesNameForbidden) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/databases/{name}][%d] deleteOrganizationsOrganizationDatabasesNameForbidden ", 403)
}

func (o *DeleteOrganizationsOrganizationDatabasesNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationsOrganizationDatabasesNameNotFound creates a DeleteOrganizationsOrganizationDatabasesNameNotFound with default headers values
func NewDeleteOrganizationsOrganizationDatabasesNameNotFound() *DeleteOrganizationsOrganizationDatabasesNameNotFound {
	return &DeleteOrganizationsOrganizationDatabasesNameNotFound{}
}

/*
DeleteOrganizationsOrganizationDatabasesNameNotFound describes a response with status code 404, with default header values.

Not Found
*/
type DeleteOrganizationsOrganizationDatabasesNameNotFound struct {
}

// IsSuccess returns true when this delete organizations organization databases name not found response has a 2xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete organizations organization databases name not found response has a 3xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organizations organization databases name not found response has a 4xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this delete organizations organization databases name not found response has a 5xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organizations organization databases name not found response a status code equal to that given
func (o *DeleteOrganizationsOrganizationDatabasesNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the delete organizations organization databases name not found response
func (o *DeleteOrganizationsOrganizationDatabasesNameNotFound) Code() int {
	return 404
}

func (o *DeleteOrganizationsOrganizationDatabasesNameNotFound) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/databases/{name}][%d] deleteOrganizationsOrganizationDatabasesNameNotFound ", 404)
}

func (o *DeleteOrganizationsOrganizationDatabasesNameNotFound) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/databases/{name}][%d] deleteOrganizationsOrganizationDatabasesNameNotFound ", 404)
}

func (o *DeleteOrganizationsOrganizationDatabasesNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteOrganizationsOrganizationDatabasesNameInternalServerError creates a DeleteOrganizationsOrganizationDatabasesNameInternalServerError with default headers values
func NewDeleteOrganizationsOrganizationDatabasesNameInternalServerError() *DeleteOrganizationsOrganizationDatabasesNameInternalServerError {
	return &DeleteOrganizationsOrganizationDatabasesNameInternalServerError{}
}

/*
DeleteOrganizationsOrganizationDatabasesNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type DeleteOrganizationsOrganizationDatabasesNameInternalServerError struct {
}

// IsSuccess returns true when this delete organizations organization databases name internal server error response has a 2xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this delete organizations organization databases name internal server error response has a 3xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organizations organization databases name internal server error response has a 4xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organizations organization databases name internal server error response has a 5xx status code
func (o *DeleteOrganizationsOrganizationDatabasesNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this delete organizations organization databases name internal server error response a status code equal to that given
func (o *DeleteOrganizationsOrganizationDatabasesNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the delete organizations organization databases name internal server error response
func (o *DeleteOrganizationsOrganizationDatabasesNameInternalServerError) Code() int {
	return 500
}

func (o *DeleteOrganizationsOrganizationDatabasesNameInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/databases/{name}][%d] deleteOrganizationsOrganizationDatabasesNameInternalServerError ", 500)
}

func (o *DeleteOrganizationsOrganizationDatabasesNameInternalServerError) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organization}/databases/{name}][%d] deleteOrganizationsOrganizationDatabasesNameInternalServerError ", 500)
}

func (o *DeleteOrganizationsOrganizationDatabasesNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}