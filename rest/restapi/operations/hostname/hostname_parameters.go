// Code generated by go-swagger; DO NOT EDIT.

package hostname

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"

	strfmt "github.com/go-openapi/strfmt"
)

// NewHostnameParams creates a new HostnameParams object
// with the default values initialized.
func NewHostnameParams() HostnameParams {
	var ()
	return HostnameParams{}
}

// HostnameParams contains all the bound params for the hostname operation
// typically these are obtained from a http.Request
//
// swagger:parameters hostname
type HostnameParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request `json:"-"`

	/*Container name
	  Required: true
	  In: path
	*/
	Container string
	/*New name
	  Required: true
	  In: path
	*/
	Name string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *HostnameParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	rContainer, rhkContainer, _ := route.Params.GetOK("container")
	if err := o.bindContainer(rContainer, rhkContainer, route.Formats); err != nil {
		res = append(res, err)
	}

	rName, rhkName, _ := route.Params.GetOK("name")
	if err := o.bindName(rName, rhkName, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *HostnameParams) bindContainer(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Container = raw

	return nil
}

func (o *HostnameParams) bindName(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.Name = raw

	return nil
}