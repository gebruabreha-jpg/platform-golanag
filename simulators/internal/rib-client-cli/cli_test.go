package ribclientcli

import (
	"testing"

	"github.com/alecthomas/kong"
	"github.com/stretchr/testify/assert"
)

// Verifies that PrintUsage does not return an error for a valid parser and context.
func TestPrintUsage_NoError(t *testing.T) {
	var commands Commands
	parser := kong.Must(&commands)
	ctx, err := parser.Parse([]string{"help"})
	assert.NoError(t, err)

	err = PrintUsage(parser, false, ctx)
	assert.NoError(t, err)
}

// Verifies that PrintUsage with allCommands=true does not return an error.
func TestPrintUsage_AllCommands(t *testing.T) {
	var commands Commands
	parser := kong.Must(&commands)
	ctx, err := parser.Parse([]string{"help"})
	assert.NoError(t, err)

	err = PrintUsage(parser, true, ctx)
	assert.NoError(t, err)
}

// Verifies that Config has the correct default value for Rib.
func TestConfig_DefaultRibServer(t *testing.T) {
	var config Config
	parser := kong.Must(&config)
	_, err := parser.Parse([]string{})
	assert.NoError(t, err)
	assert.Equal(t, "eric-pc-routing-information-base", config.RibServer)
}

// Verifies that Config accepts a custom Rib server argument.
func TestConfig_CustomRibServer(t *testing.T) {
	var config Config
	parser := kong.Must(&config)
	_, err := parser.Parse([]string{"my-custom-rib"})
	assert.NoError(t, err)
	assert.Equal(t, "my-custom-rib", config.RibServer)
}
