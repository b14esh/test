#!/bin/bash

#flopiflop_v4
# Выключает порт и через время указанное в переменной "time_to_switch" его включает
# пока не закончатся порты в "list_port"
# после берется следующий из списка "list_ip"
# Все повторяется

#perem
count=1
time_to_switch=5
list_ip="192.168.88.100 192.168.88.101"
list_name="ip_switch_1 ip_switch2"
list_port="49154 49155"
community="YOU_COMMUNITY"

(
trap printout SIGINT
printout() {
    echo -e "\t \t At end of loop: count=$count"
    exit
}

floap_switch(){
for i in ${list_name}; do
        for x in ${list_ip}; do
          for s in ${list_port}; do
                  snmpset -v 2c -c ${community} ${x} ifAdminStatus.${s} i 2
                  echo "For switch ${i} port ${s} off"
                  echo "Wait ${time_to_switch} sec"
                  sleep ${time_to_switch}
                  echo "For switch ${i} port ${s} on"
                  snmpset -v 2c -c ${community} ${x} ifAdminStatus.${s} i 1
                  echo "Press [CTRL+C] to stop.."
          done
  done
done

}


#for start loop
echo -e "\t \t Start script  $0"



while :
do
  floap_switch
  ((count++))
done 
)

#for end
echo -e "\t \t Return all ports to UP"
for i in ${list_ip}; do
        for x in ${list_port}; do
                snmpset -v 2c -c ${community} ${i} ifAdminStatus.${x} i 1
        done
done


