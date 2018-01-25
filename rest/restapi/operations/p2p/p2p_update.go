// Code generated by go-swagger; DO NOT EDIT.

package p2p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// P2pUpdateHandlerFunc turns a function with the right signature into a p2p update handler
type P2pUpdateHandlerFunc func(P2pUpdateParams) middleware.Responder

// Handle executing the request and returning a response
func (fn P2pUpdateHandlerFunc) Handle(params P2pUpdateParams) middleware.Responder {
	return fn(params)
}

// P2pUpdateHandler interface for that can handle valid p2p update params
type P2pUpdateHandler interface {
	Handle(P2pUpdateParams) middleware.Responder
}

// NewP2pUpdate creates a new http.Handler for the p2p update operation
func NewP2pUpdate(ctx *middleware.Context, handler P2pUpdateHandler) *P2pUpdate {
	return &P2pUpdate{Context: ctx, Handler: handler}
}

/*P2pUpdate swagger:route PUT /p2p p2p p2pUpdate

P2P function controls and configures the peer-to-peer network structure: the swarm which includes all hosts with same the same swarm hash and secret key.
P2P is a base layer for Subutai environment networking: all containers in same environment are connected to each other via VXLAN tunnels and are accesses as if they were in one LAN. It doesn't matter where the containers are physically located.

*/
type P2pUpdate struct {
	Context *middleware.Context
	Handler P2pUpdateHandler
}

func (o *P2pUpdate) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewP2pUpdateParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
