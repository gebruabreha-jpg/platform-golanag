#!/bin/bash

MY_IP=`ip add show | grep eth0 | grep inet | awk '{print $2}' | cut -d/ -f1`
LAST_OCTET=`echo $MY_IP | cut -d. -f4`
echo "nameserver ${MY_IP}" > /etc/resolv.conf
echo "address=/ericsson.com/ericsson.se/coyote-ericsson.com/${MY_IP}" >> /etc/dnsmasq.conf
