// Code generated by go-swagger; DO NOT EDIT.

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	graceful "github.com/tylerb/graceful"

	"github.com/subutai-io/agent/rest/restapi/operations"
	"github.com/subutai-io/agent/rest/restapi/operations/attach"
	"github.com/subutai-io/agent/rest/restapi/operations/backup"
	"github.com/subutai-io/agent/rest/restapi/operations/batch"
	"github.com/subutai-io/agent/rest/restapi/operations/cleanup"
	"github.com/subutai-io/agent/rest/restapi/operations/clone"
	"github.com/subutai-io/agent/rest/restapi/operations/config"
	"github.com/subutai-io/agent/rest/restapi/operations/configuration"
	"github.com/subutai-io/agent/rest/restapi/operations/container"
	"github.com/subutai-io/agent/rest/restapi/operations/demote"
	"github.com/subutai-io/agent/rest/restapi/operations/destroy"
	"github.com/subutai-io/agent/rest/restapi/operations/export"
	"github.com/subutai-io/agent/rest/restapi/operations/hostname"
	"github.com/subutai-io/agent/rest/restapi/operations/import_operations"
	"github.com/subutai-io/agent/rest/restapi/operations/info"
	"github.com/subutai-io/agent/rest/restapi/operations/list"
	"github.com/subutai-io/agent/rest/restapi/operations/metrics"
	"github.com/subutai-io/agent/rest/restapi/operations/p2p"
	"github.com/subutai-io/agent/rest/restapi/operations/promote"
	"github.com/subutai-io/agent/rest/restapi/operations/proxy"
	"github.com/subutai-io/agent/rest/restapi/operations/quota"
	"github.com/subutai-io/agent/rest/restapi/operations/rename"
	"github.com/subutai-io/agent/rest/restapi/operations/resource_host"
	"github.com/subutai-io/agent/rest/restapi/operations/restore"
	"github.com/subutai-io/agent/rest/restapi/operations/start"
	"github.com/subutai-io/agent/rest/restapi/operations/stop"
	"github.com/subutai-io/agent/rest/restapi/operations/tunnel"
	"github.com/subutai-io/agent/rest/restapi/operations/update"
	"github.com/subutai-io/agent/rest/restapi/operations/vxlan"
)

// This file is safe to edit. Once it exists it will not be overwritten

//go:generate swagger generate server --target .. --name  --spec ../swagger.yml

