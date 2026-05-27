package ribclientcli

// Config holds the CLI configuration parsed from command-line arguments.
type Config struct {
	RibServer string `arg:"" default:"eric-pc-routing-information-base" help:"The service name of the RIB."`
}
