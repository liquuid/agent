// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/subutai-io/agent/cli"
	"github.com/subutai-io/agent/rest/models"
	"github.com/subutai-io/agent/rest/restapi/operations"
	"github.com/subutai-io/agent/rest/restapi/operations/container"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../swagger.yml

func configureFlags(api *operations.AgentAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.AgentAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.CliListHandler = operations.CliListHandlerFunc(func(params operations.CliListParams) middleware.Responder {
		var containerOnlyFlag bool = false
		var templatesOnlyFlag bool = false
		var withAncestorsFlag bool = false
		var withParentFlag bool = false

		if params.ContainersOnly != nil {
			containerOnlyFlag = *params.ContainersOnly
		}

		if params.TemplatesOnly != nil {
			templatesOnlyFlag = *params.TemplatesOnly
		}

		if params.WithAncestors != nil {
			withAncestorsFlag = *params.WithAncestors
		}

		if params.WithParents != nil {
			withParentFlag = *params.WithParents
		}

		list := cli.LxcListJSON(containerOnlyFlag,
			templatesOnlyFlag,
			withAncestorsFlag,
			withParentFlag)

		okList := models.CliListOKBody{}

		for name, body := range list {

			okList = append(okList, &models.Container{body["ancestors"], name, body["parent"]})

		}

		return &operations.CliListOK{Payload: okList}

	})
	api.ContainerDestroyOneHandler = container.DestroyOneHandlerFunc(func(params container.DestroyOneParams) middleware.Responder {
		return middleware.NotImplemented("operation container.DestroyOne has not yet been implemented")
	})
	api.GetContainerInfoHandler = operations.GetContainerInfoHandlerFunc(func(params operations.GetContainerInfoParams) middleware.Responder {
		return middleware.NotImplemented("operation .GetContainerInfo has not yet been implemented")
	})

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *graceful.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
