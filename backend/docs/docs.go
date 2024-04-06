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
        "/api/v1/tours": {
            "get": {
                "description": "Get tours according to search and filter params",
                "produces": [
                    "application/json",
                    "application/json"
                ],
                "tags": [
                    "tours"
                ],
                "summary": "Get tours according to search and filter params",
                "parameters": [
                    {
                        "type": "string",
                        "description": "From location name",
                        "name": "fromName",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "To location name",
                        "name": "toName",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Date of the tour",
                        "name": "when",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Number of nights in the tour",
                        "name": "nightsCnt",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Number of adults",
                        "name": "adults",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Number of children",
                        "name": "childrens",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Type of the tour",
                        "name": "tourType",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Minimum price",
                        "name": "priceFrom",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Maximum price",
                        "name": "priceTo",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Minimum rating",
                        "name": "rating",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "boolean",
                        "description": "Guaranteed availability",
                        "name": "guaranteed",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Feature 1",
                        "name": "feature1",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Feature 2",
                        "name": "feature2",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "Feature 3",
                        "name": "feature3",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Age group ID",
                        "name": "ageGroupId",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Tour difficulty level",
                        "name": "difficulty",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Comfort level",
                        "name": "comfort",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Food ID",
                        "name": "foodId",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/github_com_kldd0_travel-hack-2024_internal_entity.SimplifiedTourView"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        },
        "/api/v1/tours/{id}": {
            "get": {
                "description": "Get tour by id",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tours"
                ],
                "summary": "Get tour by id",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/github_com_kldd0_travel-hack-2024_internal_entity.Tour"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/echo.HTTPError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "echo.HTTPError": {
            "type": "object",
            "properties": {
                "message": {}
            }
        },
        "github_com_kldd0_travel-hack-2024_internal_entity.Image": {
            "type": "object",
            "properties": {
                "caption": {
                    "type": "string"
                },
                "url": {
                    "type": "string"
                }
            }
        },
        "github_com_kldd0_travel-hack-2024_internal_entity.SimplifiedTourView": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "location": {
                    "type": "string"
                },
                "night_count": {
                    "type": "integer"
                },
                "rating": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_kldd0_travel-hack-2024_internal_entity.Tag"
                    }
                }
            }
        },
        "github_com_kldd0_travel-hack-2024_internal_entity.Tag": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string"
                }
            }
        },
        "github_com_kldd0_travel-hack-2024_internal_entity.Tour": {
            "type": "object",
            "properties": {
                "category": {
                    "type": "string"
                },
                "comfort_level": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "difficulty_level": {
                    "type": "string"
                },
                "faq": {
                    "type": "string"
                },
                "group_dates": {
                    "description": "must be updated every req",
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_kldd0_travel-hack-2024_internal_entity.TourDate"
                    }
                },
                "id": {
                    "type": "integer"
                },
                "important_info": {
                    "type": "string"
                },
                "included": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "location": {
                    "type": "string"
                },
                "media": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_kldd0_travel-hack-2024_internal_entity.Image"
                    }
                },
                "nearest_date": {
                    "description": "must be updated every req",
                    "type": "string"
                },
                "night_count": {
                    "type": "integer"
                },
                "not_included": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "program": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "rating": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                },
                "type": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/github_com_kldd0_travel-hack-2024_internal_entity.Tag"
                    }
                }
            }
        },
        "github_com_kldd0_travel-hack-2024_internal_entity.TourDate": {
            "type": "object",
            "properties": {
                "end": {
                    "type": "string"
                },
                "start": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "JWT": {
            "description": "JWT token",
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Tour Management Service",
	Description:      "Additional service for RUSSPASS",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	// LeftDelim:        "{{",
	// RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
