// Code generated by go-swagger; DO NOT EDIT.

package destroy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// DestroyHandlerFunc turns a function with the right signature into a destroy handler
type DestroyHandlerFunc func(DestroyParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DestroyHandlerFunc) Handle(params DestroyParams) middleware.Responder {
	return fn(params)
}

// DestroyHandler interface for that can handle valid destroy params
type DestroyHandler interface {
	Handle(DestroyParams) middleware.Responder
}

// NewDestroy creates a new http.Handler for the destroy operation
func NewDestroy(ctx *middleware.Context, handler DestroyHandler) *Destroy {
	return &Destroy{Context: ctx, Handler: handler}
}

/*Destroy swagger:route GET /destroy/{ID} destroy destroy

Destroy Subutai container

*/
type Destroy struct {
	Context *middleware.Context
	Handler DestroyHandler
}

func (o *Destroy) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDestroyParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
