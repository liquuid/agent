// Code generated by go-swagger; DO NOT EDIT.

package promote

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// PromoteHandlerFunc turns a function with the right signature into a promote handler
type PromoteHandlerFunc func(PromoteParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PromoteHandlerFunc) Handle(params PromoteParams) middleware.Responder {
	return fn(params)
}

// PromoteHandler interface for that can handle valid promote params
type PromoteHandler interface {
	Handle(PromoteParams) middleware.Responder
}

// NewPromote creates a new http.Handler for the promote operation
func NewPromote(ctx *middleware.Context, handler PromoteHandler) *Promote {
	return &Promote{Context: ctx, Handler: handler}
}

/*Promote swagger:route POST /promote/{container} promote promote

Promote turns a Subutai container into container template which may be cloned with "clone" command. Promote executes several simple steps, such as dropping a container's configuration to default values, dumping the list of installed packages (this step requires the target container to still be running), and setting the container's filesystem to read-only to prevent changes.

*/
type Promote struct {
	Context *middleware.Context
	Handler PromoteHandler
}

func (o *Promote) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewPromoteParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}