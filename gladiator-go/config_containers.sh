#! /usr/bin/bash
echo restart http3_client and http3_server
docker restart http3_client http3_server
echo configure client
docker exec --privileged http3_client ip r add local 192.168.0.0/16 table local dev lo
docker exec --privileged http3_client ip -6 r add local fd80::/16 table local dev lo
docker exec --privileged http3_client sh -c "echo `docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' http3_server` www.test.com >> /etc/hosts"   
docker exec --privileged http3_client sh -c "echo `docker inspect -f '{{range .NetworkSettings.Networks}}{{.GlobalIPv6Address}}{{end}}' http3_server` www.test6.com >> /etc/hosts" 
docker exec --privileged http3_client sh -c "echo  216.155.158.183 quic.rocks >> /etc/hosts" 
docker exec --privileged http3_client sh -c "echo  45.77.96.66     quic.tech >> /etc/hosts" 
echo configure server
docker exec --privileged http3_server ip r add 192.168.0.0/16 via `docker inspect -f '{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' http3_client`
docker exec --privileged http3_server ip -6 r add fd80::1/16 via `docker inspect -f '{{range .NetworkSettings.Networks}}{{.GlobalIPv6Address}}{{end}}' http3_client`
