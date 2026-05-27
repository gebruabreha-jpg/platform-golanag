package vetosim

import (
	"net/rpc"

	"github.com/hashicorp/go-plugin"
)

// This is the RPC arguments for the SendEvent function.
type EvArgs struct {
	Name  string
	Event string
}

// Sim is the interface that we're exposing as a plugin.
type Sim interface {
	Start() error
	GetStats() (map[string]uint64, error)
	Configure(string) error
	RecvEvent(string) error
	Command(string) (string, error)
}

// Here is an implementation that talks over RPC.
type SimRPC struct{ client *rpc.Client }

func (g *SimRPC) Start() error {
	var resp interface{}
	err := g.client.Call("Plugin.Start", new(interface{}), &resp)

	return err
}

func (g *SimRPC) GetStats() (map[string]uint64, error) {
	var resp map[string]uint64
	err := g.client.Call("Plugin.GetStats", new(interface{}), &resp)

	return resp, err
}

func (g *SimRPC) Configure(cfg string) error {
	// We don't expect a response, so we can just use interface{}
	var resp interface{}
	err := g.client.Call("Plugin.Configure", cfg, &resp)

	return err
}

func (g *SimRPC) RecvEvent(ev string) error {
	// We don't expect a response, so we can just use interface{}
	var resp interface{}
	err := g.client.Call("Plugin.RecvEvent", ev, &resp)

	return err
}

func (g *SimRPC) Command(ev string) (string, error) {
	var resp string
	err := g.client.Call("Plugin.Command", ev, &resp)

	return resp, err
}

// Here is the RPC server that SimRPC talks to, conforming to
// the requirements of net/rpc.
type SimRPCServer struct {
	// This is the real implementation
	Impl Sim
}

func (s *SimRPCServer) Start(_ interface{}, _ *interface{}) error {
	return s.Impl.Start()
}

func (s *SimRPCServer) GetStats(_ interface{}, resp *map[string]uint64) error {
	v, err := s.Impl.GetStats()
	*resp = v

	return err
}

func (s *SimRPCServer) Configure(cfg string, _ *interface{}) error {
	return s.Impl.Configure(cfg)
}

func (s *SimRPCServer) RecvEvent(ev string, _ *interface{}) error {
	return s.Impl.RecvEvent(ev)
}

func (s *SimRPCServer) Command(cmd string, resp *string) error {
	v, err := s.Impl.Command(cmd)
	*resp = v

	return err
}

// This is the implementation of plugin.Plugin so we can serve/consume this
//
// This has two methods: Server must return an RPC server for this plugin
// type. We construct a SimRPCServer for this.
//
// Client must return an implementation of our interface that communicates
// over an RPC client. We return SimRPC for this.
//
// Ignore MuxBroker. That is used to create more multiplexed streams on our
// plugin connection and is a more advanced use case.
type SimPlugin struct {
	// Impl Injection
	Impl Sim
}

func (p *SimPlugin) Server(*plugin.MuxBroker) (interface{}, error) {
	return &SimRPCServer{Impl: p.Impl}, nil
}

func (SimPlugin) Client(_ *plugin.MuxBroker, c *rpc.Client) (interface{}, error) {
	return &SimRPC{client: c}, nil
}

// HandshakeConfig is used to just do a basic handshake between
// a plugin and host. If the handshake fails, a user friendly error is shown.
// This prevents users from executing bad plugins or executing a plugin
// directory. It is a UX feature, not a security feature.
var HandshakeConfig = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SIM_PLUGIN",
	MagicCookieValue: "hello",
}
