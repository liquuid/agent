// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/subutai-io/agent/rest/models"
)

// CloneOKCode is the HTTP code returned for type CloneOK
const CloneOKCode int = 200

/*CloneOK OK

swagger:response cloneOK
*/
type CloneOK struct {

	/*
	  In: Body
	*/
	Payload *models.Message `json:"body,omitempty"`
}

// NewCloneOK creates CloneOK with default headers values
func NewCloneOK() *CloneOK {
	return &CloneOK{}
}

// WithPayload adds the payload to the clone o k response
func (o *CloneOK) WithPayload(payload *models.Message) *CloneOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the clone o k response
func (o *CloneOK) SetPayload(payload *models.Message) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CloneOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*CloneDefault error

swagger:response cloneDefault
*/
type CloneDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewCloneDefault creates CloneDefault with default headers values
func NewCloneDefault(code int) *CloneDefault {
	if code <= 0 {
		code = 500
	}

	return &CloneDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the clone default response
func (o *CloneDefault) WithStatusCode(code int) *CloneDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the clone default response
func (o *CloneDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the clone default response
func (o *CloneDefault) WithPayload(payload *models.Error) *CloneDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the clone default response
func (o *CloneDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CloneDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
