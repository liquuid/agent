// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"net/http"
	"strings"

	errors "github.com/go-openapi/errors"
	loads "github.com/go-openapi/loads"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"
	security "github.com/go-openapi/runtime/security"
	spec "github.com/go-openapi/spec"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/subutai-io/agent/rest/restapi/operations/config"
	"github.com/subutai-io/agent/rest/restapi/operations/container"
)

// NewAgentAPI creates a new Agent instance
func NewAgentAPI(spec *loads.Document) *AgentAPI {
	return &AgentAPI{
		handlers:            make(map[string]map[string]http.Handler),
		formats:             strfmt.Default,
		defaultConsumes:     "application/json",
		defaultProduces:     "application/json",
		ServerShutdown:      func() {},
		spec:                spec,
		ServeError:          errors.ServeError,
		BasicAuthenticator:  security.BasicAuth,
		APIKeyAuthenticator: security.APIKeyAuth,
		BearerAuthenticator: security.BearerAuth,
		JSONConsumer:        runtime.JSONConsumer(),
		JSONProducer:        runtime.JSONProducer(),
		BackupContainerHandler: BackupContainerHandlerFunc(func(params BackupContainerParams) middleware.Responder {
			return middleware.NotImplemented("operation BackupContainer has not yet been implemented")
		}),
		BatchHandler: BatchHandlerFunc(func(params BatchParams) middleware.Responder {
			return middleware.NotImplemented("operation Batch has not yet been implemented")
		}),
		CleanupHandler: CleanupHandlerFunc(func(params CleanupParams) middleware.Responder {
			return middleware.NotImplemented("operation Cleanup has not yet been implemented")
		}),
		CliListHandler: CliListHandlerFunc(func(params CliListParams) middleware.Responder {
			return middleware.NotImplemented("operation CliList has not yet been implemented")
		}),
		CloneHandler: CloneHandlerFunc(func(params CloneParams) middleware.Responder {
			return middleware.NotImplemented("operation Clone has not yet been implemented")
		}),
		ConfigHandler: ConfigHandlerFunc(func(params ConfigParams) middleware.Responder {
			return middleware.NotImplemented("operation Config has not yet been implemented")
		}),
		DemoteHandler: DemoteHandlerFunc(func(params DemoteParams) middleware.Responder {
			return middleware.NotImplemented("operation Demote has not yet been implemented")
		}),
		DestroyHandler: DestroyHandlerFunc(func(params DestroyParams) middleware.Responder {
			return middleware.NotImplemented("operation Destroy has not yet been implemented")
		}),
		ConfigDestroyEntryHandler: config.DestroyEntryHandlerFunc(func(params config.DestroyEntryParams) middleware.Responder {
			return middleware.NotImplemented("operation ConfigDestroyEntry has not yet been implemented")
		}),
		ContainerDestroyOneHandler: container.DestroyOneHandlerFunc(func(params container.DestroyOneParams) middleware.Responder {
			return middleware.NotImplemented("operation ContainerDestroyOne has not yet been implemented")
		}),
		ExportHandler: ExportHandlerFunc(func(params ExportParams) middleware.Responder {
			return middleware.NotImplemented("operation Export has not yet been implemented")
		}),
		GetContainerInfoHandler: GetContainerInfoHandlerFunc(func(params GetContainerInfoParams) middleware.Responder {
			return middleware.NotImplemented("operation GetContainerInfo has not yet been implemented")
		}),
		HostnameHandler: HostnameHandlerFunc(func(params HostnameParams) middleware.Responder {
			return middleware.NotImplemented("operation Hostname has not yet been implemented")
		}),
		ImportHandler: ImportHandlerFunc(func(params ImportParams) middleware.Responder {
			return middleware.NotImplemented("operation Import has not yet been implemented")
		}),
		InfoHandler: InfoHandlerFunc(func(params InfoParams) middleware.Responder {
			return middleware.NotImplemented("operation Info has not yet been implemented")
		}),
		MetricsHandler: MetricsHandlerFunc(func(params MetricsParams) middleware.Responder {
			return middleware.NotImplemented("operation Metrics has not yet been implemented")
		}),
		P2pCreateHandler: P2pCreateHandlerFunc(func(params P2pCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation P2pCreate has not yet been implemented")
		}),
		P2pDeleteHandler: P2pDeleteHandlerFunc(func(params P2pDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation P2pDelete has not yet been implemented")
		}),
		P2pListHandler: P2pListHandlerFunc(func(params P2pListParams) middleware.Responder {
			return middleware.NotImplemented("operation P2pList has not yet been implemented")
		}),
		P2pUpdateHandler: P2pUpdateHandlerFunc(func(params P2pUpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation P2pUpdate has not yet been implemented")
		}),
		PromoteHandler: PromoteHandlerFunc(func(params PromoteParams) middleware.Responder {
			return middleware.NotImplemented("operation Promote has not yet been implemented")
		}),
		ProxyCheckHandler: ProxyCheckHandlerFunc(func(params ProxyCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation ProxyCheck has not yet been implemented")
		}),
		ProxyCreateHandler: ProxyCreateHandlerFunc(func(params ProxyCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation ProxyCreate has not yet been implemented")
		}),
		ProxyDeleteHandler: ProxyDeleteHandlerFunc(func(params ProxyDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation ProxyDelete has not yet been implemented")
		}),
		QuotaHandler: QuotaHandlerFunc(func(params QuotaParams) middleware.Responder {
			return middleware.NotImplemented("operation Quota has not yet been implemented")
		}),
		RenameHandler: RenameHandlerFunc(func(params RenameParams) middleware.Responder {
			return middleware.NotImplemented("operation Rename has not yet been implemented")
		}),
		RestoreHandler: RestoreHandlerFunc(func(params RestoreParams) middleware.Responder {
			return middleware.NotImplemented("operation Restore has not yet been implemented")
		}),
		RhIDHandler: RhIDHandlerFunc(func(params RhIDParams) middleware.Responder {
			return middleware.NotImplemented("operation RhID has not yet been implemented")
		}),
		StartHandler: StartHandlerFunc(func(params StartParams) middleware.Responder {
			return middleware.NotImplemented("operation Start has not yet been implemented")
		}),
		StopHandler: StopHandlerFunc(func(params StopParams) middleware.Responder {
			return middleware.NotImplemented("operation Stop has not yet been implemented")
		}),
		TunnelAddHandler: TunnelAddHandlerFunc(func(params TunnelAddParams) middleware.Responder {
			return middleware.NotImplemented("operation TunnelAdd has not yet been implemented")
		}),
		TunnelCheckHandler: TunnelCheckHandlerFunc(func(params TunnelCheckParams) middleware.Responder {
			return middleware.NotImplemented("operation TunnelCheck has not yet been implemented")
		}),
		TunnelDeleteHandler: TunnelDeleteHandlerFunc(func(params TunnelDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation TunnelDelete has not yet been implemented")
		}),
		TunnelListHandler: TunnelListHandlerFunc(func(params TunnelListParams) middleware.Responder {
			return middleware.NotImplemented("operation TunnelList has not yet been implemented")
		}),
		UpdateHandler: UpdateHandlerFunc(func(params UpdateParams) middleware.Responder {
			return middleware.NotImplemented("operation Update has not yet been implemented")
		}),
		VxlanCreateHandler: VxlanCreateHandlerFunc(func(params VxlanCreateParams) middleware.Responder {
			return middleware.NotImplemented("operation VxlanCreate has not yet been implemented")
		}),
		VxlanDeleteHandler: VxlanDeleteHandlerFunc(func(params VxlanDeleteParams) middleware.Responder {
			return middleware.NotImplemented("operation VxlanDelete has not yet been implemented")
		}),
		VxlanListHandler: VxlanListHandlerFunc(func(params VxlanListParams) middleware.Responder {
			return middleware.NotImplemented("operation VxlanList has not yet been implemented")
		}),
	}
}

