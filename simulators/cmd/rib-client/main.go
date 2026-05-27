package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"

	"gerrit-gamma.gic.ericsson.se/pc/eric-pc-routing-information-base/tests/veto/simulators/internal/httputils"
	"gerrit-gamma.gic.ericsson.se/pc/eric-pc-routing-information-base/tests/veto/simulators/internal/vetosim"
)

const (
	plugName      = "rib_client"
	serverAddress = ":8888"
)

var errCmdNotImplemented = errors.New("that command is not implemented by the simulator")

type ribClientSim struct{}

func main() {
	vetosim.InitLogger(plugName)
	vetosim.ServePlugin(ribClientSim{})
}

func (se ribClientSim) Start() error {
	mux := http.NewServeMux()

	srv, err := httputils.NewHTTPServer(serverAddress, mux)
	if err != nil {
		return fmt.Errorf("creating TLS server: %w", err)
	}

	go func() {
		if err := srv.ListenAndServeTLS("", ""); err != nil {
			slog.Error("ListenAndServeTLS failed", "error", err)
		}
	}()

	slog.Debug("RIB Client simulator started successfully")

	return nil
}

func (se ribClientSim) GetStats() (map[string]uint64, error) {
	return map[string]uint64{}, nil
}

func (se ribClientSim) Configure(_ string) error {
	return nil
}

func (se ribClientSim) RecvEvent(_ string) error {
	return nil
}

func (se ribClientSim) Command(str string) (string, error) {
	command, err := vetosim.ExtractCommand(str)
	if err != nil {
		return "", err
	}
	handled, ret, err := vetosim.TryHandlingCommand(command)
	if handled {
		return ret, err
	}

	return "", errCmdNotImplemented
}
