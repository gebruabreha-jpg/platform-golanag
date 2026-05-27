# CMM Simulator for RIB Service

## Overview

This is a minimal CMM (Configuration Management Mediator) simulator designed
specifically for the RIB service. It simulates only the essential interaction
that RIB has with CMM.

## Functionality

The simulator implements a single HTTP endpoint:

- **PUT** `/cm/api/v1/schemas/ietf-network-instance/data-sources/rib`
  - Accepts PUT requests from RIB service
  - Responds with `200 OK`
  - Tracks the number of requests received

## Configuration

No configuration is required for this simulator. The `Configure()` function is
not used.

## Statistics

The simulator tracks the following statistics accessible via `GetStats()`:

- `rxPutRequests`: Number of PUT requests received from RIB service

## Usage in Tests

```python
# RIB will automatically send PUT request to CMM during startup
# Verify it was received:
stats = framework.get_sim_stats("cmm")
assert stats["rxPutRequests"] > 0
