// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/subutai-io/agent/rest/models"
)

// StartOKCode is the HTTP code returned for type StartOK
const StartOKCode int = 200

/*StartOK Ok

swagger:response startOK
*/
type StartOK struct {

	/*
	  In: Body
	*/
	Payload models.StartOKBody `json:"body,omitempty"`
}

// NewStartOK creates StartOK with default headers values
func NewStartOK() *StartOK {
	return &StartOK{}
}

// WithPayload adds the payload to the start o k response
func (o *StartOK) WithPayload(payload models.StartOKBody) *StartOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the start o k response
func (o *StartOK) SetPayload(payload models.StartOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *StartOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.StartOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*StartDefault generic error response

swagger:response startDefault
*/
type StartDefault struct {
	_statusCode int
}

// NewStartDefault creates StartDefault with default headers values
func NewStartDefault(code int) *StartDefault {
	if code <= 0 {
		code = 500
	}

	return &StartDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the start default response
func (o *StartDefault) WithStatusCode(code int) *StartDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the start default response
func (o *StartDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *StartDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
