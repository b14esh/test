#!/bin/bash

list_net_ip="192.168.55 192.168.80 192.168.88 10.12.88"
ip_router="192.168.16.55"
i=99

for x in ${list_net_ip} ; do
      i=$(ip route | grep "${x}" 1>/dev/null ; echo $?)
      if ((i > 0)) ; then 
       #echo -e "ip route add ${x} via ${ip_router}"
       ip route add ${x}.0/24 via ${ip_router}
      fi
done
~         
