// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProxyArgs proxy args
// swagger:model ProxyArgs
type ProxyArgs struct {

	// SSL certificate
	// Min Length: 1
	Cert string `json:"cert,omitempty"`

	// Domain name
	// Min Length: 1
	Domain string `json:"domain,omitempty"`

	// Node Name
	// Min Length: 1
	Node string `json:"node,omitempty"`

	// policy
	Policy string `json:"policy,omitempty"`

	// VLan Name
	// Required: true
	// Read Only: true
	Vlan string `json:"vlan"`
}

// Validate validates this proxy args
func (m *ProxyArgs) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCert(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateDomain(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validatePolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateVlan(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProxyArgs) validateCert(formats strfmt.Registry) error {

	if swag.IsZero(m.Cert) { // not required
		return nil
	}

	if err := validate.MinLength("cert", "body", string(m.Cert), 1); err != nil {
		return err
	}

	return nil
}

func (m *ProxyArgs) validateDomain(formats strfmt.Registry) error {

	if swag.IsZero(m.Domain) { // not required
		return nil
	}

	if err := validate.MinLength("domain", "body", string(m.Domain), 1); err != nil {
		return err
	}

	return nil
}

func (m *ProxyArgs) validateNode(formats strfmt.Registry) error {

	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if err := validate.MinLength("node", "body", string(m.Node), 1); err != nil {
		return err
	}

	return nil
}

var proxyArgsTypePolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["rr","lb","hash"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		proxyArgsTypePolicyPropEnum = append(proxyArgsTypePolicyPropEnum, v)
	}
}

const (
	// ProxyArgsPolicyRr captures enum value "rr"
	ProxyArgsPolicyRr string = "rr"
	// ProxyArgsPolicyLb captures enum value "lb"
	ProxyArgsPolicyLb string = "lb"
	// ProxyArgsPolicyHash captures enum value "hash"
	ProxyArgsPolicyHash string = "hash"
)

// prop value enum
func (m *ProxyArgs) validatePolicyEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, proxyArgsTypePolicyPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProxyArgs) validatePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.Policy) { // not required
		return nil
	}

	// value enum
	if err := m.validatePolicyEnum("policy", "body", m.Policy); err != nil {
		return err
	}

	return nil
}

func (m *ProxyArgs) validateVlan(formats strfmt.Registry) error {

	if err := validate.RequiredString("vlan", "body", string(m.Vlan)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProxyArgs) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProxyArgs) UnmarshalBinary(b []byte) error {
	var res ProxyArgs
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
