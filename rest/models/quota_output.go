// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// QuotaOutput quota output
// swagger:model quotaOutput
type QuotaOutput struct {

	// quota
	// Required: true
	// Read Only: true
	Quota string `json:"quota"`

	// threshold
	// Required: true
	// Min Length: 1
	Threshold *string `json:"threshold"`
}

// Validate validates this quota output
func (m *QuotaOutput) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateQuota(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateThreshold(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *QuotaOutput) validateQuota(formats strfmt.Registry) error {

	if err := validate.RequiredString("quota", "body", string(m.Quota)); err != nil {
		return err
	}

	return nil
}

func (m *QuotaOutput) validateThreshold(formats strfmt.Registry) error {

	if err := validate.Required("threshold", "body", m.Threshold); err != nil {
		return err
	}

	if err := validate.MinLength("threshold", "body", string(*m.Threshold), 1); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *QuotaOutput) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *QuotaOutput) UnmarshalBinary(b []byte) error {
	var res QuotaOutput
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}