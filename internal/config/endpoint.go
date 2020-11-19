package config

import "strings"

const defaultStatusCode = 200

// Endpoint : Represents an endpoint that will run in the server.
type Endpoint struct {
	Method     string `json:"method"`
	StatusCode int    `json:"status_code"`
	Route      string `json:"route"`
	BodyPath   string `json:"body_path"`
}

// IsNoContent : If no path is provided endpoint will be a no content type.
func (e Endpoint) IsNoContent() bool {
	return e.BodyPath == ""
}

// GetStatusCode : Returns the "default" 200 if status has a zero value otherwise returns the given status code.
func (e Endpoint) GetStatusCode() int {
	if e.StatusCode == 0 {
		return defaultStatusCode
	}
	return e.StatusCode
}

// GetMethod : Returns the endpoint method (lowercase).
func (e Endpoint) GetMethod() string {
	return strings.ToLower(e.Method)
}
