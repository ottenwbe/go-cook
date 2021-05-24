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
        "license": {
            "name": "MIT",
            "url": "https://github.com/ottenwbe/recipes-manager/blob/master/LICENSE"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/recipes": {
            "get": {
                "description": "A list of ids of recipes is returned",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "Get Recipes",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search for a specific name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search for a specific term in a description",
                        "name": "description",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Search for a specific ingredient",
                        "name": "ingredient",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/recipes.RecipeList"
                        }
                    }
                }
            },
            "post": {
                "description": "Adds a new recipe, the id will automatically overriden by the backend",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "Add a new Recipe",
                "parameters": [
                    {
                        "description": "Recipe",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/recipes.Recipe"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": ""
                    }
                }
            }
        },
        "/recipes/num": {
            "get": {
                "description": "The number of recipes is returned that is managed by the service.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "Get the number of recipes",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "/recipes/r/{recipe}": {
            "get": {
                "description": "A specific recipe is returned",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "Get a specific Recipe",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Number of Servings",
                        "name": "servings",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Recipe ID",
                        "name": "recipe",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/recipes.Recipe"
                        }
                    }
                }
            },
            "put": {
                "description": "A specific recipe is updates",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "Update a specific Recipe",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recipe ID",
                        "name": "recipe",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Recipe",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/recipes.Recipe"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            },
            "delete": {
                "description": "Deletes a recipe by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "Delete a Recipe",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recipe ID",
                        "name": "recipe",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/recipes/r/{recipe}/pictures/{name}": {
            "get": {
                "description": "A specific picture of a specific recipe is returned",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "Get a picture of a",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Recipe ID",
                        "name": "recipe",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Name of Picture",
                        "name": "name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/recipes.RecipePicture"
                        }
                    }
                }
            }
        },
        "/recipes/rand": {
            "get": {
                "description": "A specific picture of a specific recipe is returned",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Recipes"
                ],
                "summary": "Get a Random Recipe",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Number of Servings",
                        "name": "servings",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/recipes.Recipe"
                        }
                    }
                }
            }
        },
        "/sources": {
            "get": {
                "description": "List sources",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sources"
                ],
                "summary": "List sources",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "$ref": "#/definitions/sources.SourceResponse"
                            }
                        }
                    }
                }
            }
        },
        "/sources/{source}/connect": {
            "get": {
                "description": "Trigger the oauth process",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sources"
                ],
                "summary": "Trigger the oauth process",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Source ID",
                        "name": "source",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/sources.SourceOAuthConnectResponse"
                        }
                    }
                }
            }
        },
        "/sources/{source}/oauth": {
            "get": {
                "description": "Handles Tokens. Typically not directly called.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sources"
                ],
                "summary": "Handles Tokens",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Source ID",
                        "name": "source",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "301": {
                        "description": "Moved Permanently",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/sources/{source}/recipes": {
            "get": {
                "description": "Download recipes from a source",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Sources"
                ],
                "summary": "Download Recipes from a Source",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Source ID",
                        "name": "source",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": ""
                    }
                }
            }
        },
        "/version": {
            "get": {
                "description": "get the current version",
                "produces": [
                    "application/json"
                ],
                "summary": "Get the curent version",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/core.Version"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "core.Version": {
            "type": "object",
            "properties": {
                "api": {
                    "description": "API is the MAJOR API Version supported by the app",
                    "type": "string"
                },
                "app": {
                    "description": "APP is the version of the current app",
                    "type": "string"
                }
            }
        },
        "recipes.Ingredients": {
            "type": "object",
            "properties": {
                "amount": {
                    "description": "Amount needed in a recipe of an ingredient",
                    "type": "number"
                },
                "name": {
                    "description": "Name of the ingredient",
                    "type": "string"
                },
                "unit": {
                    "description": "Unit of the Amount",
                    "type": "string"
                }
            }
        },
        "recipes.Recipe": {
            "type": "object",
            "properties": {
                "components": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/recipes.Ingredients"
                    }
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pictureLink": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "servings": {
                    "type": "integer"
                }
            }
        },
        "recipes.RecipeList": {
            "type": "object",
            "properties": {
                "recipes": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "recipes.RecipePicture": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "picture": {
                    "type": "string"
                }
            }
        },
        "sources.SourceOAuthConnectResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "string"
                },
                "oAuthURL": {
                    "type": "string"
                }
            }
        },
        "sources.SourceResponse": {
            "type": "object",
            "properties": {
                "connected": {
                    "type": "boolean"
                },
                "id": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
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
	Version:     "1.0",
	Host:        "",
	BasePath:    "/api/v1",
	Schemes:     []string{},
	Title:       "Swagger API documentation for recipes-manager",
	Description: "This is the API documentation for recipes-manager.",
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
