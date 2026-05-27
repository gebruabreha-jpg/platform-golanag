package ribclientcli

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime/debug"
	"strings"

	"gerrit-gamma.gic.ericsson.se/pc/eric-pc-routing-information-base/tests/veto/simulators/internal/logger"

	"github.com/alecthomas/kong"
)

const logFilePermissions = 0o666

// PrintUsage displays the CLI usage information.
func PrintUsage(parser *kong.Kong, allCommands bool, ctx *kong.Context) error {
	ctx, err := kong.Trace(parser, ctx.Args)
	if err != nil {
		return err
	}
	err = kong.DefaultHelpPrinter(kong.HelpOptions{
		Compact:             true,
		Summary:             false,
		Tree:                allCommands,
		NoExpandSubcommands: !allCommands,
		FlagsLast:           true,
	}, ctx)

	return err
}

// Run initializes logging, parses configuration, and starts the CLI REPL loop.
func Run() {
	defaultLogFile, err := os.OpenFile(
		logger.LogsFolder+"rib-client-simulator.log",
		os.O_RDWR|os.O_CREATE|os.O_APPEND,
		logFilePermissions)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer defaultLogFile.Close()
	log.SetOutput(defaultLogFile)

	defer func() {
		if r := recover(); r != nil {
			log.Printf("Panic: %v\nStack trace:\n%s", r, debug.Stack())
		}
	}()

	var config Config
	kongCtx := kong.Parse(&config,
		kong.Name("rib-client-simulator"),
		kong.Description("A simulator for producers and consumers in the RIB."),
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))
	fmt.Println("RIB server:", config.RibServer)

	cliCommands := Commands{}
	parser := kong.Must(&cliCommands,
		kong.UsageOnError(),
		kong.ConfigureHelp(kong.HelpOptions{
			Compact: true,
			Summary: true,
		}))
	if err := PrintUsage(parser, false, kongCtx); err != nil {
		log.Printf("Error printing usage: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		ctxKong, err := parser.Parse(args)
		if err != nil {
			fmt.Printf("Error parsing command: %v\n", err)

			continue
		}

		switch ctxKong.Command() {
		case "exit":
			return

		case "help":
			if err := PrintUsage(parser, true, kongCtx); err != nil {
				fmt.Printf("Error printing usage: %v\n", err)
			}

		default:
			if err := ctxKong.Run(); err != nil {
				fmt.Printf("Error running command: %v\n", err)

				continue
			}
		}
	}
}
