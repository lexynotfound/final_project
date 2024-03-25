package docs

// Documentation API MY GRAM
var Docs = `{
	"schemes": {{ marshal .Schemes }},
	"swagger": "2.0",
	"info": {
		"description": "{{escape .Description}}",
		"title": "{{.Title}}",
		"version": "{{.Version}}"
	},
	"host": "{{.Host}}",
	"basePath": "{{.BasePath}}",
	"paths": {
		"/users":{
			"get": {
				"description": "Get list of users",
				"tags"; ["users"],
				"parameters": [],
				"response": {
					"200": {
						"description": "Sucess"
					},
					"400": {
						"description": "Bad Request"
					},
					"500": {
						"description": "Internal server Error"
					}
				}
			},
			"put": {
				"description": "Update user",
				"tags": ["users"],
				"parameters": [
					{
						"name": "Authorization",
                        "in": "header",
                        "description": "Bearer token string",
                        "required": true,
						"type": "string
                        "schema": {
							"$ref": "#/definitions/UserUpdate"
						}
					},
					{
						"name": "id",
						"in": "path",
						"description": "ID of users update",
						"required": true,
						"type": "integer"
					}
					{
						"name": "body",
                        "in": "body",
                        "description": "Users Detail,
                        "required": true,
                        "schema": {
							"$ref": "#/definitions/UserUpdate"
						}
					}
				],
				"response": {
					"200": {
						"description": "User updated successfully"
					},
					"400": {
						"description": "Bad Request"
					},
					"401": {
						"description": "Unauthorized"
					},
					"500": {
						"description": "Internal server error"
					}
				}
			},
			"delete" : {
				"description": "Delet user",
				"tags": ["users"],
				"parameters": [],
				"response": {
					"200": {
						"description": "User updated successfully"
					},
					"400": {
						"description": "Bad Request"
					},
					"401": {
						"description": "Unauthorized"
					},
					"500": {
						"description": "Internal server error"
					}
				}
			}
		},
		"/photos": {
            "get": {
                "description": "Get list of photos",
                "tags": ["photos"],
                "parameters": [],
                "responses": {
                    "200": {
                        "description": "Success"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            },
            "post": {
                "description": "Create a new photo",
                "tags": ["photos"],
                "parameters": [
                    {
                        "name": "Authorization",
                        "in": "header",
                        "description": "Bearer token string",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Photo"
                        }
                    },
					{
						"name": "body",
                        "in": "body",
                        "description": "Photo details",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Photo"
                        }
					}
                ],
                "responses": {
                    "201": {
                        "description": "Photo created successfully"
                    },
                    "400": {
                        "description": "Bad request"
                    },
                    "500": {
                        "description": "Internal server error"
                    }
                }
            }
        },
		"/photos/{photoId}": {
			"put": {
				"description": "Update photo",
                "tags": ["photos"],
                "parameters": [
					{
						"name": "Authorization",
                        "in": "header",
                        "description": "Bearer token string",
                        "required": true,
                        "type": "string"
					},
					{
						"name": "photoId",
                        "in": "path",
                        "description": "ID of the photo to update",
                        "required": true,
                        "type": "integer"
					},
					{
						"name": "body",
                        "in": "body",
                        "description": "Updated photo details",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/Photo"
                        }
					}
				],
				"response": {
					"200": {
						"description": "Photo updated successfully"
					},
					"400": {
						"description": "Bad Request"
					},
					"401": {
						"description": "Unauthorized"
					},
					"500": {
						"description": "Internal server error"
					},
				}
			},
			"delete": {
				"description": "Delete photo",
                "tags": ["photos"],
                "parameters": [
					{
						"name": "Authorization",
                        "in": "header",
                        "description": "Bearer token string",
                        "required": true,
                        "type": "string"
					},
					{
						"name": "photoId",
                        "in": "path",
                        "description": "ID of the photo to delete",
                        "required": true,
                        "type": "integer"
					}
				],
				"response" : {
					"200": {
						"description": "Photo deleted successfully"
					},
					"400": {
						"description": "Bad Request"
					},
					"500": {
						"description": "Internal server error"
					}
				}
			}
		},
		"/comments": {
			"get": {
				"description": "Get list of comments",
                "tags": ["comments"],
                "parameters": [
                    {
                        "name": "Authorization",
                        "in": "header",
						"description": "Bearer token string",
						"required": true,
						"type": "string"
                    }
                ],
				"response": {
					"200": {
						"description": "Success",
						"schema": {
							"type": "object"
						}
					},
					"400": {
						"description": "Bad Request"
					},
					"500": {
						"description": "Internal server error",
					}
				}
			}
		},
		"/comments/{commentId}": {
			"put": { 
				"description": "Update comment",
				"tags": ["comments"],
				"parameters": [
					{
						"name": "Authorization",
						"in": "header",
						"description": "Bearer token string",
						"required": true,
						"type": "string"
					},

					{
						"name": "commentId",
						"in": "path",
						"description": "ID of the comment to update",
						"required": true,
						"type": "integer"
					},
					{
						"name": "body",
						"in": "body",
						"description": "Updated comment details",
						"required": true,
						"schema": {
							"$ref": "#/definitions/Comment"
						}
					}
				],

				"response": {
					"200": {
						"description": "Comment updated successfully"
					},
					"400": {
						"description": "Bad Request"
					},
					"500": {
						"description": "Internal server error"
					},
				}
			},
			"delete" : {
				"description": "Delete comment",
				"tags": ["comments"],
				"parameters": [
					{ 
						"name": "Authorization",
						"in": "header",
						"description": "Bearer token string",
						"required": true,
						"type": "string"
					},
					{ 
						"name": "commentId",
						"in": "path",
						"description": "ID of the comment to delete",
						"required": true,
						"type": "integer"
					}
				],
				"responses": {
					"200": {
						"description": "Comment deleted successfully"
					},
					"400": {
						"description": "Bad Request"
					},
					"500": {
						"description": "Internal server error"
					}
				}
			}
		},
		"/socialmedias": {
			"get": {
					"description": "Create a new social media entry",
					"tags": ["socialmedias"],
					"parameters": [
						{
							"name": "Authorization",
							"in": "header",
							"description": "Bearer token string",
							"required": true,
							"type": "string"
						}
					],
				"responses": 
				{
					"201": {
						"description": "Social media entry created successfully"
					},
					"400": {
						"description": "Bad request"
					},
					"500": {
						"description": "Internal server error"
					}
				}
			},
			"post": {
					"description": "Create a new social media entry",
					"tags": ["socialmedias"],
					"parameters": [
						{
							"name": "Authorization",
							"in": "header",
							"description": "Bearer token string",
							"required": true,
							"type": "string"
						},
						{
							"name": "body",
							"in": "body",
							"description": "Social media details",
							"required": true,
							"schema": {
								"$ref": "#/definitions/SocialMedia"
							}
						}
					],
				"responses": 
				{
					"201": {
						"description": "Social media entry created successfully"
					},
					"400": {
						"description": "Bad request"
					},
					"500": {
						"description": "Internal server error"
					}
				}
			}
		},
		"/socialmedias/{socialMediaId}:{
			"put": {
				"description": "Update social media link",
				"tags": ["socialmeidas"],
				"parameters": [
					{
						"name": "Authorization",
						"in": "header",
						"description": "Bearer token string",
						"required": true,
						"type": "string",
					},
					{
						"name": "socialMediaId",
						"in": "path",
						"description": "ID of the social media entry to update",
						"required": true,
						"type": "integer",
					},
					{
						"name": "body",
						"in": "body",
						"description": "update social media details",
						"required": true,
						"schema": {
							"$ref": "#/definitions/SocialMedia"
						}
					}
				],
				"responses": {
					"200": {
						"description": "Social media entry updated successfully"
					},
					"400": {
						"description": "Bad request"
					},
					"500": {
						"description": "Internal server error"
					}
				}
			},
			"delete": {
				"description": "Delete social media",
				"tags": ["socialmeidas"],
				"parameters": [
					{
						"name": "Authorization",
						"in": "header",
						"description": "Bearer token string",
						"required": true,
						"type": "string",
					},
					{
						"name": "socialMediaId",
						"in": "path",
						"description": "ID of the social media entry to update",
						"required": true,
						"type": "integer",
					}
				],
				"responses": {
					"200": {
						"description": "Social media entry updated successfully"
					},
					"400": {
						"description": "Bad request"
					},
					"500": {
						"description": "Internal server error"
					}
				}
			}
		}
	}
}`
