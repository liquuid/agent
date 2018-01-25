// Code generated by go-swagger; DO NOT EDIT.

package vxlan

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// VxlanCreateHandlerFunc turns a function with the right signature into a vxlan create handler
type VxlanCreateHandlerFunc func(VxlanCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn VxlanCreateHandlerFunc) Handle(params VxlanCreateParams) middleware.Responder {
	return fn(params)
}

// VxlanCreateHandler interface for that can handle valid vxlan create params
type VxlanCreateHandler interface {
	Handle(VxlanCreateParams) middleware.Responder
}

// NewVxlanCreate creates a new http.Handler for the vxlan create operation
func NewVxlanCreate(ctx *middleware.Context, handler VxlanCreateHandler) *VxlanCreate {
	return &VxlanCreate{Context: ctx, Handler: handler}
}

/*VxlanCreate swagger:route POST /vxlan/{tunnel} vxlan vxlanCreate

Creates VXLAN tunnel

*/
type VxlanCreate struct {
	Context *middleware.Context
	Handler VxlanCreateHandler
}

func (o *VxlanCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewVxlanCreateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
