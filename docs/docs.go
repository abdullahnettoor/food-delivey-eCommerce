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
    "securityDefinitions": {
        "Bearer": {
        "type": "apiKey",
        "name": "Authorization",
        "in": "header"
        }
    },
    "paths": {
        "/admin/categories": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Retrieve a list of all categories",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Get all categories",
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/res.AllCategoriesRes"
                        }
                    },
                    "401": {
                        "description": "Unauthorized Access",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    }
                }
            }
        },
        "/admin/categories/addCategory": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Create a new category",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Add a category",
                "parameters": [
                    {
                        "description": "Category creation request",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.CreateCategoryReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Category successfully created",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "401": {
                        "description": "Unauthorized Access",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    }
                }
            }
        },
        "/admin/categories/{id}/edit": {
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Update an existing category by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Edit a category",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Category ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Category update request",
                        "name": "req",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.UpdateCategoryReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Category successfully updated",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "401": {
                        "description": "Unauthorized Access",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Admin login",
                "parameters": [
                    {
                        "description": "Admin Login Request",
                        "name": "adminLoginReq",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/req.AdminLoginReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Successful login",
                        "schema": {
                            "$ref": "#/definitions/res.AdminLoginRes"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    }
                }
            }
        },
        "/admin/sellers": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get a list of all sellers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Get all sellers",
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/res.SellerListRes"
                        }
                    },
                    "401": {
                        "description": "Unauthorized Access",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    }
                }
            }
        },
        "/admin/sellers/{id}/block": {
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Block a specific seller by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Block a seller",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Seller ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Seller successfully blocked",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "401": {
                        "description": "Unauthorized Access",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    }
                }
            }
        },
        "/admin/sellers/{id}/unblock": {
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Unblock a specific seller by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Unblock a seller",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Seller ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Seller successfully unblocked",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "401": {
                        "description": "Unauthorized Access",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    }
                }
            }
        },
        "/admin/sellers/{id}/verify": {
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    },
                    {
                        "Bearer": []
                    }
                ],
                "description": "Verify a specific seller by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Verify a seller",
                "parameters": [
                    {
                        "type": "string",
                        "format": "uuid",
                        "description": "Seller ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Seller successfully verified",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "401": {
                        "description": "Unauthorized Access",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    }
                }
            }
        },
        "/admin/users": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Get a list of all users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Get all users",
                "responses": {
                    "200": {
                        "description": "Successful operation",
                        "schema": {
                            "$ref": "#/definitions/res.UserListRes"
                        }
                    },
                    "401": {
                        "description": "Unauthorized Access",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    }
                }
            }
        },
        "/admin/users/{id}/block": {
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Block a specific user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Block a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User successfully blocked",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "401": {
                        "description": "Unauthorized Access",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    }
                }
            }
        },
        "/admin/users/{id}/unblock": {
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "description": "Unblock a specific user by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Admin"
                ],
                "summary": "Unblock a user",
                "parameters": [
                    {
                        "type": "string",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "User successfully unblocked",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "401": {
                        "description": "Unauthorized Access",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/res.CommonRes"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "entities.Category": {
            "type": "object",
            "properties": {
                "categoryId": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "entities.Seller": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pinCode": {
                    "type": "string"
                },
                "sellerId": {
                    "type": "integer"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "entities.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "firstName": {
                    "type": "string"
                },
                "lastName": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "userId": {
                    "type": "integer"
                }
            }
        },
        "req.AdminLoginReq": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 3
                }
            }
        },
        "req.CreateCategoryReq": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "minLength": 3
                }
            }
        },
        "req.UpdateCategoryReq": {
            "type": "object",
            "required": [
                "name"
            ],
            "properties": {
                "name": {
                    "type": "string",
                    "minLength": 3
                }
            }
        },
        "res.AdminLoginRes": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "token": {
                    "type": "string"
                }
            }
        },
        "res.AllCategoriesRes": {
            "type": "object",
            "properties": {
                "categories": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Category"
                    }
                },
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "res.CommonRes": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "result": {},
                "status": {
                    "type": "string"
                }
            }
        },
        "res.SellerListRes": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "sellerList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.Seller"
                    }
                },
                "status": {
                    "type": "string"
                }
            }
        },
        "res.UserListRes": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "userList": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/entities.User"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
