// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetOrganizationsReader is a Reader for the GetOrganizations structure.
type GetOrganizationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetOrganizationsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetOrganizationsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetOrganizationsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetOrganizationsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[GET /organizations] get_organizations", response, response.Code())
	}
}

// NewGetOrganizationsOK creates a GetOrganizationsOK with default headers values
func NewGetOrganizationsOK() *GetOrganizationsOK {
	return &GetOrganizationsOK{}
}

/*
GetOrganizationsOK describes a response with status code 200, with default header values.

Gets the organizations for the current user
*/
type GetOrganizationsOK struct {
	Payload *GetOrganizationsOKBody
}

// IsSuccess returns true when this get organizations o k response has a 2xx status code
func (o *GetOrganizationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organizations o k response has a 3xx status code
func (o *GetOrganizationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations o k response has a 4xx status code
func (o *GetOrganizationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organizations o k response has a 5xx status code
func (o *GetOrganizationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations o k response a status code equal to that given
func (o *GetOrganizationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organizations o k response
func (o *GetOrganizationsOK) Code() int {
	return 200
}

func (o *GetOrganizationsOK) Error() string {
	return fmt.Sprintf("[GET /organizations][%d] getOrganizationsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationsOK) String() string {
	return fmt.Sprintf("[GET /organizations][%d] getOrganizationsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationsOK) GetPayload() *GetOrganizationsOKBody {
	return o.Payload
}

func (o *GetOrganizationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrganizationsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationsUnauthorized creates a GetOrganizationsUnauthorized with default headers values
func NewGetOrganizationsUnauthorized() *GetOrganizationsUnauthorized {
	return &GetOrganizationsUnauthorized{}
}

/*
GetOrganizationsUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type GetOrganizationsUnauthorized struct {
}

// IsSuccess returns true when this get organizations unauthorized response has a 2xx status code
func (o *GetOrganizationsUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations unauthorized response has a 3xx status code
func (o *GetOrganizationsUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations unauthorized response has a 4xx status code
func (o *GetOrganizationsUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations unauthorized response has a 5xx status code
func (o *GetOrganizationsUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations unauthorized response a status code equal to that given
func (o *GetOrganizationsUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the get organizations unauthorized response
func (o *GetOrganizationsUnauthorized) Code() int {
	return 401
}

func (o *GetOrganizationsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /organizations][%d] getOrganizationsUnauthorized ", 401)
}

func (o *GetOrganizationsUnauthorized) String() string {
	return fmt.Sprintf("[GET /organizations][%d] getOrganizationsUnauthorized ", 401)
}

func (o *GetOrganizationsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationsForbidden creates a GetOrganizationsForbidden with default headers values
func NewGetOrganizationsForbidden() *GetOrganizationsForbidden {
	return &GetOrganizationsForbidden{}
}

/*
GetOrganizationsForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type GetOrganizationsForbidden struct {
}

// IsSuccess returns true when this get organizations forbidden response has a 2xx status code
func (o *GetOrganizationsForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations forbidden response has a 3xx status code
func (o *GetOrganizationsForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations forbidden response has a 4xx status code
func (o *GetOrganizationsForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations forbidden response has a 5xx status code
func (o *GetOrganizationsForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations forbidden response a status code equal to that given
func (o *GetOrganizationsForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the get organizations forbidden response
func (o *GetOrganizationsForbidden) Code() int {
	return 403
}

func (o *GetOrganizationsForbidden) Error() string {
	return fmt.Sprintf("[GET /organizations][%d] getOrganizationsForbidden ", 403)
}

func (o *GetOrganizationsForbidden) String() string {
	return fmt.Sprintf("[GET /organizations][%d] getOrganizationsForbidden ", 403)
}

func (o *GetOrganizationsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationsNotFound creates a GetOrganizationsNotFound with default headers values
func NewGetOrganizationsNotFound() *GetOrganizationsNotFound {
	return &GetOrganizationsNotFound{}
}

/*
GetOrganizationsNotFound describes a response with status code 404, with default header values.

Not Found
*/
type GetOrganizationsNotFound struct {
}

// IsSuccess returns true when this get organizations not found response has a 2xx status code
func (o *GetOrganizationsNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations not found response has a 3xx status code
func (o *GetOrganizationsNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations not found response has a 4xx status code
func (o *GetOrganizationsNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this get organizations not found response has a 5xx status code
func (o *GetOrganizationsNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this get organizations not found response a status code equal to that given
func (o *GetOrganizationsNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the get organizations not found response
func (o *GetOrganizationsNotFound) Code() int {
	return 404
}

func (o *GetOrganizationsNotFound) Error() string {
	return fmt.Sprintf("[GET /organizations][%d] getOrganizationsNotFound ", 404)
}

func (o *GetOrganizationsNotFound) String() string {
	return fmt.Sprintf("[GET /organizations][%d] getOrganizationsNotFound ", 404)
}

func (o *GetOrganizationsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetOrganizationsInternalServerError creates a GetOrganizationsInternalServerError with default headers values
func NewGetOrganizationsInternalServerError() *GetOrganizationsInternalServerError {
	return &GetOrganizationsInternalServerError{}
}

/*
GetOrganizationsInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type GetOrganizationsInternalServerError struct {
}

// IsSuccess returns true when this get organizations internal server error response has a 2xx status code
func (o *GetOrganizationsInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this get organizations internal server error response has a 3xx status code
func (o *GetOrganizationsInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organizations internal server error response has a 4xx status code
func (o *GetOrganizationsInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organizations internal server error response has a 5xx status code
func (o *GetOrganizationsInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this get organizations internal server error response a status code equal to that given
func (o *GetOrganizationsInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the get organizations internal server error response
func (o *GetOrganizationsInternalServerError) Code() int {
	return 500
}

func (o *GetOrganizationsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /organizations][%d] getOrganizationsInternalServerError ", 500)
}

func (o *GetOrganizationsInternalServerError) String() string {
	return fmt.Sprintf("[GET /organizations][%d] getOrganizationsInternalServerError ", 500)
}

func (o *GetOrganizationsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
GetOrganizationsOKBody get organizations o k body
swagger:model GetOrganizationsOKBody
*/
type GetOrganizationsOKBody struct {

	// data
	// Required: true
	Data []*GetOrganizationsOKBodyDataItems0 `json:"data"`
}

// Validate validates this get organizations o k body
func (o *GetOrganizationsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationsOKBody) validateData(formats strfmt.Registry) error {

	if err := validate.Required("getOrganizationsOK"+"."+"data", "body", o.Data); err != nil {
		return err
	}

	for i := 0; i < len(o.Data); i++ {
		if swag.IsZero(o.Data[i]) { // not required
			continue
		}

		if o.Data[i] != nil {
			if err := o.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrganizationsOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getOrganizationsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get organizations o k body based on the context it is used
func (o *GetOrganizationsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationsOKBody) contextValidateData(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Data); i++ {

		if o.Data[i] != nil {

			if swag.IsZero(o.Data[i]) { // not required
				return nil
			}

			if err := o.Data[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getOrganizationsOK" + "." + "data" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getOrganizationsOK" + "." + "data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationsOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrganizationsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationsOKBodyDataItems0 get organizations o k body data items0
swagger:model GetOrganizationsOKBodyDataItems0
*/
type GetOrganizationsOKBodyDataItems0 struct {

	// The billing email of the organization
	BillingEmail string `json:"billing_email,omitempty"`

	// When the organization was created
	// Required: true
	CreatedAt *string `json:"created_at"`

	// The number of databases in the organization
	// Required: true
	DatabaseCount *float64 `json:"database_count"`

	// features
	Features *GetOrganizationsOKBodyDataItems0Features `json:"features,omitempty"`

	// flags
	Flags *GetOrganizationsOKBodyDataItems0Flags `json:"flags,omitempty"`

	// Whether or not the organization has past due billing invoices
	// Required: true
	HasPastDueInvoices *bool `json:"has_past_due_invoices"`

	// The ID for the organization
	// Required: true
	ID *string `json:"id"`

	// Whether or not the IdP provider is be responsible for managing roles in PlanetScale
	// Required: true
	IdpManagedRoles *bool `json:"idp_managed_roles"`

	// The expected monthly budget for the organization
	// Required: true
	InvoiceBudgetAmount *float64 `json:"invoice_budget_amount"`

	// The name of the organization
	// Required: true
	Name *string `json:"name"`

	// The billing plan of the organization
	// Required: true
	Plan *string `json:"plan"`

	// Whether or not the organization has single tenancy enabled
	// Required: true
	SingleTenancy *bool `json:"single_tenancy"`

	// The number of sleeping databases in the organization
	// Required: true
	SleepingDatabaseCount *float64 `json:"sleeping_database_count"`

	// Whether or not SSO is enabled on the organization
	// Required: true
	Sso *bool `json:"sso"`

	// Whether or not the organization uses a WorkOS directory
	// Required: true
	SsoDirectory *bool `json:"sso_directory"`

	// The URL of the organization's SSO portal
	SsoPortalURL string `json:"sso_portal_url,omitempty"`

	// When the organization was last updated
	// Required: true
	UpdatedAt *string `json:"updated_at"`

	// Whether or not the organization's billing information is valid
	// Required: true
	ValidBillingInfo *bool `json:"valid_billing_info"`
}

// Validate validates this get organizations o k body data items0
func (o *GetOrganizationsOKBodyDataItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDatabaseCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFeatures(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFlags(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHasPastDueInvoices(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIdpManagedRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInvoiceBudgetAmount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePlan(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSingleTenancy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSleepingDatabaseCount(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSso(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSsoDirectory(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValidBillingInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("created_at", "body", o.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validateDatabaseCount(formats strfmt.Registry) error {

	if err := validate.Required("database_count", "body", o.DatabaseCount); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validateFeatures(formats strfmt.Registry) error {
	if swag.IsZero(o.Features) { // not required
		return nil
	}

	if o.Features != nil {
		if err := o.Features.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validateFlags(formats strfmt.Registry) error {
	if swag.IsZero(o.Flags) { // not required
		return nil
	}

	if o.Flags != nil {
		if err := o.Flags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flags")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validateHasPastDueInvoices(formats strfmt.Registry) error {

	if err := validate.Required("has_past_due_invoices", "body", o.HasPastDueInvoices); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validateIdpManagedRoles(formats strfmt.Registry) error {

	if err := validate.Required("idp_managed_roles", "body", o.IdpManagedRoles); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validateInvoiceBudgetAmount(formats strfmt.Registry) error {

	if err := validate.Required("invoice_budget_amount", "body", o.InvoiceBudgetAmount); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validatePlan(formats strfmt.Registry) error {

	if err := validate.Required("plan", "body", o.Plan); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validateSingleTenancy(formats strfmt.Registry) error {

	if err := validate.Required("single_tenancy", "body", o.SingleTenancy); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validateSleepingDatabaseCount(formats strfmt.Registry) error {

	if err := validate.Required("sleeping_database_count", "body", o.SleepingDatabaseCount); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validateSso(formats strfmt.Registry) error {

	if err := validate.Required("sso", "body", o.Sso); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validateSsoDirectory(formats strfmt.Registry) error {

	if err := validate.Required("sso_directory", "body", o.SsoDirectory); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("updated_at", "body", o.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) validateValidBillingInfo(formats strfmt.Registry) error {

	if err := validate.Required("valid_billing_info", "body", o.ValidBillingInfo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get organizations o k body data items0 based on the context it is used
func (o *GetOrganizationsOKBodyDataItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFeatures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateFlags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) contextValidateFeatures(ctx context.Context, formats strfmt.Registry) error {

	if o.Features != nil {

		if swag.IsZero(o.Features) { // not required
			return nil
		}

		if err := o.Features.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("features")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("features")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationsOKBodyDataItems0) contextValidateFlags(ctx context.Context, formats strfmt.Registry) error {

	if o.Flags != nil {

		if swag.IsZero(o.Flags) { // not required
			return nil
		}

		if err := o.Flags.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("flags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("flags")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationsOKBodyDataItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationsOKBodyDataItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationsOKBodyDataItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationsOKBodyDataItems0Features get organizations o k body data items0 features
swagger:model GetOrganizationsOKBodyDataItems0Features
*/
type GetOrganizationsOKBodyDataItems0Features struct {

	// insights
	Insights bool `json:"insights,omitempty"`

	// single tenancy
	SingleTenancy bool `json:"single_tenancy,omitempty"`

	// sso
	Sso bool `json:"sso,omitempty"`
}

// Validate validates this get organizations o k body data items0 features
func (o *GetOrganizationsOKBodyDataItems0Features) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organizations o k body data items0 features based on context it is used
func (o *GetOrganizationsOKBodyDataItems0Features) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationsOKBodyDataItems0Features) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationsOKBodyDataItems0Features) UnmarshalBinary(b []byte) error {
	var res GetOrganizationsOKBodyDataItems0Features
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationsOKBodyDataItems0Flags get organizations o k body data items0 flags
swagger:model GetOrganizationsOKBodyDataItems0Flags
*/
type GetOrganizationsOKBodyDataItems0Flags struct {

	// example flag
	// Enum: [full on]
	ExampleFlag string `json:"example_flag,omitempty"`
}

// Validate validates this get organizations o k body data items0 flags
func (o *GetOrganizationsOKBodyDataItems0Flags) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExampleFlag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getOrganizationsOKBodyDataItems0FlagsTypeExampleFlagPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["full","on"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getOrganizationsOKBodyDataItems0FlagsTypeExampleFlagPropEnum = append(getOrganizationsOKBodyDataItems0FlagsTypeExampleFlagPropEnum, v)
	}
}

const (

	// GetOrganizationsOKBodyDataItems0FlagsExampleFlagFull captures enum value "full"
	GetOrganizationsOKBodyDataItems0FlagsExampleFlagFull string = "full"

	// GetOrganizationsOKBodyDataItems0FlagsExampleFlagOn captures enum value "on"
	GetOrganizationsOKBodyDataItems0FlagsExampleFlagOn string = "on"
)

// prop value enum
func (o *GetOrganizationsOKBodyDataItems0Flags) validateExampleFlagEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getOrganizationsOKBodyDataItems0FlagsTypeExampleFlagPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetOrganizationsOKBodyDataItems0Flags) validateExampleFlag(formats strfmt.Registry) error {
	if swag.IsZero(o.ExampleFlag) { // not required
		return nil
	}

	// value enum
	if err := o.validateExampleFlagEnum("flags"+"."+"example_flag", "body", o.ExampleFlag); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get organizations o k body data items0 flags based on context it is used
func (o *GetOrganizationsOKBodyDataItems0Flags) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationsOKBodyDataItems0Flags) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationsOKBodyDataItems0Flags) UnmarshalBinary(b []byte) error {
	var res GetOrganizationsOKBodyDataItems0Flags
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
