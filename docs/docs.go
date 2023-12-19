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
        "/track": {
            "post": {
                "description": "Create a new track record in the database",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Create a new track",
                "operationId": "create-track",
                "parameters": [
                    {
                        "description": "Track details to create",
                        "name": "trackDetails",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TrackDetails"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/model.TrackDetails"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    },
                    "409": {
                        "description": "Track with ISRC code already exists",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/track/artist/{artist_name}": {
            "get": {
                "description": "Search for tracks by artist name using the Spotify API",
                "produces": [
                    "application/json"
                ],
                "summary": "Search for tracks by artist name",
                "operationId": "search-track-by-artist",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Name of the artist",
                        "name": "artist_name",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/model.TrackDetails"
                            }
                        }
                    },
                    "404": {
                        "description": "tracks not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            }
        },
        "/track/{isrc}": {
            "get": {
                "description": "Get track details from the database or Spotify by ISRC code",
                "produces": [
                    "application/json"
                ],
                "summary": "Get track details by ISRC",
                "operationId": "get-track-by-isrc",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ISRC code of the track",
                        "name": "isrc",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TrackDetails"
                        }
                    },
                    "404": {
                        "description": "Track not found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": true
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing track record in the database by ISRC",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Update a track by ISRC",
                "parameters": [
                    {
                        "type": "string",
                        "description": "ISRC code of the track to be updated",
                        "name": "isrc",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Updated track details",
                        "name": "trackDetails",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.TrackDetails"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/model.TrackDetails"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
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
        "model.TrackDetails": {
            "type": "object",
            "properties": {
                "artist_name": {
                    "type": "string"
                },
                "isrc": {
                    "type": "string"
                },
                "spotify_image": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Spotify API",
	Description:      "Golang application using Gin and GORM to interact with the Spotify API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
