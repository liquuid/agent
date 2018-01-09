// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// DemoteOptions demote options
// swagger:model demoteOptions
type DemoteOptions struct {

	// ipaddr
	// Read Only: true
	Ipaddr string `json:"ipaddr,omitempty"`

	// vlan
	// Read Only: true
	Vlan string `json:"vlan,omitempty"`
}

// Validate validates this demote options
func (m *DemoteOptions) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *DemoteOptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DemoteOptions) UnmarshalBinary(b []byte) error {
	var res DemoteOptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
