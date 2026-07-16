// Package response defines a consistent JSON envelope for API responses so
// every handler returns the same shape regardless of success or error.
package response

import "github.com/gin-gonic/gin"

// Envelope is the standard wrapper for all API payloads.
type Envelope struct {
	Data  any    `json:"data,omitempty"`
	Error string `json:"error,omitempty"`
}

// OK writes a 200 response with data in the envelope.
func OK(c *gin.Context, status int, data any) {
	c.JSON(status, Envelope{Data: data})
}

// Fail writes an error response with the given HTTP status and message.
func Fail(c *gin.Context, status int, message string) {
	c.JSON(status, Envelope{Error: message})
}
