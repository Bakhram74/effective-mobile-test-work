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
        "/create": {
            "post": {
                "description": "Handler for Createting new song",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Create Song",
                "parameters": [
                    {
                        "description": "song fields",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.songCreateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.Song"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "func"
                        }
                    }
                }
            }
        },
        "/delete": {
            "delete": {
                "description": "Handler for Deleting song by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Delete Song",
                "parameters": [
                    {
                        "description": "song id",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.songDeleteReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "ok",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "func"
                        }
                    }
                }
            }
        },
        "/songs": {
            "get": {
                "description": "Handler for getting Song verses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Get filtered Song",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ATTENTION!!! use date Instead of release_date, default dir is asc",
                        "name": "sort",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "asc or desc",
                        "name": "dir",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "page of songs",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "count of songs",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.SongResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "func"
                        }
                    }
                }
            }
        },
        "/update": {
            "put": {
                "description": "Handler for Updating song by id",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Update Song",
                "parameters": [
                    {
                        "description": "all feilds are optional exept id",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.songUpdateReq"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/db.Song"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "func"
                        }
                    }
                }
            }
        },
        "/verses": {
            "post": {
                "description": "Handler for getting Song verses",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "song"
                ],
                "summary": "Get paginated Song verses",
                "parameters": [
                    {
                        "description": "group, name",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/http.paginateSongVersesReq"
                        }
                    },
                    {
                        "type": "string",
                        "description": "page of verses",
                        "name": "page",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "count of verses",
                        "name": "limit",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.VerseResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "func"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "func"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "db.Song": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "entity.Pagination": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer"
                },
                "total_count": {
                    "type": "integer"
                },
                "total_page": {
                    "type": "integer"
                }
            }
        },
        "entity.SongResponse": {
            "type": "object",
            "properties": {
                "pagination": {
                    "$ref": "#/definitions/entity.Pagination"
                },
                "songs": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/db.Song"
                    }
                }
            }
        },
        "entity.VerseResponse": {
            "type": "object",
            "properties": {
                "group": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "pagination": {
                    "$ref": "#/definitions/entity.Pagination"
                },
                "verse": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "http.paginateSongVersesReq": {
            "type": "object",
            "required": [
                "group",
                "name"
            ],
            "properties": {
                "group": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                }
            }
        },
        "http.songCreateReq": {
            "type": "object",
            "required": [
                "group",
                "link",
                "name",
                "release_date",
                "text"
            ],
            "properties": {
                "group": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        },
        "http.songDeleteReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "id": {
                    "type": "string"
                }
            }
        },
        "http.songUpdateReq": {
            "type": "object",
            "required": [
                "id"
            ],
            "properties": {
                "group": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "link": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "release_date": {
                    "type": "string"
                },
                "text": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.0.1",
	Host:             "localhost:8080",
	BasePath:         "/song",
	Schemes:          []string{},
	Title:            "Music info",
	Description:      "API docs for Effective-mobile-test-work.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
