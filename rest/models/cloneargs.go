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

// Cloneargs cloneargs
// swagger:model Cloneargs
type Cloneargs struct {

	// Child name
	// Required: true
	// Read Only: true
	Child string `json:"child"`

	// Environment ID
	// Read Only: true
	EnvID string `json:"envID,omitempty"`

	// IP address
	// Read Only: true
	Ipaddr string `json:"ipaddr,omitempty"`

	// Kurjun Token
	// Read Only: true
	KurjunToken string `json:"kurjunToken,omitempty"`

	// Container/Template parent name
	// Required: true
	// Read Only: true
	Parent string `json:"parent"`

	// token
	// Read Only: true
	Token string `json:"token,omitempty"`
}

// Validate validates this cloneargs
func (m *Cloneargs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateChild(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cloneargs) validateChild(formats strfmt.Registry) error {

	if err := validate.RequiredString("child", "body", string(m.Child)); err != nil {
		return err
	}

	return nil
}

func (m *Cloneargs) validateParent(formats strfmt.Registry) error {

	if err := validate.RequiredString("parent", "body", string(m.Parent)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cloneargs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cloneargs) UnmarshalBinary(b []byte) error {
	var res Cloneargs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
