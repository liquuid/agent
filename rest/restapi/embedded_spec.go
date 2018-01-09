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
    }
  }
}`))
}