/*AgentAPI API to handle agent
 */
type AgentAPI struct {
	spec            *loads.Document
	context         *middleware.Context
	handlers        map[string]map[string]http.Handler
	formats         strfmt.Registry
	defaultConsumes string
	defaultProduces string
	Middleware      func(middleware.Builder) http.Handler

	// BasicAuthenticator generates a runtime.Authenticator from the supplied basic auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BasicAuthenticator func(security.UserPassAuthentication) runtime.Authenticator
	// APIKeyAuthenticator generates a runtime.Authenticator from the supplied token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	APIKeyAuthenticator func(string, string, security.TokenAuthentication) runtime.Authenticator
	// BearerAuthenticator generates a runtime.Authenticator from the supplied bearer token auth function.
	// It has a default implemention in the security package, however you can replace it for your particular usage.
	BearerAuthenticator func(string, security.ScopedTokenAuthentication) runtime.Authenticator

	// JSONConsumer registers a consumer for a "application/json" mime type
	JSONConsumer runtime.Consumer

	// JSONProducer registers a producer for a "application/json" mime type
	JSONProducer runtime.Producer

	// BackupContainerHandler sets the operation handler for the backup container operation
	BackupContainerHandler BackupContainerHandler
	// BatchHandler sets the operation handler for the batch operation
	BatchHandler BatchHandler
	// CleanupHandler sets the operation handler for the cleanup operation
	CleanupHandler CleanupHandler
	// CliListHandler sets the operation handler for the cli list operation
	CliListHandler CliListHandler
	// CloneHandler sets the operation handler for the clone operation
	CloneHandler CloneHandler
	// ConfigHandler sets the operation handler for the config operation
	ConfigHandler ConfigHandler
	// DemoteHandler sets the operation handler for the demote operation
	DemoteHandler DemoteHandler
	// DestroyHandler sets the operation handler for the destroy operation
	DestroyHandler DestroyHandler
	// ConfigDestroyEntryHandler sets the operation handler for the destroy entry operation
	ConfigDestroyEntryHandler config.DestroyEntryHandler
	// ContainerDestroyOneHandler sets the operation handler for the destroy one operation
	ContainerDestroyOneHandler container.DestroyOneHandler
	// ExportHandler sets the operation handler for the export operation
	ExportHandler ExportHandler
	// GetContainerInfoHandler sets the operation handler for the get container info operation
	GetContainerInfoHandler GetContainerInfoHandler
	// HostnameHandler sets the operation handler for the hostname operation
	HostnameHandler HostnameHandler
	// ImportHandler sets the operation handler for the import operation
	ImportHandler ImportHandler
	// InfoHandler sets the operation handler for the info operation
	InfoHandler InfoHandler
	// MetricsHandler sets the operation handler for the metrics operation
	MetricsHandler MetricsHandler
	// P2pCreateHandler sets the operation handler for the p2p create operation
	P2pCreateHandler P2pCreateHandler
	// P2pDeleteHandler sets the operation handler for the p2p delete operation
	P2pDeleteHandler P2pDeleteHandler
	// P2pListHandler sets the operation handler for the p2p list operation
	P2pListHandler P2pListHandler
	// P2pUpdateHandler sets the operation handler for the p2p update operation
	P2pUpdateHandler P2pUpdateHandler
	// PromoteHandler sets the operation handler for the promote operation
	PromoteHandler PromoteHandler
	// ProxyCheckHandler sets the operation handler for the proxy check operation
	ProxyCheckHandler ProxyCheckHandler
	// ProxyCreateHandler sets the operation handler for the proxy create operation
	ProxyCreateHandler ProxyCreateHandler
	// ProxyDeleteHandler sets the operation handler for the proxy delete operation
	ProxyDeleteHandler ProxyDeleteHandler
	// QuotaHandler sets the operation handler for the quota operation
	QuotaHandler QuotaHandler
	// RenameHandler sets the operation handler for the rename operation
	RenameHandler RenameHandler
	// RestoreHandler sets the operation handler for the restore operation
	RestoreHandler RestoreHandler
	// RhIDHandler sets the operation handler for the rh ID operation
	RhIDHandler RhIDHandler
	// StartHandler sets the operation handler for the start operation
	StartHandler StartHandler
	// StopHandler sets the operation handler for the stop operation
	StopHandler StopHandler
	// TunnelAddHandler sets the operation handler for the tunnel add operation
	TunnelAddHandler TunnelAddHandler
	// TunnelCheckHandler sets the operation handler for the tunnel check operation
	TunnelCheckHandler TunnelCheckHandler
	// TunnelDeleteHandler sets the operation handler for the tunnel delete operation
	TunnelDeleteHandler TunnelDeleteHandler
	// TunnelListHandler sets the operation handler for the tunnel list operation
	TunnelListHandler TunnelListHandler
	// UpdateHandler sets the operation handler for the update operation
	UpdateHandler UpdateHandler
	// VxlanCreateHandler sets the operation handler for the vxlan create operation
	VxlanCreateHandler VxlanCreateHandler
	// VxlanDeleteHandler sets the operation handler for the vxlan delete operation
	VxlanDeleteHandler VxlanDeleteHandler
	// VxlanListHandler sets the operation handler for the vxlan list operation
	VxlanListHandler VxlanListHandler

	// ServeError is called when an error is received, there is a default handler
	// but you can set your own with this
	ServeError func(http.ResponseWriter, *http.Request, error)

	// ServerShutdown is called when the HTTP(S) server is shut down and done
	// handling all active connections and does not accept connections any more
	ServerShutdown func()

	// Custom command line argument groups with their descriptions
	CommandLineOptionsGroups []swag.CommandLineOptionsGroup

	// User defined logger function.
	Logger func(string, ...interface{})
}

