package vetosim

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

var levelVar = &slog.LevelVar{}
var errUnrecognizedLevel = errors.New("unrecognized log level")

const bufInitialCapacity = 512

type LoggingOptions struct {
	Level      slog.Leveler
	PluginName string
}

type LoggingHandler struct {
	options LoggingOptions
	out     io.Writer
	ioMutex *sync.Mutex
}

type SetLogLevelParams struct {
	Level string `json:"level"`
}

func InitLogger(plugName string) {
	handlerOptions := LoggingOptions{
		Level:      levelVar,
		PluginName: plugName,
	}
	logger := slog.New(NewLoggingHandler(os.Stderr, &handlerOptions))
	slog.SetDefault(logger)
}

func HandleSetLogLevelCommand(cmd SimCommand) (string, error) {
	var level SetLogLevelParams
	err := json.Unmarshal(cmd.Parameters, &level)
	if err != nil {
		return "", err
	}

	prevLevel := levelVar.Level()

	switch level.Level {
	case "debug":
		levelVar.Set(slog.LevelDebug)
	case "info":
		levelVar.Set(slog.LevelInfo)
	case "warn":
		levelVar.Set(slog.LevelWarn)
	case "error":
		levelVar.Set(slog.LevelError)
	default:
		return "", fmt.Errorf("%w: %s", errUnrecognizedLevel, level.Level)
	}

	slog.Info("Set log level", "prev", prevLevel, "new", levelVar.Level())

	return string(cmd.Parameters), nil
}

func NewLoggingHandler(out io.Writer, options *LoggingOptions) *LoggingHandler {
	h := &LoggingHandler{out: out, ioMutex: &sync.Mutex{}}

	if options != nil {
		h.options = *options
	}
	if h.options.Level == nil {
		h.options.Level = slog.LevelInfo
	}

	return h
}

func (h *LoggingHandler) Enabled(_ context.Context, lvl slog.Level) bool {
	return lvl >= h.options.Level.Level()
}

func (h *LoggingHandler) Handle(_ context.Context, r slog.Record) error {
	buf := make([]byte, 0, bufInitialCapacity)

	if !r.Time.IsZero() {
		buf = fmt.Appendf(buf, "%s ", r.Time.Format(time.RFC3339Nano))
	}

	buf = fmt.Appendf(buf, "%s ", r.Level)

	if h.options.PluginName != "" {
		buf = fmt.Appendf(buf, "%s ", h.options.PluginName)
	}

	if r.PC != 0 {
		fs := runtime.CallersFrames([]uintptr{r.PC})
		f, _ := fs.Next()
		buf = fmt.Appendf(buf, "[%s:%d] ", shortenPath(f.File), f.Line)
	}

	buf = fmt.Appendf(buf, "%s", r.Message)

	r.Attrs(func(a slog.Attr) bool {
		buf = fmt.Appendf(buf, " %s=%v", a.Key, a.Value)

		return true
	})

	buf = append(buf, "\n"...)

	h.ioMutex.Lock()
	defer h.ioMutex.Unlock()
	_, err := h.out.Write(buf)

	return err
}

// NOTE: Not implemented!
func (h *LoggingHandler) WithAttrs(_ []slog.Attr) slog.Handler {
	return h
}

// NOTE: Not implemented!
func (h *LoggingHandler) WithGroup(_ string) slog.Handler {
	return h
}

func shortenPath(path string) string {
	before, after, found := strings.Cut(path, "simulators/")
	if found {
		return after
	}

	return before
}
