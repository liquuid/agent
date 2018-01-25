// Code generated by go-swagger; DO NOT EDIT.

package tunnel

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// TunnelDeleteURL generates an URL for the tunnel delete operation
type TunnelDeleteURL struct {
	Socket string

	Pid *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *TunnelDeleteURL) WithBasePath(bp string) *TunnelDeleteURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *TunnelDeleteURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *TunnelDeleteURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/tunnel/{socket}"

	socket := o.Socket
	if socket != "" {
		_path = strings.Replace(_path, "{socket}", socket, -1)
	} else {
		return nil, errors.New("Socket is required on TunnelDeleteURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/rest/v1/agent"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var pid string
	if o.Pid != nil {
		pid = *o.Pid
	}
	if pid != "" {
		qs.Set("pid", pid)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *TunnelDeleteURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *TunnelDeleteURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *TunnelDeleteURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on TunnelDeleteURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on TunnelDeleteURL")
	}

	base, err := o.Build()
	if err != nil {
		return nil, err
	}

	base.Scheme = scheme
	base.Host = host
	return base, nil
}

// StringFull returns the string representation of a complete url
func (o *TunnelDeleteURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}