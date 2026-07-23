// Package response provides framework-agnostic HTTP response helpers.
// It uses net/http instead of Gin so it stays reusable across any HTTP framework
// (Gin, Echo, Fiber, etc.) without tying shared code to a specific framework.
package response

import (
	"encoding/json"
	"net/http"
)

// Envelope is the standard wrapper for all API payloads.
// Every response uses the same shape: { "data": ..., "error": ... }.
// The omitempty tags ensure empty fields are omitted from JSON
// so the client doesn't receive null values unnecessarily.
type Envelope struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

// OK writes a successful JSON response with the given status and data.
// It sets Content-Type to application/json and writes the envelope.
func OK(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Envelope{Data: data})
}

// Fail writes an error JSON response with the given status and message.
// It sets Content-Type to application/json and writes the envelope.
func Fail(w http.ResponseWriter, status int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(Envelope{Error: message})
}