func configureFlags(api *operations.HubAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.HubAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	api.AttachAttachHandler = attach.AttachHandlerFunc(func(params attach.AttachParams) middleware.Responder {
		return middleware.NotImplemented("operation attach.Attach has not yet been implemented")
	})
	api.BackupBackupContainerHandler = backup.BackupContainerHandlerFunc(func(params backup.BackupContainerParams) middleware.Responder {
		return middleware.NotImplemented("operation backup.BackupContainer has not yet been implemented")
	})
	api.BatchBatchHandler = batch.BatchHandlerFunc(func(params batch.BatchParams) middleware.Responder {
		return middleware.NotImplemented("operation batch.Batch has not yet been implemented")
	})
	api.CleanupCleanupHandler = cleanup.CleanupHandlerFunc(func(params cleanup.CleanupParams) middleware.Responder {
		return middleware.NotImplemented("operation cleanup.Cleanup has not yet been implemented")
	})
	api.ListCliListHandler = list.CliListHandlerFunc(func(params list.CliListParams) middleware.Responder {
		return middleware.NotImplemented("operation list.CliList has not yet been implemented")
	})
	api.CloneCloneHandler = clone.CloneHandlerFunc(func(params clone.CloneParams) middleware.Responder {
		return middleware.NotImplemented("operation clone.Clone has not yet been implemented")
	})
	api.ConfigurationConfigHandler = configuration.ConfigHandlerFunc(func(params configuration.ConfigParams) middleware.Responder {
		return middleware.NotImplemented("operation configuration.Config has not yet been implemented")
	})
	api.DemoteDemoteHandler = demote.DemoteHandlerFunc(func(params demote.DemoteParams) middleware.Responder {
		return middleware.NotImplemented("operation demote.Demote has not yet been implemented")
	})
	api.DestroyDestroyHandler = destroy.DestroyHandlerFunc(func(params destroy.DestroyParams) middleware.Responder {
		return middleware.NotImplemented("operation destroy.Destroy has not yet been implemented")
	})
	api.ConfigDestroyEntryHandler = config.DestroyEntryHandlerFunc(func(params config.DestroyEntryParams) middleware.Responder {
		return middleware.NotImplemented("operation config.DestroyEntry has not yet been implemented")
	})
	api.ContainerDestroyOneHandler = container.DestroyOneHandlerFunc(func(params container.DestroyOneParams) middleware.Responder {
		return middleware.NotImplemented("operation container.DestroyOne has not yet been implemented")
	})
	api.ExportExportHandler = export.ExportHandlerFunc(func(params export.ExportParams) middleware.Responder {
		return middleware.NotImplemented("operation export.Export has not yet been implemented")
	})
	api.ContainerGetContainerInfoHandler = container.GetContainerInfoHandlerFunc(func(params container.GetContainerInfoParams) middleware.Responder {
		return middleware.NotImplemented("operation container.GetContainerInfo has not yet been implemented")
	})
	api.HostnameHostnameHandler = hostname.HostnameHandlerFunc(func(params hostname.HostnameParams) middleware.Responder {
		return middleware.NotImplemented("operation hostname.Hostname has not yet been implemented")
	})
	api.ImportOperationsImportHandler = import_operations.ImportHandlerFunc(func(params import_operations.ImportParams) middleware.Responder {
		return middleware.NotImplemented("operation import_operations.Import has not yet been implemented")
	})
	api.InfoInfoHandler = info.InfoHandlerFunc(func(params info.InfoParams) middleware.Responder {
		return middleware.NotImplemented("operation info.Info has not yet been implemented")
	})
	api.MetricsMetricsHandler = metrics.MetricsHandlerFunc(func(params metrics.MetricsParams) middleware.Responder {
		return middleware.NotImplemented("operation metrics.Metrics has not yet been implemented")
	})
	api.P2pP2pCreateHandler = p2p.P2pCreateHandlerFunc(func(params p2p.P2pCreateParams) middleware.Responder {
		return middleware.NotImplemented("operation p2p.P2pCreate has not yet been implemented")
	})
	api.P2pP2pDeleteHandler = p2p.P2pDeleteHandlerFunc(func(params p2p.P2pDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation p2p.P2pDelete has not yet been implemented")
	})
	api.P2pP2pListHandler = p2p.P2pListHandlerFunc(func(params p2p.P2pListParams) middleware.Responder {
		return middleware.NotImplemented("operation p2p.P2pList has not yet been implemented")
	})
	api.P2pP2pUpdateHandler = p2p.P2pUpdateHandlerFunc(func(params p2p.P2pUpdateParams) middleware.Responder {
		return middleware.NotImplemented("operation p2p.P2pUpdate has not yet been implemented")
	})
	api.PromotePromoteHandler = promote.PromoteHandlerFunc(func(params promote.PromoteParams) middleware.Responder {
		return middleware.NotImplemented("operation promote.Promote has not yet been implemented")
	})
	api.ProxyProxyCheckHandler = proxy.ProxyCheckHandlerFunc(func(params proxy.ProxyCheckParams) middleware.Responder {
		return middleware.NotImplemented("operation proxy.ProxyCheck has not yet been implemented")
	})
	api.ProxyProxyCreateHandler = proxy.ProxyCreateHandlerFunc(func(params proxy.ProxyCreateParams) middleware.Responder {
		return middleware.NotImplemented("operation proxy.ProxyCreate has not yet been implemented")
	})
	api.ProxyProxyDeleteHandler = proxy.ProxyDeleteHandlerFunc(func(params proxy.ProxyDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation proxy.ProxyDelete has not yet been implemented")
	})
	api.QuotaQuotaHandler = quota.QuotaHandlerFunc(func(params quota.QuotaParams) middleware.Responder {
		return middleware.NotImplemented("operation quota.Quota has not yet been implemented")
	})
	api.RenameRenameHandler = rename.RenameHandlerFunc(func(params rename.RenameParams) middleware.Responder {
		return middleware.NotImplemented("operation rename.Rename has not yet been implemented")
	})
	api.RestoreRestoreHandler = restore.RestoreHandlerFunc(func(params restore.RestoreParams) middleware.Responder {
		return middleware.NotImplemented("operation restore.Restore has not yet been implemented")
	})
	api.ResourceHostRhIDHandler = resource_host.RhIDHandlerFunc(func(params resource_host.RhIDParams) middleware.Responder {
		return middleware.NotImplemented("operation resource_host.RhID has not yet been implemented")
	})
	api.StartStartHandler = start.StartHandlerFunc(func(params start.StartParams) middleware.Responder {
		return middleware.NotImplemented("operation start.Start has not yet been implemented")
	})
	api.StopStopHandler = stop.StopHandlerFunc(func(params stop.StopParams) middleware.Responder {
		return middleware.NotImplemented("operation stop.Stop has not yet been implemented")
	})
	api.TunnelTunnelAddHandler = tunnel.TunnelAddHandlerFunc(func(params tunnel.TunnelAddParams) middleware.Responder {
		return middleware.NotImplemented("operation tunnel.TunnelAdd has not yet been implemented")
	})
	api.TunnelTunnelCheckHandler = tunnel.TunnelCheckHandlerFunc(func(params tunnel.TunnelCheckParams) middleware.Responder {
		return middleware.NotImplemented("operation tunnel.TunnelCheck has not yet been implemented")
	})
	api.TunnelTunnelDeleteHandler = tunnel.TunnelDeleteHandlerFunc(func(params tunnel.TunnelDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation tunnel.TunnelDelete has not yet been implemented")
	})
	api.TunnelTunnelListHandler = tunnel.TunnelListHandlerFunc(func(params tunnel.TunnelListParams) middleware.Responder {
		return middleware.NotImplemented("operation tunnel.TunnelList has not yet been implemented")
	})
	api.UpdateUpdateHandler = update.UpdateHandlerFunc(func(params update.UpdateParams) middleware.Responder {
		return middleware.NotImplemented("operation update.Update has not yet been implemented")
	})
	api.VxlanVxlanCreateHandler = vxlan.VxlanCreateHandlerFunc(func(params vxlan.VxlanCreateParams) middleware.Responder {
		return middleware.NotImplemented("operation vxlan.VxlanCreate has not yet been implemented")
	})
	api.VxlanVxlanDeleteHandler = vxlan.VxlanDeleteHandlerFunc(func(params vxlan.VxlanDeleteParams) middleware.Responder {
		return middleware.NotImplemented("operation vxlan.VxlanDelete has not yet been implemented")
	})
	api.VxlanVxlanListHandler = vxlan.VxlanListHandlerFunc(func(params vxlan.VxlanListParams) middleware.Responder {
		return middleware.NotImplemented("operation vxlan.VxlanList has not yet been implemented")
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
