// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

// SwaggerJSON embedded version of the swagger document used at generation time
var SwaggerJSON json.RawMessage

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API to handle agent\n",
    "title": "Agent API",
    "version": "0.0.1"
  },
  "host": "localhost:8080",
  "basePath": "/",
  "paths": {
    "/rest/v1/backup/{name}": {
      "get": {
        "description": "BackupContainer takes a snapshots of each container's volume and stores it in the ` + "`" + `/mnt/backups/container_name/datetime/` + "`" + ` directory. A full backup creates a delta-file of each BTRFS subvolume. An incremental backup (default) creates a delta-file with the difference of changes between the current and last snapshots. All deltas are compressed to archives in ` + "`" + `/mnt/backups/` + "`" + ` directory (container_datetime.tar.gz or container_datetime_Full.tar.gz for full backup). A changelog file can be found next to backups archive (container_datetime_changelog.txt or container_datetime_Full_changelog.txt) which contains a list of changes made between two backups.",
        "tags": [
          "agent",
          "cli",
          "container"
        ],
        "operationId": "backupContainer",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "name",
          "in": "path",
          "required": true
        },
        {
          "type": "boolean",
          "name": "full",
          "in": "query"
        },
        {
          "type": "boolean",
          "name": "stop",
          "in": "query"
        }
      ]
    },
    "/rest/v1/batch": {
      "post": {
        "description": "Batch binding provides a mechanism to perform several Subutai commands in the container in batch, passed in a single JSON message. Initially, the purpose of this command was internal for SS \u003c-\u003e Agent communication, yet it may be invoked manually from the CLI. The response from a batch command returns a JSON array with each element representing the results (response) from each command (request) in the batch: the positions of responses correlate with the request position in the array",
        "tags": [
          "agent",
          "cli",
          "container"
        ],
        "operationId": "batch",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/batchLine"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/rest/v1/cleanup": {
      "get": {
        "description": "Cleanup simply removes every resource associated with a Subutai container or template: data, network, configs, etc. The destroy command always runs each step in \"force\" mode to provide reliable deletion results; even if some instance components were already removed, the destroy command will continue to perform all operations once again while ignoring possible underlying errors: i.e. missing configuration files.",
        "tags": [
          "agent",
          "cli",
          "container"
        ],
        "operationId": "cleanup",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/rest/v1/clone/{parent}/{child}": {
      "post": {
        "description": "Clone function creates new ` + "`" + `child` + "`" + ` container from a Subutai ` + "`" + `parent` + "`" + ` template. If the specified template argument is not deployed in system, Subutai first tries to import it, and if import succeeds, it then continues to clone from the imported template image. By default, clone will use the NAT-ed network interface with IP address received from the Subutai DHCP server, but this behavior can be changed with command options described below.\nIf ` + "`" + `ipaddr` + "`" + ` option is defined, separate bridge interface will be created in specified VLAN and new container will receive static IP address. Option ` + "`" + `envID` + "`" + ` writes the environment ID string inside new container. Option ` + "`" + `token` + "`" + ` is intended to check the origin of new container creation request during environment build. This is one of the security checks which makes sure that each container creation request is authorized by registered user.\nOption ` + "`" + `kurjunToken` + "`" + ` set kurjun token to clone private and shared templates\nThe clone options are not intended for manual use: unless you're confident about what you're doing. Use default clone format without additional options to create Subutai containers.",
        "tags": [
          "agent",
          "cli",
          "container"
        ],
        "operationId": "clone",
        "parameters": [
          {
            "type": "string",
            "name": "parent",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "child",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/cloneArgs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/rest/v1/config": {
      "post": {
        "description": "Allows read and write container's configuration file",
        "tags": [
          "agent",
          "cli",
          "container"
        ],
        "operationId": "config",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/configOptions"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "description": "Delete entry in configuration file",
        "tags": [
          "config"
        ],
        "operationId": "destroyEntry",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/configOptions"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/rest/v1/container/{name}": {
      "get": {
        "description": "Get container Info in JSON formatted object",
        "tags": [
          "agent",
          "cli",
          "container"
        ],
        "operationId": "getContainerInfo",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/container"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "delete": {
        "tags": [
          "container"
        ],
        "operationId": "destroyOne",
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      },
      "parameters": [
        {
          "type": "string",
          "name": "name",
          "in": "path",
          "required": true
        },
        {
          "type": "boolean",
          "name": "containersOnly",
          "in": "query"
        },
        {
          "type": "boolean",
          "name": "templatesOnly",
          "in": "query"
        },
        {
          "type": "boolean",
          "name": "detailedInfo",
          "in": "query"
        },
        {
          "type": "boolean",
          "name": "withAncestors",
          "in": "query"
        },
        {
          "type": "boolean",
          "name": "withParents",
          "in": "query"
        }
      ]
    },
    "/rest/v1/demote/{container}": {
      "post": {
        "description": "Converts template into regular Subutai container.\nA Subutai template is a \"locked down\" container only to be used for cloning purposes. It cannot be started, and its file system cannot be modified: it's read-only. Normal operational containers are promoted into templates, but sometimes you might want to demote them back to regular containers. This is what the demote sub command does: it reverts a template without children back into a normal container. Demoted container will use NAT network interface and dynamic IP address if opposite options are not specified.",
        "tags": [
          "agent",
          "cli",
          "container"
        ],
        "operationId": "demote",
        "parameters": [
          {
            "type": "string",
            "name": "container",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/demoteOptions"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/rest/v1/destroy/{ID}": {
      "get": {
        "description": "Destroy Subutai container",
        "tags": [
          "agent",
          "cli",
          "list"
        ],
        "operationId": "destroy",
        "parameters": [
          {
            "type": "string",
            "name": "ID",
            "in": "path",
            "required": true
          },
          {
            "type": "boolean",
            "name": "vlan",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/message"
              }
            }
          },
          "default": {
            "description": "generic error response"
          }
        },
        "x-swagger-router-controller": "Default"
      }
    },
    "/rest/v1/export/{container}": {
      "post": {
        "description": "Export prepares an archive from a template in the ` + "`" + `/mnt/lib/lxc/tmpdir/` + "`" + ` path. This archive can be moved to another Subutai peer and deployed as ready-to-use template or uploaded to Subutai's global template repository to make it widely available for others to use.\nExport consist of two steps if the target is a container: container promotion to template (see \"promote\" command) and packing the template into the archive. If already a template just the packing of the archive takes place.\nConfiguration values for template metadata parameters can be overridden on export, like the recommended container size when the template is cloned using ` + "`" + `-s` + "`" + ` option.",
        "tags": [
          "agent",
          "cli",
          "container"
        ],
        "operationId": "export",
        "parameters": [
          {
            "type": "string",
            "name": "container",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/exportOptions"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/rest/v1/hostname/{container}/{name}": {
      "post": {
        "description": "        - type: string name: container in: path required: true",
        "tags": [
          "agent",
          "cli",
          "container"
        ],
        "operationId": "hostname",
        "parameters": [
          {
            "type": "string",
            "name": "container",
            "in": "path",
            "required": true
          },
          {
            "type": "string",
            "name": "name",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/rest/v1/import/{container}": {
      "get": {
        "description": "Import Subutai template",
        "tags": [
          "agent",
          "cli",
          "list"
        ],
        "operationId": "import",
        "parameters": [
          {
            "type": "string",
            "name": "container",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/importOptions"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/message"
              }
            }
          },
          "default": {
            "description": "generic error response"
          }
        },
        "x-swagger-router-controller": "Default"
      }
    },
    "/rest/v1/info": {
      "get": {
        "description": "Info command's purposed is to display common system information, such as external IP address to access the container host quotas, its CPU model, RAM size, etc. It's mainly used for internal SS needs.",
        "tags": [
          "agent",
          "cli",
          "list"
        ],
        "operationId": "info",
        "parameters": [
          {
            "type": "string",
            "name": "command",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/infoHostStat"
              }
            }
          },
          "default": {
            "description": "generic error response"
          }
        },
        "x-swagger-router-controller": "Default"
      }
    },
    "/rest/v1/list": {
      "get": {
        "description": "Info returns JSON formatted list of Subutai instances with information such as IP address, parent template, etc.",
        "tags": [
          "agent",
          "cli",
          "list"
        ],
        "operationId": "cliList",
        "parameters": [
          {
            "type": "integer",
            "format": "int32",
            "default": 20,
            "name": "limit",
            "in": "query"
          },
          {
            "type": "boolean",
            "name": "containersOnly",
            "in": "query"
          },
          {
            "type": "boolean",
            "name": "templatesOnly",
            "in": "query"
          },
          {
            "type": "boolean",
            "name": "withAncestors",
            "in": "query"
          },
          {
            "type": "boolean",
            "name": "withParents",
            "in": "query"
          }
        ],
        "responses": {
          "200": {
            "description": "List all subutai instances",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/container"
              }
            }
          },
          "default": {
            "description": "generic error response"
          }
        },
        "x-swagger-router-controller": "Default"
      }
    },
    "/rest/v1/metrics": {
      "get": {
        "description": "HostMetrics function retrieves monitoring data from a time-series database deployed in the SS Management server for container hosts and Resource Hosts. Statistics are being collected by the Subutai daemon and includes common information like CPU utilization, network load, RAM and disk usage for both containers and hosts. Since the database is located on the SS Management Host, hosts which are not a part of a Subutai peer have no access to this information. Data aggregation in the time-series database has following configuration: last hour statistic is stored \"as is\", last day data aggregates to 1 minute interval, last week is in 5 minute intervals, After 7 days all statistics is are overwritten by new incoming data.",
        "tags": [
          "agent",
          "cli",
          "list"
        ],
        "operationId": "metrics",
        "parameters": [
          {
            "type": "string",
            "name": "start",
            "in": "query",
            "required": true
          },
          {
            "type": "string",
            "name": "end",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/metricInfo"
              }
            }
          },
          "default": {
            "description": "generic error response"
          }
        },
        "x-swagger-router-controller": "Default"
      }
    },
    "/rest/v1/p2p": {
      "get": {
        "description": "P2P function controls and configures the peer-to-peer network structure: the swarm which includes all hosts with same the same swarm hash and secret key.\nP2P is a base layer for Subutai environment networking: all containers in same environment are connected to each other via VXLAN tunnels and are accesses as if they were in one LAN. It doesn't matter where the containers are physically located.",
        "tags": [
          "agent",
          "cli",
          "list"
        ],
        "operationId": "p2pList",
        "parameters": [
          {
            "type": "string",
            "description": "Accepted commands: - list: list of p2p instances - interfaces: list of p2p interfaces - peers: list of p2p swarm participants by hash - version: print p2p version",
            "name": "command",
            "in": "query",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/text"
              }
            }
          },
          "default": {
            "description": "generic error response"
          }
        },
        "x-swagger-router-controller": "Default"
      },
      "put": {
        "description": "P2P function controls and configures the peer-to-peer network structure: the swarm which includes all hosts with same the same swarm hash and secret key.\nP2P is a base layer for Subutai environment networking: all containers in same environment are connected to each other via VXLAN tunnels and are accesses as if they were in one LAN. It doesn't matter where the containers are physically located.",
        "tags": [
          "agent",
          "cli",
          "list"
        ],
        "operationId": "p2pUpdate",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/p2pArgs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/text"
              }
            }
          },
          "default": {
            "description": "generic error response"
          }
        },
        "x-swagger-router-controller": "Default"
      },
      "post": {
        "description": "P2P function controls and configures the peer-to-peer network structure: the swarm which includes all hosts with same the same swarm hash and secret key.\nP2P is a base layer for Subutai environment networking: all containers in same environment are connected to each other via VXLAN tunnels and are accesses as if they were in one LAN. It doesn't matter where the containers are physically located.",
        "tags": [
          "agent",
          "cli",
          "list"
        ],
        "operationId": "p2pCreate",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/p2pArgs"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/item"
            }
          },
          "default": {
            "description": "generic error response"
          }
        },
        "x-swagger-router-controller": "Default"
      },
      "delete": {
        "description": "P2P function controls and configures the peer-to-peer network structure: the swarm which includes all hosts with same the same swarm hash and secret key.\nP2P is a base layer for Subutai environment networking: all containers in same environment are connected to each other via VXLAN tunnels and are accesses as if they were in one LAN. It doesn't matter where the containers are physically located.",
        "tags": [
          "agent",
          "cli",
          "list"
        ],
        "operationId": "p2pDelete",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/p2pArgs"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "generic error response"
          }
        },
        "x-swagger-router-controller": "Default"
      }
    },
    "/rest/v1/promote/{container}": {
      "post": {
        "description": "Promote turns a Subutai container into container template which may be cloned with \"clone\" command. Promote executes several simple steps, such as dropping a container's configuration to default values, dumping the list of installed packages (this step requires the target container to still be running), and setting the container's filesystem to read-only to prevent changes.",
        "tags": [
          "agent",
          "cli",
          "container"
        ],
        "operationId": "promote",
        "parameters": [
          {
            "type": "string",
            "name": "container",
            "in": "path",
            "required": true
          },
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/promoteOptions"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "error",
            "schema": {
              "$ref": "#/definitions/error"
            }
          }
        }
      }
    },
    "/rest/v1/proxy": {
      "get": {
        "description": "ProxyCheck exits with 0 code if domain or node is exists in specified vlan, otherwise exitcode is 1",
        "tags": [
          "agent",
          "cli",
          "list"
        ],
        "operationId": "proxyCheck",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/proxyArgs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/message"
              }
            }
          },
          "default": {
            "description": "generic error response"
          }
        },
        "x-swagger-router-controller": "Default"
      },
      "post": {
        "description": "ProxyAdd checks input args and perform required operations to configure reverse proxy",
        "tags": [
          "agent",
          "cli",
          "list"
        ],
        "operationId": "proxyCreate",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/proxyArgs"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created",
            "schema": {
              "$ref": "#/definitions/message"
            }
          },
          "default": {
            "description": "generic error response"
          }
        },
        "x-swagger-router-controller": "Default"
      },
      "delete": {
        "description": "ProxyDel checks what need to be removed - domain or node and pass args to required functions",
        "tags": [
          "agent",
          "cli",
          "list"
        ],
        "operationId": "proxyDelete",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/proxyArgs"
            }
          }
        ],
        "responses": {
          "204": {
            "description": "Deleted"
          },
          "default": {
            "description": "generic error response"
          }
        },
        "x-swagger-router-controller": "Default"
      }
    },
    "/rest/v1/quota": {
      "get": {
        "description": "Quota function controls container's quotas and thresholds. Available resources: cpu, % cpuset, available cores ram, Mb network, Kbps rootfs/home/var/opt, Gb The threshold value represents a percentage for each resource. Once resource consumption exceeds this threshold it triggers an alert. The clone operation, sets no quotas and thresholds for new containers; quotas need to be configured with quota command after a clone operation.",
        "tags": [
          "agent",
          "cli",
          "list"
        ],
        "operationId": "quota",
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "schema": {
              "$ref": "#/definitions/quotaArgs"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "Ok",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/quotaOutput"
              }
            }
          },
          "default": {
            "description": "generic error response"
          }
        },
        "x-swagger-router-controller": "Default"
      }
    },
    "/rest/v1/rh/id": {
      "get": {
        "description": "Returns JSON formatted Id of RH, UUID which is the PGP fingerprint",
        "tags": [
          "agent",
          "rh"
        ],
        "operationId": "rhID",
        "responses": {
          "200": {
            "description": "JSON formatted Id of RH",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/fingerprint"
              }
            }
          },
          "default": {
            "description": "generic error response"
          }
        },
        "x-swagger-router-controller": "Default"
      }
    }
  },
  "definitions": {
    "batchLine": {
      "type": "object",
      "required": [
        "action",
        "args"
      ],
      "properties": {
        "action": {
          "type": "string",
          "readOnly": true
        },
        "args": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/item"
          },
          "readOnly": true
        }
      }
    },
    "cloneArgs": {
      "type": "object",
      "required": [
        "parent",
        "child"
      ],
      "properties": {
        "child": {
          "type": "string",
          "readOnly": true
        },
        "envID": {
          "type": "string",
          "readOnly": true
        },
        "ipaddr": {
          "type": "string",
          "readOnly": true
        },
        "kurjunToken": {
          "type": "string",
          "readOnly": true
        },
        "parent": {
          "type": "string",
          "readOnly": true
        },
        "token": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "configOptions": {
      "type": "object",
      "required": [
        "key",
        "value"
      ],
      "properties": {
        "key": {
          "type": "string",
          "readOnly": true
        },
        "value": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "container": {
      "type": "object",
      "required": [
        "name"
      ],
      "properties": {
        "ancestor": {
          "type": "string",
          "minLength": 1
        },
        "name": {
          "type": "string",
          "readOnly": true
        },
        "parent": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "cpuInfo": {
      "type": "object",
      "properties": {
        "coreCount": {
          "type": "string",
          "readOnly": true
        },
        "frequency": {
          "type": "string",
          "readOnly": true
        },
        "idle": {
          "type": "string",
          "readOnly": true
        },
        "model": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "demoteOptions": {
      "type": "object",
      "properties": {
        "ipaddr": {
          "type": "string",
          "readOnly": true
        },
        "vlan": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "diskInfo": {
      "type": "object",
      "properties": {
        "total": {
          "type": "string",
          "readOnly": true
        },
        "used": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "diskStructInfo": {
      "type": "object",
      "properties": {
        "home": {
          "type": "string",
          "readOnly": true
        },
        "opt": {
          "type": "string",
          "readOnly": true
        },
        "rootfs": {
          "type": "string",
          "readOnly": true
        },
        "var": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "error": {
      "type": "object",
      "required": [
        "message"
      ],
      "properties": {
        "code": {
          "type": "integer",
          "format": "int64"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "exportOptions": {
      "type": "object",
      "properties": {
        "description": {
          "type": "string",
          "readOnly": true
        },
        "private": {
          "type": "boolean",
          "readOnly": true
        },
        "size": {
          "type": "integer",
          "readOnly": true
        },
        "token": {
          "type": "string",
          "readOnly": true
        },
        "version": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "fingerprint": {
      "type": "object",
      "required": [
        "hash"
      ],
      "properties": {
        "hash": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "importOptions": {
      "type": "object",
      "properties": {
        "token": {
          "type": "string",
          "readOnly": true
        },
        "torrent": {
          "type": "boolean",
          "readOnly": true
        },
        "version": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "infoHostStat": {
      "type": "object",
      "properties": {
        "cpu": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/cpuInfo"
          },
          "readOnly": true
        },
        "disk": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/diskInfo"
          },
          "readOnly": true
        },
        "hostname": {
          "type": "string",
          "readOnly": true
        },
        "quota": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/quotaInfo"
          },
          "readOnly": true
        },
        "ram": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/ramInfo"
          },
          "readOnly": true
        }
      }
    },
    "item": {
      "type": "object",
      "required": [
        "text"
      ],
      "properties": {
        "text": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "message": {
      "type": "object",
      "required": [
        "text"
      ],
      "properties": {
        "exitcode": {
          "type": "string",
          "readOnly": true
        },
        "text": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "metricInfo": {
      "type": "object",
      "properties": {
        "messages": {
          "type": "string",
          "readOnly": true
        },
        "series": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "p2pArgs": {
      "type": "object",
      "required": [
        "interfaceName",
        "hash",
        "key",
        "ttl"
      ],
      "properties": {
        "hash": {
          "type": "string",
          "readOnly": true
        },
        "interfaceName": {
          "type": "string",
          "readOnly": true
        },
        "key": {
          "type": "string",
          "readOnly": true
        },
        "localPeepIPAddr": {
          "type": "string",
          "readOnly": true
        },
        "portRange": {
          "type": "string",
          "readOnly": true
        },
        "ttl": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "promoteOptions": {
      "type": "object",
      "properties": {
        "source": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "proxyArgs": {
      "type": "object",
      "required": [
        "vlan"
      ],
      "properties": {
        "cert": {
          "type": "string",
          "minLength": 1
        },
        "domain": {
          "type": "boolean",
          "minLength": 1
        },
        "node": {
          "type": "string",
          "minLength": 1
        },
        "policy": {
          "type": "string",
          "minLength": 1
        },
        "vlan": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "quotaArgs": {
      "type": "object",
      "required": [
        "name",
        "resource"
      ],
      "properties": {
        "name": {
          "type": "string",
          "readOnly": true
        },
        "resource": {
          "type": "string",
          "minLength": 1
        },
        "size": {
          "type": "string",
          "minLength": 1
        },
        "threshold": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "quotaInfo": {
      "type": "object",
      "properties": {
        "container": {
          "type": "string",
          "readOnly": true
        },
        "cpu": {
          "type": "string",
          "readOnly": true
        },
        "disk": {
          "type": "array",
          "items": {
            "$ref": "#/definitions/diskStructInfo"
          },
          "readOnly": true
        },
        "ram": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "quotaOutput": {
      "type": "object",
      "required": [
        "quota",
        "threshold"
      ],
      "properties": {
        "quota": {
          "type": "string",
          "readOnly": true
        },
        "threshold": {
          "type": "string",
          "minLength": 1
        }
      }
    },
    "ramInfo": {
      "type": "object",
      "properties": {
        "cached": {
          "type": "string",
          "readOnly": true
        },
        "free": {
          "type": "string",
          "readOnly": true
        },
        "total": {
          "type": "string",
          "readOnly": true
        }
      }
    },
    "text": {
      "type": "object",
      "properties": {
        "text": {
          "type": "string",
          "readOnly": true
        }
      }
    }
  }
}`))
}
