// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/subutai-io/agent/rest/models"
)

// P2pPostOKCode is the HTTP code returned for type P2pPostOK
const P2pPostOKCode int = 200

/*P2pPostOK Ok

swagger:response p2pPostOK
*/
type P2pPostOK struct {

	/*
	  In: Body
	*/
	Payload models.P2pPostOKBody `json:"body,omitempty"`
}

// NewP2pPostOK creates P2pPostOK with default headers values
func NewP2pPostOK() *P2pPostOK {
	return &P2pPostOK{}
}

// WithPayload adds the payload to the p2p post o k response
func (o *P2pPostOK) WithPayload(payload models.P2pPostOKBody) *P2pPostOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the p2p post o k response
func (o *P2pPostOK) SetPayload(payload models.P2pPostOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *P2pPostOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.P2pPostOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*P2pPostDefault generic error response

swagger:response p2pPostDefault
*/
type P2pPostDefault struct {
	_statusCode int
}

// NewP2pPostDefault creates P2pPostDefault with default headers values
func NewP2pPostDefault(code int) *P2pPostDefault {
	if code <= 0 {
		code = 500
	}

	return &P2pPostDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the p2p post default response
func (o *P2pPostDefault) WithStatusCode(code int) *P2pPostDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the p2p post default response
func (o *P2pPostDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *P2pPostDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
