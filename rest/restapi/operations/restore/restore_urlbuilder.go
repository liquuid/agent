// Code generated by go-swagger; DO NOT EDIT.

package restore

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"errors"
	"net/url"
	golangswaggerpaths "path"
	"strings"
)

// RestoreURL generates an URL for the restore operation
type RestoreURL struct {
	Container string

	Date         *string
	Newcontainer *string

	_basePath string
	// avoid unkeyed usage
	_ struct{}
}

// WithBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RestoreURL) WithBasePath(bp string) *RestoreURL {
	o.SetBasePath(bp)
	return o
}

// SetBasePath sets the base path for this url builder, only required when it's different from the
// base path specified in the swagger spec.
// When the value of the base path is an empty string
func (o *RestoreURL) SetBasePath(bp string) {
	o._basePath = bp
}

// Build a url path and query string
func (o *RestoreURL) Build() (*url.URL, error) {
	var result url.URL

	var _path = "/restore/{container}"

	container := o.Container
	if container != "" {
		_path = strings.Replace(_path, "{container}", container, -1)
	} else {
		return nil, errors.New("Container is required on RestoreURL")
	}
	_basePath := o._basePath
	if _basePath == "" {
		_basePath = "/rest/v1/agent"
	}
	result.Path = golangswaggerpaths.Join(_basePath, _path)

	qs := make(url.Values)

	var date string
	if o.Date != nil {
		date = *o.Date
	}
	if date != "" {
		qs.Set("date", date)
	}

	var newcontainer string
	if o.Newcontainer != nil {
		newcontainer = *o.Newcontainer
	}
	if newcontainer != "" {
		qs.Set("newcontainer", newcontainer)
	}

	result.RawQuery = qs.Encode()

	return &result, nil
}

// Must is a helper function to panic when the url builder returns an error
func (o *RestoreURL) Must(u *url.URL, err error) *url.URL {
	if err != nil {
		panic(err)
	}
	if u == nil {
		panic("url can't be nil")
	}
	return u
}

// String returns the string representation of the path with query string
func (o *RestoreURL) String() string {
	return o.Must(o.Build()).String()
}

// BuildFull builds a full url with scheme, host, path and query string
func (o *RestoreURL) BuildFull(scheme, host string) (*url.URL, error) {
	if scheme == "" {
		return nil, errors.New("scheme is required for a full url on RestoreURL")
	}
	if host == "" {
		return nil, errors.New("host is required for a full url on RestoreURL")
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
func (o *RestoreURL) StringFull(scheme, host string) string {
	return o.Must(o.BuildFull(scheme, host)).String()
}
