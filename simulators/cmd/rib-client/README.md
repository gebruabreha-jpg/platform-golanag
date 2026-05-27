# RIB Client Simulator

## About Simulator

This simulator emulates clients of the Routing Information Base (RIB). These
clients interact with the RIB through its API, acting as either producers
(injecting routes, nexthops, and prefixes) or consumers (subscribing to routing
information). The simulator supports both v1 and v2 API versions.

## Functionality

### Initial Configuration

The initial configuration for every test case is performed through the
`Configure` function. It receives a JSON payload defining the full test
environment setup:

- **Producer parameters:** identifier, route hold time, keepalive timeout,
  heartbeat interval, stale route flag, callback URL, and a list of route types
  (connected, service, BGP, etc.) per network instance.
- **Consumer parameters:** identifier, keepalive timeout, heartbeat interval,
  route filters, network instance filters, and callback URL.
- **Network instances:** definitions and their associated address families,
  indicating which route types they carry.
- **Global rate parameters:** producers/consumers registered per second, routes
  injected per second, and nexthops registered per second.

Based on the declared network instances and route types, the simulator
automatically generates the corresponding nexthops, prefixes, and routes.

Upon receiving the configuration, the simulator concurrently executes the full
registration sequence against the SUT — producer/consumer registration,
nexthop/prefix registration with their respective EOFs, route injection with
route EOF, and heartbeat initialization — before the test scenario begins.

### Commands

Commands are dispatched through the `Command(string)` function, which expects a
JSON string with the following structure:

```json
{
  "command": "<command-name>",
  "parameters": { ... }
}
```

The function first parses the JSON into a `SimCommand` struct via
`ExtractCommand`. It then passes the command through two handler layers in
order:

1. **Shared handler** (`TryHandlingCommand`) — resolves commands common to all
   simulators. Currently handles `set-log-level`.
2. **RIB client handler** (`TryHandlingRibClientCommand`) — resolves
   RIB-specific commands by mapping the command name to its corresponding
   handler and invoking `RunJSON` with the raw JSON parameters.

If neither layer recognises the command, the `Command` function returns an
`errCmdNotImplemented` error. On success, the raw parameters JSON is returned
as the result string.

## Stats

`GetStats` is part of the `Sim` interface. Currently, the simulator does
not collect internal metrics, it always returns an empty map with no error.

## Events

`RecvEvent` is part of the `Sim` interface. For this simulator, the function is
a no-op and always returns `nil`.
