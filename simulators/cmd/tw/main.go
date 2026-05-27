package main

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strings"
	"sync"
	"sync/atomic"
	"time"

	"gerrit-gamma.gic.ericsson.se/pc/eric-pc-routing-information-base/tests/veto/simulators/internal/httputils"
	"gerrit-gamma.gic.ericsson.se/pc/eric-pc-routing-information-base/tests/veto/simulators/internal/vetosim"
)

const (
	plugName          = "TW"
	serverAddress     = ":8443"
	timeoutDuration   = 1000000 // nanoseconds per millisecond
	oneshotTimerType  = "oneshot"
	periodicTimerType = "periodic"
)

var (
	timerList   = make(map[string]*timer)
	timerListMu sync.RWMutex
	httpClient  *http.Client

	// use atomic counter to avoid races.
	activeTimers uint64

	errCmdNotImplemented = errors.New("that command is not implemented by the simulator")
)

type timerTime struct {
	Epoch int64 `json:"epoch"`
	Ms    int64 `json:"ms"`
}

type timerData struct {
	ID        string    `json:"id"`
	Type      string    `json:"type"`
	Geored    bool      `json:"geored"`
	Timeout   timerTime `json:"timeout"`
	Duration  int32     `json:"duration"`
	ExpireURI string    `json:"expire_uri"`
	ExpireARG string    `json:"expire_arg"`
	Update    bool      `json:"update"`
}

type timer struct {
	mu     sync.Mutex
	info   timerData
	tmr    *time.Timer
	ticker *time.Ticker
	cancel context.CancelFunc
}

type TwSim struct {
	server *http.Server
}

func calculateTimeout(epoch, ms int64) time.Duration {
	timeout := time.Unix(epoch, ms*timeoutDuration)

	d := time.Until(timeout)
	if d <= 0 {
		return 0
	}

	return d
}

func sendTimerExpireReq(info timerData) {
	if httpClient == nil {
		// initialize HTTP client when starting
		slog.Error("HTTP client is nil in sendTimerExpireReq")

		return
	}

	byteReq, err := json.Marshal(info)
	if err != nil {
		slog.Error("Error in Marshal", "error", err)

		return
	}

	request, err := http.NewRequestWithContext(
		context.Background(),
		http.MethodPost,
		info.ExpireURI,
		bytes.NewReader(byteReq),
	)
	if err != nil {
		slog.Error("New request error", "error", err.Error())

		return
	}

	response, err := httpClient.Do(request)
	if err != nil {
		slog.Error("Error while sending timer callback", "url", request.URL, "error", err)

		return
	}
	defer response.Body.Close()
	_, _ = io.Copy(io.Discard, response.Body) // drain body without allocating
}

func runTimer(ctx context.Context, t *timer) {
	defer func() {
		if t.tmr != nil {
			t.tmr.Stop()
		}
		if t.ticker != nil {
			t.ticker.Stop()
		}
	}()

	if t.tmr != nil {
		select {
		case <-t.tmr.C:
			t.mu.Lock()
			info := t.info
			t.mu.Unlock()
			sendTimerExpireReq(info)
		case <-ctx.Done():
		}

		return
	}

	if t.ticker != nil {
		for {
			select {
			case <-t.ticker.C:
				t.mu.Lock()
				info := t.info
				t.mu.Unlock()
				sendTimerExpireReq(info)
			case <-ctx.Done():

				return
			}
		}
	}
}

func (tw *TwSim) handleCreateTw(w http.ResponseWriter, r *http.Request) {
	var input timerData

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		slog.Error("Error in Decode", "error", err)
		http.Error(w, "invalid json", http.StatusBadRequest)

		return
	}

	// After decoding input
	if input.Type != oneshotTimerType && input.Type != periodicTimerType {
		http.Error(w, "invalid timer type", http.StatusBadRequest)

		return
	}

	if input.ID == "" {
		slog.Error("missing timer id")
		http.Error(w, "missing timer id", http.StatusBadRequest)

		return
	}

	if input.Update {
		tw.handleUpdateTimer(w, &input)

		return
	}

	tw.handleCreateTimer(w, &input)
}

func (tw *TwSim) handleUpdateTimer(w http.ResponseWriter, input *timerData) {
	timerListMu.RLock()
	t, ok := timerList[input.ID]
	timerListMu.RUnlock()

	if !ok {
		slog.Error("Error not find the update timer id")
		http.Error(w, "timer id not found for update", http.StatusNotFound)

		return
	}

	if t == nil {
		slog.Error("Loaded timer is nil", "id", input.ID)
		http.Error(w, "found timer but it was nil", http.StatusInternalServerError)

		return
	}

	if t.tmr != nil {
		t.mu.Lock()
		d := calculateTimeout(input.Timeout.Epoch, input.Timeout.Ms)
		if d <= 0 {
			slog.Warn("Oneshot timer timeout is in the past, firing immediately", "id", input.ID)
		}
		t.info = *input
		t.tmr.Reset(d)
		t.mu.Unlock()
		writeResponse(w, http.StatusOK, input)

		return
	}

	if t.ticker != nil {
		if input.Duration <= 0 {
			slog.Error("Error invalid duration value")
			http.Error(w, "invalid duration for periodic timer", http.StatusBadRequest)

			return
		}

		t.mu.Lock()
		t.info = *input
		t.mu.Unlock()
		d := time.Duration(input.Duration) * time.Millisecond
		t.ticker.Reset(d)
	}

	writeResponse(w, http.StatusOK, input)
}

