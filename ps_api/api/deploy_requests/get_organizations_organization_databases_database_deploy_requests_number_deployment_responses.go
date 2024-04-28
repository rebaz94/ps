// Code generated by go-swagger; DO NOT EDIT.

package deploy_requests

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentReader is a Reader for the GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeployment structure.
type GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organization}/databases/{database}/deploy-requests/{number}/deployment] get_organizations_organization_databases_database_deploy-requests_number_deployment", response, response.Code())
	}
}

// NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK creates a GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK with default headers values
func NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK() *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK {
	return &GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK{}
}

/*
GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK describes a response with status code 200, with default header values.

Returns the deployment for a deploy request
*/
type GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK struct {
	Payload *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody
}

// IsSuccess returns true when this get organizations organization databases database deploy requests number deployment o k response has a 2xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organizations organization databases database deploy requests number deployment o k response has a 3xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations organization databases database deploy requests number deployment o k response has a 4xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organizations organization databases database deploy requests number deployment o k response has a 5xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations organization databases database deploy requests number deployment o k response a status code equal to that given
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organizations organization databases database deploy requests number deployment o k response
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK) Code() int {
	return 200
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{database}/deploy-requests/{number}/deployment][%d] getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{database}/deploy-requests/{number}/deployment][%d] getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK) GetPayload() *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody {
	return o.Payload
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized creates a GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized with default headers values
func NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized() *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized {
	return &GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized{}
}

/*
GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized struct {
}

// IsSuccess returns true when this get organizations organization databases database deploy requests number deployment unauthorized response has a 2xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations organization databases database deploy requests number deployment unauthorized response has a 3xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations organization databases database deploy requests number deployment unauthorized response has a 4xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations organization databases database deploy requests number deployment unauthorized response has a 5xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations organization databases database deploy requests number deployment unauthorized response a status code equal to that given
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get organizations organization databases database deploy requests number deployment unauthorized response
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized) Code() int {
	return 401
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{database}/deploy-requests/{number}/deployment][%d] getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized ", 401)
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized) String() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{database}/deploy-requests/{number}/deployment][%d] getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized ", 401)
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden creates a GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden with default headers values
func NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden() *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden {
	return &GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden{}
}

/*
GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden struct {
}

// IsSuccess returns true when this get organizations organization databases database deploy requests number deployment forbidden response has a 2xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations organization databases database deploy requests number deployment forbidden response has a 3xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations organization databases database deploy requests number deployment forbidden response has a 4xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations organization databases database deploy requests number deployment forbidden response has a 5xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations organization databases database deploy requests number deployment forbidden response a status code equal to that given
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get organizations organization databases database deploy requests number deployment forbidden response
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden) Code() int {
	return 403
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{database}/deploy-requests/{number}/deployment][%d] getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden ", 403)
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden) String() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{database}/deploy-requests/{number}/deployment][%d] getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden ", 403)
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound creates a GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound with default headers values
func NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound() *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound {
	return &GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound{}
}

/*
GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound struct {
}

// IsSuccess returns true when this get organizations organization databases database deploy requests number deployment not found response has a 2xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations organization databases database deploy requests number deployment not found response has a 3xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations organization databases database deploy requests number deployment not found response has a 4xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations organization databases database deploy requests number deployment not found response has a 5xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations organization databases database deploy requests number deployment not found response a status code equal to that given
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get organizations organization databases database deploy requests number deployment not found response
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound) Code() int {
	return 404
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{database}/deploy-requests/{number}/deployment][%d] getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound ", 404)
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound) String() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{database}/deploy-requests/{number}/deployment][%d] getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound ", 404)
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError creates a GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError with default headers values
func NewGetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError() *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError {
	return &GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError{}
}

/*
GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError struct {
}

// IsSuccess returns true when this get organizations organization databases database deploy requests number deployment internal server error response has a 2xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations organization databases database deploy requests number deployment internal server error response has a 3xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations organization databases database deploy requests number deployment internal server error response has a 4xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organizations organization databases database deploy requests number deployment internal server error response has a 5xx status code
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get organizations organization databases database deploy requests number deployment internal server error response a status code equal to that given
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get organizations organization databases database deploy requests number deployment internal server error response
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError) Code() int {
	return 500
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{database}/deploy-requests/{number}/deployment][%d] getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError ", 500)
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError) String() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{database}/deploy-requests/{number}/deployment][%d] getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError ", 500)
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody get organizations organization databases database deploy requests number deployment o k body
swagger:model GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody
*/
type GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody struct {

	// Whether or not to automatically cutover once deployment is finished
	// Required: true
	AutoCutover *bool `json:"auto_cutover"`

	// When the deployment was created
	// Required: true
	CreatedAt *string `json:"created_at"`

	// When the cutover for the deployment was initiated
	CutoverAt string `json:"cutover_at,omitempty"`

	// Whether or not the deployment cutover will expire soon and be cancelled
	// Required: true
	CutoverExpiring *bool `json:"cutover_expiring"`

	// Deploy check errors for the deployment
	DeployCheckErrors string `json:"deploy_check_errors,omitempty"`

	// When the deployment was finished
	FinishedAt string `json:"finished_at,omitempty"`

	// The ID for a deployment
	// Required: true
	ID *string `json:"id"`

	// When the deployment was queued
	QueuedAt string `json:"queued_at,omitempty"`

	// When the deployment was ready for cutover
	ReadyToCutoverAt string `json:"ready_to_cutover_at,omitempty"`

	// When the deployment was started
	StartedAt string `json:"started_at,omitempty"`

	// The state the deployment is in
	// Required: true
	// Enum: [pending ready no_changes queued submitting in_progress pending_cutover in_progress_vschema in_progress_cancel in_progress_cutover complete complete_cancel complete_error complete_pending_revert in_progress_revert complete_revert complete_revert_error cancelled error]
	State *string `json:"state"`

	// When the deployment was submitted
	// Required: true
	SubmittedAt *string `json:"submitted_at"`

	// When the deployment was last updated
	// Required: true
	UpdatedAt *string `json:"updated_at"`
}

// Validate validates this get organizations organization databases database deploy requests number deployment o k body
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAutoCutover(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCutoverExpiring(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubmittedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody) validateAutoCutover(formats strfmt.Registry) error {

	if err := validate.Required("getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK"+"."+"auto_cutover", "body", o.AutoCutover); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK"+"."+"created_at", "body", o.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody) validateCutoverExpiring(formats strfmt.Registry) error {

	if err := validate.Required("getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK"+"."+"cutover_expiring", "body", o.CutoverExpiring); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

var getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["pending","ready","no_changes","queued","submitting","in_progress","pending_cutover","in_progress_vschema","in_progress_cancel","in_progress_cutover","complete","complete_cancel","complete_error","complete_pending_revert","in_progress_revert","complete_revert","complete_revert_error","cancelled","error"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyTypeStatePropEnum = append(getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyTypeStatePropEnum, v)
	}
}

const (

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStatePending captures enum value "pending"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStatePending string = "pending"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateReady captures enum value "ready"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateReady string = "ready"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateNoChanges captures enum value "no_changes"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateNoChanges string = "no_changes"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateQueued captures enum value "queued"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateQueued string = "queued"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateSubmitting captures enum value "submitting"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateSubmitting string = "submitting"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateInProgress captures enum value "in_progress"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateInProgress string = "in_progress"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStatePendingCutover captures enum value "pending_cutover"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStatePendingCutover string = "pending_cutover"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateInProgressVschema captures enum value "in_progress_vschema"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateInProgressVschema string = "in_progress_vschema"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateInProgressCancel captures enum value "in_progress_cancel"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateInProgressCancel string = "in_progress_cancel"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateInProgressCutover captures enum value "in_progress_cutover"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateInProgressCutover string = "in_progress_cutover"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateComplete captures enum value "complete"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateComplete string = "complete"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateCompleteCancel captures enum value "complete_cancel"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateCompleteCancel string = "complete_cancel"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateCompleteError captures enum value "complete_error"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateCompleteError string = "complete_error"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateCompletePendingRevert captures enum value "complete_pending_revert"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateCompletePendingRevert string = "complete_pending_revert"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateInProgressRevert captures enum value "in_progress_revert"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateInProgressRevert string = "in_progress_revert"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateCompleteRevert captures enum value "complete_revert"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateCompleteRevert string = "complete_revert"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateCompleteRevertError captures enum value "complete_revert_error"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateCompleteRevertError string = "complete_revert_error"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateCancelled captures enum value "cancelled"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateCancelled string = "cancelled"

	// GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateError captures enum value "error"
	GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyStateError string = "error"
)

// prop value enum
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBodyTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody) validateState(formats strfmt.Registry) error {

	if err := validate.Required("getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK"+"."+"state", "body", o.State); err != nil {
		return err
	}

	// value enum
	if err := o.validateStateEnum("getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK"+"."+"state", "body", *o.State); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody) validateSubmittedAt(formats strfmt.Registry) error {

	if err := validate.Required("getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK"+"."+"submitted_at", "body", o.SubmittedAt); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("getOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOK"+"."+"updated_at", "body", o.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get organizations organization databases database deploy requests number deployment o k body based on context it is used
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrganizationsOrganizationDatabasesDatabaseDeployRequestsNumberDeploymentOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
