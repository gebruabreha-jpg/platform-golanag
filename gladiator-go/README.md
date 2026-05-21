# MASQUE Client

This project develops a client based on the [HTTP/3] protocol, which can establish connections to multiple target servers either directly or indirectly via a [MASQUE] proxy. When work with proxy, the client can operate in two modes: the [MASQUE] tunnel mode and the [QUIC-Aware Forwarding][QUIC-Aware] mode, the latter being an extension of the [MASQUE] protocol. It can also support communicating with the PVD server.

Additionally, it offers the capability to simulate multiple users with distinct IPv4 or IPv6 addresses.

## System view

@startuml
'allowmixing

node "Masque Client(+)" as client {

component "Inner Layer(+)" as inn_layer{
component Inner [
    Http3
    ----
    QUIC
]

}

component "Outer Layer(?)" as Outer_layer{
component Outer [
    Http3 (masque)
    ----
    QUIC [datagram,...]
]
}

component "PVD Client(?)" as pvd_client {
    component Pvd [
    Http3
    ----
    QUIC
]
    
}

component "UDP Layer(+)" as udp_layer {
    component udp [
        UDP
    ]
}

component "client" as masque_client {
control " " as controller
}

Inner <--> Outer : Tunnel Mode
Outer <--> udp 
Pvd <--> udp
Inner <--> udp : Quic-Aware Forwarding Mode

controller -(0 inn_layer: ""
controller -(0 Outer_layer: ""
controller --(0 pvd_client: ""
controller -(0 udp_layer: ""

}

node "MASQUE Proxy" as proxy {
  component p1  [
    Http3 (masque)
    -----
    QUIC [datagram, stream,...]
    -----
    UDP
    ====
    IP (proxy IP)
   ]
}

node "PVD Server" as pvd_server {
    component p3 [
    Http3
    -----
    QUIC [stream,...]
    -----
    UDP
    ====
    IP
    ]    
}

node "Target Server" as target {
    component p2 [
    Http3
    -----
    QUIC [stream,...]
    -----
    UDP
    ====
    IP
    ]    
}

udp <.> proxy : "           1010010001        "
udp <..> pvd_server : " 1101001001"

p1 <.> p2 : "    0010100010      "
@enduml

![system view](https://plantuml.lmera.ericsson.se/svg/jL9HIyCm47xFhpZwuaGxMdqREXbJiU2mOywNCIIiOHIwjARft1Z-TtD9itJhK7mm57gvktpttTrtHhCi_robcvGl2U5vmi0RqVAZOZ3CKiPbv-BS0rh2GjzWHyWYNnKvnuiNSSu4FDCj4pOlvVmzGyUkaMZoPWJKUPAokDQMhuusnjDeIEQ4V1s0lrHo3tvdmBMh3Myo6_3GMUaUpjjrO9PGIPU2hYwxtUvSbJsi6h_TsyxheiKwUTztZp0kErudltQ5vb6DwNryEBGgAKnNhGegKbV0envg-vYqeQUXJEEEy2BdKkGPU63MQAuP4mXXzdBZ-_rwXZruhPGlWv5I0Z4-7KOU1Dkrj-PR3FPWNAKBV_191OF7NAX_eZHbW8Wh2tudS2BfWUST1njBReKnxlPnwlL-b3g8JTQY7ap6qmU8HRxPchMYfUVj1bc4yBECJke8IYcOrTDXfbtOhJhQY6BecCfHh67dZY9ARyB4ceczxafzRQb2KJibkoQxFt4oeh9ADMnUgLWosIGaZQT5vke_oQ2sRhhzUbrggU1E68G1Va5eFLwTZm_S-F1LQ9D3fGgrPurvPG4nee626kSR)

## Design Details

Please refer to [Design Document](docs/design.md)

## Build

### Get source code and build Docker images

Clone the beets and checkout the change

- `git clone ssh://gerrit.ericsson.se:29418/pc/beets`

- checkout the latest version of gladiator

- run `setup_workspace -d`

- run `cd src/main/resources/com/ericsson/pc/beets/files/gladiator`

Then, execute the following command:

```bash
docker build --target=server_stage -t masque_test_server .
docker build --target=client_stage -t masque_test_client .
```

For the first time, building the Docker images may take a long time.

Then, execute these commands to actually create the image files:

```bash
docker save -o masque_test_client.tar masque_test_client
docker save -o masque_test_server.tar masque_test_server
```


### Create and configure the client and server container

#### Create containers

1. create a container as http3 server (for test purpose) and start it

```bash
docker run --name http3_server --hostname server -t -d masque_test_server
```

2. create a container as http3 client and start it

```bash
docker run --name http3_client --hostname client --sysctl net.ipv6.ip_nonlocal_bind=1 -t -d masque_test_client
```

#### Configuration (for test purpose)

1. configure the route on client

```bash
docker exec --privileged http3_client ip r add local 192.168.0.0/16 table local dev lo
docker exec --privileged http3_client ip -6 r add local fd80::/16 table local dev lo
```

2. configure local dns on client

```bash
docker exec --privileged http3_client sh -c "echo `docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' http3_server` www.test.com >> /etc/hosts"
docker exec --privileged http3_client sh -c "echo `docker inspect -f '{{range .NetworkSettings.Networks}}{{.GlobalIPv6Address}}{{end}}' http3_server` www.test6.com >> /etc/hosts"
docker exec --privileged http3_client sh -c "echo  216.155.158.183 quic.rocks >> /etc/hosts"
docker exec --privileged http3_client sh -c "echo  45.77.96.66     quic.tech >> /etc/hosts"
```

3. configure the route on server

```bash
docker exec --privileged http3_server ip r add 192.168.0.0/16 via `docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' http3_client`

docker exec --privileged http3_server ip -6 r add fd80::1/16 via `docker inspect -f '{{range .NetworkSettings.Networks}}{{.GlobalIPv6Address}}{{end}}' http3_client`
```

`Note`: you can run script `config_containers.sh` to execute above commands.

## Test

### Visit the official http3 website

`Note`: It maybe fail if the firewall do not allow traffic on UDP port 8443 or 4433.

For example, visit the <https://quic.rocks:4433> and <https://quic.tech:4433>, and save the output to folder `/tmp`.

```bash
docker exec http3_client /usr/app/bin/http3client https://quic.rocks:4433 https://quic.tech:4433 -o /tmp

docker exec http3_client ls /tmp
quic.rocks:4433_1_0_index.html  quic.tech:4433_1_0_index.html
```

**Note**:

1. The file name has the format: `{hostname}:{port}_{client_id}_{stream_id}_{file_name}`
2. option `--quiet` or `-q` can disables the debug output.

The client can send multiple requests to multiple target servers, allowing for the concurrent download of resources via multiple [QUIC][QUIC] streams within a single connection.

For example, visit the <https://quic.tech:8443/> to get the `index.html` and the picture `quic.svg`.

```bash
docker exec http3_client /usr/app/bin/http3client https://quic.tech:4433 https://quic.tech:4433/quic.svg https://quic.rocks:4433 -o /tmp

docker exec http3_client ls -l /tmp
total 16
-rw-r--r-- 1 borisw borisw  164 Jul 15 11:56 quic.rocks:4433_1_0_index.html
-rw-r--r-- 1 borisw borisw  462 Jul 15 11:56 quic.tech:4433_1_0_index.html
-rw-r--r-- 1 borisw borisw 5126 Jul 15 11:56 quic.tech:4433_1_4_quic.svg
```

### Visit target server via MASQUE Proxy server

The `--connect-udp` option specifies the hostname/IP address and port of the [MASQUE][MASQUE] Proxy. By default, the client operates in [MASQUE] tunnel mode.

The `--quic-forwarding` option enables the client to operate in [QUIC-Aware Forwarding][QUIC-Aware] mode.

#### Configure and start the MASQUE Proxy container

Besides the docker containers `http3_client` and `http3_server` mentioned above, the docker container `masque_proxy` built from project `aioquic-masque` will be used.

After you get the `masque_proxy` Docker image (e.g. masque_proxy_image.tar), load it and create/start the `masque_proxy` container.

```bash
docker load -i masque_proxy_image.tar
docker run --name masque_proxy --hostname proxy -t -d masque_proxy
```

Add route and local DNS

```bash
docker exec --privileged masque_proxy ip r add 192.168.0.0/16 via `docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' http3_client`
docker exec --privileged masque_proxy ip -6 r add fd80::1/16 via `docker inspect -f '{{range .NetworkSettings.Networks}}{{.GlobalIPv6Address}}{{end}}' http3_client`

docker exec --privileged masque_proxy sh -c "echo `docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' http3_server` www.test.com >> /etc/hosts"
docker exec --privileged masque_proxy sh -c "echo `docker inspect -f '{{range .NetworkSettings.Networks}}{{.GlobalIPv6Address}}{{end}}' http3_server` www.test6.com >> /etc/hosts"

docker exec --privileged http3_client sh -c "echo `docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' masque_proxy` www.masque-proxy.com >> /etc/hosts"
```

Start the proxy: `python3 masque_examples/masque_server.py --port 6060 --certificate tests/ssl_cert.pem --private-key tests/ssl_key.pem  --host 0.0.0.0` on the `masque_proxy`.

```bash
 docker exec masque_proxy python3 masque_examples/masque_server.py --port 6060 --certificate tests/ssl_cert.pem --private-key tests/ssl_key.pem  --host 0.0.0.0
```

#### Example

visit the `https://quic.tech:4433` via `masque_proxy`

Run below command:

```bash

# Add local DNS resolve for quic.rocks and quic.tech
docker exec --privileged masque_proxy sh -c "echo  216.155.158.183 quic.rocks >> /etc/hosts"
docker exec --privileged masque_proxy sh -c "echo  45.77.96.66  quic.tech >> /etc/hosts"

# clear folder /tmp
docker exec http3_client bash -c 'rm -rf /tmp/*'

# visit quic.tech via masque_proxy
docker exec http3_client /usr/app/bin/http3client --connect-udp www.masque-proxy.com:6060 https://quic.tech:4433 https://quic.tech:4433/quic.svg  -o /tmp -q

# check the download files
docker exec http3_client ls -l /tmp
```

#### Get MASQUE Proxy information from the PVD server

Instead of geting MASQUE Proxy address from option `--connect-udp`, the client can get the MASQUE proxy information from the PVD server.

If the option `--pvd-server <host:port>` is specified, the client will send the `HTTP GET` to the PVD server and fetch the information of the MASQUE proxy. The client will parse the information from the HTTP response's body and visit the MASQUE proxy according to the information provided by the PVD server.

e.g.

`docker exec http3_client /usr/app/bin/http3client --pvd-server www.test.com:443 https://www.test.com/README.md -o /tmp -q`

If the option `--pvd-only` is specified, the client will stop process after get the HTTP response from the PVD server.

e.g.

`docker exec http3_client /usr/app/bin/http3client --pvd-server www.test.com:443 --pvd-only -q`

It can test the entire process or just test the PVD server.

>NOTE: A PVD server simulator has been built with the http3 server.

### Simulate Multiple Users

**Note**:

On the target server, there are some pre-installed files for download.

- `README.md`
- `16K`,`128K`
- `2M`,`16M`

#### Case 1: IPv4

```bash
# 2000 users download files. The connect setup speed is 200/s.
docker exec http3_client /usr/bin/app/http3client https://www.test.com/README.md https://www.test.com/16K https://www.test.com/128K -o /tmp  --clients-num 2000 --connect-speed=200 --ipv4-pool=192.168.0.0/16 -q

# check the files downloaded.
docker exec http3_client ls -lh /tmp
```

This command creates 2000 clients to visit the site `www.test.com` and each client will create 3 [QUIC][QUIC] streams to download the files `README.md` , `16K` and  `128K`.
The downloaded files are stored under folder `/tmp`. Each client has different IPv4 address which is allocated from the pool specified by `--ipv4-pool`.

#### Case 2: IPv6

```bash
docker exec http3_client bash -c 'rm -rf /tmp/*'

# 1000 users download files. The connect setup speed is 50/s.
docker exec http3_client /usr/app/bin/http3client https://www.test6.com/README.md https://www.test6.com/2M -o /tmp  --clients-num 1000 --connect-speed=50 --ipv6-pool=fd80::1/16  -q

docker exec http3_client ls -lh /tmp
```

This command creates 2000 clients to visit the site `www.test6.com` and each client will create 2 [QUIC][QUIC] streams to download the `README.md` and file `2M`.
The downloaded files are stored under folder `/tmp`. Each client has different IPv6 address which is allocated from the pool specified by `--ipv6-pool`.

The output can show the download speed, opened and closed connections/streams.

![output](docs/interface.png)

#### Case 3: IPv4 with MASQUE Proxy

```bash
docker exec http3_client /usr/app/bin/http3client https://www.test.com/README.md https://www.test.com/128K -o /tmp  --clients-num 2000 --connect-speed=10 --ipv4-pool=192.168.0.0/16 -q --connect-udp www.masque-proxy.com:6060

docker exec http3_client ls -lh /tmp
```

The output can show the download speed, opened and closed connections of inner and tunnel layers.

![output](docs/UI_tunnel.png)

### Some concepts and tips

#### User
Different user has different source IP address.  The option --clients-num and option --ipv4-pool / ipv6-pool control how many users the gladiator will create.

e.g. Here are some examples:

- case 1: `--clients-num 100`
    create 100 clients, but all clients have the same source IP address(system assigned by default).  In this case, there is only one User!  (one session)


- case 2: `--clients-num 100 --ipv4-pool  172.16.0.1/16`
    create 100 clients,  each client has the different source IP address allocated from the pool. In this case, there are 100 different Users!

#### Connection

A connection is 4-tuple  (source ip, source port, dest ip, dest port) entity.

Examples:
- case 1:
`--clients-num 100  <request to site A>`
    One user, 100 connections, 100 streams. (same source ip, different source port)

- case 2:
`--clients-num 100  <request to site A>   --request-dup-factor 20`
    One user, 100 connections, 2000 streams.  (each connection has 20 streams)

- case 3:
`--clients-num 100  <request to site A>  <request to site B>  --request-dup-factor 20`
    One user, 200 connections, 4000 streams.  (each connection has 20 streams)

- case 4:
`--clients-num 100 --ipv4-pool  172.16.0.1/16 <request to site A>  <request to site B>  --request-dup-factor 20`
    100 users, 200 connections, 4000 streams.  (each user has 2 connections (different sites), and each connection has 20 streams (20 requests))

The above connections refer to the connections between user and target server.   (please check the CLI output INNER CONNECTIONS)

The connections between gladiator client and proxy is only impacted by the `--clients-num`  (please check the CLI output OUTER CONNECTIONS, one outer layer connection only has one stream for data transmission)

Finally, here is an example to inject traffic with multiple Users (10), multiple connections (each user has 2,  total 20), multiple steams (each has 2, total 40 streams).

![output](docs/user_connect_stream.png)

```bash
user 1:

    connection 1 to 172.17.0.1:4433

      two streams for same request  https://172.17.0.1:4433/seconds/5

    connection 2 to 172.17.0.1:8443

      two streams for same request https://172.17.0.1:8443/seconds/5

...
user 10:

    connection 1 to 172.17.0.1:4433

      two streams for same request  https://172.17.0.1:4433/seconds/5

    connection 2 to 172.17.0.1:8443

      two streams for same request https://172.17.0.1:8443/seconds/5
```


### Capture packets by tcpdump

If you want to save the secrets into log file (for wireshark pcap decryption), add `env SSLKEYLOGFILE=/tmp/ngtcp2.keys` before the command.
For example

```bash

# clear tmp folder
docker exec http3_client bash -c 'rm -rf /tmp/*'
docker exec http3_server bash -c 'rm -rf /tmp/*'

# capture packets on http3_client
docker exec http3_client tcpdump -i eth0 udp port 6060 -w /tmp/test01.pcap

# capture packets on http3_server
docker exec http3_server tcpdump -i eth0 udp port 443 -w /tmp/test02.pcap

# start http3 client to visit web server via masque proxy
docker exec http3_client env SSLKEYLOGFILE=/tmp/ngtcp2.keys /usr/app/bin/http3client https://www.test.com/README.md -o /tmp  --clients-num 2 --ipv4-pool=192.168.0.0/16 -q --connect-udp www.masque-proxy.com:6060

# stop the tcpdump
docker exec http3_client bash -c 'pkill tcpdump'
docker exec http3_server bash -c 'pkill tcpdump'

# cp the test01.pcap to host
docker cp http3_client:/tmp/test01.pcap .
docker cp http3_server:/tmp/test02.pcap .

# cp the key file to host
docker cp http3_client:/tmp/ngtcp2.keys .

# open the pcap files with wireshark and set the ngtcp2.keys as TLS master secrete log file
```

## CLIENT_WRAPPER



[MASQUE]: https://www.rfc-editor.org/rfc/rfc9298

[QUIC]: https://www.rfc-editor.org/rfc/rfc9000


[HTTP/3]: https://www.rfc-editor.org/rfc/rfc9114


[QUIC-Aware]: https://www.ietf.org/archive/id/draft-ietf-masque-quic-proxy-04.html
