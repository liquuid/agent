// Code generated by go-swagger; DO NOT EDIT.

package p2p

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/subutai-io/agent/rest/models"
)

// P2pCreateCreatedCode is the HTTP code returned for type P2pCreateCreated
const P2pCreateCreatedCode int = 201

/*P2pCreateCreated Created

swagger:response p2pCreateCreated
*/
type P2pCreateCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Item `json:"body,omitempty"`
}

// NewP2pCreateCreated creates P2pCreateCreated with default headers values
func NewP2pCreateCreated() *P2pCreateCreated {
	return &P2pCreateCreated{}
}

// WithPayload adds the payload to the p2p create created response
func (o *P2pCreateCreated) WithPayload(payload *models.Item) *P2pCreateCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the p2p create created response
func (o *P2pCreateCreated) SetPayload(payload *models.Item) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *P2pCreateCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*P2pCreateDefault generic error response

swagger:response p2pCreateDefault
*/
type P2pCreateDefault struct {
	_statusCode int
}

// NewP2pCreateDefault creates P2pCreateDefault with default headers values
func NewP2pCreateDefault(code int) *P2pCreateDefault {
	if code <= 0 {
		code = 500
	}

	return &P2pCreateDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the p2p create default response
func (o *P2pCreateDefault) WithStatusCode(code int) *P2pCreateDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the p2p create default response
func (o *P2pCreateDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *P2pCreateDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}