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

// Message message
// swagger:model message
type Message struct {

	// exitcode
	// Read Only: true
	Exitcode string `json:"exitcode,omitempty"`

	// text
	// Required: true
	// Read Only: true
	Text string `json:"text"`
}

// Validate validates this message
func (m *Message) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateText(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Message) validateText(formats strfmt.Registry) error {

	if err := validate.RequiredString("text", "body", string(m.Text)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Message) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Message) UnmarshalBinary(b []byte) error {
	var res Message
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
