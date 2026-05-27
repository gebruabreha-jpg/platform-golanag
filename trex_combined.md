# TRex Dallas Network Issue - Complete Summary
## Problem
When TRex starts, it binds NIC1/NIC2 to DPDK, removing them from the kernel. This destroys all VLANs built on those NICs, causing:
- All kernel routes (signaling, media, om_cn, ran, sgi) to disappear
- iptables and NAT rules to become useless (no route to destination)  
- Yang/Netconf, Outline, Prometheus, and Search Engine to become unreachable
- Only raw kubectl commands (via Kubernetes API) continue to work

Current Dallas Implementation on Network Stack:-
Physical Interfaces:-INTERFACE_SWITCH1 (enp59s0f0) and INTERFACE_SWITCH2 (enp59s0f1) — used by both kernel VLANs.

Network Setup in configure_ip_address_on_tap.sh:-
VLANs built on both physical NICs:
Media:     vlan<media_if1/2>
Signaling: vlan<signaling_if1/2>
O&M:       vlan<om_cn_if1/2>
RAN:       vlan<ran_if1/2>
SGi:       vlan<sgi1_if1> through vlan<sgi5_if2>

Routes on those VLANs:-
AMF/UDM/NRF/SMF/PCF → signaling VLANs
PCG UP N4/N9/N3 → signaling/media/RAN VLANs
ADP COMMON, ECFE → signaling/om_cn VLANs
Outline IPs (100.1.x.x, 100.2.x.x, 100.3.x.x) bound to media/signaling/om_cn VLANs
iptables NAT for Prometheus/Search Engine/CNOM access via signaling VLAN IPs

Services Requiring Dallas Routes/Needs Dallas Routes (reliant on VLANs created on physical NICs):-
    Yang Provider (AdpCm/Netconf)
    Outline (Legacy) 
    PM/Prometheus (Ingress)
    Search Engine
    Object Storage (SFTP)
Does NOT Need Dallas Routes (uses Kubernetes API):
PM/Prometheus (kubectl port-forward) -kubectl port-forward goes through Kubernetes API server, not dallas VLAN routes as an alternative access method specifically for LLV test steps that need direct pod access. However, this is NOT the primary communication method - it's only used for specific test scenarios. The standard/primary access for PM/Prometheus uses ingress with loadBalancerIP, which DOES require Dallas routes and becomes unreachable when TRex is running.

## Services Requiring Dallas Routes (reliant on VLANs created on physical NICs)
Java File Paths where they are implemented when we use dallas:-

Yang Provider (AdpCm/Netconf):-
File: src/main/java/com/ericsson/pc/beets/testcases/steps/DeploymentSteps.java
File:`DeploymentSteps.java`-gets yang-provider external service IP, creates SSHtunnel
Method: SSH tunnel via DallasPortForwarder to yang-provider external service IP
File:`AdpOamStepDefinition.java:567-568` - creates SSH tunnel via DallasPortForwarder
File:`BeetsGuiceModule.java:92`- DallasPortForwarder injection
Method:SSH tunnel through DallasPortForwarder to reach yang-provider external service I


Yang Provider (AdpOam)
File: src/main/java/com/ericsson/pc/beets/testcases/steps/AdpOamStepDefinition.java
Method: SSH tunnel via DallasPortForwarder using adpCmProperties.getHost() and port


PM/Prometheus
File: src/main/java/com/ericsson/pc/beets/testcases/steps/AdpOamLlvStepDefinition.java
Method: kubectl port-forward -n beets pods/eric-pm-server-0 9091:9090

PM/Prometheus (Ingress)
File: ActionsHelper.java
File: `ActionsHelper.java:1020-1035` - patches ingress-nginx loadBalancerIP, creates ingress, Method: Ingress with loadBalancerIP via `kubectlApi.patchService()` and `kubectlApi.createHttpIngress()`

Search Engine
File: ActionsHelper.java
File: `ActionsHelper.java:1061-1062` - creates ingress, Method: Ingress with loadBalancerIP via `kubectlApi.createHttpIngress()` accessed through `searchEngineProperties.getServer()`

Outline (Legacy)
File: Runs on dallas tool server (no specific Java file path documented)
Method: Binds IPs (100.1.x.x, 100.2.x.x, 100.3.x.x) to VLAN interfaces. SUT sends traffic (charging records, IPFIX, syslog) to these IPs. OutlineServerNavigator SSHes to tool server to read files.

Object Storage (SFTP)
File: src/main/java/com/ericsson/pc/beets/testcases/steps/DeploymentSteps.java
Method: SSH tunnel via DallasPortForwarder to adpObjectStorageProperties.getHost()

DallasPortForwarder Injection
File: src/main/java/com/ericsson/pc/beets/fw/BeetsGuiceModule.java
Method: @Named("DallasPortForwarder") final Provider<PortForwarder> portForwarderProvider - provides the SSH tunnel capability

PM/Prometheus (kubectl port-forward)**
- File: `AdpOamLlvStepDefinition.java:1166`
- Method: `kubectl port-forward -n beets pods/eric-pm-server-0 9091:9090`
- Note: Alternative access for specific LLV test steps, NOT primary communication method

## Solution Proposal
1. Use cOutline instead of legacy Outline - eliminates dependency on TRex-bound VLANs
2. Use a management interface (third NIC) that TRex doesn't touch - configure routes to kubernetes service IPs through the management NIC instead of VLANs on TRex-bound NICs
3. For Yang/Netconf - add routes via management interface to reach yang-provider external service IP and tm_ingress_controller_cr_ecfe IP