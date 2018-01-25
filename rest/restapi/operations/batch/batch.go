// Code generated by go-swagger; DO NOT EDIT.

package batch

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// BatchHandlerFunc turns a function with the right signature into a batch handler
type BatchHandlerFunc func(BatchParams) middleware.Responder

// Handle executing the request and returning a response
func (fn BatchHandlerFunc) Handle(params BatchParams) middleware.Responder {
	return fn(params)
}

// BatchHandler interface for that can handle valid batch params
type BatchHandler interface {
	Handle(BatchParams) middleware.Responder
}

// NewBatch creates a new http.Handler for the batch operation
func NewBatch(ctx *middleware.Context, handler BatchHandler) *Batch {
	return &Batch{Context: ctx, Handler: handler}
}

/*Batch swagger:route POST /batch batch batch

Batch binding provides a mechanism to perform several Subutai commands in the container in batch, passed in a single JSON message. Initially, the purpose of this command was internal for SS <-> Agent communication, yet it may be invoked manually from the CLI. The response from a batch command returns a JSON array with each element representing the results (response) from each command (request) in the batch: the positions of responses correlate with the request position in the array

*/
type Batch struct {
	Context *middleware.Context
	Handler BatchHandler
}

func (o *Batch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewBatchParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
