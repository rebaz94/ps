// Code generated by go-swagger; DO NOT EDIT.

package regions

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

// GetRegionsReader is a Reader for the GetRegions structure.
type GetRegionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRegionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetRegionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetRegionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetRegionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetRegionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetRegionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /regions] get_regions", response, response.Code())
	}
}

// NewGetRegionsOK creates a GetRegionsOK with default headers values
func NewGetRegionsOK() *GetRegionsOK {
	return &GetRegionsOK{}
}

/*
GetRegionsOK describes a response with status code 200, with default header values.

Returns the available public PlanetScale regions
*/
type GetRegionsOK struct {
	Payload *GetRegionsOKBody
}

// IsSuccess returns true when this get regions o k response has a 2xx status code
func (o *GetRegionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get regions o k response has a 3xx status code
func (o *GetRegionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get regions o k response has a 4xx status code
func (o *GetRegionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get regions o k response has a 5xx status code
func (o *GetRegionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get regions o k response a status code equal to that given
func (o *GetRegionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get regions o k response
func (o *GetRegionsOK) Code() int {
	return 200
}

func (o *GetRegionsOK) Error() string {
	return fmt.Sprintf("[GET /regions][%d] getRegionsOK  %+v", 200, o.Payload)
}

func (o *GetRegionsOK) String() string {
	return fmt.Sprintf("[GET /regions][%d] getRegionsOK  %+v", 200, o.Payload)
}

func (o *GetRegionsOK) GetPayload() *GetRegionsOKBody {
	return o.Payload
}

func (o *GetRegionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetRegionsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRegionsUnauthorized creates a GetRegionsUnauthorized with default headers values
func NewGetRegionsUnauthorized() *GetRegionsUnauthorized {
	return &GetRegionsUnauthorized{}
}

/*
GetRegionsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetRegionsUnauthorized struct {
}

// IsSuccess returns true when this get regions unauthorized response has a 2xx status code
func (o *GetRegionsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get regions unauthorized response has a 3xx status code
func (o *GetRegionsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get regions unauthorized response has a 4xx status code
func (o *GetRegionsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get regions unauthorized response has a 5xx status code
func (o *GetRegionsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get regions unauthorized response a status code equal to that given
func (o *GetRegionsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get regions unauthorized response
func (o *GetRegionsUnauthorized) Code() int {
	return 401
}

func (o *GetRegionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /regions][%d] getRegionsUnauthorized ", 401)
}

func (o *GetRegionsUnauthorized) String() string {
	return fmt.Sprintf("[GET /regions][%d] getRegionsUnauthorized ", 401)
}

func (o *GetRegionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegionsForbidden creates a GetRegionsForbidden with default headers values
func NewGetRegionsForbidden() *GetRegionsForbidden {
	return &GetRegionsForbidden{}
}

/*
GetRegionsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetRegionsForbidden struct {
}

// IsSuccess returns true when this get regions forbidden response has a 2xx status code
func (o *GetRegionsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get regions forbidden response has a 3xx status code
func (o *GetRegionsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get regions forbidden response has a 4xx status code
func (o *GetRegionsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get regions forbidden response has a 5xx status code
func (o *GetRegionsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get regions forbidden response a status code equal to that given
func (o *GetRegionsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get regions forbidden response
func (o *GetRegionsForbidden) Code() int {
	return 403
}

func (o *GetRegionsForbidden) Error() string {
	return fmt.Sprintf("[GET /regions][%d] getRegionsForbidden ", 403)
}

func (o *GetRegionsForbidden) String() string {
	return fmt.Sprintf("[GET /regions][%d] getRegionsForbidden ", 403)
}

func (o *GetRegionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegionsNotFound creates a GetRegionsNotFound with default headers values
func NewGetRegionsNotFound() *GetRegionsNotFound {
	return &GetRegionsNotFound{}
}

/*
GetRegionsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetRegionsNotFound struct {
}

// IsSuccess returns true when this get regions not found response has a 2xx status code
func (o *GetRegionsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get regions not found response has a 3xx status code
func (o *GetRegionsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get regions not found response has a 4xx status code
func (o *GetRegionsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get regions not found response has a 5xx status code
func (o *GetRegionsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get regions not found response a status code equal to that given
func (o *GetRegionsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get regions not found response
func (o *GetRegionsNotFound) Code() int {
	return 404
}

func (o *GetRegionsNotFound) Error() string {
	return fmt.Sprintf("[GET /regions][%d] getRegionsNotFound ", 404)
}

func (o *GetRegionsNotFound) String() string {
	return fmt.Sprintf("[GET /regions][%d] getRegionsNotFound ", 404)
}

func (o *GetRegionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetRegionsInternalServerError creates a GetRegionsInternalServerError with default headers values
func NewGetRegionsInternalServerError() *GetRegionsInternalServerError {
	return &GetRegionsInternalServerError{}
}

/*
GetRegionsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetRegionsInternalServerError struct {
}

// IsSuccess returns true when this get regions internal server error response has a 2xx status code
func (o *GetRegionsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get regions internal server error response has a 3xx status code
func (o *GetRegionsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get regions internal server error response has a 4xx status code
func (o *GetRegionsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get regions internal server error response has a 5xx status code
func (o *GetRegionsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get regions internal server error response a status code equal to that given
func (o *GetRegionsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get regions internal server error response
func (o *GetRegionsInternalServerError) Code() int {
	return 500
}

func (o *GetRegionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /regions][%d] getRegionsInternalServerError ", 500)
}

func (o *GetRegionsInternalServerError) String() string {
	return fmt.Sprintf("[GET /regions][%d] getRegionsInternalServerError ", 500)
}

func (o *GetRegionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
GetRegionsOKBody get regions o k body
swagger:model GetRegionsOKBody
*/
type GetRegionsOKBody struct {

	// data
	// Required: true
	Data []*GetRegionsOKBodyDataItems0 `json:"data"`
}

// Validate validates this get regions o k body
func (o *GetRegionsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRegionsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getRegionsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRegionsOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getRegionsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get regions o k body based on the context it is used
func (o *GetRegionsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetRegionsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getRegionsOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getRegionsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetRegionsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRegionsOKBody) UnmarshalBinary(b []byte) error {
	var res GetRegionsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetRegionsOKBodyDataItems0 get regions o k body data items0
swagger:model GetRegionsOKBodyDataItems0
*/
type GetRegionsOKBodyDataItems0 struct {

	// Name of the region
	// Required: true
	DisplayName *string `json:"display_name"`

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

// Validate validates this get regions o k body data items0
func (o *GetRegionsOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDisplayName(formats); err != nil {
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

func (o *GetRegionsOKBodyDataItems0) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("display_name", "body", o.DisplayName); err != nil {
		return err
	}

	return nil
}

func (o *GetRegionsOKBodyDataItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *GetRegionsOKBodyDataItems0) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("location", "body", o.Location); err != nil {
		return err
	}

	return nil
}

func (o *GetRegionsOKBodyDataItems0) validateProvider(formats strfmt.Registry) error {

	if err := validate.Required("provider", "body", o.Provider); err != nil {
		return err
	}

	return nil
}

func (o *GetRegionsOKBodyDataItems0) validatePublicIPAddresses(formats strfmt.Registry) error {

	if err := validate.Required("public_ip_addresses", "body", o.PublicIPAddresses); err != nil {
		return err
	}

	return nil
}

func (o *GetRegionsOKBodyDataItems0) validateSlug(formats strfmt.Registry) error {

	if err := validate.Required("slug", "body", o.Slug); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get regions o k body data items0 based on context it is used
func (o *GetRegionsOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetRegionsOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetRegionsOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res GetRegionsOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
