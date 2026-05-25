# TRex Implementation Summary

## Current Dallas Implementation

### Dallas Route Workflow
- `BeetsDallasSetupHandler.java`
- `BeetsSuiteFixture.java`
- `BeetsDallasPostInstall.java`

## Dallas Routes for ADP CM and YANG Provider Interaction

**Primary Route:** `configure_ip_address_on_tap.sh.template`

### ADP Service Controller Routes
- **pc_sm_controller_ecfe routes** (IPv4/IPv6) - Service Controller ECFE interface
- **pc_sm_controller routes** (IPv4/IPv6) - Main Service Controller interface

### Routing Entries

**For ADP Service Controller ECFE (Signaling interface):**
```
ip route add {{ adp.stp.service.pc_sm_controller_ecfe.ipv4_supernet }} via {{ bgw.stp.link.signaling_if1.ipv4.ip }} dev vlan{{ dallas.stp.link.signaling_if1.vlan }}
```

**For ADP Service Controller (OM/CN interface):**
```
ip route add {{ adp.stp.service.pc_sm_controller.ipv4_supernet }} via {{ bgw.stp.link.om_cn_if1.ipv4.ip }} dev vlan{{ dallas.stp.link.om_cn_if1.vlan }}
```

## TRex Java Setup

### Implementation Status
- ✅ `prepareTRex()` method implemented
- ✅ `uploadTrex()` and `startTRex()` methods working
- ✅ Logs confirm TRex uploads and starts correctly

### Routing Logic
- `prepareTRex()` calls `dallasSetupHandler.createDallasRoutesSystemService()`
- Same route template used for both Dallas and TRex
- Same network interfaces (tool server interfaces, not traffic generator specific)

### Current Implementation Flow
```java
prepareTRex() {
    // 1. Uninstall Dallas traffic generator
    dallasMaster.getDallasControl().uninstall();
    
    // 2. Keep Dallas routes (reuse routing logic)
    dallasSetupHandler.createDallasRoutesSystemService();
    
    // 3. Setup TRex traffic generator
    // (via framework step definitions)
}
```

## Template Selection Mechanism

### How Templates Are Chosen
The configuration system automatically selects templates based on the current context:

1. **Dallas context:** Uses `templates/dallas/dallas_env/configure_ip_address_on_tap.sh.template`
2. **TRex context:** Uses `templates/trex/configure_ip_address_on_tap.sh.template`

### Template Mapping (trex_sub_config.yaml)
```yaml
template:
  trex:
    ip_routing:
      template: templates/trex/configure_ip_address_on_tap.sh.template
      output: configure_ip_address_on_tap.sh
```

## Dallas Variables in TRex Templates

### Source of Variables
The `dallas.stp.link.*` variables come from STP (System Test Platform) configuration, NOT from Dallas software installation.

### Configuration Context Flow
1. When TRex runs, `beetsConfigurationProvider.get().getConfiguration()` returns TRex configuration context
2. `configuration.getFile("configure_ip_address_on_tap.sh")` resolves to the TRex template
3. Template variables resolve to STP network configuration values

### Why This Works
- Network infrastructure stays the same (VLANs, IPs, interfaces)
- Only traffic generator changes (Dallas → TRex)
- STP provides network configuration used by both tools

## STP Configuration Files

| File | Path | Purpose |
|------|------|---------|
| catapult_keys.yaml | `subprofiles/confix/catapult_keys.yaml` | Contains dallas.stp.link.* definitions |
| geored_2_clusters.yaml | `subprofiles/geored/geored_2_clusters.yaml` | Additional STP network topology for geo-redundant setup |

### Dallas STP Configuration Structure
```yaml
dallas:
  stp:
    default_ci: pccc-dallas
    link:
      media_if1: {name: if-media-tools-1}
      signaling_if1: {name: if-signaling-tools-1}
      om_cn_if1: {name: if-om_cn-tools-1}
      ran_if1: {name: if-ran-tools-1}
      sgi1_if1: {name: if-sgi1-tools-1}
      # ... additional interfaces
```

These provide variables like `{{ dallas.stp.link.sgi3_if1.vlan }}` used by the TRex template.