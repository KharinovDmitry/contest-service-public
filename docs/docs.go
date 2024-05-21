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
        "/api/launches/success/{user_id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "tags": [
                    "Launch"
                ],
                "summary": "Получение результатов успешных решений задач конкретным пользователем",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
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
                                "$ref": "#/definitions/dto.LaunchDTO"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request"
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/launches/{user_id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "tags": [
                    "Launch"
                ],
                "summary": "Получение результатов решения задач конкретным пользователем",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
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
                                "$ref": "#/definitions/dto.LaunchDTO"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/launches/{user_id}/{contest_id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "tags": [
                    "Launch"
                ],
                "summary": "Получение результатов решения конкретной задачи конкретным пользователем",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "user_id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Contest ID",
                        "name": "contest_id",
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
                                "$ref": "#/definitions/dto.LaunchDTO"
                            }
                        }
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/run": {
            "post": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Run Test"
                ],
                "summary": "Отправка кода на проверку",
                "parameters": [
                    {
                        "description": "инормация о проверяемом коде",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.RunTestRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.RunTestResponse"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/handlers.RunError"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.RunError"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.RunError"
                        }
                    }
                }
            }
        },
        "/api/task": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Добавление задачи",
                "parameters": [
                    {
                        "description": "Информация о задаче",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.AddTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/task/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Получение задачи",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.TaskDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Удаление задачи",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Изменение задачи",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Новая информация о задаче",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.AddTaskRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/tasks": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "tags": [
                    "Task"
                ],
                "summary": "Получение всех задач",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.TaskDTO"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/test": {
            "put": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "Test"
                ],
                "summary": "Добавление теста",
                "parameters": [
                    {
                        "description": "Информация о тесте",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handlers.AddTestRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/test/{id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "tags": [
                    "Test"
                ],
                "summary": "Получение теста по айди",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Test ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.TestDTO"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "tags": [
                    "Test"
                ],
                "summary": "Удаление теста",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Test ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            },
            "patch": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "tags": [
                    "Test"
                ],
                "summary": "Обновление информации о тесте",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Test ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Новая нформация о тесте",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.TestDTO"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "403": {
                        "description": "Forbidden"
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/tests": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "tags": [
                    "Test"
                ],
                "summary": "Получение всех тестов",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/dto.TestDTO"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        },
        "/api/tests/{task_id}": {
            "get": {
                "security": [
                    {
                        "Bearer": []
                    }
                ],
                "tags": [
                    "Test"
                ],
                "summary": "Получение всех тестов для конкретной задачи",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Task ID",
                        "name": "task_id",
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
                                "$ref": "#/definitions/dto.TestDTO"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error"
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.LaunchDTO": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "contest_id": {
                    "type": "integer"
                },
                "date": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "points": {
                    "type": "integer"
                },
                "result_code": {
                    "description": "@Model domain.TestResultCode",
                    "allOf": [
                        {
                            "$ref": "#/definitions/enum.TestResultCode"
                        }
                    ]
                },
                "user_id": {
                    "type": "integer"
                }
            }
        },
        "dto.TaskDTO": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "memory_limit": {
                    "type": "integer"
                },
                "text": {
                    "type": "string"
                },
                "time_limit": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "dto.TestDTO": {
            "type": "object",
            "properties": {
                "expectedResult": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "input": {
                    "type": "string"
                },
                "points": {
                    "type": "integer"
                },
                "taskID": {
                    "type": "integer"
                }
            }
        },
        "enum.Language": {
            "description": "CPP - C++, Python - Python,",
            "type": "string",
            "enum": [
                "CPP",
                "Python",
                "C#",
                "JS"
            ],
            "x-enum-varnames": [
                "CPP",
                "Python",
                "CSharp",
                "JavaScript"
            ]
        },
        "enum.TestResultCode": {
            "description": "TL - Превышено ограничение по времени, ML - Превышено ограничение по памяти, CE - Ошибка компиляции, RE - Ошибка во время выполнения, SC - Успешное выполнение, IA - Неверный ответ",
            "type": "string",
            "enum": [
                "TL",
                "ML",
                "CE",
                "RE",
                "SC",
                "IA"
            ],
            "x-enum-varnames": [
                "TimeLimitCode",
                "MemoryLimitCode",
                "CompileErrorCode",
                "RuntimeErrorCode",
                "SuccessCode",
                "IncorrectAnswerCode"
            ]
        },
        "handlers.AddTaskRequest": {
            "type": "object",
            "properties": {
                "text": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "handlers.AddTestRequest": {
            "type": "object",
            "properties": {
                "expectedResult": {
                    "type": "string"
                },
                "input": {
                    "type": "string"
                },
                "points": {
                    "type": "string",
                    "example": "0"
                },
                "taskID": {
                    "type": "string",
                    "example": "0"
                }
            }
        },
        "handlers.RunError": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "message": {
                    "type": "string"
                }
            }
        },
        "handlers.RunTestRequest": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "string"
                },
                "language": {
                    "description": "@Model enum.Language",
                    "allOf": [
                        {
                            "$ref": "#/definitions/enum.Language"
                        }
                    ]
                },
                "task_id": {
                    "type": "integer"
                }
            }
        },
        "handlers.RunTestResponse": {
            "type": "object",
            "properties": {
                "description": {
                    "type": "string"
                },
                "points": {
                    "type": "string",
                    "example": "0"
                },
                "result_code": {
                    "description": "@Model enum.TestResultCode",
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "Bearer": {
            "description": "Type \"Bearer\" followed by a space and JWT token.",
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
	Title:            "Contest Service API",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}