#!/bin/bash

MY_IP=`ip add show | grep eth0 | grep inet | awk '{print $2}' | cut -d/ -f1`
SECOND_OCTET=`echo $MY_IP | cut -d. -f2`
case "$SECOND_OCTET" in
  38)
    LAST_OCTET=1
    ;;
  39)
    LAST_OCTET=2
    ;;
  *)
    LAST_OCTET=$(( $SECOND_OCTET + 3 ))
    ;;
esac
echo "nameserver 192.0.0.$LAST_OCTET"  > /etc/resolv.conf
