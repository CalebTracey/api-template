// Package openapi3 provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.12.0 DO NOT EDIT.
package openapi3

// ErrorLog defines model for ErrorLog.
type ErrorLog struct {
	Query     *string `json:"query"`
	RootCause *string `json:"rootCause"`
	Scope     *string `json:"scope"`
	Status    *string `json:"status"`
	Trace     *string `json:"trace"`
}

// ErrorLogs defines model for ErrorLogs.
type ErrorLogs = []interface{}

// Message defines model for Message.
type Message struct {
	Count     *string    `json:"count"`
	ErrorLog  *ErrorLogs `json:"errorLog,omitempty"`
	HostName  *string    `json:"hostName"`
	Status    *string    `json:"status"`
	TimeTaken *string    `json:"timeTaken"`
}

// PSQLResponse defines model for PSQLResponse.
type PSQLResponse struct {
	LastInsertID *string  `json:"lastInsertID,omitempty"`
	Message      *Message `json:"message,omitempty"`
	RowsAffected *string  `json:"rowsAffected"`
}

// PSQLRequest defines model for PSQLRequest.
type PSQLRequest struct {
	Id          *string `json:"id,omitempty"`
	RequestType *string `json:"requestType,omitempty"`
	Table       *string `json:"table,omitempty"`
}

// AddToDatabaseJSONBody defines parameters for AddToDatabase.
type AddToDatabaseJSONBody struct {
	Id          *string `json:"id,omitempty"`
	RequestType *string `json:"requestType,omitempty"`
	Table       *string `json:"table,omitempty"`
}

// AddToDatabaseJSONRequestBody defines body for AddToDatabase for application/json ContentType.
type AddToDatabaseJSONRequestBody AddToDatabaseJSONBody
