// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// PatchOrganizationsNameReader is a Reader for the PatchOrganizationsName structure.
type PatchOrganizationsNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchOrganizationsNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPatchOrganizationsNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewPatchOrganizationsNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPatchOrganizationsNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPatchOrganizationsNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPatchOrganizationsNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[PATCH /organizations/{name}] patch_organizations_name", response, response.Code())
	}
}

// NewPatchOrganizationsNameOK creates a PatchOrganizationsNameOK with default headers values
func NewPatchOrganizationsNameOK() *PatchOrganizationsNameOK {
	return &PatchOrganizationsNameOK{}
}

/*
PatchOrganizationsNameOK describes a response with status code 200, with default header values.

Returns the updated organization
*/
type PatchOrganizationsNameOK struct {
	Payload *PatchOrganizationsNameOKBody
}

// IsSuccess returns true when this patch organizations name o k response has a 2xx status code
func (o *PatchOrganizationsNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this patch organizations name o k response has a 3xx status code
func (o *PatchOrganizationsNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch organizations name o k response has a 4xx status code
func (o *PatchOrganizationsNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch organizations name o k response has a 5xx status code
func (o *PatchOrganizationsNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this patch organizations name o k response a status code equal to that given
func (o *PatchOrganizationsNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the patch organizations name o k response
func (o *PatchOrganizationsNameOK) Code() int {
	return 200
}

func (o *PatchOrganizationsNameOK) Error() string {
	return fmt.Sprintf("[PATCH /organizations/{name}][%d] patchOrganizationsNameOK  %+v", 200, o.Payload)
}

func (o *PatchOrganizationsNameOK) String() string {
	return fmt.Sprintf("[PATCH /organizations/{name}][%d] patchOrganizationsNameOK  %+v", 200, o.Payload)
}

func (o *PatchOrganizationsNameOK) GetPayload() *PatchOrganizationsNameOKBody {
	return o.Payload
}

func (o *PatchOrganizationsNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(PatchOrganizationsNameOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchOrganizationsNameUnauthorized creates a PatchOrganizationsNameUnauthorized with default headers values
func NewPatchOrganizationsNameUnauthorized() *PatchOrganizationsNameUnauthorized {
	return &PatchOrganizationsNameUnauthorized{}
}

/*
PatchOrganizationsNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PatchOrganizationsNameUnauthorized struct {
}

// IsSuccess returns true when this patch organizations name unauthorized response has a 2xx status code
func (o *PatchOrganizationsNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch organizations name unauthorized response has a 3xx status code
func (o *PatchOrganizationsNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch organizations name unauthorized response has a 4xx status code
func (o *PatchOrganizationsNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch organizations name unauthorized response has a 5xx status code
func (o *PatchOrganizationsNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this patch organizations name unauthorized response a status code equal to that given
func (o *PatchOrganizationsNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the patch organizations name unauthorized response
func (o *PatchOrganizationsNameUnauthorized) Code() int {
	return 401
}

func (o *PatchOrganizationsNameUnauthorized) Error() string {
	return fmt.Sprintf("[PATCH /organizations/{name}][%d] patchOrganizationsNameUnauthorized ", 401)
}

func (o *PatchOrganizationsNameUnauthorized) String() string {
	return fmt.Sprintf("[PATCH /organizations/{name}][%d] patchOrganizationsNameUnauthorized ", 401)
}

func (o *PatchOrganizationsNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchOrganizationsNameForbidden creates a PatchOrganizationsNameForbidden with default headers values
func NewPatchOrganizationsNameForbidden() *PatchOrganizationsNameForbidden {
	return &PatchOrganizationsNameForbidden{}
}

/*
PatchOrganizationsNameForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PatchOrganizationsNameForbidden struct {
}

// IsSuccess returns true when this patch organizations name forbidden response has a 2xx status code
func (o *PatchOrganizationsNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch organizations name forbidden response has a 3xx status code
func (o *PatchOrganizationsNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch organizations name forbidden response has a 4xx status code
func (o *PatchOrganizationsNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch organizations name forbidden response has a 5xx status code
func (o *PatchOrganizationsNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this patch organizations name forbidden response a status code equal to that given
func (o *PatchOrganizationsNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the patch organizations name forbidden response
func (o *PatchOrganizationsNameForbidden) Code() int {
	return 403
}

func (o *PatchOrganizationsNameForbidden) Error() string {
	return fmt.Sprintf("[PATCH /organizations/{name}][%d] patchOrganizationsNameForbidden ", 403)
}

func (o *PatchOrganizationsNameForbidden) String() string {
	return fmt.Sprintf("[PATCH /organizations/{name}][%d] patchOrganizationsNameForbidden ", 403)
}

func (o *PatchOrganizationsNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchOrganizationsNameNotFound creates a PatchOrganizationsNameNotFound with default headers values
func NewPatchOrganizationsNameNotFound() *PatchOrganizationsNameNotFound {
	return &PatchOrganizationsNameNotFound{}
}

/*
PatchOrganizationsNameNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PatchOrganizationsNameNotFound struct {
}

// IsSuccess returns true when this patch organizations name not found response has a 2xx status code
func (o *PatchOrganizationsNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch organizations name not found response has a 3xx status code
func (o *PatchOrganizationsNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch organizations name not found response has a 4xx status code
func (o *PatchOrganizationsNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this patch organizations name not found response has a 5xx status code
func (o *PatchOrganizationsNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this patch organizations name not found response a status code equal to that given
func (o *PatchOrganizationsNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the patch organizations name not found response
func (o *PatchOrganizationsNameNotFound) Code() int {
	return 404
}

func (o *PatchOrganizationsNameNotFound) Error() string {
	return fmt.Sprintf("[PATCH /organizations/{name}][%d] patchOrganizationsNameNotFound ", 404)
}

func (o *PatchOrganizationsNameNotFound) String() string {
	return fmt.Sprintf("[PATCH /organizations/{name}][%d] patchOrganizationsNameNotFound ", 404)
}

func (o *PatchOrganizationsNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewPatchOrganizationsNameInternalServerError creates a PatchOrganizationsNameInternalServerError with default headers values
func NewPatchOrganizationsNameInternalServerError() *PatchOrganizationsNameInternalServerError {
	return &PatchOrganizationsNameInternalServerError{}
}

/*
PatchOrganizationsNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PatchOrganizationsNameInternalServerError struct {
}

// IsSuccess returns true when this patch organizations name internal server error response has a 2xx status code
func (o *PatchOrganizationsNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this patch organizations name internal server error response has a 3xx status code
func (o *PatchOrganizationsNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this patch organizations name internal server error response has a 4xx status code
func (o *PatchOrganizationsNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this patch organizations name internal server error response has a 5xx status code
func (o *PatchOrganizationsNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this patch organizations name internal server error response a status code equal to that given
func (o *PatchOrganizationsNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the patch organizations name internal server error response
func (o *PatchOrganizationsNameInternalServerError) Code() int {
	return 500
}

func (o *PatchOrganizationsNameInternalServerError) Error() string {
	return fmt.Sprintf("[PATCH /organizations/{name}][%d] patchOrganizationsNameInternalServerError ", 500)
}

func (o *PatchOrganizationsNameInternalServerError) String() string {
	return fmt.Sprintf("[PATCH /organizations/{name}][%d] patchOrganizationsNameInternalServerError ", 500)
}

func (o *PatchOrganizationsNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*
PatchOrganizationsNameBody patch organizations name body
swagger:model PatchOrganizationsNameBody
*/
type PatchOrganizationsNameBody struct {

	// The billing email for the organization
	BillingEmail string `json:"billing_email,omitempty"`

	// Whether or not the IdP provider is be responsible for managing roles in PlanetScale
	IdpManagedRoles bool `json:"idp_managed_roles,omitempty"`

	// The expected monthly budget for the organization
	InvoiceBudgetAmount float64 `json:"invoice_budget_amount,omitempty"`
}

// Validate validates this patch organizations name body
func (o *PatchOrganizationsNameBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch organizations name body based on context it is used
func (o *PatchOrganizationsNameBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchOrganizationsNameBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchOrganizationsNameBody) UnmarshalBinary(b []byte) error {
	var res PatchOrganizationsNameBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PatchOrganizationsNameOKBody patch organizations name o k body
swagger:model PatchOrganizationsNameOKBody
*/
type PatchOrganizationsNameOKBody struct {

	// The billing email of the organization
	BillingEmail string `json:"billing_email,omitempty"`

	// When the organization was created
	// Required: true
	CreatedAt *string `json:"created_at"`

	// The number of databases in the organization
	// Required: true
	DatabaseCount *float64 `json:"database_count"`

	// features
	Features *PatchOrganizationsNameOKBodyFeatures `json:"features,omitempty"`

	// flags
	Flags *PatchOrganizationsNameOKBodyFlags `json:"flags,omitempty"`

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

// Validate validates this patch organizations name o k body
func (o *PatchOrganizationsNameOKBody) Validate(formats strfmt.Registry) error {
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

func (o *PatchOrganizationsNameOKBody) validateCreatedAt(formats strfmt.Registry) error {

	if err := validate.Required("patchOrganizationsNameOK"+"."+"created_at", "body", o.CreatedAt); err != nil {
		return err
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) validateDatabaseCount(formats strfmt.Registry) error {

	if err := validate.Required("patchOrganizationsNameOK"+"."+"database_count", "body", o.DatabaseCount); err != nil {
		return err
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) validateFeatures(formats strfmt.Registry) error {
	if swag.IsZero(o.Features) { // not required
		return nil
	}

	if o.Features != nil {
		if err := o.Features.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patchOrganizationsNameOK" + "." + "features")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("patchOrganizationsNameOK" + "." + "features")
			}
			return err
		}
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) validateFlags(formats strfmt.Registry) error {
	if swag.IsZero(o.Flags) { // not required
		return nil
	}

	if o.Flags != nil {
		if err := o.Flags.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patchOrganizationsNameOK" + "." + "flags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("patchOrganizationsNameOK" + "." + "flags")
			}
			return err
		}
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) validateHasPastDueInvoices(formats strfmt.Registry) error {

	if err := validate.Required("patchOrganizationsNameOK"+"."+"has_past_due_invoices", "body", o.HasPastDueInvoices); err != nil {
		return err
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) validateID(formats strfmt.Registry) error {

	if err := validate.Required("patchOrganizationsNameOK"+"."+"id", "body", o.ID); err != nil {
		return err
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) validateIdpManagedRoles(formats strfmt.Registry) error {

	if err := validate.Required("patchOrganizationsNameOK"+"."+"idp_managed_roles", "body", o.IdpManagedRoles); err != nil {
		return err
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) validateInvoiceBudgetAmount(formats strfmt.Registry) error {

	if err := validate.Required("patchOrganizationsNameOK"+"."+"invoice_budget_amount", "body", o.InvoiceBudgetAmount); err != nil {
		return err
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("patchOrganizationsNameOK"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) validatePlan(formats strfmt.Registry) error {

	if err := validate.Required("patchOrganizationsNameOK"+"."+"plan", "body", o.Plan); err != nil {
		return err
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) validateSingleTenancy(formats strfmt.Registry) error {

	if err := validate.Required("patchOrganizationsNameOK"+"."+"single_tenancy", "body", o.SingleTenancy); err != nil {
		return err
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) validateSleepingDatabaseCount(formats strfmt.Registry) error {

	if err := validate.Required("patchOrganizationsNameOK"+"."+"sleeping_database_count", "body", o.SleepingDatabaseCount); err != nil {
		return err
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) validateSso(formats strfmt.Registry) error {

	if err := validate.Required("patchOrganizationsNameOK"+"."+"sso", "body", o.Sso); err != nil {
		return err
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) validateSsoDirectory(formats strfmt.Registry) error {

	if err := validate.Required("patchOrganizationsNameOK"+"."+"sso_directory", "body", o.SsoDirectory); err != nil {
		return err
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) validateUpdatedAt(formats strfmt.Registry) error {

	if err := validate.Required("patchOrganizationsNameOK"+"."+"updated_at", "body", o.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) validateValidBillingInfo(formats strfmt.Registry) error {

	if err := validate.Required("patchOrganizationsNameOK"+"."+"valid_billing_info", "body", o.ValidBillingInfo); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this patch organizations name o k body based on the context it is used
func (o *PatchOrganizationsNameOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *PatchOrganizationsNameOKBody) contextValidateFeatures(ctx context.Context, formats strfmt.Registry) error {

	if o.Features != nil {

		if swag.IsZero(o.Features) { // not required
			return nil
		}

		if err := o.Features.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patchOrganizationsNameOK" + "." + "features")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("patchOrganizationsNameOK" + "." + "features")
			}
			return err
		}
	}

	return nil
}

func (o *PatchOrganizationsNameOKBody) contextValidateFlags(ctx context.Context, formats strfmt.Registry) error {

	if o.Flags != nil {

		if swag.IsZero(o.Flags) { // not required
			return nil
		}

		if err := o.Flags.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("patchOrganizationsNameOK" + "." + "flags")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("patchOrganizationsNameOK" + "." + "flags")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *PatchOrganizationsNameOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchOrganizationsNameOKBody) UnmarshalBinary(b []byte) error {
	var res PatchOrganizationsNameOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PatchOrganizationsNameOKBodyFeatures patch organizations name o k body features
swagger:model PatchOrganizationsNameOKBodyFeatures
*/
type PatchOrganizationsNameOKBodyFeatures struct {

	// insights
	Insights bool `json:"insights,omitempty"`

	// single tenancy
	SingleTenancy bool `json:"single_tenancy,omitempty"`

	// sso
	Sso bool `json:"sso,omitempty"`
}

// Validate validates this patch organizations name o k body features
func (o *PatchOrganizationsNameOKBodyFeatures) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this patch organizations name o k body features based on context it is used
func (o *PatchOrganizationsNameOKBodyFeatures) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchOrganizationsNameOKBodyFeatures) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchOrganizationsNameOKBodyFeatures) UnmarshalBinary(b []byte) error {
	var res PatchOrganizationsNameOKBodyFeatures
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
PatchOrganizationsNameOKBodyFlags patch organizations name o k body flags
swagger:model PatchOrganizationsNameOKBodyFlags
*/
type PatchOrganizationsNameOKBodyFlags struct {

	// example flag
	// Enum: [full on]
	ExampleFlag string `json:"example_flag,omitempty"`
}

// Validate validates this patch organizations name o k body flags
func (o *PatchOrganizationsNameOKBodyFlags) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExampleFlag(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var patchOrganizationsNameOKBodyFlagsTypeExampleFlagPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["full","on"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		patchOrganizationsNameOKBodyFlagsTypeExampleFlagPropEnum = append(patchOrganizationsNameOKBodyFlagsTypeExampleFlagPropEnum, v)
	}
}

const (

	// PatchOrganizationsNameOKBodyFlagsExampleFlagFull captures enum value "full"
	PatchOrganizationsNameOKBodyFlagsExampleFlagFull string = "full"

	// PatchOrganizationsNameOKBodyFlagsExampleFlagOn captures enum value "on"
	PatchOrganizationsNameOKBodyFlagsExampleFlagOn string = "on"
)

// prop value enum
func (o *PatchOrganizationsNameOKBodyFlags) validateExampleFlagEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, patchOrganizationsNameOKBodyFlagsTypeExampleFlagPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *PatchOrganizationsNameOKBodyFlags) validateExampleFlag(formats strfmt.Registry) error {
	if swag.IsZero(o.ExampleFlag) { // not required
		return nil
	}

	// value enum
	if err := o.validateExampleFlagEnum("patchOrganizationsNameOK"+"."+"flags"+"."+"example_flag", "body", o.ExampleFlag); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this patch organizations name o k body flags based on context it is used
func (o *PatchOrganizationsNameOKBodyFlags) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PatchOrganizationsNameOKBodyFlags) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PatchOrganizationsNameOKBodyFlags) UnmarshalBinary(b []byte) error {
	var res PatchOrganizationsNameOKBodyFlags
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
