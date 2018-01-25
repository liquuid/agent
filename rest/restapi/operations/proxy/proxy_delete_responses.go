// Code generated by go-swagger; DO NOT EDIT.

package proxy

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// ProxyDeleteNoContentCode is the HTTP code returned for type ProxyDeleteNoContent
const ProxyDeleteNoContentCode int = 204

/*ProxyDeleteNoContent Deleted

swagger:response proxyDeleteNoContent
*/
type ProxyDeleteNoContent struct {
}

// NewProxyDeleteNoContent creates ProxyDeleteNoContent with default headers values
func NewProxyDeleteNoContent() *ProxyDeleteNoContent {
	return &ProxyDeleteNoContent{}
}

// WriteResponse to the client
func (o *ProxyDeleteNoContent) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(204)
}

/*ProxyDeleteDefault generic error response

swagger:response proxyDeleteDefault
*/
type ProxyDeleteDefault struct {
	_statusCode int
}

// NewProxyDeleteDefault creates ProxyDeleteDefault with default headers values
func NewProxyDeleteDefault(code int) *ProxyDeleteDefault {
	if code <= 0 {
		code = 500
	}

	return &ProxyDeleteDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the proxy delete default response
func (o *ProxyDeleteDefault) WithStatusCode(code int) *ProxyDeleteDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the proxy delete default response
func (o *ProxyDeleteDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *ProxyDeleteDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}