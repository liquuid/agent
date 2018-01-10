// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewTunnelAddParams creates a new TunnelAddParams object
// with the default values initialized.
func NewTunnelAddParams() TunnelAddParams {
	var ()
	return TunnelAddParams{}
}

// TunnelAddParams contains all the bound params for the tunnel add operation
// typically these are obtained from a http.Request
//
// swagger:parameters tunnelAdd
type TunnelAddParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*
	  In: query
	*/
	Global *bool
	/*
	  Required: true
	  In: path
	*/
	Socket string
	/*
	  In: query
	*/
	Timeout *string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *TunnelAddParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	qGlobal, qhkGlobal, _ := qs.GetOK("global")
	if err := o.bindGlobal(qGlobal, qhkGlobal, route.Formats); err != nil {
		res = append(res, err)
	}

	rSocket, rhkSocket, _ := route.Params.GetOK("socket")
	if err := o.bindSocket(rSocket, rhkSocket, route.Formats); err != nil {
		res = append(res, err)
	}

	qTimeout, qhkTimeout, _ := qs.GetOK("timeout")
	if err := o.bindTimeout(qTimeout, qhkTimeout, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *TunnelAddParams) bindGlobal(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	value, err := swag.ConvertBool(raw)
	if err != nil {
		return errors.InvalidType("global", "query", "bool", raw)
	}
	o.Global = &value

	return nil
}

func (o *TunnelAddParams) bindSocket(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Socket = raw

	return nil
}

func (o *TunnelAddParams) bindTimeout(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if raw == "" { // empty values pass all other validations
		return nil
	}

	o.Timeout = &raw

	return nil
}