// SetDefaultProduces sets the default produces media type
func (o *AgentAPI) SetDefaultProduces(mediaType string) {
	o.defaultProduces = mediaType
}

// SetDefaultConsumes returns the default consumes media type
func (o *AgentAPI) SetDefaultConsumes(mediaType string) {
	o.defaultConsumes = mediaType
}

// SetSpec sets a spec that will be served for the clients.
func (o *AgentAPI) SetSpec(spec *loads.Document) {
	o.spec = spec
}

// DefaultProduces returns the default produces media type
func (o *AgentAPI) DefaultProduces() string {
	return o.defaultProduces
}

// DefaultConsumes returns the default consumes media type
func (o *AgentAPI) DefaultConsumes() string {
	return o.defaultConsumes
}

// Formats returns the registered string formats
func (o *AgentAPI) Formats() strfmt.Registry {
	return o.formats
}

// RegisterFormat registers a custom format validator
func (o *AgentAPI) RegisterFormat(name string, format strfmt.Format, validator strfmt.Validator) {
	o.formats.Add(name, format, validator)
}

// Validate validates the registrations in the AgentAPI
func (o *AgentAPI) Validate() error {
	var unregistered []string

	if o.JSONConsumer == nil {
		unregistered = append(unregistered, "JSONConsumer")
	}

	if o.JSONProducer == nil {
		unregistered = append(unregistered, "JSONProducer")
	}

	if o.BackupContainerHandler == nil {
		unregistered = append(unregistered, "BackupContainerHandler")
	}

	if o.BatchHandler == nil {
		unregistered = append(unregistered, "BatchHandler")
	}

	if o.CleanupHandler == nil {
		unregistered = append(unregistered, "CleanupHandler")
	}

	if o.CliListHandler == nil {
		unregistered = append(unregistered, "CliListHandler")
	}

	if o.CloneHandler == nil {
		unregistered = append(unregistered, "CloneHandler")
	}

	if o.ConfigHandler == nil {
		unregistered = append(unregistered, "ConfigHandler")
	}

	if o.DemoteHandler == nil {
		unregistered = append(unregistered, "DemoteHandler")
	}

	if o.DestroyHandler == nil {
		unregistered = append(unregistered, "DestroyHandler")
	}

	if o.ConfigDestroyEntryHandler == nil {
		unregistered = append(unregistered, "config.DestroyEntryHandler")
	}

	if o.ContainerDestroyOneHandler == nil {
		unregistered = append(unregistered, "container.DestroyOneHandler")
	}

	if o.ExportHandler == nil {
		unregistered = append(unregistered, "ExportHandler")
	}

	if o.GetContainerInfoHandler == nil {
		unregistered = append(unregistered, "GetContainerInfoHandler")
	}

	if o.HostnameHandler == nil {
		unregistered = append(unregistered, "HostnameHandler")
	}

	if o.ImportHandler == nil {
		unregistered = append(unregistered, "ImportHandler")
	}

	if o.InfoHandler == nil {
		unregistered = append(unregistered, "InfoHandler")
	}

	if o.MetricsHandler == nil {
		unregistered = append(unregistered, "MetricsHandler")
	}

	if o.P2pCreateHandler == nil {
		unregistered = append(unregistered, "P2pCreateHandler")
	}

	if o.P2pDeleteHandler == nil {
		unregistered = append(unregistered, "P2pDeleteHandler")
	}

	if o.P2pListHandler == nil {
		unregistered = append(unregistered, "P2pListHandler")
	}

	if o.P2pUpdateHandler == nil {
		unregistered = append(unregistered, "P2pUpdateHandler")
	}

	if o.PromoteHandler == nil {
		unregistered = append(unregistered, "PromoteHandler")
	}

	if o.ProxyCheckHandler == nil {
		unregistered = append(unregistered, "ProxyCheckHandler")
	}

	if o.ProxyCreateHandler == nil {
		unregistered = append(unregistered, "ProxyCreateHandler")
	}

	if o.ProxyDeleteHandler == nil {
		unregistered = append(unregistered, "ProxyDeleteHandler")
	}

	if o.QuotaHandler == nil {
		unregistered = append(unregistered, "QuotaHandler")
	}

	if o.RenameHandler == nil {
		unregistered = append(unregistered, "RenameHandler")
	}

	if o.RestoreHandler == nil {
		unregistered = append(unregistered, "RestoreHandler")
	}

	if o.RhIDHandler == nil {
		unregistered = append(unregistered, "RhIDHandler")
	}

	if o.StartHandler == nil {
		unregistered = append(unregistered, "StartHandler")
	}

	if o.StopHandler == nil {
		unregistered = append(unregistered, "StopHandler")
	}

	if o.TunnelAddHandler == nil {
		unregistered = append(unregistered, "TunnelAddHandler")
	}

	if o.TunnelCheckHandler == nil {
		unregistered = append(unregistered, "TunnelCheckHandler")
	}

	if o.TunnelDeleteHandler == nil {
		unregistered = append(unregistered, "TunnelDeleteHandler")
	}

	if o.TunnelListHandler == nil {
		unregistered = append(unregistered, "TunnelListHandler")
	}

	if o.UpdateHandler == nil {
		unregistered = append(unregistered, "UpdateHandler")
	}

	if o.VxlanCreateHandler == nil {
		unregistered = append(unregistered, "VxlanCreateHandler")
	}

	if o.VxlanDeleteHandler == nil {
		unregistered = append(unregistered, "VxlanDeleteHandler")
	}

	if o.VxlanListHandler == nil {
		unregistered = append(unregistered, "VxlanListHandler")
	}

	if len(unregistered) > 0 {
		return fmt.Errorf("missing registration: %s", strings.Join(unregistered, ", "))
	}

	return nil
}

