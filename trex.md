  1. private final DallasProperties dallasProperties; — stores the reference so you can use it later (it was previously only used in the constructor and discarded)
  2. this.dallasProperties = dallasProperties; — saves the constructor parameter into the field
  3. reCreateDallasRoutes(); in startTRex() — after TRex starts and takes over the NICs (binds them to DPDK), the TAP interface routes are lost, so you re-create
  them
  4. reCreateDallasRoutes() method — runs configure_ip_address_on_tap.sh on the tool server to restore the Dallas routes on TAP interfaces that TRex destroyed when
  it started
  The logic is sound: TRex start → NICs bound to DPDK → TAP routes gone → re-create them. This is the correct place to do it (after TRex is confirmed running).
  And as we established, this code cannot cause the prepareTRex() failure since it only executes inside startTRex(), which runs much later. The failure you're
  seeing is a pre-existing race condition in BeetsSuiteFixture.prepareTRex() where getDstMac() is called immediately after the script reports "Not all interfaces
  are up".

  The core problem:-
  When TRex starts → NICs unbound from kernel → all VLANs built on those NICs are destroyed → all routes, outline IPs, everything is gone. The script can't
  recreate them because the physical interfaces no longer exist in the kernel. TAP interfaces created by TRex are a different thing — they only carry TRex traffic,
  not the full VLAN/routing stack.

  Proposed solution:-
  You need a separate management interface that is NOT used by TRex. The tool server should have:
  - NIC1/NIC2 → dedicated to TRex (unbound from kernel, bound to DPDK)
  - A third interface (or the management interface already present) → keeps the VLAN/route infrastructure alive
  OR alternatively:
  Split the script into two parts:
  1. Part A — routes that go through TRex NICs (media/ran traffic) → these are handled by TRex itself via its traffic profiles, no need for kernel routes
  2. Part B — routes needed for framework operations (PM, yang, outline, signaling to microservices) → these should use a different interface that TRex doesn't
  touch,Does the tool server have a third NIC or management interface that stays in the kernel while TRex takes over NIC1/NIC2?


  How each service connects — based on actual code:
  1. PM Server:-
  Connection method: kubectl port-forward from the local machine (where tests run) to the kubernetes cluster.
  kubectl port-forward -n beets pods/eric-pm-server-0 9091:9090
  Needs dallas routes? NO — uses kubernetes API networking, not dallas VLAN routes.
  2. Yang Provider (AdpCm / Netconf)
  Connection method: SSH port-forwarding through the DallasPortForwarder. The flow is:
  Local machine → SSH tunnel via DallasPortForwarder → adpCmProperties.getHost():adpCmProperties.getPort()
  The DallasPortForwarder creates an SSH tunnel through the dallas tool server to reach adpCmProperties.getHost() (which is the eric-cm-yang-provider service IP
  inside the cluster).
  Needs dallas routes? YES — the dallas tool server needs routes to reach the kubernetes service IP (adpCmProperties.getHost()). The SSH tunnel goes through the
  dallas tool server, so the tool server must be able to route to the node's service IPs. These routes are in the script under "ADP COMMON" and "SIGNALING"
  sections.
  3. Outline (Legacy):-
  Connection method: Outline runs on the dallas tool server itself. It binds IPs (100.1.x.x, 100.2.x.x, 100.3.x.x) to VLAN interfaces on the tool server. The SUT
  (PCG/PCC) sends traffic to these IPs. The OutlineServerNavigator is an SSH connection to the tool server to read files/run commands.
  Needs dallas routes? YES — Outline needs:
  - VLAN interfaces to exist (created on INTERFACE1/INTERFACE2)
  - IPs bound to those VLANs
  - The SUT needs routes to reach those IPs (configured on the node side)
  4. Other microservices (kubectl operations):-
  Connection method: kubectlApi.exec(), kubectlApi.get(), kubectlApi.logs() — all go through the kubernetes API server.
  Needs dallas routes? NO — uses kubernetes API, not dallas network.


