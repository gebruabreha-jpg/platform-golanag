package logger

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Verifies that CreateLogger returns a non-nil logger.
func TestCreateLogger_ReturnsNonNil(t *testing.T) {
	l, err := CreateLogger("Test", 1, "test-logger")
	defer os.Remove("test-logger.log")

	assert.NoError(t, err)
	assert.NotNil(t, l)
}

// Verifies that CreateLogger creates the log file on disk.
func TestCreateLogger_CreatesLogFile(t *testing.T) {
	name := "test-create-file"
	_, err := CreateLogger("Test", 1, name)
	defer os.Remove(name + logExtension)

	assert.NoError(t, err)
	assert.FileExists(t, name+logExtension)
}

// Verifies that CreateLogger sets the correct prefix on the logger.
func TestCreateLogger_SetsPrefix(t *testing.T) {
	name := "test-prefix"
	l, err := CreateLogger("Producer", 2, name)
	defer os.Remove(name + logExtension)

	assert.NoError(t, err)
	assert.NotNil(t, l)
	assert.Contains(t, l.Prefix(), "[Producer:V2:test-prefix]")
}

// Verifies that CreateLogger writes log output to the file.
func TestCreateLogger_WritesToFile(t *testing.T) {
	name := "test-write"
	l, err := CreateLogger("Test", 1, name)
	defer os.Remove(name + logExtension)

	assert.NoError(t, err)
	l.Println("hello")

	data, err := os.ReadFile(name + logExtension)
	assert.NoError(t, err)
	assert.Contains(t, string(data), "hello")
}

// Verifies that CreateLogger returns an error when the file path is invalid.
func TestCreateLogger_ReturnsErrorOnInvalidPath(t *testing.T) {
	l, err := CreateLogger("Test", 1, string([]byte{0}))
	assert.Error(t, err)
	assert.Nil(t, l)
}

// Verifies that Close closes the underlying file.
func TestLogger_Close(t *testing.T) {
	name := "test-close"
	l, err := CreateLogger("Test", 1, name)
	defer os.Remove(name + logExtension)

	assert.NoError(t, err)
	assert.NotNil(t, l)
	assert.NoError(t, l.Close())
}

// Verifies that Close on a nil file returns nil.
func TestLogger_Close_NilFile(t *testing.T) {
	l := &Logger{}
	assert.NoError(t, l.Close())
}