func (tw *TwSim) handleCreateTimer(w http.ResponseWriter, input *timerData) {
	timerListMu.RLock()
	_, ok := timerList[input.ID]
	timerListMu.RUnlock()

	if ok {
		writeResponse(w, http.StatusFound, input)

		return
	}

	t := &timer{
		info: *input,
	}

	switch input.Type {
	case "oneshot":
		d := calculateTimeout(input.Timeout.Epoch, input.Timeout.Ms)
		t.tmr = time.NewTimer(d)

	case "periodic":
		if input.Duration <= 0 {
			http.Error(w, "invalid duration for periodic timer", http.StatusBadRequest)

			return
		}

		d := time.Duration(input.Duration) * time.Millisecond
		t.ticker = time.NewTicker(d)
	}

	ctx, cancel := context.WithCancel(context.Background())
	t.cancel = cancel
	timerListMu.Lock()
	timerList[input.ID] = t
	timerListMu.Unlock()
	go runTimer(ctx, t)

	writeResponse(w, http.StatusCreated, input)

	atomic.AddUint64(&activeTimers, 1)
}

func writeResponse(w http.ResponseWriter, status int, v any) {
	var (
		byteReq []byte
		err     error
	)

	byteReq, err = json.Marshal(v)
	if err != nil {
		slog.Error("Error in Marshal", "error", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	fmt.Fprintf(w, "%s\n", string(byteReq))
}

func removeTw(id string) bool {
	timerListMu.RLock()
	t, ok := timerList[id]
	timerListMu.RUnlock()

	if !ok {
		slog.Warn("removeTw: timer not found", "id", id)

		return false
	}

	if t == nil {
		slog.Warn("removeTw: timer is nil", "id", id)
		timerListMu.Lock()
		delete(timerList, id)
		timerListMu.Unlock()

		return false
	}

	if t.tmr != nil {
		t.tmr.Stop()
	}
	if t.ticker != nil {
		t.ticker.Stop()
	}

	t.cancel()
	timerListMu.Lock()
	delete(timerList, id)
	timerListMu.Unlock()

	return true
}

func (tw *TwSim) handleTimers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		if r.URL.Path != "/v1/timers" {
			http.NotFound(w, r)

			return
		}
		tw.handleCreateTw(w, r)

	case http.MethodDelete:
		tw.handleRemoveTw(w, r)

	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
	}
}

func (tw *TwSim) handleRemoveTw(w http.ResponseWriter, r *http.Request) {
	const base = "/v1/timers"

	path := r.URL.Path
	if !strings.HasPrefix(path, base) {
		http.NotFound(w, r)

		return
	}

	id := path[len(base):]
	if id == "" || id == "/" {
		slog.Warn("missing timer id in delete path", "path", r.URL.Path)
		http.Error(w, "missing timer id in delete path", http.StatusBadRequest)

		return
	}

	rawID := strings.TrimPrefix(id, "/")
	realID := rawID
	if strings.HasPrefix(rawID, "v1/") {
		realID = "/" + rawID
	}

	var input timerData

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil && !errors.Is(err, io.EOF) {
		slog.Error("Error in Decode", "error", err)
	}

	removed := removeTw(realID)

	byteReq, err := json.Marshal(input)
	if err != nil {
		slog.Error("Error in Marshal", "error", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", string(byteReq))

	if removed {
		atomic.AddUint64(&activeTimers, ^uint64(0)) // -1
	}
}

func main() {
	vetosim.InitLogger(plugName)
	twSim := &TwSim{}
	vetosim.ServePlugin(twSim)
}

func (tw *TwSim) Start() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/v1/timers", tw.handleTimers)
	mux.HandleFunc("/v1/timers/", tw.handleTimers)

	var err error

	httpClient, err = httputils.NewHTTPClient()
	if err != nil {
		return fmt.Errorf("creating HTTP client: %w", err)
	}

	tw.server, err = httputils.NewHTTPServer(serverAddress, mux)
	if err != nil {
		return fmt.Errorf("creating HTTP server: %w", err)
	}

	go func() {
		if serveErr := tw.server.ListenAndServeTLS("", ""); serveErr != nil {
			slog.Error("Error in ListenAndServe", "error", serveErr)
		}
	}()

	slog.Info("Started TW simulator")

	return nil
}

func (tw *TwSim) GetStats() (map[string]uint64, error) {
	res := make(map[string]uint64, 1)
	res["activeTimers"] = atomic.LoadUint64(&activeTimers)

	return res, nil
}

func (tw *TwSim) Configure(str string) error {
	if str == "" {
		return nil
	}

	return nil
}

func (tw *TwSim) RecvEvent(str string) error {
	slog.Debug("Receive event", "raw", str)

	return nil
}

func (tw *TwSim) Command(str string) (string, error) {
	command, err := vetosim.ExtractCommand(str)
	if err != nil {
		return "", err
	}

	handled, ret, err := vetosim.TryHandlingCommand(command)
	if handled {
		return ret, err
	}

	slog.Debug("New command request", "command", command)

	return "", errCmdNotImplemented
}
