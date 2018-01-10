// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/subutai-io/agent/rest/models"
)

// ProxyCheckOKCode is the HTTP code returned for type ProxyCheckOK
const ProxyCheckOKCode int = 200

/*ProxyCheckOK Ok

swagger:response proxyCheckOK
*/
type ProxyCheckOK struct {

	/*
	  In: Body
	*/
	Payload models.ProxyCheckOKBody `json:"body,omitempty"`
}

// NewProxyCheckOK creates ProxyCheckOK with default headers values
func NewProxyCheckOK() *ProxyCheckOK {
	return &ProxyCheckOK{}
}

// WithPayload adds the payload to the proxy check o k response
func (o *ProxyCheckOK) WithPayload(payload models.ProxyCheckOKBody) *ProxyCheckOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the proxy check o k response
func (o *ProxyCheckOK) SetPayload(payload models.ProxyCheckOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ProxyCheckOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if payload == nil {
		payload = make(models.ProxyCheckOKBody, 0, 50)
	}

	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*ProxyCheckDefault generic error response

swagger:response proxyCheckDefault
*/
type ProxyCheckDefault struct {
	_statusCode int
}

// NewProxyCheckDefault creates ProxyCheckDefault with default headers values
func NewProxyCheckDefault(code int) *ProxyCheckDefault {
	if code <= 0 {
		code = 500
	}

	return &ProxyCheckDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the proxy check default response
func (o *ProxyCheckDefault) WithStatusCode(code int) *ProxyCheckDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the proxy check default response
func (o *ProxyCheckDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *ProxyCheckDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
