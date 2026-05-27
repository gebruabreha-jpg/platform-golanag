package vetosim

import (
	"encoding/json"
	"fmt"
	"log/slog"
)

type SimCommand struct {
	Command    string          `json:"command"`
	Parameters json.RawMessage `json:"parameters"`
}

func ExtractCommand(str string) (SimCommand, error) {
	var command SimCommand

	err := json.Unmarshal([]byte(str), &command)
	if err != nil {
		slog.Error("Error in Unmarshal when extracting command", "error", err, "raw", str)

		return SimCommand{}, fmt.Errorf("Command err: %w", err)
	}

	return command, nil
}

func TryHandlingCommand(cmd SimCommand) (bool, string, error) {
	var handlerRetVal string
	var handlerErr error

	switch cmd.Command {
	case "set-log-level":
		handlerRetVal, handlerErr = HandleSetLogLevelCommand(cmd)
	default:
		return false, "", nil
	}

	return true, handlerRetVal, handlerErr
}