// ServeErrorFor gets a error handler for a given operation id
func (o *AgentAPI) ServeErrorFor(operationID string) func(http.ResponseWriter, *http.Request, error) {
	return o.ServeError
}

// AuthenticatorsFor gets the authenticators for the specified security schemes
func (o *AgentAPI) AuthenticatorsFor(schemes map[string]spec.SecurityScheme) map[string]runtime.Authenticator {

	return nil

}

// Authorizer returns the registered authorizer
func (o *AgentAPI) Authorizer() runtime.Authorizer {

	return nil

}

// ConsumersFor gets the consumers for the specified media types
func (o *AgentAPI) ConsumersFor(mediaTypes []string) map[string]runtime.Consumer {

	result := make(map[string]runtime.Consumer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONConsumer

		}
	}
	return result

}

// ProducersFor gets the producers for the specified media types
func (o *AgentAPI) ProducersFor(mediaTypes []string) map[string]runtime.Producer {

	result := make(map[string]runtime.Producer)
	for _, mt := range mediaTypes {
		switch mt {

		case "application/json":
			result["application/json"] = o.JSONProducer

		}
	}
	return result

}

// HandlerFor gets a http.Handler for the provided operation method and path
func (o *AgentAPI) HandlerFor(method, path string) (http.Handler, bool) {
	if o.handlers == nil {
		return nil, false
	}
	um := strings.ToUpper(method)
	if _, ok := o.handlers[um]; !ok {
		return nil, false
	}
	if path == "/" {
		path = ""
	}
	h, ok := o.handlers[um][path]
	return h, ok
}

