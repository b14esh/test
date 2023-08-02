#!/bin/bash

i=30 # time sleep 
a="swich100"
b="switch101"

#pp(){
#  echo "For switch ${a} off port 2"
#  echo "witing ${i}"
#  echo "Press [CTRL+C] to stop.."
#       sleep ${i}
#}


while :
do
  snmpset -v 2c -c YOU_COMMUNITY  192.168.0.100 ifAdminStatus.49154 i 2
  snmpset -v 2c -c YOU_COMMUNITY  192.168.0.100 ifAdminStatus.49155 i 2
  echo "Press [CTRL+C] to stop.."
  echo "For ${a} off port 2"
  echo "Wait ${i}"
  sleep ${i}
  snmpset -v 2c -c YOU_COMMUNITY  192.168.0.100 ifAdminStatus.49154 i 1
  snmpset -v 2c -c YOU_COMMUNITY  192.168.0.100 ifAdminStatus.49155 i 1
  echo "interface up for ${a}"
  sleep ${i}

  snmpset -v 2c -c YOU_COMMUNITY  192.168.0.101 ifAdminStatus.49154 i 2
  snmpset -v 2c -c YOU_COMMUNITY  192.168.0.101 ifAdminStatus.49155 i 2
  echo "Press [CTRL+C] to stop.."
  echo "For  ${b} off port 2"
  echo "Wait ${i}"
  sleep ${i}
  snmpset -v 2c -c YOU_COMMUNITY  192.168.0.101 ifAdminStatus.49154 i 1
  snmpset -v 2c -c YOU_COMMUNITY  192.168.0.101 ifAdminStatus.49155 i 1
  echo "interface up for ${b}"
  sleep ${i}

done
