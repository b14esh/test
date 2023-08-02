#!/bin/sh
disk_name=$1
if [ "$1" == "" ]
then
         echo -e "ВНИМАНИЕ !!!"
         echo -e "\t Убедись что это не твойcbcntvysq диск!!!"
         echo -e "Передайте имя диска аргументом при запусе. \nПример:\n\t $0 sdb"
         echo -e "Доступные диски для записи: \n\t\t`ls /dev/sd*`"       
         exit
fi

# 0
# Если аргумент был передан, то пошли дальше  
# Для начала очистим диск
echo "0. Стираем начало диска ${disk_name}"
dd if=/dev/zero of=/dev/${disk_name} bs=8M count=1

