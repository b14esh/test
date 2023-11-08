#!/bin/bash
#set -x #DEBUG ON

#sed  -i 's/\#NTP=/NTP=0.ru.pool.ntp.org1.ru.pool.ntp.org/g' /etc/systemd/timesyncd.conf
#cp /usr/share/zoneinfo/Europe/Moscow /etc/localtime
#systemctl restart  systemd-timesyncd.service


hosts=" root@192.168.55.112 root@192.168.16.106 user@192.168.16.115"
#commands="LANG="en_US.UTF-8" LC_TIME="en_US.utf8"  date"
commands="LANG="en_US.UTF-8" LC_TIME="C"  date"

for i in ${hosts}; do
        echo ${i} $(ssh ${i} ${commands} 2> /dev/null)
done

