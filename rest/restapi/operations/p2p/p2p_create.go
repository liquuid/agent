// Code generated by go-swagger; DO NOT EDIT.

package p2p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// P2pCreateHandlerFunc turns a function with the right signature into a p2p create handler
type P2pCreateHandlerFunc func(P2pCreateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn P2pCreateHandlerFunc) Handle(params P2pCreateParams) middleware.Responder {
	return fn(params)
}

// P2pCreateHandler interface for that can handle valid p2p create params
type P2pCreateHandler interface {
	Handle(P2pCreateParams) middleware.Responder
}

// NewP2pCreate creates a new http.Handler for the p2p create operation
func NewP2pCreate(ctx *middleware.Context, handler P2pCreateHandler) *P2pCreate {
	return &P2pCreate{Context: ctx, Handler: handler}
}

/*P2pCreate swagger:route POST /p2p p2p p2pCreate

P2P function controls and configures the peer-to-peer network structure: the swarm which includes all hosts with same the same swarm hash and secret key.
P2P is a base layer for Subutai environment networking: all containers in same environment are connected to each other via VXLAN tunnels and are accesses as if they were in one LAN. It doesn't matter where the containers are physically located.

*/
type P2pCreate struct {
	Context *middleware.Context
	Handler P2pCreateHandler
}

func (o *P2pCreate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewP2pCreateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