// Context returns the middleware context for the agent API
func (o *AgentAPI) Context() *middleware.Context {
	if o.context == nil {
		o.context = middleware.NewRoutableContext(o.spec, o, nil)
	}

	return o.context
}

func (o *AgentAPI) initHandlerCache() {
	o.Context() // don't care about the result, just that the initialization happened

	if o.handlers == nil {
		o.handlers = make(map[string]map[string]http.Handler)
	}

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/backup/{name}"] = NewBackupContainer(o.context, o.BackupContainerHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/rest/v1/batch"] = NewBatch(o.context, o.BatchHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/cleanup"] = NewCleanup(o.context, o.CleanupHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/list"] = NewCliList(o.context, o.CliListHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/rest/v1/clone/{parent}/{child}"] = NewClone(o.context, o.CloneHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/rest/v1/config"] = NewConfig(o.context, o.ConfigHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/rest/v1/demote/{container}"] = NewDemote(o.context, o.DemoteHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/destroy/{ID}"] = NewDestroy(o.context, o.DestroyHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/rest/v1/config"] = config.NewDestroyEntry(o.context, o.ConfigDestroyEntryHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/rest/v1/container/{name}"] = container.NewDestroyOne(o.context, o.ContainerDestroyOneHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/rest/v1/export/{container}"] = NewExport(o.context, o.ExportHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/container/{name}"] = NewGetContainerInfo(o.context, o.GetContainerInfoHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/rest/v1/hostname/{container}/{name}"] = NewHostname(o.context, o.HostnameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/import/{container}"] = NewImport(o.context, o.ImportHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/info"] = NewInfo(o.context, o.InfoHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/metrics"] = NewMetrics(o.context, o.MetricsHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/rest/v1/p2p"] = NewP2pCreate(o.context, o.P2pCreateHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/rest/v1/p2p"] = NewP2pDelete(o.context, o.P2pDeleteHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/p2p"] = NewP2pList(o.context, o.P2pListHandler)

	if o.handlers["PUT"] == nil {
		o.handlers["PUT"] = make(map[string]http.Handler)
	}
	o.handlers["PUT"]["/rest/v1/p2p"] = NewP2pUpdate(o.context, o.P2pUpdateHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/rest/v1/promote/{container}"] = NewPromote(o.context, o.PromoteHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/proxy"] = NewProxyCheck(o.context, o.ProxyCheckHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/rest/v1/proxy"] = NewProxyCreate(o.context, o.ProxyCreateHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/rest/v1/proxy"] = NewProxyDelete(o.context, o.ProxyDeleteHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/quota"] = NewQuota(o.context, o.QuotaHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/rename/{source}/{newname}"] = NewRename(o.context, o.RenameHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/restore/{container}"] = NewRestore(o.context, o.RestoreHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/rh/id"] = NewRhID(o.context, o.RhIDHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/start/{container}"] = NewStart(o.context, o.StartHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/stop/{container}"] = NewStop(o.context, o.StopHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/rest/v1/tunnel/{socket}"] = NewTunnelAdd(o.context, o.TunnelAddHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/tunnel/check"] = NewTunnelCheck(o.context, o.TunnelCheckHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/rest/v1/tunnel/{socket}"] = NewTunnelDelete(o.context, o.TunnelDeleteHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/tunnel/list"] = NewTunnelList(o.context, o.TunnelListHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/update/{container}"] = NewUpdate(o.context, o.UpdateHandler)

	if o.handlers["POST"] == nil {
		o.handlers["POST"] = make(map[string]http.Handler)
	}
	o.handlers["POST"]["/rest/v1/vxlan/{tunnel}"] = NewVxlanCreate(o.context, o.VxlanCreateHandler)

	if o.handlers["DELETE"] == nil {
		o.handlers["DELETE"] = make(map[string]http.Handler)
	}
	o.handlers["DELETE"]["/rest/v1/vxlan/{iface}"] = NewVxlanDelete(o.context, o.VxlanDeleteHandler)

	if o.handlers["GET"] == nil {
		o.handlers["GET"] = make(map[string]http.Handler)
	}
	o.handlers["GET"]["/rest/v1/vxlan/list"] = NewVxlanList(o.context, o.VxlanListHandler)

}

// Serve creates a http handler to serve the API over HTTP
// can be used directly in http.ListenAndServe(":8000", api.Serve(nil))
func (o *AgentAPI) Serve(builder middleware.Builder) http.Handler {
	o.Init()

	if o.Middleware != nil {
		return o.Middleware(builder)
	}
	return o.context.APIHandler(builder)
}

// Init allows you to just initialize the handler cache, you can then recompose the middelware as you see fit
func (o *AgentAPI) Init() {
	if len(o.handlers) == 0 {
		o.initHandlerCache()
	}
}
