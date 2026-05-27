package logger

import (
	"fmt"
	"log"
	"os"
)

const (
	LogsFolder      = "./"
	logExtension    = ".log"
	filePermissions = 0o666
)

type Logger struct {
	*log.Logger
	file *os.File
}

func (l *Logger) Close() error {
	if l.file != nil {
		return l.file.Close()
	}

	return nil
}

func CreateLogger(clientType string, version int, name string) (*Logger, error) {
	fileName := LogsFolder + name + logExtension
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, filePermissions)
	if err != nil {
		return nil, fmt.Errorf("failed to create logger for %s: %w", name, err)
	}
	prefix := fmt.Sprintf("[%s:V%d:%s] ", clientType, version, name)

	return &Logger{
		Logger: log.New(file, prefix, log.Ldate|log.Ltime|log.Lshortfile),
		file:   file,
	}, nil
}
