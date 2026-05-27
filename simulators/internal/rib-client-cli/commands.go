package ribclientcli

// Commands defines the top-level commands available in the CLI REPL.
type Commands struct {
	Help HelpCmd `cmd:"" aliases:"h" help:"Show help."`
	Exit ExitCmd `cmd:"" aliases:"x" help:"Exit the simulator."`
}

// HelpCmd triggers the help display.
type HelpCmd struct{}

// ExitCmd exits the simulator.
type ExitCmd struct{}
