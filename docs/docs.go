// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/product": {
            "post": {
                "description": "Add a new product",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create product",
                "parameters": [
                    {
                        "type": "string",
                        "name": "Name",
                        "in": "path"
                    },
                    {
                        "type": "string",
                        "name": "PK",
                        "in": "path"
                    },
                    {
                        "type": "number",
                        "name": "Value",
                        "in": "path"
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Product created successfully",
                        "schema": {
                            "$ref": "#/definitions/aggregate.Product"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/product/{id}": {
            "get": {
                "description": "Retorna os detalhes de um produto com base no ID fornecido.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Recupera um produto por ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ID do produto",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Product created successfully",
                        "schema": {
                            "$ref": "#/definitions/aggregate.Product"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "aggregate.Product": {
            "type": "object",
            "properties": {
                "categoryID": {
                    "type": "string"
                },
                "createdAt": {
                    "type": "string"
                },
                "isActive": {
                    "type": "boolean"
                },
                "name": {
                    "type": "string"
                },
                "pk": {
                    "type": "string"
                },
                "updatedAt": {
                    "type": "string"
                },
                "validity": {
                    "type": "string"
                },
                "value": {
                    "type": "number"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8081",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger Example API",
	Description:      "This is a sample server for using Swagger with Echo.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
