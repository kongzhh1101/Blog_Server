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
        "/images": {
            "get": {
                "description": "获取图片列表",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "查询图片",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "页码",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "每页显示所少",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "更改图片名字",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "更新图片",
                "parameters": [
                    {
                        "description": "要修改的图片id",
                        "name": "id",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/image_api.ImageUpdateRequest"
                        }
                    }
                ],
                "responses": {}
            },
            "post": {
                "description": "上传图片到数据库",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "上传图片",
                "parameters": [
                    {
                        "type": "file",
                        "description": "图片文件",
                        "name": "images",
                        "in": "formData"
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "删除图片",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Image"
                ],
                "summary": "删除图片",
                "parameters": [
                    {
                        "description": "要删除的图片id列表",
                        "name": "some_id",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/models.RemoveList"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/settings/{name}": {
            "get": {
                "description": "根据传入的种类，返回对应的设置信息",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "设置信息"
                ],
                "summary": "查看设置信息",
                "parameters": [
                    {
                        "type": "string",
                        "description": "设置种类",
                        "name": "name",
                        "in": "path"
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "根据传入的种类，和数据，修改对应设置的信息",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "设置信息"
                ],
                "summary": "跟新设置",
                "parameters": [
                    {
                        "description": "设置种类",
                        "name": "data",
                        "in": "body",
                        "schema": {
                            "$ref": "#/definitions/config.Config"
                        }
                    },
                    {
                        "type": "string",
                        "description": "设置种类",
                        "name": "name",
                        "in": "path"
                    }
                ],
                "responses": {}
            }
        }
    },
    "definitions": {
        "config.Config": {
            "type": "object",
            "properties": {
                "email": {
                    "$ref": "#/definitions/config.Email"
                },
                "jwt": {
                    "$ref": "#/definitions/config.JWT"
                },
                "logger": {
                    "$ref": "#/definitions/config.Logger"
                },
                "mysql": {
                    "$ref": "#/definitions/config.Mysql"
                },
                "qiNiu": {
                    "$ref": "#/definitions/config.QiNiu"
                },
                "qq": {
                    "$ref": "#/definitions/config.QQ"
                },
                "siteInfo": {
                    "$ref": "#/definitions/config.SiteInfo"
                },
                "system": {
                    "$ref": "#/definitions/config.System"
                },
                "upload": {
                    "$ref": "#/definitions/config.Upload"
                }
            }
        },
        "config.Email": {
            "type": "object",
            "properties": {
                "default-from-email": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "passwrod": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                },
                "user": {
                    "type": "string"
                },
                "user-ssl": {
                    "type": "boolean"
                },
                "user-tls": {
                    "type": "boolean"
                }
            }
        },
        "config.JWT": {
            "type": "object",
            "properties": {
                "expire": {
                    "type": "integer"
                },
                "issure": {
                    "type": "string"
                },
                "secret": {
                    "type": "string"
                }
            }
        },
        "config.Logger": {
            "type": "object",
            "properties": {
                "director": {
                    "type": "string"
                },
                "level": {
                    "type": "string"
                },
                "logInConsol": {
                    "type": "boolean"
                },
                "perfix": {
                    "type": "string"
                },
                "showLine": {
                    "description": "是否显示行号",
                    "type": "boolean"
                }
            }
        },
        "config.Mysql": {
            "type": "object",
            "properties": {
                "config": {
                    "description": "高级配置",
                    "type": "string"
                },
                "db-name": {
                    "description": "数据库名",
                    "type": "string"
                },
                "engine": {
                    "description": "数据库引擎，默认InnoDB",
                    "type": "string",
                    "default": "InnoDB"
                },
                "log-mode": {
                    "description": "是否开启Gorm全局日志",
                    "type": "string"
                },
                "log-zap": {
                    "description": "是否通过zap写入日志文件",
                    "type": "boolean"
                },
                "max-idle-conns": {
                    "description": "空闲中的最大连接数",
                    "type": "integer"
                },
                "max-open-conns": {
                    "description": "打开到数据库的最大连接数",
                    "type": "integer"
                },
                "password": {
                    "description": "数据库密码",
                    "type": "string"
                },
                "path": {
                    "description": "数据库地址",
                    "type": "string"
                },
                "port": {
                    "description": "数据库端口",
                    "type": "string"
                },
                "prefix": {
                    "description": "数据库前缀",
                    "type": "string"
                },
                "singular": {
                    "description": "是否开启全局禁用复数，true表示开启",
                    "type": "boolean"
                },
                "username": {
                    "description": "数据库账号",
                    "type": "string"
                }
            }
        },
        "config.QQ": {
            "type": "object",
            "properties": {
                "app_id": {
                    "type": "string"
                },
                "key": {
                    "type": "string"
                },
                "redirect": {
                    "type": "string"
                }
            }
        },
        "config.QiNiu": {
            "type": "object",
            "properties": {
                "access_key": {
                    "type": "string"
                },
                "bucket": {
                    "type": "string"
                },
                "cdn": {
                    "type": "string"
                },
                "enabeld": {
                    "type": "boolean"
                },
                "secret_key": {
                    "type": "string"
                },
                "size": {
                    "type": "number"
                },
                "zone": {
                    "type": "string"
                }
            }
        },
        "config.SiteInfo": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "version": {
                    "type": "string"
                }
            }
        },
        "config.System": {
            "type": "object",
            "properties": {
                "env": {
                    "type": "string"
                },
                "host": {
                    "type": "string"
                },
                "port": {
                    "type": "integer"
                }
            }
        },
        "config.Upload": {
            "type": "object",
            "properties": {
                "path": {
                    "type": "string"
                },
                "size": {
                    "type": "integer"
                }
            }
        },
        "image_api.ImageUpdateRequest": {
            "type": "object",
            "required": [
                "id",
                "name"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "models.RemoveList": {
            "type": "object",
            "properties": {
                "id_list": {
                    "type": "array",
                    "items": {
                        "type": "integer"
                    }
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "博客后端",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}