// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/",
        "contact": {
            "name": "API Support",
            "url": "http://www.swagger.io/support",
            "email": "support@swagger.io"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/admin/problem/": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Admin create a problem",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "admin"
                ],
                "summary": "Create a Problem",
                "parameters": [
                    {
                        "description": "create problem",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.InsertProblem"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Success: Problem created",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    },
                    "400": {
                        "description": "Bad request",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/user/go/exec": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "The code the user sent will be executed",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Execute code",
                "parameters": [
                    {
                        "description": "Go code to execute",
                        "name": "code",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/user/go/{id}": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "The code the user sent will be executed and the result will be given",
                "consumes": [
                    "text/plain"
                ],
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "user"
                ],
                "summary": "Execute code",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Problem ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Go code to execute",
                        "name": "code",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "success",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.JsonResponse"
                        }
                    }
                }
            }
        },
        "/user/problem/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "get one problem to display",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user",
                    "admin"
                ],
                "summary": "get one problems",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Problem ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.Problem"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.Problem"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.Problem"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.Problem"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.Problem"
                        }
                    }
                }
            }
        },
        "/user/signin": {
            "post": {
                "description": "signin to code-up",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user",
                    "admin"
                ],
                "summary": "User signin",
                "parameters": [
                    {
                        "description": "user details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserSignInRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.SignUpWrapper"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.SignUpWrapper"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.SignUpWrapper"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.SignUpWrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.SignUpWrapper"
                        }
                    }
                }
            }
        },
        "/user/signup": {
            "post": {
                "description": "signup to code-up",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "User signup",
                "parameters": [
                    {
                        "description": "user details",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UserSignUpRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.SignUpWrapper"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/response.SignUpWrapper"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/response.SignUpWrapper"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/response.SignUpWrapper"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/response.SignUpWrapper"
                        }
                    }
                }
            }
        },
        "/user/view": {
            "get": {
                "description": "View all problems code-up",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user",
                    "admin"
                ],
                "summary": "View problems",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Problem"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Problem"
                            }
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Problem"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Problem"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/response.Problem"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.InsertProblem": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "difficulty": {
                    "type": "string"
                },
                "memory_limit": {
                    "type": "integer"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "test_cases": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/request.TestCase"
                    }
                },
                "time_limit": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "request.TestCase": {
            "type": "object",
            "properties": {
                "input": {
                    "type": "string"
                },
                "output": {
                    "type": "string"
                }
            }
        },
        "request.UserSignInRequest": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        },
        "request.UserSignUpRequest": {
            "type": "object",
            "required": [
                "confirmpassword",
                "email",
                "first_name",
                "last_name",
                "password",
                "username"
            ],
            "properties": {
                "confirmpassword": {
                    "type": "string",
                    "minLength": 8
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string",
                    "minLength": 3
                },
                "last_name": {
                    "type": "string",
                    "minLength": 3
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 6
                }
            }
        },
        "response.JsonResponse": {
            "type": "object",
            "properties": {
                "data": {},
                "error": {},
                "message": {
                    "type": "string"
                }
            }
        },
        "response.Problem": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "difficulty": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "memory_limit": {
                    "type": "integer"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "test_cases": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/response.TestCase"
                    }
                },
                "time_limit": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "response.SignUpWrapper": {
            "type": "object",
            "properties": {
                "token": {
                    "type": "string"
                },
                "user": {
                    "$ref": "#/definitions/response.UserSignUpResponse"
                }
            }
        },
        "response.TestCase": {
            "type": "object",
            "properties": {
                "input": {
                    "type": "string"
                },
                "output": {
                    "type": "string"
                }
            }
        },
        "response.UserSignUpResponse": {
            "type": "object",
            "properties": {
                "created_at": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "first_name": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "last_name": {
                    "description": "for proper representaion of null value in go sql.Null is used",
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Code Up Project API Documentation",
	Description:      "This is a sample code execution platform.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
