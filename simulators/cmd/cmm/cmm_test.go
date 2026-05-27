package main

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleDataRegister(t *testing.T) {
	sim := NewCmmSim()

	req := httptest.NewRequest(http.MethodPut, "/cm/api/v1/schemas/ietf-network-instance/data-sources/rib", http.NoBody)
	w := httptest.NewRecorder()

	sim.handleDataRegister(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

func TestHandleDataRegisterWrongMethod(t *testing.T) {
	sim := NewCmmSim()

	req := httptest.NewRequest(http.MethodGet, "/cm/api/v1/schemas/ietf-network-instance/data-sources/rib", http.NoBody)
	w := httptest.NewRecorder()

	sim.handleDataRegister(w, req)

	resp := w.Result()
	defer resp.Body.Close()

	require.Equal(t, http.StatusMethodNotAllowed, resp.StatusCode)
}

func TestGetStats(t *testing.T) {
	sim := NewCmmSim()

	// Initial stats should be 0
	stats, err := sim.GetStats()
	require.NoError(t, err)
	assert.Equal(t, uint64(0), stats["rxRequests"])

	// Simulate a PUT request
	req := httptest.NewRequest(http.MethodPut, "/cm/api/v1/schemas/ietf-network-instance/data-sources/rib", http.NoBody)
	w := httptest.NewRecorder()
	sim.handleDataRegister(w, req)

	// Stats should increment
	stats, err = sim.GetStats()
	require.NoError(t, err)
	assert.Equal(t, uint64(1), stats["rxRequests"])
}

func TestCommand(t *testing.T) {
	sim := NewCmmSim()

	// Test set-log-level command (handled by shared)
	result, err := sim.Command(`{"command": "set-log-level", "parameters": {"level": "debug"}}`)
	require.NoError(t, err)
	assert.NotEmpty(t, result)

	// Test unknown command (should return empty JSON)
	result, err = sim.Command(`{"command": "unknown-command", "parameters": {}}`)
	require.NoError(t, err)
	assert.Equal(t, "{}", result)
}
