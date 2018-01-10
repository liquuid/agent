// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/subutai-io/agent/rest/models"
)

// TunnelAddCreatedCode is the HTTP code returned for type TunnelAddCreated
const TunnelAddCreatedCode int = 201

/*TunnelAddCreated Created

swagger:response tunnelAddCreated
*/
type TunnelAddCreated struct {

	/*
	  In: Body
	*/
	Payload *models.Message `json:"body,omitempty"`
}

// NewTunnelAddCreated creates TunnelAddCreated with default headers values
func NewTunnelAddCreated() *TunnelAddCreated {
	return &TunnelAddCreated{}
}

// WithPayload adds the payload to the tunnel add created response
func (o *TunnelAddCreated) WithPayload(payload *models.Message) *TunnelAddCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the tunnel add created response
func (o *TunnelAddCreated) SetPayload(payload *models.Message) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *TunnelAddCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*TunnelAddDefault generic error response

swagger:response tunnelAddDefault
*/
type TunnelAddDefault struct {
	_statusCode int
}

// NewTunnelAddDefault creates TunnelAddDefault with default headers values
func NewTunnelAddDefault(code int) *TunnelAddDefault {
	if code <= 0 {
		code = 500
	}

	return &TunnelAddDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the tunnel add default response
func (o *TunnelAddDefault) WithStatusCode(code int) *TunnelAddDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the tunnel add default response
func (o *TunnelAddDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *TunnelAddDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
