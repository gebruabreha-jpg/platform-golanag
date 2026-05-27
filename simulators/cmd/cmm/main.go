// CMM simulator for RIB service.
package main

import (
	"errors"
	"fmt"
	"gerrit-gamma.gic.ericsson.se/pc/eric-pc-routing-information-base/tests/veto/simulators/internal/httputils"
	"gerrit-gamma.gic.ericsson.se/pc/eric-pc-routing-information-base/tests/veto/simulators/internal/vetosim"
	"log/slog"
	"net/http"
	"sync/atomic"
)

const (
	pluginName    = "CMM"
	serverAddress = ":5004"
	registerPath  = "/cm/api/v1/schemas/ietf-network-instance/data-sources/rib"
)

type CmmSim struct {
	server     *http.Server
	rxRequests atomic.Uint64
}

func NewCmmSim() *CmmSim {
	return &CmmSim{}
}

// Handles the PUT request to register RIB yang paths.
func (s *CmmSim) handleDataRegister(w http.ResponseWriter, r *http.Request) {
	slog.Debug("Received request", "method", r.Method, "path", r.URL.Path)

	if r.Method != http.MethodPut {
		slog.Warn("Unexpected HTTP method", "method", r.Method, "expected", http.MethodPut)
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	count := s.rxRequests.Add(1)
	slog.Info("RIB data source registration received", "count", count)

	// Respond with 200 OK
	w.WriteHeader(http.StatusOK)
}

// Starts the simulator HTTP server.
func (s *CmmSim) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc(registerPath, s.handleDataRegister)

	srv, err := httputils.NewHTTPServer(serverAddress, mux)
	if err != nil {
		return fmt.Errorf("creating TLS server: %w", err)
	}

	s.server = srv

	go func() {
		slog.Info("Starting CMM simulator", "addr", srv.Addr)
		if err := srv.ListenAndServeTLS("", ""); err != nil && !errors.Is(err, http.ErrServerClosed) {
			slog.Error("Server failed", "error", err)
		}
	}()

	slog.Info("CMM simulator started successfully")

	return nil
}

// Returns the statistics of the simulator.
func (s *CmmSim) GetStats() (map[string]uint64, error) {
	stats := make(map[string]uint64)
	stats["rxRequests"] = s.rxRequests.Load()

	return stats, nil
}

// Configures the simulator (no configuration needed for this simulator).
func (s *CmmSim) Configure(str string) error {
	if str == "" {
		slog.Debug("Configure called with empty string (reset)")

		return nil
	}

	slog.Debug("Configure called", "config", str)

	return nil
}

// Receives an event (not used).
func (s *CmmSim) RecvEvent(str string) error {
	slog.Debug("RecvEvent called", "event", str)

	return nil
}

// Executes a command (no internal commands needed for this simulator).
func (s *CmmSim) Command(str string) (string, error) {
	command, err := vetosim.ExtractCommand(str)
	if err != nil {
		return "", fmt.Errorf("command extraction failed: %w", err)
	}

	// Try handling common commands
	if handled, ret, err := vetosim.TryHandlingCommand(command); handled {
		return ret, err
	}

	slog.Debug("Received command", "command", command.Command)

	return "{}", nil
}

func main() {
	vetosim.InitLogger(pluginName)
	cmmSim := NewCmmSim()
	vetosim.ServePlugin(cmmSim)
}
