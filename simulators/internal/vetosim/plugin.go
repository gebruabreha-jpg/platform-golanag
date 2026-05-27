package vetosim

import (
	plugin "github.com/hashicorp/go-plugin"
)

func ServePlugin(impl Sim) {
	// pluginMap is the map of plugins we can dispense.
	var pluginMap = map[string]plugin.Plugin{
		"sim": &SimPlugin{Impl: impl},
	}

	plugin.Serve(&plugin.ServeConfig{
		HandshakeConfig: HandshakeConfig,
		Plugins:         pluginMap,
	})
}