The real problem for your TRex scenario:
  When TRex unbinds NICs from kernel → VLANs on those NICs are destroyed → Yang/Netconf loses connectivity (can't route to service IPs) and legacy Outline stops
  working (its IPs disappear).
  What you need to keep working after TRex starts:
  1. Yang/Netconf — needs a route from the tool server to the cluster service IPs
  2. Legacy Outline — needs VLAN interfaces with IPs bound
  Both depend on VLANs that are built on the physical NICs that TRex takes over.
  Possible solutions:
  1. Use cOutline instead of legacy Outline (eliminates that dependency)
  2. For Yang/Netconf: route the service IPs through a different interface (management NIC) that TRex doesn't touch
  3. For Yang/Netconf: use kubectl port-forward instead of SSH tunneling through dallas (like PM server does)

 Actual connection methods:

  1. Yang Provider / AdpCm (Netconf)

  Connection: localPortRegistryService.addLocalPortForwarder(portForwarder, host, port, 3) → creates an SSH tunnel through the DallasPortForwarder to reach
  adpCmProperties.getHost().

  The host is determined by:

  kubectl.kubectl("get svc -n %s | grep eric-cm-yang-provider-external| awk '{print $4}'")

  This gets the EXTERNAL-IP of the eric-cm-yang-provider-external service (a LoadBalancer or NodePort service). The connection path is:

  Test runner → SSH tunnel via DallasPortForwarder (dallas tool server) → yang-provider external service IP

  Needs dallas routes? YES — the dallas tool server must be able to reach the yang-provider external service IP.

  2. PM Server

  Connection: kubectl port-forward -n beets pods/eric-pm-server-0 9091:9090

  This is a kubectl command run from the test runner machine (or dallas tool server). It uses the kubernetes API (via kubeconfig) to set up a port-forward tunnel.

  Needs dallas routes? NO — kubectl port-forward goes through the kubernetes API server, not through the data plane network.

  3. kubectl operations (get pods, exec, logs, etc.)

  Connection: Uses kubectlProperties.getAdminConfPath() — a kubeconfig file that points to the kubernetes API server.

  Needs dallas routes? NO — goes through kubernetes API server.

  4. Outline (legacy)

  Connection: Runs on the dallas tool server. Binds IPs to VLAN interfaces. The SUT sends traffic TO outline (charging records, IPFIX, syslog). The
  OutlineServerNavigator SSHes to the tool server to read files.

  Needs dallas routes? YES — needs VLANs with IPs for the SUT to send data to.





  1. Yang/Netconf (AdpCm) — SSH tunnel through dallas tool server

  File: src/main/java/com/ericsson/pc/beets/testcases/steps/DeploymentSteps.java

  - Line 1343: Gets yang-provider external service IP:

    final String hostIp = kubectl.kubectl(yangProviderHostString).trim();

  - Line 1376: Creates SSH tunnel through DallasPortForwarder to that IP:

    final int localSshPort = localPortRegistryService.addLocalPortForwarder(portForwarder.get(), hostIp, port, 3);

  File: src/main/java/com/ericsson/pc/beets/testcases/steps/AdpOamStepDefinition.java

  - Line 567-568: Same pattern — SSH tunnel through DallasPortForwarder:

    final int localSshPort = localPortRegistryService.addLocalPortForwarder(portForwarder,
            adpCmProperties.getHost(), adpCmProperties.getPort(), 3);

  File: src/main/java/com/ericsson/pc/beets/fw/BeetsGuiceModule.java

  - Line 92: DallasPortForwarder injected into AdpCmProviderClusterMap:

    @Named("DallasPortForwarder") final Provider<PortForwarder> portForwarderProvider

  ─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

  2. PM Server — kubectl port-forward

  File: src/main/java/com/ericsson/pc/beets/testcases/steps/AdpOamLlvStepDefinition.java

  - Line 1155-1166: `setPort



PM Server

  File: AdpOamLlvStepDefinition.java line 1166

  kubectlApi.kubectl("port-forward -n " + ns + " pods/eric-pm-server-0 9091:9090");

  Uses kubectl port-forward — goes through kubernetes API. No dallas routes needed.

  Prometheus

  File: ActionsHelper.java lines 1020-1035

  kubectlApi.patchService(ingressNginx, "{\"spec\":{\"loadBalancerIP\":\"" + prometheusProperties.getServer() + "\"}}");
  kubectlApi.createHttpIngress(kubeNamespace, prometheusProperties.getServiceName(),
          prometheusProperties.getServicePort(), "/prometheus");

  Uses ingress with a loadBalancerIP — patches the ingress-nginx service with a specific IP, then creates an ingress. The framework accesses prometheus via that
  loadBalancerIP. Needs a route from the test runner to that IP.

  Search Engine

  File: ActionsHelper.java lines 1061-1062

  kubectlApi.createHttpIngress(kubeNamespace, searchEngineProperties.getServiceName(),
          searchEngineProperties.getServicePort(), searchEngineProperties.getPathPrefix());

  Same pattern — uses ingress. Accessed via searchEngineProperties.getServer(). Needs a route to the ingress IP.

  Yang/Netconf (AdpCm)

  File: DeploymentSteps.java line 1376, AdpOamStepDefinition.java line 567

  localPortRegistryService.addLocalPortForwarder(portForwarder, adpCmProperties.getHost(), port, 3);

  Uses SSH tunnel through DallasPortForwarder to reach the yang-provider external service IP. Needs the dallas tool server to route to that IP.

  Object Storage (SFTP)

  File: DeploymentSteps.java line 712

  localPortRegistryService.addLocalPortForwarder(portForwarder.get(), adpObjectStorageProperties.getHost(), sutTargetPort, nRetries);

  Same — SSH tunnel through DallasPortForwarder. Needs dallas tool server to route to that IP.



Yes, you're right — they are the same thing. eric-pm-server IS Prometheus. It's Ericsson's packaged version of Prometheus.

  The code shows two different ways to access the same eric-pm-server:

  1. kubectl port-forward (in AdpOamLlvStepDefinition.java line 1166) — used for specific LLV test steps that need direct pod access
  2. Ingress (in ActionsHelper.java line 1030) — the standard setup that patches ingress-nginx with a loadBalancerIP and creates an HTTP ingress route

  So the normal/primary access method for PM/Prometheus is via ingress (loadBalancerIP) — which needs a route from wherever the test runner queries it. The kubectl
  port-forward is just an alternative used in specific test steps.

  Corrected: PM Server / Prometheus → accessed via ingress (loadBalancerIP) → YES, needs a route to the ingress IP.







PM/Prometheus (ingress) → unreachable
  - Search Engine (ingress) → unreachable
  - Yang/Netconf (via dallas tool server) → unreachable
  - Object Storage (via dallas tool server) → unreachable
  - Outline (legacy) → unreachable

  The only things that work while TRex is running are raw kubectl commands (via kubeconfig/kubernetes API).

  Possible solutions:

  1. Use a third NIC on the tool server (management interface) that TRex doesn't touch — create routes on that interface instead
  2. Move all service access to kubectl-based methods (port-forward, exec) instead of ingress/SSH-tunnel — but this requires significant framework changes
  3. Don't query PM/Yang/etc during TRex traffic — only query before TRex starts and after TRex stops (rebind NICs → recreate routes → query)
  4. Stop TRex temporarily when you need to query services, then restart it




















 Network setup created by configure_ip_address_on_tap.sh:

  Interfaces used:

  From /etc/testtool.interfaces.conf:

  - INTERFACE_SWITCH1 (e.g. enp59s0f0) — physical NIC 1
  - INTERFACE_SWITCH2 (e.g. enp59s0f1) — physical NIC 2
  - NIC1_NAME / NIC2_NAME — same physical NICs (used by TRex via PCI)
  - NIC1_NAME_TAP / NIC2_NAME_TAP — TAP interfaces created by TRex

  What's built on those interfaces:

  1. VLANs (on INTERFACE1/INTERFACE2):

  - vlan<media_if1>, vlan<media_if2> — media network
  - vlan<signaling_if1>, vlan<signaling_if2> — signaling network
  - vlan<om_cn_if1>, vlan<om_cn_if2> — O&M network
  - vlan<ran_if1>, vlan<ran_if2> — RAN network
  - vlan<sgi1_if1> ... vlan<sgi5_if2> — SGi networks

  2. Routes (on those VLANs):

  - AMF, UDM, NRF, SMF, PCF, VSMF → via signaling VLANs
  - PCG UP signaling (N4) → via signaling VLANs
  - PCG UP media (N9) → via media VLANs
  - PCG UP RAN (N3) → via RAN VLANs
  - PCG O&M → via om_cn VLANs
  - ADP COMMON (yang-provider, pm-controller) → via signaling + om_cn VLANs
  - ECFE → via signaling VLANs
  - Outline → via media/signaling/om_cn VLANs

  3. iptables NAT (for Prometheus/Search Engine/CNOM access):

  tool-server-hostname:80  → DNAT → tm_ingress_controller_cr_ecfe IP:80
  tool-server-hostname:443 → DNAT → tm_ingress_controller_cr_ecfe IP:443

  With SNAT source = dallas.stp.link.signaling_if1.ipv4_range[0].ip (an IP on the signaling VLAN)

  4. Outline IPs (bound to VLANs):

  - 100.1.x.x on media VLAN
  - 100.2.x.x on signaling VLAN
  - 100.3.x.x on om_cn VLAN

  ─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

  The key insight:

  Everything — Prometheus, Search Engine, Yang/Netconf, Outline — routes through VLANs that are built on INTERFACE_SWITCH1/INTERFACE_SWITCH2. These are the same
  physical NICs that TRex takes over.

  The iptables NAT for Prometheus/Search Engine uses:

  - Destination: tm_ingress_controller_cr_ecfe IP (reachable via signaling VLAN)
  - Source NAT: dallas.stp.link.signaling_if1 IP (an IP on the signaling VLAN)

  So when TRex unbinds the NICs → VLANs gone → iptables rules become useless (no route to destination) → Prometheus, Search Engine, Yang, Outline all unreachable.

  ─────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────────

  The real question:

  The tm_ingress_controller_cr_ecfe IP and adpCmProperties.getHost() (yang-provider external IP) — are these IPs only reachable through the signaling/om_cn VLANs
  on the tool server? Or is there another path to reach them (e.g., from the management network of the tool server)?

  If the kubernetes cluster's worker nodes are on a network that the tool server's management interface (the one used for SSH) can also reach, then you could add
  routes to those service IPs via the management interface instead of through the VLANs. That would be the simplest fix without needing a third NIC.




  I confirm everything I said is correct:

  1. ✅ INTERFACE_SWITCH1/2 and NIC1_NAME/NIC2_NAME are the same physical NICs — one is the kernel name, the other is used by TRex
  2. ✅ All VLANs are built on those physical NICs (or their TAP equivalents)
  3. ✅ All routes (signaling, media, om_cn, ran, sgi) go through those VLANs
  4. ✅ iptables NAT for Prometheus/Search Engine uses signaling VLAN IP as source — needs the VLAN to exist
  5. ✅ Yang/Netconf SSH tunnel goes through the tool server which needs routes via those VLANs
  6. ✅ Outline binds IPs to those VLANs
  7. ✅ When TRex unbinds NICs from kernel → all VLANs destroyed → all routes/iptables/outline gone → everything unreachable
  One correction: I said NIC1_NAME_TAP / NIC2_NAME_TAP are "TAP interfaces created by TRex" — this is what the script assumes (it checks if they're up), but you
  confirmed that you cannot use TAP interfaces for dallas routes because TRex won't work. So that path in the script is not viable for your case.
