// GENERATED BY THE COMMAND ABOVE; DO NOT EDIT
// This file was generated by swaggo/swag

package docs

import (
	"bytes"
	"encoding/json"
	"strings"

	"github.com/alecthomas/template"
	"github.com/swaggo/swag"
)

var doc = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{.Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/order/create": {
            "post": {
                "description": "Create new resource to be saved to database",
                "consumes": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create new resource to be saved to database",
                "operationId": "order-create",
                "parameters": [
                    {
                        "description": "request body json",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.OrderDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/order/delete/{id}": {
            "delete": {
                "description": "Delete existing resource",
                "produces": [
                    "application/json"
                ],
                "summary": "Delete existing resource",
                "operationId": "order-delete",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Resource primary key",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/order/edit/{id}": {
            "patch": {
                "description": "Edit existing resource",
                "consumes": [
                    "application/json",
                    "application/x-www-form-urlencoded"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Edit existing resource",
                "operationId": "order-edit",
                "parameters": [
                    {
                        "description": "request body json",
                        "name": "order",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.OrderDTO"
                        }
                    },
                    {
                        "type": "string",
                        "description": "Resource primary key",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/order/index": {
            "get": {
                "description": "Index list of resources",
                "produces": [
                    "application/json"
                ],
                "summary": "Index list of resources",
                "operationId": "order-index",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        },
        "/api/v1/order/show/{id}": {
            "get": {
                "description": "Show detail of resource",
                "produces": [
                    "application/json"
                ],
                "summary": "Show detail of resource",
                "operationId": "order-show",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Resource primary key",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/controllers.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "controllers.Meta": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "error": {
                    "type": "string"
                },
                "errors": {
                    "type": "object",
                    "additionalProperties": {
                        "type": "string"
                    }
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "controllers.Response": {
            "type": "object",
            "properties": {
                "data": {
                    "type": "object"
                },
                "meta": {
                    "$ref": "#/definitions/controllers.Meta"
                }
            }
        },
        "dto.ItemDTO": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "desc": {
                    "type": "string"
                },
                "qty": {
                    "type": "integer"
                }
            }
        },
        "dto.OrderDTO": {
            "type": "object",
            "properties": {
                "customer_name": {
                    "type": "string"
                },
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/dto.ItemDTO"
                    }
                }
            }
        }
    }
}`

type swaggerInfo struct {
	Version     string
	Host        string
	BasePath    string
	Schemes     []string
	Title       string
	Description string
}

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = swaggerInfo{
	Version:     "",
	Host:        "",
	BasePath:    "",
	Schemes:     []string{},
	Title:       "",
	Description: "",
}

type s struct{}

func (s *s) ReadDoc() string {
	sInfo := SwaggerInfo
	sInfo.Description = strings.Replace(sInfo.Description, "\n", "\\n", -1)

	t, err := template.New("swagger_info").Funcs(template.FuncMap{
		"marshal": func(v interface{}) string {
			a, _ := json.Marshal(v)
			return string(a)
		},
	}).Parse(doc)
	if err != nil {
		return doc
	}

	var tpl bytes.Buffer
	if err := t.Execute(&tpl, sInfo); err != nil {
		return doc
	}

	return tpl.String()
}

func init() {
	swag.Register(swag.Name, &s{})
}
