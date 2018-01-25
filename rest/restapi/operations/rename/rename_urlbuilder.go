// Code generated by go-swagger; DO NOT EDIT.

package rename

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// RenameURL generates an URL for the rename operation
type RenameURL struct {
	Newname string
	Source  string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RenameURL) WithBasePath(bp string) *RenameURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RenameURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *RenameURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/rename/{source}/{newname}"

	newname := o.Newname
	if newname != "" {
		_path = strings.Replace(_path, "{newname}", newname, -1)
	} else {
		return nil, errors.New("Newname is required on RenameURL")
	}
	source := o.Source
	if source != "" {
		_path = strings.Replace(_path, "{source}", source, -1)
	} else {
		return nil, errors.New("Source is required on RenameURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/rest/v1/agent"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *RenameURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *RenameURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *RenameURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on RenameURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on RenameURL")
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
func (o *RenameURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}