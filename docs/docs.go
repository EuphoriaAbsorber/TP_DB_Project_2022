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
        "termsOfService": "http://swagger.io/terms/",
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
        "/forum/create": {
            "post": {
                "description": "Creates Forum",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Forum"
                ],
                "summary": "Creates Forum",
                "operationId": "CreateForum",
                "parameters": [
                    {
                        "description": "Forum params",
                        "name": "forum",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ForumCreateModel"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request - Problem with the request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "404": {
                        "description": "Not found - Requested entity is not found in database",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "409": {
                        "description": "Conflict - User already exists",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/forum/{slug}/create": {
            "post": {
                "description": "creates thread",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Forum"
                ],
                "summary": "creates thread",
                "operationId": "CreateThread",
                "parameters": [
                    {
                        "type": "string",
                        "description": "slug of forum",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Thread params",
                        "name": "thread",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ThreadCreateModel"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.Thread"
                        }
                    },
                    "404": {
                        "description": "Not found - Requested entity is not found in database",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "409": {
                        "description": "Conflict",
                        "schema": {
                            "$ref": "#/definitions/model.Thread"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/forum/{slug}/details": {
            "get": {
                "description": "Gets forum info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Forum"
                ],
                "summary": "Gets forum info",
                "operationId": "GetForumInfo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "slug of forum",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Forum"
                        }
                    },
                    "404": {
                        "description": "Not found - Requested entity is not found in database",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/forum/{slug}/threads": {
            "get": {
                "description": "Gets forum threads",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Forum"
                ],
                "summary": "Gets forum threads",
                "operationId": "GetForumThreads",
                "parameters": [
                    {
                        "type": "string",
                        "description": "slug of forum",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "since",
                        "name": "since",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "desc",
                        "name": "desc",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Thread"
                            }
                        }
                    },
                    "404": {
                        "description": "Not found - Requested entity is not found in database",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/forum/{slug}/users": {
            "get": {
                "description": "Gets forum users",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Forum"
                ],
                "summary": "Gets forum users",
                "operationId": "GetForumUsers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "slug of forum",
                        "name": "slug",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "since",
                        "name": "since",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "desc",
                        "name": "desc",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Forum"
                        }
                    },
                    "404": {
                        "description": "Not found - Requested entity is not found in database",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/post/{id}/details": {
            "get": {
                "description": "Gets post by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Gets post by id",
                "operationId": "GetPostById",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of post",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "related",
                        "name": "related",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PostFull"
                        }
                    },
                    "404": {
                        "description": "Not found - Requested entity is not found in database",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Edits post by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Post"
                ],
                "summary": "Editss post by id",
                "operationId": "UpdatePost",
                "parameters": [
                    {
                        "type": "string",
                        "description": "id of post",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "PostUpdate params",
                        "name": "message",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.PostUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.PostFull"
                        }
                    },
                    "404": {
                        "description": "Not found - Requested entity is not found in database",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/service/clear": {
            "post": {
                "description": "Clears Service info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Service"
                ],
                "summary": "Clears Service info",
                "operationId": "ServiceClear",
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/service/status": {
            "get": {
                "description": "Gets Service info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Service"
                ],
                "summary": "Gets Service info",
                "operationId": "GetServiceStatus",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Status"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/thread/{slug_or_id}/create": {
            "post": {
                "description": "Creates Posts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "Creates Posts",
                "operationId": "CreatePosts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "slug or id",
                        "name": "slug_or_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Posts params",
                        "name": "posts",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Post"
                            }
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request - Problem with the request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "404": {
                        "description": "Not found - Requested entity is not found in database",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "409": {
                        "description": "Conflict - User already exists",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/thread/{slug_or_id}/details": {
            "get": {
                "description": "Gets thread info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "Gets thread info",
                "operationId": "GetThreadInfo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "slug or id of thread",
                        "name": "slug_or_id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Thread"
                        }
                    },
                    "404": {
                        "description": "Not found - Requested entity is not found in database",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Updates thread info",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "Updates thread info",
                "operationId": "UpdateThreadInfo",
                "parameters": [
                    {
                        "type": "string",
                        "description": "slug or id of thread",
                        "name": "slug_or_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "ThreadUpdate params",
                        "name": "threadUpdate",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.ThreadUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Thread"
                        }
                    },
                    "404": {
                        "description": "Not found - Requested entity is not found in database",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/thread/{slug_or_id}/posts": {
            "get": {
                "description": "GetThreadPosts",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "GetThreadPosts",
                "operationId": "GetThreadPosts",
                "parameters": [
                    {
                        "type": "string",
                        "description": "slug or id of thread",
                        "name": "slug_or_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "since",
                        "name": "since",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "sort",
                        "name": "sort",
                        "in": "query"
                    },
                    {
                        "type": "boolean",
                        "description": "desc",
                        "name": "desc",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.Thread"
                            }
                        }
                    },
                    "404": {
                        "description": "Not found - Requested entity is not found in database",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/thread/{slug_or_id}/vote": {
            "post": {
                "description": "VoteForThread",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Thread"
                ],
                "summary": "VoteForThread",
                "operationId": "VoteForThread",
                "parameters": [
                    {
                        "type": "string",
                        "description": "slug or id of thread",
                        "name": "slug_or_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "vote params",
                        "name": "vote",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Vote"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Thread"
                        }
                    },
                    "404": {
                        "description": "Not found - Requested entity is not found in database",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/user/{nickname}/create": {
            "post": {
                "description": "Creates User",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Creates User",
                "operationId": "CreateUser",
                "parameters": [
                    {
                        "type": "string",
                        "description": "nickname of user",
                        "name": "nickname",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User params",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "400": {
                        "description": "Bad request - Problem with the request",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "409": {
                        "description": "Conflict - User already exists",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        },
        "/user/{nickname}/profile": {
            "get": {
                "description": "Gets Users profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Gets Users profile",
                "operationId": "GetProfile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "nickname of user",
                        "name": "nickname",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    },
                    "404": {
                        "description": "Not found - Requested entity is not found in database",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            },
            "post": {
                "description": "Changes Users profile",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "User"
                ],
                "summary": "Changes Users profile",
                "operationId": "PostProfile",
                "parameters": [
                    {
                        "type": "string",
                        "description": "nickname of user",
                        "name": "nickname",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "User params",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.Response"
                        }
                    },
                    "404": {
                        "description": "Not found - Requested entity is not found in database",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "409": {
                        "description": "Conflict - User already exists",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error - Request is valid but operation failed at server side",
                        "schema": {
                            "$ref": "#/definitions/model.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "model.Error": {
            "type": "object",
            "properties": {
                "error": {}
            }
        },
        "model.Forum": {
            "type": "object",
            "properties": {
                "posts": {
                    "type": "integer"
                },
                "slug": {
                    "type": "string"
                },
                "threads": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "model.ForumCreateModel": {
            "type": "object",
            "properties": {
                "slug": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "user": {
                    "type": "string"
                }
            }
        },
        "model.Post": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "forum": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "isedited": {
                    "type": "boolean"
                },
                "message": {
                    "type": "string"
                },
                "parent": {
                    "type": "integer"
                },
                "thread": {
                    "type": "integer"
                }
            }
        },
        "model.PostFull": {
            "type": "object",
            "properties": {
                "author": {
                    "$ref": "#/definitions/model.User"
                },
                "forum": {
                    "$ref": "#/definitions/model.Forum"
                },
                "post": {
                    "$ref": "#/definitions/model.Post"
                },
                "thread": {
                    "$ref": "#/definitions/model.Thread"
                }
            }
        },
        "model.PostUpdate": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                }
            }
        },
        "model.Response": {
            "type": "object",
            "properties": {
                "body": {}
            }
        },
        "model.Status": {
            "type": "object",
            "properties": {
                "forum": {
                    "type": "integer"
                },
                "post": {
                    "type": "integer"
                },
                "thread": {
                    "type": "integer"
                },
                "user": {
                    "type": "integer"
                }
            }
        },
        "model.Thread": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "forum": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                },
                "slug": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                },
                "votes": {
                    "type": "integer"
                }
            }
        },
        "model.ThreadCreateModel": {
            "type": "object",
            "properties": {
                "author": {
                    "type": "string"
                },
                "created": {
                    "type": "string"
                },
                "message": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.ThreadUpdate": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "model.User": {
            "type": "object",
            "properties": {
                "about": {
                    "type": "string"
                },
                "email": {
                    "type": "string"
                },
                "fullname": {
                    "type": "string"
                },
                "nickname": {
                    "type": "string"
                }
            }
        },
        "model.Vote": {
            "type": "object",
            "properties": {
                "nickname": {
                    "type": "string"
                },
                "voice": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "127.0.0.1:5000",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "DB project API",
	Description:      "DB project server.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
