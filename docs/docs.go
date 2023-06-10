// Package docs GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
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
        "/admin/authTest": {
            "get": {
                "description": "用户鉴权测试",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台用户管理"
                ],
                "summary": "用户鉴权测试",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_common.GinCommonResponse"
                        }
                    }
                }
            }
        },
        "/admin/info": {
            "get": {
                "security": [
                    {
                        "GinJWTMiddleware": []
                    }
                ],
                "description": "根据用户 ID 获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台用户管理"
                ],
                "summary": "根据用户 ID 获取用户信息",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_common.GinCommonResponse"
                        }
                    }
                }
            }
        },
        "/admin/list": {
            "post": {
                "security": [
                    {
                        "GinJWTMiddleware": []
                    }
                ],
                "description": "分页查询用户",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台用户管理"
                ],
                "summary": "分页查询用户",
                "parameters": [
                    {
                        "description": "分页查询用户",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ums_admin.UmsAdminPage"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_common.GinCommonResponse"
                        }
                    }
                }
            }
        },
        "/admin/login": {
            "post": {
                "description": "用户登录",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台用户管理"
                ],
                "summary": "用户登录",
                "parameters": [
                    {
                        "description": "用户登录",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ums_admin.UmsAdminLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_common.GinCommonResponse"
                        }
                    }
                }
            }
        },
        "/admin/refreshToken": {
            "post": {
                "security": [
                    {
                        "GinJWTMiddleware": []
                    }
                ],
                "description": "刷新 token",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台用户管理"
                ],
                "summary": "刷新 token",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_common.GinCommonResponse"
                        }
                    }
                }
            }
        },
        "/admin/register": {
            "post": {
                "description": "用户注册",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台用户管理"
                ],
                "summary": "用户注册",
                "parameters": [
                    {
                        "description": "用户注册",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ums_admin.UmsAdminRegisterRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_common.GinCommonResponse"
                        }
                    }
                }
            }
        },
        "/admin/update/{user_id}": {
            "post": {
                "security": [
                    {
                        "GinJWTMiddleware": []
                    }
                ],
                "description": "更新用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台用户管理"
                ],
                "summary": "更新用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户 ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "更新用户信息",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/ums_admin.UmsAdminUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_common.GinCommonResponse"
                        }
                    }
                }
            }
        },
        "/admin/{user_id}": {
            "get": {
                "security": [
                    {
                        "GinJWTMiddleware": []
                    }
                ],
                "description": "根据用户 ID 获取用户信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "后台用户管理"
                ],
                "summary": "根据用户 ID 获取用户信息",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "用户 ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/gin_common.GinCommonResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "gin_common.GinCommonResponse": {
            "type": "object",
            "properties": {
                "data": {
                    "description": "返回的数据是任意类型\t如果有错误，则把错误信息也封装在此\n\n\t\t{\n\t\t\t\"err_code\": 300000,\n\t\t\t\"err_msg\": \"用户名已存在\"\n\t\t}"
                },
                "status": {
                    "description": "success or fail",
                    "type": "string"
                }
            }
        },
        "ums_admin.UmsAdminLogin": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "password": {
                    "description": "密文密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "ums_admin.UmsAdminPage": {
            "type": "object",
            "required": [
                "page_num",
                "page_size"
            ],
            "properties": {
                "page_num": {
                    "description": "页码",
                    "type": "integer",
                    "default": 1
                },
                "page_size": {
                    "description": "每页数量",
                    "type": "integer",
                    "default": 10
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "ums_admin.UmsAdminRegisterRequest": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "icon": {
                    "description": "用户头像",
                    "type": "string"
                },
                "nickname": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "note": {
                    "description": "备注",
                    "type": "string"
                },
                "password": {
                    "description": "密文密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        },
        "ums_admin.UmsAdminUpdate": {
            "type": "object",
            "required": [
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "description": "邮箱",
                    "type": "string"
                },
                "icon": {
                    "description": "用户头像",
                    "type": "string"
                },
                "nickname": {
                    "description": "用户昵称",
                    "type": "string"
                },
                "note": {
                    "description": "备注",
                    "type": "string"
                },
                "password": {
                    "description": "密文密码",
                    "type": "string"
                },
                "username": {
                    "description": "用户名",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "GinJWTMiddleware": {
            "description": "Description for what is this security definition being used",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "v1",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{"http", "https"},
	Title:            "mall-go API",
	Description:      "mall-go API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
