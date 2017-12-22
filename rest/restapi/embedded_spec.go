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
    "/agent/container/{name}": {
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
    "/agent/rest/list": {
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
    }
  },
  "definitions": {
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
    }
  }
}`))
}
