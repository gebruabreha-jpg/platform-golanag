#!/usr/bin/env python3
"""Generate configfile_fabric_capacity.yaml for PCPB-23628.

Service model (Table 1):
  40 Producers × 512 NIs × 20 routes/NI = ~400K routes
  Each DP consumer is paired 1-to-1 with a producer (receives only its producer's routes)
  16 IPSEC consumers receive all routes
"""
import yaml
import os

NUM_PRODUCERS = 40
NUM_NIS = 512
NUM_IPSEC_CONSUMERS = 16
ROUTES_PER_NI = 20
PRODUCER_KEEPALIVE = 15
CONSUMER_KEEPALIVE = 120
CONSUMER_HEARTBEAT = 5

NI_NAMES = [f"NI{i}" for i in range(1, NUM_NIS + 1)]


def generate_producers():
    producers = []
    for i in range(1, NUM_PRODUCERS + 1):
        producers.append({
            "name": f"producer-{i}",
            "routeHoldTime": 30,
            "keepaliveTimeout": PRODUCER_KEEPALIVE,
            "staleRoute": False,
            "routeTypes": [{"networkInstance": ni, "types": ["service_local"]}
                           for ni in NI_NAMES],
        })
    return producers


def generate_consumers():
    consumers = []
    # 40 DP consumers (paired 1-to-1 with producers)
    for i in range(1, NUM_PRODUCERS + 1):
        consumers.append({
            "name": f"consumer-dp-{i}",
            "keepaliveTimeout": CONSUMER_KEEPALIVE,
            "heartbeatInterval": CONSUMER_HEARTBEAT,
            "routeFilters": ["service_local"],
            "networkInstanceFilters": NI_NAMES,
        })
    # 16 IPSEC consumers (receive all routes)
    for i in range(1, NUM_IPSEC_CONSUMERS + 1):
        consumers.append({
            "name": f"consumer-ipsec-{i}",
            "keepaliveTimeout": CONSUMER_KEEPALIVE,
            "heartbeatInterval": CONSUMER_HEARTBEAT,
            "routeFilters": ["service_local"],
            "networkInstanceFilters": NI_NAMES,
        })
    return consumers


def generate_initial_routes():
    """Generate initial routes: 1 route per producer per NI for initial sync."""
    routes = []
    for i in range(1, NUM_PRODUCERS + 1):
        producer = f"producer-{i}"
        for ni_idx, ni in enumerate(NI_NAMES[:1], 1):  # Only first NI for initial config
            routes.append({
                "networkInstanceName": ni,
                "producerName": producer,
                "items": [{
                    "prefix": f"10.0.{i}.0/24",
                    "addressFamily": "ipv4",
                    "lsp": "non_lsp",
                    "routeType": "service_local",
                    "dist": 0,
                    "metric": 0,
                    "nexthops": [{
                        "addressInfo": {
                            "type": "service",
                            "serviceName": producer,
                        },
                        "attr": ["service_local"],
                    }],
                }],
                "sendEof": True,
                "eofAddressFamilies": ["ipv4"],
            })
    return routes


def generate_config():
    config = {
        "rib": {
            "parameter": {
                "producers": generate_producers(),
                "consumers": generate_consumers(),
                "routes": generate_initial_routes(),
                "rate": {
                    "producersPerSecond": 40,
                    "consumersPerSecond": 56,
                    "routesPerSecond": 5000,
                    "nexthopsPerSecond": 5000,
                },
            },
        },
        "template": {
            "sut_config": {
                "template": "config/sut_config.template",
                "output": "sut_configFile.json",
            },
        },
    }
    return config


if __name__ == "__main__":
    config = generate_config()
    output_path = os.path.join(os.path.dirname(os.path.abspath(__file__)),
                               "configfile_fabric_capacity.yaml")
    with open(output_path, "w") as f:
        yaml.dump(config, f, default_flow_style=False, sort_keys=False)
    print(f"Generated {output_path}")
    total = NUM_PRODUCERS * NUM_NIS * ROUTES_PER_NI
    print(f"  Producers: {NUM_PRODUCERS}")
    print(f"  NIs: {NUM_NIS}")
    print(f"  Routes/NI: {ROUTES_PER_NI}")
    print(f"  Total routes: {total}")
    print(f"  DP Consumers: {NUM_PRODUCERS} (paired)")
    print(f"  IPSEC Consumers: {NUM_IPSEC_CONSUMERS}")
