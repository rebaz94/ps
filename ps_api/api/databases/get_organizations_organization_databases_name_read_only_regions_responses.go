// Code generated by go-swagger; DO NOT EDIT.

package databases

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsReader is a Reader for the GetOrganizationsOrganizationDatabasesNameReadOnlyRegions structure.
type GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organization}/databases/{name}/read-only-regions] get_organizations_organization_databases_name_read-only-regions", response, response.Code())
	}
}

// NewGetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK creates a GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK with default headers values
func NewGetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK() *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK {
	return &GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK{}
}

/*
GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK describes a response with status code 200, with default header values.

List of the database's read-only regions
*/
type GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK struct {
	Payload *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBody
}

// IsSuccess returns true when this get organizations organization databases name read only regions o k response has a 2xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organizations organization databases name read only regions o k response has a 3xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations organization databases name read only regions o k response has a 4xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organizations organization databases name read only regions o k response has a 5xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations organization databases name read only regions o k response a status code equal to that given
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organizations organization databases name read only regions o k response
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK) Code() int {
	return 200
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{name}/read-only-regions][%d] getOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{name}/read-only-regions][%d] getOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK) GetPayload() *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBody {
	return o.Payload
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized creates a GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized with default headers values
func NewGetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized() *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized {
	return &GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized{}
}

/*
GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized struct {
}

// IsSuccess returns true when this get organizations organization databases name read only regions unauthorized response has a 2xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations organization databases name read only regions unauthorized response has a 3xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations organization databases name read only regions unauthorized response has a 4xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations organization databases name read only regions unauthorized response has a 5xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations organization databases name read only regions unauthorized response a status code equal to that given
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get organizations organization databases name read only regions unauthorized response
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized) Code() int {
	return 401
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{name}/read-only-regions][%d] getOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized ", 401)
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{name}/read-only-regions][%d] getOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized ", 401)
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden creates a GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden with default headers values
func NewGetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden() *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden {
	return &GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden{}
}

/*
GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden struct {
}

// IsSuccess returns true when this get organizations organization databases name read only regions forbidden response has a 2xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations organization databases name read only regions forbidden response has a 3xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations organization databases name read only regions forbidden response has a 4xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations organization databases name read only regions forbidden response has a 5xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations organization databases name read only regions forbidden response a status code equal to that given
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get organizations organization databases name read only regions forbidden response
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden) Code() int {
	return 403
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{name}/read-only-regions][%d] getOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden ", 403)
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden) String() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{name}/read-only-regions][%d] getOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden ", 403)
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound creates a GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound with default headers values
func NewGetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound() *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound {
	return &GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound{}
}

/*
GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound struct {
}

// IsSuccess returns true when this get organizations organization databases name read only regions not found response has a 2xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations organization databases name read only regions not found response has a 3xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations organization databases name read only regions not found response has a 4xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations organization databases name read only regions not found response has a 5xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations organization databases name read only regions not found response a status code equal to that given
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get organizations organization databases name read only regions not found response
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound) Code() int {
	return 404
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{name}/read-only-regions][%d] getOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound ", 404)
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound) String() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{name}/read-only-regions][%d] getOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound ", 404)
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError creates a GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError with default headers values
func NewGetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError() *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError {
	return &GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError{}
}

/*
GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError struct {
}

// IsSuccess returns true when this get organizations organization databases name read only regions internal server error response has a 2xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations organization databases name read only regions internal server error response has a 3xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations organization databases name read only regions internal server error response has a 4xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organizations organization databases name read only regions internal server error response has a 5xx status code
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get organizations organization databases name read only regions internal server error response a status code equal to that given
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get organizations organization databases name read only regions internal server error response
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError) Code() int {
	return 500
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{name}/read-only-regions][%d] getOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError ", 500)
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /organizations/{organization}/databases/{name}/read-only-regions][%d] getOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError ", 500)
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBody get organizations organization databases name read only regions o k body
swagger:model GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBody
*/
type GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBody struct {

	// data
	// Required: true
	Data []*GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0 `json:"data"`
}

// Validate validates this get organizations organization databases name read only regions o k body
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get organizations organization databases name read only regions o k body based on the context it is used
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getOrganizationsOrganizationDatabasesNameReadOnlyRegionsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0 get organizations organization databases name read only regions o k body data items0
swagger:model GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0
*/
type GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0 struct {

	// actor
	// Required: true
	Actor *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Actor `json:"actor"`

	// When the read-only region was created
	// Required: true
	CreatedAt *string `json:"created_at"`

	// The name of the read-only region
	// Required: true
	DisplayName *string `json:"display_name"`

	// The ID of the read-only region
	// Required: true
	ID *string `json:"id"`

	// Whether or not the read-only region is ready to serve queries
	// Required: true
	Ready *bool `json:"ready"`

	// When the read-only region was ready to serve queries
	// Required: true
	ReadyAt *string `json:"ready_at"`

	// region
	// Required: true
	Region *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region `json:"region"`

	// When the read-only region was last updated
	// Required: true
	UpdatedAt *string `json:"updated_at"`
}

// Validate validates this get organizations organization databases name read only regions o k body data items0
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActor(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReady(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReadyAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRegion(formats); err != nil {
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

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0) validateActor(formats strfmt.Registry) error {

	if err := validate.Required("actor", "body", o.Actor); err != nil {
		return err
	}

	if o.Actor != nil {
		if err := o.Actor.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actor")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", o.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("display_name", "body", o.DisplayName); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0) validateReady(formats strfmt.Registry) error {

	if err := validate.Required("ready", "body", o.Ready); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0) validateReadyAt(formats strfmt.Registry) error {

	if err := validate.Required("ready_at", "body", o.ReadyAt); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", o.Region); err != nil {
		return err
	}

	if o.Region != nil {
		if err := o.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", o.UpdatedAt); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get organizations organization databases name read only regions o k body data items0 based on the context it is used
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateActor(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0) contextValidateActor(ctx context.Context, formats strfmt.Registry) error {

	if o.Actor != nil {

		if err := o.Actor.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("actor")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("actor")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if o.Region != nil {

		if err := o.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Actor get organizations organization databases name read only regions o k body data items0 actor
swagger:model GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Actor
*/
type GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Actor struct {

	// The URL of the actor's avatar
	// Required: true
	AvatarURL *string `json:"avatar_url"`

	// The name of the actor
	// Required: true
	DisplayName *string `json:"display_name"`

	// The ID of the actor
	// Required: true
	ID *string `json:"id"`
}

// Validate validates this get organizations organization databases name read only regions o k body data items0 actor
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Actor) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAvatarURL(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Actor) validateAvatarURL(formats strfmt.Registry) error {

	if err := validate.Required("actor"+"."+"avatar_url", "body", o.AvatarURL); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Actor) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("actor"+"."+"display_name", "body", o.DisplayName); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Actor) validateID(formats strfmt.Registry) error {

	if err := validate.Required("actor"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get organizations organization databases name read only regions o k body data items0 actor based on context it is used
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Actor) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Actor) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Actor) UnmarshalBinary(b []byte) error {
	var res GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Actor
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region get organizations organization databases name read only regions o k body data items0 region
swagger:model GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region
*/
type GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region struct {

	// Name of the region
	// Required: true
	DisplayName *string `json:"display_name"`

	// Whether or not the region is currently active
	// Required: true
	Enabled *bool `json:"enabled"`

	// The ID of the region
	// Required: true
	ID *string `json:"id"`

	// Location of the region
	// Required: true
	Location *string `json:"location"`

	// Provider for the region (ex. AWS)
	// Required: true
	Provider *string `json:"provider"`

	// Public IP addresses for the region
	// Required: true
	PublicIPAddresses []string `json:"public_ip_addresses"`

	// The slug of the region
	// Required: true
	Slug *string `json:"slug"`
}

// Validate validates this get organizations organization databases name read only regions o k body data items0 region
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProvider(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePublicIPAddresses(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSlug(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("region"+"."+"display_name", "body", o.DisplayName); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("region"+"."+"enabled", "body", o.Enabled); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region) validateID(formats strfmt.Registry) error {

	if err := validate.Required("region"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("region"+"."+"location", "body", o.Location); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("region"+"."+"provider", "body", o.Provider); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region) validatePublicIPAddresses(formats strfmt.Registry) error {

	if err := validate.Required("region"+"."+"public_ip_addresses", "body", o.PublicIPAddresses); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("region"+"."+"slug", "body", o.Slug); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get organizations organization databases name read only regions o k body data items0 region based on context it is used
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region) UnmarshalBinary(b []byte) error {
	var res GetOrganizationsOrganizationDatabasesNameReadOnlyRegionsOKBodyDataItems0Region
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
