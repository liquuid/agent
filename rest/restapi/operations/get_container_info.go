// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// GetContainerInfoHandlerFunc turns a function with the right signature into a get container info handler
type GetContainerInfoHandlerFunc func(GetContainerInfoParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GetContainerInfoHandlerFunc) Handle(params GetContainerInfoParams) middleware.Responder {
	return fn(params)
}

// GetContainerInfoHandler interface for that can handle valid get container info params
type GetContainerInfoHandler interface {
	Handle(GetContainerInfoParams) middleware.Responder
}

// NewGetContainerInfo creates a new http.Handler for the get container info operation
func NewGetContainerInfo(ctx *middleware.Context, handler GetContainerInfoHandler) *GetContainerInfo {
	return &GetContainerInfo{Context: ctx, Handler: handler}
}

/*GetContainerInfo swagger:route GET /agent/container/{name} agent cli container getContainerInfo

Get container Info in JSON formatted object

*/
type GetContainerInfo struct {
	Context *middleware.Context
	Handler GetContainerInfoHandler
}

func (o *GetContainerInfo) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewGetContainerInfoParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
